<script setup lang="ts">
import type { FunctionalComponent } from 'vue';

interface Props {
    label: string;
    value: string | number;
    icon: FunctionalComponent;
    color?: 'blue' | 'violet' | 'emerald' | 'amber' | 'rose';
    change?: string;
    status?: string;
}

defineProps<Props>();
</script>

<template>
    <div
        class="border border-border bg-surface p-4 transition-colors duration-200 hover:border-border-strong lg:p-5"
    >
        <div class="flex items-start justify-between">
            <span class="caption-mono-upper text-dim">{{ label }}</span>
            <component :is="icon" class="h-4 w-4 text-dim" />
        </div>
        <div class="mt-4 flex items-baseline gap-2 lg:mt-5">
            <span
                class="text-2xl font-semibold tracking-tight text-fg lg:text-3xl xl:text-[32px]"
            >
                {{ value }}
            </span>
            <span v-if="change" class="caption-mono text-muted">
                {{ change }}
            </span>
        </div>
    </div>
</template>
