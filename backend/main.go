package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/f1gopher/f1gopherlib/flowControl"
	"github.com/f1gopher/f1gopherlib/parser"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var (
// 	upgrader = websocket.Upgrader{}
// )

var (
	replayConnection F1GopherLib
	err              error
	lock             = sync.Mutex{}
)

func main() {
	// GetData()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://5173-victorgomez09-f1tracker-rkauu33dlup.ws-eu118.gitpod.io", "https://stunning-system-j4wxj4p5v4j3555p-5173.app.github.dev", "http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// e.GET("/ws", HandleWebsocket)
	e.POST("/actions", HandleActions)
	e.GET("/historical", HandleHistorical)
	e.GET("/historical/:eventName", HandleHistoricalEvent)
	e.Logger.Fatal(e.Start(":3000"))
}

type actions struct {
	SkipToStart   bool
	Skip5Secs     bool
	SkipMinute    bool
	Skip10Minutes bool
	Pause         bool
}

func HandleActions(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	var a actions

	if err := c.Bind(&a); err != nil {
		fmt.Println("erro", err)
		return err
	}

	if a.SkipToStart {
		replayConnection.SkipToSessionStart()
	} else if a.SkipMinute {
		replayConnection.IncrementTime(time.Minute * 1)
	} else if a.Skip10Minutes {
		replayConnection.IncrementTime(time.Minute * 10)
	} else if a.Skip5Secs {
		replayConnection.IncrementTime(time.Second * 5)
	} else {
		replayConnection.TogglePause()
	}

	return c.JSON(http.StatusCreated, replayConnection.IsPaused())
}

// func HandleWebsocket(c echo.Context) error {
// 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	// defer ws.Close()
// 	// websocket.Handler(func(ws *websocket.Conn) {
// 	defer ws.Close()

// 	// GetData(c, ws)
// 	// }).ServeHTTP(c.Response(), c.Request())

// 	return nil
// }

func HandleHistorical(c echo.Context) error {
	var result []any
	for _, r := range RaceHistory() {
		result = append(result, r)
	}
	c.JSON(http.StatusOK, result)

	return nil
}

func HandleHistoricalEvent(c echo.Context) error {
	var result []any
	eventName := c.Param("eventName")
	for _, r := range RaceHistory() {
		if r.Name == eventName {
			fmt.Println("getting ws")
			var origins = []string{"http://127.0.0.1:18081", "http://localhost:18081", "https://stunning-system-j4wxj4p5v4j3555p-5173.app.github.dev"}
			var upgrader = websocket.Upgrader{
				// Resolve cross-domain problems
				CheckOrigin: func(r *http.Request) bool {
					var origin = r.Header.Get("origin")
					for _, allowOrigin := range origins {
						if origin == allowOrigin {
							return true
						}
					}
					return false
				}}
			ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
			if err != nil {
				fmt.Println("Error getting WS", err)
				return err
			}
			// defer ws.Close()

			GetHistoricalData(c, ws, &r)

			// // Read
			// _, msg, err := ws.ReadMessage()
			// if err != nil {
			// 	c.Logger().Error(err)
			// }
			// fmt.Printf("%s\n", msg)
		}
	}
	c.JSON(http.StatusOK, result)

	return nil
}

type DataStruct struct {
	DataType string `json:"dataType"`
	Data     any    `json:"data"`
}

// func GetData(c echo.Context, ws *websocket.Conn) F1GopherLib {
// 	const dataSources = parser.EventTime | parser.Timing | parser.Event | parser.RaceControl |
// 		parser.TeamRadio | parser.Weather | parser.Location | parser.Telemetry | parser.Drivers
// 	liveConnection, _ := CreateLive(
// 		dataSources,
// 		"",
// 		"./.cache")
// 	if liveConnection == nil {
// 		fmt.Println("There is no live session currently happening.")
// 		return nil
// 	}

// 	var dataLock sync.Mutex
// 	// Send session data
// 	dataLock.Lock()
// 	err := websocket.JSON.Send(ws, &DataStruct{
// 		DataType: "SESSION",
// 		Data:     liveConnection.Session(),
// 	})
// 	if err != nil {
// 		c.Logger().Error(err)
// 	}
// 	dataLock.Unlock()
// 	// General data
// 	dataLock.Lock()
// 	err2 := websocket.JSON.Send(ws, &DataStruct{
// 		DataType: "GENERAL",
// 		Data:     liveConnection.TimeLostInPitlane(),
// 	})
// 	if err2 != nil {
// 		c.Logger().Error(err)
// 	}
// 	dataLock.Unlock()

// 	for {
// 		select {
// 		case msg := <-liveConnection.Drivers():
// 			dataLock.Lock()
// 			var numbers []int
// 			for _, driver := range msg.Drivers {
// 				numbers = append(numbers, driver.Number)
// 			}
// 			liveConnection.SelectTelemetrySources(numbers)
// 			// fmt.Println("msg", msg)
// 			err := websocket.JSON.Send(ws, &DataStruct{
// 				DataType: "DRIVERS",
// 				Data:     msg,
// 			})
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			dataLock.Unlock()
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessDrivers(msg)
// 			// }

// 		case msg := <-liveConnection.Timing():
// 			dataLock.Lock()
// 			err := websocket.JSON.Send(ws, &DataStruct{
// 				DataType: "TIMING",
// 				Data:     msg,
// 			})
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			dataLock.Unlock()
// 			// TODO - sometimes get empty records on shutdown so filter these out
// 			// if msg.Position == 0 {
// 			// 	continue
// 			// }

// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessTiming(msg)
// 			// }

// 		case msg := <-liveConnection.Event():
// 			dataLock.Lock()
// 			err := websocket.JSON.Send(ws, &DataStruct{
// 				DataType: "EVENT",
// 				Data:     msg,
// 			})
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			dataLock.Unlock()

// 		case <-liveConnection.Time():
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessEventTime(msg)
// 			// }

// 		case <-liveConnection.RaceControlMessages():
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessRaceControlMessages(msg)
// 			// }

// 		case <-liveConnection.Weather():
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessWeather(msg)
// 			// }

// 		case <-liveConnection.Radio():
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessRadio(msg)
// 			// }

// 		case <-liveConnection.Location():
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessLocation(msg)
// 			// }

// 		case msg := <-liveConnection.Telemetry():
// 			dataLock.Lock()
// 			// fmt.Println("msg", msg)
// 			err := websocket.JSON.Send(ws, &DataStruct{
// 				DataType: "TELEMETRY",
// 				Data:     msg,
// 			})
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			dataLock.Unlock()
// 			// for x := range d.panels {
// 			// 	d.panels[x].ProcessTelemetry(msg)
// 			// }
// 		}

// 		// Data has changed so force a UI redraw
// 		// giu.Update()
// 	}
// }

func GetHistoricalData(c echo.Context, ws *websocket.Conn, event *RaceEvent) F1GopherLib {
	const dataSources = parser.EventTime | parser.Timing | parser.Event | parser.RaceControl |
		parser.TeamRadio | parser.Weather | parser.Location | parser.Telemetry | parser.Drivers
	replayConnection, _ = CreateReplay(
		dataSources,
		*event,
		"./.cache", flowControl.Realtime)
	if replayConnection == nil {
		fmt.Println("There is no replay session.")
		return nil
	}

	var dataLock sync.Mutex
	// Send session data
	dataLock.Lock()
	// err := websocket.JSON.Send(ws, &DataStruct{
	err := ws.WriteJSON(&DataStruct{
		DataType: "SESSION",
		Data:     replayConnection.Session(),
	})
	if err != nil {
		c.Logger().Error(err)
	}
	dataLock.Unlock()
	// General data
	dataLock.Lock()
	err2 := ws.WriteJSON(&DataStruct{
		DataType: "GENERAL",
		Data:     replayConnection.TimeLostInPitlane(),
	})
	if err2 != nil {
		c.Logger().Error(err2)
	}
	dataLock.Unlock()
	// Circuit data
	dataLock.Lock()
	err3 := ws.WriteJSON(&DataStruct{
		DataType: "INFORMATION",
		Data:     replayConnection.CircuitTimezone().String(),
	})
	if err3 != nil {
		c.Logger().Error(err3)
	}
	dataLock.Unlock()

	for {
		select {
		case msg := <-replayConnection.Drivers():
			dataLock.Lock()
			var numbers []int
			for _, driver := range msg.Drivers {
				numbers = append(numbers, driver.Number)
			}
			replayConnection.SelectTelemetrySources(numbers)
			err := ws.WriteJSON(&DataStruct{
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

		case msg := <-replayConnection.Timing():
			dataLock.Lock()
			err := ws.WriteJSON(&DataStruct{
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

		case msg := <-replayConnection.Event():
			dataLock.Lock()
			err := ws.WriteJSON(&DataStruct{
				DataType: "EVENT",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()

		case msg := <-replayConnection.Time():
			dataLock.Lock()
			err := ws.WriteJSON(&DataStruct{
				DataType: "TIME",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()
			// for x := range d.panels {
			// 	d.panels[x].ProcessEventTime(msg)
			// }

		case msg := <-replayConnection.RaceControlMessages():
			dataLock.Lock()
			err := ws.WriteJSON(&DataStruct{
				DataType: "RACE_CONTROL",
				Data:     msg,
			})
			if err != nil {
				c.Logger().Error(err)
			}
			dataLock.Unlock()

		case <-replayConnection.Weather():
			// for x := range d.panels {
			// 	d.panels[x].ProcessWeather(msg)
			// }

		case <-replayConnection.Radio():
			// for x := range d.panels {
			// 	d.panels[x].ProcessRadio(msg)
			// }

		case <-replayConnection.Location():
			// for x := range d.panels {
			// 	d.panels[x].ProcessLocation(msg)
			// }

		case msg := <-replayConnection.Telemetry():
			dataLock.Lock()
			err := ws.WriteJSON(&DataStruct{
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

		// Read
		// _, read, errRes := ws.ReadMessage()
		// if errRes != nil {
		// 	c.Logger().Error(errRes)
		// }
		// if len(read) > 0 {
		// 	fmt.Printf("%s\n", read)
		// }
	}
}
