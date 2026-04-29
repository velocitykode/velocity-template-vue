import { onScopeDispose, ref } from 'vue';

const MOBILE_BREAKPOINT = 768;

export function useIsMobile() {
    const isMobile = ref(false);

    if (typeof window === 'undefined') {
        return isMobile;
    }

    const mql = window.matchMedia(
        `(max-width: ${MOBILE_BREAKPOINT - 1}px)`,
    );

    const update = () => {
        isMobile.value = window.innerWidth < MOBILE_BREAKPOINT;
    };

    update();
    mql.addEventListener('change', update);

    onScopeDispose(() => {
        mql.removeEventListener('change', update);
    });

    return isMobile;
}
