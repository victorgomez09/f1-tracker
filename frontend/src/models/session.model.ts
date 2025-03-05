export enum Session {
	Practice1 = 0,
	Practice2 = 1,
	Practice3 = 2,
	Qualifying0 = 3,
	Qualifying1 = 4,
	Qualifying2 = 5,
	Qualifying3 = 6,
	Sprint = 7,
	Race = 8,
	PreSeason = 9
}

export enum HistoricalSession {
	Practice1Session = 0,
	Practice2Session = 1,
	Practice3Session = 2,
	QualifyingSession = 3,
	SprintSession = 4,
	RaceSession = 5,
	PreSeasonSession = 6
}

export enum Status {
  UnknownState = 0,
	Inactive = 1,
	Started = 2,
	Aborted = 3,
	Finished = 4,
	Finalised = 5,
	Ended = 6
}

export enum Drs {
  DRSUnknown = 0,
	DRSEnabled = 1,
	DRSDisabled = 2
}

export enum SafeftyCar {
  Clear = 0,
	VirtualSafetyCar = 1,
	VirtualSafetyCarEnding = 2,
	SafetyCar = 3,
	SafetyCarEnding = 4
}