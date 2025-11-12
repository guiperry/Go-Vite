package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

type ProjectConfig struct {
	Name        string
	Module      string
	Description string
	Author      string
	Port        int
	BackendPort int
}

type ProjectType int

const (
	Unknown ProjectType = iota
	GoProject
	NodeProject
)

type CLIData struct {
	InstalledModules map[string][]string `json:"installed_modules"` // project path -> modules
}

var rootCmd = &cobra.Command{
	Use:   "go-vite",
	Short: "Go-Vite Framework Generator",
	Long:  `A CLI tool to generate Go + Vite + React projects with embedded webview support.`,
}

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Go-Vite project",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-vite v%s\n", version)
	},
}

var installCmd = &cobra.Command{
	Use:   "install [module]",
	Short: "Install a module",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstall,
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [module]",
	Short: "Uninstall a module",
	Args:  cobra.ExactArgs(1),
	RunE:  runUninstall,
}

var installLocalCmd = &cobra.Command{
	Use:   "install-local [path]",
	Short: "Install a module from local directory",
	Args:  cobra.ExactArgs(1),
	RunE:  runInstallLocal,
}

var importModuleCmd = &cobra.Command{
	Use:   "import-module [path]",
	Short: "Import and register a local module",
	Args:  cobra.ExactArgs(1),
	RunE:  runImportModule,
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(installLocalCmd)
	rootCmd.AddCommand(importModuleCmd)

	initCmd.Flags().StringP("module", "m", "", "Go module name (e.g., github.com/user/project)")
	initCmd.Flags().StringP("description", "d", "A Go-Vite desktop application", "Project description")
	initCmd.Flags().StringP("author", "a", "", "Author name")
	initCmd.Flags().IntP("port", "p", 5173, "Frontend port")
	initCmd.Flags().IntP("backend-port", "b", 8080, "Backend port")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runInit(cmd *cobra.Command, args []string) error {
	// Get project name
	projectName := "my-app"
	if len(args) > 0 {
		projectName = args[0]
	}

	// Get current directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	projectPath := filepath.Join(cwd, projectName)

	// Check if directory exists
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", projectName)
	}

	// Get flags
	moduleName, _ := cmd.Flags().GetString("module")
	if moduleName == "" {
		moduleName = projectName
	}

	description, _ := cmd.Flags().GetString("description")
	author, _ := cmd.Flags().GetString("author")
	port, _ := cmd.Flags().GetInt("port")
	backendPort, _ := cmd.Flags().GetInt("backend-port")

	config := ProjectConfig{
		Name:        projectName,
		Module:      moduleName,
		Description: description,
		Author:      author,
		Port:        port,
		BackendPort: backendPort,
	}

	fmt.Printf("🚀 Creating new Go-Vite project: %s\n", projectName)
	fmt.Printf("📦 Module: %s\n", moduleName)
	fmt.Printf("🔧 Frontend port: %d\n", port)
	fmt.Printf("🔧 Backend port: %d\n\n", backendPort)

	// Create project structure
	if err := createProjectStructure(projectPath, config); err != nil {
		return fmt.Errorf("failed to create project structure: %w", err)
	}

	fmt.Println("✅ Project created successfully!")
	fmt.Printf("\n📝 Next steps:\n")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Printf("   make deps      # Install dependencies\n")
	fmt.Printf("   make binary    # Build the application\n")
	fmt.Printf("   ./dist/%s  # Run the application\n\n", projectName)

	return nil
}

func runInstall(cmd *cobra.Command, args []string) error {
	installModule(args[0])
	return nil
}

func runUninstall(cmd *cobra.Command, args []string) error {
	uninstallModule(args[0])
	return nil
}

func runInstallLocal(cmd *cobra.Command, args []string) error {
	installLocalModule(args[0])
	return nil
}

func runImportModule(cmd *cobra.Command, args []string) error {
	importModule(args[0])
	return nil
}

func detectProjectType() ProjectType {
	if _, err := os.Stat("go.mod"); err == nil {
		return GoProject
	}
	if _, err := os.Stat("package.json"); err == nil {
		return NodeProject
	}
	return Unknown
}

