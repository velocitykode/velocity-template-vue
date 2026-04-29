<script setup lang="ts">
const { error = false, type = 'text' } = defineProps<{
    error?: boolean;
    type?: string;
}>();

const model = defineModel<string | number>();
</script>

<template>
    <input
        v-model="model"
        :type="type"
        :class="[
            'h-11 w-full rounded-none border bg-surface px-3.5 text-[14px] text-fg placeholder:text-dim transition-colors duration-150 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 lg:h-12',
            error
                ? 'border-error text-error placeholder:text-error focus:border-error'
                : 'border-border focus:border-fg',
        ]"
    />
</template>
