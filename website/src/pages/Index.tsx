import { Button } from "@/components/ui/button";
import { FeatureCard } from "@/components/FeatureCard";
import { CodeBlock } from "@/components/CodeBlock";
import { Link } from "react-router-dom";
import {
  Zap,
  Package,
  Boxes,
  Palette,
  Settings,
  Plug,
  Monitor,
  Code,
  RefreshCw,
  Github,
  Terminal,
  Rocket,
} from "lucide-react";
import logo from "@/assets/logo.png";
import heroBg from "@/assets/hero-bg.jpg";

const Index = () => {
  const features = [
    {
      icon: Rocket,
      title: "One Command Setup",
      description: "Generate a complete project structure instantly with a single CLI command.",
    },
    {
      icon: Package,
      title: "Embedded Architecture",
      description: "Single binary contains frontend, backend, and webview for easy distribution.",
    },
    {
      icon: Zap,
      title: "Lightning Fast",
      description: "Vite for blazing-fast frontend development with hot module replacement.",
    },
    {
      icon: Palette,
      title: "Modern UI",
      description: "Tailwind CSS pre-configured with custom theme for beautiful interfaces.",
    },
    {
      icon: Settings,
      title: "Production Ready",
      description: "Includes build system, error handling, and comprehensive logging out of the box.",
    },
    {
      icon: Plug,
      title: "Extensible",
      description: "Modular architecture allows easy addition of features and functionality.",
    },
    {
      icon: Monitor,
      title: "Cross-Platform",
      description: "Build native desktop apps for Windows, macOS, and Linux from one codebase.",
    },
    {
      icon: Code,
      title: "Type Safe",
      description: "Full TypeScript support for enhanced developer experience and code quality.",
    },
    {
      icon: Boxes,
      title: "Complete Backend",
      description: "RESTful API with Gin framework ready for your backend logic.",
    },
    {
      icon: RefreshCw,
      title: "Hot Reload",
      description: "Development mode with automatic reloading for rapid iteration.",
    },
  ];

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
              <Link to="/docs">
                <Button variant="ghost">Documentation</Button>
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

      {/* Hero Section */}
      <section className="relative pt-32 pb-20 overflow-hidden">
        <div
          className="absolute inset-0 opacity-20"
          style={{
            backgroundImage: `url(${heroBg})`,
            backgroundSize: "cover",
            backgroundPosition: "center",
          }}
        />
        <div className="absolute inset-0 bg-gradient-to-b from-background via-background/95 to-background" />
        
        <div className="container mx-auto px-6 relative z-10">
          <div className="max-w-4xl mx-auto text-center">
            <div className="flex justify-center mb-8">
              <img src={logo} alt="Go-Vite" className="w-32 h-32 animate-pulse" />
            </div>
            <h1 className="text-5xl md:text-7xl font-bold mb-6 bg-gradient-hero bg-clip-text text-transparent">
              Build Desktop Apps at Lightning Speed
            </h1>
            <p className="text-xl md:text-2xl text-muted-foreground mb-8 leading-relaxed">
              A powerful CLI tool for generating production-ready desktop applications using Go,
              React + Vite, and native webview. Create complete, self-contained apps in seconds.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center mb-12">
              <Button variant="hero" size="lg" asChild>
                <Link to="/docs">
                  <Terminal className="w-5 h-5 mr-2" />
                  Get Started
                </Link>
              </Button>
              <Button variant="outline" size="lg" asChild>
                <a href="https://github.com/guiperry/go-vite" target="_blank" rel="noopener noreferrer">
                  <Github className="w-5 h-5 mr-2" />
                  View on GitHub
                </a>
              </Button>
            </div>

            {/* Quick Install */}
            <div className="max-w-2xl mx-auto">
              <CodeBlock code="go install github.com/yourusername/go-vite@latest" />
            </div>
          </div>
        </div>
      </section>

      {/* Quick Start Section */}
      <section className="py-20 bg-gradient-feature">
        <div className="container mx-auto px-6">
          <div className="max-w-5xl mx-auto">
            <h2 className="text-3xl md:text-4xl font-bold text-center mb-4">
              From Zero to App in 30 Seconds
            </h2>
            <p className="text-center text-muted-foreground mb-12 text-lg">
              Create and run your first desktop application in just a few commands
            </p>
            
            <div className="grid gap-6">
              <div>
                <h3 className="text-lg font-semibold mb-3 flex items-center gap-2">
                  <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                    1
                  </span>
                  Create a new project
                </h3>
                <CodeBlock code="go-vite init my-awesome-app" />
              </div>
              
              <div>
                <h3 className="text-lg font-semibold mb-3 flex items-center gap-2">
                  <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                    2
                  </span>
                  Navigate and install dependencies
                </h3>
                <CodeBlock code="cd my-awesome-app && make deps" />
              </div>
              
              <div>
                <h3 className="text-lg font-semibold mb-3 flex items-center gap-2">
                  <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground text-sm font-bold">
                    3
                  </span>
                  Build and run
                </h3>
                <CodeBlock code="make binary && ./dist/my-awesome-app" />
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20">
        <div className="container mx-auto px-6">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              Everything You Need, Out of the Box
            </h2>
            <p className="text-muted-foreground text-lg max-w-2xl mx-auto">
              Go-Vite comes packed with modern tools and features to help you build amazing desktop applications
            </p>
          </div>
          
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-6xl mx-auto">
            {features.map((feature, index) => (
              <FeatureCard key={index} {...feature} />
            ))}
          </div>
        </div>
      </section>

      {/* Tech Stack Section */}
      <section className="py-20 bg-gradient-feature">
        <div className="container mx-auto px-6">
          <div className="max-w-4xl mx-auto text-center">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              Built with Modern Technologies
            </h2>
            <p className="text-muted-foreground text-lg mb-12">
              Leveraging the best tools for frontend and backend development
            </p>
            
            <div className="grid md:grid-cols-3 gap-8">
              <div className="p-6">
                <div className="text-4xl mb-4">🐹</div>
                <h3 className="text-xl font-semibold mb-2">Go 1.24+</h3>
                <p className="text-muted-foreground">
                  Fast, efficient backend with the power of Go's concurrency
                </p>
              </div>
              
              <div className="p-6">
                <div className="text-4xl mb-4">⚡</div>
                <h3 className="text-xl font-semibold mb-2">React + Vite</h3>
                <p className="text-muted-foreground">
                  Lightning-fast frontend with modern React and Vite tooling
                </p>
              </div>
              
              <div className="p-6">
                <div className="text-4xl mb-4">🎨</div>
                <h3 className="text-xl font-semibold mb-2">Tailwind CSS</h3>
                <p className="text-muted-foreground">
                  Beautiful, responsive UIs with utility-first CSS framework
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20">
        <div className="container mx-auto px-6">
          <div className="max-w-4xl mx-auto text-center">
            <h2 className="text-3xl md:text-4xl font-bold mb-6">
              Ready to Build Something Amazing?
            </h2>
            <p className="text-xl text-muted-foreground mb-8">
              Start creating powerful desktop applications today with Go-Vite
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Button variant="hero" size="lg" asChild>
                <Link to="/docs">
                  Read the Documentation
                </Link>
              </Button>
              <Button variant="outline" size="lg" asChild>
                <a href="https://github.com/guiperry/go-vite" target="_blank" rel="noopener noreferrer">
                  Star on GitHub
                </a>
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-border/50 py-8">
        <div className="container mx-auto px-6">
          <div className="flex flex-col md:flex-row items-center justify-between gap-4">
            <div className="flex items-center gap-2">
              <img src={logo} alt="Go-Vite" className="w-6 h-6" />
              <span className="text-sm text-muted-foreground">
                © 2024 Go-Vite. MIT License.
              </span>
            </div>
            <div className="flex gap-6">
              <Link to="/docs" className="text-sm text-muted-foreground hover:text-primary transition-colors">
                Documentation
              </Link>
              <a
                href="https://github.com/guiperry/go-vite"
                target="_blank"
                rel="noopener noreferrer"
                className="text-sm text-muted-foreground hover:text-primary transition-colors"
              >
                GitHub
              </a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default Index;
