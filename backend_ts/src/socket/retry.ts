import { WebSocket } from "ws";

import { setupF1 } from "./init";
import { sleep } from "../utils/sleep.util";

export const retrySetup = async (active: boolean, wss: WebSocket) => {
  if (!active) return;
  console.log("F1: retrying to setup in 1s");

  await sleep(1000);
  setupF1(wss);
};
