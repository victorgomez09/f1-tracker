package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/f1gopher/f1gopherlib/parser"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	// GetData()

	e := echo.New()
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		go GetData(c)
	// 		return next(c)
	// 	}
	// })
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://5173-victorgomez09-f1tracker-rkauu33dlup.ws-eu118.gitpod.io"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/ws", HandleWebsocket)
	e.Logger.Fatal(e.Start(":3000"))
}

func HandleWebsocket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		GetData(c, ws)
		// for {
		// 	// Write
		// 	err := websocket.Message.Send(ws, "Hello, Client!")
		// 	if err != nil {
		// 		c.Logger().Error(err)
		// 	}

		// }
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

type DataStruct struct {
	DataType string `json:"dataType"`
	Data     any    `json:"data"`
}

func GetData(c echo.Context, ws *websocket.Conn) F1GopherLib {
	const dataSources = parser.EventTime | parser.Timing | parser.Event | parser.RaceControl |
		parser.TeamRadio | parser.Weather | parser.Location | parser.Telemetry | parser.Drivers
	liveConnection, _ := CreateLive(
		dataSources,
		"",
		"./.cache")
	if liveConnection == nil {
		fmt.Println("There is no live session currently happening.")
	}

	var dataLock sync.Mutex
	// Send session data
	dataLock.Lock()
	err := websocket.JSON.Send(ws, &DataStruct{
		DataType: "SESSION",
		Data:     liveConnection.Session(),
	})
	if err != nil {
		c.Logger().Error(err)
	}
	dataLock.Unlock()
	// General data
	dataLock.Lock()
	err2 := websocket.JSON.Send(ws, &DataStruct{
		DataType: "GENERAL",
		Data:     liveConnection.TimeLostInPitlane(),
	})
	if err2 != nil {
		c.Logger().Error(err)
	}
	dataLock.Unlock()

	for {
		select {
		case msg := <-liveConnection.Drivers():
			dataLock.Lock()
			var numbers []int
			for _, driver := range msg.Drivers {
				numbers = append(numbers, driver.Number)
			}
			liveConnection.SelectTelemetrySources(numbers)
			// fmt.Println("msg", msg)
			err := websocket.JSON.Send(ws, &DataStruct{
				DataType: "DRIVERS",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()
			// for x := range d.panels {
			// 	d.panels[x].ProcessDrivers(msg)
			// }

		case msg := <-liveConnection.Timing():
			dataLock.Lock()
			err := websocket.JSON.Send(ws, &DataStruct{
				DataType: "TIMING",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()
			// TODO - sometimes get empty records on shutdown so filter these out
			// if msg.Position == 0 {
			// 	continue
			// }

			// for x := range d.panels {
			// 	d.panels[x].ProcessTiming(msg)
			// }

		case msg := <-liveConnection.Event():
			dataLock.Lock()
			err := websocket.JSON.Send(ws, &DataStruct{
				DataType: "EVENT",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()

		case <-liveConnection.Time():
			// for x := range d.panels {
			// 	d.panels[x].ProcessEventTime(msg)
			// }

		case <-liveConnection.RaceControlMessages():
			// for x := range d.panels {
			// 	d.panels[x].ProcessRaceControlMessages(msg)
			// }

		case <-liveConnection.Weather():
			// for x := range d.panels {
			// 	d.panels[x].ProcessWeather(msg)
			// }

		case <-liveConnection.Radio():
			// for x := range d.panels {
			// 	d.panels[x].ProcessRadio(msg)
			// }

		case <-liveConnection.Location():
			// for x := range d.panels {
			// 	d.panels[x].ProcessLocation(msg)
			// }

		case msg := <-liveConnection.Telemetry():
			dataLock.Lock()
			// fmt.Println("msg", msg)
			err := websocket.JSON.Send(ws, &DataStruct{
				DataType: "TELEMETRY",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()
			// for x := range d.panels {
			// 	d.panels[x].ProcessTelemetry(msg)
			// }
		}

		// Data has changed so force a UI redraw
		// giu.Update()
	}
}
