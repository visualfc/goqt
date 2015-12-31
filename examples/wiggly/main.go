package main

import (
	"os"

	"github.com/visualfc/goqt/ui"
)

type WigglyWidget struct {
	*ui.QWidget
	step int
	text string
	exit chan bool
}

func NewWiggleWidget() *WigglyWidget {
	w := &WigglyWidget{ui.NewQWidget(), 0, "Hello World", make(chan bool)}
	w.SetBackgroundRole(ui.QPalette_Midlight)
	w.SetAutoFillBackground(true)

	font := w.Font()
	font.SetPointSize(font.PointSize() + 20)
	w.SetFont(font)

	w.InstallEventFilter(w)

	t := ui.NewQTimer()
	t.SetInterval(60)
	t.OnTimeout(func() {
		w.step++
		w.Update()
	})
	t.Start()

	return w
}

func (w *WigglyWidget) SetText(text string) {
	w.text = text
}

var sineTablep = [16]int{0, 38, 71, 92, 100, 92, 71, 38, 0, -38, -71, -92, -100, -92, -71, -38}

func (w *WigglyWidget) OnPaintEvent(e *ui.QPaintEvent) bool {
	w.PaintEvent(e)
	merics := ui.NewQFontMetrics(w.Font())
	defer merics.Delete()

	x := (w.Width() - merics.WidthWithString(w.text)) / 2
	y := (w.Height() + merics.Ascent() - merics.Descent()) / 2

	color := ui.NewQColor()
	defer color.Delete()

	painter := ui.NewQPainterWithPaintDevice(w)
	defer painter.Delete()

	for i, s := range w.text {
		index := (w.step + i) % 16
		color.SetHsv((15-index)*16, 255, 191, 255)
		painter.SetPen(color)
		painter.DrawTextWithXYText(x, y-((sineTablep[index]*merics.Height())/400), string(s))
		x += merics.Width(s)
	}
	return true
}

type Dialog struct {
	*ui.QDialog
	widdly *WigglyWidget
}

func NewDialog() *Dialog {
	dlg := &Dialog{}
	dlg.QDialog = ui.NewQDialog()

	vbox := ui.NewQVBoxLayout()
	dlg.SetLayout(vbox)

	dlg.widdly = NewWiggleWidget()
	edit := ui.NewQLineEdit()

	vbox.AddWidget(dlg.widdly)
	vbox.AddWidget(edit)

	edit.OnTextChanged(func(text string) {
		dlg.widdly.SetText(text)
	})
	edit.SetText("Hello World!")

	dlg.ResizeWithWidthHeight(360, 145)
	dlg.SetWindowTitle("Wiggly")

	return dlg
}

func main() {
	ui.RunEx(os.Args, func() {
		dlg := NewDialog()
		dlg.Show()
	})
}
