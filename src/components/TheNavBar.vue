<template>
    <nav class="border-b border-slate-500 bg-slate-700 flex justify-center content-center">
      <div class="m-5 text-3xl font-bold grow">FleetDash</div>
      <TheAuthenticator/>
    </nav>
</template>

<script>
import TheAuthenticator from "./TheAuthenticator.vue";
export default {
  name: "TheNavBar",
  components: {TheAuthenticator},
}
</script>

<style scoped>

</style>
