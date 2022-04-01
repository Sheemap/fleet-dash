<script setup lang="ts">
import {computed} from 'vue';
import * as numeral from 'numeral';

const props = defineProps<{
    title: string;
    count: number;
    countColor: string;
}>();

let readableCount = computed(() => numeral(props.count).format('0[.]00a'));

</script>

<template>
    <div class="border rounded drop-shadow m-3 p-2 text-center bg-zinc-800">
        <div class="text-2xl mt-4">{{title}}</div>
        <div class="text-6xl font-extrabold my-10" :class="[countColor]">{{readableCount}}</div>
    </div>
</template>
