package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	socket "github.com/victorgomez09/f1-tracker.git/internal/socket"
)

var (
	upgrader = websocket.Upgrader{}
)

func handleWebsocket(c echo.Context) error {
	var origins = []string{"http://127.0.0.1:18081", "http://localhost:18081", "https://5173-victorgomez09-f1tracker-a7zkp3fi4k8.ws-eu118.gitpod.io"}
	// ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	var up = websocket.Upgrader{
		// Resolve cross-domain problems
		CheckOrigin: func(r *http.Request) bool {
			var origin = r.Header.Get("origin")
			fmt.Println("origin", origin)
			for _, allowOrigin := range origins {
				if origin == allowOrigin {
					return true
				}
			}
			return false
		}}
	ws, err := up.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	socket.InitStream(ws)

	// for {
	// 	// Write
	// 	err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}

	// 	// Read
	// 	_, msg, err := ws.ReadMessage()
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}
	// 	fmt.Printf("%s\n", msg)
	// }

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Static("/", "../public")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://5173-victorgomez09-f1tracker-a7zkp3fi4k8.ws-eu118.gitpod.io"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/ws", handleWebsocket)
	e.Logger.Fatal(e.Start(":3000"))
}
