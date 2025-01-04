package main

import (
	"net/http"

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
