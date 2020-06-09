import { ApiService } from './networkManager.js' 
import {setToken} from '../helpers/storage'

export const openWebSocket = async (gameId, token) => {
    ApiService.wsConnect(gameId, token)
}

export const login = async (username, password) => {
    const path = '/login'
    const creds = Buffer.from(username + ':' + password).toString('base64');
    const header = {Authorization: `Basic ${creds}`}
    const res = await ApiService.get(path, null, null, header)
    setToken(res['token'])
}
