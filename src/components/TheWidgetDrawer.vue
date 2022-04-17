<script setup lang="ts">
import Counter from "./Counter.vue";
import SliderMeter from "./SliderMeter.vue";
import RecentIncomingJams from "./RecentIncomingJams.vue";
import RecentOutgoingJams from "./RecentOutgoingJams.vue";
import FleetOverview from "./FleetOverview.vue";
import {Ref} from "vue";
import {createUUID} from "../js/shared";

import {useUserStore} from "../js/userStore";
import {INCOMING_DAMAGE_WIDGET_ID, SLIDER_WIDGET_ID, INCOMING_JAM_WIDGET_ID, OUTGOING_JAM_WIDGET_ID, OVERVIEW_WIDGET_ID} from "../js/constants";

const props = defineProps<{
  dashboardRef: Ref;
}>();

const userStore = useUserStore();

function append(widgetId: string) {
  let max = 0;
  for(let item of userStore.dash_layout){
    if(item.y > max){
      max = item.y;
    }
  }
  userStore.appendToLayout({
    x: 0,
    y: max + 1,
    w: 3,
    h: 1,
    i: `${widgetId}_${createUUID()}`,
  });
}
</script>

<template>
  <div class="flex flex-no-wrap bg-gray-200 h-full">
    <div class="w-full absolute sm:relative bg-gray-800 shadow md:h-full grid grid-cols-1 gap-3">
      <span class="mx-2 mt-2" @click="() => append(INCOMING_DAMAGE_WIDGET_ID)">
        <Counter title="Counter" subtitle="" :count=0 />
      </span>
      <div class="mx-2 mt-2" @click="() => append(SLIDER_WIDGET_ID)">
        <SliderMeter title="Sliding Meter" subtitle="" :left-value=0 :right-value=0 text-suffix="" />
      </div>
      <div class="mx-2 mt-2" @click="() => append(INCOMING_JAM_WIDGET_ID)">
        <RecentIncomingJams />
      </div>
      <div class="mx-2 mt-2" @click="() => append(OUTGOING_JAM_WIDGET_ID)">
        <RecentOutgoingJams />
      </div>
      <div class="mx-2 mt-2" @click="() => append(OVERVIEW_WIDGET_ID)">
        <FleetOverview />
      </div>
    </div>
  </div>
</template>
