import { useEffect, useRef, useState } from "react";
import moment from "moment";

import "./App.css";
import { F1 } from "./models/f1.model";

function App() {
  const [connected, setConnected] = useState(false);
  const [liveState, setLiveState] = useState<F1>({} as F1);
  const [updated, setUpdated] = useState(new Date());
  const [delayMs, setDelayMs] = useState(0);
  const [delayTarget, setDelayTarget] = useState(0);
  const [blocking, setBlocking] = useState(false);
  const [triggerConnection, setTriggerConnection] = useState(0);
  const [triggerTick, setTriggerTick] = useState(0);

  const socket = useRef<WebSocket>();
  const retry = useRef<number>();

  const sortPosition = (a: any, b: any) => {
    const [, aLine] = a;
    const [, bLine] = b;
    const aPos = Number(aLine.Position);
    const bPos = Number(bLine.Position);
    return aPos - bPos;
  };

  const initWebsocket = (handleMessage: any) => {
    if (retry.current) {
      clearTimeout(retry.current);
      retry.current = undefined;
    }

    // const wsUrl =
    //   `${window.location.protocol.replace("http", "ws")}//` +
    //   window.location.hostname +
    //   (window.location.port ? `:${window.location.port}` : "") +
    //   "/ws";
    const wsUrl =
      "wss://3001-victorgomez09-f1tracker-zuw71jj3ixf.ws-eu108.gitpod.io";
    // const wsUrl = "ws://localhost:3001";

    const ws = new WebSocket(wsUrl);

    ws.addEventListener("open", () => {
      setConnected(true);
    });

    ws.addEventListener("close", () => {
      setConnected(false);
      setBlocking((isBlocking) => {
        if (!retry.current && !isBlocking)
          retry.current = window.setTimeout(() => {
            initWebsocket(handleMessage);
          }, 1000);

        return isBlocking;
      });
    });

    ws.addEventListener("error", () => {
      ws.close();
    });

    ws.addEventListener("message", ({ data }) => {
      setTimeout(() => {
        // console.log("mesasge", data);
        // const d = JSON.parse(data);
        // setLiveState(d);
        // setUpdated(new Date());
        handleMessage(data);
      }, delayMs);
    });

    socket.current = ws;
  };

  useEffect(() => {
    if ("serviceWorker" in navigator) {
      navigator.serviceWorker.register("/worker.js");
    }
  }, []);

  useEffect(() => {
    setLiveState({} as F1);
    setBlocking(false);
    initWebsocket((data: any) => {
      try {
        const d = JSON.parse(data);
        setLiveState(d);
        setUpdated(new Date());
      } catch (e) {
        console.error(`could not process message: ${e}`);
      }
    });
  }, [triggerConnection]);

  useEffect(() => {
    if (blocking) {
      socket.current?.close();
      setTimeout(() => {
        console.log("trigger connection");
        setTriggerConnection((n) => n + 1);
      }, 100);
    }
  }, [blocking]);

  useEffect(() => {
    let interval: number;
    if (Date.now() < delayTarget) {
      interval = setInterval(() => {
        setTriggerTick((n) => n + 1);
        if (Date.now() >= delayTarget) clearInterval(interval);
      }, 250);
    }
  }, [delayTarget]);

  // RENDER
  if (!connected)
    return (
      <>
        {(document.title = "No connection")}
        <main>
          <div
            style={{
              width: "100vw",
              height: "100vh",
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <p style={{ marginBottom: "var(--space-4)" }}>
              <strong>NO CONNECTION</strong>
            </p>
            <button onClick={() => window.location.reload()}>RELOAD</button>
          </div>
        </main>
      </>
    );

  const {
    extrapolatedClock,
    session,
    trackStatus,
    lapCount,
    weather,
    positionBatches,
  } = liveState;
  if (session) {
    const extrapolatedTimeRemaining =
      extrapolatedClock.utc && extrapolatedClock.remaining
        ? extrapolatedClock.extrapolating
          ? moment
              .utc(
                Math.max(
                  moment
                    .duration(extrapolatedClock.remaining)
                    .subtract(
                      moment.utc().diff(moment.utc(extrapolatedClock.utc))
                    )
                    .asMilliseconds() + delayMs,
                  0
                )
              )
              .format("HH:mm:ss")
          : extrapolatedClock.remaining
        : undefined;

    return (
      <div className="bg-base-300 w-full h-full">
        {/* {(document.title = `${session.circuitName}: ${session.typeName}`)} */}

        <div className="flex items-center justify-between gap-6 overflow-x-auto">
          <div className="flex flex-start gap-4">
            <p>
              <strong>{session.officialName}</strong>, {session.circuitName},{" "}
              {session.countryName}
            </p>

            <p>Session: {session.typeName}</p>

            {trackStatus.message && <p>Status: {trackStatus.message}</p>}

            {lapCount && (
              <p>
                Lap: {lapCount.current}/{lapCount.total}
              </p>
            )}

            {extrapolatedClock && <p>Remaining: {extrapolatedTimeRemaining}</p>}
          </div>

          <div className="flex items-center">
            <p>
              Data updated: {moment.utc(updated).format("HH:mm:ss.SSS")} UTC
            </p>
            <p style={{ color: "limegreen", marginRight: "var(--space-4)" }}>
              CONNECTED
            </p>
            <form
              onSubmit={(e) => {
                e.preventDefault();
                const form = new FormData(e.target);
                const delayMsValue = Number(form.get("delayMs"));
                setBlocking(true);
                setDelayMs(delayMsValue);
                setDelayTarget(Date.now() + delayMsValue);
              }}
              style={{ display: "flex", alignItems: "center" }}
            >
              <p style={{ marginRight: "var(--space-2)" }}>Delay</p>
              <input
                type="number"
                name="delayMs"
                defaultValue={delayMs}
                style={{ width: "75px", marginRight: "var(--space-2)" }}
              />
              <p style={{ marginRight: "var(--space-4)" }}>ms</p>
            </form>
          </div>
        </div>

        {weather && (
          <div className="flex items-center gap-4 overflow-x-auto">
            <p>
              <strong>WEATHER</strong>
            </p>

            <p>Air temp: {weather.air_temp}ºC</p>
            <p>Track temp: {weather.track_temp}</p>
            <p>Rainfall: {weather.rainfall}</p>
            <p>Wind dir: {weather.wind_direction}</p>
            <p>Wind speed: {weather.wind_speed}</p>
            <p>Humidity: {weather.humidity}%</p>
            <p>Pressure: {weather.pressure} mb</p>
          </div>
        )}

        <div className="divider divider-horizontal m-0"></div>

        <div>
          <div className="bg-base-100">
            <h2 className="font-semibold text-left">LIVE TIMING DATA</h2>
          </div>

          <div>
            {(() => {
              const lines = Object.entries(positionBatches).sort(sortPosition);

              return (
                <div>
                  <div
                    className="grid items-center gap-6 overflow-x-auto"
                    style={{
                      gridTemplateColumns:
                        "18px 42px 60px 60px 18px 74px 74px 44px 38px auto",
                    }}
                  >
                    <p>POS</p>
                    <p className="text-right">DRIVER</p>
                    <p>GEAR/RPM</p>
                    <p>SPD/PDL</p>
                    <p>DRS</p>
                    <p>TIME</p>
                    <p>GAP</p>
                    <p>TYRE</p>
                    <p>INFO</p>
                    <p>SECTORS</p>
                  </div>

                  {lines.map((line, index) => {
                    return (
                      
                    )
                  })}
                </div>
              );
            })()}
          </div>
        </div>
      </div>
    );
  }
}

export default App;
