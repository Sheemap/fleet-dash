<script setup lang="ts">
import {computed} from "vue";

const props = defineProps<{
    title: string;
    subtitle: string;
    textSuffix: string;
    leftValue: number;
    rightValue: number;
}>();



let leftPercent = computed(() => zeroIfNaN(Math.round((props.leftValue / getTotal()) * 100)));
let rightPercent = computed(() => zeroIfNaN(Math.round((props.rightValue / getTotal()) * 100)));

function getTotal() {
    return parseFloat(props.leftValue.toString()) + parseFloat(props.rightValue.toString());
}

function zeroIfNaN(value: number) {
    return isNaN(value) ? 0 : value;
}
</script>

<template>
    <div class="h-full border rounded drop-shadow text-center bg-zinc-800 flex flex-col justify-center">
        <div class="text-2xl mt-4">{{ title }}</div>
        <div class="text-zinc-400 italic">{{ subtitle }}</div>
      <div class="flex pt-1">
        <div class="h-7 w-1/2 rounded-l">{{ leftValue }}{{ textSuffix }}</div>
        <div class="h-7 w-1/2 rounded-r">{{ rightValue }}{{ textSuffix }}</div>
      </div>
      <div v-if="rightPercent !== 0 || leftPercent !== 0" class="flex pb-3 px-3">
        <div class="bg-gradient-to-r from-emerald-400 to-emerald-600 h-7 rounded-l border-zinc-400" :class="leftPercent !== 100 || 'rounded'" :style="`width: ${leftPercent || 0}%;`"></div>
        <div class="bg-gradient-to-l from-red-500 to-red-600 h-7 w-1/2 rounded-r" :class="rightPercent !== 100 || 'rounded'" :style="`width: ${rightPercent || 0}%;`"></div>
      </div>
      <div v-else class="flex pb-3 px-3">
        <div class="bg-gradient-to-r from-emerald-400 to-emerald-600 h-7 rounded-l border-zinc-400" style="width: 50%"></div>
        <div class="bg-gradient-to-l from-red-500 to-red-600 h-7 w-1/2 rounded-r" style="width: 50%"></div>
      </div>
    </div>
</template>
