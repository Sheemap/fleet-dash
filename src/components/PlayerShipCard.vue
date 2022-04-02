<script setup lang="ts">
import {useFleetStore} from "../js/fleetStore";
import {useUserStore} from "../js/userStore";
import {ref} from "vue";

const props = defineProps<{
    shipId: number;
    playerName: string;
    value: string;
}>();

const fleetStore = useFleetStore();
const userStore = useUserStore();

let shipName = ref("");

userStore.getActiveToken().then(token => {
    fleetStore.fetchItemName(props.shipId, token).then(ship => {
        shipName.value = ship;
    });
});

const imgUrl = `https://images.evetech.net/types/${props.shipId}/icon`
</script>

<template>
    <div class="p-1 bg-slate-600 w-full rounded">
        <div class="w-full lg:max-w-full lg:flex">
            <div class="h-12 w-12" ><img :alt="shipName" class="rounded" :src="imgUrl" /></div>
            <div class="px-3">
                <div>{{ playerName }}</div>
                <div class="text-xs">{{ shipName }}</div>
            </div>
            <div class="p-3 flex-grow text-xl text-right">{{ value }}</div>
        </div>
        <slot></slot>
    </div>
</template>