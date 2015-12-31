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
	btn := ui.NewQPushButton()
	btn.SetText("Async")

	clear := ui.NewQPushButton()
	clear.SetText("Clear")

	edit := ui.NewQPlainTextEdit()
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

	hbox := ui.NewQHBoxLayout()
	hbox.AddWidget(btn)
	hbox.AddWidget(clear)

	vbox := ui.NewQVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddWidget(edit)

	widget := ui.NewQWidget()
	widget.SetLayout(vbox)
	widget.Show()
}
