import '../css/app.css';

import { createInertiaApp, Link, router } from '@inertiajs/vue3';
import { createApp, h, type DefineComponent } from 'vue';
import { initializeTheme } from './composables/use-appearance';

let csrfToken: string | null = null;

router.on('navigate', (event) => {
    const props = event.detail.page.props as { csrf_token?: string };
    if (props.csrf_token) {
        csrfToken = props.csrf_token;
        const meta = document.querySelector('meta[name="csrf-token"]');
        if (meta) {
            meta.setAttribute('content', csrfToken);
        }
    }
});

router.on('before', (event) => {
    if (!csrfToken) {
        csrfToken = readInitialCsrfToken();
    }
    if (csrfToken && event.detail.visit?.headers) {
        event.detail.visit.headers['X-CSRF-Token'] = csrfToken;
    }
});

createInertiaApp({
    progress: { color: '#4B5563' },
    resolve: async (name) => {
        const pages = import.meta.glob<DefineComponent>('./pages/**/*.vue');
        const path = `./pages/${name}.vue`;
        const page = pages[path];
        if (!page) {
            throw new Error(`Page not found: ${path}`);
        }
        return await page();
    },
    setup({ el, App, props, plugin }) {
        createApp({ render: () => h(App, props) })
            .use(plugin)
            .component('Link', Link)
            .mount(el);
    },
});

initializeTheme();

function readInitialCsrfToken(): string | null {
    if (typeof document === 'undefined') {
        return null;
    }
    const el = document.getElementById('app');
    const raw = el?.dataset.page;
    if (!raw) return null;
    try {
        const page = JSON.parse(raw) as {
            props?: { csrf_token?: string };
        };
        return page.props?.csrf_token ?? null;
    } catch {
        return null;
    }
}
