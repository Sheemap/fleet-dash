<template>
    <nav class="border flex">
        <a class="grow m-5 text-3xl font-bold">FleetDash</a>
        <a v-if="this.user.authenticated" class="m-5 text-3xl font-bold">Yuh</a>
        <a v-else class="m-5 text-3xl font-bold"><img src="../assets/eve-login.png"/></a>
    </nav>
</template>

<script>
export default {
  name: "TheNavBar",
  props: {
    user: Object,
  },
}
</script>

<style scoped>

</style>
