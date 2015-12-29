package main

import (
	"fmt"
	"os"
	"time"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(os.Args, main_ui)
}

func main_ui() {
	btn := ui.NewPushButton()
	btn.SetText("Async")

	clear := ui.NewPushButton()
	clear.SetText("Clear")

	edit := ui.NewPlainTextEdit()
	edit.SetReadOnly(true)

	btn.OnClicked(func() {
		for i := 0; i < 10; i++ {
			go func(i int) {
				start := time.Now()
				<-time.After(1e7)
				offset := time.Now().Sub(start)
				ui.Async(func() {
					edit.AppendPlainText(fmt.Sprintf("Task %v %v", i, offset))
				})
			}(i)
		}
	})

	clear.OnClicked(func() {
		edit.Clear()
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(btn)
	hbox.AddWidget(clear)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddWidget(edit)

	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.Show()
}
