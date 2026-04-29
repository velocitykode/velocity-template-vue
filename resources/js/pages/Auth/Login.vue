<script setup lang="ts">
import { Button, FormField, Input } from '@/components/ui';
import { usePageTitle } from '@/composables/use-page-title';
import AuthLayout from '@/layouts/AuthLayout.vue';
import { Link, useForm } from '@inertiajs/vue3';
import { ArrowRight, Eye, EyeOff } from 'lucide-vue-next';
import { ref } from 'vue';

interface Props {
    status?: string;
    canResetPassword?: boolean;
    errors?: {
        email?: string;
        password?: string;
    };
    old?: {
        email?: string;
    };
}

const props = withDefaults(defineProps<Props>(), {
    canResetPassword: false,
    errors: () => ({}),
    old: () => ({}),
});

usePageTitle('Sign in');
const showPassword = ref(false);

const form = useForm({
    email: props.old.email || '',
    password: '',
    remember: false,
});

const handleSubmit = () => {
    form.post('/login');
};
</script>

<template>
    <AuthLayout
        title="Sign in"
        description="Welcome back - manage your application and continue where you left off."
        switch-prompt="New to Velocity?"
        switch-href="/register"
        switch-label="Create an account"
    >
        <div
            v-if="status"
            class="mb-6 border border-emerald-600/40 bg-emerald-600/10 px-4 py-3"
        >
            <p
                class="caption-mono text-emerald-700 dark:text-emerald-400"
            >
                {{ status }}
            </p>
        </div>

        <form class="space-y-4" @submit.prevent="handleSubmit">
            <FormField
                label="Email"
                html-for="email"
                :error="errors.email"
            >
                <Input
                    id="email"
                    v-model="form.email"
                    type="email"
                    placeholder="you@example.com"
                    autocomplete="email"
                    autofocus
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
                        autocomplete="current-password"
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

            <div class="flex items-center justify-between">
                <label class="flex cursor-pointer items-center gap-2.5">
                    <input
                        v-model="form.remember"
                        type="checkbox"
                        class="h-4 w-4 rounded-none border-border-strong bg-surface accent-fg"
                    />
                    <span class="caption-mono text-muted">Remember me</span>
                </label>

                <Link
                    v-if="canResetPassword"
                    href="/password/request"
                    class="link-underline caption-mono text-muted"
                >
                    Trouble logging in?
                </Link>
            </div>

            <Button
                type="submit"
                full-width
                size="lg"
                :loading="form.processing"
                :icon-right="ArrowRight"
                class="mt-2"
            >
                Sign in
            </Button>
        </form>
    </AuthLayout>
</template>
