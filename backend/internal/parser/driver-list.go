package parser

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
	"github.com/f1gopher/f1gopherlib/connection"
)

func (p *Parser) parseDriverList(dat map[string]interface{}, timestamp time.Time) []Messages.Drivers {
	var driver []Messages.Drivers = nil

	for driverNum, info := range dat {
		if driverNum == "_kf" {
			continue
		}

		record := info.(map[string]interface{})

		current, exists := p.driverTimes[driverNum]

		if !exists {
			number, _ := strconv.Atoi(driverNum)

			line := 0
			rawLine, exists := record["Line"]
			if exists {
				line = int(rawLine.(float64))
			}

			fullName, _ := record["FullName"].(string)
			shortName, _ := record["Tla"].(string)
			// TeamName and TeamColor do not always exist
			teamName, _ := record["TeamName"].(string)
			teamHexColour, colorExists := record["TeamColour"].(string)

			// Default colors
			teamColor := color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
			if colorExists {
				_, err := fmt.Sscanf(teamHexColour, "%02x%02x%02x", &teamColor.R, &teamColor.G, &teamColor.B)
				if err != nil {
					p.ParseErrorf(connection.DriverListFile, timestamp, "Unable to parse team color: '%s', %v", teamColor, err)
				}
			}

			current = Messages.Timing{
				Number:    number,
				Position:  line,
				Name:      fullName,
				ShortName: shortName,
				Team:      teamName,
				HexColor:  "#" + teamHexColour,
				Color:     teamColor,
			}

			if driver == nil {
				driver = []Messages.Drivers{
					{
						Timestamp: timestamp,
						Drivers:   nil,
					},
				}
			}

			driver[0].Drivers = append(driver[0].Drivers, Messages.DriverInfo{
				StartPosition: current.Position,
				Name:          current.Name,
				ShortName:     current.ShortName,
				Number:        current.Number,
				Team:          current.Team,
				HexColor:      current.HexColor,
				Color:         current.Color,
			})
		}

		p.driverTimes[driverNum] = current
	}

	return driver
}
