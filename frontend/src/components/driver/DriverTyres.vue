<script setup lang="ts">
import { PropType, defineProps, computed, ref, watchEffect } from "vue";

import { Stint } from "../../models/driver.model";

const props = defineProps({
  stints: Object as PropType<Stint[]>,
});

const stops = computed(() => (props.stints ? props.stints.length - 1 : 0));
const currentStint = computed(() =>
  props.stints ? props.stints[props.stints.length - 1] : null
);
const unknownCompound = ![
  "soft",
  "medium",
  "hard",
  "intermediate",
  "wet",
].includes(currentStint.value?.compound ?? "");

const tyreCompound = ref();
watchEffect(async () => {
  tyreCompound.value = (
    await import(
      /* @vite-ignore */ `../../assets/tyres/${currentStint.value?.compound.toLowerCase()}.svg`
    )
  ).default;
});
</script>

<template>
  <div
    className="grid grid-rows-1 grid-flow-col items-center gap-3 place-self-start"
  >
    <div class="indicator">
      <span
        class="indicator-item badge badge-xs"
        :class="currentStint?.new ? 'badge-success' : 'badge-error'"
      ></span>
      <img
        :src="tyreCompound"
        :alt="currentStint.compound"
        class="h-7 w-7"
        v-if="currentStint && !unknownCompound"
      />
      <div
        className="flex h-8 w-8 items-center justify-center"
        v-if="currentStint && unknownCompound"
      >
        <p>?</p>
      </div>
      <div
        className="h-8 w-8 animate-pulse rounded-md bg-gray-700 font-semibold"
        v-if="!currentStint"
      />
    </div>
    <!-- TODO move this to a tooltip -->
    <div>
      <p className="font-bold leading-none">L {{ currentStint?.laps ?? 0 }}</p>
      <p className="text-sm font-medium leading-none text-gray-500">
        St {{ stops }}
      </p>
    </div>
  </div>
</template>
../../models/driver.model
