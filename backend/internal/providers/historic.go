package provider

import (
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
)

func RaceHistory() []RaceEvent {
	result := make([]RaceEvent, 0)

	for _, session := range sessionHistory {
		sessionEnd := session.EventTime
		switch session.Type {
		case Messages.Practice1Session, Messages.Practice2Session, Messages.Practice3Session:
			sessionEnd = sessionEnd.Add(time.Hour * 1)

		case Messages.QualifyingSession:
			sessionEnd = sessionEnd.Add(time.Hour * 1)

		case Messages.SprintSession:
			sessionEnd = sessionEnd.Add(time.Hour * 1)

		case Messages.RaceSession:
			sessionEnd = sessionEnd.Add(time.Hour * 3)
		}

		if sessionEnd.Before(time.Now()) {
			result = append(result, session)
		}
	}

	return result
}

func HappeningSessions() (liveSession RaceEvent, nextSession RaceEvent, hasLiveSession bool, hasNextSession bool) {
	all := sessionHistory
	utcNow := time.Now().UTC()

	for x := 0; x < len(all); x++ {
		// If we are the same day as a session see if it is live
		if all[x].EventTime.Year() == utcNow.Year() &&
			all[x].EventTime.Month() == utcNow.Month() &&
			all[x].EventTime.Day() == utcNow.Day() {

			// Start up to 55 mins before the start  of the event (50 mins is when they go to grid)
			if utcNow.After(all[x].EventTime.Add(-time.Minute * 55)) {

				duringEvent := false
				switch all[x].Type {
				case Messages.Practice1Session, Messages.Practice2Session, Messages.Practice3Session:
					// Usually 60 mins but tire tests are 90 so cover both since it won't overlap with anything else
					duringEvent = utcNow.Before(all[x].EventTime.Add(time.Hour * 2))

				case Messages.QualifyingSession, Messages.SprintSession, Messages.RaceSession, Messages.PreSeasonSession:
					// Last events in the day so just assume it's that event
					duringEvent = true

				default:
					panic("History: Unhandled session type: " + all[x].Type.String())
				}

				if duringEvent {
					if x == 0 {
						return all[x], RaceEvent{}, true, false
					}

					return all[x], all[x-1], true, true
				}
			}
		} else if all[x].EventTime.Before(utcNow) {
			if x == 0 {
				// No live or upcoming sessions
				break
			}

			// If this is the first session before now then the next session is the one that came before it
			return RaceEvent{}, all[x-1], false, true
		}
	}

	return RaceEvent{}, RaceEvent{}, false, false
}

func liveEvent() (event RaceEvent, exists bool) {
	live, _, exists, _ := HappeningSessions()

	return live, exists
}
