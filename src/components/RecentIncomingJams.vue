<script setup lang="ts">
import {reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';


interface JammedTarget {
  name: string
  shipId: number
  seconds: number
}


let jams : JammedTarget[] = reactive([]);

setTimeout(() => {
  let item = {
    name: 'TestPilot',
    shipId: 587,
    seconds: 10
  };
  jams.push(item);

  setTimeout(() => {
    jams.splice(jams.indexOf(item), 1);
  }, 10000);

}, 1000);

setTimeout(() => {
  let item = {
    name: 'TestPilot',
    shipId: 587,
    seconds: 8
  };
  jams.push(item);

  setTimeout(() => {
    jams.splice(jams.indexOf(item), 1);
  }, 8000);
}, 2000);

setTimeout(() => {
  let item = {
    name: 'TestPilot',
    shipId: 587,
    seconds: 10
  };
  jams.push(item);

  setTimeout(() => {
    jams.splice(jams.indexOf(item), 1);
  }, 10000);
}, 4000);

</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full">
    <div class="text-2xl my-3">Incoming Jams</div>

    <div v-if="jams.length === 0" class="py-5 text-zinc-400 italic">No one currently jammed</div>
    <div v-for="jammed in jams" class="py-2">
      <PlayerShipCardProgressCountdown :player-name="jammed.name" :ship-id="jammed.shipId" :seconds="jammed.seconds" />
    </div>

  </div>
</template>
