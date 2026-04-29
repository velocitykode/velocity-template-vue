import { usePage } from '@inertiajs/vue3';
import { watchEffect } from 'vue';

const appName = 'Velocity App';

export function usePageTitle(title?: string) {
    const page = usePage();

    watchEffect(() => {
        const pageTitle =
            title || (page.props as Record<string, unknown>).title || '';
        document.title = pageTitle ? `${pageTitle} - ${appName}` : appName;
    });
}
