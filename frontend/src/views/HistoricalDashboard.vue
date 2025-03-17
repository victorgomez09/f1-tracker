<script setup lang="ts">
import RaceControl from '@/components/race-control/RaceControl.vue';
import RaceDetails from '@/components/race-details/RaceDetails.vue';
import RaceMap from '@/components/race-map/RaceMap.vue';
import TelemetryTable from '@/components/telemetry/TelemetryTable.vue';
import { useDriverStore, useEventStore } from '@/store/data.store';
import { generateWsUrl, initWs } from '@/utils/ws.utils';
import { onMounted } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const raceName = route.params.eventName

onMounted(() => {
  generateWsUrl(raceName)
  initWs()
});
</script>

<template>

  <div class="flex flex-col gap-1.5 w-full"
    v-if="useDriverStore.drivers.length > 0 && Object.keys(useEventStore.event).length > 0">
    <div class="bg-base-100 rounded-md">
      <RaceDetails :isReplay="true"></RaceDetails>
    </div>

    <div class="row-span-8 grid grid-cols-8 gap-1.5">
      <div class="bg-base-100 w-full col-span-6 rounded-md">
        <TelemetryTable></TelemetryTable>
      </div>

      <div class="bg-base-100 p-1 w-full col-span-2 rounded-md overflow-auto" :style="{height: '34em'}">
        <RaceControl></RaceControl>
      </div>
    </div>

    <div class="grid">
      <div class="bg-base-100 w-full col-span-6 rounded-md">
        <RaceMap></RaceMap>
      </div>

      <!-- <div class="bg-base-100 p-1 w-full col-span-2 rounded-md overflow-auto">
        <RaceControl></RaceControl>
      </div> -->
    </div>
  </div>
  <div v-else class="flex items-center justify-center w-full h-full">
    <span class="loading loading-spinner loading-md"></span>
  </div>
</template>