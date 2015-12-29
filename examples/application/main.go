package main

import (
	"os"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(os.Args, func() {
		app := ui.Application()
		app.SetOrganizationName("GoQt")
		app.SetApplicationName("Application Example")

		w := NewMainWindow()
		w.Show()
	})
}
