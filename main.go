package main

import (
	"html/template"
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

	pinLevel1 := rpio.Pin(4)
	pinLevel1.Input()
	pinLevel1.PullUp()

	tmpl, err := template.ParseFiles("strona.template")
	if err != nil {
		panic(err)
	}

	dane1 := data1{
		Switches: []Switch{
			{
				Checked: false,
			},
			{
				Checked: false,
			},
		},
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
	}
	// dupa := 1

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//sprawdzam stan switchow przed zaladowaniem strony i ustawiam switch na odpowiedni
		if pinPomp.Read() == rpio.High {
			dane1.Switches[0].Checked = true
		} else {
			dane1.Switches[0].Checked = false
		}
		tmpl.Execute(w, dane1)

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
	Circles []Circle
}

type Circle struct {
	Filled bool
}

type data1 struct {
	Switches []Switch
}

type Switch struct {
	Checked bool
}
