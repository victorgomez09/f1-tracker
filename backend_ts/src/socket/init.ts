import { WebSocket } from "ws";
import { F1State } from "../models/f1.model";
import { negotiate } from "./negotiate";
import { config } from "../config";
import { updateState } from "./state";
import { subscribeRequest } from "./subscrition";
import { translate } from "./translators";
import { retrySetup } from "./retry";

export let f1_ws: WebSocket | null;
export let state: F1State = {};
let active: boolean = false;

export const setupF1 = async (wss: WebSocket) => {
  if (f1_ws) return;
  state = {};

  console.log("F1: setting up socket");

  const hub = encodeURIComponent(JSON.stringify([{ name: "Streaming" }]));
  const { body, cookie } = await negotiate(hub);

  const token = encodeURIComponent(body.ConnectionToken);
  const url = `${config.f1BaseUrl}/connect?clientProtocol=1.5&transport=webSockets&connectionToken=${token}&connectionData=${hub}`;

  console.log("F1: connecting!");

  f1_ws = new WebSocket(url, {
    headers: {
      "User-Agent": "BestHTTP",
      "Accept-Encoding": "gzip,identity",
      Cookie: cookie,
    },
  });

  f1_ws.onmessage = (rawData) => {
    if (typeof rawData.data !== "string") return;
    const data = JSON.parse(rawData.data);

    state = updateState(state, data);
    wss.emit("f1-socket", JSON.stringify(translate(state)));
  };

  f1_ws.onopen = () => f1_ws?.send(subscribeRequest());

  f1_ws.onerror = () => {
    console.log("F1: got error");
    console.log("F1: closing socket");
    f1_ws?.close();
  };

  f1_ws.onclose = () => {
    console.log("F1: got close");
    console.log("F1: killing socket");
    f1_ws = null;

    retrySetup(active, wss);
  };

  if (!f1_ws || f1_ws.readyState === 3) {
    console.log("F1: failed to setup socket");
    retrySetup(active, wss);
  }
};
