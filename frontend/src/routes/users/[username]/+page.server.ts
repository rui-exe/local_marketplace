import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
  const { username } = params;
  return {
    username
  };
};