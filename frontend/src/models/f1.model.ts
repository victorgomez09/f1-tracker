import { F1ExtrapolatedClock } from "./clock.model";
import { F1Driver } from "./driver.model";
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
  drivers: F1Driver[];
  raceControlMessages: F1RaceControlMessage[];
  positionBatches: DriverPositionBatch[];
  teamRadios: TeamRadioType[];
};
