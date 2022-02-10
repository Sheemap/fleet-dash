<script setup>
  import { inject } from "vue";
  import { useUserStore } from "../js/userStore";
  import { useEventStore } from "../js/eventStore";

  const emitter = inject("emitter");
  const userStore = useUserStore();
  const eventStore = useEventStore();


  function tryStartStream(){
    userStore.getActiveToken().then(tokenSet => {
      eventStore.getStreamTicket(tokenSet).then(ticket => {
        const socketScheme = import.meta.env.VITE_FLEETDASH_CORE_SCHEME === 'https' ? 'wss' : 'ws';
        const ws = new WebSocket(`${socketScheme}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`);

        ws.addEventListener('open', function open() {
          console.log('Connection has been established.');

          function timeoutFunc() {
            ws.send(JSON.stringify({
              type: 'ping'
            }));

            if(ws.readyState === WebSocket.OPEN) {
              setTimeout(timeoutFunc, 5000);
            }
          }

          timeoutFunc();
        });

        ws.addEventListener('message', function incoming(data) {
          let jsonData = JSON.parse(data.data)
          console.log(data.data);
          emitter.emit(jsonData.Type, jsonData);
        });

        ws.addEventListener('close', function close(event) {
          console.log('Connection has been closed.');
          console.log(event);
          tryStartStream();
        });
      });
    });
  }

  tryStartStream();
</script>

<template>
</template>
