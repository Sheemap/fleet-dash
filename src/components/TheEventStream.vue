<script setup>
  import { useUserStore } from "../js/userStore";
  import { useEventStore } from "../js/eventStore";

  const userStore = useUserStore();
  const eventStore = useEventStore();

  userStore.getActiveToken().then(tokenSet => {
    eventStore.getStreamTicket(tokenSet).then(ticket => {
      const ws = new WebSocket(`ws://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`);

      ws.addEventListener('open', function open() {
        console.log('Connection has been established.');
    });

      ws.addEventListener('message', function incoming(data) {
        console.log(data.data);
      });

      ws.addEventListener('close', function close() {
        console.log('Connection has been closed.');
      });
    });
  });
</script>

<template>
Streaming :)
</template>
