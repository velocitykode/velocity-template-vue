<script setup lang="ts">
interface Props {
    label?: string;
    error?: string;
    hint?: string;
    required?: boolean;
    htmlFor?: string;
}

withDefaults(defineProps<Props>(), {
    required: false,
});
</script>

<template>
    <div class="space-y-1.5">
        <label
            v-if="label"
            :for="htmlFor"
            class="caption-mono-upper block text-muted"
        >
            {{ label }}
            <span v-if="required" class="ml-1 text-error">*</span>
        </label>
        <slot />
        <p v-if="error" class="caption-mono text-error">{{ error }}</p>
        <p v-else-if="hint" class="caption-mono text-dim">{{ hint }}</p>
    </div>
</template>
