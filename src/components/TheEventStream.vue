<script setup>
  import { useUserStore } from "../js/userStore";
  import { useEventStore } from "../js/eventStore";

  const userStore = useUserStore();
  const eventStore = useEventStore();

  userStore.getActiveToken().then(tokenSet => {
    eventStore.getStreamTicket(tokenSet).then(ticket => {
      const socketScheme = import.meta.env.VITE_FLEETDASH_CORE_SCHEME === 'https' ? 'wss' : 'ws';
      const ws = new WebSocket(`${socketScheme}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`);

      ws.addEventListener('open', function open() {
        console.log('Connection has been established.');
      });

      ws.addEventListener('message', function incoming(data) {
        console.log(data.data);
      });

      ws.addEventListener('close', function close(event) {
        console.log('Connection has been closed. Reason:', event.reason);
      });
    });
  });
</script>

<template>
Streaming :)
</template>
