package parser

import (
	"math"
	"strconv"
	"time"

	"github.com/f1gopher/f1gopherlib/Messages"
	"github.com/f1gopher/f1gopherlib/connection"
)

func (p *Parser) parsePositionData(dat map[string]interface{}, timestamp time.Time) ([]Messages.Location, error) {

	result := make([]Messages.Location, 0)
	const tolerance = 0.000001

	for _, record := range dat["Position"].([]interface{}) {
		timestampStr := record.(map[string]interface{})["Timestamp"].(string)
		dataTimestamp, err := parseTime(timestampStr)
		if err != nil {
			p.ParseTimeError(connection.PositionFile, timestamp, "Timestamp", err)
		}

		for key, entry := range record.(map[string]interface{})["Entries"].(map[string]interface{}) {
			driver, _ := strconv.ParseInt(key, 10, 8)
			//status := entry.(map[string]interface{})["Status"].(string)

			x := entry.(map[string]interface{})["X"].(float64)
			y := entry.(map[string]interface{})["Y"].(float64)
			z := entry.(map[string]interface{})["Z"].(float64)

			// Ignore locations which are (0, 0) because it means we don't have a location for them
			if math.Abs(x) < tolerance && math.Abs(y) < tolerance {
				continue
			}

			result = append(result, Messages.Location{
				Timestamp:    dataTimestamp,
				DriverNumber: int(driver),
				X:            x,
				Y:            y,
				Z:            z,
			})
		}
	}

	return result, nil
}
