<script setup>
  import { useUserStore } from "../js/userStore";
  import { useEventStore } from "../js/eventStore";
  import {ref} from "vue";

  const userStore = useUserStore();
  const eventStore = useEventStore();

  let active = ref(false);
  userStore.getActiveToken().then(tokenSet => {
    eventStore.getStreamTicket(tokenSet).then(ticket => {
      let ws = new WebSocket(`ws://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`);

      ws.addEventListener('open', function open() {
        console.log('Connection has been established.');
        active.value = true;

        ws.send("Hey friend!")
        ws.send("How it be? :)")
    });


    });
  });
</script>

<template>
{{active}}
</template>
