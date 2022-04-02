<script setup lang="ts">
import {computed} from 'vue';
import * as numeral_ from 'numeral';
const numeral = numeral_;

const props = defineProps<{
    title: string;
    count: number;
}>();

let readableCount = computed(() => numeral(props.count).format('0[.]00a'));

</script>

<template>
    <div class="border rounded drop-shadow text-center bg-zinc-800">
        <div class="text-2xl mt-4">{{title}}</div>
        <div class="text-6xl font-extrabold my-10">{{readableCount}}</div>
    </div>
</template>
