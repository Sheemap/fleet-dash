<script setup lang="ts">
import {defineProps, computed} from 'vue';
import numeral from 'numeral';
import PlayerShipCardProgressCountdown from './PlayerShipCardProgressCountdown.vue'

const props = defineProps<{
    title: string;
    count: number;
    countColor: string;
}>();

let readableCount = computed(() => numeral(props.count).format('0[.]00a'));

</script>

<template>
    <div class="border rounded drop-shadow m-3 p-2 max-w-xs text-center bg-zinc-800">
        <div class="text-2xl mt-4">{{title}}</div>
        <div class="text-6xl font-extrabold my-10" :class="[countColor]">{{readableCount}}</div>
    <PlayerShipCardProgressCountdown player-name="Hello" ship-id="587" seconds="10" />

    </div>
</template>