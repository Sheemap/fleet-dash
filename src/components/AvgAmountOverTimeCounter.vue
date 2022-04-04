<script setup lang="ts">
import Counter from "../components/Counter.vue"

import {inject, ref} from 'vue';

const props = defineProps<{
    eventType: string;
    periodSeconds: number;
}>();

let runningAvg = ref(0);
let count = ref(0);

function removeValue(value){
    if (count.value - 1  <= 0) {
        runningAvg.value = 0;
    } else {
        runningAvg.value = ((runningAvg.value * count.value) - value) / (count.value - 1) || 0;
    }
    count.value--;
}

const emitter = inject("emitter");
emitter.on(props.eventType, (evt) => {
    count.value++;
    runningAvg.value = runningAvg.value + ((evt.Amount - runningAvg.value) / count.value) || 0;
  setTimeout(removeValue, props.periodSeconds * 1000, evt.Amount);
});
</script>

<template>
    <Counter :title="eventType" :count="runningAvg" />
</template>