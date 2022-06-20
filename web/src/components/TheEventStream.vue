<script setup lang="ts">
  import { inject, ref } from "vue";
  import { useUserStore } from "../js/userStore";
  import { useEventStore } from "../js/eventStore";
  import useToast from "../js/toast";

  const emitter = inject("emitter");
  const userStore = useUserStore();
  const eventStore = useEventStore();
  let toast = useToast();

  let authenticated = ref(false);
  let starting = ref(false);

  function tryStartStream(startSession = true){
    starting.value = true;

    userStore.getActiveToken().then(tokenSet => {
      eventStore.getStreamTicket(tokenSet).then(ticket => {
        const socketScheme = import.meta.env.VITE_FLEETDASH_CORE_SCHEME === 'https' ? 'wss' : 'ws';
        const url = `${socketScheme}://${import.meta.env.VITE_FLEETDASH_CORE_URL}/api/event-stream?ticket=${ticket}`;

        const ws = new WebSocket(url);

        ws.addEventListener('open', function open() {
          console.log('Connection has been established.');
          eventStore.active = true;

          let pingInterval = setInterval(pingFunc, 5000);

          function pingFunc() {
            if(ws.readyState !== WebSocket.OPEN) {
              return;
            }
            if (!eventStore.active) {
              starting.value = false;
              ws.close(1000, 'Event stream closed');
              clearInterval(pingInterval);
              return;
            }

            ws.send(JSON.stringify({
              type: 'ping'
            }));
          }
        });

        ws.addEventListener('message', function incoming(data) {
          let jsonData = JSON.parse(data.data)
          console.log(jsonData);
          for (let event of jsonData){
            emitter.emit(event.Type, event);
          }
        });

        ws.addEventListener('close', function close(event) {
          console.log('Connection has been closed.');
          console.log(event);

          if (event.reason === "invalid ticket"){
            eventStore.active = false;
            eventStore.resetTicket();
          }

          // If its been over a minute since last error, reset the counter
          if (eventStore.last_stream_failure_time < Date.now() - (1000 * 60)){
            eventStore.stream_failure_count = 0;
          }

          // If no reason for close, we assume it's a request timeout, or some other retryable event
          // else its a permanent error
          if(event.reason !== "" && event.reason !== "invalid ticket"){
            toast.error('Connection closed: ' + event.reason);
            eventStore.shutdownLocalStream();
            starting.value = false;
          } else if (!event.wasClean) {
            // If we have tried to reconnect too many times, we dont want to keep trying
            if(eventStore.stream_failure_count < 5)
            {
              if (eventStore.stream_failure_count > 0){
                toast.warning("Connection closed, retrying. Attempt: " + eventStore.stream_failure_count);
              }
              // Try to reconnect with an exponential backoff
              setTimeout(tryStartStream, Math.pow(2, eventStore.stream_failure_count) * 1000, startSession);
            } else{
              toast.error('We have tried to reconnect too many times. Please refresh the page to try again.', {
                duration: 0
              });
              eventStore.active = false;
              eventStore.resetTicket();
              starting.value = false;
            }

            // Increment the failure counter and set timer
            eventStore.stream_failure_count++;
            eventStore.last_stream_failure_time = Date.now();
          }
        });
      })
        .catch(err => {
          if (startSession && err === 'not in active session') {
            eventStore.startSession(tokenSet)
                .then(() => {
                  tryStartStream();
                });
          } else if (startSession) {
            if (err == "not in fleet"){
              toast.error("You are not in a fleet!\nFleets are refreshed every 60 seconds. Give it some time if you believe this is a mistake.", {
                duration: 15000
              });
            } else {
              toast.error("Could not start stream: " + err);
            }
            starting.value = false;
          } else{
            starting.value = false;
          }
        });
    })
    .catch(_ => starting.value = false);
  }

  // Try to start the stream if we are authenticated
  userStore.getActiveToken().then(_ =>{
      authenticated.value = true;
      tryStartStream(false);
  })
  .catch(_ =>{
    // No token stored
    authenticated.value = false;
  })
</script>

<template>
  <div v-if="!eventStore.active && authenticated" class="grid place-content-center my-5">
    <button v-if="!starting" class="bg-emerald-600 border border-emerald-800 py-2 px-3 rounded hover:bg-emerald-700 transition-all"
            v-on:click="tryStartStream">Join Session</button>
    <button v-else class="bg-emerald-700 border border-emerald-800 py-2 px-3 rounded" disabled>
      <svg class="animate-spin inline-block w-5 h-5 mb-1 mr-2 border-b-2 rounded-full" viewBox="0 0 24 24"></svg>Join Session</button>
  </div>
</template>
