<script setup lang="ts">

let props = defineProps<{
  text: string;
  count: number;
}>();

</script>

<template>
    <span>
      <span class="font-bold">{{ count }}</span>
      {{ text }}
    </span>
</template>