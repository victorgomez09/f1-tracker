<script setup lang="ts">
import { dashboardData } from "../store/data.store";

const { session } = dashboardData.data;
const { trackStatus } = dashboardData.data;
const { weatherData } = dashboardData.data;
</script>

<template>
  <div class="flex flex-col flex-1 w-full h-full">
    <div
      class="flex flex-1 flex-wrap items-center gap-2 border border-b-primary p-2"
    >
      <!-- Session info -->
      <div class="flex items-center">
        <h5 class="font-semibold">{{ session.officialName }}</h5>
        <span>{{ session.circuitName }}, {{ session.countryName }}</span>
        <span>Session: {{ session.type }}</span>
        <span>Status: {{ trackStatus.message }}</span>
      </div>

      <!-- Weather -->
      <div class="flex items-center">
        <span>Air temp: {{ weatherData.airTemp }}</span>
        <span>Humidity: {{ weatherData.humidity }}</span>
        <span>Pressure: {{ weatherData.pressure }}</span>
        <span>Rainfall: {{ weatherData.rainfall }}</span>
        <span>Wind dir: {{ weatherData.windDirection }}</span>
        <span>Wind speed: {{ weatherData.windSpeed }}</span>
      </div>
    </div>
  </div>
</template>
