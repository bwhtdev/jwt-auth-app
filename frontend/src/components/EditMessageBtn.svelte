<script lang='ts'>
  const { id, text } = $props();
</script>

<button x-on:click={`messageId='${id}';messageText='${text}';$store.modalOpen='edit-message'`} class="inline-flex items-center justify-center h-10 px-4 py-2 text-sm font-medium transition-colors bg-white border rounded-md hover:bg-neutral-100 active:bg-white focus:bg-white focus:outline-none focus:ring-2 focus:ring-neutral-200/60 focus:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none">Edit</button>