func installModule(module string) {
	projectType := detectProjectType()
	switch projectType {
	case GoProject:
		installGoModule(module)
	case NodeProject:
		installNodeModule(module)
	default:
		fmt.Println("Unknown project type. Ensure you have go.mod or package.json in the current directory.")
		os.Exit(1)
	}
}

func uninstallModule(module string) {
	projectType := detectProjectType()
	switch projectType {
	case GoProject:
		uninstallGoModule(module)
	case NodeProject:
		uninstallNodeModule(module)
	default:
		fmt.Println("Unknown project type. Ensure you have go.mod or package.json in the current directory.")
		os.Exit(1)
	}
}

func installGoModule(module string) {
	fmt.Printf("Installing Go module: %s\n", module)
	cmd := exec.Command("go", "get", module)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to install Go module: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Go module installed successfully.")
	saveInstalledModule(module)
}

func uninstallGoModule(module string) {
	fmt.Printf("Uninstalling Go module: %s\n", module)
	// For Go, we can try to remove from go.mod
	cmd := exec.Command("go", "mod", "edit", "-droprequire", module)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to remove from go.mod: %v\n", err)
		os.Exit(1)
	}
	// Then tidy
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to tidy go.mod: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Go module uninstalled successfully.")
	removeInstalledModule(module)
}

func installNodeModule(module string) {
	fmt.Printf("Installing Node.js module: %s\n", module)
	cmd := exec.Command("npm", "install", module)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to install Node.js module: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Node.js module installed successfully.")
	saveInstalledModule(module)
}

func uninstallNodeModule(module string) {
	fmt.Printf("Uninstalling Node.js module: %s\n", module)
	cmd := exec.Command("npm", "uninstall", module)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to uninstall Node.js module: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Node.js module uninstalled successfully.")
	removeInstalledModule(module)
}

func getDataFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.TempDir()
	}
	return filepath.Join(configDir, "automationgenie", "cli.json")
}

func loadData() (*CLIData, error) {
	path := getDataFilePath()
	data := &CLIData{
		InstalledModules: make(map[string][]string),
	}
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}

func saveData(data *CLIData) error {
	path := getDataFilePath()
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(data)
}

func saveInstalledModule(module string) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	data, err := loadData()
	if err != nil {
		return
	}
	if data.InstalledModules[wd] == nil {
		data.InstalledModules[wd] = []string{}
	}
	// Check if already exists
	for _, m := range data.InstalledModules[wd] {
		if m == module {
			return
		}
	}
	data.InstalledModules[wd] = append(data.InstalledModules[wd], module)
	saveData(data)
}

func removeInstalledModule(module string) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	data, err := loadData()
	if err != nil {
		return
	}
	modules := data.InstalledModules[wd]
	for i, m := range modules {
		if m == module {
			data.InstalledModules[wd] = append(modules[:i], modules[i+1:]...)
			break
		}
	}
	saveData(data)
}

func installLocalModule(sourcePath string) {
	// Check if source path exists
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		fmt.Printf("Error: Source path '%s' does not exist\n", sourcePath)
		os.Exit(1)
	}

	// Detect module type
	moduleType := detectLocalModuleType(sourcePath)
	if moduleType == Unknown {
		fmt.Println("Error: Cannot determine module type. Ensure the directory contains go.mod or package.json")
		os.Exit(1)
	}

	// Get module name from source
	moduleName := getModuleName(sourcePath, moduleType)
	if moduleName == "" {
		fmt.Println("Error: Cannot determine module name")
		os.Exit(1)
	}

	// Determine destination path
	destPath := getModuleDestinationPath(moduleName, moduleType)

	// Copy module to destination
	fmt.Printf("Installing local module '%s' to %s\n", moduleName, destPath)
	if err := copyDir(sourcePath, destPath); err != nil {
		fmt.Printf("Error copying module: %v\n", err)
		os.Exit(1)
	}

	// Register the module
	if err := registerLocalModule(moduleName, moduleType, destPath); err != nil {
		fmt.Printf("Error registering module: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Local module '%s' installed successfully\n", moduleName)
	saveInstalledModule("local:" + moduleName)
}

