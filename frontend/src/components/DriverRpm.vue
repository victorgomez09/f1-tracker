<script setup lang="ts">
const props = defineProps({
  rpm: Number,
  gear: Number,
  speed: Number,
  status: String,
});
</script>

<template>
  <div class="flex flex-col">
    <span
      class="text-sm inline-flex items-center justify-center rounded-md border-2 font-bold px-1 h-6 w-full"
      :class="[
        {
          'border-gray-500 text-gray-500':
            !props.status && !props.gear && !props.speed,
        },
        {
          'border-gray-300 text-gray-300':
            !props.status && !props.gear && props.speed,
        },
        { 'border-emerald-500 text-emerald-500': !props.status && props.gear },
        { 'border-cyan-500 text-cyan-500': props.status },
      ]"
    >
      {{ props.status ? "DRS" : "RPM" }}
    </span>

    <span
      :class="[
        {
          'text-success': props.rpm! > 0,
          'text-error': props.rpm! < 0,
          'text-neutral-content': !props.rpm,
        },
      ]"
    >
      {{ props.rpm }}
    </span>
  </div>
</template>
