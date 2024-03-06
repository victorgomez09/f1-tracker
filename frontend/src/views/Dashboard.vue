<script setup lang="ts">
import { computed, isProxy, toRaw } from "vue";

import Driver from "../components/driver/Driver.vue";
import RaceDetails from "../components/race-details/RaceDetails.vue";
import RaceControl from "../components/race-control/RaceControl.vue";
import RaceMap from "../components/race-map/RaceMap.vue";

import { dashboardData } from "../store/data.store";
import { sortPos } from "../utils/position.utils";
import RaceRadio from "../components/race-radios/RaceRadio.vue";

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

const driversSorted = computed(() => {
  return drivers.sort(sortPos);
});
</script>

<template>
  <div class="flex flex-col flex-1 bg-base-300">
    <RaceDetails :session="session" :trackStatus="trackStatus" :weather="weather" :extrapolatedClock="extrapolatedClock"
      :laps="lapCount" />

    <div class="grid overflow-auto">
      <h3 class="font-bold text-lg bg-base-100 p-2 w-full">Live Timming</h3>
      <div class="col-span-2 overflow-auto">
        <Driver v-for="driver in driversSorted" :driver="driver" :position="driver.position" />
      </div>

      <div class="grid grid-rows-1 grid-cols-6 overflow-auto" :style="{ maxHeight: '81vh' }">
        <div class="flex flex-col col-span-2">
          <h3 class="sticky top-0 font-bold text-lg bg-base-100 bg-fixed p-2">
            Race Control
          </h3>
          <div class="overflow-auto h-full">
              <RaceControl :messages="raceControlMessages" />
          </div>
          </div>

        <div class="flex flex-col col-span-3">
            <h3 class="sticky top-0 font-bold text-lg bg-base-100 bg-fixed p-2">
              Race Map
            </h3>

            <div class="overflow-auto h-full">
              <RaceMap :circuit="session.circuitKey" :trackStatus="trackStatus" :windDirection="weather.wind_direction"
                :position-batches="positionBatches" />
            </div>
          </div>

        <div class="flex flex-col ">
          <h3 class="sticky top-0 font-bold text-lg bg-base-100 bg-fixed p-2 z-[1]">
            Race Radios
          </h3>

          <div class="overflow-auto h-full">
            <RaceRadio :team-radios="teamRadios" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
