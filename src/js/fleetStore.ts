import {defineStore} from "pinia";
import {TokenSet} from "./userStore";

export type FleetState = {
    _character_fleet_expires_at: number;
    character_fleet: CharacterFleet;
    _members_expires_at: number;
    members: FleetMember[];
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
    station_id: number;
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
        } as FleetState;
    },
    actions: {
        fetchFleet(charID: number, token: TokenSet) : Promise<CharacterFleet> {
            if (this._character_fleet_expires_at > Date.now()) {
                return Promise.resolve(this.character_fleet);
            }

            return new Promise((resolve, reject) => {
                    fetch(`https://esi.evetech.net/latest/characters/${charID}/fleet/`, {
                        headers: {
                            'Authorization': `Bearer ${token.access_token}`
                        }
                    })
                    .then(res => {
                        let expires = res.headers.get('Expires');
                        if (expires) {
                            this._character_fleet_expires_at = new Date(expires).getTime();
                        } else {
                            this._character_fleet_expires_at = 0;
                        }

                        return res.json();
                    })
                    .then(data => {
                        this.character_fleet = data;
                        resolve(data);
                    })
                    .catch(err => {
                        reject(err);
                    });
            });
        },
        fetchMembers(token: TokenSet) : Promise<FleetMember[]> {
            if (this._members_expires_at > Date.now()) {
                return Promise.resolve(this.members);
            }

            if (this.character_fleet.fleet_id === 0) {
                return Promise.reject('No fleet ID');
            }

            return new Promise((resolve, reject) => {
                fetch(`https://esi.evetech.net/latest/fleets/${this.character_fleet.fleet_id}/members/`, {
                    headers: {
                        'Authorization': `Bearer ${token.access_token}`
                    }
                })
                .then(res => {
                    let expires = res.headers.get('Expires');
                    if (expires) {
                        this._members_expires_at = new Date(expires).getTime();
                    } else {
                        this._members_expires_at = 0;
                    }

                    return res.json();
                })
                .then(data => {
                    this.members = data;
                    resolve(data);
                })
                .catch(err => {
                    reject(err);
                });
            });
        },
    },
});
