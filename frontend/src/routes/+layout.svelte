<script>
    import "../app.css";
    export let title = "Local Marketplace";
  </script>
  <header class="bg-blue-600 text-white shadow-md">
    <div class="container mx-auto px-6 py-4 flex justify-between items-center">
      <h1 class="text-2xl font-bold">
        {title}
      </h1>
      <nav>
        <ul class="flex space-x-4">
          <li>
            <a href="/" class="hover:text-blue-300">Home</a>
          </li>
          <li>
            <a href="/about" class="hover:text-blue-300">About</a>
          </li>
          <li>
            <a href="/profile" class="hover:text-blue-300">Profile</a>
          </li>
          <li>
            <a href="/login" class="hover:text-blue-300">Login</a>
          </li>
          <li>
            <a href="/signup" class="hover:text-blue-300">Sign Up</a>
          </li>
        </ul>
      </nav>
    </div>
  </header>
  
  <main>
    <slot></slot>
  </main>