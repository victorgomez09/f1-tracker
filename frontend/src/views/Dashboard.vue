<script setup lang="ts">
import { computed, isProxy, toRaw } from "vue";

import RaceDetails from "../components/race-details/RaceDetails.vue";
import Driver from "../components/driver/Driver.vue";
import RaceControl from "../components/race-control/RaceControl.vue";

import { dashboardData } from "../store/data.store";
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
const { raceControlMessages } = data;

const driversSorted = computed(() => {
  return drivers.sort(sortPos);
});
</script>

<template>
  <!-- <div class="flex flex-col flex-1 bg-base-300 min-h-screen">
    

    <div class="relative flex flex-grow">
      <main class="flex flex-1 overflow-scroll">
        <div class="grid grid-cols-3 w-full h-full">
          <div class="col-span-2 overflow-scroll h-dvh w-dvw">
            <h3 class="font-bold text-lg bg-base-100 p-2">Live Timming</h3>
            <Driver
              v-for="driver in driversSorted"
              :driver="driver"
              :position="driver.position"
            />
          </div>

          <div class="w-full overflow-auto">
            <h3 class="font-bold text-lg bg-base-100 p-2">Race Control</h3>
            <RaceControl :messages="raceControlMessages" />
          </div>
        </div>
      </main>
    </div>
  </div> -->

  <div class="flex flex-col flex-1 bg-base-300">
    <RaceDetails
      :session="session"
      :trackStatus="trackStatus"
      :weather="weather"
      :extrapolatedClock="extrapolatedClock"
    />

    <div class="grid grid-cols-3 overflow-auto">
      <div class="col-span-2 overflow-auto">
        <h3 class="font-bold text-lg bg-base-100 p-2">Live Timming</h3>
        <Driver
          v-for="driver in driversSorted"
          :driver="driver"
          :position="driver.position"
        />
      </div>

      <div class="overflow-auto">
        <h3 class="font-bold text-lg bg-base-100 p-2">Race Control</h3>
        <RaceControl :messages="raceControlMessages" />
      </div>
      <!-- <div
        class="p-4 border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700"
      >
        <div class="grid grid-cols-3 gap-4 mb-4">
          <div
            class="flex items-center justify-center h-24 rounded bg-gray-50 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center h-24 rounded bg-gray-50 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center h-24 rounded bg-gray-50 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
        </div>
        <div
          class="flex items-center justify-center h-48 mb-4 rounded bg-gray-50 dark:bg-gray-800"
        >
          <p class="text-2xl text-gray-400 dark:text-gray-500">
            <svg
              class="w-3.5 h-3.5"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 18 18"
            >
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 1v16M1 9h16"
              />
            </svg>
          </p>
        </div>
        <div class="grid grid-cols-2 gap-4 mb-4">
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
        </div>
        <div
          class="flex items-center justify-center h-48 mb-4 rounded bg-gray-50 dark:bg-gray-800"
        >
          <p class="text-2xl text-gray-400 dark:text-gray-500">
            <svg
              class="w-3.5 h-3.5"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 18 18"
            >
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 1v16M1 9h16"
              />
            </svg>
          </p>
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
          <div
            class="flex items-center justify-center rounded bg-gray-50 h-28 dark:bg-gray-800"
          >
            <p class="text-2xl text-gray-400 dark:text-gray-500">
              <svg
                class="w-3.5 h-3.5"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 18 18"
              >
                <path
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 1v16M1 9h16"
                />
              </svg>
            </p>
          </div>
        </div>
      </div> -->
    </div>
  </div>
</template>
