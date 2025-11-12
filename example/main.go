package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"backend/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/webview/webview_go"
)

//go:embed all:frontend/dist
var distFS embed.FS

//go:embed bin/backend
var backendBinary []byte

var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

type EmbeddedFS struct {
	files fs.FS
}

func NewEmbeddedFS() (*EmbeddedFS, error) {
	files, err := fs.Sub(distFS, "frontend/dist")
	if err != nil {
		return nil, err
	}
	return &EmbeddedFS{files: files}, nil
}

func (efs *EmbeddedFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}

	file, err := efs.files.Open(path)
	if err != nil {
		if !strings.Contains(path, ".") {
			htmlPath := path + ".html"
			if file, err = efs.files.Open(htmlPath); err != nil {
				if file, err = efs.files.Open("index.html"); err != nil {
					http.NotFound(w, r)
					return
				}
			}
		} else {
			http.NotFound(w, r)
			return
		}
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	}

	http.ServeContent(w, r, stat.Name(), stat.ModTime(), file.(io.ReadSeeker))
}

type App struct {
	config      *config.Config
	router      *gin.Engine
	server      *http.Server
	backendCmd  *exec.Cmd
	backendPath string
	tempDir     string
}

func NewApp(cfg *config.Config) (*App, error) {
	app := &App{
		config: cfg,
		router: gin.New(),
	}

	if err := app.extractBackend(); err != nil {
		return nil, fmt.Errorf("failed to extract backend: %w", err)
	}

	app.router.Use(gin.Logger())
	app.router.Use(gin.Recovery())
	app.router.Use(corsMiddleware())

	if err := app.setupRoutes(); err != nil {
		return nil, fmt.Errorf("failed to setup routes: %w", err)
	}

	return app, nil
}

func (app *App) extractBackend() error {
	tempDir, err := os.MkdirTemp("", "app-*")
	if err != nil {
		return err
	}
	app.tempDir = tempDir

	app.backendPath = filepath.Join(tempDir, "backend")
	file, err := os.Create(app.backendPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(backendBinary); err != nil {
		return err
	}

	return os.Chmod(app.backendPath, 0755)
}

func (app *App) setupRoutes() error {
	embeddedFS, err := NewEmbeddedFS()
	if err != nil {
		return err
	}

	app.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":     "healthy",
			"version":    Version,
			"build_time": BuildTime,
			"timestamp":  time.Now().UTC().Format(time.RFC3339),
		})
	})

	api := app.router.Group("/api")
	{
		api.Any("/*path", func(c *gin.Context) {
			backendURL := fmt.Sprintf("http://localhost:%d%s", app.config.BackendPort, c.Request.URL.Path)
			req, err := http.NewRequest(c.Request.Method, backendURL, c.Request.Body)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to create proxy request"})
				return
			}

			for key, values := range c.Request.Header {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}

			client := &http.Client{Timeout: 30 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				c.JSON(500, gin.H{"error": "Backend service unavailable"})
				return
			}
			defer resp.Body.Close()

			for key, values := range resp.Header {
				for _, value := range values {
					c.Header(key, value)
				}
			}

			c.Status(resp.StatusCode)
			io.Copy(c.Writer, resp.Body)
		})
	}

	app.router.NoRoute(func(c *gin.Context) {
		embeddedFS.ServeHTTP(c.Writer, c.Request)
	})

	return nil
}

func (app *App) startBackend() error {
	log.Printf("Starting backend service on port %d...", app.config.BackendPort)

	app.backendCmd = exec.Command(app.backendPath)
	env := append(os.Environ(),
		fmt.Sprintf("PORT=%d", app.config.BackendPort),
		"GIN_MODE=release",
	)
	app.backendCmd.Env = env
	app.backendCmd.Stdout = os.Stdout
	app.backendCmd.Stderr = os.Stderr

	if err := app.backendCmd.Start(); err != nil {
		return err
	}

	log.Printf("Backend started (PID: %d)", app.backendCmd.Process.Pid)
	time.Sleep(3 * time.Second)
	return nil
}

func (app *App) stopBackend() {
	if app.backendCmd != nil && app.backendCmd.Process != nil {
		log.Printf("Stopping backend (PID: %d)", app.backendCmd.Process.Pid)
		app.backendCmd.Process.Signal(syscall.SIGTERM)
		app.backendCmd.Wait()
	}
}

func (app *App) Start() error {
	if err := app.startBackend(); err != nil {
		return err
	}

	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.FrontendPort),
		Handler:      app.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Starting application on port %d", app.config.FrontendPort)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return nil
}

func (app *App) Stop() error {
	if app.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := app.server.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown error: %v", err)
		}
	}

	app.stopBackend()

	if app.tempDir != "" {
		os.RemoveAll(app.tempDir)
	}

	return nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	fmt.Printf("Application v%s (built %s, commit %s)\n", Version, BuildTime, GitCommit)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg := config.GetConfig()

	app, err := NewApp(&cfg)
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	url := fmt.Sprintf("http://localhost:%d", cfg.FrontendPort)
	log.Printf("Opening webview at %s", url)

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Application")
	w.SetSize(1200, 800, webview.HintNone)
	w.Navigate(url)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal, terminating webview...")
		w.Terminate()
	}()

	w.Run()

	log.Println("Shutting down...")
	if err := app.Stop(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}

	log.Println("Application stopped")
}
