<script>
  import "../app.css";
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let isLoggedIn = false;
  let username = '';

  /**
	 * @param {string} token
	 */
  function decodeJWT(token) {
    const payload = token.split('.')[1];
    return JSON.parse(atob(payload));
  }

  /**
	 * @param {string} token
	 */
  function isTokenExpired(token) {
    const { exp } = decodeJWT(token);
    const currentTime = Math.floor(Date.now() / 1000);
    return exp < currentTime;
  }

  /**
	 * @param {string} token
	 */
  function getUsernameFromToken(token) {
    const decodedToken = decodeJWT(token);
    return decodedToken.Username; // Adjust this if the username is stored under a different claim
  }

  onMount(() => {
    const token = localStorage.getItem('token');
    if (token) {
      if (isTokenExpired(token)) {
        console.log('Token has expired');
        localStorage.removeItem('token');
      } else {
        isLoggedIn = true;
        username = getUsernameFromToken(token);
        console.log(username);
      }
    }
  });

  function logout() {
    localStorage.removeItem('token');
    isLoggedIn = false;
    goto('/login');
  }
</script>

<div class="flex flex-col min-h-screen">
  <header class="bg-white">
    <nav class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
      <div class="flex lg:flex-1">
        <a href="/" class="-m-2.5 p-2">
          <span class="text-lg font-semibold leading-6 text-gray-900">Local Marketplace</span>
        </a>
      </div>
      <div class="hidden lg:flex lg:gap-x-12">
        <a href="/marketplace" class="text-sm font-semibold leading-6 text-gray-900">Marketplace</a>
      </div>
      <div class="hidden lg:flex lg:flex-1 lg:justify-end">
        {#if isLoggedIn}
          <a href="/users/{username}" class="text-sm font-semibold leading-6 text-gray-900">Profile</a>
          <button on:click={logout} class="text-sm font-semibold leading-6 text-gray-900 ml-4">Log out</button>
        {:else}
          <a href="/login" class="text-sm font-semibold leading-6 text-gray-900">Log in <span aria-hidden="true">&rarr;</span></a>
        {/if}
      </div>
    </nav>
  </header>

  <main class="flex-grow">
    <slot></slot>
  </main>

  <footer class="bg-white text-black py-4">
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
