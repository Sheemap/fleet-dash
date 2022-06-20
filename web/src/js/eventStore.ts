import {defineStore} from "pinia";
import {TokenSet} from "./userStore";

export type EventState = {
    active: boolean;
    stream_failure_count: number;
    last_stream_failure_time: number;
    _stream_ticket: string;
    session_fleet_id: number;
}

export const useEventStore = defineStore('event', {
    state: () => {
        return {
            active: false,
            stream_failure_count: 0,
            last_stream_failure_time: 0,
            _stream_ticket: localStorage.getItem('stream_ticket') || '',
            session_fleet_id: parseInt(localStorage.getItem('session_fleet_id') || '0'),
        } as EventState;
    },
    actions: {
        getStreamTicket(tokenSet : TokenSet) : Promise<string> {
            return new Promise((resolve, reject) => {
                if (this._stream_ticket !== '') {
                    resolve(this._stream_ticket);
                } else {
                    const options = {
                        method: 'POST',
                        headers: {
                            "Authorization": `Bearer ${tokenSet.access_token}`,
                        }
                    }
                    fetch(`${import.meta.env.VITE_FLEETDASH_CORE_SCHEME}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream/ticket`, options)
                        .then(response => response.json())
                        .then(response => {
                            if(typeof(response.error) === 'undefined') {
                                this._stream_ticket = response.ticket;
                                this.session_fleet_id = response.fleetId;
                                localStorage.setItem('stream_ticket', this._stream_ticket);
                                localStorage.setItem('session_fleet_id', this.session_fleet_id.toString());
                                resolve(this._stream_ticket);
                            } else {
                                reject(response.error);
                            }
                        })
                        .catch(error => {
                            reject(error);
                        });
                }
            });
        },
        resetTicket() {
            this._stream_ticket = '';
            this.session_fleet_id = 0;
            localStorage.removeItem('stream_ticket');
            localStorage.removeItem('session_fleet_id');
        },
        startSession(tokenSet : TokenSet) : Promise<void> {
            const options = {
                method: 'POST',
                headers: {
                    "Authorization": `Bearer ${tokenSet.access_token}`,
                }
            }
            return fetch(`${import.meta.env.VITE_FLEETDASH_CORE_SCHEME}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/session/start`, options)
                .then(response => response.json())
                .then(response => {
                    if(typeof(response.error) !== 'undefined') {
                        throw response.error;
                    }
                })
        },
        shutdownLocalStream() : void {
            this.resetTicket();
            this.active = false;
        },
    }
});
