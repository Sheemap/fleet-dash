<script setup lang="ts">
  import TotalAmountOverTimeCounter from "./TotalAmountOverTimeCounter.vue";
  const incomingLogiEvents = ["IncomingShieldEvent", "IncomingArmorEvent", "IncomingHullEvent"];

</script>

<template>
  <TotalAmountOverTimeCounter title="Incoming Logi" :event-type="incomingLogiEvents" :period-seconds=15 />

</template>
