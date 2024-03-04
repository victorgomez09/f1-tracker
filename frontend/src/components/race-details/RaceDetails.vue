<script setup lang="ts">
import { PropType, computed, ref, watchEffect } from "vue";
import moment from "moment";

import { F1Session } from "../../models/session.model";
import { F1ExtrapolatedClock } from "../../models/clock.model";
import { getTrackStatusMessage } from "../../utils/track.util";
import { F1TrackStatus } from "../../models/track-status.model";
import { F1WeatherData } from "../../models/weather";
import { getWindDirection } from "../../utils/wind.util";

const props = defineProps({
  session: Object as PropType<F1Session>,
  extrapolatedClock: Object as PropType<F1ExtrapolatedClock>,
  trackStatus: Object as PropType<F1TrackStatus>,
  weather: Object as PropType<F1WeatherData>,
});

const countryFlag = ref();
watchEffect(async () => {
  countryFlag.value = (
    await import(
      /* @vite-ignore */ `../../assets/country-flags/${props.session?.countryCode.toLowerCase()}.svg`
    )
  ).default;
});
const timeRemaining = computed(() => {
  return !!props.extrapolatedClock && !!props.extrapolatedClock?.remaining
    ? props.extrapolatedClock.extrapolating
      ? moment
          .utc(
            moment
              .duration(props.extrapolatedClock.remaining)
              .subtract(
                moment.utc().diff(moment.utc(props.extrapolatedClock.utc))
              )
              // .asMilliseconds() + (delay ? delay * 1000 : 0),
              .asMilliseconds()
          )
          .format("HH:mm:ss")
      : props.extrapolatedClock.remaining
    : undefined;
});
const trackStatusInfo = computed(() => {
  return getTrackStatusMessage(props.trackStatus?.status);
});
const windDirection = computed(() => {
  return getWindDirection(props.weather?.wind_direction!);
});
</script>

<template>
  <div class="flex flex-wrap items-center gap-2 border-b border-b-primary p-4">
    <div class="flex flex-col gap-2">
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-2">
          <img
            class="relative overflow-hidden rounded h-12 w-16"
            :src="countryFlag"
            alt="Country flag"
          />

          <div class="flex flex-col">
            <span class="font-semibold"
              >{{ props.session?.name }}: {{ props.session?.type }}</span
            >
            <div class="flex items-center gap-2">
              <span class="text-3xl font-bold">{{ timeRemaining }}</span>

              <span class="badge badge-lg" :class="trackStatusInfo?.color">{{
                trackStatusInfo?.message
              }}</span>
            </div>
          </div>

          <div class="flex gap-4">
            <div class="flex flex-col">
              <span class="font-semibold">Wind speed</span>
              <span class="text-xl font-bold">
                {{ props.weather?.wind_speed }} km/h
              </span>
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Wind direction</span>
              <span class="flex items-center text-xl font-bold">
                {{ windDirection }}
                <img
                  src="../../assets/icons/arrow-up.svg"
                  alt="arrow"
                  :style="[
                    { rotate: `${props.weather?.wind_direction}deg` },
                    { transition: '1s linear' },
                  ]"
                />
              </span>
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Air temp</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.air_temp }}ºC</span
              >
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Track temp</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.track_temp }}ºC</span
              >
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Wind</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.track_temp }}ºC</span
              >
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Humidity</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.humidity }}%</span
              >
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Pressure</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.pressure }} mb</span
              >
            </div>

            <div class="flex flex-col">
              <span class="font-semibold">Rainfall</span>
              <span class="text-xl font-bold"
                >{{ props.weather?.rainfall }}%</span
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
