<script>
	// @ts-nocheck
	import Icon from 'svelte-icons-pack/Icon.svelte';
	import { fly } from 'svelte/transition';
	import HiOutlineHome from 'svelte-icons-pack/hi/HiOutlineHome';
	import HiSolidHome from 'svelte-icons-pack/hi/HiSolidHome';
	import HiOutlineInbox from 'svelte-icons-pack/hi/HiOutlineInbox';
	import HiSolidInbox from 'svelte-icons-pack/hi/HiSolidInbox';
	import { page } from '$app/stores';
  import {clickOutside} from '$lib/utils/clickOutside';

	let routes = [
		{ name: 'Home', href: '/', icon: HiOutlineHome, currentIcon: HiSolidHome },
		{ name: 'Repositories', href: '/repositories', icon: HiOutlineInbox, currentIcon: HiSolidInbox }
	];
	let settings = [
		{ name: 'Profile', href: '/profile' },
		{ name: 'Settings', href: '/settings' }
	];
	let open = false;

	function toggle() {
		open = !open;
	}
</script>

<div class=" fixed inset-y-0 flex w-64 flex-col border-r border-gray-200 bg-gray-100 pt-5 pb-4">
	<div class="flex flex-shrink-0 items-center px-6">
		<img
			class="h-8 w-auto"
			src="https://tailwindui.com/img/logos/mark.svg?color=blue&shade=500"
			alt="Your Company"
		/>
	</div>
	<!-- Sidebar component, swap this element with another sidebar if you like -->
	<div class="mt-5 flex h-0 flex-1 flex-col overflow-y-auto pt-1">
		<!-- User account dropdown -->
		<div class="relative inline-block px-3 text-left">
			<div>
				<button
					on:click|stopPropagation={() => toggle()}
          use:clickOutside
          on:click_outside={() => open = false}
					type="button"
					class="group w-full rounded-md bg-gray-100 px-3.5 py-2 text-left text-sm font-medium text-gray-700 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-gray-100"
					id="options-menu-button"
					aria-expanded="false"
					aria-haspopup="true"
				>
					<span class="flex w-full items-center justify-between">
						<span class="flex min-w-0 items-center justify-between space-x-3">
							<img
								class="h-10 w-10 flex-shrink-0 rounded-full bg-gray-300"
								src="https://images.unsplash.com/photo-1502685104226-ee32379fefbe?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=3&w=256&h=256&q=80"
								alt=""
							/>
							<span class="flex min-w-0 flex-1 flex-col">
								<span class="truncate text-sm font-medium text-gray-900">Jessy Schwarz</span>
								<span class="truncate text-sm text-gray-500">@jessyschwarz</span>
							</span>
						</span>
						<svg
							class="h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M10 3a.75.75 0 01.55.24l3.25 3.5a.75.75 0 11-1.1 1.02L10 4.852 7.3 7.76a.75.75 0 01-1.1-1.02l3.25-3.5A.75.75 0 0110 3zm-3.76 9.2a.75.75 0 011.06.04l2.7 2.908 2.7-2.908a.75.75 0 111.1 1.02l-3.25 3.5a.75.75 0 01-1.1 0l-3.25-3.5a.75.75 0 01.04-1.06z"
								clip-rule="evenodd"
							/>
						</svg>
					</span>
				</button>
			</div>

			{#if open}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<div
					on:click|stopPropagation={() => toggle()}
					transition:fly={{ y: 20, duration: 200 }}
					class="absolute right-0 left-0 z-10 mx-3 mt-1 origin-top divide-y divide-gray-200 rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none"
					role="menu"
					aria-orientation="vertical"
					aria-labelledby="options-menu-button"
					tabindex="-1"
				>
					<div class="py-1" role="none">
						{#each settings as setting}
							<a
								href={setting.href}
								class="text-gray-700 block px-4 py-2 text-sm"
								role="menuitem"
								tabindex="-1"
								id="options-menu-item-0">{setting.name}</a
							>
						{/each}
					</div>
					<div class="py-1" role="none">
						<a
							href="/logout"
							class="text-gray-700 block px-4 py-2 text-sm"
							role="menuitem"
							tabindex="-1"
							id="options-menu-item-5">Logout</a
						>
					</div>
				</div>
			{/if}
		</div>
		<!-- Navigation -->
		<nav class="mt-6 px-3">
			<div class="space-y-1">
				{#each routes as route}
					<!-- Current: "bg-gray-200 text-gray-900", Default: "text-gray-700 hover:bg-gray-50 hover:text-gray-900" -->
					<a
						href={route.href}
						class="text-gray-900 group flex items-center rounded-md px-2 py-2 text-sm font-medium {$page
							.route.id === route.href
							? 'bg-gray-200'
							: ''}"
						aria-current="page"
					>
						{#if $page.route.id === route.href}
							<Icon className="w-5 h-5" src={route.currentIcon} />{@html '&nbsp;'}{route.name}
						{:else}
							<Icon className="w-5 h-5" src={route.icon} />{@html '&nbsp;'}{route.name}
						{/if}
					</a>
				{/each}
			</div>
		</nav>
	</div>
</div>
