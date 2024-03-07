import express from "express";
import { WebSocketServer, WebSocket } from "ws";

import { f1_ws, setupF1, state } from "./socket/init";
import { translate } from "./socket/translators";

const server = express();
const serverPort = process.env.SERVER_PORT || 3000;
const wsPort = process.env.WS_PORT || 3001;

server.listen(serverPort, async () => {
  console.log(`Server running on port. ${serverPort}`);

  const ws = new WebSocketServer({ port: Number(wsPort) });
  console.log(`Running ws on port. ${Number(wsPort)}`);

  ws.on("connection", async (ws: WebSocket) => {
    console.log("client connected");
    ws.on("f1-socket", (data) => {
      ws.send(data);
    });

    if (!f1_ws) {
      await setupF1(ws);
    } else {
      ws.send(JSON.stringify(translate(state)));
    }
  });
});
