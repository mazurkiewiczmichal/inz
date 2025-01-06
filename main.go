package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {

	err := rpio.Open()
	if err != nil {
		if os.Getenv("GOOS") == "linux" {
			panic(err)
		}

	}
	pinPomp := rpio.Pin(10)
	pinPomp.Output()

	pinValeve := rpio.Pin(22)
	pinValeve.Output()

	pinLevel1 := rpio.Pin(4)
	pinLevel1.Input()
	pinLevel1.PullUp()

	tmpl, err := template.ParseFiles("strona.template")
	if err != nil {
		panic(err)
	}

	dane := data{
		Circles: []Circle{
			{
				Filled: false,
			},
			{
				Filled: false,
			},
			{
				Filled: false,
			},
			{
				Filled: false,
			},
		},
		PumpSwitchID:       "pumpSwitch",
		PumpSwitchChecked:  false,
		ValveSwitchID:      "valveSwitch",
		ValveSwitchChecked: false,
	}

	// dupa := 1

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//skorygowanie stanu switcha od zaworu przed zaladowaniem strony
		if pinValeve.Read() == rpio.High {
			dane.ValveSwitchChecked = true
		} else {
			dane.ValveSwitchChecked = false
		}

		//skorygowanie stanu switcha od pompy przed zaladowaniem strony
		if pinPomp.Read() == rpio.High {
			dane.PumpSwitchChecked = true
		} else {
			dane.PumpSwitchChecked = false
		}
		//sprawdzam stan inputow i wypelniam kolka na stronie
		if pinLevel1.Read() == rpio.High {
			// if pinPomp.Read() == 1 {
			dane.Circles[3].Filled = true
		} else {
			dane.Circles[3].Filled = false
		}
		tmpl.Execute(w, dane)

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

type data struct {
	Circles            []Circle
	PumpSwitchID       template.JS
	PumpSwitchChecked  bool
	ValveSwitchID      template.JS
	ValveSwitchChecked bool
}

type Circle struct {
	Filled bool
}
