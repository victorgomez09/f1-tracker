<script setup lang="ts">
import { onMounted, ref } from "vue";

import TelemetryTable from "@/components/telemetry/TelemetryTable.vue";
import { parseData } from "@/utils/parse-data.utils";

// Connect to websocket
const socket = ref();
const retry = ref();

const delayMs = ref(0);
const connected = ref(false);
const blocking = ref(false);

const dataUpdated = ref(false);

const initWebsocket = (handleMessage: any) => {
  if (retry.value) {
    clearTimeout(retry.value);
    retry.value = undefined;
  }

  // const wsUrl =
  //   `${window.location.protocol.replace("http", "ws")}//` +
  //   window.location.hostname +
  //   (window.location.port ? `:${window.location.port}` : "") +
  //   "/ws";

  // const wsUrl =
  //   "wss://victorgomez09-f1tracker-rkauu33dlup.ws-eu118.gitpod.io/wss";
  const wsUrl = "ws://localhost:3001";

  const ws = new WebSocket(wsUrl);

  ws.addEventListener("open", () => {
    connected.value = true;
  });

  ws.addEventListener("close", () => {
    connected.value = false;
    blocking.value = true;
    () => {
      if (!retry.value && !blocking.value)
        retry.value = window.setTimeout(() => {
          initWebsocket(handleMessage);
        }, 1000);
    };
  });

  ws.addEventListener("error", () => {
    ws.close();
  });

  ws.addEventListener("message", ({ data }) => {
    setTimeout(() => {
      handleMessage(data);
      console.log("message", data);
    }, delayMs.value);
  });

  socket.value = ws;
};

onMounted(() => {
  // setInterval(() => {
  //   initWebsocket((data: any) => {
  //     try {
  //       const d = JSON.parse(data);
  //       liveState.value = d;
  //       updated.value = new Date();
  //       dashboardData.data = d;
  //       dataUpdated.value = true;
  //     } catch (e) {
  //       console.error(`could not process message: ${e}`);
  //     }
  //   });
  // }, 100)
  // const ws = new WebSocket(
  //   "wss://3000-victorgomez09-f1tracker-rkauu33dlup.ws-eu118.gitpod.io/ws"
  // );
  const ws = new WebSocket(
    "ws://localhost:3000/ws"
  );

  ws.addEventListener("open", () => {
    console.log("open");
    connected.value = true;
  });

  // ws.addEventListener("close", () => {
  //   connected.value = false;
  //   blocking.value = true;
  //   () => {
  //     if (!retry.value && !blocking.value)
  //       retry.value = window.setTimeout(() => {
  //         // initWebsocket(handleMessage);
  //       }, 1000);
  //   };
  // });

  // ws.addEventListener("error", () => {
  //   ws.close();
  // });

  // ws.addEventListener("message", ({ data }) => {
  //   setTimeout(() => {
  //     // handleMessage(data);
  //     console.log("message", data);
  //   }, delayMs.value);
  // });

  ws.onmessage = (data) => {
    try {
      // const d: ServerResponse = JSON.parse(data.data);
      // console.log("d", d);
      // console.log("onmessage", data.data);
      parseData(JSON.parse(data.data));
      // liveState.value = d;
      // updated.value = new Date();
      // dashboardData.value = d;
      dataUpdated.value = true;
    } catch (e) {
      console.error(`could not process message: ${e}`);
    }
  };
});

// const dataProxy = isProxy(dashboardData)
//   ? toRaw(dashboardData)
//   : dashboardData.value;
// console.log("dataProxy", dataProxy);
// const data =
//   dataProxy == null ||
//     dataProxy === undefined ||
//     Object.keys(dataProxy).length === 0
//     ? JSON.parse(mockData)
//     : dataProxy;
// const { session } = data;
// const { trackStatus } = data;
// const { weather } = data;
// const { extrapolatedClock } = data;
// const { drivers } = data;
// const { raceControlMessages } = data;
// const { positionBatches } = data;
// const { teamRadios } = data;
// const { lapCount } = data;
// const { carData } = data;

// const liveTimingViewOption = ref<boolean>(false);
// const checkChange = (element: any) => {
//   viewMode.value = element.target.checked ? "TELEMETRY" : "PRETTY";
// };

// const driversSorted = computed(() => {
//   return drivers.sort(sortPos);
// });

// const radiosSorted = computed(() => {
//   return teamRadios.sort(sortUtc);
// });
</script>

<template>
  <!--<div v-if="!liveTimingViewOption">
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
          <div class="form-control">
            <label class="cursor-pointer label">
              <span class="label-text mr-2">
                {{ liveTimingViewOption ? "TELEMETRY" : "PRETTY" }}
              </span>
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
                  <input
                    type="checkbox"
                    class="toggle toggle-primary"
                    checked
                  />
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
  </div>

  <div v-else class="flex flex-col flex-1 gap-2 w-full h-full">
    <RaceDetails
      :session="session"
      :trackStatus="trackStatus"
      :weather="weather"
      :extrapolatedClock="extrapolatedClock"
      :laps="lapCount"
    />

    <div>
      <TelemetryTable :data="driversSorted" />
    </div>
  </div> -->
  <div>
    <TelemetryTable />
  </div>
</template>
