import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/kit/vite';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: [
		vitePreprocess(),
		preprocess({
			postcss: true
		})
	],
	server: {
		port: 8080
	},
	envPrefix: 'NATRIUM_',
	kit: {
		adapter: adapter({
			out: 'build',
			envPrefix: 'NATRIUM_',
			polyfill: false
		}),
		csrf: {
			checkOrigin: false
		}
	},
	onwarn: (warning, handler) => {
		if (warning.code.startsWith('a11y-')) return;
		handler(warning);
	}
};

export default config;
