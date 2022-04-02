<script setup lang="ts">
import {useFleetStore} from "../js/fleetStore";
import {useUserStore} from "../js/userStore";
import SummaryTextItem from "./SummaryTextItem.vue";
import {ref} from "vue";

const props = defineProps<{
    shipId: number;
    count: number;
}>();

const imgUrl = `https://images.evetech.net/types/${props.shipId}/icon`;
let shipName = ref('');

const fleetStore = useFleetStore();
const userStore = useUserStore();

userStore.getActiveToken().then(token => {
  fleetStore.fetchItemName(props.shipId, token).then(name => {
    shipName.value = name;
  }).catch(err => {
    console.error(err);
    shipName.value = 'Unknown';
  });
}).catch(() => {
  shipName.value = 'Unknown';
});

</script>

<template>
  <div class="max-w-full flex">
      <div class="h-6 w-6" ><img alt="hello" class="rounded" :src="imgUrl" /></div>
      <span class="px-3">
        <SummaryTextItem :count=count :text="shipName" />
      </span>
  </div>
</template>