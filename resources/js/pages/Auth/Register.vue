<script setup lang="ts">
import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/composables/use-page-title';
import AuthLayout from '@/layouts/AuthLayout.vue';
import { useForm } from '@inertiajs/vue3';
import { ArrowRight, Eye, EyeOff } from 'lucide-vue-next';
import { ref } from 'vue';

interface Props {
    errors?: {
        name?: string;
        email?: string;
        password?: string;
        password_confirmation?: string;
    };
    old?: {
        name?: string;
        email?: string;
    };
}

const props = withDefaults(defineProps<Props>(), {
    errors: () => ({}),
    old: () => ({}),
});

usePageTitle('Create account');
const showPassword = ref(false);

const form = useForm({
    name: props.old.name || '',
    email: props.old.email || '',
    password: '',
    password_confirmation: '',
});

const handleSubmit = () => {
    form.post('/register');
};
</script>

<template>
    <AuthLayout
        title="Create account"
        description="Get started with Velocity in seconds - no credit card, no ceremony."
        switch-prompt="Already have an account?"
        switch-href="/login"
        switch-label="Sign in"
    >
        <form class="space-y-4" @submit.prevent="handleSubmit">
            <FormField label="Name" html-for="name" :error="errors.name">
                <Input
                    id="name"
                    v-model="form.name"
                    type="text"
                    placeholder="Your name"
                    autocomplete="name"
                    autofocus
                    required
                    :error="!!errors.name"
                />
            </FormField>

            <FormField label="Email" html-for="email" :error="errors.email">
                <Input
                    id="email"
                    v-model="form.email"
                    type="email"
                    placeholder="you@example.com"
                    autocomplete="email"
                    required
                    :error="!!errors.email"
                />
            </FormField>

            <FormField
                label="Password"
                html-for="password"
                :error="errors.password"
            >
                <div class="relative">
                    <Input
                        id="password"
                        v-model="form.password"
                        :type="showPassword ? 'text' : 'password'"
                        autocomplete="new-password"
                        required
                        :error="!!errors.password"
                    />
                    <button
                        type="button"
                        :tabindex="-1"
                        :aria-label="
                            showPassword ? 'Hide password' : 'Show password'
                        "
                        class="absolute right-3 top-1/2 -translate-y-1/2 text-muted transition-colors hover:text-fg"
                        @click="showPassword = !showPassword"
                    >
                        <EyeOff v-if="showPassword" class="h-4 w-4" />
                        <Eye v-else class="h-4 w-4" />
                    </button>
                </div>
            </FormField>

            <FormField
                label="Confirm password"
                html-for="password_confirmation"
                :error="errors.password_confirmation"
            >
                <Input
                    id="password_confirmation"
                    v-model="form.password_confirmation"
                    :type="showPassword ? 'text' : 'password'"
                    autocomplete="new-password"
                    required
                    :error="!!errors.password_confirmation"
                />
            </FormField>

            <Button
                type="submit"
                full-width
                size="lg"
                :loading="form.processing"
                :icon-right="ArrowRight"
                class="mt-2"
            >
                Create account
            </Button>
        </form>
    </AuthLayout>
</template>
