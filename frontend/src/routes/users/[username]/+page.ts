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

    return {
        user: userData
    };
};
