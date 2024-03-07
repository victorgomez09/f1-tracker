<script setup lang="ts">
import { computed, type PropType } from "vue";

import { getTimeColor } from "../../utils/time.util";
import { viewMode } from "../../store/viewMode.store";
import DriverTag from "./DriverTag.vue";
import DriverDrs from "./DriverDrs.vue";
import DriverRpm from "./DriverRpm.vue";
import DriverSpeed from "./DriverSpeed.vue";
import DriverLaps from "./DriverLaps.vue";
import DriverGap from "./DriverGap.vue";
import DriverSectors from "./DriverSectors.vue";
import DriverTyres from "./DriverTyres.vue";
import { F1Driver } from "../../models/driver.model";
import { F1CarData } from "../../models/car.model";

const props = defineProps({
  driver: Object as PropType<F1Driver>,
  carData: Object as PropType<F1CarData>,
  position: String,
});

const gridCols = "21px 52px 64px 64px 21px 90px 90px 52px 45px auto";
const gridColsSmall = "18px 42px 60px 60px 18px 74px 74px 44px 38px auto";

const car = computed(
  () =>
    props.carData?.Entries[props.carData.Entries.length - 1].Cars[
      props.driver?.nr!
    ].Channels
);
const rpmPercent = computed(() => (car.value![0] / 15000) * 100);
const throttlePercent = computed(() => Math.min(100, car.value![4]));
const brakeApplied = computed(() => car.value![5] > 0);
const drs = computed(() => car.value![45]);
</script>

<template>
  <div
    v-if="viewMode === 'PRETTY'"
    class="grid place-items-center items-center gap-1 py-1 grid-cols-[6em_4.5em_5.5em_5em_8em_8em_28em_5.5em] lg:grid-cols-[1.5fr_0.5fr_1.5fr_1fr_2fr_2fr_3fr_1.5fr] w-fit"
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

  <div v-else>
    <div
      class="grid gap-6 overflow-x-auto border-b border-b-base-content"
      :style="{ gridTemplateColumns: gridCols }"
    >
      <span :class="{ 'text-violet-500': driver?.lapTimes.last.pb }">
        <span>P{{ position }}</span>
        <br />
        <span
          :class="[
            {
              'text-success': driver?.positionChange! > 0,
              'text-error': driver?.positionChange! < 0,
              'text-base-content': driver?.positionChange! == 0,
            },
          ]"
        >
          {{
            driver?.positionChange! > 0
              ? `+${driver?.positionChange}`
              : driver?.positionChange
          }}
        </span>
      </span>

      <span
        class="flex flex-col items-end text-nowrap text-end"
        :style="{ color: '#' + driver?.teamColor! }"
      >
        {{ driver?.nr }} {{ driver?.short }}

        <span>
          {{
            driver?.status === "OUT"
              ? "OUT"
              : driver?.status === "RETIRED"
              ? "RETIRED"
              : driver?.status === "STOPPED"
              ? "STOPPED"
              : driver?.status === "PIT"
              ? "PIT"
              : driver?.status === "PIT OUT"
              ? "PIT OUT"
              : null
          }}
        </span>
      </span>

      <span>
        <span title="Gear"> {{ car?.[3] }} </span>
        {{ " " }}
        <span title="RPM">{{ car?.[0] }}</span>
        <br />
        <progress
          class="progress progress-info"
          :value="rpmPercent"
          max="100"
        ></progress>
      </span>

      <span>
        <span class="w-full">{{ car![2].toString() }} km/h</span>
        <br />
        <progress
          class="progress progress-success"
          :value="throttlePercent"
          max="100"
        ></progress>
        <br />
        <progress
          class="progress progress-error"
          :value="brakeApplied ? 100 : 0"
          max="100"
        ></progress>
      </span>

      <span :class="{ 'text-success': drs, 'text-info': driver?.drs.possible }">
        DRS
      </span>

      <span>
        <span title="Last lap">
          Lst{{ " " }}
          <span
            :class="getTimeColor(driver?.lapTimes.last?.fastest!, driver?.lapTimes.last!.pb!)"
          >
            {{ driver?.lapTimes.last.value || "—" }}
          </span>
        </span>
        <br />
        <span title="Best lap">
          Bst{{ " " }}
          <span
            :class="getTimeColor(driver?.lapTimes.best?.fastest!, driver?.lapTimes.best!.pb!)"
          >
            {{ driver?.lapTimes.best?.value || "—" }}
          </span>
        </span>
      </span>

      <span class="flex flex-col items-end">
        <span
          :title="
            'Gap to car ahead ' + driver?.catchingFront ? ' (catching)' : ''
          "
        >
          Int{{ " " }}
          <span
            :class="{
              'text-success': driver?.catchingFront,
            }"
          >
            {{ driver?.gapToFront || "-" }}
          </span>
        </span>

        <span title="Gap to leader">
          Ldr{{ " " }} {{ driver?.gapToLeader || "-" }}
        </span>
      </span>
    </div>
  </div>
</template>
