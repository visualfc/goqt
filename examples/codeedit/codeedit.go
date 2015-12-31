package main

import (
	"strconv"
	"strings"

	"github.com/visualfc/goqt/ui"
)

type CodeEdit struct {
	*ui.QWidget
	edit     *ui.QPlainTextEdit
	lineArea *ui.QWidget
	syntax   *ui.QSyntaxHighlighterHook
	rules    map[*ui.QRegExp]*ui.QTextCharFormat
}

func NewCodeEdit() *CodeEdit {
	w := &CodeEdit{}
	w.QWidget = ui.NewQWidget()
	w.edit = ui.NewQPlainTextEdit()
	w.edit.SetLineWrapMode(ui.QPlainTextEdit_NoWrap)
	w.syntax = ui.NewQSyntaxHighlighterHookWithDoc(w.edit.Document())
	w.lineArea = ui.NewQWidget()
	w.lineArea.SetFixedWidth(0)

	hbox := ui.NewQHBoxLayout()
	hbox.SetMargin(0)
	hbox.SetSpacing(0)
	hbox.AddWidget(w.lineArea)
	hbox.AddWidget(w.edit)
	w.SetLayout(hbox)

	w.lineArea.InstallEventFilter(w)

	w.edit.OnBlockCountChanged(func(int) {
		w.UpdateLineNumberAreaWidth()
		w.lineArea.Update()
	})
	w.edit.OnUpdateRequest(func(*ui.QRect, int) {
		w.lineArea.Update()
	})

	w.UpdateLineNumberAreaWidth()

	w.MakeRules()

	w.syntax.OnHighlightBlock(w.OnHighlightBlock)

	font := w.edit.Font()
	font.SetPointSize(12)
	w.edit.SetFont(font)
	w.lineArea.SetFont(font)

	return w
}

var keywords = `
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var 
`

func (w *CodeEdit) MakeRules() {
	w.rules = make(map[*ui.QRegExp]*ui.QTextCharFormat)
	keyword := ui.NewQTextCharFormat()
	b := ui.NewQBrush()
	b.SetStyle(ui.Qt_SolidPattern)
	b.SetColorWithGlobalcolor(ui.Qt_darkBlue)
	keyword.SetForeground(b)
	keyword.SetFontWeight(int(ui.QFont_Bold))
	for _, line := range strings.Split(keywords, "\n") {
		for _, v := range strings.Split(line, " ") {
			if len(v) == 0 {
				continue
			}
			r := ui.NewQRegExp()
			r.SetPattern("\\b" + v + "\\b")
			w.rules[r] = keyword
		}
	}
}

func (w *CodeEdit) OnHighlightBlock(text string) {
	for k, v := range w.rules {
		index := k.IndexIn(text)
		for index >= 0 {
			length := k.MatchedLength()
			w.syntax.SetFormatWithStartCountFormat(index, length, v)
			index = k.IndexInWithTextOffsetCaretmode(text, index+length, ui.QRegExp_CaretAtOffset)
		}
	}
}

func (w *CodeEdit) UpdateLineNumberAreaWidth() {
	digits := 1
	max := 1
	count := w.edit.BlockCount()
	if count > max {
		max = count
	}
	for max >= 10 {
		max /= 10
		digits++
	}
	space := 10 + w.edit.FontMetrics().Width('9')*digits
	w.lineArea.SetFixedWidth(space)
}

func (w *CodeEdit) OnPaintEvent(obj *ui.QObject, e *ui.QPaintEvent) bool {
	if ui.Equal(obj, w.lineArea) {
		w.paintLineArea(e)
	}
	return true
}

func (w *CodeEdit) paintLineArea(event *ui.QPaintEvent) {
	painter := ui.NewQPainterWithPaintDevice(w.lineArea)
	defer painter.Delete()
	painter.FillRectWithRectColor(w.lineArea.Rect(), ui.NewQColorWithGlobalcolor(ui.Qt_lightGray))

	block := w.edit.FirstVisibleBlock()
	blockNumber := block.BlockNumber() + 1
	top := int(w.edit.BlockBoundingGeometry(block).Translated(w.edit.ContentOffset()).Top())
	bottom := top + int(w.edit.BlockBoundingRect(block).Height())
	height := w.lineArea.FontMetrics().Height()
	for block.IsValid() && top < event.Rect().Bottom() {
		if block.IsVisible() && bottom >= event.Rect().Top() {
			painter.SetPen(ui.NewQColorWithGlobalcolor(ui.Qt_black))
			painter.DrawTextWithXYWidthHeightFlagsTextRect(0, top, w.lineArea.Width(), height, int(ui.Qt_AlignRight), strconv.Itoa(blockNumber), ui.NewQRect())
		}
		block = block.Next()
		top = bottom
		bottom = top + int(w.edit.BlockBoundingRect(block).Height())
		blockNumber++
	}
}