func importModule(sourcePath string) {
	// Check if source path exists
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		fmt.Printf("Error: Source path '%s' does not exist\n", sourcePath)
		os.Exit(1)
	}

	// Detect module type
	moduleType := detectLocalModuleType(sourcePath)
	if moduleType == Unknown {
		fmt.Println("Error: Cannot determine module type. Ensure the directory contains go.mod or package.json")
		os.Exit(1)
	}

	// Get module name from source
	moduleName := getModuleName(sourcePath, moduleType)
	if moduleName == "" {
		fmt.Println("Error: Cannot determine module name")
		os.Exit(1)
	}

	// Check if module already exists
	destPath := getModuleDestinationPath(moduleName, moduleType)
	if _, err := os.Stat(destPath); err == nil {
		fmt.Printf("Module '%s' already exists. Use --force to overwrite\n", moduleName)
		os.Exit(1)
	}

	// Copy module to destination
	fmt.Printf("Importing module '%s' to %s\n", moduleName, destPath)
	if err := copyDir(sourcePath, destPath); err != nil {
		fmt.Printf("Error copying module: %v\n", err)
		os.Exit(1)
	}

	// Register the module
	if err := registerLocalModule(moduleName, moduleType, destPath); err != nil {
		fmt.Printf("Error registering module: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Module '%s' imported and registered successfully\n", moduleName)
	saveInstalledModule("imported:" + moduleName)
}

func detectLocalModuleType(sourcePath string) ProjectType {
	goModPath := filepath.Join(sourcePath, "go.mod")
	packageJsonPath := filepath.Join(sourcePath, "package.json")

	if _, err := os.Stat(goModPath); err == nil {
		return GoProject
	}
	if _, err := os.Stat(packageJsonPath); err == nil {
		return NodeProject
	}
	return Unknown
}

func getModuleName(sourcePath string, moduleType ProjectType) string {
	switch moduleType {
	case GoProject:
		return getGoModuleName(sourcePath)
	case NodeProject:
		return getNodeModuleName(sourcePath)
	default:
		return ""
	}
}

func getGoModuleName(sourcePath string) string {
	goModPath := filepath.Join(sourcePath, "go.mod")
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1]
			}
		}
	}
	return ""
}

func getNodeModuleName(sourcePath string) string {
	packageJsonPath := filepath.Join(sourcePath, "package.json")
	content, err := os.ReadFile(packageJsonPath)
	if err != nil {
		return ""
	}

	// Simple JSON parsing to extract name field
	contentStr := string(content)
	nameStart := strings.Index(contentStr, `"name"`)
	if nameStart == -1 {
		return ""
	}

	colonIdx := strings.Index(contentStr[nameStart:], ":")
	if colonIdx == -1 {
		return ""
	}

	valueStart := nameStart + colonIdx + 1
	quoteStart := strings.Index(contentStr[valueStart:], `"`)
	if quoteStart == -1 {
		return ""
	}

	valueStart += quoteStart + 1
	quoteEnd := strings.Index(contentStr[valueStart:], `"`)
	if quoteEnd == -1 {
		return ""
	}

	return contentStr[valueStart : valueStart+quoteEnd]
}

func getModuleDestinationPath(moduleName string, moduleType ProjectType) string {
	// Assume we're running from the AutomationGenie root
	basePath := "backend/internal/modules"

	switch moduleType {
	case GoProject:
		return filepath.Join(basePath, moduleName)
	case NodeProject:
		return filepath.Join(basePath, moduleName)
	default:
		return ""
	}
}

func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func registerLocalModule(moduleName string, moduleType ProjectType, modulePath string) error {
	// For now, just log the registration
	// In a full implementation, this would update the module registry
	fmt.Printf("Registering module: %s (%s) at %s\n", moduleName, moduleTypeString(moduleType), modulePath)

	// TODO: Update backend/internal/modules/builtin.go or modules.go to register the new module
	// This would require parsing and modifying Go source files

	return nil
}

