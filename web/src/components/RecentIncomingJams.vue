<script setup lang="ts">
import {computed, getCurrentInstance, inject, reactive} from 'vue';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue';
import {useUserStore} from "../js/userStore";
import {useFleetStore} from "../js/fleetStore";

const keyStr = getCurrentInstance()?.vnode.key;
const id = typeof keyStr === "string" ? keyStr.split("_")[1] : "";

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

function removeFromGrid(): void {
  userStore.dash_layout = userStore.dash_layout.filter((item) => !item.i.endsWith(id));
  localStorage.removeItem("widget_settings_" + id);
}
</script>

<template>
  <div class="border rounded drop-shadow text-center bg-zinc-800 h-full overflow-auto">
    <svg v-if="userStore.widget_drawer_open && id !== ''"
         @click="removeFromGrid"
         class="fixed z-50 top-0 left-0 m-1 config-button cursor-pointer transition ease-in-out opacity-70 hover:opacity-100"
         xmlns="http://www.w3.org/2000/svg"
         height="30px"
         viewBox="0 0 24 24"
         width="30px"
         fill="#DC2626">
      <path d="M0 0h24v24H0z" fill="none"/>
      <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"/>
    </svg>

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
