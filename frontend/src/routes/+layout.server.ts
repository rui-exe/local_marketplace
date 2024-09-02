export function load({ cookies }) {
  const loggedIn = cookies.get('loggedIn');
  const username = cookies.get('username');
    return { loggedIn , username };
}