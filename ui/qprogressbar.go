// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QProgressBar_Direction - QProgressBar::Direction
type QProgressBar_Direction uint32
const (
	QProgressBar_TopToBottom QProgressBar_Direction = 0
	QProgressBar_BottomToTop QProgressBar_Direction = 1
)
//struct QProgressBar : QProgressBar
type QProgressBar struct {
	QWidget
}
func (q *QProgressBar) OnValueChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(327000,327102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QProgressBar::QProgressBar()	
func NewProgressBar() *QProgressBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),327000,327103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProgressBar{}
	_p.SetDriver(__rv,327000,false)
	return _p
} 
//QProgressBar::QProgressBar(QWidget*)	
func NewProgressBarWithParent(parent QWidgetInterface) *QProgressBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),327000,327104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProgressBar{}
	_p.SetDriver(__rv,327000,false)
	return _p
} 
//QProgressBar::alignment()
func (q *QProgressBar) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(327000,327105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::event(QEvent*)
func (q *QProgressBar) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(327000,327106,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::format()
func (q *QProgressBar) Format() string {
	var __rv string
	q.Drv(327000,327107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::invertedAppearance()
func (q *QProgressBar) InvertedAppearance() bool {
	var __rv bool
	q.Drv(327000,327108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::isTextVisible()
func (q *QProgressBar) IsTextVisible() bool {
	var __rv bool
	q.Drv(327000,327109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::maximum()
func (q *QProgressBar) Maximum() int {
	var __rv int
	q.Drv(327000,327110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::minimum()
func (q *QProgressBar) Minimum() int {
	var __rv int
	q.Drv(327000,327111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::minimumSizeHint()
func (q *QProgressBar) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(327000,327112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QProgressBar::orientation()
func (q *QProgressBar) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(327000,327113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::paintEvent(QPaintEvent*)
func (q *QProgressBar) PaintEvent(value *QPaintEvent)  {
	q.Drv(327000,327114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::reset()
func (q *QProgressBar) Reset()  {
	q.Drv(327000,327115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QProgressBar) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(327000,327116,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setFormat(QString const&)
func (q *QProgressBar) SetFormat(format string)  {
	q.Drv(327000,327117,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setInvertedAppearance(bool)
func (q *QProgressBar) SetInvertedAppearance(invert bool)  {
	q.Drv(327000,327118,unsafe.Pointer(&invert),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setMaximum(int)
func (q *QProgressBar) SetMaximum(maximum int)  {
	q.Drv(327000,327119,unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setMinimum(int)
func (q *QProgressBar) SetMinimum(minimum int)  {
	q.Drv(327000,327120,unsafe.Pointer(&minimum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setOrientation(Qt::Orientation)
func (q *QProgressBar) SetOrientation(value Qt_Orientation)  {
	q.Drv(327000,327121,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setRange(int,int)
func (q *QProgressBar) SetRange(minimum int,maximum int)  {
	q.Drv(327000,327122,unsafe.Pointer(&minimum),unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setTextDirection(QProgressBar::Direction)
func (q *QProgressBar) SetTextDirection(textDirection QProgressBar_Direction)  {
	q.Drv(327000,327123,unsafe.Pointer(&textDirection),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setTextVisible(bool)
func (q *QProgressBar) SetTextVisible(visible bool)  {
	q.Drv(327000,327124,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::setValue(int)
func (q *QProgressBar) SetValue(value int)  {
	q.Drv(327000,327125,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressBar::sizeHint()
func (q *QProgressBar) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(327000,327126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QProgressBar::text()
func (q *QProgressBar) Text() string {
	var __rv string
	q.Drv(327000,327127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::textDirection()
func (q *QProgressBar) TextDirection() QProgressBar_Direction {
	var __rv QProgressBar_Direction
	q.Drv(327000,327128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressBar::value()
func (q *QProgressBar) Value() int {
	var __rv int
	q.Drv(327000,327129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
