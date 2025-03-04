<script setup lang="ts">
import { useDriverStore } from '@/store/data.store';
import { parseData } from '@/utils/parse-data.utils';
import { computed, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import Telemetry from './Telemetry.vue'
import TelemetryTable from '@/components/telemetry/TelemetryTable.vue';
import RaceDetails from '@/components/race-details/RaceDetails.vue';
import { generateWsUrl, getWs, initWs } from '@/utils/ws.utils';

const route = useRoute()
const raceName = route.params.eventName
// const socket = ref();
// const retry = ref();

// const delayMs = ref(0);
// const connected = ref(false);
// const blocking = ref(false);

// const dataUpdated = ref(false);

  //   const wsUrl = "ws://localhost:3001";
// const initWebsocket = (handleMessage: any) => {
//   if (retry.value) {
//     clearTimeout(retry.value);
//     retry.value = undefined;
//   }

//   const ws = new WebSocket(wsUrl);

//   ws.addEventListener("open", () => {
//     connected.value = true;
//   });

//   ws.addEventListener("close", () => {
//     connected.value = false;
//     blocking.value = true;
//     () => {
//       if (!retry.value && !blocking.value)
//         retry.value = window.setTimeout(() => {
//           initWebsocket(handleMessage);
//         }, 1000);
//     };
//   });

//   ws.addEventListener("error", () => {
//     ws.close();
//   });

//   ws.addEventListener("message", ({ data }) => {
//     setTimeout(() => {
//       handleMessage(data);
//       console.log("message", data);
//     }, delayMs.value);
//   });

//   socket.value = ws;
// };

onMounted(() => {
  // setInterval(() => {
  // const ws = new WebSocket(
  //   wsUrl
  // );

  // ws.addEventListener("open", () => {
  //   console.log("open");
  //   connected.value = true;
  // });

  // ws.onmessage = (data) => {
  //   try {
  //       // console.log("JSON.parse(data.data)",JSON.parse(data.data))
  //     parseData(JSON.parse(data.data));
  //     // liveState.value = d;
  //     // updated.value = new Date();
  //     // dashboardData.value = d;
  //     dataUpdated.value = true;
  //   } catch (e) {
  //     console.error(`could not process message: ${e}`);
  //   }
  // };
  generateWsUrl(raceName)
  initWs()
});
</script>

<template>
  <div class="flex flex-col gap-1 w-full h-full">
    <div class="bg-secondary p-2 rounded-md">
      <RaceDetails :isReplay="true"></RaceDetails>
    </div>

    <div class="bg-secondary p-2 rounded-md">
      <TelemetryTable></TelemetryTable>
    </div>
  </div>
</template>