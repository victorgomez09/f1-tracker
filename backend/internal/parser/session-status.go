package parser

import (
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
	"github.com/f1gopher/f1gopherlib/connection"
)

func (p *Parser) parseSessionStatusData(dat map[string]interface{}, timestamp time.Time) (Messages.Event, error) {

	status := dat["Status"].(string)

	switch status {
	case "Inactive":
		p.eventState.Status = Messages.Inactive
	case "Started":
		p.eventState.Status = Messages.Started
	case "Aborted":
		p.eventState.Status = Messages.Aborted
	case "Finished":
		p.eventState.Status = Messages.Finished
	case "Finalised":
		p.eventState.Status = Messages.Finalised
	case "Ends":
		p.eventState.Status = Messages.Ended
	default:
		p.ParseErrorf(connection.SessionStatusFile, timestamp, "SessionStatus: Unhandled Status '%s'", status)
	}

	p.eventState.Timestamp = timestamp

	return p.eventState, nil
}