func moduleTypeString(moduleType ProjectType) string {
	switch moduleType {
	case GoProject:
		return "Go"
	case NodeProject:
		return "Node.js"
	default:
		return "Unknown"
	}
}

func createProjectStructure(projectPath string, config ProjectConfig) error {
	// Create directories
	dirs := []string{
		"backend/cmd/server",
		"backend/config",
		"backend/internal/api/handlers",
		"backend/internal/api/middleware",
		"backend/internal/models",
		"backend/internal/modules",
		"backend/internal/storage",
		"backend/internal/utils",
		"backend/tests",
		"frontend/src/components",
		"frontend/src/pages",
		"frontend/src/hooks",
		"frontend/src/services",
		"frontend/src/utils",
		"frontend/public",
		"netlify/functions",
		"dist",
		"bin",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectPath, dir), 0755); err != nil {
			return err
		}
	}

	// Generate files from templates
	files := map[string]string{
		// Root files
		"go.mod":          generateGoMod(config),
		"main.go":         generateMainGo(),
		"Makefile":        generateMakefile(config),
		"README.md":       generateReadme(config),
		".gitignore":      generateGitignore(),
		".gitattributes":  generateGitattributes(),
		".env.example":    generateEnvExample(config),
		"netlify.toml":    generateNetlifyToml(config),

		// Backend files
		"backend/go.mod":                          generateBackendGoMod(),
		"backend/cmd/server/main.go":              generateBackendMain(),
		"backend/config/config.go":                generateConfig(),
		"backend/internal/api/routes.go":          generateRoutes(),
		"backend/internal/api/handlers/handlers.go": generateHandlers(),
		"backend/internal/api/middleware/cors.go":   generateCorsMiddleware(),
		"backend/internal/api/middleware/logger.go": generateLoggerMiddleware(),
		"backend/internal/models/pipeline.go":       generatePipelineModel(),
		"backend/internal/models/project.go":        generateProjectModel(),
		"backend/internal/models/user.go":           generateUserModel(),
		"backend/internal/modules/modules.go":       generateModulesManager(),
		"backend/internal/modules/builtin.go":       generateBuiltinModules(),
		"backend/internal/storage/database.go":      generateDatabase(),
		"backend/internal/utils/logger.go":          generateLogger(),

		// Frontend files
		"frontend/package.json":       generatePackageJson(config),
		"frontend/vite.config.js":     generateViteConfig(config),
		"frontend/tailwind.config.js": generateTailwindConfig(),
		"frontend/postcss.config.js":  generatePostcssConfig(),
		"frontend/index.html":         generateIndexHtml(config),
		"frontend/src/main.tsx":       generateMainTsx(),
		"frontend/src/App.tsx":        generateAppTsx(config),
		"frontend/src/index.css":      generateIndexCss(),
		"frontend/.eslintrc.cjs":      generateEslintrc(),
		"frontend/.prettierrc":        generatePrettierrc(),

		// Netlify files
		"netlify/functions/api.js":    generateNetlifyApiFunction(),
	}

	for path, content := range files {
		fullPath := filepath.Join(projectPath, path)
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", path, err)
		}
	}

	fmt.Println("✅ Project structure created")
	return nil
}

// Template generators
func generateGoMod(config ProjectConfig) string {
	return fmt.Sprintf(`module %s

go 1.24.0

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
	github.com/webview/webview_go v0.0.0-20240831120633-6173450d4dd6
)

replace backend => ./backend
`, config.Module)
}

func generateMainGo() string {
	return `package main

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
`
}

