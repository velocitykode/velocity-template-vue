<script setup lang="ts">
import {
    AppLogo,
    Badge,
    IconButton,
    NavItem,
    StatCard,
    ThemeToggle,
    UserAvatar,
} from '@/components/ui';
import { usePageTitle } from '@/composables/use-page-title';
import { Link, usePage } from '@inertiajs/vue3';
import {
    Activity,
    ArrowRight,
    BookOpen,
    Check,
    Database,
    Gauge,
    LayoutGrid,
    LogOut,
    Menu,
    Route,
    Server,
    Settings,
    Shield,
    Terminal,
    X,
} from 'lucide-vue-next';
import { computed, ref } from 'vue';

usePageTitle('Dashboard');

const page = usePage<{ auth: { user: { name: string; email: string } } }>();
const user = computed(() => page.props.auth?.user);
const sidebarOpen = ref(false);

const navigation = [
    { name: 'Dashboard', href: '/dashboard', icon: LayoutGrid, active: true },
    { name: 'Routes', href: '#', icon: Route },
    { name: 'Middleware', href: '#', icon: Shield },
    { name: 'Database', href: '#', icon: Database },
];

const resources = [
    {
        name: 'Documentation',
        href: 'https://velocity.velocitykode.com/',
        icon: BookOpen,
        external: true,
    },
    { name: 'Settings', href: '#', icon: Settings },
];

const stats = [
    { label: 'Routes', value: '12', icon: Route, change: '+2' },
    { label: 'Middleware', value: '6', icon: Shield },
    { label: 'Cache hit', value: '94%', icon: Gauge },
    { label: 'Uptime', value: '99.9%', icon: Server },
];

const tasks = [
    { task: 'Configure routes', done: true },
    { task: 'Set up middleware', done: true },
    { task: 'Connect database', done: false },
    { task: 'Configure cache', done: false },
    { task: 'Deploy to production', done: false },
];

const activities = [
    { action: 'Migration completed', time: '2m', icon: Database },
    { action: 'Cache invalidated', time: '15m', icon: Server },
    { action: 'Route registered', time: '1h', icon: Route },
    { action: 'Middleware updated', time: '3h', icon: Shield },
    { action: 'Server restarted', time: '5h', icon: Terminal },
];

const completed = computed(() => tasks.filter((t) => t.done).length);
const progressPct = computed(() =>
    Math.round((completed.value / tasks.length) * 100),
);
</script>

