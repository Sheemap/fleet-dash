<script setup>
import {inject, ref} from "vue";

let totalIncomingShield = ref(0);

const emitter = inject("emitter");
emitter.on("IncomingShieldEvent", (evt) => {
  totalIncomingShield.value += evt.Amount;
});

</script>

<template>
<div class="border p-5 m-5 border-blue-400 text-center">
  <h1>Incoming Shield Reps</h1>
  {{totalIncomingShield}}
</div>
</template>
