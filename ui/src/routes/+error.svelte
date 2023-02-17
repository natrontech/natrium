<script>
  import { page } from '$app/stores';
</script>

<div class="max-w-xl mx-auto py-8 px-4">
  <h1 class="text-2xl mb-8">Error</h1>
  <div class="card">
    <h2 class="text-xl mb-4">{$page.status}</h2>
    <p>{$page.error?.message}</p>
  </div>
</div>
