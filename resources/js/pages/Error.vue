<script setup lang="ts">
import { AppLogo, Button, ThemeToggle } from '@/components/ui';
import { usePageTitle } from '@/composables/use-page-title';
import { Link } from '@inertiajs/vue3';
import { ArrowLeft, House } from 'lucide-vue-next';

const props = defineProps<{ status: number; message: string }>();

usePageTitle(`${props.status} ${props.message}`);

function goBack() {
    window.history.back();
}
</script>

<template>
    <div class="flex min-h-screen flex-col bg-bg text-fg">
        <header
            class="flex h-12 items-center justify-between border-b border-border px-4 sm:px-5 lg:px-6"
        >
            <AppLogo variant="full" size="md" href="/" />
            <ThemeToggle />
        </header>

        <main
            class="flex flex-1 items-center justify-center px-4 py-8 sm:px-5 lg:px-6"
        >
            <div class="w-full max-w-3xl">
                <h1
                    class="font-sans text-6xl font-semibold leading-none tracking-tight sm:text-7xl lg:text-8xl"
                >
                    {{ status }}
                </h1>
                <p
                    class="mt-4 max-w-xl text-[15px] leading-relaxed text-muted lg:text-[17px]"
                >
                    {{ message }}
                </p>

                <div class="mt-6 flex flex-col gap-2.5 sm:flex-row">
                    <Link href="/">
                        <Button size="lg" :icon-left="House" full-width>
                            Go home
                        </Button>
                    </Link>
                    <Button
                        variant="secondary"
                        size="lg"
                        :icon-left="ArrowLeft"
                        @click="goBack"
                    >
                        Go back
                    </Button>
                </div>
            </div>
        </main>
    </div>
</template>
