<script setup lang="ts">

import SummaryTextItem from "./SummaryTextItem.vue";

const props = defineProps<{
    shipId: number;
    count: number;
}>();

const imgUrl = `https://images.evetech.net/types/${props.shipId}/icon`
</script>

<template>
  <div class="max-w-full flex">
      <div class="h-6 w-6" ><img alt="hello" class="rounded" :src="imgUrl" /></div>
      <span class="px-3">
        <SummaryTextItem count="5" text="Ship Name"/>
      </span>
  </div>
</template>