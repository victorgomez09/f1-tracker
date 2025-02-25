<script setup lang="ts">
import { onMounted, ref, withScopeId } from "vue";
import { dashboardData } from "./store/data.store";

const socket = ref();
const retry = ref();

const delayMs = ref(0);
const connected = ref(false);
const blocking = ref(false);

const liveState = ref({});
const dataUpdated = ref(false);
const updated = ref(new Date());

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

  const wsUrl =
    "wss://3000-victorgomez09-f1tracker-a7zkp3fi4k8.ws-eu118.gitpod.io/wss";
  // const wsUrl = "ws://localhost:3001";

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
  const ws = new WebSocket(
    "wss://3000-victorgomez09-f1tracker-a7zkp3fi4k8.ws-eu118.gitpod.io/ws"
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
    console.log("onmessage", data.data);
    try {
      const d = JSON.parse(data.data);
      console.log("d", d);
      liveState.value = d;
      updated.value = new Date();
      dashboardData.data = d;
      dataUpdated.value = true;
    } catch (e) {
      console.error(`could not process message: ${e}`);
    }
  };
});
</script>

<template>
  <div class="flex flex-1 h-full w-full">
    <!--<router-view v-if="connected && dataUpdated"></router-view>-->
    <router-view></router-view>
  </div>
</template>
