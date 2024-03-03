<script setup lang="ts">
import { computed, isProxy, ref, toRaw, watchEffect } from "vue";
import moment from "moment";

import Driver from "../components/Driver.vue";
import { dashboardData } from "../store/data.store";
import { getTrackStatusMessage } from "../utils/track.util";
import { getWindDirection } from "../utils/wind.util";
import { sortPos } from "../utils/position.utils";

const data = isProxy(dashboardData)
  ? toRaw(dashboardData.data)
  : dashboardData.data;
console.log("data from dash", data);
const { session } = data;
const { trackStatus } = data;
const { weather } = data;
const { extrapolatedClock } = data;
const { drivers } = data;

const countryFlag = ref();
watchEffect(async () => {
  countryFlag.value = (
    await import(
      /* @vite-ignore */ `../assets/country-flags/${session.countryCode.toLowerCase()}.svg`
    )
  ).default;
});
const timeRemaining = computed(() => {
  return !!extrapolatedClock && !!extrapolatedClock.remaining
    ? extrapolatedClock.extrapolating
      ? moment
          .utc(
            moment
              .duration(extrapolatedClock.remaining)
              .subtract(moment.utc().diff(moment.utc(extrapolatedClock.utc)))
              // .asMilliseconds() + (delay ? delay * 1000 : 0),
              .asMilliseconds()
          )
          .format("HH:mm:ss")
      : extrapolatedClock.remaining
    : undefined;
});
const trackStatusInfo = computed(() => {
  return getTrackStatusMessage(trackStatus.status);
});
const windDirection = computed(() => {
  return getWindDirection(weather.wind_direction);
});
const driversSorted = computed(() => {
  return drivers.sort(sortPos);
});
</script>

<template>
  <div class="flex flex-col flex-1 w-full h-full">
    <!-- Header -->
    <div
      class="flex flex-1 flex-wrap items-center gap-2 border-b border-b-primary p-4"
    >
      <!-- Session info -->
      <div class="flex flex-col gap-2">
        <!-- <h5 class="font-semibold">{{ session.officialName }}</h5> -->
        <div class="flex items-center gap-2">
          <div class="flex items-center gap-2">
            <img
              class="relative overflow-hidden rounded h-12 w-16"
              :src="countryFlag"
              alt="Country flag"
            />

            <div class="flex flex-col">
              <span class="font-semibold"
                >{{ session.name }}: {{ session.type }}</span
              >
              <div class="flex items-center gap-2">
                <span class="text-3xl font-bold">{{ timeRemaining }}</span>

                <span class="badge badge-lg" :class="trackStatusInfo?.color">{{
                  trackStatusInfo?.message
                }}</span>
              </div>
            </div>

            <!-- Weather -->
            <div class="flex gap-4">
              <div class="flex flex-col">
                <span class="font-semibold">Wind speed</span>
                <span class="text-xl font-bold">
                  {{ weather.wind_speed }} km/h
                </span>
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Wind direction</span>
                <span class="flex items-center text-xl font-bold">
                  {{ windDirection }}
                  <img
                    src="../assets/icons/arrow-up.svg"
                    alt="arrow"
                    :style="[
                      { rotate: `${weather.wind_direction}deg` },
                      { transition: '1s linear' },
                    ]"
                  />
                </span>
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Air temp</span>
                <span class="text-xl font-bold">{{ weather.air_temp }}ºC</span>
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Track temp</span>
                <span class="text-xl font-bold"
                  >{{ weather.track_temp }}ºC</span
                >
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Wind</span>
                <span class="text-xl font-bold"
                  >{{ weather.track_temp }}ºC</span
                >
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Humidity</span>
                <span class="text-xl font-bold">{{ weather.humidity }}%</span>
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Pressure</span>
                <span class="text-xl font-bold">{{ weather.pressure }} mb</span>
              </div>

              <div class="flex flex-col">
                <span class="font-semibold">Rainfall</span>
                <span class="text-xl font-bold">{{ weather.rainfall }}%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div>
      <Driver
        v-for="driver in driversSorted"
        :driver="driver"
        :position="driver.position"
      />
    </div>
  </div>
</template>
