<script setup lang="ts">
import { PropType, computed, ref, watchEffect } from "vue";
import moment from "moment";

import { F1Session } from "../../models/session.model";
import { F1ExtrapolatedClock } from "../../models/clock.model";
import { getTrackStatusMessage } from "../../utils/track.util";
import { F1TrackStatus } from "../../models/track-status.model";
import { F1WeatherData } from "../../models/weather";
import { getWindDirection } from "../../utils/wind.util";
import { F1Laps } from "../../models/laps.model";

const props = defineProps({
  session: Object as PropType<F1Session>,
  extrapolatedClock: Object as PropType<F1ExtrapolatedClock>,
  trackStatus: Object as PropType<F1TrackStatus>,
  weather: Object as PropType<F1WeatherData>,
  laps: Object as PropType<F1Laps>,
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
  <div class="w-full">
    <div class="flex items-center self-start gap-2 p-2">
      <img
        class="relative overflow-hidden rounded h-6 w-12"
        :src="countryFlag"
        alt="Country flag"
      />

      <div class="flex items-center gap-6 w-full">
        <p>
          <span class="font-semibold">{{ props.session?.officialName }}</span>
          , {{ session?.location }}, {{ session?.countryName }}
        </p>

        <p><span class="font-semibold">Session:</span> {{ session?.type }}</p>
        <p v-if="!laps">
          <span class="font-semibold">Remaining:</span> {{ timeRemaining }}
        </p>
        <p v-else>
          <span class="font-semibold">Laps:</span> {{ laps.current }} /
          {{ laps.total }}
        </p>

        <!-- <span class="badge badge-lg" :class="trackStatusInfo?.color">
          {{ trackStatusInfo?.message }}
        </span> -->

        <div class="flex items-center gap-2 ml-auto">
          <div class="dropdown dropdown-end">
            <div tabindex="0" role="button" class="btn m-1">
              Theme
              <svg
                width="12px"
                height="12px"
                class="h-2 w-2 fill-current opacity-60 inline-block"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 2048 2048"
              >
                <path
                  d="M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z"
                ></path>
              </svg>
            </div>
            <ul
              tabindex="0"
              class="dropdown-content z-[2] p-2 shadow-2xl bg-base-300 rounded-box w-52"
            >
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  class="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Default"
                  value="default"
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  class="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Retro"
                  value="retro"
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  class="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Cyberpunk"
                  value="cyberpunk"
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  class="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Valentine"
                  value="valentine"
                />
              </li>
              <li>
                <input
                  type="radio"
                  name="theme-dropdown"
                  class="theme-controller btn btn-sm btn-block btn-ghost justify-start"
                  aria-label="Aqua"
                  value="aqua"
                />
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <div class="divider m-0 h-0"></div>

    <div class="flex items-center gap-4 p-2 text-sm">
      <h2 class="font-semibold">WEATHER</h2>

      <p>Wind speed: {{ props.weather?.wind_speed }} km/h</p>

      <p class="flex items-center">
        Wind dir: {{ windDirection }}
        <img
          class="text-base-content"
          src="../../assets/icons/arrow-up.svg"
          alt="arrow"
          :style="[
            { rotate: `${props.weather?.wind_direction}deg` },
            { transition: '1s linear' },
          ]"
        />
      </p>

      <p>Air temp: {{ props.weather?.air_temp }}ºC</p>
      <p>Track temp: {{ props.weather?.track_temp }}ºC</p>
      <p>Wind: {{ props.weather?.track_temp }}ºC</p>
      <p>Humidity: {{ props.weather?.humidity }}%</p>
      <p>Pressure: {{ props.weather?.pressure }} mb</p>
      <p>Rainfall: {{ props.weather?.rainfall }}%</p>
    </div>

    <div class="divider m-0 h-0"></div>
  </div>

  <!-- <div class="flex items-center gap-2 border-b border-b-primary p-4 w-full">
    <div class="flex flex-col gap-2">
      <div class="flex flex-col md:flex-row items-center gap-2">
        <div class="flex items-center self-start gap-2">
          <img class="relative overflow-hidden rounded h-8 lg:h-12 w-12 lg:w-16" :src="countryFlag" alt="Country flag" />
  
          <div class="flex flex-col">
            <span class="font-semibold">{{ props.session?.name }}: {{ props.session?.type }}</span>
            <div class="flex items-center gap-2">
              <span class="lg:text-3xl font-bold" v-if="!laps">{{ timeRemaining }}</span>
              <span class="lg:text-3xl font-bold" v-if="laps">
                {{ laps.current }} / {{ laps.total }}
              </span>
  
              <span class="badge badge-lg" :class="trackStatusInfo?.color">{{
                trackStatusInfo?.message
                }}</span>
            </div>
          </div>
        </div>

        <div class="divider lg:hidden"></div>

        <div class="flex items-center gap-4 lg:ml-4">
          <div class="flex flex-col">
            <span class="font-semibold">Wind speed</span>
            <span class="lg:text-xl font-bold">
              {{ props.weather?.wind_speed }} km/h
            </span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Wind dir</span>
            <span class="flex items-center lg:text-xl font-bold">
              {{ windDirection }}
              <img class="text-base-content" src="../../assets/icons/arrow-up.svg" alt="arrow" :style="[
                    { rotate: `${props.weather?.wind_direction}deg` },
                    { transition: '1s linear' },
                  ]" />
            </span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Air temp</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.air_temp }}ºC</span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Track temp</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.track_temp }}ºC</span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Wind</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.track_temp }}ºC</span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Humidity</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.humidity }}%</span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Pressure</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.pressure }} mb</span>
          </div>

          <div class="flex flex-col">
            <span class="font-semibold">Rainfall</span>
            <span class="lg:text-xl font-bold">{{ props.weather?.rainfall }}%</span>
          </div>
        </div>
      </div>
    </div>
  </div> -->
</template>
