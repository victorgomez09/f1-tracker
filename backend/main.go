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
	liveConnection, _ := CreateLive(
		parser.EventTime|parser.Timing|parser.Event|parser.RaceControl|parser.TeamRadio|parser.Weather,
		"",
		"./.cache")
	if liveConnection == nil {
		fmt.Println("There is no live session currently happening.")
	}

	var dataLock sync.Mutex
	for {
		select {
		case <-liveConnection.Drivers():
			fmt.Println("DRIVERS")
			// for x := range d.panels {
			// 	d.panels[x].ProcessDrivers(msg)
			// }

		case msg := <-liveConnection.Timing():
			fmt.Println("TIMING")
			dataLock.Lock()
			// fmt.Println("msg", msg)
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

		case <-liveConnection.Event():
			fmt.Println("EVENTS")
			// d.eventLock.Lock()
			// d.event = msg
			// d.eventLock.Unlock()

			// for x := range d.panels {
			// 	d.panels[x].ProcessEvent(msg)
			// }

		case <-liveConnection.Time():
			fmt.Println("TIME")
			// for x := range d.panels {
			// 	d.panels[x].ProcessEventTime(msg)
			// }

		case msg := <-liveConnection.RaceControlMessages():
			fmt.Println("RCMS")
			dataLock.Lock()
			fmt.Println("msg", msg)
			dataLock.Unlock()
			// for x := range d.panels {
			// 	d.panels[x].ProcessRaceControlMessages(msg)
			// }

		case <-liveConnection.Weather():
			fmt.Println("WEATHER")
			// for x := range d.panels {
			// 	d.panels[x].ProcessWeather(msg)
			// }

		case <-liveConnection.Radio():
			fmt.Println("RADIO")
			// for x := range d.panels {
			// 	d.panels[x].ProcessRadio(msg)
			// }

		case <-liveConnection.Location():
			fmt.Println("LOCATION")
			// for x := range d.panels {
			// 	d.panels[x].ProcessLocation(msg)
			// }

		case <-liveConnection.Telemetry():
			fmt.Println("TELEMETRY")
			// for x := range d.panels {
			// 	d.panels[x].ProcessTelemetry(msg)
			// }
		}

		// Data has changed so force a UI redraw
		// giu.Update()
	}

	return liveConnection
}
