<script setup>
import {inject, ref} from "vue";

let totalOutgoingDamage = ref(0);

const emitter = inject("emitter");
emitter.on("OutgoingDamageEvent", (evt) => {
  totalOutgoingDamage.value += evt.Amount;
});

</script>

<template>
<div class="border p-5 m-5 border-blue-400 text-center">
  <h1>Outgoing Damage</h1>
  {{totalOutgoingDamage}}
</div>
</template>
