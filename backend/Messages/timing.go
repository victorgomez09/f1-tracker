package Messages

import (
	"image/color"
	"time"
)

type CarLocation int

// TODO - add garage and grid - need to calculate these based on speed and session type
const (
	NoLocation CarLocation = iota
	Pitlane
	PitOut
	OutLap
	OnTrack
	OutOfRace
	Stopped
)

func (c CarLocation) String() string {
	return [...]string{"Unknown", "Pitlane", "Pit Exit", "Out Lap", "On Track", "Out", "Stopped"}[c]
}

type TireType int

const (
	Unknown TireType = iota
	Soft
	Medium
	Hard
	Intermediate
	Wet
	Test
	HYPERSOFT
	ULTRASOFT
	SUPERSOFT
)

func (t TireType) String() string {
	return [...]string{"", "Soft", "Medium", "Hard", "Inter", "Wet", "Test", "Hyp Soft", "Ult Soft", "Sup Soft"}[t]
}

type SegmentType int

const (
	None SegmentType = iota
	YellowSegment
	GreenSegment
	InvalidSegment // Doesn't get displayed, cut corner/boundaries or invalid segment time?
	PurpleSegment
	RedSegment     // After chequered flag/stopped on track
	PitlaneSegment // In pitlane
	Mystery
	Mystery2 // ??? 2021 - Turkey Practice_2
	Mystery3 // ??? 2020 - Italy Race
)

type PitStop struct {
	Lap          int
	PitlaneEntry time.Time
	PitlaneExit  time.Time
	PitlaneTime  time.Duration
}

type Timing struct {
	Timestamp time.Time

	Position int

	Name      string
	ShortName string
	Number    int
	Team      string
	HexColor  string
	Color     color.RGBA

	TimeDiffToFastest       int64
	TimeDiffToPositionAhead int64
	GapToLeader             int64

	PreviousSegmentIndex   int
	Segment                [MaxSegments]SegmentType
	Sector1                int64
	Sector1PersonalFastest bool
	Sector1OverallFastest  bool
	Sector2                int64
	Sector2PersonalFastest bool
	Sector2OverallFastest  bool
	Sector3                int64
	Sector3PersonalFastest bool
	Sector3OverallFastest  bool
	LastLap                int64
	LastLapPersonalFastest bool
	LastLapOverallFastest  bool

	FastestLap        int64
	OverallFastestLap bool

	KnockedOutOfQualifying bool
	ChequeredFlag          bool

	Tire       TireType
	LapsOnTire int
	Lap        int

	DRSOpen bool

	Pitstops     int
	PitStopTimes []PitStop

	Location CarLocation

	SpeedTrap                int
	SpeedTrapPersonalFastest bool
	SpeedTrapOverallFastest  bool
}
