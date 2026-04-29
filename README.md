# Velocity Template - Vue

Vue 3 + Inertia.js starter template for the [Velocity](https://github.com/velocitykode/velocity) Go web framework, with SSR enabled out of the box.

This repo is a **template** consumed by the Velocity installer. To start a new project:

```bash
velocity new myapp --stack=vue
```

The installer clones this template, rewrites the module placeholders, installs dependencies, builds the project's `vel` binary, and runs the initial migrations.

## Stack

- **Backend**: Velocity Go framework
- **Frontend**: Vue 3.5 + TypeScript 5 (Composition API, `<script setup>`)
- **Rendering**: Inertia.js 3 (`@inertiajs/vue3`) with SSR
- **Icons**: lucide-vue-next
- **Styling**: Tailwind CSS 4
- **Build**: Vite 7 (with `@velocitykode/velocity-vite-plugin`)

## SSR

SSR is wired up by default. After scaffolding, `npm run build` in your project produces:

- `public/build/` - client bundle
- `public/build/ssr/ssr.js` - SSR entry consumed by the Inertia SSR runtime

Toggle at runtime with `INERTIA_SSR_ENABLED` in the project's `.env`.

## Documentation

Full documentation at **[velocity.velocitykode.com/docs](https://velocity.velocitykode.com/docs)**

## Sibling Templates

- [`velocity-template-react`](https://github.com/velocitykode/velocity-template-react) - React 19 + Inertia.js
- [`velocity-template-api`](https://github.com/velocitykode/velocity-template-api) - API only (no frontend)

## License

MIT
