<template>
    <nav class="border flex">
        <a class="grow m-5 text-3xl font-bold">FleetDash</a>
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
