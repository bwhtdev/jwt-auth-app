<script lang='ts'>
  let { id } = $props();

  interface Message {
    id: string;
    text: string;
    username: string;
    createdAt: string;
  }
  
  const getMessage = async (): Promise<Message> => {
    const res = await fetch(`/api/v1/message/${id}`, {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getMessage()}
  <p>Loading message...</p>
{:then message}
  <div>
    <p>{message.text}</p>
    <p>{message.username} - {message.createdAt}</p>
  </div>
{:catch error}
  <p>Cannot load message</p>
  <p>{error.message}</p>
{/await}  
