<script lang='ts'>
  import SignUpBtn from './SignUpBtn.svelte';
  import LogInBtn from './LogInBtn.svelte';

  import { onMount } from 'svelte';
  import { get } from 'svelte/store';

  import { username, loggedIn } from '@components/authStore';
  
  import { isLoggedIn, logOut, getUsername } from '../alpine/authUtils';
  
  onMount(() => {
    loggedIn.set(isLoggedIn());
    if (get(loggedIn)) username.set(getUsername());
  });
</script>

{#if $loggedIn}
  <p class='mr-2'>{$username}</p>
  <button on:click={logOut} class="inline-flex items-center justify-center px-4 py-2 text-base font-medium leading-6 text-gray-600 whitespace-no-wrap bg-white border border-gray-200 rounded-md shadow-sm hover:bg-gray-50 focus:outline-none focus:shadow-none">Log Out</button>
{:else}
  <LogInBtn />
  <span class="inline-flex rounded-md shadow-sm">
    <SignUpBtn />
  </span>
{/if}
