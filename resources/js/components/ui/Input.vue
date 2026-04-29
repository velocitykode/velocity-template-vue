<script setup lang="ts">
const {
    error = false,
    type = 'text',
    modelValue,
} = defineProps<{
    error?: boolean;
    type?: string;
    modelValue?: string | number;
}>();

const emit = defineEmits<{
    'update:modelValue': [value: string | number];
}>();

const handleInput = (event: Event) => {
    emit('update:modelValue', (event.target as HTMLInputElement).value);
};
</script>

<template>
    <input
        :type="type"
        :value="modelValue"
        :class="[
            'h-11 w-full rounded-none border bg-surface px-3.5 text-[14px] text-fg placeholder:text-dim transition-colors duration-150 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 lg:h-12',
            error
                ? 'border-error text-error placeholder:text-error focus:border-error'
                : 'border-border focus:border-fg',
        ]"
        @input="handleInput"
    />
</template>
