# Velocity Template - Vue

Full-stack starter template for the [Velocity](https://github.com/velocitykode/velocity) Go web framework using Vue 3 + Inertia.js, with SSR support.

## Stack

- **Backend**: Velocity Go framework
- **Frontend**: Vue 3.5 + TypeScript 5 (Composition API, `<script setup>`)
- **Rendering**: Inertia.js 3 (`@inertiajs/vue3`) with SSR
- **Icons**: lucide-vue-next
- **Styling**: Tailwind CSS 4
- **Build**: Vite 7 (with `@velocitykode/velocity-vite-plugin`)

## Usage

This template is used automatically by the Velocity CLI:

```bash
velocity new myapp --stack=vue
cd myapp
./vel serve
```

## Development Commands

```bash
# Start dev server with hot reload
./vel serve

# Run database migrations
./vel migrate

# Generate a new controller
./vel make:controller User

# Build for production (client + SSR bundle)
npm run build

# Client-only build
npm run build:client

# SSR-only build
npm run build:ssr
```

## SSR

SSR is enabled out of the box. `npm run build` produces:

- `public/build/` - client bundle
- `public/build/ssr/ssr.js` - SSR entry consumed by the Inertia SSR runtime

## Documentation

Full documentation at **[velocity.velocitykode.com/docs](https://velocity.velocitykode.com/docs)**

## License

MIT
