<script setup lang="ts">
import { Link } from '@inertiajs/vue3';
import { ExternalLink } from 'lucide-vue-next';
import type { FunctionalComponent } from 'vue';

interface Props {
    name: string;
    href: string;
    icon: FunctionalComponent;
    active?: boolean;
    external?: boolean;
}

withDefaults(defineProps<Props>(), {
    active: false,
    external: false,
});

const baseClasses =
    'group flex items-center gap-3 rounded-none px-3 py-2.5 text-sm transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg';
</script>

<template>
    <a
        v-if="external"
        :href="href"
        target="_blank"
        rel="noopener noreferrer"
        :class="[
            baseClasses,
            active
                ? 'border-l-2 border-accent pl-[10px] text-fg font-medium'
                : 'border-l-2 border-transparent pl-[10px] text-muted hover:text-fg',
        ]"
    >
        <component :is="icon" class="h-4 w-4" />
        <span class="flex-1 truncate">{{ name }}</span>
        <ExternalLink class="h-3.5 w-3.5 opacity-50" />
    </a>
    <Link
        v-else
        :href="href"
        :class="[
            baseClasses,
            active
                ? 'border-l-2 border-accent pl-[10px] text-fg font-medium'
                : 'border-l-2 border-transparent pl-[10px] text-muted hover:text-fg',
        ]"
    >
        <component :is="icon" class="h-4 w-4" />
        <span class="flex-1 truncate">{{ name }}</span>
    </Link>
</template>
