# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is the default project template for the Velocity Go web framework. It provides a full-stack application with a Go backend and React frontend connected via Inertia.js.

**Stack**: Go (Velocity framework) + React 19 + TypeScript + Inertia.js + Tailwind CSS + Vite

## Development Commands

```bash
# Start development server (Go with hot reload + Vite)
./vel serve

# Build for production (TypeScript check + Vite bundle)
npm run build

# Run database migrations
./vel migrate

# Generate a new controller
./vel make:controller <Name>

# Frontend dev server only (port 5173)
npm run dev
```

## Architecture

### Request Flow

```
Request → Global Middleware → Web/API Middleware → Route Handler → Inertia → React Page
```

### Backend Structure (`internal/`)

- **`app/`** - Application bootstrap and HTTP kernel
  - `bootstrap.go` - Initializes crypto, database, auth, CSRF, views
  - `kernel.go` - HTTP lifecycle and route loading
  - `middleware.go` - Middleware stack definitions (Global, Web, API)
- **`handlers/`** - Route handlers that render Inertia pages
- **`middleware/`** - Custom middleware implementations
- **`models/`** - Database models

### Frontend Structure (`resources/js/`)

- **`pages/`** - Inertia page components (mapped to server routes)
- **`components/ui/`** - Reusable UI components (Radix-based)
- **`layouts/`** - Shared layout components
- **`hooks/`** - Custom React hooks (appearance, mobile detection, etc.)

### Route Registration

Routes are registered in `routes/web.go` using init-based loading:

```go
router.Register(func(r router.Router) {
    r.Group("", func(guest router.Router) {
        guest.Get("/login", handlers.AuthShowLoginForm)
    }).Use(middleware.Guest)
})
```

### Middleware Stacks

Three middleware tiers defined in `internal/app/middleware.go`:

- **Global**: Recovery, logging, CORS, maintenance mode, request validation
- **Web**: Session, CSRF, Inertia (for browser requests)
- **API**: JSON response formatting

### Rendering Pages

Handlers use `view.Render` to pass props to React components:

```go
view.Render(ctx.Response, ctx.Request, "Dashboard", view.Props{
    "auth": map[string]interface{}{"user": userMap},
})
```

The component name ("Dashboard") maps to `resources/js/pages/Dashboard.tsx`.

### Configuration

- **Environment**: Copy `.env.example` to `.env`, configure database and app settings
- **Hot reload**: Backend uses Air (`.air.toml`), frontend uses Vite
- **Static files**: Served from `public/`, Vite outputs to `public/build/`

## Key Patterns

- **Auth guards**: `middleware.Auth` and `middleware.Guest` protect route groups
- **CSRF**: Automatically handled via session storage and Inertia headers
- **Props passing**: Server data flows to React via Inertia props (includes `auth.user`, `csrf_token`)
- **Path alias**: Use `@/` in imports to reference `resources/js/`
