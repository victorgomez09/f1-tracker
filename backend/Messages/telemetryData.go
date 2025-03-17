package Messages

import (
	"time"
)

type Telemetry struct {
	Timestamp    time.Time
	DriverNumber int

	RPM      int16
	Speed    float32
	Gear     byte
	Throttle float32
	Brake    float32
	DRS      bool
}
