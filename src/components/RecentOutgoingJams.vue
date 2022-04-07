<script setup lang="ts">
import {inject, reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';
import {useFleetStore} from "../js/fleetStore";
import {useUserStore} from "../js/userStore";

interface JammedTarget {
  name: string
  shipId: number
  seconds: number
}

const fleetStore = useFleetStore();
const userStore = useUserStore();

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

  userStore.getActiveToken().then(token => {
    fleetStore.fetchItemId(evt.Ship, token).then(shipId => {
      let item = {
        name: evt.Pilot,
        shipId: shipId,
        seconds: seconds
      };
      let dupes = jams.filter(jam => jam.name === evt.Pilot);
      if (dupes.length > 0) {
        dupes.map(x => x.seconds = seconds);
      }
      else{
        jams.push(item);
      }
    });
  });
});

function removeJam(pilot: string) {
  jams.filter(x => x.name === pilot).forEach(jam => {
    jams.splice(jams.indexOf(jam), 1);
  });
}
</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full overflow-auto">
      <div class="text-2xl my-3">Outgoing Jams</div>

      <div v-if="jams.length === 0" class="py-5 text-zinc-400 italic">No one currently jammed</div>
      <div v-for="jammed in jams" class="py-1">
        <PlayerShipCardProgressCountdown :player-name="jammed.name" :ship-id="jammed.shipId" :seconds="jammed.seconds" :expiration-callback="removeJam" :item-key="jammed.name" />
      </div>
  </div>
</template>
