<script lang="ts">
  interface Message {
    id: string;
    text: string;
    username: string;
    createdAt: string;
  }
  
  const getMessages = async (): Promise<Message[]> => {
    const res = await fetch('/api/v1/messages', {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getMessages()}
  <p>Loading messages...</p>
{:then messages}
  {#each messages as message}
    <a href={`/${message.id}`}>
      <p>{message.text}</p>
      <p>{message.username} - {message.createdAt}</p>
      <br/>
    </a>
  {/each}
{:catch error}
  <p>Cannot load message</p>
  <p>{error.message}</p>
{/await}
