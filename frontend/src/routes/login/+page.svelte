<script>
    let email = '';
    let password = '';
    let error = '';
    let success = '';
  
    async function handleSubmit(event) {
      event.preventDefault();
  
      // Clear previous messages
      error = '';
      success = '';
  
      // Send login request
      try {
        const response = await fetch('http://localhost:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email, password }),
        });
  
        if (!response.ok) {
          error = await response.json().then((data) => data.error);
          throw new Error(error);
        }
  
        const data = await response.json();
        success = 'Login successful!';
        error = '';
        console.log('Login successful:', data);
        // Store the token in local storage
        localStorage.setItem('token', data.token);
        // Redirect to the dashboard
        setTimeout(() => {
          window.location.href = '/';
        }, 2000);
      } catch (err) {
            success = '';
            console.error('Login failed:', err);
      }
    }
</script>

<section class="bg-gray-50 dark:bg-gray-900">
  <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
      <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
              <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
                  Sign in to your account
              </h1>
              <form class="space-y-4 md:space-y-6" on:submit|preventDefault={handleSubmit}>
                  <div>
                      <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your email</label>
                      <input type="email" name="email" id="email" bind:value={email} class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@company.com" required="">
                  </div>
                  <div>
                      <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                      <input type="password" name="password" id="password" bind:value={password} placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required="">
                  </div>
                  <button type="submit" class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>
                  {#if error}
                      <p class="text-sm text-red-500 dark:text-red-400">{error}</p>
                  {/if}
                  {#if success}
                      <p class="text-sm text-green-500 dark:text-green-400">{success}</p>
                  {/if}
                  <p class="text-sm font-light text-gray-500 dark:text-gray-400">
                      Don’t have an account yet? <a href="/signup" class="text-primary-600 hover:underline dark:text-white">Sign up</a>
                  </p>
              </form>
          </div>
      </div>
  </div>
</section>
  

  