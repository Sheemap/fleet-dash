import { defineStore } from 'pinia';
import {
    TOTAL_OVER_TIME_WIDGET_ID,
    INCOMING_JAM_WIDGET_ID,
    OUTGOING_JAM_WIDGET_ID,
    OVERVIEW_WIDGET_ID,
    SLIDER_WIDGET_ID
} from './constants'
import {createUUID} from "./shared";

export type TokenSet = {
    access_token: string;
    refresh_token: string;
    expires_at: number;
};

type ApiTokenSet = {
    access_token: string;
    refresh_token: string;
    expires_in: number;
};

function appendUniqueId(widget: string): string {
    return widget + "_" + createUUID();
}

function generateDefaultLayout(){
    let layout : any[] = [];

    let uuid = createUUID();
    let settings = {"title": "Incoming Damage", "periodSeconds": 15, "eventType": "IncomingDamageEvent"};
    localStorage.setItem("widget_settings_" + uuid, JSON.stringify(settings));
    layout.push({"x":0,"y":0,"w":3,"h":1,"i":`${TOTAL_OVER_TIME_WIDGET_ID}_${uuid}`});

    let uuid2 = createUUID();
    let settings2 = {"title": "Outgoing Damage", "periodSeconds": 15, "eventType": "OutgoingDamageEvent"};
    localStorage.setItem("widget_settings_" + uuid2, JSON.stringify(settings2));
    layout.push({"x":3,"y":0,"w":3,"h":1,"i":`${TOTAL_OVER_TIME_WIDGET_ID}_${uuid2}`});

    layout.push({"x":6,"y":0,"w":3,"h":3,"i":appendUniqueId(INCOMING_JAM_WIDGET_ID)});
    layout.push({"x":9,"y":0,"w":3,"h":3,"i":appendUniqueId(OUTGOING_JAM_WIDGET_ID)});
    layout.push({"x":0,"y":1,"w":3,"h":2,"i":appendUniqueId(OVERVIEW_WIDGET_ID)});
    layout.push({"x":3,"y":1,"w":3,"h":1,"i":appendUniqueId(SLIDER_WIDGET_ID)});

    return layout;
}

export const useUserStore = defineStore('user', {
    state: () => {
        let dashLayout = JSON.parse(localStorage.getItem('dash_layout') || JSON.stringify([])) || [];
        if(dashLayout.length === 0){
            dashLayout = generateDefaultLayout(dashLayout);
        }
        
        return {
            auth_state: localStorage.getItem('auth_state') || '',
            widget_drawer_open: false,
            dash_layout: dashLayout,
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
        },
    },
    actions: {
        appendToLayout(newItem){
            this.dash_layout.push(newItem);
            localStorage.setItem('dash_layout', JSON.stringify(this.dash_layout));
        },
        updateLayout(newLayout) {
            this.dash_layout = newLayout;
            localStorage.setItem('dash_layout', JSON.stringify(newLayout));
        },
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
        getActiveToken() : Promise<TokenSet> {
            return new Promise((resolve, reject) => {
                // If no tokenset, reject
                if (Object.keys(this._token_set).length === 0) {
                    reject();
                }
                // If tokenset is not expired or close, return it
                else if (this._token_set.expires_at > (Date.now() / 1000) + 60) {
                    resolve(this._token_set);
                }
                // If tokenset is expired, or close, refresh it
                else {
                    this.refreshToken()
                        .then(resolve)
                        .catch(reject);
                }
            });
        },
        refreshToken() : Promise<TokenSet> {
            return new Promise((resolve, reject) => {
                if (Object.keys(this._token_set).length === 0) {
                    reject();
                }

                const params = new URLSearchParams();
                params.append('client_id', import.meta.env.VITE_EVE_CLIENT_ID);
                params.append('client_secret', import.meta.env.VITE_EVE_CLIENT_SECRET);
                params.append('grant_type', 'refresh_token');
                params.append('redirect_uri', import.meta.env.VITE_EVE_REDIRECT_URI);
                params.append('refresh_token', this._token_set.refresh_token);

                const options = {
                    method: 'POST',
                    body: params,
                }
                fetch('https://login.eveonline.com/v2/oauth/token', options)
                    .then(response => response.json())
                    .then(res => {
                        this.setToken(res)
                        resolve(this._token_set);
                    })
                    .catch(err => {
                        reject(err);
                    });
            })
        },
        getPortraitUrl() : Promise<string> {
            if(this._portrait_url !== ''){
                return new Promise((resolve, _) => resolve(this._portrait_url));
            }

            return new Promise((resolve, _) => {
                this.getActiveToken().then(token => {
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
                })
                    .catch(err => {
                        console.error(err);
                        resolve('');
                    });
            });
        },
    }
})
