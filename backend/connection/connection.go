package connection

import (
	"time"
)

type Payload struct {
	Name      string
	Data      []byte
	Timestamp string
}

type Connection interface {
	Connect() (error, <-chan Payload)

	IncrementTime(amount time.Duration)

	JumpToStart() time.Time
}
