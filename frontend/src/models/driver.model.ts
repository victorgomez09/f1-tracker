export type Driver = {
  Timestamp: Date;
  Position: number;
  Name: string;
  ShortName: string;
  Number: number;
  Team: string;
  HexColor: string;
  Color: { R: number; G: number; B: number; A: number };
  TimeDiffToFastest: number;
  TimeDiffToPositionAhead: number;
  GapToLeader: number;
  PreviousSegmentIndex: number;
  Segment: number[];
  Sector1: number;
  Sector1PersonalFastest: boolean;
  Sector1OverallFastest: boolean;
  Sector2: number;
  Sector2PersonalFastest: boolean;
  Sector2OverallFastest: boolean;
  Sector3: number;
  Sector3PersonalFastest: boolean;
  Sector3OverallFastest: boolean;
  LastLap: number;
  LastLapPersonalFastest: boolean;
  LastLapOverallFastest: boolean;
  FastestLap: number;
  OverallFastestLap: boolean;
  KnockedOutOfQualifying: boolean;
  ChequeredFlag: boolean;
  Tire: number;
  LapsOnTire: number;
  Lap: number;
  DRSOpen: boolean;
  Pitstops: number;
  PitStopTimes: [];
  Location: DriverLocation;
  SpeedTrap: number;
  SpeedTrapPersonalFastest: boolean;
  SpeedTrapOverallFastest: boolean;
};

export enum DriverLocation {
  NoLocation = 0,
  Pitlane = 1,
  PitOut = 2,
  OutLap = 3,
  OnTrack = 4,
  OutOfRace = 5,
  Stopped = 6,
}

export type PitStop = {
  Lap: number;
  PitlaneEntry: number;
  PitlaneExit: number;
  PitlaneTime: number;
};
