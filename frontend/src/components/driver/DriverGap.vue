<script setup lang="ts">
const props = defineProps({
  gapToLeader: String,
  gapToFront: String,
});
</script>

<template>
  <div class="grid grid-rows-2 grid-flow-col items-center justify-items-end">
    <!-- Gap to pilot ahead -->
    <div class="grid grid-rows-1 grid-flow-col items-center gap-2 w-full">
      <span class="font-semibold justify-self-start">Gap</span>
      <p class="text-lg font-semibold leading-none justify-self-end">
        {{ !!props.gapToFront ? props.gapToFront : "+--.---" }}
      </p>
    </div>

    <!-- Gap to leader -->
    <div class="grid grid-rows-1 grid-flow-col items-center gap-2 w-full">
      <span class="font-semibold justify-self-start">Ldr</span>
      <p class="text-lg font-semibold leading-none justify-self-end">
        {{ !!props.gapToLeader ? props.gapToLeader : "---.---" }}
      </p>
    </div>
  </div>
</template>
