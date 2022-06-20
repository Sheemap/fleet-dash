<script setup lang="ts">
import { ref } from 'vue'
import PlayerShipCardProgress from './PlayerShipCardProgress.vue'

const props = defineProps<{
    shipId: number;
    playerName: string;
    seconds: number;
    timestamp: number;
    itemKey: any;
    expirationCallback: (key: any) => void;
}>();

let curSecond = ref(props.seconds);
let progPercent = ref(100);

let endTime = props.timestamp + props.seconds * 1000;

function updateProgress(){
    const now = Date.now();
    curSecond.value = Math.floor((endTime - now) / 1000);
    progPercent.value = Math.ceil( curSecond.value / props.seconds * 100);

    if (progPercent.value > 0){
        setTimeout(updateProgress, 1000)
    }
    else{
        props.expirationCallback(props.itemKey);
    }
}

updateProgress();
</script>

<template>
    <PlayerShipCardProgress :ship-id="shipId" :player-name="playerName" :value="curSecond + 's'" :progress-percent="progPercent" />
</template>