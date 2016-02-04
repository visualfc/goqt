package main

import (
	"fmt"
	"runtime"

	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.Run(func() {
		info := fmt.Sprintf("GoQt Version %v \nQt Version %v\ngo verison %v %v/%v",
			ui.Version(),
			ui.QtVersion(),
			runtime.Version(), runtime.GOOS, runtime.GOARCH)

		lable := ui.NewLabel()
		lable.SetText(info)

		hbox := ui.NewHBoxLayout()
		hbox.AddWidget(lable)

		widget := ui.NewWidget()
		widget.SetLayout(hbox)
		widget.Show()
	})
}
