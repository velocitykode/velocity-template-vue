<script setup lang="ts">
import { Link } from '@inertiajs/vue3';
import { computed } from 'vue';

interface Props {
    variant?: 'icon' | 'full' | 'responsive';
    size?: 'sm' | 'md' | 'lg' | 'xl';
    href?: string;
    onDarkSurface?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    variant: 'responsive',
    size: 'md',
    href: '/',
    onDarkSurface: false,
});

const sizeClasses = {
    sm: { icon: 'text-sm', full: 'text-base' },
    md: { icon: 'text-base', full: 'text-lg' },
    lg: { icon: 'text-lg', full: 'text-xl' },
    xl: { icon: 'text-xl', full: 'text-2xl' },
} as const;

const sizes = computed(() => sizeClasses[props.size]);
const fg = computed(() =>
    props.onDarkSurface ? 'text-panel-fg' : 'text-fg',
);
</script>

<template>
    <Link
        v-if="href"
        :href="href"
        class="inline-flex items-center"
    >
        <span
            v-if="variant === 'icon'"
            :class="[
                'font-mono font-semibold tracking-tight leading-none',
                sizes.icon,
                fg,
            ]"
        >
            v<span class="text-phosphor">_</span>
        </span>
        <span
            v-else-if="variant === 'full'"
            :class="[
                'font-mono font-semibold tracking-tight leading-none',
                sizes.full,
                fg,
            ]"
        >
            velocity<span class="text-phosphor">_</span>
        </span>
        <template v-else>
            <span class="lg:hidden">
                <span
                    :class="[
                        'font-mono font-semibold tracking-tight leading-none',
                        sizes.icon,
                        fg,
                    ]"
                >
                    v<span class="text-phosphor">_</span>
                </span>
            </span>
            <span class="hidden lg:inline">
                <span
                    :class="[
                        'font-mono font-semibold tracking-tight leading-none',
                        sizes.full,
                        fg,
                    ]"
                >
                    velocity<span class="text-phosphor">_</span>
                </span>
            </span>
        </template>
    </Link>
    <div v-else class="inline-flex items-center">
        <span
            v-if="variant === 'icon'"
            :class="[
                'font-mono font-semibold tracking-tight leading-none',
                sizes.icon,
                fg,
            ]"
        >
            v<span class="text-phosphor">_</span>
        </span>
        <span
            v-else-if="variant === 'full'"
            :class="[
                'font-mono font-semibold tracking-tight leading-none',
                sizes.full,
                fg,
            ]"
        >
            velocity<span class="text-phosphor">_</span>
        </span>
        <template v-else>
            <span class="lg:hidden">
                <span
                    :class="[
                        'font-mono font-semibold tracking-tight leading-none',
                        sizes.icon,
                        fg,
                    ]"
                >
                    v<span class="text-phosphor">_</span>
                </span>
            </span>
            <span class="hidden lg:inline">
                <span
                    :class="[
                        'font-mono font-semibold tracking-tight leading-none',
                        sizes.full,
                        fg,
                    ]"
                >
                    velocity<span class="text-phosphor">_</span>
                </span>
            </span>
        </template>
    </div>
</template>
