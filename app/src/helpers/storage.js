// Local storage helper. It will be used to store auth tokens

export const setToken = (token) => {
    localStorage.setItem('auth_token', token);
}

export const getToken = () => {
    const token = localStorage.getItem('auth_token') || ''
    if (token === '' || token === 'undefined') return null
    const token_payload = token.split('.')[1];
    const token_claims = JSON.parse(window.atob(token_payload));
    // TODO: check expiry date or we could let the server return 4xx
    return token 
}

export const removeToken = () => {
    localStorage.removeItem('auth_token');
}
