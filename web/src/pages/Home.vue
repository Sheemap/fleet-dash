<script setup lang="ts">
import TheNavBar from "../components/TheNavBar.vue";
import TheEventStream from "../components/TheEventStream.vue";
import { useEventStore } from "../js/eventStore";
import {useUserStore} from "../js/userStore";
import {ValidModalConfigs} from "../js/shared";
import TheDashboard from "../components/TheDashboard.vue";
import TheWidgetDrawer from "../components/TheWidgetDrawer.vue";
import {ref} from "vue";
import TheConfigurationModal from "../components/TheConfigurationModal.vue";

const eventStore = useEventStore();
const userStore = useUserStore();

const dashboardRef = ref(null);


let showingModal = ref(false);
let configType = ref("");
let currentConfig = ref({});
let configureCallback = ref((_) => {});
function showModal(newConfigType: string, config: ValidModalConfigs, updateCallback: (config: ValidModalConfigs) => void) {
  configType.value = newConfigType
  currentConfig.value = config;
  configureCallback.value = updateCallback;
  showingModal.value = true;
}

function closeModal() {
  showingModal.value = false;
}
</script>

<template>
  <div class="bg-zinc-900 text-slate-100 pb-96 md:pb-0 h-screen h-full">

    <TheNavBar>
      <TheEventStream />
    </TheNavBar>

    <div v-if="eventStore.active">
      <div id="content" class="flex gap-7">
        <div :class="userStore.widget_drawer_open ? 'w-3/4' : 'w-full'">
          <TheDashboard :update-ref="(el) => dashboardRef = el" :show-configure-modal="showModal" />
        </div>
        <div v-if="userStore.widget_drawer_open" class="w-1/4 fixed right-0 top-0 h-full">
          <TheWidgetDrawer :dashboard-ref="dashboardRef" />
        </div>
        <div v-if="userStore.widget_drawer_open"
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

    <div v-else>
      <div class="flex flex-col items-center justify-center text-center h-full pt-10">
        <h1 class="text-3xl font-bold">Welcome to FleetDash</h1>
        <p class="text-center text-slate-100 py-1">
          A <span class="text-emerald-300">real-time</span> dashboard for your Eve Online fleets.
        </p>
        <p class="text-lg text-center p-5 font-bold">With easily digestible and customizable widgets, FleetDash shows you exactly how your fleet is performing</p>
        <div class="px-10">
          <h2 class="text-2xl font-bold pt-16">How does it work?</h2>
          <p>Each fleet member needs to run the FleetDash client and login to their accounts. This tool will read their toon's combat logs, and send any significant events to the core server. These are then collected and displayed to the FC right here, through the FleetDash website.</p>
        </div>
        <div class="px-10">
          <h2 class="text-2xl font-bold pt-16">Is this against the EULA?</h2>
          <p>No! FleetDash will only parse your combat logs, which is acceptable under the Eve EULA. FleetDash never sends any input to the game, and only collects data through the combat logs.</p>
        </div>
      </div>

    </div>

      <TheConfigurationModal v-if="showingModal" :config-type="configType" :config="currentConfig" :update-configuration="configureCallback" :close="closeModal" />

  </div>
</template>
