<script setup lang="ts">
const props = defineProps<{
    shipId: number;
    playerName: string;
    value: string;
}>();

const imgUrl = `https://images.evetech.net/types/${props.shipId}/icon`
</script>

<template>
    <div class="p-1 bg-slate-600 w-full rounded">
        <div class="w-full lg:max-w-full lg:flex">
            <div class="h-12 w-12" ><img class="rounded" :src="imgUrl" /></div>
            <div class="px-3">
                <div>{{ playerName }}</div>
                <div class="text-xs">{{ shipId }}</div>
            </div>
            <div class="p-3 flex-grow text-xl text-right">{{ value }}</div>
        </div>
        <slot></slot>
    </div>
</template>