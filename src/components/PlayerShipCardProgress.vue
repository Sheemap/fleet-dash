<script setup lang="ts">
import PlayerShipCard from './PlayerShipCard.vue'

const props = defineProps<{
    shipId: number;
    playerName: string;
    value: string;
    progressPercent: number;
}>();
</script>

<template>
    <PlayerShipCard :ship-id="shipId" :player-name="playerName" :value="value" >
        <div class="bg-slate-500 rounded">
            <div class="bg-green-500 h-1 rounded" :style="{ width: progressPercent + '%' }"></div>
        </div>
    </PlayerShipCard>
</template>