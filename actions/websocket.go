package actions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
)

var websocketConn *websocket.Conn

// WebSocketHandler - Handle websocket request
func WebSocketHandler(c buffalo.Context) error {
	var err error
	websocketConn, err = GetWebsocketConnection(c)
	if err != nil {
		return nil
	}
	return nil
}

// GetWebsocketConnection - Returns a connection to the websocket.
func GetWebsocketConnection(c buffalo.Context) (*websocket.Conn, error) {
	r := c.Request()
	w := c.Response()
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("conn")
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func onMessage() {
	for {
		messageType, p, err := websocketConn.ReadMessage()
		if err != nil {
			fmt.Println("close")
			log.Println(err)
		}
		if messageType == 1 {
			fmt.Println("read : ", p)
			go onSend("shit")
		}
	}
}

func onSend(message string) {
	if err := websocketConn.WriteMessage(1, []byte(message)); err != nil {
		fmt.Println("write")
		log.Println(err)
	}
}
