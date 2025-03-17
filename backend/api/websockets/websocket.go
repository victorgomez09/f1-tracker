package api

import (
	"fmt"
	"sync"

	"github.com/f1gopher/f1gopherlib/flowControl"
	"github.com/f1gopher/f1gopherlib/internal/parser"
	providers "github.com/f1gopher/f1gopherlib/internal/providers"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

type DataStruct struct {
	DataType string `json:"dataType"`
	Data     any    `json:"data"`
}

func HandleHistoricalWs(c echo.Context) error {
	eventName := c.Param("eventName")
	for _, r := range providers.RaceHistory() {
		if r.Name == eventName {
			const dataSources = parser.EventTime | parser.Timing | parser.Event | parser.RaceControl |
				parser.TeamRadio | parser.Weather | parser.Location | parser.Telemetry | parser.Drivers
			replayConnection, _ := providers.CreateReplay(
				dataSources,
				r,
				"./.cache", flowControl.Realtime)
			if replayConnection == nil {
				fmt.Println("There is no replay session.")
				return nil
			}

			websocket.Handler(func(ws *websocket.Conn) {
				defer ws.Close()
				for {
					var dataLock sync.Mutex
					// Send session data
					dataLock.Lock()
					err := websocket.JSON.Send(ws, &DataStruct{
						DataType: "SESSION",
						Data:     replayConnection.Session(),
					})
					if err != nil {
						ws.Close()
						c.Logger().Debug(err)
					}
					dataLock.Unlock()
					// General data
					dataLock.Lock()
					err2 := websocket.JSON.Send(ws, &DataStruct{
						DataType: "GENERAL",
						Data:     replayConnection.TimeLostInPitlane(),
					})
					if err2 != nil {
						ws.Close()
						c.Logger().Debug(err2)
					}
					dataLock.Unlock()
					// Circuit data
					dataLock.Lock()
					err3 := websocket.JSON.Send(ws, &DataStruct{
						DataType: "INFORMATION",
						Data:     replayConnection.CircuitTimezone().String(),
					})
					if err3 != nil {
						ws.Close()
						c.Logger().Debug(err3)
					}
					dataLock.Unlock()

					go func() {
						for {
							select {
							case msg := <-replayConnection.Drivers():
								dataLock.Lock()
								var numbers []int
								for _, driver := range msg.Drivers {
									numbers = append(numbers, driver.Number)
								}
								replayConnection.SelectTelemetrySources(numbers)
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "DRIVERS",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()

							case msg := <-replayConnection.Timing():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "TIMING",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()
							// 	// TODO - sometimes get empty records on shutdown so filter these out
							// 	// if msg.Position == 0 {
							// 	// 	continue
							// 	// }

							case msg := <-replayConnection.Event():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "EVENT",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()

							case msg := <-replayConnection.Time():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "TIME",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()

							case msg := <-replayConnection.RaceControlMessages():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "RACE_CONTROL",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
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

							case msg := <-replayConnection.Location():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "LOCATION",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()

							case msg := <-replayConnection.Telemetry():
								dataLock.Lock()
								err := websocket.JSON.Send(ws, &DataStruct{
									DataType: "TELEMETRY",
									Data:     msg,
								})
								if err != nil {
									c.Logger().Debug(err)
									break
								}
								dataLock.Unlock()
							}
						}
					}()

					// Read
					msg := ""
					err = websocket.Message.Receive(ws, &msg)
					if err != nil {
						c.Logger().Error(err)
					}
					fmt.Printf("%s\n", msg)
				}
			}).ServeHTTP(c.Response(), c.Request())
		}
	}

	return nil
}
