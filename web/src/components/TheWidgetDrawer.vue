<script setup lang="ts">
import Counter from "./Counter.vue";
import SliderMeter from "./SliderMeter.vue";
import RecentIncomingJams from "./RecentIncomingJams.vue";
import RecentOutgoingJams from "./RecentOutgoingJams.vue";
import {Ref} from "vue";
import {createUUID} from "../js/shared";

import {useUserStore} from "../js/userStore";
import {TOTAL_OVER_TIME_WIDGET_ID, TOTAL_SLIDER_WIDGET_ID, INCOMING_JAM_WIDGET_ID, OUTGOING_JAM_WIDGET_ID, OVERVIEW_WIDGET_ID} from "../js/constants";
import AddOverlay from "./AddOverlay.vue";
import WidgetDisplayFleetOverview from "./WidgetDisplayFleetOverview.vue";

const props = defineProps<{
  dashboardRef: Ref;
}>();

const userStore = useUserStore();

function append(widgetId: string) {
  let max = 0;
  for(let item of userStore.dash_layout){
    if(item.y + item.h > max){
      max = item.y + item.h;
    }
  }
  userStore.appendToLayout({
    x: 0,
    y: max,
    w: 3,
    h: 1,
    i: `${widgetId}_${createUUID()}`,
  });
}
</script>

<template>
  <div class="h-full px-3 py-5 rounded bg-opacity-80 bg-gray-800 grid grid-cols-1 gap-3">
    <AddOverlay :widget-id="TOTAL_OVER_TIME_WIDGET_ID" :append-func="append">
      <Counter title="Total Counter" subtitle="Tracks totals for an event" :count=0 />
    </AddOverlay>
    <AddOverlay :widget-id="TOTAL_SLIDER_WIDGET_ID" :append-func="append">
      <SliderMeter title="Sliding Meter" subtitle="Compare two totals" :left-value=0 :right-value=0 text-suffix="" />
    </AddOverlay>
    <AddOverlay :widget-id="INCOMING_JAM_WIDGET_ID" :append-func="append">
      <RecentIncomingJams />
    </AddOverlay>
    <AddOverlay :widget-id="OUTGOING_JAM_WIDGET_ID" :append-func="append">
      <RecentOutgoingJams />
    </AddOverlay>
    <AddOverlay :widget-id="OVERVIEW_WIDGET_ID" :append-func="append">
      <WidgetDisplayFleetOverview />
    </AddOverlay>
  </div>
</template>
