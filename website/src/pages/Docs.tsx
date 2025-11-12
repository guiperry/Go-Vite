import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Link } from "react-router-dom";
import { Home, Github, Menu } from "lucide-react";
import logo from "@/assets/logo.png";
import { useState } from "react";
import { cn } from "@/lib/utils";

const Docs = () => {
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const sections = [
    { id: "intro", title: "Introduction" },
    { id: "features", title: "Features" },
    { id: "prerequisites", title: "Prerequisites" },
    { id: "installation", title: "Installation" },
    { id: "quick-start", title: "Quick Start" },
    { id: "cli-commands", title: "CLI Commands" },
    { id: "project-structure", title: "Project Structure" },
    { id: "development", title: "Development Workflow" },
    { id: "building", title: "Building for Production" },
    { id: "configuration", title: "Configuration" },
  ];

  const scrollToSection = (id: string) => {
    const element = document.getElementById(id);
    if (element) {
      element.scrollIntoView({ behavior: "smooth", block: "start" });
      setSidebarOpen(false);
    }
  };

  return (
    <div className="min-h-screen">
      {/* Navigation */}
      <nav className="fixed top-0 left-0 right-0 z-50 bg-background/80 backdrop-blur-md border-b border-border/50">
        <div className="container mx-auto px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <img src={logo} alt="Go-Vite Logo" className="w-10 h-10" />
              <span className="text-2xl font-bold bg-gradient-hero bg-clip-text text-transparent">
                Go-Vite
              </span>
            </div>
            <div className="flex items-center gap-4">
              <Button variant="ghost" size="icon" className="md:hidden" onClick={() => setSidebarOpen(!sidebarOpen)}>
                <Menu className="w-5 h-5" />
              </Button>
              <Link to="/">
                <Button variant="ghost">
                  <Home className="w-4 h-4 mr-2" />
                  Home
                </Button>
              </Link>
              <Button variant="outline" asChild>
                <a href="https://github.com/guiperry/go-vite" target="_blank" rel="noopener noreferrer">
                  <Github className="w-4 h-4 mr-2" />
                  GitHub
                </a>
              </Button>
            </div>
          </div>
        </div>
      </nav>

      <div className="flex pt-20">
        {/* Sidebar */}
        <aside
          className={cn(
            "fixed left-0 top-20 bottom-0 w-64 bg-card/50 backdrop-blur-sm border-r border-border/50 p-6 overflow-y-auto transition-transform duration-300 z-40",
            sidebarOpen ? "translate-x-0" : "-translate-x-full md:translate-x-0"
          )}
        >
          <h2 className="text-lg font-semibold mb-4">Documentation</h2>
          <nav className="space-y-1">
            {sections.map((section) => (
              <button
                key={section.id}
                onClick={() => scrollToSection(section.id)}
                className="block w-full text-left px-3 py-2 text-sm text-muted-foreground hover:text-foreground hover:bg-accent/20 rounded-md transition-colors"
              >
                {section.title}
              </button>
            ))}
          </nav>
        </aside>

        {/* Main Content */}
        <main className="flex-1 md:ml-64 px-6 py-8">
          <div className="max-w-4xl mx-auto prose prose-invert">
            <section id="intro" className="mb-16 scroll-mt-24">
              <h1 className="text-4xl font-bold mb-4 bg-gradient-hero bg-clip-text text-transparent">
                Go-Vite Documentation
              </h1>
              <p className="text-xl text-muted-foreground mb-6">
                A powerful CLI tool for generating production-ready desktop applications using Go (backend),
                React + Vite (frontend), and native webview.
              </p>
              <Card className="p-6 bg-gradient-feature border-border/50">
                <div className="flex items-center gap-2 mb-2">
                  <span className="inline-block px-3 py-1 text-xs font-semibold rounded-full bg-primary/20 text-primary border border-primary/30">
                    v1.0.0
                  </span>
                  <span className="inline-block px-3 py-1 text-xs font-semibold rounded-full bg-secondary/20 text-secondary border border-secondary/30">
                    MIT License
                  </span>
                </div>
                <p className="text-sm text-muted-foreground">
                  Create complete, self-contained application architecture in seconds
                </p>
              </Card>
            </section>

            <section id="features" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Features</h2>
              <div className="grid gap-4">
                {[
                  "🚀 One Command Setup - Generate a complete project structure instantly",
                  "📦 Embedded Architecture - Single binary contains frontend, backend, and webview",
                  "⚡ Lightning Fast - Vite for blazing-fast frontend development",
                  "🎨 Modern UI - Tailwind CSS pre-configured with custom theme",
                  "🔧 Production Ready - Includes build system, error handling, and logging",
                  "🔌 Extensible - Module system for easy feature addition",
                  "📦 Module Management - Install, uninstall, and manage local/remote modules",
                  "🖥️ Cross-Platform - Build native desktop apps for Windows, macOS, and Linux",
                  "🎯 Type Safe - TypeScript support out of the box",
                  "📊 Complete Backend - RESTful API with Gin framework",
                  "🔄 Hot Reload - Development mode with automatic reloading",
                ].map((feature, i) => (
                  <Card key={i} className="p-4 bg-card/50 border-border/50">
                    <p className="text-foreground">{feature}</p>
                  </Card>
                ))}
              </div>
            </section>

            <section id="prerequisites" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Prerequisites</h2>
              <p className="text-muted-foreground mb-4">
                Before using Go-Vite, ensure you have the following installed:
              </p>
              
              <h3 className="text-2xl font-semibold mb-4 text-foreground">Required</h3>
              <ul className="space-y-2 mb-6">
                <li className="text-foreground">• <strong>Go</strong> 1.24 or later</li>
                <li className="text-foreground">• <strong>Node.js</strong> 18.x or later</li>
                <li className="text-foreground">• <strong>npm</strong> or <strong>yarn</strong> (comes with Node.js)</li>
                <li className="text-foreground">• <strong>Git</strong> (for version control)</li>
              </ul>

              <h3 className="text-2xl font-semibold mb-4 text-foreground">Platform-Specific Requirements</h3>
              
              <div className="space-y-4">
                <Card className="p-4 bg-card/50 border-border/50">
                  <h4 className="font-semibold mb-2 text-foreground">Linux</h4>
                  <pre className="text-sm bg-background/50 p-3 rounded overflow-x-auto">
                    <code className="text-primary">{`# Ubuntu/Debian
sudo apt-get install webkit2gtk-4.0-dev

# Fedora
sudo dnf install webkit2gtk3-devel

# Arch
sudo pacman -S webkit2gtk`}</code>
                  </pre>
                </Card>

                <Card className="p-4 bg-card/50 border-border/50">
                  <h4 className="font-semibold mb-2 text-foreground">macOS</h4>
                  <pre className="text-sm bg-background/50 p-3 rounded overflow-x-auto">
                    <code className="text-primary">{`# Xcode Command Line Tools
xcode-select --install`}</code>
                  </pre>
                </Card>

                <Card className="p-4 bg-card/50 border-border/50">
                  <h4 className="font-semibold mb-2 text-foreground">Windows</h4>
                  <ul className="text-sm space-y-1 text-muted-foreground">
                    <li>• MinGW-w64 or MSYS2 for CGO support</li>
                    <li>• Windows SDK</li>
                  </ul>
                </Card>
              </div>
            </section>

            <section id="installation" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Installation</h2>
              
              <h3 className="text-2xl font-semibold mb-4 text-foreground">Install via Go</h3>
              <Card className="p-4 bg-card/50 border-border/50 mb-6">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">go install github.com/guiperry/go-vite@latest</code>
                </pre>
              </Card>

              <h3 className="text-2xl font-semibold mb-4 text-foreground">Build from Source</h3>
              <Card className="p-4 bg-card/50 border-border/50 mb-6">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">{`git clone https://github.com/guiperry/go-vite.git
cd go-vite
go build -o go-vite .
sudo mv go-vite /usr/local/bin/`}</code>
                </pre>
              </Card>

              <h3 className="text-2xl font-semibold mb-4 text-foreground">Verify Installation</h3>
              <Card className="p-4 bg-card/50 border-border/50">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">go-vite version</code>
                </pre>
              </Card>
            </section>

            <section id="quick-start" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Quick Start</h2>
              
              <div className="space-y-6">
                <div>
                  <h3 className="text-xl font-semibold mb-3 text-foreground flex items-center gap-2">
                    <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                      1
                    </span>
                    Create a New Project
                  </h3>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">{`# Basic usage
go-vite init my-app

# With custom options
go-vite init my-app \\
  --module github.com/myuser/my-app \\
  --description "My awesome desktop application" \\
  --author "Your Name" \\
  --port 5173 \\
  --backend-port 8080`}</code>
                    </pre>
                  </Card>
                </div>

                <div>
                  <h3 className="text-xl font-semibold mb-3 text-foreground flex items-center gap-2">
                    <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                      2
                    </span>
                    Navigate to Project
                  </h3>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">cd my-app</code>
                    </pre>
                  </Card>
                </div>

                <div>
                  <h3 className="text-xl font-semibold mb-3 text-foreground flex items-center gap-2">
                    <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                      3
                    </span>
                    Install Dependencies
                  </h3>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">make deps</code>
                    </pre>
                  </Card>
                </div>

                <div>
                  <h3 className="text-xl font-semibold mb-3 text-foreground flex items-center gap-2">
                    <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                      4
                    </span>
                    Build the Application
                  </h3>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">make binary</code>
                    </pre>
                  </Card>
                </div>

                <div>
                  <h3 className="text-xl font-semibold mb-3 text-foreground flex items-center gap-2">
                    <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                      5
                    </span>
                    Run Your App
                  </h3>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">./dist/my-app</code>
                    </pre>
                  </Card>
                </div>
              </div>

              <Card className="p-6 mt-6 bg-primary/10 border-primary/30">
                <p className="text-lg font-semibold text-foreground">
                  🎉 That's it! Your desktop application is now running.
                </p>
              </Card>
            </section>

            <section id="cli-commands" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">CLI Commands</h2>

              <div className="space-y-8">
                <div>
                  <h3 className="text-2xl font-semibold mb-4 text-foreground">
                    <code className="text-primary">go-vite init [project-name]</code>
                  </h3>
                  <p className="text-muted-foreground mb-4">Initialize a new Go-Vite project.</p>
                  
                  <h4 className="font-semibold mb-2 text-foreground">Flags:</h4>
                  <Card className="p-4 bg-card/50 border-border/50 overflow-x-auto">
                    <table className="w-full text-sm">
                      <thead>
                        <tr className="border-b border-border/50">
                          <th className="text-left py-2 px-2 text-foreground">Flag</th>
                          <th className="text-left py-2 px-2 text-foreground">Short</th>
                          <th className="text-left py-2 px-2 text-foreground">Default</th>
                          <th className="text-left py-2 px-2 text-foreground">Description</th>
                        </tr>
                      </thead>
                      <tbody className="text-muted-foreground">
                        <tr className="border-b border-border/30">
                          <td className="py-2 px-2 font-mono text-primary">--module</td>
                          <td className="py-2 px-2 font-mono">-m</td>
                          <td className="py-2 px-2">[project-name]</td>
                          <td className="py-2 px-2">Go module name</td>
                        </tr>
                        <tr className="border-b border-border/30">
                          <td className="py-2 px-2 font-mono text-primary">--description</td>
                          <td className="py-2 px-2 font-mono">-d</td>
                          <td className="py-2 px-2">"A Go-Vite desktop application"</td>
                          <td className="py-2 px-2">Project description</td>
                        </tr>
                        <tr className="border-b border-border/30">
                          <td className="py-2 px-2 font-mono text-primary">--author</td>
                          <td className="py-2 px-2 font-mono">-a</td>
                          <td className="py-2 px-2">""</td>
                          <td className="py-2 px-2">Author name</td>
                        </tr>
                        <tr className="border-b border-border/30">
                          <td className="py-2 px-2 font-mono text-primary">--port</td>
                          <td className="py-2 px-2 font-mono">-p</td>
                          <td className="py-2 px-2">5173</td>
                          <td className="py-2 px-2">Frontend development port</td>
                        </tr>
                        <tr>
                          <td className="py-2 px-2 font-mono text-primary">--backend-port</td>
                          <td className="py-2 px-2 font-mono">-b</td>
                          <td className="py-2 px-2">8080</td>
                          <td className="py-2 px-2">Backend API port</td>
                        </tr>
                      </tbody>
                    </table>
                  </Card>
                </div>

                <div>
                  <h3 className="text-2xl font-semibold mb-4 text-foreground">
                    <code className="text-primary">go-vite install [module]</code>
                  </h3>
                  <p className="text-muted-foreground mb-4">
                    Install a module from a remote repository. Automatically detects project type.
                  </p>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">{`# Install a Go module
go-vite install github.com/gin-gonic/gin

# Install a Node.js package
go-vite install axios`}</code>
                    </pre>
                  </Card>
                </div>

                <div>
                  <h3 className="text-2xl font-semibold mb-4 text-foreground">
                    <code className="text-primary">go-vite uninstall [module]</code>
                  </h3>
                  <p className="text-muted-foreground mb-4">
                    Uninstall a module from the current project.
                  </p>
                  <Card className="p-4 bg-card/50 border-border/50">
                    <pre className="text-sm overflow-x-auto">
                      <code className="text-primary">{`# Uninstall a Go module
go-vite uninstall github.com/gin-gonic/gin

# Uninstall a Node.js package
go-vite uninstall axios`}</code>
                    </pre>
                  </Card>
                </div>
              </div>
            </section>

            <section id="project-structure" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Project Structure</h2>
              <Card className="p-6 bg-card/50 border-border/50">
                <pre className="text-sm overflow-x-auto text-muted-foreground font-mono">
{`my-app/
├── main.go                    # Desktop app entry point
├── go.mod                     # Root Go module
├── Makefile                   # Build automation
├── README.md                  # Project documentation
├── .env.example               # Environment variables
├── .gitignore                 # Git ignore rules
│
├── backend/                   # Go backend service
│   ├── go.mod
│   ├── cmd/
│   │   └── server/
│   │       └── main.go        # Backend server entry
│   ├── config/
│   │   └── config.go          # Configuration
│   ├── internal/
│   │   ├── api/
│   │   │   ├── routes.go      # API routes
│   │   │   ├── handlers/
│   │   │   └── middleware/
│   │   ├── models/            # Data models
│   │   ├── modules/           # Module system
│   │   ├── storage/           # Database layer
│   │   └── utils/
│   └── tests/
│
├── frontend/                  # React + Vite frontend
│   ├── src/
│   │   ├── main.tsx
│   │   ├── App.tsx
│   │   ├── index.css
│   │   ├── components/
│   │   ├── pages/
│   │   ├── hooks/
│   │   ├── services/
│   │   └── utils/
│   ├── public/
│   ├── package.json
│   ├── vite.config.js
│   └── tailwind.config.js
│
└── dist/                      # Final application binary
    └── my-app`}
                </pre>
              </Card>
            </section>

            <section id="development" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Development Workflow</h2>
              
              <h3 className="text-2xl font-semibold mb-4 text-foreground">Makefile Commands</h3>
              <Card className="p-4 bg-card/50 border-border/50 mb-6 overflow-x-auto">
                <table className="w-full text-sm">
                  <thead>
                    <tr className="border-b border-border/50">
                      <th className="text-left py-2 px-2 text-foreground">Command</th>
                      <th className="text-left py-2 px-2 text-foreground">Description</th>
                    </tr>
                  </thead>
                  <tbody className="text-muted-foreground">
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make all</td>
                      <td className="py-2 px-2">Clean, install deps, and build everything</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make deps</td>
                      <td className="py-2 px-2">Install all dependencies (Go + Node.js)</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make frontend</td>
                      <td className="py-2 px-2">Build React frontend</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make backend</td>
                      <td className="py-2 px-2">Build Go backend</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make binary</td>
                      <td className="py-2 px-2">Build unified desktop application</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make clean</td>
                      <td className="py-2 px-2">Remove build artifacts</td>
                    </tr>
                    <tr className="border-b border-border/30">
                      <td className="py-2 px-2 font-mono text-primary">make test</td>
                      <td className="py-2 px-2">Run all tests</td>
                    </tr>
                    <tr>
                      <td className="py-2 px-2 font-mono text-primary">make run</td>
                      <td className="py-2 px-2">Build and run the application</td>
                    </tr>
                  </tbody>
                </table>
              </Card>
            </section>

            <section id="building" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Building for Production</h2>
              
              <h3 className="text-2xl font-semibold mb-4 text-foreground">Single Platform Build</h3>
              <Card className="p-4 bg-card/50 border-border/50 mb-6">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">make binary</code>
                </pre>
              </Card>

              <h3 className="text-2xl font-semibold mb-4 text-foreground">Cross-Platform Builds</h3>
              <Card className="p-4 bg-card/50 border-border/50 mb-6">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">{`# Linux
GOOS=linux GOARCH=amd64 make binary

# Windows
GOOS=windows GOARCH=amd64 make binary

# macOS (Intel)
GOOS=darwin GOARCH=amd64 make binary

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 make binary`}</code>
                </pre>
              </Card>

              <h3 className="text-2xl font-semibold mb-4 text-foreground">Build Optimization</h3>
              <Card className="p-4 bg-card/50 border-border/50">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">{`# Build with optimizations
go build -ldflags="-s -w" -o dist/my-app .

# Further compress with UPX (optional)
upx --best --lzma dist/my-app`}</code>
                </pre>
              </Card>
            </section>

            <section id="configuration" className="mb-16 scroll-mt-24">
              <h2 className="text-3xl font-bold mb-6 text-foreground">Configuration</h2>
              
              <h3 className="text-2xl font-semibold mb-4 text-foreground">Environment Variables</h3>
              <p className="text-muted-foreground mb-4">
                Create a <code className="text-primary">.env</code> file in the project root:
              </p>
              <Card className="p-4 bg-card/50 border-border/50">
                <pre className="text-sm overflow-x-auto">
                  <code className="text-primary">{`# Frontend
VITE_API_URL=http://localhost:8080
VITE_APP_NAME=My Application
VITE_APP_VERSION=1.0.0

# Backend
PORT=8080
LOG_LEVEL=info

# Database (if needed)
DATABASE_URL=postgresql://user:password@localhost:5432/myapp

# API Keys
API_KEY=your_api_key_here`}</code>
                </pre>
              </Card>
            </section>
          </div>
        </main>
      </div>
    </div>
  );
};

export default Docs;
