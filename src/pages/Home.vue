<script setup lang="ts">
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import { useEventStore } from "../js/eventStore";
import TheDashboard from "../components/TheDashboard.vue";

const eventStore = useEventStore();

</script>

<template>
  <div class="bg-zinc-900 text-slate-100 h-full pb-96 md:pb-0 md:h-screen align-middle">
    <TheNavBar />

    <TheDashboard v-if="eventStore.active"/>

    <TheEventStream />
  </div>
</template>
