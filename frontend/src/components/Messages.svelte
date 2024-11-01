<script lang="ts">
  import { onMount } from 'svelte';
  import { messageData, messageStatus, reloadMessageData } from './messageStore';

  onMount(() => {
     reloadMessageData();
  });
</script>

{#if $messageStatus == 'loading'}
  <p>Loading messages...</p>
{:else if $messageStatus == 'error' || $messageData.error}
  <p>Cannot load message</p>
  <p>{$messageData.error}</p>
{:else}
  {#each $messageData as message}
    <a href={`/${message.id}`}>
      <p>{message.text}</p>
      <p>{message.username} - {new Date(message.createdAt).toDateString()}</p>
      <br/>
    </a>
  {/each}
{/if}
