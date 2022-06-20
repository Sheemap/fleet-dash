<script setup lang="ts">

defineProps<{
  widgetId: string,
  appendFunc: (widgetId: string) => void,
}>();

</script>

<template>
  <div class="relative group cursor-pointer" @click="() => appendFunc(widgetId)">
    <div class="h-full transition group-hover:opacity-50">
      <slot></slot>
    </div>
    <div class="absolute left-0 top-0 transition opacity-0 group-hover:opacity-100 bg-emerald-300 m-2 rounded-full">
      <img alt="Add Widget" src="https://img.icons8.com/material-outlined/48/000000/add.png"/>
    </div>
  </div>
</template>