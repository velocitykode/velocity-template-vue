import { onScopeDispose, ref } from 'vue';

export type Appearance = 'light' | 'dark' | 'system';

const isClient = () => typeof window !== 'undefined';

const prefersDark = (): boolean => {
    if (!isClient()) return false;
    return window.matchMedia('(prefers-color-scheme: dark)').matches;
};

const setCookie = (name: string, value: string, days = 365): void => {
    if (typeof document === 'undefined') return;
    const maxAge = days * 24 * 60 * 60;
    document.cookie = `${name}=${value};path=/;max-age=${maxAge};SameSite=Lax`;
};

const applyTheme = (appearance: Appearance): void => {
    if (typeof document === 'undefined') return;
    const isDark =
        appearance === 'dark' ||
        (appearance === 'system' && prefersDark());
    document.documentElement.classList.toggle('dark', isDark);
    document.documentElement.style.colorScheme = isDark ? 'dark' : 'light';
};

const getMediaQuery = (): MediaQueryList | null => {
    if (!isClient()) return null;
    return window.matchMedia('(prefers-color-scheme: dark)');
};

const handleSystemThemeChange = () => {
    const saved = localStorage.getItem('appearance') as Appearance | null;
    applyTheme(saved ?? 'system');
};

export function initializeTheme(): void {
    if (!isClient()) return;
    const saved =
        (localStorage.getItem('appearance') as Appearance) ?? 'system';
    applyTheme(saved);
    getMediaQuery()?.addEventListener('change', handleSystemThemeChange);
}

export function useAppearance() {
    const appearance = ref<Appearance>('system');

    const updateAppearance = (mode: Appearance): void => {
        appearance.value = mode;
        if (isClient()) {
            localStorage.setItem('appearance', mode);
            setCookie('appearance', mode);
        }
        applyTheme(mode);
    };

    if (isClient()) {
        const saved = localStorage.getItem(
            'appearance',
        ) as Appearance | null;
        appearance.value = saved ?? 'system';
        applyTheme(appearance.value);
    }

    onScopeDispose(() => {
        getMediaQuery()?.removeEventListener('change', handleSystemThemeChange);
    });

    return { appearance, updateAppearance } as const;
}
