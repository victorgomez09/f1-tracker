package main

import (
	historic "github.com/f1gopher/f1gopherlib/api/handlers/historic"
	websocket "github.com/f1gopher/f1gopherlib/api/websockets"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/historical", historic.HandleHistoric)
	e.GET("/historical/:eventName", websocket.HandleHistoricalWs)

	go e.Logger.Fatal(e.Start(":3000"))
}

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/f1gopher/f1gopherlib/flowControl"
// 	"github.com/f1gopher/f1gopherlib/parser"
// 	"github.com/gorilla/websocket"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// // var (
// // 	upgrader = websocket.Upgrader{}
// // )

// type Server struct {
// 	clients map[*websocket.Conn]bool
// }

// var (
// 	replayConnection F1GopherLib
// 	lock             = sync.Mutex{}
// )

// func StartServer() *Server {
// 	e := echo.New()
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"http://localhost:5173"},
// 		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
// 	}))

// 	server := Server{
// 		make(map[*websocket.Conn]bool),
// 	}

// 	// e.GET("/ws", HandleWebsocket)
// 	e.POST("/actions", HandleActions)
// 	e.GET("/historical", HandleHistorical)
// 	e.GET("/historical/:eventName", server.HandleHistoricalEvent)
// 	go e.Logger.Fatal(e.Start(":3000"))
// 	// go http.ListenAndServe(":8080", nil)

// 	return &server
// }

// func main() {
// 	server := StartServer()

// 	for {
// 		server.WriteMessage([]byte("Hello"))
// 	}
// }

// type actions struct {
// 	SkipToStart   bool
// 	Skip5Secs     bool
// 	SkipMinute    bool
// 	Skip10Minutes bool
// 	Pause         bool
// }

// func HandleActions(c echo.Context) error {
// 	lock.Lock()
// 	defer lock.Unlock()

// 	var a actions

// 	if err := c.Bind(&a); err != nil {
// 		fmt.Println("erro", err)
// 		return err
// 	}

// 	if a.SkipToStart {
// 		replayConnection.SkipToSessionStart()
// 	} else if a.SkipMinute {
// 		replayConnection.IncrementTime(time.Minute * 1)
// 	} else if a.Skip10Minutes {
// 		replayConnection.IncrementTime(time.Minute * 10)
// 	} else if a.Skip5Secs {
// 		replayConnection.IncrementTime(time.Second * 5)
// 	} else {
// 		replayConnection.TogglePause()
// 	}

// 	return c.JSON(http.StatusCreated, replayConnection.IsPaused())
// }

// // func HandleWebsocket(c echo.Context) error {
// // 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// // 	if err != nil {
// // 		return err
// // 	}
// // 	// defer ws.Close()
// // 	// websocket.Handler(func(ws *websocket.Conn) {
// // 	defer ws.Close()

// // 	// GetData(c, ws)
// // 	// }).ServeHTTP(c.Response(), c.Request())

// // 	return nil
// // }

// func HandleHistorical(c echo.Context) error {
// 	var result []any
// 	for _, r := range RaceHistory() {
// 		result = append(result, r)
// 	}
// 	c.JSON(http.StatusOK, result)

// 	return nil
// }

// func (server *Server) HandleHistoricalEvent(c echo.Context) error {
// 	var result []any
// 	eventName := c.Param("eventName")
// 	for _, r := range RaceHistory() {
// 		if r.Name == eventName {
// 			// var origins = []string{"http://localhost:5173"}
// 			// var upgrader = websocket.Upgrader{
// 			// 	// Resolve cross-domain problems
// 			// 	CheckOrigin: func(r *http.Request) bool {
// 			// 		var origin = r.Header.Get("origin")
// 			// 		return slices.Contains(origins, origin)
// 			// 	}}

// 			// ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 			var upgrader = websocket.Upgrader{
// 				CheckOrigin: func(r *http.Request) bool {
// 					return true // Accepting all requests
// 				},
// 			}
// 			ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 			if err != nil {
// 				fmt.Println("Error getting WS", err)
// 				return err
// 			}
// 			server.clients[ws] = true
// 			defer ws.Close()

// 			// GetHistoricalData(c, ws, &r)
// 			go server.WriteMessage([]byte("adsfasdfasdfasd"))
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)

// 	return nil
// }

