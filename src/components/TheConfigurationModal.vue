<script setup lang="ts">
import TheCounterConfigForm from "./modalContents/TheCounterConfigForm.vue";
import TheSliderConfigForm from "./modalContents/TheSliderConfigForm.vue";
import {TotalCounterConfig, ValidModalConfigs} from "../js/shared";

const props = defineProps<{
  configType: string;
  config: ValidModalConfigs;
  updateConfiguration: (config: ValidModalConfigs) => void;
  close: () => void;
}>();

let contents;
if (props.configType == "TotalCounterConfig") {
  contents = TheCounterConfigForm;
}
if (props.configType == "SliderMeterConfig") {
  contents = TheSliderConfigForm;
}
</script>

<template>

<div class="fixed top-0 left-0 right-0 bottom-0 flex flex-col justify-center items-center">
    <div class="bg-slate-700 rounded-lg shadow-lg p-5">
      <div class="flex justify-center">
        <div class="text-center">
          <h1 class="text-3xl font-bold">Modify Widget</h1>
        </div>
      </div>
      <div class="flex justify-center">
        <div class="text-center pb-5 w-full">
          <component :is="contents" :initial-config="config" :save-configuration="updateConfiguration" />
        </div>
      </div>
      <div class="flex justify-center">
        <div class="text-center">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" @click="close">
            Close
          </button>
        </div>
      </div>
    </div>
  </div>

</template>
