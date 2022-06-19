<script setup lang="ts">
import {ValidModalConfigs,SliderMeterConfig} from "../../js/shared";
import Multiselect from "@vueform/multiselect";
import {EVENT_TYPES, MULTI_SELECT_STYLE} from "../../js/constants";
import {reactive} from "vue";

const props = defineProps<{
    initialConfig: SliderMeterConfig;
    saveConfiguration: (config: ValidModalConfigs) => void;
}>();

let currentConfig = reactive(props.initialConfig);

currentConfig.title = props.initialConfig.title;
currentConfig.subtitle = props.initialConfig.subtitle;
currentConfig.leftEvents = props.initialConfig.leftEvents || [];
currentConfig.rightEvents = props.initialConfig.rightEvents || [];

function updateConfig(){
  props.saveConfiguration(currentConfig);
}

</script>

<template>
  <input v-model="currentConfig.title" class="bg-slate-700 mt-0 block w-full my-2 text-center px-0.5 border-0 border-b-2 border-zinc-500 focus:ring-0 focus:border-zinc-300 focus:bg-slate-800 hover:bg-slate-800" type="text" placeholder="Title" @change="updateConfig" />
  <input v-model="currentConfig.subtitle" class="bg-slate-700 mt-0 block w-full my-2 text-center px-0.5 border-0 border-b-2 border-zinc-500 focus:ring-0 focus:border-zinc-300 focus:bg-slate-800 hover:bg-slate-800" type="text" placeholder="Subtitle" @change="updateConfig" />
  <input v-model="currentConfig.periodSeconds" class="bg-slate-700 mt-0 block w-full my-2 text-center px-0.5 border-0 border-b-2 border-zinc-500 focus:ring-0 focus:border-zinc-300 focus:bg-slate-800 hover:bg-slate-800" type="number" placeholder="Period Seconds" @keyup="updateConfig" />
  <span class="w-[90%] p-3">
    <Multiselect v-model="currentConfig.leftEvents" mode="multiple" :searchable="true" placeholder="Left side events" :multipleLabel="(vals) => `${vals.length} left side events`" :options="EVENT_TYPES" :classes="MULTI_SELECT_STYLE" @change="updateConfig" />
  </span>
  <span class="w-[90%] p-3">
    <Multiselect v-model="currentConfig.rightEvents" mode="multiple" :searchable="true" placeholder="Right side events" :multipleLabel="(vals) => `${vals.length} right side events`" :options="EVENT_TYPES" :classes="MULTI_SELECT_STYLE" @change="updateConfig" />
  </span>
</template>