func generateMakefile(config ProjectConfig) string {
	return fmt.Sprintf(`# %s Build System
PROJECT_NAME := %s
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +"%%Y-%%m-%%dT%%H:%%M:%%SZ")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

ROOT_DIR := $(shell pwd)
BACKEND_DIR := $(ROOT_DIR)/backend
FRONTEND_DIR := $(ROOT_DIR)/frontend
BUILD_DIR := $(ROOT_DIR)/build
DIST_DIR := $(ROOT_DIR)/dist

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
CGO_ENABLED := 0
CGO_ENABLED_WEBVIEW := 1

LDFLAGS := -X main.Version=$(VERSION) \
           -X main.BuildTime=$(BUILD_TIME) \
           -X main.GitCommit=$(GIT_COMMIT) \
           -w -s

NPM_CMD := npm

.PHONY: help all clean build test deps frontend backend binary

all: clean deps frontend backend binary

help:
	@echo "Available targets:"
	@echo "  all       - Build everything"
	@echo "  deps      - Install dependencies"
	@echo "  frontend  - Build React frontend"
	@echo "  backend   - Build Go backend"
	@echo "  binary    - Build unified binary"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"

deps: deps-go deps-node

deps-go:
	@echo "Installing Go dependencies..."
	cd $(BACKEND_DIR) && go mod download && go mod tidy
	go mod download

deps-node:
	@echo "Installing Node.js dependencies..."
	cd $(FRONTEND_DIR) && $(NPM_CMD) install

frontend: deps-node
	@echo "Building React frontend..."
	cd $(FRONTEND_DIR) && $(NPM_CMD) run build
	@echo "Frontend build completed"

backend: deps-go
	@echo "Building Go backend..."
	cd $(BACKEND_DIR) && go build -ldflags "$(LDFLAGS)" -o bin/backend ./cmd/server
	@mkdir -p bin
	cp $(BACKEND_DIR)/bin/backend bin/
	@echo "Backend build completed"

binary: frontend backend
	@echo "Creating unified binary..."
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=$(CGO_ENABLED_WEBVIEW) go build -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/%s .
	@echo "Unified binary created: $(DIST_DIR)/%s"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)
	rm -rf $(BACKEND_DIR)/bin
	rm -rf bin

test:
	@echo "Running tests..."
	cd $(BACKEND_DIR) && go test -v ./...

run: binary
	./$(DIST_DIR)/%s
`, config.Name, config.Name, config.Name, config.Name, config.Name)
}

func generateReadme(config ProjectConfig) string {
	return fmt.Sprintf(`# %s

%s

## 🚀 Quick Start

### Prerequisites
- Go 1.24+
- Node.js 18+
- npm/yarn

### Installation & Running

` + "```" + `bash
# Install dependencies
make deps

# Build the application
make binary

# Run the desktop application
./dist/%s
` + "```" + `

## 📁 Project Structure

` + "```" + `
%s/
├── main.go                    # Desktop app entry point
├── Makefile                   # Build automation
├── backend/                   # Go backend service
│   ├── cmd/server/           # Backend server entry
│   ├── config/               # Configuration
│   ├── internal/
│   │   ├── api/             # API routes and handlers
│   │   ├── models/          # Data models
│   │   ├── modules/         # Business logic modules
│   │   ├── storage/         # Database and cache
│   │   └── utils/           # Utilities
│   └── tests/               # Tests
├── frontend/                  # React frontend
│   ├── src/
│   │   ├── components/      # React components
│   │   ├── pages/           # Page components
│   │   ├── hooks/           # Custom hooks
│   │   ├── services/        # API services
│   │   └── utils/           # Utilities
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
└── dist/                      # Final binary output
` + "```" + `

## 🛠️ Development

` + "```" + `bash
# Run frontend dev server
cd frontend && npm run dev

# Run backend dev server
cd backend && go run ./cmd/server

# Build for production
make all
` + "```" + `

## 📄 License

MIT License
`, config.Name, config.Description, config.Name, config.Name)
}

func generateGitignore() string {
	return `# Dependencies
node_modules/
vendor/

# Build outputs
dist/
build/
*.exe
*.dll
*.so
*.dylib

# Environment files
.env
.env.local
.env.production
.env.development

# Logs
*.log
logs/

# Editor directories
.vscode/
.idea/
*.swp

# Testing
coverage/

# Backend specific
backend/tmp/
backend/bin/
*.test

# Database
*.db
*.sqlite
data/

# OS
.DS_Store
Thumbs.db

bin/*
`
}

func generateGitattributes() string {
	return `* text=auto
`
}

