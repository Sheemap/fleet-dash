<script setup lang="ts">
import {computed} from "vue";

const props = defineProps<{
    title: string;
    subtitle: string;
    textSuffix: string;
    leftValue: number;
    rightValue: number;
}>();



let leftPercent = computed(() => Math.round((props.leftValue / getTotal()) * 100));
let rightPercent = computed(() => Math.round((props.rightValue / getTotal()) * 100));

function getTotal() {
    return parseFloat(props.leftValue) + parseFloat(props.rightValue);
}
</script>

<template>
    <div class="border rounded drop-shadow text-center bg-zinc-800">
        <div class="text-2xl mt-4">{{ title }}</div>
        <div class="text-zinc-400 italic">{{ subtitle }}</div>
      <div class="flex pt-1">
        <div class="h-7 w-1/2 rounded-l">{{ leftValue }}{{ textSuffix }}</div>
        <div class="h-7 w-1/2 rounded-r">{{ rightValue }}{{ textSuffix }}</div>
      </div>
      <div class="flex pb-3 px-3">
        <div class="bg-gradient-to-r from-emerald-400 to-emerald-600 h-7 rounded-l border-zinc-400" :class="leftPercent !== 100 || 'rounded'" :style="`width: ${leftPercent || 0}%;`"></div>
        <div class="bg-gradient-to-l from-red-500 to-red-600 h-7 w-1/2 rounded-r" :class="rightPercent !== 100 || 'rounded'" :style="`width: ${rightPercent || 0}%;`"></div>
      </div>
    </div>
</template>
