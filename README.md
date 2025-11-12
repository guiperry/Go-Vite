# Go-Vite CLI Tool

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Go](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)
![React](https://img.shields.io/badge/React-18.x-61DAFB.svg)

**Go-Vite** is a powerful CLI tool for generating production-ready desktop applications using Go (backend), React + Vite (frontend), and native webview. It creates a complete, self-contained application architecture in seconds.

---

## 🌟 Features

- **🚀 One Command Setup** - Generate a complete project structure instantly
- **📦 Embedded Architecture** - Single binary contains frontend, backend, and webview
- **⚡ Lightning Fast** - Vite for blazing-fast frontend development
- **🎨 Modern UI** - Tailwind CSS pre-configured with custom theme
- **🔧 Production Ready** - Includes build system, error handling, and logging
- **🔌 Extensible** - Module system for easy feature addition
- **📦 Module Management** - Install, uninstall, and manage local/remote modules
- **🖥️ Cross-Platform** - Build native desktop apps for Windows, macOS, and Linux
- **🎯 Type Safe** - TypeScript support out of the box
- **📊 Complete Backend** - RESTful API with Gin framework
- **🔄 Hot Reload** - Development mode with automatic reloading

---

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [CLI Commands](#cli-commands)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Building for Production](#building-for-production)
- [Configuration](#configuration)
- [Module System](#module-system)
- [API Development](#api-development)
- [Frontend Development](#frontend-development)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)
- [Best Practices](#best-practices)
- [Contributing](#contributing)
- [License](#license)

---

## 📦 Prerequisites

Before using Go-Vite, ensure you have the following installed:

### Required

- **Go** 1.24 or later ([Download](https://golang.org/dl/))
- **Node.js** 18.x or later ([Download](https://nodejs.org/))
- **npm** or **yarn** (comes with Node.js)
- **Git** (for version control)

### Platform-Specific Requirements

#### Linux
```bash
# Ubuntu/Debian
sudo apt-get install webkit2gtk-4.0-dev

# Fedora
sudo dnf install webkit2gtk3-devel

# Arch
sudo pacman -S webkit2gtk
```

#### macOS
```bash
# Xcode Command Line Tools
xcode-select --install
```

#### Windows
- **MinGW-w64** or **MSYS2** for CGO support
- Windows SDK

---

## 🚀 Installation

### Install via Go

```bash
go install github.com/yourusername/go-vite@latest
```

### Build from Source

```bash
git clone https://github.com/yourusername/go-vite.git
cd go-vite
go build -o go-vite .
sudo mv go-vite /usr/local/bin/
```

### Verify Installation

```bash
go-vite version
```

---

## ⚡ Quick Start

### 1. Create a New Project

```bash
# Basic usage
go-vite init my-app

# With custom options
go-vite init my-app \
  --module github.com/myuser/my-app \
  --description "My awesome desktop application" \
  --author "Your Name" \
  --port 5173 \
  --backend-port 8080
```

### 2. Navigate to Project

```bash
cd my-app
```

### 3. Install Dependencies

```bash
make deps
```

### 4. Build the Application

```bash
make binary
```

### 5. Run Your App

```bash
./dist/my-app
```

🎉 **That's it!** Your desktop application is now running.

---

## 🎯 CLI Commands

### `go-vite init [project-name]`

Initialize a new Go-Vite project.

**Usage:**
```bash
go-vite init [project-name] [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--module` | `-m` | `[project-name]` | Go module name |
| `--description` | `-d` | `"A Go-Vite desktop application"` | Project description |
| `--author` | `-a` | `""` | Author name |
| `--port` | `-p` | `5173` | Frontend development port |
| `--backend-port` | `-b` | `8080` | Backend API port |

**Examples:**

```bash
# Minimal
go-vite init todo-app

# Full configuration
go-vite init todo-app \
  --module github.com/john/todo-app \
  --description "A simple todo list application" \
  --author "John Doe" \
  --port 3000 \
  --backend-port 9000
```

### `go-vite version`

Display version information.

```bash
go-vite version
```

### `go-vite install [module]`

Install a module from a remote repository. Automatically detects whether the current project is a Go or Node.js project and uses the appropriate package manager.

**Usage:**
```bash
go-vite install [module]
```

**Examples:**

```bash
# Install a Go module
go-vite install github.com/gin-gonic/gin

# Install a Node.js package
go-vite install axios

# Install with version constraint (Go)
go-vite install github.com/gin-gonic/gin@v1.9.1
```

### `go-vite uninstall [module]`

Uninstall a module from the current project. Automatically detects the project type and uses the appropriate package manager.

**Usage:**
```bash
go-vite uninstall [module]
```

**Examples:**

```bash
# Uninstall a Go module
go-vite uninstall github.com/gin-gonic/gin

# Uninstall a Node.js package
go-vite uninstall axios
```

### `go-vite install-local [path]`

Install a module from a local directory. Copies the module files to the project's modules directory and registers it for use.

**Usage:**
```bash
go-vite install-local [path]
```

**Examples:**

```bash
# Install from a local directory
go-vite install-local ./my-custom-module

# Install from an absolute path
go-vite install-local /home/user/projects/my-module
```

**Requirements:**
- The source directory must contain either `go.mod` (for Go modules) or `package.json` (for Node.js modules)
- The module will be copied to `backend/internal/modules/[module-name]`

### `go-vite import-module [path]`

Import and register a local module without copying files. Similar to `install-local` but checks for existing modules first.

**Usage:**
```bash
go-vite import-module [path]
```

**Examples:**

```bash
# Import a local module
go-vite import-module ./my-module

# Import from an absolute path
go-vite import-module /path/to/module
```

**Note:** If a module with the same name already exists, the command will fail. Use `install-local` to overwrite existing modules.

---

## 📁 Project Structure

```
my-app/
├── main.go                          # Desktop app entry point
├── go.mod                           # Root Go module
├── Makefile                         # Build automation
├── README.md                        # Project documentation
├── .env.example                     # Environment variables template
├── .gitignore                       # Git ignore rules
├── .gitattributes                   # Git attributes
│
├── backend/                         # Go backend service
│   ├── go.mod                       # Backend Go module
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Backend server entry
│   ├── config/
│   │   └── config.go                # Configuration management
│   ├── internal/
│   │   ├── api/
│   │   │   ├── routes.go            # API route definitions
│   │   │   ├── handlers/
│   │   │   │   └── handlers.go      # Request handlers
│   │   │   └── middleware/
│   │   │       ├── cors.go          # CORS middleware
│   │   │       └── logger.go        # Logging middleware
│   │   ├── models/
│   │   │   ├── pipeline.go          # Pipeline data model
│   │   │   ├── project.go           # Project data model
│   │   │   └── user.go              # User data model
│   │   ├── modules/
│   │   │   ├── modules.go           # Module manager
│   │   │   └── builtin.go           # Built-in modules
│   │   ├── storage/
│   │   │   └── database.go          # Database layer
│   │   └── utils/
│   │       └── logger.go            # Utility functions
│   └── tests/
│       └── ...                      # Test files
│
├── frontend/                        # React + Vite frontend
│   ├── src/
│   │   ├── main.tsx                 # React entry point
│   │   ├── App.tsx                  # Main App component
│   │   ├── index.css                # Global styles
│   │   ├── components/              # Reusable components
│   │   ├── pages/                   # Page components
│   │   ├── hooks/                   # Custom React hooks
│   │   ├── services/                # API services
│   │   └── utils/                   # Utility functions
│   ├── public/                      # Static assets
│   ├── index.html                   # HTML template
│   ├── package.json                 # Node dependencies
│   ├── vite.config.js               # Vite configuration
│   ├── tailwind.config.js           # Tailwind CSS config
│   ├── postcss.config.js            # PostCSS config
│   ├── .eslintrc.cjs                # ESLint rules
│   └── .prettierrc                  # Prettier config
│
├── bin/                             # Backend binary (temp)
└── dist/                            # Final application binary
    └── my-app                       # Single executable
```

---

## 💻 Development Workflow

### Development Mode

Run frontend and backend separately for hot-reload during development:

**Terminal 1 - Frontend:**
```bash
cd frontend
npm run dev
```

**Terminal 2 - Backend:**
```bash
cd backend
go run ./cmd/server
```

**Terminal 3 - Build Desktop App:**
```bash
make binary
./dist/my-app
```

### Makefile Commands

| Command | Description |
|---------|-------------|
| `make all` | Clean, install deps, and build everything |
| `make deps` | Install all dependencies (Go + Node.js) |
| `make deps-go` | Install Go dependencies only |
| `make deps-node` | Install Node.js dependencies only |
| `make frontend` | Build React frontend |
| `make backend` | Build Go backend |
| `make binary` | Build unified desktop application |
| `make clean` | Remove build artifacts |
| `make test` | Run all tests |
| `make run` | Build and run the application |

### Quick Development Commands

```bash
# Install dependencies
make deps

# Build everything
make all

# Run the application
make run

# Clean and rebuild
make clean && make binary
```

---

## 🏗️ Building for Production

### Single Platform Build

```bash
# Build for your current platform
make binary
```

### Cross-Platform Builds

```bash
# Linux
GOOS=linux GOARCH=amd64 make binary

# Windows
GOOS=windows GOARCH=amd64 make binary

# macOS (Intel)
GOOS=darwin GOARCH=amd64 make binary

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 make binary
```

### Build Optimization

For smaller binary sizes:

```bash
# Build with optimizations
go build -ldflags="-s -w" -o dist/my-app .

# Further compress with UPX (optional)
upx --best --lzma dist/my-app
```

---

## ⚙️ Configuration

### Environment Variables

Create a `.env` file in the project root:

```bash
# Frontend
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=My Application
VITE_APP_VERSION=1.0.0

# Backend
PORT=8080
LOG_LEVEL=info

# Database (if needed)
DATABASE_URL=postgresql://user:password@localhost:5432/myapp

# API Keys
API_KEY=your_api_key_here
```

### Backend Configuration

Edit `backend/config/config.go`:

```go
type Config struct {
    FrontendPort int
    BackendPort  int
    LogLevel     string
    DatabaseURL  string
    APIKey       string
}
```

### Frontend Configuration

Edit `frontend/vite.config.js`:

```javascript
export default defineConfig({
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

---

## 🔌 Module System

The generated project includes an extensible module system for organizing business logic. Go-Vite also provides CLI commands for managing both remote and local modules.

### Module Management Commands

Go-Vite provides powerful module management capabilities:

#### Installing Remote Modules

```bash
# Install from remote repositories
go-vite install github.com/gin-gonic/gin
go-vite install axios

# Uninstall modules
go-vite uninstall github.com/gin-gonic/gin
go-vite uninstall axios
```

#### Installing Local Modules

```bash
# Install from local directory (copies files)
go-vite install-local ./my-custom-module

# Import without copying (references existing location)
go-vite import-module ./my-existing-module
```

Local modules are automatically:
- Detected as Go or Node.js modules
- Copied to `backend/internal/modules/[module-name]`
- Registered in the module system
- Available for use in your application

### Creating a Custom Module

**1. Create module file:** `backend/internal/modules/mymodule.go`

```go
package modules

type MyModule struct {
    config map[string]interface{}
}

func NewMyModule() *MyModule {
    return &MyModule{
        config: make(map[string]interface{}),
    }
}

func (m *MyModule) Name() string {
    return "mymodule"
}

func (m *MyModule) Execute(input map[string]interface{}) (map[string]interface{}, error) {
    // Your business logic here
    result := map[string]interface{}{
        "status": "success",
        "data":   input,
    }
    return result, nil
}

func (m *MyModule) Validate(config map[string]interface{}) error {
    // Validation logic
    return nil
}
```

**2. Register the module:** `backend/internal/modules/builtin.go`

```go
func LoadBuiltinModules(manager *Manager) {
    manager.Register("example", &ExampleModule{})
    manager.Register("mymodule", NewMyModule())
}
```

**3. Use in handlers:** `backend/internal/api/handlers/handlers.go`

```go
func ExecuteModule(c *gin.Context) {
    moduleManager := c.MustGet("modules").(*modules.Manager)
    
    module, exists := moduleManager.Get("mymodule")
    if !exists {
        c.JSON(404, gin.H{"error": "Module not found"})
        return
    }
    
    input := map[string]interface{}{
        "data": "test",
    }
    
    result, err := module.Execute(input)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, result)
}
```

---

## 🔗 API Development

### Adding New Endpoints

**1. Define route:** `backend/internal/api/routes.go`

```go
func SetupRoutes(router *gin.Engine) {
    v1 := router.Group("/api/v1")
    {
        v1.GET("/users", handlers.ListUsers)
        v1.POST("/users", handlers.CreateUser)
        v1.GET("/users/:id", handlers.GetUser)
    }
}
```

**2. Create handler:** `backend/internal/api/handlers/handlers.go`

```go
func ListUsers(c *gin.Context) {
    users := []map[string]interface{}{
        {"id": 1, "name": "John"},
        {"id": 2, "name": "Jane"},
    }
    c.JSON(200, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
    var user map[string]interface{}
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(201, gin.H{"message": "User created", "user": user})
}
```

### API Testing

```bash
# Health check
curl http://localhost:8080/health

# List items
curl http://localhost:8080/api/v1/items

# Create item
curl -X POST http://localhost:8080/api/v1/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Test Item"}'
```

---

## 🎨 Frontend Development

### Creating Components

**1. Create component:** `frontend/src/components/Button.tsx`

```typescript
import React from 'react';

interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  variant?: 'primary' | 'secondary';
}

export const Button: React.FC<ButtonProps> = ({ 
  children, 
  onClick, 
  variant = 'primary' 
}) => {
  const baseClasses = 'px-6 py-3 rounded-lg font-semibold transition-all';
  const variantClasses = variant === 'primary'
    ? 'bg-gradient-to-r from-brand-600 to-blue-600 hover:from-brand-500'
    : 'bg-slate-700 hover:bg-slate-600';

  return (
    <button
      onClick={onClick}
      className={`${baseClasses} ${variantClasses}`}
    >
      {children}
    </button>
  );
};
```

**2. Use component:** `frontend/src/App.tsx`

```typescript
import { Button } from './components/Button';

function App() {
  return (
    <div>
      <Button onClick={() => console.log('Clicked!')}>
        Click Me
      </Button>
    </div>
  );
}
```

### API Integration

**1. Create service:** `frontend/src/services/api.ts`

```typescript
const API_BASE = import.meta.env.VITE_API_URL || '/api/v1';

export const api = {
  async get(endpoint: string) {
    const response = await fetch(`${API_BASE}${endpoint}`);
    return response.json();
  },

  async post(endpoint: string, data: any) {
    const response = await fetch(`${API_BASE}${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return response.json();
  },
};
```

**2. Use in component:**

```typescript
import { useEffect, useState } from 'react';
import { api } from './services/api';

function ItemsList() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    api.get('/items').then(data => setItems(data.items));
  }, []);

  return (
    <ul>
      {items.map(item => (
        <li key={item.id}>{item.name}</li>
      ))}
    </ul>
  );
}
```

### Styling with Tailwind

The project comes with Tailwind CSS pre-configured with a custom brand color palette:

```typescript
// Use brand colors
<div className="bg-brand-500 text-white">
  Brand colored background
</div>

// Gradient text
<h1 className="bg-gradient-to-r from-brand-400 to-blue-400 bg-clip-text text-transparent">
  Gradient Text
</h1>

// Custom animations
<div className="animate-fade-in">
  Fades in on load
</div>
```

---

## 🚢 Deployment

### Desktop Application Distribution

**1. Build for target platform:**
```bash
# macOS
GOOS=darwin GOARCH=amd64 make binary

# Windows
GOOS=windows GOARCH=amd64 make binary

# Linux
GOOS=linux GOARCH=amd64 make binary
```

**2. Package the binary:**

**macOS:**
```bash
# Create .app bundle
mkdir -p MyApp.app/Contents/MacOS
cp dist/my-app MyApp.app/Contents/MacOS/
# Add Info.plist and icon
```

**Windows:**
```bash
# Use tools like NSIS or Inno Setup
# Or simply distribute the .exe
```

**Linux:**
```bash
# Create AppImage, Snap, or Flatpak
# Or distribute as binary with .desktop file
```

### Web Deployment (Optional)

If you want to deploy as a web app instead:

```bash
# Build frontend only
cd frontend && npm run build

# Deploy dist/ folder to:
# - Netlify
# - Vercel
# - GitHub Pages
# - AWS S3
```

---

## 🐛 Troubleshooting

### Common Issues

#### 1. **Webview not loading**

**Issue:** Blank window or "Failed to load" error

**Solution:**
```bash
# Linux: Install webkit2gtk
sudo apt-get install webkit2gtk-4.0-dev

# macOS: Install Xcode Command Line Tools
xcode-select --install

# Windows: Ensure WebView2 Runtime is installed
```

#### 2. **CGO errors during build**

**Issue:** `CGO_ENABLED=1` required but not set

**Solution:**
```bash
# Install build tools
# Ubuntu/Debian
sudo apt-get install build-essential

# macOS
xcode-select --install

# Windows
# Install MinGW-w64 or MSYS2
```

#### 3. **Port already in use**

**Issue:** `bind: address already in use`

**Solution:**
```bash
# Find and kill process using port
lsof -ti:8080 | xargs kill -9

# Or change port in .env
echo "PORT=9000" >> .env
```

#### 4. **Frontend not connecting to backend**

**Issue:** API calls returning 404 or connection refused

**Solution:**
- Check `vite.config.js` proxy settings
- Ensure backend is running on correct port
- Verify CORS middleware is enabled
- Check `.env` file has correct `VITE_API_URL`

#### 5. **Module not found errors**

**Issue:** `cannot find package` during Go build

**Solution:**
```bash
# Clean and reinstall
rm -rf go.sum
go clean -modcache
go mod download
go mod tidy
```

### Debug Mode

Enable verbose logging:

```bash
# Backend
LOG_LEVEL=debug go run ./cmd/server

# Frontend
DEBUG=vite:* npm run dev
```

---

## ✅ Best Practices

### Project Organization

1. **Keep business logic in modules** - Don't put logic in handlers
2. **Use consistent naming** - Follow Go and React conventions
3. **Write tests** - Aim for >80% coverage
4. **Document APIs** - Use OpenAPI/Swagger for backend
5. **Type everything** - Use TypeScript strictly

### Code Quality

```bash
# Go
go fmt ./...
go vet ./...
golangci-lint run

# Frontend
npm run lint
npm run format
```

### Security

1. **Never commit secrets** - Use `.env` files
2. **Validate all inputs** - Backend and frontend
3. **Use HTTPS in production**
4. **Implement rate limiting**
5. **Keep dependencies updated**

```bash
# Update Go dependencies
go get -u ./...
go mod tidy

# Update Node dependencies
npm update
npm audit fix
```

### Performance

1. **Optimize images** - Use WebP format
2. **Lazy load components** - React.lazy() and Suspense
3. **Bundle splitting** - Configure Vite code splitting
4. **Database indexing** - Add indexes for common queries
5. **Enable compression** - gzip middleware in production

---

## 🤝 Contributing

We welcome contributions! Here's how to get started:

### Development Setup

```bash
# Fork and clone
git clone https://github.com/yourusername/go-vite.git
cd go-vite

# Create branch
git checkout -b feature/amazing-feature

# Make changes and test
go build .
./go-vite init test-app

# Commit and push
git commit -m "Add amazing feature"
git push origin feature/amazing-feature
```

### Pull Request Process

1. Update README.md with details of changes
2. Add tests for new functionality
3. Ensure all tests pass
4. Update version numbers following SemVer
5. Submit PR with clear description

### Reporting Issues

When reporting issues, include:
- Go version: `go version`
- Node version: `node --version`
- Operating system
- Error messages and logs
- Steps to reproduce

---

## 📝 Examples

Check out example projects:

- **[Todo App](https://github.com/yourusername/go-vite-todo)** - Simple task manager
- **[Dashboard](https://github.com/yourusername/go-vite-dashboard)** - Analytics dashboard
- **[Chat App](https://github.com/yourusername/go-vite-chat)** - Real-time messaging

---

## 🗺️ Roadmap

- [ ] Support for additional frontend frameworks (Vue, Svelte)
- [ ] Built-in database migrations
- [ ] Authentication scaffolding
- [ ] Docker deployment templates
- [ ] CI/CD pipeline templates
- [ ] Plugin system for custom generators
- [ ] GUI configuration tool
- [ ] Auto-update mechanism for apps

---

## 📚 Resources

- **[Go Documentation](https://golang.org/doc/)**
- **[Vite Documentation](https://vitejs.dev/)**
- **[React Documentation](https://react.dev/)**
- **[Gin Framework](https://gin-gonic.com/)**
- **[Tailwind CSS](https://tailwindcss.com/)**
- **[Webview](https://github.com/webview/webview)**

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- **[Webview](https://github.com/webview/webview)** - For the native webview library
- **[Gin](https://github.com/gin-gonic/gin)** - For the excellent web framework
- **[Vite](https://vitejs.dev/)** - For the blazing-fast build tool
- **[Cobra](https://github.com/spf13/cobra)** - For CLI framework
- All contributors and users of this project

---

## 💬 Community & Support

- **GitHub Issues:** [Report bugs or request features](https://github.com/yourusername/go-vite/issues)
- **Discussions:** [Community forum](https://github.com/yourusername/go-vite/discussions)
- **Twitter:** [@govitecli](https://twitter.com/govitecli)
- **Discord:** [Join our server](https://discord.gg/govite)

---

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=yourusername/go-vite&type=Date)](https://star-history.com/#yourusername/go-vite&Date)

---

<div align="center">

**[Website](https://govite.dev)** • **[Documentation](https://docs.govite.dev)** • **[Blog](https://blog.govite.dev)**

Made with ❤️ by the Go-Vite team

</div>
