import {defineStore} from "pinia";
import {TokenSet} from "./userStore";

export const useEventStore = defineStore('event', {
    state: () => {
        return {
            active: false,
            stream_failure_count: 0,
            last_stream_failure_time: 0,
            _stream_ticket: localStorage.getItem('stream_ticket') || '',
        }
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
                            this._stream_ticket = response.ticket;
                            localStorage.setItem('stream_ticket', this._stream_ticket);
                            resolve(this._stream_ticket);
                        })
                        .catch(error => {
                            reject(error);
                        });
                }
            });
        },
        resetTicket() {
            this._stream_ticket = '';
            localStorage.removeItem('stream_ticket');
        },
    }
});
