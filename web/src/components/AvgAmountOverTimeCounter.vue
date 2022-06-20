<script setup lang="ts">
import Counter from "../components/Counter.vue"

import {inject, ref} from 'vue';
import {addToRunningAvg, removeFromRunningAvg} from "../js/shared";

const props = defineProps<{
    eventType: string;
    periodSeconds: number;
}>();

const subtitle = `Average over past ${props.periodSeconds} seconds`;

let runningAvg = ref(0);
let count = ref(0);

function removeValue(value){
    runningAvg.value = removeFromRunningAvg(runningAvg.value, count.value, value);
    count.value--;
}

const emitter = inject("emitter");
emitter.on(props.eventType, (evt) => {
    count.value++;
    runningAvg.value = addToRunningAvg(runningAvg.value, count.value, evt.Amount);
    setTimeout(removeValue, props.periodSeconds * 1000, evt.Amount);
});
</script>

<template>
    <Counter :title="eventType" :count="runningAvg" :subtitle="subtitle" />
</template>