func generateEnvExample(config ProjectConfig) string {
	return fmt.Sprintf(`# Frontend
VITE_API_URL=http://localhost:%d
VITE_APP_NAME=%s

# Backend
PORT=%d
LOG_LEVEL=info
`, config.BackendPort, config.Name, config.BackendPort)
}

func generateNetlifyToml(config ProjectConfig) string {
	return fmt.Sprintf(`[build]
  publish = "frontend/dist"
  functions = "netlify/functions"

[build.environment]
  NODE_VERSION = "18"
  VITE_APP_NAME = "%s"
  VITE_API_URL = "http://localhost:%d"

[[redirects]]
  from = "/api/*"
  to = "/.netlify/functions/api/:splat"
  status = 200

[[redirects]]
  from = "/*"
  to = "/index.html"
  status = 200
`, config.Name, config.BackendPort)
}

func generateNetlifyApiFunction() string {
	return `const axios = require('axios');

exports.handler = async (event, context) => {
  // Set CORS headers
  const headers = {
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Headers': 'Content-Type',
    'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
    'Content-Type': 'application/json'
  };

  // Handle preflight requests
  if (event.httpMethod === 'OPTIONS') {
    return {
      statusCode: 200,
      headers,
      body: ''
    };
  }

  try {
    // Extract the path after /api/
    const path = event.path.replace('/.netlify/functions/api/', '');

    // Forward the request to your backend service
    // In production, replace this with your actual backend URL
    const backendUrl = process.env.BACKEND_URL || 'http://localhost:8080';

    const response = await axios({
      method: event.httpMethod,
      url: backendUrl + '/api/' + path,
      data: event.body,
      headers: {
        'Content-Type': event.headers['content-type'] || 'application/json',
        'Authorization': event.headers.authorization || '',
      },
      params: event.queryStringParameters
    });

    return {
      statusCode: response.status,
      headers,
      body: JSON.stringify(response.data)
    };

  } catch (error) {
    console.error('API proxy error:', error);

    return {
      statusCode: error.response?.status || 500,
      headers,
      body: JSON.stringify({
        error: 'Internal Server Error',
        message: error.message
      })
    };
  }
};
`
}

// Backend generators
func generateBackendGoMod() string {
	return `module backend

go 1.24

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/joho/godotenv v1.5.1
)
`
}

func generateBackendMain() string {
	return `package main

import (
	"log"
	"os"

	"backend/config"
	"backend/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg := config.GetConfig()

	router := gin.Default()
	router.Use(api.CORSMiddleware())
	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Backend server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
`
}

func generateConfig() string {
	return `package config

import (
	"os"
	"strconv"
)

type Config struct {
	FrontendPort int
	BackendPort  int
	LogLevel     string
}

var globalConfig Config

func LoadConfig() error {
	frontendPort, _ := strconv.Atoi(getEnv("FRONTEND_PORT", "5173"))
	backendPort, _ := strconv.Atoi(getEnv("BACKEND_PORT", "8080"))

	globalConfig = Config{
		FrontendPort: frontendPort,
		BackendPort:  backendPort,
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}

	return nil
}

func GetConfig() Config {
	return globalConfig
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
`
}

func generateRoutes() string {
	return `package api

import (
	"backend/internal/api/handlers"
	"backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	v1.Use(middleware.Logger())
	{
		v1.GET("/items", handlers.ListItems)
		v1.POST("/items", handlers.CreateItem)
		v1.GET("/items/:id", handlers.GetItem)
		v1.PUT("/items/:id", handlers.UpdateItem)
		v1.DELETE("/items/:id", handlers.DeleteItem)
	}
}

func CORSMiddleware() gin.HandlerFunc {
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
`
}

func generateHandlers() string {
	return `package handlers

import (
	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {
	c.JSON(200, gin.H{"items": []string{}})
}

func CreateItem(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Item created"})
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id})
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "message": "Item updated"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "message": "Item deleted"})
}
`
}

func generateCorsMiddleware() string {
	return `package middleware

import "github.com/gin-gonic/gin"

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}
`
}

