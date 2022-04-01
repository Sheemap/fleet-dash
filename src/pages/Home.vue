<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import TotalDamageOut from "../components/TotalDamageOut.vue";
import TotalDamageIn from "../components/TotalDamageIn.vue";
import TotalShieldIn from "../components/TotalShieldIn.vue";
import AvgAmountOverTimeCounter from "../components/AvgAmountOverTimeCounter.vue"
import RecentIncomingJams from "../components/RecentIncomingJams.vue"

import { useEventStore } from "../js/eventStore";
import {inject} from "vue";
import IncomingDPS from "../components/IncomingDPS.vue";
import FleetOverview from "../components/FleetOverview.vue";
import TheFleetUpdater from "../components/TheFleetUpdater.vue";

const eventStore = useEventStore();

const emitter = inject("emitter");

function test(){
  emitter.emit('Test-Event', {Amount: Math.floor(Math.random() * 3000)})
  setTimeout(test, 2000)
}

// test();
</script>

<template>
  <div class="bg-zinc-900 text-slate-100 h-full pb-96 md:pb-0 md:h-screen">
    <TheNavBar />

    <div class="p-3">

    <div class="grid  gap-3 grid-cols-1 grid-rows-auto md:grid-cols-4 md:grid-rows-3">
      <AvgAmountOverTimeCounter event-type="IncomingDamageEvent" period-seconds=5 />
      <AvgAmountOverTimeCounter event-type="OutgoingDamageEvent" period-seconds=5 />
      <!-- <PlayerShipCardProgressCountdown player-name="Hello" ship-id="587" seconds="10" /> -->
      <div class="row-span-3">
        <RecentIncomingJams />
      </div>
      <div class="row-span-3">
        <RecentIncomingJams />
      </div>

      <div class="row-span-2">
        <FleetOverview/>
      </div>

      <div v-if="eventStore.active">
        <IncomingDPS />
        <TotalDamageOut />
        <TotalDamageIn />
        <TotalShieldIn />
      </div>

      <TheEventStream />
      <TheFleetUpdater />
    </div>
  </div>
  </div>
</template>
