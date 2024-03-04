<script setup lang="ts">
import { PropType, computed } from "vue";
import { Drs } from "../../models/driver.model";

const props = defineProps({
  positionsChanged: Number,
  drs: Object as PropType<Drs>,
  status: String,
});

const gain = computed(() => props.positionsChanged! > 0);
const loss = computed(() => props.positionsChanged! < 0);
const pit = props.status === "PIT" || props.status === "PIT OUT";
</script>

<template>
  <div class="flex flex-col">
    <span
      class="front-semibold text-center"
      :class="[
        {
          'text-success': gain,
          'text-error': loss,
          'text-neutral-content': !gain && !loss,
        },
      ]"
    >
      {{
        gain
          ? `+${props.positionsChanged}`
          : loss
          ? props.positionsChanged
          : "+0"
      }}
    </span>

    <span
      class="text-sm inline-flex items-center justify-center rounded-md border-2 font-bold h-6"
      :class="[
        {
          'border-gray-500 text-gray-500':
            !pit && !props.drs?.on && !props.drs?.possible,
        },
        {
          'border-gray-300 text-gray-300':
            !pit && !props.drs?.on && props.drs?.possible,
        },
        { 'border-emerald-500 text-emerald-500': !pit && props.drs?.on },
        { 'border-cyan-500 text-cyan-500': pit },
      ]"
    >
      {{ pit ? "PIT" : "DRS" }}
    </span>
  </div>
</template>
../../models/driver.model
