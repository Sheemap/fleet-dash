import {defineStore} from "pinia";

export const useEventStore = defineStore('event', {
    state: () => {
        return {
            event_log: [],

        }
    },
});
