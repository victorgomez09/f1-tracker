<script setup lang="ts">
import { computed, isProxy, ref, toRaw } from "vue";

import Driver from "../components/driver/Driver.vue";
import RaceDetails from "../components/race-details/RaceDetails.vue";
import RaceControl from "../components/race-control/RaceControl.vue";
import RaceMap from "../components/race-map/RaceMap.vue";

import { viewMode } from "../store/viewMode.store";
import { dashboardData } from "../store/data.store";
import { sortPos } from "../utils/position.utils";
import RaceRadio from "../components/race-radios/RaceRadio.vue";
import { sortUtc } from "../utils/time.util";

const data = isProxy(dashboardData)
  ? toRaw(dashboardData.data)
  : dashboardData.data;
console.log("data from dash", data);
const { session } = data;
const { trackStatus } = data;
const { weather } = data;
const { extrapolatedClock } = data;
const { drivers } = data;
const { raceControlMessages } = data;
const { positionBatches } = data;
const { teamRadios } = data;
const { lapCount } = data;
const { carData } = data;

const liveTimingViewOption = ref<boolean>(false);
const checkChange = (element: any) => {
  viewMode.value = element.target.checked ? "TELEMETRY" : "PRETTY";
};

const driversSorted = computed(() => {
  return drivers.sort(sortPos);
});

const radiosSorted = computed(() => {
  return teamRadios.sort(sortUtc);
});
</script>

<template>
  <div class="flex flex-col flex-1 bg-base-300">
    <RaceDetails
      :session="session"
      :trackStatus="trackStatus"
      :weather="weather"
      :extrapolatedClock="extrapolatedClock"
      :laps="lapCount"
    />

    <div class="grid overflow-auto">
      <div class="flex justify-between items-center bg-base-100 bg-fixed">
        <h3 class="sticky top-0 font-bold text-lg p-2">Live Timing</h3>

        <div class="form-control">
          <label class="cursor-pointer label">
            <span class="label-text mr-2">
              {{ liveTimingViewOption ? "TELEMETRY" : "PRETTY" }}</span
            >
            <input
              type="checkbox"
              class="toggle toggle-primary"
              :checked="liveTimingViewOption"
              v-model="liveTimingViewOption"
              @change="checkChange"
            />
          </label>
        </div>
      </div>
      <div class="col-span-2 overflow-auto">
        <div
          class="grid gap-6 bg-base-100"
          :style="{
            gridTemplateColumns:
              '21px 52px 64px 64px 21px 90px 90px 52px 45px auto',
          }"
        >
          <p>POS</p>
          <p :style="{ textAlign: 'right' }">DRIVER</p>
          <p>GEAR/RPM</p>
          <p>SPD/PDL</p>
          <p>DRS</p>
          <p>TIME</p>
          <p>GAP</p>
          <p>TYRE</p>
          <p>INFO</p>
          <p>SECTORS</p>
        </div>

        <Driver
          v-for="driver in driversSorted"
          :driver="driver"
          :car-data="carData"
          :position="driver.position"
        />
      </div>

      <div
        class="grid grid-rows-1 grid-cols-8 overflow-auto"
        :style="{ maxHeight: '81vh' }"
      >
        <div class="flex flex-col col-span-3">
          <div class="flex justify-between items-center bg-base-100 bg-fixed">
            <h3 class="sticky top-0 font-bold text-lg p-2">Race Control</h3>

            <div class="form-control w-52">
              <label class="cursor-pointer label">
                <span class="label-text">Remember me</span>
                <input type="checkbox" class="toggle toggle-primary" checked />
              </label>
            </div>
          </div>
          <div class="overflow-auto h-full">
            <RaceControl :messages="raceControlMessages" />
          </div>
        </div>

        <div class="flex flex-col col-span-3">
          <h3 class="sticky top-0 font-bold text-lg bg-base-100 bg-fixed p-2">
            Race Map
          </h3>

          <div class="overflow-auto h-full">
            <RaceMap
              :circuit="session.circuitKey"
              :trackStatus="trackStatus"
              :windDirection="weather.wind_direction"
              :position-batches="positionBatches"
            />
          </div>
        </div>

        <div class="flex flex-col col-span-2">
          <h3
            class="sticky top-0 font-bold text-lg bg-base-100 bg-fixed p-2 z-[1]"
          >
            Race Radios
          </h3>

          <div class="overflow-auto h-full">
            <RaceRadio :team-radios="radiosSorted" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
