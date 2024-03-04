<script setup lang="ts">
import type { PropType } from "vue";

import DriverTag from "./DriverTag.vue";
import DriverDrs from "./DriverDrs.vue";
import DriverRpm from "./DriverRpm.vue";
import { F1Driver } from "../models/driver.model";

const props = defineProps({
  driver: Object as PropType<F1Driver>,
  position: String,
});
</script>

<template>
  <div
    class="flex items-center gap-2 p-2 h-18"
    :class="[
      {
        'opacity-50':
          props.driver?.status === 'OUT' ||
          props.driver?.status === 'RETIRED' ||
          props.driver?.status === 'STOPPED',
      },
      { 'bg-violet-800 bg-opacity-30': props.driver?.lapTimes.best.fastest },
      {
        'bg-red-800 bg-opacity-30': false, // TODO use this for danger zone in quali
      },
    ]"
  >
    <DriverTag
      :position="props.position"
      :teamColor="props.driver?.teamColor"
      :short="props.driver?.short"
    />

    <DriverDrs
      :drs="props.driver?.drs"
      :positionsChanged="props.driver?.positionChange"
    />

    <DriverRpm
      :rpm="props.driver?.metrics.rpm"
      :gear="props.driver?.metrics.gear"
      :speed="props.driver?.metrics.speed"
      :status="props.driver?.status!"
    />
  </div>
</template>
