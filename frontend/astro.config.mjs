// @ts-check
import { defineConfig } from 'astro/config';

import svelte from '@astrojs/svelte';
import sitemap from '@astrojs/sitemap';
import tailwind from '@astrojs/tailwind';
import alpinejs from '@astrojs/alpinejs';

import node from '@astrojs/node';

// https://astro.build/config
export default defineConfig({
  integrations: [svelte(), sitemap(), tailwind(), alpinejs({ entrypoint: '/src/entrypoint' })],
  output: 'server',
  adapter: node({
    mode: 'standalone'
  })
});
