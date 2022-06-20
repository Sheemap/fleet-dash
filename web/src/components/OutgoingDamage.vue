<script setup lang="ts">
  import TotalAmountOverTimeCounter from "./TotalAmountOverTimeCounter.vue";
</script>

<template>
  <TotalAmountOverTimeCounter title="Outgoing Damage" event-type="OutgoingDamageEvent" :period-seconds=15 />
</template>
