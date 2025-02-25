package model

type SocketData struct {
	C string
	M []SocketDataPayload
	G string
	H string
	I string
	R F1Recap
}

type SocketDataPayload struct {
	A []string
	H string
	M string
}

type ParsedRecap struct {
	CarData  F1CarData
	Position F1Position
}

type F1State struct {
	ParsedRecap
	F1Recap
}

type F1CarData struct {
	Entries []F1Entry
}

type F1Entry struct {
	Utc  string
	Cars []F1CarDataChannels
}

type F1Position struct {
	Position []F1PositionItem
}

type F1PositionItem struct {
	Timestamp string
	Entries   []F1PositionCar
}

type F1PositionCar struct {
	Status string
	X      int
	Y      int
	Z      int
}

/**
 * @namespace
 * @property {number} 0 - RPM
 * @property {number} 2 - Speed number km/h
 * @property {number} 3 - gear number
 * @property {number} 4 - Throttle int 0-100
 * @property {number} 5 - Brake number boolean
 * @property {number} 45 - DRS
 */
type F1CarDataChannels struct {
	RPM         int "json:0"
	SpeedNumber int "json:2"
	GearNumber  int "json:3"
	Throttle    int "json:4"
	Brake       int "json:4"
	DRS         int "json:45"
}

type F1Recap struct {
	Heartbeat           F1Heartbeat
	ExtrapolatedClock   F1ExtrapolatedClock
	TopThree            F1TopThree
	TimingStats         F1TimingStats
	TimingAppData       F1TimingAppData
	WeatherData         F1WeatherData
	TrackStatus         F1TrackStatus
	DriverList          []F1Driver
	RaceControlMessages []F1Message
	SessionInfo         F1SessionInfo
	SessionData         F1SessionData
	LapCount            F1LapCount
	TimingData          F1TimingData
	TeamRadio           F1TeamRadio

	CarData  string `json:"CarData.z"`
	Position string `json:"Position.z"`
}

type F1TeamRadio struct {
	Captures []F1RadioCapture
	_kf      bool
}

type F1RadioCapture struct {
	Utc          string
	RacingNumber string
	Path         string
}

type F1TimingData struct {
	NoEntries        []int
	SessionPart      int
	CutOffTime       string
	CutOffPercentage string
	Lines            []F1TimingDataDriver
	Withheld         bool
	_kf              bool
}

type F1TimingDataDriver struct {
	Stats []struct {
		TimeDiffToFastest       string
		TimeDifftoPositionAhead string
	}
	TimeDiffToFastest       string
	TimeDiffToPositionAhead string
	GapToLeader             string
	IntervalToPositionAhead struct {
		Value    string
		Catching bool
	}
	Line         int
	Position     string
	ShowPosition bool
	RacingNumber string
	Retired      bool
	InPit        bool
	PitOut       bool
	Stopped      bool
	Status       int
	Sectors      []F1Sector
	Speeds       F1Speeds
	BestLapTime  F1PersonalBestLapTime
	LastLapTime  F1I1
	NumberOfLaps int
	KnockedOut   bool
	Cutoff       bool
}

type F1Sector struct {
	Stopped         bool
	Value           string
	PreviousValue   string
	Status          int
	OverallFastest  bool
	PersonalFastest bool
	Segments        []struct {
		Status int
	}
}

type F1Speeds struct {
	I1 F1I1
	I2 F1I1
	FL F1I1
	ST F1I1
}

type F1I1 struct {
	Value           string
	Status          int
	OverallFastest  bool
	PersonalFastest bool
}

type F1LapCount struct {
	CurrentLap int8
	TotalLaps  int8
	_kf        bool
}

type F1SessionData struct {
	Series       []F1Series
	StatusSeries []F1StatusSeries
	_kf          bool
}

type F1StatusSeries struct {
	Utc         string
	TrackStatus string
}

type F1Series struct {
	Utc      string
	Lap      int
	Position string "json:Position.z"
}

type F1SessionInfo struct {
	Meeting       F1Meeting
	ArchiveStatus F1ArchiveStatus
	Key           int
	Type          string
	Name          string
	StartDate     string
	EndDate       string
	GmtOffset     string
	Path          string
	Number        int
	_kf           bool
}

type F1ArchiveStatus struct {
	Status string
}

type F1Meeting struct {
	Key          int
	Name         string
	OfficialName string
	Location     string
	Country      F1Country
	Circuit      F1Circuit
}

type F1Circuit struct {
	Key       int
	ShortName string
}

type F1Country struct {
	Key  int
	Code string
	Name string
}

type RadioStatus string

const (
	ENABLED  RadioStatus = "ENABLED"
	DISABLED RadioStatus = "DISABLED"
)

type F1Message struct {
	Utc      string
	Lap      int8
	Category string
	Message  string
	Flag     string
	Scope    string
	Sector   int
	Status   RadioStatus
}

type F1Driver struct {
	RacingNumber  string
	BroadcastName string
	FullName      string
	Tla           string
	Line          int
	TeamName      string
	TeamColour    string
	FirstName     string
	LastName      string
	Reference     string
	HeadshotUrl   string
	CountryCode   string
}

type F1TimingAppData struct {
	Lines []Line
	_kf   bool
}

type Line struct {
	RacingNumber string
	Stints       []F1Stint
	Line         int
	GridPos      string
}

type F1Heartbeat struct {
	Utc string
	_kf bool
}

type F1ExtrapolatedClock struct {
	Utc           string
	Remaining     string
	Extrapolating bool
	_kf           bool
}

type F1TopThree struct {
	Withheld bool
	Lines    []F1TopThreeDriver
	_kf      bool
}

type F1TopThreeDriver struct {
	Position        string
	ShowPosition    bool
	RacingNumber    string
	Tla             string
	BroadcastName   string
	FullName        string
	Team            string
	TeamColour      string
	LapTime         string
	LapState        int8
	DiffToAhead     string
	DiffToLeader    string
	OverallFastest  bool
	PersonalFastest bool
}

type F1TimingStats struct {
	Withheld    bool
	Lines       []F1TimingStatsDriver
	Sessiontype string
	_kf         bool
}

type F1TimingStatsDriver struct {
	Line                int16
	RacingNumber        string
	PersonalBestLapTime F1PersonalBestLapTime
	BestSectors         []F1PersonalBestLapTime
	BestSpeeds          struct {
		I1 F1PersonalBestLapTime
		I2 F1PersonalBestLapTime
		FL F1PersonalBestLapTime
		ST F1PersonalBestLapTime
	}
}

type F1PersonalBestLapTime struct {
	Value    string
	Position int8
}

type TyreType string

const (
	SOFT         TyreType = "SOFT"
	MEDIUM       TyreType = "MEDIUM"
	HARD         TyreType = "HARD"
	INTERMEDIATE TyreType = "INTERMEDIATE"
	WET          TyreType = "WET"
)

type F1Stint struct {
	TotalLaps int8
	Compound  TyreType
	New       string // TRUE | FALSE
}

type F1WeatherData struct {
	AirTemp       string
	Humidity      string
	Pressure      string
	Rainfall      string
	TrackTemp     string
	WindDirection string
	WindSpeed     string
	_kf           bool
}

type F1TrackStatus struct {
	Status  string
	Message string
	_kf     bool
}
