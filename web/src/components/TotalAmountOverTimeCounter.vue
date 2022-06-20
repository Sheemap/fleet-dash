<script setup lang="ts">
import Counter from "../components/Counter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
    title: string;
    eventType: string | string[];
    periodSeconds: number;
}>();

const subtitle = `Total over past ${props.periodSeconds} seconds`;

let runningTotal = ref(0);

function removeValue(value){
  runningTotal.value -= value;
}

function addValue(evt){
  runningTotal.value += evt.Amount;
  setTimeout(removeValue, props.periodSeconds * 1000, evt.Amount);
}

const emitter = inject("emitter");

if (props.eventType instanceof Array) {
  props.eventType.forEach(eventType => {
    emitter.on(eventType, addValue);
  });
}else{
  emitter.on(props.eventType, addValue);
}
</script>

<template>
    <Counter :title="title" :count="runningTotal" :subtitle="subtitle" />
</template>