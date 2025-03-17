package parser

import (
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
	"github.com/f1gopher/f1gopherlib/connection"
)

func (p *Parser) parseHeartbeatData(dat map[string]interface{}, timestamp time.Time) (Messages.Event, error) {

	time := dat["Utc"].(string)
	value, err := parseTime(time)
	if err != nil {
		p.ParseTimeError(connection.HeartbeatFile, timestamp, "Utc", err)
	} else {
		p.eventState.Timestamp = value
	}

	// TODO - ignore _kf or flag when no heartbeat recieved?

	p.eventState.Heartbeat = true

	return p.eventState, nil
}
