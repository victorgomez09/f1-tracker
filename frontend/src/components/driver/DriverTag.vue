<script setup lang="ts">
const props = defineProps({
  color: String,
  short: String,
  position: Number,
});
</script>

<template>
  <div
    class="flex w-fit items-center justify-between gap-0.5 rounded-lg px-1 py-1 font-black bg-zinc-700 !min-w-[3.5rem]"
    :class="{ '!min-w-[5.5rem]': props.position }" :style="{ backgroundColor: `#${props.color}` }">
    <p v-if="props.position" class="px-1 text-xl leading-none text-white">
      {{ props.position }}
    </p>

    <div class="flex h-min w-min items-center justify-center rounded-md bg-white px-1"
      :class="{ '!w-full': !props.position }">
      <p class="text-zinc-700" :style="{ color: `#${props.color}` }">
        {{ props.short }}
      </p>
    </div>
  </div>
</template>
