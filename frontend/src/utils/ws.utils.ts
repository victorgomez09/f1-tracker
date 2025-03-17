import { ref } from "vue";
import { parseData } from "./parse-data.utils";

let wsUrl = '';

let ws = ref<WebSocket | null>(null)
const connected = ref<boolean>(false)

export const generateWsUrl = (ops?: string | string[]) => {
  // if (ops) wsUrl = `wss://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/historical/${ops}`
  // else wsUrl = `wss://stunning-system-j4wxj4p5v4j3555p-3000.app.github.dev/ws`
  if (ops) wsUrl = `ws://localhost:3000/historical/${ops}`
  else wsUrl = `ws://localhost:3000/ws`
}

export const initWs = () => {
  ws.value = new WebSocket(
    wsUrl
  );

  ws.value.addEventListener("open", () => {
    console.log("open");
    connected.value = true;
  });

  ws.value.onmessage = (data) => {
    try {
      // console.log("JSON.parse(data.data)",JSON.parse(data.data))
      console.log("data", JSON.parse(data.data))
      parseData(JSON.parse(data.data));
      // liveState.value = d;
      // updated.value = new Date();
      // dashboardData.value = d;
      //   dataUpdated.value = true;
    } catch (e) {
      console.error(`could not process message: ${e}`);
    }
  };
}

export const getWs = () => {
  if (ws.value !== null) return ws.value
  return null
}
