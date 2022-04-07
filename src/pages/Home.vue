<script setup lang="ts">
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import AvgAmountOverTimeCounter from "../components/AvgAmountOverTimeCounter.vue"
import RecentIncomingJams from "../components/RecentIncomingJams.vue"
import RecentOutgoingJams from "../components/RecentOutgoingJams.vue"

import { useEventStore } from "../js/eventStore";
import {inject} from "vue";
import FleetOverview from "../components/FleetOverview.vue";
import TheFleetUpdater from "../components/TheFleetUpdater.vue";
import LogiDpsSlider from "../components/LogiDpsSlider.vue";

const eventStore = useEventStore();

const emitter = inject("emitter");
</script>

<template>
  <div class="bg-zinc-900 text-slate-100 h-full pb-96 md:pb-0 md:h-screen">
    <TheNavBar />
    <div v-if="eventStore.active">

        <div class="grid gap-3 grid-cols-1 grid-rows-auto md:grid-cols-4 md:grid-rows-dash md:w-full md:p-3">
          <AvgAmountOverTimeCounter title="Incoming DPS" event-type="IncomingDamageEvent" :period-seconds=15 />
          <AvgAmountOverTimeCounter title="Outgoing DPS" event-type="OutgoingDamageEvent" :period-seconds=15 />
          <div class="row-span-3">
            <RecentIncomingJams />
          </div>
          <div class="row-span-3">
            <RecentOutgoingJams />
          </div>
          <div class="row-span-2">
            <FleetOverview/>
          </div>
          <LogiDpsSlider :period-seconds="15" />

        </div>
      </div>

    <TheEventStream />
    <TheFleetUpdater />
  </div>
</template>
