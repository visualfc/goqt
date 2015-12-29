package main

import (
	"log"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(func() {
		w, err := NewCalclatorForm()
		if err != nil {
			log.Fatalln(err)
		}
		w.Show()
	})
}
