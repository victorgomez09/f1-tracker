<script setup lang="ts">
defineProps({
  playing: Boolean,
  onClick: Function,
});
</script>

<template>
  <div
    class="flex h-8 w-8 cursor-pointer items-center justify-center"
    @click="onClick!"
  >
    <svg
      v-if="!playing"
      width="24"
      height="24"
      view-box="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M5 3L19 12L5 21V3Z"
        fill="white"
        stroke="white"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
    <svg
      v-if="playing"
      width="24"
      height="24"
      view-box="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M10 4H6V20H10V4Z"
        fill="white"
        stroke="white"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <path
        d="M18 4H14V20H18V4Z"
        fill="white"
        stroke="white"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </div>
</template>
