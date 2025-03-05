<script setup lang="ts">
import RaceDetails from '@/components/race-details/RaceDetails.vue';
import TelemetryTable from '@/components/telemetry/TelemetryTable.vue';
import RaceControl from '@/components/race-control/RaceControl.vue'
import { generateWsUrl, initWs } from '@/utils/ws.utils';
import { onMounted } from 'vue';
import { useRoute } from 'vue-router';

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
  <div class="flex flex-col gap-1.5 w-full">
    <div class="bg-base-100 p-2 rounded-md">
      <RaceDetails :isReplay="true"></RaceDetails>
    </div>

    <div class="grid grid-cols-8 gap-1.5 max-h-6/12">
      <div class="bg-base-100 w-full col-span-6 rounded-md">
        <TelemetryTable></TelemetryTable>
      </div>

      <div class="bg-base-100 p-1 w-full col-span-2 rounded-md overflow-auto">
        <RaceControl></RaceControl>
      </div>
    </div>
  </div>
</template>