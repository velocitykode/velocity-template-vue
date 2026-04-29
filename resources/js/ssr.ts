import { createInertiaApp } from '@inertiajs/vue3';
import createServer from '@inertiajs/vue3/server';
import type { Page } from '@inertiajs/core';
import { createSSRApp, h, type DefineComponent } from 'vue';
import { renderToString } from 'vue/server-renderer';

createServer((page: Page) =>
    createInertiaApp({
        page,
        render: renderToString,
        resolve: (name) => {
            const pages = import.meta.glob<DefineComponent>(
                './pages/**/*.vue',
                { eager: true },
            );
            const path = `./pages/${name}.vue`;
            const component = pages[path];
            if (!component) {
                throw new Error(`Page not found: ${path}`);
            }
            return component;
        },
        setup({ App, props, plugin }) {
            return createSSRApp({ render: () => h(App, props) }).use(
                plugin,
            );
        },
    }),
);
