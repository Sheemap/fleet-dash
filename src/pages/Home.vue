<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import HelloWorld from '../components/HelloWorld.vue'
import TheNavBar from "../components/TheNavBar.vue";
</script>

<template>
  <div class="bg-zinc-900 h-screen text-slate-100">
    <TheNavBar />
    <img alt="Vue logo" src="../assets/logo.png" />
  </div>

</template>

<script>
export default {
  name: "Home",
  computed: {
    user() {
      return {
        "name": "Sekap",
        "profile_pic": "yuhhhh",
        "authenticated": true,
      }
      // return this.$store.state.user;
    }
  }
}
</script>

<style>
/*#app {*/
/*  font-family: Avenir, Helvetica, Arial, sans-serif;*/
/*  -webkit-font-smoothing: antialiased;*/
/*  -moz-osx-font-smoothing: grayscale;*/
/*  text-align: center;*/
/*  color: #2c3e50;*/
/*}*/
</style>
