<script setup lang="ts">
import SliderMeter from "../components/SliderMeter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
  title: string;
  subtitle: string;
  leftEvent: string;
  rightEvent: string;
  periodSeconds: number;
}>();

let runningRightTotal = ref(0);
let runningLeftTotal = ref(0);

const emitter = inject("emitter");

function removeRightValue(value){
  runningRightTotal.value -= value;
}

emitter.on(props.rightEvent, (evt) => {
  runningRightTotal.value += evt.Amount;
  setTimeout(removeRightValue, props.periodSeconds * 1000, evt.Amount);
});

function removeLeftValue(value){
  runningLeftTotal.value -= value;
}

emitter.on(props.leftEvent, (evt) => {
  runningLeftTotal.value += evt.Amount;
  setTimeout(removeLeftValue, props.periodSeconds * 1000, evt.Amount);
});
</script>

<template>
  <SliderMeter :title="title" :subtitle="subtitle" :left-value="runningLeftTotal" :right-value="runningRightTotal" text-suffix="" />
</template>
