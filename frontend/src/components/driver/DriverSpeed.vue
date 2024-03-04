<script setup lang="ts">
const props = defineProps({
  speed: Number,
  status: String,
});
</script>

<template>
  <div class="flex flex-col justify-center text-center">
    <p class="font-bold text-lg">{{ props.speed }}</p>
    <p class="font-semibold">km/h</p>
  </div>
</template>
