import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params, fetch }) => {
    const productID = params.productID;

    const response = await fetch(`http://localhost:8080/products/${productID}`, {
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to load product data');
    }

    const productData = await response.json();

    return {
        product: productData
    };
}
