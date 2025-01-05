package main

import (
	"net/http"

	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	pinPomp := rpio.Pin(10)
	pinPomp.Output()

	pinValeve := rpio.Pin(22)
	pinValeve.Output()

	pinLevel1 := rpio.Pin(2)
	pinLevel1.Input()

	// Handler dla WebSocket
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("Błąd upgradera WebSocket: %s\n", err)
			return
		}
		defer conn.Close()

		for {
			// Sprawdź stan pinu
			if pinLevel1.Read() == rpio.High {
				conn.WriteMessage(websocket.TextMessage, []byte("onLevel1"))
			} else {
				conn.WriteMessage(websocket.TextMessage, []byte("offLevel1"))
			}
			time.Sleep(100 * time.Millisecond)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "strona.html")
	})

	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		pinPomp.High()
	})

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		pinPomp.Low()
	})

	http.HandleFunc("/on1", func(w http.ResponseWriter, r *http.Request) {
		pinValeve.High()
	})

	http.HandleFunc("/off1", func(w http.ResponseWriter, r *http.Request) {
		pinValeve.Low()
	})

	err = http.ListenAndServe(":2137", nil)
	if err != nil {
		panic(err)
	}

}
