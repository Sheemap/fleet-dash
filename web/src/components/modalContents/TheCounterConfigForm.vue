<script setup lang="ts">
import {ValidModalConfigs,TotalCounterConfig} from "../../js/shared";
import Multiselect from "@vueform/multiselect";
import {EVENT_TYPES, MULTI_SELECT_STYLE} from "../../js/constants";
import {reactive} from "vue";

const props = defineProps<{
    initialConfig: TotalCounterConfig;
    saveConfiguration: (config: ValidModalConfigs) => void;
}>();

let currentConfig = reactive(props.initialConfig);

currentConfig.title = props.initialConfig.title || "Title";
currentConfig.eventTypes = props.initialConfig.eventTypes || [];

function updateConfig(){
  props.saveConfiguration(currentConfig);
}

</script>

<template>
      <input v-model="currentConfig.title" class="bg-slate-700 w-full mt-0 block text-center px-0.5 border-0 border-b-2 border-zinc-500 my-2 focus:ring-0 focus:border-zinc-300 focus:bg-slate-800 hover:bg-slate-800" type="text" placeholder="Title" @change="updateConfig" />
      <input v-model="currentConfig.periodSeconds" class="bg-slate-700 w-full mt-0 block text-center px-0.5 border-0 border-b-2 border-zinc-500 mb-2 focus:ring-0 focus:border-zinc-300 focus:bg-slate-800 hover:bg-slate-800" type="number" placeholder="Period Seconds" @keyup="updateConfig" />
      <span class="w-[90%]">
        <Multiselect v-model="currentConfig.eventTypes" mode="multiple" :searchable="true" placeholder="Event types" :options="EVENT_TYPES" :classes="MULTI_SELECT_STYLE" @change="updateConfig" />
      </span>
</template>