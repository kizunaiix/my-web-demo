# AGENTS.md

This file contains guidelines and commands for agentic coding agents working in this repository.

## Project Overview

This is a habit tracking/check-in application with:
- **Frontend**: React with TypeScript, Vite, Tailwind CSS, Base UI components
- **Backend**: Go with Gin framework, GORM, Swagger documentation
- **API Style**: RPC-style using only GET and POST methods (no PUT/DELETE)
- **Database**: PostgreSQL (configured but not actively used in current state)

## Build Commands

### Frontend (React/TypeScript)
```bash
cd frontend
npm install              # Install dependencies
npm run dev             # Start development server
npm run build           # Build for production
npm run lint            # Run ESLint
npm run preview         # Preview production build
```

### Backend (Go/Gin)
```bash
# Using Task (recommended)
task swag:init          # Initialize Swagger documentation
task go:run-dev         # Run in development mode
task go:run-prod        # Run in production mode
task go:build           # Build binary to backend/bin/
task go:test            # Run Go tests

# Direct Go commands
cd backend/cmd
swag init -d ./,../internal/  # Generate Swagger docs
go run .                 # Run application
go build -o ../bin/my-web-demo .  # Build binary
go test ./...           # Run all tests
go test -v ./test       # Run specific test with verbose output
go test -run TestSpecific ./test  # Run specific test function
```

### Docker Commands
```bash
task docker:build       # Build Docker image
task docker:run-dev     # Run container in development mode
task docker:run-prod    # Run container in production mode
```

## Code Style Guidelines

### Frontend (React/TypeScript)

#### Imports
- Use absolute imports with `@/` alias for src directory
- Group imports: React libraries first, then local components, then UI components
```typescript
import * as React from "react"
import { ComponentExample } from "@/components/component-example"
import { Button } from "@/components/ui/button"
import { PlusIcon } from "lucide-react"
```

#### Component Structure
- Use function components with React hooks
- Export components separately: `export function ComponentName() {}`
- Add `export default ComponentName` at the end
- Use TypeScript interfaces for props

#### Styling
- Use Tailwind CSS classes exclusively
- Follow Base UI component patterns from shadcn/ui
- Use `cn()` utility for conditional classes
- Component variants use class-variance-authority

#### State Management
- Use `React.useState` for local state
- Use `const [state, setState] = React.useState<Type>(initialValue)`
- Prefer explicit typing for state

#### File Naming
- Component files: `component-name.tsx`
- Use kebab-case for file names
- Use PascalCase for component names

### Backend (Go)

#### Package Structure
- `cmd/` - Application entry points
- `internal/` - Private application code
  - `config/` - Configuration
  - `dto/` - Data transfer objects
  - `middleware/` - HTTP middleware
  - `server/` - Server setup
  - `{feature}/` - Feature-specific packages (handler, service, repository, model)
- `pkg/` - Public library code
- `test/` - Test files

#### Import Organization
```go
import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "ki9.com/gin_demo/internal/dto"
    "ki9.com/gin_demo/internal/middleware/myerr"
)
```

#### Naming Conventions
- Package names: lowercase, single word when possible
- Struct names: PascalCase (e.g., `TaskHandler`)
- Interface names: PascalCase, often with `-er` suffix
- Function names: PascalCase for exported, camelCase for unexported
- Variable names: camelCase
- Constants: UPPER_SNAKE_CASE

#### Error Handling
- Use custom error types from `internal/middleware/myerr`
- Return errors from functions, don't panic unless unrecoverable
- Use `ctx.Error()` for Gin context errors
- Log errors with appropriate levels using zap logger

#### API Handlers
- Use Swagger comments for all API endpoints
- Follow pattern: `func (h *Handler) HandlerName(ctx *gin.Context)`
- Use `ctx.ShouldBindJSON()` instead of `ctx.BindJSON()` for custom error handling
- Return standardized responses using `dto.UniResponseBody`

#### Testing
- Test files: `feature_test.go` in `test/` package or alongside source
- Test functions: `func TestFeature(t *testing.T)`
- Use subtests with `t.Run()` for multiple scenarios
- Table-driven tests for multiple test cases

## Development Workflow

### Adding New Features
1. **Backend**: Create handler, service, repository, model in feature package
2. **Frontend**: Create components in `src/components/` or `src/pages/`
3. **API**: Add Swagger comments and run `task swag:init`
4. **Testing**: Write tests for both backend and frontend
5. **Linting**: Run `npm run lint` (frontend) and ensure Go code passes `go vet`

### Code Quality
- Frontend: ESLint configuration enforces React and TypeScript best practices
- Backend: Use `go vet`, `go fmt`, and proper error handling
- All code should be type-safe (TypeScript strict mode, Go strong typing)

### File Organization
- Keep related files together in feature directories
- Use index files for re-exports when appropriate
- Separate UI components from business logic components

## Running Single Tests

### Frontend
```bash
cd frontend
npm test  # If test framework is added
```

### Backend
```bash
cd backend
go test -v ./test                    # Run all tests in test package
go test -v -run TestTimeInit ./test  # Run specific test
go test -run TestUuid ./test         # Run specific test without verbose
go test ./...                        # Run all tests in project
```

## Environment Variables

### Backend
- `ENV`: Set to "dev" or "prod" for different configurations
- `GIN_MODE`: Set to "release" for production

### Frontend
- Uses Vite for environment configuration
- Base path set to "/my-web-demo" in vite.config.ts

## Notes

- The Taskfile.yml contains outdated paths that need updating (react → frontend, gin → backend)
- Swagger documentation is generated from comments and available in dev mode
- The project uses RPC-style API design with method field in request bodies
- Frontend uses Base UI components with custom styling via Tailwind CSS