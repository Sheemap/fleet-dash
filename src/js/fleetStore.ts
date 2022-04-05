import {defineStore} from "pinia";
import {TokenSet} from "./userStore";

export type FleetState = {
    _character_fleet_expires_at: number;
    character_fleet: CharacterFleet;
    _members_expires_at: number;
    members: FleetMember[];
    able_to_fetch_members: boolean;
}

export type CharacterFleet = {
    fleet_id: number;
    role: string;
    squad_id: number;
    wing_id: number;
}

export type FleetMember = {
    character_id: number;
    join_time: string;
    role: string;
    role_name: string;
    ship_type_id: number;
    solar_system_id: number;
    squad_id: number;
    station_id: number | null;
    takes_fleet_warp: boolean;
    wing_id: number;
}

export const useFleetStore = defineStore('fleet', {
    state: () => {
        return {
            _character_fleet_expires_at: 0,
            character_fleet: {
                fleet_id: 0,
                role: '',
                squad_id: 0,
                wing_id: 0,
            },
            _members_expires_at: 0,
            members: [],
            able_to_fetch_members: false,
        } as FleetState;
    },
    actions: {
        fetchFleet(charID: number, token: TokenSet) : Promise<CharacterFleet> {
            if (this._character_fleet_expires_at > Date.now() + 3000) {
                return Promise.resolve(this.character_fleet);
            }

            return new Promise((resolve, reject) => {
                    fetch(`https://esi.evetech.net/latest/characters/${charID}/fleet/`, {
                        headers: {
                            'Authorization': `Bearer ${token.access_token}`
                        }
                    })
                    .then(res => {
                        if (res.status >= 400 && res.status < 500) {
                            this.character_fleet = {
                                fleet_id: 0,
                                role: '',
                                squad_id: 0,
                                wing_id: 0,
                            };
                        }

                        let expires = res.headers.get('Expires');
                        if (expires) {
                            this._character_fleet_expires_at = new Date(expires).getTime();
                        }

                        return res.json();
                    })
                    .then(data => {
                        if (data.error){
                            reject(data.error);
                        } else {
                            this.character_fleet = data;
                            resolve(data);
                        }
                    })
                    .catch(err => {
                        reject(err);
                    });
            });
        },
        fetchMembers(token: TokenSet) : Promise<FleetMember[]> {
            if (this._members_expires_at > Date.now() + 3000) {
                return Promise.resolve(this.members);
            }

            if (isNaN(this.character_fleet.fleet_id) || this.character_fleet.fleet_id === 0) {
                return Promise.reject('No fleet ID');
            }

            return new Promise((resolve, reject) => {
                fetch(`https://esi.evetech.net/latest/fleets/${this.character_fleet.fleet_id}/members/`, {
                    headers: {
                        'Authorization': `Bearer ${token.access_token}`
                    }
                })
                .then(res => {
                    if (res.status >= 400 && res.status < 500) {
                        this.members = [];
                    }

                    let expires = res.headers.get('Expires');
                    if (expires) {
                        this._members_expires_at = new Date(expires).getTime();
                    }

                    return res.json();
                })
                .then(data => {
                    if (data.error){
                        reject(data.error);
                    } else {
                        this.members = data;
                        resolve(data);
                    }
                })
                .catch(err => {
                    reject(err);
                });
            });
        },
        fetchItemName(itemId: number, token: TokenSet) : Promise<string> {
            const cacheKey = 'item_'+ itemId.toString();
            let cachedName = localStorage.getItem(cacheKey);

            if  (cachedName) {
                return Promise.resolve(cachedName);
            }

            return new Promise((resolve, reject) => {
                fetch(`${import.meta.env.VITE_FLEETDASH_CORE_SCHEME}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/static/item?id=${itemId}`, {
                    headers: {
                        'Authorization': `Bearer ${token.access_token}`
                    }
                })
                .then(res => {
                    return res.json();
                })
                .then(data => {
                    if (data.Name && typeof data.Name === 'string') {
                        localStorage.setItem(cacheKey, data.Name);
                        resolve(data.Name);
                    } else {
                        reject("Item name is empty.");
                    }
                })
                .catch(err => {
                    reject(err);
                });
            });
        },
        fetchItemId(itemName: string, token: TokenSet) : Promise<number> {
            const cacheKey = 'item_name_'+ itemName;
            let cachedId = localStorage.getItem(cacheKey);

            if  (cachedId) {
                let id = parseInt(cachedId);

                if (isNaN(id)) {
                    localStorage.removeItem(cacheKey);
                }else{
                    return Promise.resolve(id);
                }
            }

            return new Promise((resolve, reject) => {
                fetch(`${import.meta.env.VITE_FLEETDASH_CORE_SCHEME}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/static/item?name=${itemName}`, {
                    headers: {
                        'Authorization': `Bearer ${token.access_token}`
                    }
                })
                .then(res => {
                    return res.json();
                })
                .then(data => {
                    if (data.ID && typeof data.ID === 'number') {
                        localStorage.setItem(cacheKey, data.ID);
                        resolve(data.ID);
                    } else {
                        reject("Item ID is empty.");
                    }
                })
                .catch(err => {
                    reject(err);
                });
            });
        },
        fetchSystemName(systemId: number, token: TokenSet) : Promise<string> {
            const cacheKey = 'system_'+ systemId.toString();
            let cachedName = localStorage.getItem(cacheKey);

            if  (cachedName) {
                return Promise.resolve(cachedName);
            }

            return new Promise((resolve, reject) => {
                fetch(`${import.meta.env.VITE_FLEETDASH_CORE_SCHEME}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/static/system?id=${systemId}`, {
                    headers: {
                        'Authorization': `Bearer ${token.access_token}`
                    }
                })
                    .then(res => {
                        return res.json();
                    })
                    .then(data => {
                        if (data.Name && typeof data.Name === 'string') {
                            localStorage.setItem(cacheKey, data.Name);
                            resolve(data.Name);
                        } else {
                            reject("System name is empty.");
                        }
                    })
                    .catch(err => {
                        reject(err);
                    });
            });
        },
    },
});
