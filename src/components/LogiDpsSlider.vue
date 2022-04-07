<script setup lang="ts">
import SliderMeter from "../components/SliderMeter.vue"

import {inject, ref} from 'vue';
import {addToRunningAvg, removeFromRunningAvg} from "../js/shared";

const props = defineProps<{
  periodSeconds: number;
}>();

let runningDpsAvg = ref(0);
let dpsCount = ref(0);

let runningLogiAvg = ref(0);
let logiCount = ref(0);

const emitter = inject("emitter");

function removeDpsValue(value){
  runningDpsAvg.value = removeFromRunningAvg(runningDpsAvg.value, dpsCount.value, value);
  dpsCount.value--;
}

emitter.on("IncomingDamageEvent", (evt) => {
  dpsCount.value++;
  runningDpsAvg.value = addToRunningAvg(runningDpsAvg.value, dpsCount.value, evt.Amount);
  setTimeout(removeDpsValue, props.periodSeconds * 1000, evt.Amount);
});

function removeLogiValue(value){
  runningLogiAvg.value = removeFromRunningAvg(runningLogiAvg.value, logiCount.value, value);
  logiCount.value--;
}

emitter.on("IncomingShieldEvent", addLogiEvent);
emitter.on("IncomingArmorEvent", addLogiEvent);
emitter.on("IncomingHullEvent", addLogiEvent);

function addLogiEvent(evt){
  logiCount.value++;
  runningLogiAvg.value = addToRunningAvg(runningLogiAvg.value, logiCount.value, evt.Amount);
  setTimeout(removeLogiValue, props.periodSeconds * 1000, evt.Amount);
}
</script>

<template>
  <SliderMeter title="Logi vs DPS" subtitle="How the logi is holding up" :left-value="runningLogiAvg" :right-value="runningDpsAvg" text-suffix="/s" />
</template>
