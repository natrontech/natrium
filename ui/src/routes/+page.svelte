<script>
  export let data;
</script>

<h1>Welcome to SvelteKit</h1>
<p>
  Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

{#if data}
  <h2>Here's some data from the server</h2>
  <pre>{JSON.stringify(data, null, 2)}</pre>
{/if}
