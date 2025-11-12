# Go-Vite Website

![Go-Vite](https://img.shields.io/badge/Go--Vite-Website-blue.svg)
![React](https://img.shields.io/badge/React-18.x-61DAFB.svg)
![Vite](https://img.shields.io/badge/Vite-5.x-646CFF.svg)
![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6.svg)

The official website and documentation hub for **Go-Vite**, a powerful CLI tool for generating production-ready desktop applications using Go (backend), React + Vite (frontend), and native webview.

This website serves as both a **sales landing page** showcasing Go-Vite's features and capabilities, and a comprehensive **documentation hub** providing installation guides, API references, and development resources.

## 🌟 Features

- **📄 Sales Landing Page** - Modern, responsive landing page highlighting Go-Vite's key features and benefits
- **📚 Documentation Hub** - Complete documentation with installation guides, CLI commands, and development workflows
- **🎨 Modern UI** - Built with React, TypeScript, and Tailwind CSS using shadcn/ui components
- **⚡ Fast & Optimized** - Powered by Vite for lightning-fast development and builds
- **📱 Responsive Design** - Works perfectly on desktop, tablet, and mobile devices
- **🔍 SEO Optimized** - Structured for search engines and developer discovery

## 🚀 Quick Start

### Prerequisites

- **Node.js** 18.x or later ([Download](https://nodejs.org/))
- **npm** or **yarn** (comes with Node.js)

### Installation & Development

```bash
# Navigate to the website directory
cd website

# Install dependencies
npm install

# Start development server
npm run dev
```

The website will be available at `http://localhost:5173` with hot reload enabled.

### Build for Production

```bash
# Build the website
npm run build

# Preview the production build
npm run preview
```

## 📁 Project Structure

```
website/
├── src/
│   ├── components/          # Reusable UI components
│   │   ├── ui/             # shadcn/ui components
│   │   ├── FeatureCard.tsx # Feature showcase cards
│   │   ├── CodeBlock.tsx   # Syntax-highlighted code blocks
│   │   └── NavLink.tsx     # Navigation components
│   ├── pages/              # Main page components
│   │   ├── Index.tsx       # Landing page
│   │   ├── Docs.tsx        # Documentation page
│   │   └── NotFound.tsx    # 404 page
│   ├── hooks/              # Custom React hooks
│   ├── lib/                # Utility functions
│   ├── assets/             # Static assets (images, icons)
│   └── App.tsx             # Main app component
├── public/                 # Public static files
├── index.html              # HTML template
├── package.json            # Dependencies and scripts
├── vite.config.ts          # Vite configuration
├── tailwind.config.ts      # Tailwind CSS configuration
└── tsconfig.json           # TypeScript configuration
```

## 🛠️ Technologies Used

This website is built with modern web technologies:

- **React 18** - Component-based UI library
- **TypeScript** - Type-safe JavaScript
- **Vite** - Fast build tool and dev server
- **Tailwind CSS** - Utility-first CSS framework
- **shadcn/ui** - High-quality React components
- **Radix UI** - Accessible component primitives
- **Lucide React** - Beautiful icon library
- **React Router** - Client-side routing

## 🎨 Customization

### Styling

The website uses Tailwind CSS with custom design tokens defined in `tailwind.config.ts`. Key design elements include:

- **Brand Colors** - Custom gradient themes for branding
- **Dark Mode** - Automatic dark/light mode support
- **Responsive Design** - Mobile-first approach
- **Custom Components** - Reusable UI components with consistent styling

### Content Management

- **Landing Page** (`src/pages/Index.tsx`) - Hero section, features, quick start guide
- **Documentation** (`src/pages/Docs.tsx`) - Comprehensive docs with sidebar navigation
- **Assets** - Logo, hero images, and other media in `src/assets/`

## 🚀 Deployment

### Build Commands

```bash
# Development
npm run dev

# Production build
npm run build

# Preview production build
npm run preview

# Type checking
npm run lint
```

### Deployment Options

The website can be deployed to any static hosting service:

- **Vercel** - `npm run build` outputs to `dist/`
- **Netlify** - Automatic deployments from Git
- **GitHub Pages** - Static hosting for the repository
- **AWS S3 + CloudFront** - Scalable static hosting
- **Railway** - Full-stack deployment platform

### Environment Variables

Create a `.env` file for local development:

```bash
# Optional: Analytics, API keys, etc.
VITE_ANALYTICS_ID=your_analytics_id
VITE_API_URL=https://api.govite.dev
```

## 🤝 Contributing

This website is part of the Go-Vite project. To contribute:

1. **Fork** the repository
2. **Clone** your fork: `git clone https://github.com/yourusername/go-vite.git`
3. **Navigate** to website: `cd go-vite/website`
4. **Install** dependencies: `npm install`
5. **Create** a feature branch: `git checkout -b feature/website-improvement`
6. **Make** your changes and test them
7. **Commit** your changes: `git commit -m "Improve website feature"`
8. **Push** to your fork: `git push origin feature/website-improvement`
9. **Open** a Pull Request

### Content Guidelines

- Keep documentation accurate and up-to-date with the CLI tool
- Ensure all code examples are tested and functional
- Maintain consistent styling and branding
- Test responsiveness across different screen sizes

## 📄 License

This website is part of the Go-Vite project and is licensed under the MIT License - see the main [LICENSE](../LICENSE) file for details.

## 🔗 Related Links

- **[Go-Vite CLI](https://github.com/guiperry/go-vite)** - The main CLI tool repository
- **[Go-Vite Documentation](https://docs.govite.dev)** - Online documentation
- **[Go-Vite Website](https://govite.dev)** - Live website
- **[Issues](https://github.com/guiperry/go-vite/issues)** - Report bugs or request features
- **[Discussions](https://github.com/guiperry/go-vite/discussions)** - Community forum

---

<div align="center">

**Built with ❤️ for the Go-Vite community**

**[CLI Tool](../README.md)** • **[Documentation](./src/pages/Docs.tsx)** • **[GitHub](https://github.com/guiperry/go-vite)**

</div>
