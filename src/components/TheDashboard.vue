<script setup lang="ts">
import IncomingDamage from "./IncomingDamage.vue";
import OutgoingDamage from "./OutgoingDamage.vue";
import FifteenSecondsLogiDpsSlider from "./FifteenSecondsLogiDpsSlider.vue";
import RecentIncomingJams from "./RecentIncomingJams.vue";
import RecentOutgoingJams from "./RecentOutgoingJams.vue";
import FleetOverview from "./FleetOverview.vue";
import IncomingLogi from "./IncomingLogi.vue";
import TheFleetUpdater from "./TheFleetUpdater.vue";

import {useUserStore} from "../js/userStore";

import {
  INCOMING_DAMAGE_WIDGET_ID,
  OUTGOING_DAMAGE_WIDGET_ID,
  INCOMING_JAM_WIDGET_ID,
  OUTGOING_JAM_WIDGET_ID,
  OVERVIEW_WIDGET_ID,
  SLIDER_WIDGET_ID,
  INCOMING_LOGI_WIDGET_ID
} from '../js/constants';

const props = defineProps<{
  updateRef: any;
}>();

const userStore = useUserStore();


function renderFromKey(key : string) {
  let widgetId = key.split('_')[0];
  switch (widgetId) {
    case INCOMING_DAMAGE_WIDGET_ID:
      return IncomingDamage;
    case OUTGOING_DAMAGE_WIDGET_ID:
      return OutgoingDamage;
    case SLIDER_WIDGET_ID:
      return FifteenSecondsLogiDpsSlider;
    case INCOMING_JAM_WIDGET_ID:
      return RecentIncomingJams;
    case OUTGOING_JAM_WIDGET_ID:
      return RecentOutgoingJams;
    case OVERVIEW_WIDGET_ID:
      return FleetOverview;
    case INCOMING_LOGI_WIDGET_ID:
      return IncomingLogi;
    default:
      return null;
  }
}

function onLayoutUpdate(layout) {
  userStore.updateLayout(layout);
}
</script>

<template>
  <grid-layout
      :ref="(el) => updateRef(el)"
      :layout.sync="userStore.dash_layout"
      :col-num="12"

      :is-draggable="true"
      :is-resizable="true"
      :is-mirrored="false"
      :vertical-compact="true"
      :margin="[10, 10]"
      :use-css-transforms="true"
      @layout-updated="onLayoutUpdate">
    <grid-item v-for="item in userStore.dash_layout"
               :x="item.x"
               :y="item.y"
               :w="item.w"
               :h="item.h"
               :i="item.i"
               :key="item.i">
      <component :is="renderFromKey(item.i)" />
    </grid-item>
  </grid-layout>

  <TheFleetUpdater />
</template>