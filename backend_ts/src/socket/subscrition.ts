export const subscribeRequest = (): string => {
  console.log("F1: sent subscribe request");
  return JSON.stringify({
    H: "Streaming",
    M: "Subscribe",
    A: [
      [
        "Heartbeat",
        "CarData.z",
        "Position.z",
        "ExtrapolatedClock",
        "TimingStats",
        "TimingAppData",
        "WeatherData",
        "TrackStatus",
        "DriverList",
        "RaceControlMessages",
        "SessionInfo",
        "SessionData",
        "LapCount",
        "TimingData",
        "TeamRadio",
      ],
    ],
    I: 1,
  });
};
