<script setup lang="ts">
import { ref } from 'vue'
import PlayerShipCardProgress from './PlayerShipCardProgress.vue'

const props = defineProps<{
    shipId: number;
    playerName: string;
    seconds: number;
}>();

let curSecond = ref(props.seconds);
let progPercent = ref(100);

function updateProgress(totalMili){
    curSecond.value = props.seconds - Math.floor(totalMili / 1000);
    progPercent.value = Math.ceil(curSecond.value / props.seconds * 100);

    if (progPercent.value > 0){
        setTimeout(updateProgress, 1000, totalMili + 1000)
    }
}

updateProgress(0);
</script>

<template>
    <PlayerShipCardProgress :ship-id="shipId" :player-name="playerName" :value="curSecond + 's'" :progress-percent="progPercent" />
</template>