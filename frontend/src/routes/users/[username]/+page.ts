// src/routes/users/[username]/+page.ts
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
    const username = params.username;
    // Fetch the user's profile data from your backend
    const response = await fetch(`http://localhost:8080/users/${username}`, {
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to load user data');
    }

    const userData = await response.json();

    const wishlistRequest = await fetch(`http://localhost:8080/users/wishlist/${username}`, {
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include'
    });

    if (!wishlistRequest.ok) {
        throw new Error('Failed to load wishlist data');
    }

    const wishlistData = await wishlistRequest.json();

    const sellingItemsRequest = await fetch(`http://localhost:8080/users/selling_items/${username}`, {
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include'
    });

    if (!sellingItemsRequest.ok) {
        throw new Error('Failed to load selling items data');
    }

    const sellingItemsData = await sellingItemsRequest.json();

    return {
        user: userData,
        wishlist: wishlistData,
        sellingItems: sellingItemsData
    };
};
