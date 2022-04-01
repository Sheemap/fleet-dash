<script setup lang="ts">
import Counter from "../components/Counter.vue"

import {computed, inject, ref} from 'vue';

const props = defineProps<{
    title: string;
    eventType: string;
    periodSeconds: number;
}>();

let runningTotal = ref(0);

function removeValue(value){
  runningTotal.value -= value;
}

const emitter = inject("emitter");
emitter.on(props.eventType, (evt) => {
    runningTotal.value += evt.Amount;
    setTimeout(removeValue, props.periodSeconds * 1000, evt.Amount);
});

let runningAvg = computed(() => runningTotal.value / props.periodSeconds);
</script>

<template>
    <Counter :title="title" :count="runningAvg" />
</template>
