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
        const url = `${socketScheme}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`;

        const ws = new WebSocket(url);

        ws.addEventListener('open', function open() {
          console.log('Connection has been established.');
          eventStore.active = true;

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

          // If its been over a minute since last error, reset the counter
          if (eventStore.last_stream_failure_time < Date.now() - (1000 * 60)){
            eventStore.stream_failure_count = 0;
          }

          // If no reason for close, we assume it's a request timeout, or some other retryable event
          // else its a permanent error
          if(event.reason !== ""){
            // TODO: dont use alert, use a snackbar
            alert('Connection closed: ' + event.reason);
            eventStore.active = false;
            eventStore.resetTicket();
          } else{
            // If we have tried to reconnect too many times, we dont want to keep trying
            if(eventStore.stream_failure_count < 5)
            {
              // Try to reconnect with an exponential backoff
              setTimeout(tryStartStream, Math.pow(2, eventStore.stream_failure_count) * 1000);
            } else{
              // TODO: dont use alert, use a snackbar
              alert('We have tried to reconnect too many times. Please refresh the page to try again.');
              eventStore.active = false;
              eventStore.resetTicket();
            }

            // Increment the failure counter and set timer
            eventStore.stream_failure_count++;
            eventStore.last_stream_failure_time = Date.now();
          }
        });
      });
    });
  }
</script>

<template>
  <div v-if="!eventStore.active" class="grid place-content-center my-5">
    <button class="bg-emerald-600 border border-emerald-800 py-2 px-3 rounded hover:bg-emerald-700 transition-all"
            v-on:click="tryStartStream">Start Session</button>
  </div>
</template>
