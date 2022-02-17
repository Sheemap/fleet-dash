<script setup lang="ts">
  import AvgPerSecondOverTimeCounter from "./AvgPerSecondOverTimeCounter.vue";
</script>

<template>
  <AvgPerSecondOverTimeCounter title="Incoming DPS" event-type="IncomingDamageEvent" period-seconds="15" />
</template>
