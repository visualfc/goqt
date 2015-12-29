package main

import (
	"os"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(os.Args, func() {
		ed := NewCodeEdit()
		ed.ResizeWithWidthHeight(400, 400)
		ed.Show()
	})
}
