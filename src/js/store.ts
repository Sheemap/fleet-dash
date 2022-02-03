import { defineStore } from 'pinia';
import { inject } from 'vue';
import {Axios, AxiosStatic} from "axios";

type TokenSet = {
    access_token: string;
    refresh_token: string;
    expires_at: number;
};

type ApiTokenSet = {
    access_token: string;
    refresh_token: string;
    expires_in: number;
};

export const useUserStore = defineStore('user', {
    state: () => {
        return {
            auth_state: localStorage.getItem('auth_state') || '',
            _portrait_url: localStorage.getItem('portrait_url') || '',
            _token_set: JSON.parse(localStorage.getItem('token_set') || "{}") || {},
        }
    },
    getters: {
        character_name: state => {
            if (Object.keys(state._token_set).length === 0) return '';

            // Parse the JWT, no validation
            let payload = JSON.parse(atob(state._token_set.access_token.split('.')[1]));

            return payload.name;
        },
        character_id: state => {
            if (Object.keys(state._token_set).length === 0) return '';

            // Parse the JWT, no validation
            let payload = JSON.parse(atob(state._token_set.access_token.split('.')[1]));

            return payload.sub.split(':')[2];
        }
    },
    actions: {
        signOut(){
            localStorage.removeItem('auth_state')
            localStorage.removeItem('portrait_url')
            localStorage.removeItem('token_set')
            this.resetState()
        },
        setState(){
            this.auth_state = Math.random().toString(16).substr(2, 32);
            localStorage.setItem('auth_state', this.auth_state);
            return this.auth_state;
        },
        resetState(){
            this.auth_state = '';
            localStorage.removeItem('auth_state');
        },
        setToken(tokenSet : ApiTokenSet){
            this._token_set = {
                access_token: tokenSet.access_token,
                refresh_token: tokenSet.refresh_token,
                expires_at: Math.floor(Date.now() / 1000) + tokenSet.expires_in,
            }

            localStorage.setItem('token_set', JSON.stringify(this._token_set));

            return this._token_set;
        },
        getActiveToken() : Promise<TokenSet|Object> {
            return new Promise((resolve, reject) => {
                // If no tokenset, resolve with empty
                if (Object.keys(this._token_set).length === 0) {
                    resolve({});
                }
                // If tokenset is not expired or close, return it
                else if (this._token_set.expires_at > (Date.now() / 1000) - 60) {
                    resolve(this._token_set);
                }
                // If tokenset is expired, or close, refresh it
                else {
                    this.refreshToken()
                        .then(resolve)
                        .catch(err => resolve({}));
                }
            });
        },
        refreshToken() : Promise<TokenSet|Object> {
            return new Promise((resolve, reject) => {
                if (Object.keys(this._token_set).length === 0) {
                    resolve({});
                }

                const axios = inject('axios') as AxiosStatic;

                const params = new URLSearchParams();
                params.append('client_id', import.meta.env.VITE_EVE_CLIENT_ID);
                params.append('client_secret', import.meta.env.VITE_EVE_CLIENT_SECRET);
                params.append('grant_type', 'refresh_token');
                params.append('redirect_uri', import.meta.env.VITE_EVE_REDIRECT_URI);
                params.append('refresh_token', this._token_set.refresh_token);
                axios.post('https://login.eveonline.com/v2/oauth/token', params)
                    .then(res => {
                        this.setToken(res.data)
                        resolve(this._token_set);
                    })
                    .catch(err => {
                        reject(err);
                    });

            })
        },
        getPortraitUrl(){
            if(this._portrait_url !== ''){
                return new Promise((resolve, reject) => resolve(this._portrait_url));
            }

            return new Promise((resolve, reject) => {
                this.getActiveToken().then(maybeToken => {
                    if(Object.keys(maybeToken).length === 0){
                        reject('');
                    }

                    let token = maybeToken as TokenSet;

                    let charID = this.character_id;
                    // Get the portrait url
                    fetch(`https://esi.evetech.net/latest/characters/${charID}/portrait/`, {
                        headers: {
                            'Authorization': `Bearer ${token.access_token}`
                        }
                    })
                        .then(res => res.json())
                        .then(data => {
                            this._portrait_url = data.px64x64;
                            localStorage.setItem('portrait_url', data.px64x64);
                            resolve(data.px64x64);
                        })
                        .catch(err => {
                            console.error(err);
                            resolve('');
                        });
                    });
            });
        },
    }
})
