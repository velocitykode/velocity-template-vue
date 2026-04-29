<script setup lang="ts">
import { useAppearance } from '@/composables/use-appearance';
import { computed } from 'vue';

interface Props {
    onDarkSurface?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    onDarkSurface: false,
});

const { appearance, updateAppearance } = useAppearance();

const effective = computed<'light' | 'dark'>(() => {
    if (appearance.value === 'dark') return 'dark';
    if (appearance.value === 'light') return 'light';
    return typeof window !== 'undefined' &&
        window.matchMedia('(prefers-color-scheme: dark)').matches
        ? 'dark'
        : 'light';
});

const tokens = computed(() =>
    props.onDarkSurface
        ? {
              border: 'border-panel-border',
              muted: 'text-panel-dim',
              active: 'text-panel-fg',
              sep: 'text-panel-dim',
          }
        : {
              border: 'border-border',
              muted: 'text-dim',
              active: 'text-fg',
              sep: 'text-dim',
          },
);
</script>

<template>
    <button
        type="button"
        :aria-label="`Switch to ${effective === 'dark' ? 'light' : 'dark'} theme`"
        :aria-pressed="effective === 'light'"
        :class="[
            'inline-flex h-9 items-center gap-1.5 border px-3 font-mono text-[11px] uppercase tracking-widest rounded-none transition-colors',
            tokens.border,
        ]"
        @click="updateAppearance(effective === 'dark' ? 'light' : 'dark')"
    >
        <span :class="tokens.sep">[</span>
        <span :class="effective === 'dark' ? tokens.active : tokens.muted">
            dark
        </span>
        <span :class="tokens.sep">|</span>
        <span :class="effective === 'light' ? tokens.active : tokens.muted">
            light
        </span>
        <span :class="tokens.sep">]</span>
    </button>
</template>
