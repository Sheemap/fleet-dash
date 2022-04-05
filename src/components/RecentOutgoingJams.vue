<script setup lang="ts">
import {inject, reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';


interface JammedTarget {
  name: string
  shipId: number
  seconds: number
}


let jams : JammedTarget[] = reactive([]);

function getSeconds(weapon: string) {
  let seconds: number;
  if (weapon.includes("EC-")) {
    seconds = 5;
  } else {
    seconds = 20;
  }
  return seconds;
}

const emitter = inject("emitter");
emitter.on("OutgoingJamEvent", (evt) => {
  let seconds = getSeconds(evt.Weapon);
  let item = {
    name: evt.Pilot,
    shipId: 587,
    seconds: seconds,
  }
  jams.push(item);

  setTimeout(() => {
    jams.splice(jams.indexOf(item), 1);
  }, item.seconds * 1000);
});

</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full overflow-auto">
      <div class="text-2xl my-3">Outgoing Jams</div>

      <div v-if="jams.length === 0" class="py-5 text-zinc-400 italic">No one currently jammed</div>
      <div v-for="jammed in jams" class="py-1">
        <PlayerShipCardProgressCountdown :player-name="jammed.name" :ship-id="jammed.shipId" :seconds="jammed.seconds" />
      </div>
  </div>
</template>
