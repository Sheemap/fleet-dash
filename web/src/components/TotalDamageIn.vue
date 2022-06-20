<script setup>
import {inject, ref} from "vue";

let totalIncomingDamage = ref(0);

const emitter = inject("emitter");
emitter.on("IncomingDamageEvent", (evt) => {
  totalIncomingDamage.value += evt.Amount;
});

</script>

<template>
<div class="border p-5 m-5 border-blue-400 text-center">
  <h1>Incoming Damage</h1>
  {{totalIncomingDamage}}
</div>
</template>
