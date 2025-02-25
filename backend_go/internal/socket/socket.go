package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/victorgomez09/f1-tracker.git/internal/model"
	"github.com/victorgomez09/f1-tracker.git/utils"
)

type NegotiationResponse struct {
	Url                     string
	ConnectionToken         string
	ConnectionId            string
	KeepAliveTimeout        float32
	DisconnectTimeout       float32
	ConnectionTimeout       float32
	TryWebSockets           bool
	ProtocolVersion         float32
	TransportConnectTimeout float32
	LongPollDelay           float32
}

var state model.F1State

func InitStream(ws *websocket.Conn) {
	log.Println(utils.SignalrUrl, " Connecting to live timing stream")

	urlParsed, err := url.Parse("https://" + utils.SignalrUrl + "/negotiate")
	if err != nil {
		log.Println("Error parsing url")
		log.Fatal("err", err)
	}
	urlParsedValues := urlParsed.Query()
	urlParsedValues.Add("connectionData", utils.SignalrHubParsed)
	urlParsedValues.Add("clientProtocol", "1.5")
	newQuery := *urlParsed
	newQuery.RawQuery = urlParsedValues.Encode()
	response, err := http.Get(newQuery.String())
	if err != nil {
		log.Fatal("error negotiating", err)
	}

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("error parsing body", err)
		}

		cookie := response.Header.Get("Set-cookie")
		var response NegotiationResponse
		json.Unmarshal(body, &response)

		u := url.URL{Scheme: "wss", Host: utils.WssUrl, Path: "/signalr/connect"}
		parsedUrl := u
		values := u.Query()
		values.Add("clientProtocol", "1.5")
		values.Add("transport", "webSockets")
		values.Add("connectionToken", response.ConnectionToken)
		values.Add("connectionData", utils.SignalrHubParsed)
		parsedUrl.RawQuery = values.Encode()
		log.Printf("connecting to %s", parsedUrl.String())

		c, resp, err := websocket.DefaultDialer.Dial(parsedUrl.String(), http.Header{
			"User-Agent":      []string{"BestHTTP"},
			"Accept-Encoding": []string{"gzip,identity"},
			"Cookie":          []string{cookie},
		})
		// fmt.Println("resp", resp)

		if err != nil {
			log.Printf("handshake failed with status %d", resp.StatusCode)
			log.Fatal("dial:", err)
		}

		if err := c.WriteMessage(websocket.TextMessage, []byte(utils.SignalrSubscribe)); err != nil {
			log.Printf("error sending message")
			log.Fatal("err:", err)
		}

		for {
			// Read message from browser
			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read error:", err)
				break
			}
			fmt.Printf("Received: %s\n", msg)

			// Write message back to browser
			// if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
			// 	fmt.Println("write error:", err)
			// 	break
			// }
			// var object = &model.SocketData{}
			// errParsing := json.Unmarshal(msg, object)
			// if errParsing != nil {
			// 	log.Fatal("Error parsing JSON", errParsing)
			// }
			// fmt.Println("onbject", *object)

			// UpdateState(state, *object)
			if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("write error:", err)
			}
		}

		fmt.Println("subscribed")

		//When the program closes close the connection
		defer c.Close()
	}
}
