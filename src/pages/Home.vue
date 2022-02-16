<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import TotalDamageOut from "../components/TotalDamageOut.vue";
import TotalDamageIn from "../components/TotalDamageIn.vue";
import TotalShieldIn from "../components/TotalShieldIn.vue";
import AvgAmountOverTimeCounter from "../components/AvgAmountOverTimeCounter.vue"
import PlayerShipCard from "../components/PlayerShipCard.vue"
import PlayerShipCardProgressCountdown from "../components/PlayerShipCardProgressCountdown.vue"

import { useEventStore } from "../js/eventStore";
import {ref, inject} from "vue";

const eventStore = useEventStore();

const emitter = inject("emitter");

function test(){
  emitter.emit('Test-Event', {Amount: Math.floor(Math.random() * 3000)})
  setTimeout(test, 2000)
}

// test();
</script>

<template>
  <div class="bg-zinc-900 h-screen text-slate-100">

    <TheNavBar />
    <AvgAmountOverTimeCounter event-type="Test-Event" period-seconds=5 />
    <!-- <PlayerShipCardProgressCountdown player-name="Hello" ship-id="587" seconds="10" /> -->


    <div v-if="eventStore.active">
      <TotalDamageOut />
      <TotalDamageIn />
      <TotalShieldIn />
    </div>

    <TheEventStream />
  </div>
</template>
