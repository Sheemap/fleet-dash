<script setup lang="ts">
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import { useEventStore } from "../js/eventStore";
import TheDashboard from "../components/TheDashboard.vue";
import TheWidgetDrawer from "../components/TheWidgetDrawer.vue";
import {ref} from "vue";

const eventStore = useEventStore();

const drawerOpened = ref(false);
const dashboardRef = ref(null);

</script>

<template>
  <div class="bg-zinc-900 text-slate-100 pb-96 md:pb-0 md:h-screen h-full">

      <TheNavBar />

      <TheEventStream />

      <div v-if="eventStore.active">
        <div id="content" class="flex gap-7 justi">
          <div class="w-full" :class="drawerOpened && 'w-3/4'">
            <TheDashboard :update-ref="(el) => dashboardRef = el" />
          </div>
          <div v-if="drawerOpened" class="w-1/4">
            <TheWidgetDrawer :dashboard-ref="dashboardRef" />
          </div>
          <div class="absolute bottom-0 right-0 m-4 p-2 border rounded bg-slate-700 border-slate-500">
            <button class="place-content-center" @click="() => drawerOpened = !drawerOpened">Add Widgets</button>
          </div>
        </div>
      </div>

  </div>
</template>
