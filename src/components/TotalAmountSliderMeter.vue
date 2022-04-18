<script setup lang="ts">
import SliderMeter from "../components/SliderMeter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
  title: string;
  subtitle: string;
  leftEvents: string[];
  rightEvents: string[];
  periodSeconds: number;
}>();

let runningRightTotal = ref(0);
let runningLeftTotal = ref(0);

const emitter = inject("emitter");

for(let event of props.rightEvents){
  emitter.on(event, addRightValue);
}

function removeRightValue(value){
  runningRightTotal.value -= value;
}

function addRightValue(evt) {
  runningRightTotal.value += evt.Amount;
  setTimeout(removeRightValue, props.periodSeconds * 1000, evt.Amount);
}


for(let event of props.leftEvents){
  emitter.on(event, addLeftValue);
}

function removeLeftValue(value){
  runningLeftTotal.value -= value;
}

function addLeftValue (evt) {
  runningLeftTotal.value += evt.Amount;
  setTimeout(removeLeftValue, props.periodSeconds * 1000, evt.Amount);
}
</script>

<template>
  <SliderMeter :title="title" :subtitle="subtitle" :left-value="runningLeftTotal" :right-value="runningRightTotal" text-suffix="" />
</template>
