<link rel="stylesheet" href="https://demos.creative-tim.com/notus-js/assets/styles/tailwind.css">
<link rel="stylesheet" href="https://demos.creative-tim.com/notus-js/assets/vendor/@fortawesome/fontawesome-free/css/all.min.css">

<script lang="ts">
	import { username } from "../../../stores/auth";

  export let data: {
    user: {
      username: string;
      email: string;
      phone: string;
      role: string;
      picture: string;
      created_at: string;
      wishlist: string[];
    };
    wishlist: {
      ID: string;
      name: string;
      description: string;
      price: number;
      category: string;
      status: string;
      seller_id: string;
      picture: string;
      created_at: string;
      updated_at: string;
    }[];
    sellingItems: {
      ID: string;
      name: string;
      description: string;
      price: number;
      category: string;
      status: string;
      seller_id: string;
      picture: string;
      created_at: string;
      updated_at: string;
    }[];
  };
</script>

<div class="profile-page">
  <section class="relative block h-500-px">
    <div class="absolute top-0 w-full h-full bg-center bg-cover" style="
            background-image: url('/marketplace_wallpaper.png');
          ">
      <span id="blackOverlay" class="w-full h-full absolute opacity-50 bg-black"></span>
    </div>
    <div class="top-auto bottom-0 left-0 right-0 w-full absolute pointer-events-none overflow-hidden h-70-px" style="transform: translateZ(0px)">
      <svg class="absolute bottom-0 overflow-hidden" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="none" version="1.1" viewBox="0 0 2560 100" x="0" y="0">
        <polygon class="text-blueGray-200 fill-current" points="2560 0 2560 100 0 100"></polygon>
      </svg>
    </div>
  </section>
  <section class="relative py-16 bg-blueGray-200">
    <div class="container mx-auto px-4">
      <div class="relative flex flex-col min-w-0 break-words bg-white w-full mb-6 shadow-xl rounded-lg -mt-64">
        <div class="px-6">
          <div class="flex flex-wrap justify-center">
            <div class="w-full lg:w-3/12 px-4 lg:order-2 flex justify-center">
              <div class="relative">
                {#if data.user.picture}
                  <img alt="..." src={data.user.picture} class="shadow-xl rounded-full h-auto align-middle border-none absolute -m-16 -ml-20 lg:-ml-16 max-w-150-px">
                {:else}
                  <img alt="..." src="https://static.vecteezy.com/system/resources/previews/020/765/399/non_2x/default-profile-account-unknown-icon-black-silhouette-free-vector.jpg" class="shadow-xl rounded-full h-auto align-middle border-none absolute -m-16 -ml-20 lg:-ml-16 max-w-150-px">
                {/if}
              </div>
            </div>
            <div class="w-full lg:w-4/12 px-4 lg:order-3 lg:text-right lg:self-center">
            </div>
            <div class="w-full lg:w-4/12 px-4 lg:order-1">
              <div class="flex justify-center py-4 lg:pt-4 pt-8">
              </div>
            </div>
          </div>
          <div class="text-center mt-12">
            <h3 class="text-4xl font-semibold leading-normal text-blueGray-700 mb-2">
            {data.user.username}
            </h3>
            <div class="text-sm leading-normal mt-0 mb-2 text-blueGray-400 font-bold uppercase">
              Member since {new Date(data.user.created_at).toLocaleDateString()}
              <br>
              Phone: {data.user.phone}
            </div>
            <div class="mb-2 text-blueGray-600 mt-10">
              Email: {data.user.email}
            </div>
            <div class="mb-2 text-blueGray-600">
              {#if data.user.role === 'BUYER'}
                <span class="text-red-500">Buyer</span>
              {:else}
                <span class="text-green">Seller</span>
              {/if}
            </div>
          </div>
          {#if data.user.role === 'BUYER'}
          <div class="mt-10 py-10 border-t border-blueGray-200 text-center">
            <div class="flex flex-wrap justify-center">
              <div class="w-full lg:w-9/12 px-4">
                <h4 class="text-2xl font-semibold leading-normal mb-4 text-blueGray-700">
                  {data.user.username}'s Wishlist
                </h4>
                <ul>
                  {#if data.wishlist.length === 0}
                    <p class="text-blueGray-400">No products in wishlist</p>
                  {/if}
                  {#each data.wishlist as product}
                    <a href="/products/{product.ID}">
                      <li class="mb-6 bg-gray-100 rounded-lg shadow-lg p-4">
                        <div class="flex">
                          <div class="w-1/4">
                            {#if product.picture}
                              <img
                                src={product.picture}
                                alt={product.name}
                                class="w-full h-48 rounded-full object-cover"
                              />
                            {:else}
                              <img
                                src="https://sudbury.legendboats.com/resource/defaultProductImage"
                                alt={product.name}
                                class="w-full h-48 rounded-full object-cover"
                              />
                            {/if}
                          </div>
                          <div class="w-3/4 pl-4">
                            <h5 class="text-xl font-semibold">{product.name}</h5>
                            <p class="text-blueGray-600">{product.description}</p>
                            <p class="text-blueGray-400">Category: {product.category}</p>
                            <p class="text-blueGray-400">Price: ${product.price.toFixed(2)}</p>
                            <p class="text-blueGray-400">Status: {product.status}</p>
                          </div>
                        </div>
                      </li>
                    </a>
                  {/each}
                </ul>
              </div>
            </div>
          </div>
          {:else}
          <div class="mt-10 py-10 border-t border-blueGray-200 text-center">
            <div class="flex flex-wrap justify-center">
              <div class="w-full lg:w-9/12 px-4">
                <h4 class="text-2xl font-semibold leading-normal mb-4 text-blueGray-700">
                  {data.user.username}'s Products for Sale
                </h4>
                <ul>
                  {#if data.sellingItems.length === 0}
                    <p class="text-blueGray-400">No products for sale</p>
                  {/if}
                  {#each data.sellingItems as product}
                    <a href="/products/{product.ID}">
                      <li class="mb-6 bg-gray-100 rounded-lg shadow-lg p-4">
                        <div class="flex">
                          <div class="w-1/4">
                            {#if product.picture}
                              <img
                                src={product.picture}
                                alt={product.name}
                                class="w-full h-48 rounded-full object-cover"
                              />
                            {:else}
                              <img
                                src="https://sudbury.legendboats.com/resource/defaultProductImage"
                                alt={product.name}
                                class="w-full h-48 rounded-full object-cover"
                              />
                            {/if}
                          </div>
                          <div class="w-3/4 pl-4">
                            <h5 class="text-xl font-semibold">{product.name}</h5>
                            <p class="text-blueGray-600">{product.description}</p>
                            <p class="text-blueGray-400">Category: {product.category}</p>
                            <p class="text-blueGray-400">Price: ${product.price.toFixed(2)}</p>
                            <p class="text-blueGray-400">Status: {product.status}</p>
                          </div>
                        </div>
                      </li>
                    </a>
                  {/each}
                </ul>
              </div>
            </div>
          </div>
          {/if}
        </div>
      </div>
    </div>
  </section>
</div>