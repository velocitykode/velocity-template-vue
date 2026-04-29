<script setup lang="ts">
import type { FunctionalComponent } from 'vue';

interface Props {
    icon: FunctionalComponent;
    size?: 'sm' | 'md' | 'lg';
    variant?: 'default' | 'danger' | 'ghost';
    as?: 'button' | 'a';
    href?: string;
    type?: 'button' | 'submit' | 'reset';
    disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    size: 'md',
    variant: 'default',
    as: 'button',
    type: 'button',
});

const sizeClasses = {
    sm: 'h-8 w-8',
    md: 'h-10 w-10',
    lg: 'h-11 w-11',
} as const;

const iconSizeClasses = {
    sm: 'h-3.5 w-3.5',
    md: 'h-4 w-4',
    lg: 'h-5 w-5',
} as const;

const variantClasses = {
    default:
        'border border-border bg-transparent text-muted hover:border-fg hover:text-fg',
    danger:
        'border border-border bg-transparent text-muted hover:border-error hover:text-error',
    ghost: 'bg-transparent text-muted hover:text-fg',
} as const;
</script>

<template>
    <a
        v-if="as === 'a'"
        :href="href"
        :class="[
            'flex items-center justify-center rounded-none transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg',
            sizeClasses[size],
            variantClasses[variant],
        ]"
    >
        <component :is="icon" :class="iconSizeClasses[size]" />
    </a>
    <button
        v-else
        :type="type"
        :disabled="disabled"
        :class="[
            'flex items-center justify-center rounded-none transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg',
            sizeClasses[size],
            variantClasses[variant],
        ]"
    >
        <component :is="icon" :class="iconSizeClasses[size]" />
    </button>
</template>
