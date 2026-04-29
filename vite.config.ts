import inertia from '@inertiajs/vite';
import tailwindcss from '@tailwindcss/vite';
import vue from '@vitejs/plugin-vue';
import { defineConfig } from 'vite';
import velocity from '@velocitykode/velocity-vite-plugin';

export default defineConfig({
    plugins: [
        velocity('resources/js/app.ts'),
        inertia({
            pages: 'resources/js/pages',
        }),
        vue(),
        tailwindcss(),
    ],
    server: {
        port: 5173,
        strictPort: true,
        host: 'localhost',
    },
});
