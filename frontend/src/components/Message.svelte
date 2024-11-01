<script lang='ts'> 
  let { id } = $props();

  import EditMessageBtn from './EditMessageBtn.svelte';
  import DeleteMessageBtn from './DeleteMessageBtn.svelte';

  import { loggedIn, username } from './authStore';

  interface Message {
    id: string;
    text: string;
    username: string;
    createdAt: string;
  }
  
  const getMessage = async (): Promise<Message> => {
    const res = await fetch(`/api/v1/message/id/${id}`, {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getMessage()}
  <p>Loading message...</p>
{:then message}
  {#if !message.error}
    <div x-init={`messageId='${message.id}';messageText='${message.text}'`}>
      <p x-text='messageText'></p>
      <p>{message.username} - {new Date(message.createdAt).toDateString()}</p>
    </div>

    {#if $loggedIn && message.username == $username}
      <EditMessageBtn />
      <DeleteMessageBtn />
    {/if}
  {:else}
    <p>Message does not exist.</p>
  {/if}
{:catch error}
  <p>Cannot load message.</p>
  <p>{error.message}</p>
{/await}  
