<script setup lang="ts">
import SliderMeter from "../components/SliderMeter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
  periodSeconds: number;
}>();

let runningDpsTotal = ref(0);
let runningLogiTotal = ref(0);

const emitter = inject("emitter");

function removeDpsValue(value){
  runningDpsTotal.value -= value;
}

emitter.on("IncomingDamageEvent", (evt) => {
  runningDpsTotal.value += evt.Amount;
  setTimeout(removeDpsValue, props.periodSeconds * 1000, evt.Amount);
});

function removeLogiValue(value){
  runningLogiTotal.value -= value;
}

emitter.on("IncomingShieldEvent", addLogiEvent);
emitter.on("IncomingArmorEvent", addLogiEvent);
emitter.on("IncomingHullEvent", addLogiEvent);

function addLogiEvent(evt){
  runningLogiTotal.value += evt.Amount;
  setTimeout(removeLogiValue, props.periodSeconds * 1000, evt.Amount);
}
</script>

<template>
  <SliderMeter title="Logi vs DPS" subtitle="How the logi is holding up" :left-value="runningLogiTotal" :right-value="runningDpsTotal" text-suffix="" />
</template>
