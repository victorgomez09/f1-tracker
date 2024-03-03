import { F1ExtrapolatedClock } from "./clock.model";
import { F1Driver } from "./driver.model";
import { F1Session } from "./session.model";
import { F1TrackStatus } from "./track-status.model";
import { F1WeatherData } from "./weather";

export type F1 = {
  session: F1Session;
  trackStatus: F1TrackStatus;
  extrapolatedClock: F1ExtrapolatedClock;
  weather: F1WeatherData;
  drivers: F1Driver[];
};
