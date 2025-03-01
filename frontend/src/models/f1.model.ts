import { F1CarData } from "./car.model";
import { F1ExtrapolatedClock } from "./clock.model";
import { Driver } from "./driver.model";
import { F1Laps } from "./laps.model";
import { DriverPositionBatch } from "./position.model";
import { F1RaceControlMessage } from "./race-control.model";
import { TeamRadioType } from "./radio.model";
import { F1Session } from "./session.model";
import { F1TrackStatus } from "./track-status.model";
import { F1WeatherData } from "./weather";

export type F1 = {
  session: F1Session;
  trackStatus: F1TrackStatus;
  extrapolatedClock: F1ExtrapolatedClock;
  weather: F1WeatherData;
  drivers: Driver[];
  raceControlMessages: F1RaceControlMessage[];
  positionBatches: DriverPositionBatch[];
  teamRadios: TeamRadioType[];
  lapCount: F1Laps;
  carData: F1CarData;
};
