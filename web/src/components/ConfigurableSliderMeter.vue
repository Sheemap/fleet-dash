<script setup lang="ts">
import {computed, getCurrentInstance, reactive} from "vue";
import {useUserStore} from "../js/userStore";

import TotalAmountSliderMeter from "./TotalAmountSliderMeter.vue";
import {SliderMeterConfig, ValidModalConfigs} from "../js/shared";

defineProps<{
  showConfigureModal: (configType: string, config: ValidModalConfigs, saveConfig: (ValidModalConfigs) => void) => void;
}>();

const userStore = useUserStore();
const keyStr = getCurrentInstance()?.vnode.key;

const id = typeof keyStr === "string" ? keyStr.split("_")[1] : "";

let settings = reactive(JSON.parse(localStorage.getItem("widget_settings_" + id) || "{}"));

const configurationIsValid = computed((): boolean => {
  let titleValid = typeof settings.title !== 'undefined' && settings.title !== "";
  let subtitleValid = typeof settings.subtitle !== 'undefined' && settings.subtitle !== "";
  let periodValid = typeof settings.periodSeconds !== 'undefined' && !isNaN(settings.periodSeconds) && settings.periodSeconds > 0;
  let leftEventsValue = typeof settings.leftEvents !== 'undefined' && settings.leftEvents.length > 0;
  let rightEventsValue = typeof settings.rightEvents !== 'undefined' && settings.rightEvents.length > 0;

  return titleValid && subtitleValid && periodValid && leftEventsValue && rightEventsValue;
});

function commitSettings(newConfig: SliderMeterConfig): void {
  settings.title = newConfig.title;
  settings.subtitle = newConfig.subtitle;
  settings.periodSeconds = newConfig.periodSeconds;
  settings.leftEvents = newConfig.leftEvents;
  settings.rightEvents = newConfig.rightEvents;
  localStorage.setItem("widget_settings_" + id, JSON.stringify(settings));
}

function removeFromGrid(): void {
  userStore.dash_layout = userStore.dash_layout.filter((item) => !item.i.endsWith(id));
  localStorage.removeItem("widget_settings_" + id);
}
</script>

<template>
  <span>
    <svg class="fixed z-50 top-0 right-0 m-1 config-button cursor-pointer transition ease-in-out opacity-50 hover:opacity-100"
         @click="() => showConfigureModal('SliderMeterConfig', settings, commitSettings)"
         xmlns="http://www.w3.org/2000/svg"
         height="24px"
         viewBox="0 0 24 24"
         width="24px"
         fill="#ffffff">
      <g>
        <path d="M0,0h24v24H0V0z" fill="none"/>
        <path d="M19.14,12.94c0.04-0.3,0.06-0.61,0.06-0.94c0-0.32-0.02-0.64-0.07-0.94l2.03-1.58c0.18-0.14,0.23-0.41,0.12-0.61 l-1.92-3.32c-0.12-0.22-0.37-0.29-0.59-0.22l-2.39,0.96c-0.5-0.38-1.03-0.7-1.62-0.94L14.4,2.81c-0.04-0.24-0.24-0.41-0.48-0.41 h-3.84c-0.24,0-0.43,0.17-0.47,0.41L9.25,5.35C8.66,5.59,8.12,5.92,7.63,6.29L5.24,5.33c-0.22-0.08-0.47,0-0.59,0.22L2.74,8.87 C2.62,9.08,2.66,9.34,2.86,9.48l2.03,1.58C4.84,11.36,4.8,11.69,4.8,12s0.02,0.64,0.07,0.94l-2.03,1.58 c-0.18,0.14-0.23,0.41-0.12,0.61l1.92,3.32c0.12,0.22,0.37,0.29,0.59,0.22l2.39-0.96c0.5,0.38,1.03,0.7,1.62,0.94l0.36,2.54 c0.05,0.24,0.24,0.41,0.48,0.41h3.84c0.24,0,0.44-0.17,0.47-0.41l0.36-2.54c0.59-0.24,1.13-0.56,1.62-0.94l2.39,0.96 c0.22,0.08,0.47,0,0.59-0.22l1.92-3.32c0.12-0.22,0.07-0.47-0.12-0.61L19.14,12.94z M12,15.6c-1.98,0-3.6-1.62-3.6-3.6 s1.62-3.6,3.6-3.6s3.6,1.62,3.6,3.6S13.98,15.6,12,15.6z"/>
      </g>
    </svg>
    <svg v-if="userStore.widget_drawer_open"
         @click="removeFromGrid"
         class="fixed z-50 top-0 left-0 m-1 config-button cursor-pointer transition ease-in-out opacity-70 hover:opacity-100"
         xmlns="http://www.w3.org/2000/svg"
         height="30px"
         viewBox="0 0 24 24"
         width="30px"
         fill="#DC2626">
      <path d="M0 0h24v24H0z" fill="none"/>
      <path d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"/>
    </svg>

    <span>
      <TotalAmountSliderMeter v-if="configurationIsValid"
                              :title="settings.title"
                              :subtitle="settings.subtitle"
                              :left-events="settings.leftEvents"
                              :right-events="settings.rightEvents"
                              :period-seconds=settings.periodSeconds />
      <span v-else class="border rounded drop-shadow text-center bg-zinc-800 text-red-500 font-bold text-xl h-full flex flex-col justify-center">
        Configuration is not valid
      </span>

    </span>
  </span>

</template>