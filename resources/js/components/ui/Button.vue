<script setup lang="ts">
import type { FunctionalComponent } from 'vue';

interface Props {
    variant?: 'primary' | 'secondary' | 'ghost' | 'danger';
    size?: 'sm' | 'md' | 'lg';
    iconLeft?: FunctionalComponent;
    iconRight?: FunctionalComponent;
    loading?: boolean;
    fullWidth?: boolean;
    disabled?: boolean;
    type?: 'button' | 'submit' | 'reset';
}

const props = withDefaults(defineProps<Props>(), {
    variant: 'primary',
    size: 'md',
    loading: false,
    fullWidth: false,
    disabled: false,
    type: 'button',
});

const variantClasses = {
    primary: 'bg-accent text-accent-fg hover:opacity-90',
    secondary:
        'border border-border-strong bg-transparent text-fg hover:bg-surface-alt',
    ghost: 'text-muted hover:text-fg hover:bg-surface-alt',
    danger: 'bg-error text-white hover:opacity-90',
} as const;

const sizeClasses = {
    sm: 'h-8 px-3.5 text-xs gap-1.5',
    md: 'h-10 px-4 text-sm gap-2',
    lg: 'h-12 px-6 text-[14px] gap-2',
} as const;

const iconSizeClasses = {
    sm: 'h-3.5 w-3.5',
    md: 'h-4 w-4',
    lg: 'h-5 w-5',
} as const;
</script>

<template>
    <button
        :type="type"
        :disabled="disabled || loading"
        :class="[
            'inline-flex items-center justify-center rounded-none font-medium tracking-wide transition-[background-color,color,opacity] duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg disabled:cursor-not-allowed disabled:opacity-50',
            variantClasses[variant],
            sizeClasses[size],
            fullWidth ? 'w-full' : '',
        ]"
    >
        <component
            v-if="iconLeft"
            :is="iconLeft"
            :class="iconSizeClasses[size]"
        />
        <slot />
        <span
            v-if="loading || iconRight"
            :class="[
                'inline-flex items-center justify-center',
                iconSizeClasses[size],
            ]"
        >
            <span v-if="loading" class="ascii-spinner" aria-hidden />
            <component
                v-else-if="iconRight"
                :is="iconRight"
                :class="iconSizeClasses[size]"
            />
        </span>
    </button>
</template>
