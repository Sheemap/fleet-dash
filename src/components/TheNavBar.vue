<script setup>
import TheAuthenticator from "./TheAuthenticator.vue";
</script>

<template>
    <nav class="border-b border-slate-500 bg-slate-700 flex">
      <span class="my-auto ml-5 text-3xl font-bold grow">FleetDash</span>
      <span class="mx-5">
        <slot/>
      </span>
      <span class="my-auto mr-5">
        <TheAuthenticator/>
      </span>
    </nav>
</template>
