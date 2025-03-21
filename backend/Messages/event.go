package Messages

import (
	"time"
)

type SessionType int

const (
	Practice1Session SessionType = iota
	Practice2Session
	Practice3Session
	QualifyingSession
	SprintSession
	RaceSession
	PreSeasonSession
)

func (s SessionType) String() string {
	return [...]string{"Practice 1", "Practice 2", "Practice 3", "Qualifying", "Sprint", "Race", "Pre-Season Test"}[s]
}

type EventType int

const (
	Practice1 EventType = iota
	Practice2
	Practice3
	Qualifying0
	Qualifying1
	Qualifying2
	Qualifying3
	Sprint
	Race
	PreSeason
)

func (e EventType) String() string {
	return [...]string{"Practice 1", "Practice 2", "Practice 3", "Qualifying 0", "Qualifying 1", "Qualifying 2", "Qualifying 3", "Sprint", "Race", "Pre-season"}[e]
}

type TrackState int

const (
	Clear TrackState = iota
	VirtualSafetyCar
	VirtualSafetyCarEnding
	SafetyCar
	SafetyCarEnding
)

func (t TrackState) String() string {
	return [...]string{"Clear", "VSC Deployed", "VSC Ending", "Deployed", "Ending"}[t]
}

type FlagState int

const (
	NoFlag FlagState = iota
	GreenFlag
	YellowFlag
	DoubleYellowFlag
	RedFlag
	ChequeredFlag
	BlueFlag
	BlackAndWhite
)

func (f FlagState) String() string {
	return [...]string{"None", "Green", "Yellow", "Double Yellow", "Red", "Chequered", "Blue", "Black and White"}[f]
}

type SessionState int

const (
	UnknownState SessionState = iota
	Inactive
	Started
	Aborted
	Finished
	Finalised
	Ended
)

func (s SessionState) String() string {
	return [...]string{"Unknown", "Inactive", "Started", "Aborted", "Finished", "Finalised", "Ended"}[s]
}

type DRSState int

const (
	DRSUnknown DRSState = iota
	DRSEnabled
	DRSDisabled
)

func (d DRSState) String() string {
	return [...]string{"Unknown", "Enabled", "Disabled"}[d]
}

const MaxSegments = 40

type Event struct {
	Timestamp time.Time

	Name    string
	Type    EventType
	Meeting map[string]interface{}

	Status    SessionState
	Heartbeat bool

	CurrentLap      int
	TotalLaps       int
	Sector1Segments int
	Sector2Segments int
	Sector3Segments int
	TotalSegments   int
	SegmentFlags    [MaxSegments]FlagState

	PitExitOpen bool
	TrackStatus FlagState
	SafetyCar   TrackState

	RemainingTime    time.Duration
	SessionStartTime time.Time
	ClockStopped     bool

	DRSEnabled DRSState
}