func generateLoggerMiddleware() string {
	return `package middleware

import (
	"log"
	"time"
	
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		
		c.Next()
		
		latency := time.Since(start)
		status := c.Writer.Status()
		
		log.Printf("[%s] %s %d %v", c.Request.Method, path, status, latency)
	}
}
`
}

func generatePipelineModel() string {
	return `package models

import "time"

type Pipeline struct {
	ID            string    ` + "`json:\"id\"`" + `
	Name          string    ` + "`json:\"name\"`" + `
	Status        string    ` + "`json:\"status\"`" + `
	CurrentStep   int       ` + "`json:\"current_step\"`" + `
	CompletedSteps []int    ` + "`json:\"completed_steps\"`" + `
	CreatedAt     time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt     time.Time ` + "`json:\"updated_at\"`" + `
}
`
}

func generateProjectModel() string {
	return `package models

import "time"

type Project struct {
	ID          string    ` + "`json:\"id\"`" + `
	Name        string    ` + "`json:\"name\"`" + `
	Description string    ` + "`json:\"description\"`" + `
	Owner       string    ` + "`json:\"owner\"`" + `
	CreatedAt   time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt   time.Time ` + "`json:\"updated_at\"`" + `
}
`
}

func generateUserModel() string {
	return `package models

import "time"

type User struct {
	ID        string    ` + "`json:\"id\"`" + `
	Email     string    ` + "`json:\"email\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
}
`
}

func generateModulesManager() string {
	return `package modules

type Module interface {
	Name() string
	Execute(input map[string]interface{}) (map[string]interface{}, error)
	Validate(config map[string]interface{}) error
}

type Manager struct {
	modules map[string]Module
}

func NewManager() *Manager {
	return &Manager{
		modules: make(map[string]Module),
	}
}

func (m *Manager) Register(name string, module Module) {
	m.modules[name] = module
}

func (m *Manager) Get(name string) (Module, bool) {
	module, ok := m.modules[name]
	return module, ok
}

func (m *Manager) List() []string {
	names := make([]string, 0, len(m.modules))
	for name := range m.modules {
		names = append(names, name)
	}
	return names
}
`
}

func generateBuiltinModules() string {
	return `package modules

type ExampleModule struct{}

func (m *ExampleModule) Name() string {
	return "example"
}

func (m *ExampleModule) Execute(input map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "success",
		"data":   input,
	}, nil
}

func (m *ExampleModule) Validate(config map[string]interface{}) error {
	return nil
}

func LoadBuiltinModules(manager *Manager) {
	manager.Register("example", &ExampleModule{})
}
`
}

func generateDatabase() string {
	return `package storage

type Database struct {
	// Add your database implementation here
}

func NewDatabase() (*Database, error) {
	return &Database{}, nil
}

func (db *Database) Close() error {
	return nil
}
`
}

func generateLogger() string {
	return `package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile)

func Info(v ...interface{}) {
	Logger.Println(v...)
}

func Error(v ...interface{}) {
	Logger.Println(v...)
}
`
}

// Frontend generators
func generatePackageJson(config ProjectConfig) string {
	return `{
  "name": "` + config.Name + `",
  "version": "1.0.0",
  "description": "` + config.Description + `",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview",
    "lint": "eslint . --ext js,jsx,ts,tsx",
    "format": "prettier --write \"src/**/*.{js,jsx,ts,tsx,json,css,md}\""
  },
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "lucide-react": "^0.294.0",
    "axios": "^1.6.2"
  },
  "devDependencies": {
    "@types/react": "^18.2.43",
    "@types/react-dom": "^18.2.17",
    "@vitejs/plugin-react": "^4.2.1",
    "autoprefixer": "^10.4.16",
    "eslint": "^8.55.0",
    "eslint-plugin-react": "^7.33.2",
    "eslint-plugin-react-hooks": "^4.6.0",
    "eslint-plugin-react-refresh": "^0.4.5",
    "postcss": "^8.4.32",
    "prettier": "^3.1.1",
    "tailwindcss": "^3.3.6",
    "vite": "^5.0.8"
  }
}
`
}

func generateViteConfig(config ProjectConfig) string {
	return `import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    port: ` + fmt.Sprintf("%d", config.Port) + `,
    host: true,
    proxy: {
      '/api': {
        target: 'http://localhost:` + fmt.Sprintf("%d", config.BackendPort) + `',
        changeOrigin: true,
        secure: false,
      },
    },
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
  },
})
`
}

func generateTailwindConfig() string {
	return `/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#ecfeff',
          100: '#cffafe',
          200: '#a5f3fc',
          300: '#67e8f9',
          400: '#22d3ee',
          500: '#06b6d4',
          600: '#0891b2',
          700: '#0e7490',
          800: '#155e75',
          900: '#164e63',
        },
      },
    },
  },
  plugins: [],
}
`
}

func generatePostcssConfig() string {
	return `export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}
