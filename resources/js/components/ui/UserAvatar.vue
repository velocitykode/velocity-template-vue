<script setup lang="ts">
import { computed } from 'vue';

interface Props {
    name?: string;
    size?: 'sm' | 'md' | 'lg';
    src?: string;
}

const props = withDefaults(defineProps<Props>(), {
    size: 'md',
});

const sizeClasses = {
    sm: 'h-8 w-8 text-[11px]',
    md: 'h-10 w-10 text-xs',
    lg: 'h-12 w-12 text-sm',
} as const;

const initials = computed(() => {
    const name = props.name;
    if (!name) return 'U';
    const parts = name.trim().split(' ');
    if (parts.length === 1) return parts[0].charAt(0).toUpperCase();
    return (
        parts[0].charAt(0) + parts[parts.length - 1].charAt(0)
    ).toUpperCase();
});
</script>

<template>
    <img
        v-if="src"
        :src="src"
        :alt="name || 'User avatar'"
        :class="['rounded-none object-cover', sizeClasses[size]]"
    />
    <div
        v-else
        :class="[
            'flex items-center justify-center border border-border-strong bg-surface font-mono tracking-widest text-fg',
            sizeClasses[size],
        ]"
    >
        {{ initials }}
    </div>
</template>
