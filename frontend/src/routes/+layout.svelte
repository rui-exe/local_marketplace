<script>
  import "../app.css";
  import { onMount } from "svelte";
  import { authenticated, username} from "../stores/auth";

  async function logout() {
    try {
      const response = await fetch('http://localhost:8080/logout/', {
        method: 'POST',
        credentials: 'include', // Include the JWT token in the request
      });

      if (response.ok) {
        window.location.href = '/login';
      } else {
        console.error('Failed to log out');
      }
    } catch (err) {
    }
  }

  onMount(async () => {
    const response = await fetch('http://localhost:8080/users/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      },
      credentials: 'include', // Include the JWT token in the request
    });

    if (response.ok) {
      authenticated.set(true);
      const data = await response.json();
      username.set(data.username);

    } else {
      authenticated.set(false);
      username.set('');
    }
  });

</script>

<div class="flex flex-col min-h-screen">
  <header class="bg-gray-800 border-b border-gray-700">
    <nav class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
      <div class="flex lg:flex-1">
        <a href="/" class="-m-2.5 p-2">
          <span class="text-lg font-semibold leading-6 text-white">Local Marketplace</span>
        </a>
      </div>
      <div class="hidden lg:flex lg:gap-x-12">
        <a href="/marketplace" class="text-sm font-semibold leading-6 text-white">Marketplace</a>
      </div>
      <div class="hidden lg:flex lg:flex-1 lg:justify-end">
        {#if $authenticated === true}
          <a href="/users/{$username}" class="text-sm font-semibold leading-6 text-white">{$username}</a>
          <button on:click|preventDefault={logout} class="text-sm font-semibold leading-6 ml-4 text-white">Log out</button>
        {:else}
          <a href="/login" class="text-sm font-semibold leading-6 text-white">Log in <span aria-hidden="true">&rarr;</span></a>
        {/if}
      </div>
    </nav>
  </header>

  <main class="flex-grow">
    <slot></slot>
  </main>

  <footer class="py-4 bg-gray-800 text-white">
    <div class="mx-auto max-w-7xl px-6 lg:px-8 flex justify-between">
      <div>
        <a href="/about" class="text-sm font-semibold leading-6 hover:underline">About Us</a>
      </div>
      <div>
        <a href="/contacts" class="text-sm font-semibold leading-6 hover:underline">Contact Us</a>
      </div>
    </div>
  </footer>
</div>
