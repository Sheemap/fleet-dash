<script setup lang="ts">
import {computed, inject, reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';
import {useUserStore} from "../js/userStore";
import {useFleetStore} from "../js/fleetStore";

interface JammedTarget {
  name: string
  characterId: number
  shipId: number
  seconds: number
  timestamp: number
}

interface JammedTargetKey {
  characterId: number
  timestamp: number
}

const userStore = useUserStore();
const fleetStore = useFleetStore();

let jams : JammedTarget[] = reactive([]);
let uniqueJams = computed(() => {
  let sortedJams = jams.sort((a, b) => (b.timestamp + b.seconds * 1000) - (a.timestamp + a.seconds * 1000));
  // Get unique jams, ordered by timestamp, with respect to character ID
  return sortedJams.reduce((uniqueJams, jam) => {
    if (uniqueJams.length === 0 || uniqueJams[uniqueJams.length - 1].characterId !== jam.characterId) {
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
emitter.on("IncomingJamEvent", (evt) => {
    let seconds = getSeconds(evt.Weapon);
    fleetStore.fetchCharacterName(evt.CharacterID).then((name) => {
      jams.push({
        name: name,
        characterId: evt.CharacterID,
        shipId: evt.CharacterShipTypeID,
        seconds: seconds,
        timestamp: new Date(evt.Timestamp).getTime(),
      });
    });
});

function removeJam(key: JammedTargetKey) {
  jams.filter(x => x.characterId === key.characterId &&
      x.timestamp === key.timestamp).forEach(jam => {
    jams.splice(jams.indexOf(jam), 1);
  });
}
</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full overflow-auto">
    <div class="text-2xl my-3">Incoming Jams</div>

    <div v-if="jams.length === 0" class="py-5 text-zinc-400 italic">No one currently jammed</div>
    <div v-for="jammed in uniqueJams" class="py-1">
      <PlayerShipCardProgressCountdown :key="jammed.characterId+jammed.timestamp"
                                       :player-name="jammed.name"
                                       :ship-id="jammed.shipId"
                                       :seconds="jammed.seconds"
                                       :timestamp="jammed.timestamp"
                                       :expiration-callback="removeJam"
                                       :item-key="{ characterId: jammed.characterId, timestamp: jammed.timestamp }" />
    </div>

  </div>
</template>
