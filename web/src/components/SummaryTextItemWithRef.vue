<script setup lang="ts">

import SummaryTextItem from "./SummaryTextItem.vue";
import {Ref} from "vue";

defineProps<{
  text:  Ref<string>;
  count: number;
}>();

</script>

<template>
    <SummaryTextItem :count="count" :text="text.value" />
</template>