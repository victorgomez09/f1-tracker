package internal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
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

func InitStream() {
	log.Println(utils.SignalrUrl, " Connecting to live timing stream")

	response, err := http.Get("https://" + utils.SignalrUrl + "/negotiate?connectionData=" + utils.SignalrHubParsed + "&clientProtocol=1.5")
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

		u := url.URL{Scheme: "wss", Host: utils.WssUrl, Path: "/signalr/connect?clientProtocol=1.5&transport=webSockets&connectionToken=" + url.QueryEscape(response.ConnectionToken) + "&connectionData=" + utils.SignalrHubParsed + ""}
		log.Printf("connecting to %s", u.String())

		c, resp, err := websocket.DefaultDialer.Dial(u.String(), http.Header{
			"User-Agent":      []string{"BestHTTP"},
			"Accept-Encoding": []string{"gzip,identity"},
			"Cookie":          []string{cookie},
		})

		if err != nil {
			log.Printf("handshake failed with status %d", resp.StatusCode)
			log.Fatal("dial:", err)
		}
		//When the program closes close the connection
		defer c.Close()
	}
}
