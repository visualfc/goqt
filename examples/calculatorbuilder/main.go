package main

import (
	"log"
	"os"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(os.Args, func() {
		w, err := NewCalculatorForm()
		if err != nil {
			log.Fatalln(err)
		}
		w.Show()
	})
}
