package socket

import (
	"github.com/dannyjhall/gotyper/typer"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("speed")
	speed := 3
	if i, err := strconv.Atoi(param); err != nil {
		log.Println(err, " - Using default type speed")
	} else {
		speed = i
	}
	goTyper := typer.Typer{
		Index: 0,
		Speed: speed,
		Out:   make(chan []byte),
	}
	go goTyper.LoadSource()

	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}

	for {
		if err := conn.WriteMessage(websocket.TextMessage, <-goTyper.Out); err != nil {
			log.Println(err)
		}
		// Not responding to actual key codes yet, just the key press
		if _, _, err := conn.ReadMessage(); err != nil {
			log.Println(err)
			break
		} else {
			go goTyper.Type()
		}
	}

}
