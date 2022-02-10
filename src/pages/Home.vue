<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import TotalDamageOut from "../components/TotalDamageOut.vue";
import TotalDamageIn from "../components/TotalDamageIn.vue";
import TotalShieldIn from "../components/TotalShieldIn.vue";
</script>

<template>
  <div class="bg-zinc-900 h-screen text-slate-100">
    <TheNavBar />
    <TotalDamageOut />
    <TotalDamageIn />
    <TotalShieldIn />
    <TheEventStream />
  </div>

</template>
