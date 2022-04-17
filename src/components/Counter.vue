<script setup lang="ts">
import {computed} from 'vue';
import numeral from 'numeral';

const props = defineProps<{
    title: string;
    subtitle: string | null;
    count: number;
}>();

let readableCount = computed(() => numeral(props.count).format('0'));

</script>

<template>
    <div class="border rounded drop-shadow text-center bg-zinc-800 h-full flex flex-col justify-center">
        <div class="text-2xl mt-4">{{title}}</div>
        <div class="text-zinc-400 italic">{{subtitle}}</div>
        <div class="text-6xl font-extrabold mb-5">{{readableCount}}</div>
    </div>
</template>
