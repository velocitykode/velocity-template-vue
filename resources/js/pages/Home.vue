<script setup lang="ts">
import { AppLogo, Badge, Button, ThemeToggle } from '@/components/ui';
import { usePageTitle } from '@/composables/use-page-title';
import { Link } from '@inertiajs/vue3';
import { ArrowRight, BookOpen, Github } from 'lucide-vue-next';

interface Props {
    message?: string;
}

defineProps<Props>();

usePageTitle('Welcome');

const year = new Date().getFullYear();
</script>

<template>
    <div class="flex min-h-screen flex-col bg-bg text-fg">
        <header
            class="flex h-12 items-center justify-between border-b border-border px-4 sm:px-5 lg:px-6"
        >
            <AppLogo variant="full" size="md" href="/" />
            <div class="flex items-center gap-2">
                <ThemeToggle />
                <Link
                    href="/login"
                    class="caption-mono-upper px-2 text-muted transition-colors hover:text-fg"
                >
                    Sign in
                </Link>
            </div>
        </header>

        <main
            class="flex flex-1 items-center justify-center px-4 py-8 sm:px-5 lg:px-6"
        >
            <div class="w-full max-w-3xl">
                <Badge variant="outline" class="mb-4">Ready to ship</Badge>
                <h1
                    class="font-sans text-4xl font-semibold leading-[1.05] tracking-tight sm:text-5xl lg:text-6xl xl:text-[68px]"
                >
                    Welcome to Velocity.
                    <br />
                    <span class="text-muted">
                        {{ message || 'Build faster with pure Go power.' }}
                    </span>
                </h1>
                <p
                    class="mt-4 max-w-xl text-[15px] leading-relaxed text-muted lg:text-[17px]"
                >
                    A batteries-included full-stack framework for Go - routing,
                    auth, db, queues, websockets. Twenty-three services, zero
                    yak-shaving.
                </p>

                <div class="mt-6 flex flex-col gap-2.5 sm:flex-row">
                    <Link href="/register">
                        <Button
                            size="lg"
                            :icon-right="ArrowRight"
                            full-width
                        >
                            Get started
                        </Button>
                    </Link>
                    <a
                        href="https://github.com/velocitykode/velocity"
                        target="_blank"
                        rel="noopener noreferrer"
                        class="inline-flex h-12 items-center justify-center gap-2 border border-border-strong px-5 text-[14px] tracking-wide text-fg transition-colors hover:bg-surface-alt"
                    >
                        <Github class="h-4 w-4" />
                        Star on GitHub
                    </a>
                </div>

                <div
                    class="mt-10 grid gap-px border border-border bg-border sm:grid-cols-2"
                >
                    <a
                        href="https://github.com/velocitykode/velocity"
                        target="_blank"
                        rel="noopener noreferrer"
                        class="group flex items-start gap-3 bg-surface p-4 transition-colors hover:bg-surface-alt"
                    >
                        <BookOpen class="mt-0.5 h-4 w-4 text-muted" />
                        <div class="flex-1">
                            <div class="caption-mono-upper text-dim">
                                01 · Documentation
                            </div>
                            <div class="mt-0.5 text-sm text-fg">
                                Read the full API reference and guides.
                            </div>
                        </div>
                        <ArrowRight
                            class="h-4 w-4 text-dim transition-transform group-hover:translate-x-0.5"
                        />
                    </a>
                    <a
                        href="https://github.com/velocitykode/velocity/tree/main/examples"
                        target="_blank"
                        rel="noopener noreferrer"
                        class="group flex items-start gap-3 bg-surface p-4 transition-colors hover:bg-surface-alt"
                    >
                        <Github class="mt-0.5 h-4 w-4 text-muted" />
                        <div class="flex-1">
                            <div class="caption-mono-upper text-dim">
                                02 · Examples
                            </div>
                            <div class="mt-0.5 text-sm text-fg">
                                Browse real-world starter apps & patterns.
                            </div>
                        </div>
                        <ArrowRight
                            class="h-4 w-4 text-dim transition-transform group-hover:translate-x-0.5"
                        />
                    </a>
                </div>
            </div>
        </main>

        <footer
            class="border-t border-border px-4 py-3 sm:px-5 lg:px-6"
        >
            <div
                class="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between"
            >
                <p class="caption-mono-upper text-dim">
                    © {{ year }} Velocity
                </p>
                <p class="caption-mono text-dim">
                    build:pre-1.0 · go/1.26+
                </p>
            </div>
        </footer>
    </div>
</template>
