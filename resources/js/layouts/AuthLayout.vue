<script setup lang="ts">
import { AppLogo, ThemeToggle } from '@/components/ui';
import { Link } from '@inertiajs/vue3';
import { computed } from 'vue';

interface Props {
    title: string;
    description: string;
    switchPrompt?: string;
    switchHref?: string;
    switchLabel?: string;
}

defineProps<Props>();

const year = computed(() => new Date().getFullYear());
</script>

<template>
    <div class="flex min-h-screen flex-col bg-bg lg:flex-row">
        <aside
            class="relative hidden overflow-hidden bg-panel text-panel-fg lg:flex lg:w-1/2 lg:flex-col lg:justify-between xl:w-[45%] 2xl:w-[40%]"
        >
            <div class="absolute inset-0 bg-paper-grid" />

            <div
                class="relative flex items-start justify-between px-6 py-5 xl:px-8 xl:py-6"
            >
                <AppLogo
                    variant="full"
                    size="xl"
                    href="/"
                    on-dark-surface
                />
                <ThemeToggle on-dark-surface />
            </div>

            <div
                class="relative flex flex-1 items-end px-6 pb-6 xl:px-8 xl:pb-8"
            >
                <div class="max-w-md">
                    <p class="caption-mono-upper mb-3 text-panel-muted">
                        Velocity · v0.22.0
                    </p>
                    <h2
                        class="font-sans text-4xl font-semibold leading-[1.05] tracking-tight text-panel-fg xl:text-5xl 2xl:text-[56px]"
                    >
                        Build Faster.
                        <br />
                        Ship Sooner.
                    </h2>
                    <p
                        class="caption-mono mt-4 max-w-sm text-panel-muted"
                    >
                        Batteries-included Go framework - routing, auth, db,
                        queues, websockets. Twenty-three services, zero
                        yak-shaving.
                    </p>
                </div>
            </div>

            <div
                class="relative border-t border-panel-border px-6 py-4 xl:px-8"
            >
                <div
                    class="caption-mono-upper flex items-center justify-between text-panel-dim"
                >
                    <span>© {{ year }} Velocity</span>
                    <span>go/1.26+</span>
                </div>
            </div>
        </aside>

        <main class="flex flex-1 flex-col bg-bg">
            <div
                class="flex h-14 items-center justify-between border-b border-border px-4 lg:hidden"
            >
                <AppLogo variant="icon" size="lg" href="/" />
                <ThemeToggle />
            </div>

            <div
                class="flex flex-1 items-center justify-center px-5 py-8 sm:px-8 lg:px-10 xl:px-16 2xl:px-24"
            >
                <div class="w-full max-w-[420px] xl:max-w-[440px]">
                    <div class="mb-6 lg:mb-8">
                        <h1
                            class="font-sans text-3xl font-semibold tracking-tight text-fg lg:text-[34px] xl:text-[38px]"
                        >
                            {{ title }}
                        </h1>
                        <p
                            class="mt-2 text-[14px] leading-relaxed text-muted"
                        >
                            {{ description }}
                        </p>
                    </div>

                    <slot />

                    <p
                        v-if="switchPrompt && switchHref && switchLabel"
                        class="caption-mono mt-6 text-muted"
                    >
                        {{ switchPrompt }}
                        <Link
                            :href="switchHref"
                            class="link-underline text-fg"
                        >
                            {{ switchLabel }}
                        </Link>
                    </p>
                </div>
            </div>

            <div
                class="border-t border-border px-5 py-4 sm:px-8 lg:px-10 xl:px-16 2xl:px-24"
            >
                <div
                    class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"
                >
                    <p class="caption-mono text-dim">
                        By continuing, you agree to our
                        <a href="#" class="link-underline text-muted">
                            Terms
                        </a>
                        and
                        <a href="#" class="link-underline text-muted">
                            Privacy Policy
                        </a>
                        .
                    </p>
                    <p class="caption-mono-upper text-dim">
                        © {{ year }} Velocity
                    </p>
                </div>
            </div>
        </main>
    </div>
</template>