`
}

func generateIndexHtml(config ProjectConfig) string {
	return `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>` + config.Name + `</title>
  </head>
  <body>
    <div id="root"></div>
    <script type="module" src="/src/main.tsx"></script>
  </body>
</html>
`
}

func generateMainTsx() string {
	return `import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
`
}

func generateAppTsx(config ProjectConfig) string {
	return `import React, { useState, useEffect } from 'react';
import { Play, Settings } from 'lucide-react';

function App() {
  const [status, setStatus] = useState<string>('idle');
  const [items, setItems] = useState<any[]>([]);

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await fetch('/api/v1/items');
      const data = await response.json();
      setItems(data.items || []);
    } catch (error) {
      console.error('Failed to fetch items:', error);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-900 via-brand-900 to-slate-900 text-white p-6">
      <div className="max-w-7xl mx-auto">
        <div className="text-center mb-8">
          <h1 className="text-5xl font-bold mb-2 bg-gradient-to-r from-brand-400 to-blue-400 bg-clip-text text-transparent">
            ` + config.Name + `
          </h1>
          <p className="text-lg text-gray-300">` + config.Description + `</p>
        </div>

        <div className="bg-slate-800/50 backdrop-blur-sm rounded-xl p-6 mb-8 border border-brand-500/30">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-2xl font-bold">Dashboard</h2>
            <span className="px-3 py-1 bg-green-600/30 text-green-300 rounded-full text-sm">
              {status}
            </span>
          </div>
          
          <div className="flex gap-3">
            <button
              onClick={() => setStatus('running')}
              className="flex items-center gap-2 bg-gradient-to-r from-brand-600 to-blue-600 px-6 py-3 rounded-lg font-semibold hover:from-brand-500 hover:to-blue-500 transition-all"
            >
              <Play className="w-5 h-5" />
              Start
            </button>
            <button className="flex items-center gap-2 bg-slate-700 px-6 py-3 rounded-lg font-semibold hover:bg-slate-600 transition-all">
              <Settings className="w-5 h-5" />
              Settings
            </button>
          </div>
        </div>

        <div className="bg-slate-800/50 backdrop-blur-sm rounded-xl p-6 border border-brand-500/30">
          <h3 className="text-xl font-bold mb-4">Items</h3>
          {items.length === 0 ? (
            <p className="text-gray-400">No items yet</p>
          ) : (
            <ul className="space-y-2">
              {items.map((item, idx) => (
                <li key={idx} className="p-3 bg-slate-700/50 rounded-lg">
                  {JSON.stringify(item)}
                </li>
              ))}
            </ul>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
`
}

func generateIndexCss() string {
	return `@tailwind base;
@tailwind components;
@tailwind utilities;
`
}

func generateEslintrc() string {
	return `module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:react/recommended',
    'plugin:react/jsx-runtime',
    'plugin:react-hooks/recommended',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs'],
  parserOptions: { ecmaVersion: 'latest', sourceType: 'module' },
  settings: { react: { version: '18.2' } },
  plugins: ['react-refresh'],
  rules: {
    'react-refresh/only-export-components': [
      'warn',
      { allowConstantExport: true },
    ],
    'react/prop-types': 'off',
  },
}
`
}

func generatePrettierrc() string {
	return `{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 100,
  "tabWidth": 2,
  "useTabs": false
}
`
}
