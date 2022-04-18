<script setup lang="ts">
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import { useEventStore } from "../js/eventStore";
import { useUserStore } from "../js/userStore";
import TheDashboard from "../components/TheDashboard.vue";
import TheWidgetDrawer from "../components/TheWidgetDrawer.vue";
import {ref} from "vue";

const eventStore = useEventStore();
const userStore = useUserStore();

const dashboardRef = ref(null);
</script>

<template>
  <div class="bg-zinc-900 text-slate-100 pb-96 md:pb-0 h-screen h-full">

      <TheNavBar />

      <TheEventStream />

      <div v-if="eventStore.active">
        <div id="content" class="flex gap-7">
          <div :class="userStore.widget_drawer_open ? 'w-3/4' : 'w-full'">
            <TheDashboard :update-ref="(el) => dashboardRef = el" />
          </div>
          <div v-if="userStore.widget_drawer_open" class="w-1/4 fixed right-0 top-0 h-full">
            <TheWidgetDrawer :dashboard-ref="dashboardRef" />
          </div>
          <div :class="userStore.widget_drawer_open ? 'opacity-100' : 'opacity-0'"
               class="fixed top-0 right-0 m-4 p-1 rounded-full bg-red-500 cursor-pointer rotate-45 transition ease-in-out hover:scale-110"
               @click="() => userStore.widget_drawer_open = !userStore.widget_drawer_open">
            <img alt="Close Widget Drawer" src="https://img.icons8.com/material-outlined/36/000000/add.png"/>
          </div>
          <div :class="!userStore.widget_drawer_open ? 'opacity-100' : 'opacity-0'"
               class="absolute bottom-0 right-0 m-6 p-1 rounded-full bg-emerald-500 cursor-pointer transition ease-in-out hover:scale-110"
               @click="() => userStore.widget_drawer_open = !userStore.widget_drawer_open">
            <img alt="Add Widget" src="https://img.icons8.com/material-outlined/30/000000/add.png"/>
          </div>
        </div>
      </div>

  </div>
</template>
