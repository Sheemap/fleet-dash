<script setup lang="ts">
import Counter from "../components/Counter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
    eventType: string;
    periodSeconds: number;
}>();

const subtitle = `Total over past ${props.periodSeconds} seconds`;

let runningTotal = ref(0);

function removeValue(value){
  runningTotal.value -= value;
}

const emitter = inject("emitter");
emitter.on(props.eventType, (evt) => {
    runningTotal.value += evt.Amount;
    setTimeout(removeValue, props.periodSeconds * 1000, evt.Amount);
});
</script>

<template>
    <Counter :title="eventType" :count="runningTotal" :subtitle="subtitle" />
</template>