<script setup lang="ts">
import {computed, inject, reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';
import {useFleetStore} from "../js/fleetStore";
import {useUserStore} from "../js/userStore";

interface JammedTarget {
  name: string
  shipId: number
  seconds: number
  timestamp: number
}

interface JammedTargetKey {
  name: string
  timestamp: number
}

const fleetStore = useFleetStore();
const userStore = useUserStore();

let jams : JammedTarget[] = reactive([]);
let uniqueJams = computed(() => {
  let sortedJams = jams.sort((a, b) => (b.timestamp + b.seconds * 1000) - (a.timestamp + a.seconds * 1000));
  // Get unique jams, ordered by timestamp, with respect to character ID
  return sortedJams.reduce((uniqueJams, jam) => {
    if (uniqueJams.length === 0 || uniqueJams[uniqueJams.length - 1].name !== jam.name) {
      uniqueJams.push(jam);
    }
    return uniqueJams;
  }, [] as JammedTarget[]);
});

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
  userStore.getActiveToken().then(token => {
    fleetStore.fetchItemId(evt.Ship, token).then(shipId => {
      let seconds = getSeconds(evt.Weapon);
      jams.push({
        name: evt.Pilot,
        shipId: shipId,
        seconds: seconds,
        timestamp: new Date(evt.Timestamp).getTime(),
      });
    });
  });
});

function removeJam(key: JammedTargetKey) {
  jams.filter(x => x.name === key.name &&
      x.timestamp === key.timestamp).forEach(jam => {
    jams.splice(jams.indexOf(jam), 1);
  });
}
</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full overflow-auto">
      <div class="text-2xl my-3">Outgoing Jams</div>

      <div v-if="jams.length === 0" class="py-5 text-zinc-400 italic">No one currently jammed</div>
      <div v-for="jammed in uniqueJams" class="py-1">
        <PlayerShipCardProgressCountdown
            :key="jammed.name + jammed.timestamp"
            :player-name="jammed.name"
            :ship-id="jammed.shipId"
            :seconds="jammed.seconds"
            :timestamp="jammed.timestamp"
            :expiration-callback="removeJam"
            :item-key="{ name: jammed.name, timestamp: jammed.timestamp }" />
      </div>
  </div>
</template>