// func (server *Server) WriteMessage(message []byte) {
// 	for conn := range server.clients {
// 		conn.WriteMessage(websocket.TextMessage, message)
// 	}
// }

// type DataStruct struct {
// 	DataType string `json:"dataType"`
// 	Data     any    `json:"data"`
// }

// // func GetData(c echo.Context, ws *websocket.Conn) F1GopherLib {
// // 	const dataSources = parser.EventTime | parser.Timing | parser.Event | parser.RaceControl |
// // 		parser.TeamRadio | parser.Weather | parser.Location | parser.Telemetry | parser.Drivers
// // 	liveConnection, _ := CreateLive(
// // 		dataSources,
// // 		"",
// // 		"./.cache")
// // 	if liveConnection == nil {
// // 		fmt.Println("There is no live session currently happening.")
// // 		return nil
// // 	}

// // 	var dataLock sync.Mutex
// // 	// Send session data
// // 	dataLock.Lock()
// // 	err := websocket.JSON.Send(ws, &DataStruct{
// // 		DataType: "SESSION",
// // 		Data:     liveConnection.Session(),
// // 	})
// // 	if err != nil {
// // 		c.Logger().Error(err)
// // 	}
// // 	dataLock.Unlock()
// // 	// General data
// // 	dataLock.Lock()
// // 	err2 := websocket.JSON.Send(ws, &DataStruct{
// // 		DataType: "GENERAL",
// // 		Data:     liveConnection.TimeLostInPitlane(),
// // 	})
// // 	if err2 != nil {
// // 		c.Logger().Error(err)
// // 	}
// // 	dataLock.Unlock()

// // 	for {
// // 		select {
// // 		case msg := <-liveConnection.Drivers():
// // 			dataLock.Lock()
// // 			var numbers []int
// // 			for _, driver := range msg.Drivers {
// // 				numbers = append(numbers, driver.Number)
// // 			}
// // 			liveConnection.SelectTelemetrySources(numbers)
// // 			// fmt.Println("msg", msg)
// // 			err := websocket.JSON.Send(ws, &DataStruct{
// // 				DataType: "DRIVERS",
// // 				Data:     msg,
// // 			})
// // 			if err != nil {
// // 				c.Logger().Error(err)
// // 			}
// // 			dataLock.Unlock()
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessDrivers(msg)
// // 			// }

// // 		case msg := <-liveConnection.Timing():
// // 			dataLock.Lock()
// // 			err := websocket.JSON.Send(ws, &DataStruct{
// // 				DataType: "TIMING",
// // 				Data:     msg,
// // 			})
// // 			if err != nil {
// // 				c.Logger().Error(err)
// // 			}
// // 			dataLock.Unlock()
// // 			// TODO - sometimes get empty records on shutdown so filter these out
// // 			// if msg.Position == 0 {
// // 			// 	continue
// // 			// }

// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessTiming(msg)
// // 			// }

// // 		case msg := <-liveConnection.Event():
// // 			dataLock.Lock()
// // 			err := websocket.JSON.Send(ws, &DataStruct{
// // 				DataType: "EVENT",
// // 				Data:     msg,
// // 			})
// // 			if err != nil {
// // 				c.Logger().Error(err)
// // 			}
// // 			dataLock.Unlock()

// // 		case <-liveConnection.Time():
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessEventTime(msg)
// // 			// }

// // 		case <-liveConnection.RaceControlMessages():
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessRaceControlMessages(msg)
// // 			// }

// // 		case <-liveConnection.Weather():
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessWeather(msg)
// // 			// }

// // 		case <-liveConnection.Radio():
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessRadio(msg)
// // 			// }

// // 		case <-liveConnection.Location():
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessLocation(msg)
// // 			// }

// // 		case msg := <-liveConnection.Telemetry():
// // 			dataLock.Lock()
// // 			// fmt.Println("msg", msg)
// // 			err := websocket.JSON.Send(ws, &DataStruct{
// // 				DataType: "TELEMETRY",
// // 				Data:     msg,
// // 			})
// // 			if err != nil {
// // 				c.Logger().Error(err)
// // 			}
// // 			dataLock.Unlock()
// // 			// for x := range d.panels {
// // 			// 	d.panels[x].ProcessTelemetry(msg)
// // 			// }
// // 		}

// // 		// Data has changed so force a UI redraw
// // 		// giu.Update()
// // 	}
// // }
