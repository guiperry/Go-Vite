# example

A Go-Vite desktop application

## 🚀 Quick Start

### Prerequisites
- Go 1.24+
- Node.js 18+
- npm/yarn

### Installation & Running

```bash
# Install dependencies
make deps

# Build the application
make binary

# Run the desktop application
./dist/example
```

## 📁 Project Structure

```
example/
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
```

## 🛠️ Development

```bash
# Run frontend dev server
cd frontend && npm run dev

# Run backend dev server
cd backend && go run ./cmd/server

# Build for production
make all
```

## 📄 License

MIT License
