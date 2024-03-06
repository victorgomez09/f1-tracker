<script setup lang="ts">
import type { PropType } from "vue";

import DriverTag from "./DriverTag.vue";
import DriverDrs from "./DriverDrs.vue";
import DriverRpm from "./DriverRpm.vue";
import DriverSpeed from "./DriverSpeed.vue";
import DriverLaps from "./DriverLaps.vue";
import DriverGap from "./DriverGap.vue";
import DriverSectors from "./DriverSectors.vue";
import DriverTyres from "./DriverTyres.vue";
import { F1Driver } from "../../models/driver.model";

const props = defineProps({
  driver: Object as PropType<F1Driver>,
  position: String,
});
</script>

<template>
  <div
    class="grid place-items-center items-center gap-1 py-1"
    :style="{
      gridTemplateColumns: '1.5fr 0.5fr 1.5fr 1fr 2fr 2fr 3fr 1.5fr',
    }"
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

    <DriverSpeed
      :speed="props.driver?.metrics.speed"
      :status="props.driver?.status!"
    />

    <DriverLaps
      :best="props.driver?.lapTimes.best"
      :last="props.driver?.lapTimes.last"
    />

    <DriverGap
      :gapToFront="props.driver?.gapToFront"
      :gapToLeader="props.driver?.gapToLeader"
    />

    <DriverSectors
      :sectors="props.driver?.sectors"
      :driverDisplayName="props.driver?.short"
    />

    <DriverTyres :stints="props.driver?.stints" />
  </div>
</template>
../../models/driver.model
