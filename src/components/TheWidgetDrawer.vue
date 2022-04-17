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
import AddOverlay from "./AddOverlay.vue";

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
  <div class="h-full px-3 py-5 rounded bg-opacity-80 bg-gray-800 grid grid-cols-1 gap-3">
    <AddOverlay :widget-id="INCOMING_DAMAGE_WIDGET_ID" :append-func="append">
      <Counter title="Counter" subtitle="" :count=0 />
    </AddOverlay>
    <AddOverlay :widget-id="SLIDER_WIDGET_ID" :append-func="append">
      <SliderMeter title="Sliding Meter" subtitle="" :left-value=0 :right-value=0 text-suffix="" />
    </AddOverlay>
    <AddOverlay :widget-id="INCOMING_JAM_WIDGET_ID" :append-func="append">
      <RecentIncomingJams />
    </AddOverlay>
    <AddOverlay :widget-id="OUTGOING_JAM_WIDGET_ID" :append-func="append">
      <RecentOutgoingJams />
    </AddOverlay>
    <AddOverlay :widget-id="OVERVIEW_WIDGET_ID" :append-func="append">
      <FleetOverview />
    </AddOverlay>
  </div>
</template>
