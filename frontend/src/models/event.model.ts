export type Event = {
  Timestamp: Date;
  Name: number;
  Type: number;
  Meeting: Meeting;
  Status: number;
  Heartbeat: boolean;
  CurrentLap: number;
  TotalLaps: number;
  Sector1Segments: number;
  Sector2Segments: number;
  Sector3Segments: number;
  TotalSegments: number;
  SegmentFlags: number[];
  PitExitOpen: boolean;
  TrackStatus: number;
  SafetyCar: number;
  RemainingTime: number;
  SessionStartTime: Date;
  ClockStopped: boolean;
  DRSEnabled: number;
};

export type Meeting = {
  Circuit: { Key: number; ShortName: string };
  Country: { Code: string; Key: number; Name: string };
  Key: number;
  Location: string;
  Name: string;
  Number: number;
  OfficialName: string;
};
