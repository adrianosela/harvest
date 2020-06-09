import axios from 'axios';
import { getToken } from '../helpers/storage.js'

const injectTokenHeader = (headers) => {
    const token = getToken()
    return {
        ...headers,
        Authorization: headers.Authorization ? headers.Authorization : `Bearer ${token}`,
        Accept: 'application/json'
    }
}

export const ApiService = {

    baseUrl: process.env.API_ENDPOINT || 'http://localhost:8080',
    wsBaseUrl: 'ws://localhost:8080',

    wsConnect: async function(gameId, token) {
        token = token || getToken() // FIXME: Remove mock token fallback
        const wsURL = `${this.wsBaseUrl}/game/${gameId}/watch?token=${token}`
        const ws = new WebSocket(wsURL);
        ws.onopen = () => {
	        console.log('connected');
	    };

	    ws.onmessage = (msg) => {
	        console.log(msg.data); // FIXME: what we sending -- state or moves?
	    };

	    ws.onclose = () => {
            console.log('disconnected');
	    };
    },

    get: async function(path, page, payload, headers) {
        const fullHeaders = injectTokenHeader({
            ...headers,
            'Content-Type': 'application/json'
        })
        if (page) { path += '?page=' + page } 
        const { data } = await axios.get(this.baseUrl + path, { headers: fullHeaders, params: payload })
        if (data['success'] === false) {
            // How to handle failed api calls?
            // check error status code? see if its unauthorized, bad, request, etc   
        }
        return data
    },

    post: async function(path, payload, headers) {
        const fullHeaders = injectTokenHeader({
            ...headers,
            'Content-Type': 'application/json'
        })
        const { data } = await axios.post(this.baseUrl + path, payload, { headers: fullHeaders })
        return data

    },

    delete: async function(path, payload, headers) {
        const fullHeaders = injectTokenHeader({
            ...headers,
            'Content-Type': 'application/json'
        })
        const { data } = await axios.delete(this.baseUrl + path,
            { 
                headers: fullHeaders,
                data: payload
            })
        if (data['success'] === false) {
            // How to handle failed api calls?
            // check error status code? see if its unauthorized, bad, request, etc   
        }
        return data
    },

    patch: async function(path, payload, headers) {
        const fullHeaders = injectTokenHeader({
            ...headers,
            'Content-Type': 'application/json'
        })
        const { data } = await axios.patch(this.baseUrl + path, payload, { headers: fullHeaders })
        return data
    }
}
