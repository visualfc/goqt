package main

import "github.com/visualfc/goqt/ui"

func main() {
	ui.Run(func() {
		widget := ui.NewWidget()
		widget.Show()
	})
}