<template>
    <div class="flex h-screen bg-bg text-fg">
        <div
            v-if="sidebarOpen"
            class="fixed inset-0 z-40 bg-black/50 lg:hidden"
            @click="sidebarOpen = false"
        />

        <aside
            :class="[
                'fixed inset-y-0 left-0 z-50 flex w-72 flex-col border-r border-border bg-surface transition-transform duration-200 ease-in-out lg:static lg:inset-auto lg:w-64 lg:translate-x-0 lg:transition-none xl:w-72',
                sidebarOpen ? 'translate-x-0' : '-translate-x-full',
            ]"
        >
            <div
                class="flex h-12 items-center justify-between border-b border-border px-4"
            >
                <AppLogo variant="full" size="md" href="/" />
                <div class="flex items-center gap-2">
                    <Badge variant="outline">v0.22</Badge>
                    <IconButton
                        :icon="X"
                        variant="ghost"
                        aria-label="Close menu"
                        class="lg:hidden"
                        @click="sidebarOpen = false"
                    />
                </div>
            </div>

            <nav class="flex-1 overflow-y-auto px-2 py-4">
                <div
                    class="caption-mono-upper mb-2 px-3 text-dim"
                >
                    &gt; Platform
                </div>
                <div class="space-y-0">
                    <NavItem
                        v-for="item in navigation"
                        :key="item.name"
                        :name="item.name"
                        :href="item.href"
                        :icon="item.icon"
                        :active="item.active"
                        @click="sidebarOpen = false"
                    />
                </div>

                <div
                    class="caption-mono-upper mt-6 mb-2 px-3 text-dim"
                >
                    &gt; Resources
                </div>
                <div class="space-y-0">
                    <NavItem
                        v-for="item in resources"
                        :key="item.name"
                        :name="item.name"
                        :href="item.href"
                        :icon="item.icon"
                        :external="item.external"
                    />
                </div>
            </nav>

            <div class="border-t border-border p-3">
                <div class="flex items-center gap-3">
                    <UserAvatar :name="user?.name" size="md" />
                    <div class="min-w-0 flex-1">
                        <div class="truncate text-sm font-medium text-fg">
                            {{ user?.name || 'User' }}
                        </div>
                        <div class="caption-mono truncate text-dim">
                            {{ user?.email }}
                        </div>
                    </div>
                    <Link href="/logout" method="post" as="button">
                        <IconButton
                            :icon="LogOut"
                            variant="danger"
                            aria-label="Log out"
                            as="button"
                            type="button"
                            size="sm"
                        />
                    </Link>
                </div>
            </div>
        </aside>

        <div class="flex flex-1 flex-col overflow-hidden">
            <header
                class="flex h-12 items-center justify-between border-b border-border bg-bg px-4 sm:px-5 lg:px-6"
            >
                <div class="flex items-center gap-3">
                    <IconButton
                        :icon="Menu"
                        variant="ghost"
                        size="sm"
                        aria-label="Open menu"
                        class="lg:hidden"
                        @click="sidebarOpen = true"
                    />
                    <h1 class="caption-mono-upper text-muted">
                        &gt; dashboard
                    </h1>
                </div>
                <div class="flex items-center gap-2">
                    <ThemeToggle />
                </div>
            </header>

            <main class="flex-1 overflow-y-auto">
                <div
                    class="mx-auto w-full max-w-[1800px] px-4 py-5 sm:px-5 sm:py-6 lg:px-6 xl:px-8 xl:py-7"
                >
                    <div class="mb-6 lg:mb-7">
                        <div class="mb-3 flex items-center gap-2">
                            <span class="relative flex h-2 w-2">
                                <span
                                    class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-60"
                                />
                                <span
                                    class="relative inline-flex h-2 w-2 rounded-full bg-emerald-500"
                                />
                            </span>
                            <span
                                class="caption-mono-upper text-emerald-700 dark:text-emerald-400"
                            >
                                Live
                            </span>
                        </div>
                        <h2
                            class="font-sans text-2xl font-semibold tracking-tight lg:text-3xl xl:text-[36px]"
                        >
                            Welcome back,
                            {{ user?.name?.split(' ')[0] || 'Developer' }}.
                        </h2>
                        <p class="mt-2 text-[14px] text-muted">
                            Your application overview - routes, middleware,
                            cache, uptime.
                        </p>
                    </div>

                    <div
                        class="mb-6 grid grid-cols-2 gap-px border border-border bg-border lg:mb-7 lg:grid-cols-4"
                    >
                        <StatCard
                            v-for="stat in stats"
                            :key="stat.label"
                            :label="stat.label"
                            :value="stat.value"
                            :icon="stat.icon"
                            :change="stat.change"
                        />
                    </div>

                    <div class="grid gap-5 lg:grid-cols-2 lg:gap-6">
                        <section class="border border-border bg-surface">
                            <header
                                class="flex items-center justify-between border-b border-border px-4 py-3 lg:px-5"
                            >
                                <div>
                                    <div class="caption-mono-upper text-dim">
                                        &gt; Getting started
                                    </div>
                                    <div class="mt-0.5 text-xs text-muted">
                                        {{ completed }} of
                                        {{ tasks.length }} complete
                                    </div>
                                </div>
                                <span
                                    class="font-sans text-2xl font-semibold tracking-tight"
                                >
                                    {{ progressPct }}%
                                </span>
                            </header>
                            <ul class="divide-y divide-border">
                                <li
                                    v-for="(item, i) in tasks"
                                    :key="i"
                                    class="flex items-center gap-3 px-4 py-2.5 lg:px-5"
                                >
                                    <div
                                        :class="[
                                            'flex h-5 w-5 items-center justify-center border',
                                            item.done
                                                ? 'border-fg bg-fg text-bg'
                                                : 'border-border-strong',
                                        ]"
                                    >
                                        <Check
                                            v-if="item.done"
                                            class="h-3 w-3"
                                            :stroke-width="3"
                                        />
                                    </div>
                                    <span
                                        :class="[
                                            'flex-1 text-sm',
                                            item.done
                                                ? 'text-dim line-through'
                                                : 'text-fg',
                                        ]"
                                    >
                                        {{ item.task }}
                                    </span>
                                    <ArrowRight class="h-4 w-4 text-dim" />
                                </li>
                            </ul>
                            <footer class="border-t border-border p-3">
                                <a
                                    href="https://github.com/velocitykode/velocity"
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    class="flex w-full items-center justify-center gap-2 bg-accent py-2.5 text-sm font-medium tracking-wide text-accent-fg transition-opacity hover:opacity-90"
                                >
                                    <BookOpen class="h-4 w-4" />
                                    View documentation
                                </a>
                            </footer>
                        </section>

                        <section class="border border-border bg-surface">
                            <header
                                class="flex items-center justify-between border-b border-border px-4 py-3 lg:px-5"
                            >
                                <div class="flex items-center gap-3">
                                    <Activity class="h-4 w-4 text-muted" />
                                    <div>
                                        <div
                                            class="caption-mono-upper text-dim"
                                        >
                                            &gt; Recent activity
                                        </div>
                                        <div
                                            class="mt-0.5 text-xs text-muted"
                                        >
                                            Live events
                                        </div>
                                    </div>
                                </div>
                                <span class="flex items-center gap-2">
                                    <span class="relative flex h-2 w-2">
                                        <span
                                            class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-60"
                                        />
                                        <span
                                            class="relative inline-flex h-2 w-2 rounded-full bg-emerald-500"
                                        />
                                    </span>
                                </span>
                            </header>
                            <ul class="divide-y divide-border">
                                <li
                                    v-for="(item, i) in activities"
                                    :key="i"
                                    class="flex items-center gap-3 px-4 py-2.5 lg:px-5"
                                >
                                    <component
                                        :is="item.icon"
                                        class="h-4 w-4 flex-shrink-0 text-muted"
                                    />
                                    <span
                                        class="flex-1 truncate text-sm text-fg"
                                    >
                                        {{ item.action }}
                                    </span>
                                    <span class="caption-mono text-dim">
                                        {{ item.time }}
                                    </span>
                                </li>
                            </ul>
                            <footer class="border-t border-border p-3">
                                <button
                                    class="caption-mono-upper w-full py-1 text-center text-muted transition-colors hover:text-fg"
                                >
                                    &gt; View all activity
                                </button>
                            </footer>
                        </section>
                    </div>
                </div>
            </main>
        </div>
    </div>
</template>
