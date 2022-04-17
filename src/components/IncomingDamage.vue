<script setup lang="ts">
  import TotalAmountOverTimeCounter from "./TotalAmountOverTimeCounter.vue";
</script>

<template>
  <TotalAmountOverTimeCounter title="Incoming Damage" event-type="IncomingDamageEvent" :period-seconds=15 />
</template>
