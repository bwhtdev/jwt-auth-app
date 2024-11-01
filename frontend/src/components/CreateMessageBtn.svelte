<script lang='ts'>
  import { loggedIn } from '@components/authStore';
</script>

{#if $loggedIn}
  <button x-on:click="messageText='';$store.modalOpen='create-message'" class="inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 bg-blue-600 rounded-md hover:bg-blue-700 focus:ring-2 focus:ring-offset-2 focus:ring-blue-700 focus:shadow-outline focus:outline-none">New Message</button>
{/if}
