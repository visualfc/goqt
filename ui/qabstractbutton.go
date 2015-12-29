// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractButton : QAbstractButton
type QAbstractButton struct {
	QWidget
}
func (q *QAbstractButton) OnClicked(fn func()) uintptr {
	var __rv uintptr
	q.Drv(193000,193102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractButton) OnClickedEx(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(193000,193103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractButton) OnPressed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(193000,193104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractButton) OnReleased(fn func()) uintptr {
	var __rv uintptr
	q.Drv(193000,193105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractButton) OnToggled(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(193000,193106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractButton::animateClick()
func (q *QAbstractButton) AnimateClick()  {
	q.Drv(193000,193107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::animateClick(int)
func (q *QAbstractButton) AnimateClickWithMsec(msec int)  {
	q.Drv(193000,193108,unsafe.Pointer(&msec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::autoExclusive()
func (q *QAbstractButton) AutoExclusive() bool {
	var __rv bool
	q.Drv(193000,193109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::autoRepeat()
func (q *QAbstractButton) AutoRepeat() bool {
	var __rv bool
	q.Drv(193000,193110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::autoRepeatDelay()
func (q *QAbstractButton) AutoRepeatDelay() int {
	var __rv int
	q.Drv(193000,193111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::autoRepeatInterval()
func (q *QAbstractButton) AutoRepeatInterval() int {
	var __rv int
	q.Drv(193000,193112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::changeEvent(QEvent*)
func (q *QAbstractButton) ChangeEvent(e *QEvent)  {
	q.Drv(193000,193113,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::checkStateSet()
func (q *QAbstractButton) CheckStateSet()  {
	q.Drv(193000,193114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::click()
func (q *QAbstractButton) Click()  {
	q.Drv(193000,193115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::event(QEvent*)
func (q *QAbstractButton) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(193000,193116,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::focusInEvent(QFocusEvent*)
func (q *QAbstractButton) FocusInEvent(e *QFocusEvent)  {
	q.Drv(193000,193117,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::focusOutEvent(QFocusEvent*)
func (q *QAbstractButton) FocusOutEvent(e *QFocusEvent)  {
	q.Drv(193000,193118,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::group()
func (q *QAbstractButton) Group() *QButtonGroup {
	var __rv uintptr
	q.Drv(193000,193119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QButtonGroup{}
	_rp.SetDriver(__rv,211000,false)
	return _rp
}	
//QAbstractButton::hitButton(QPoint const&)
func (q *QAbstractButton) HitButton(pos *QPoint) bool {
	var __rv bool
	q.Drv(193000,193120,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::icon()
func (q *QAbstractButton) Icon() *QIcon {
	var __rv uintptr
	q.Drv(193000,193121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QAbstractButton::iconSize()
func (q *QAbstractButton) IconSize() *QSize {
	var __rv uintptr
	q.Drv(193000,193122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractButton::isCheckable()
func (q *QAbstractButton) IsCheckable() bool {
	var __rv bool
	q.Drv(193000,193123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::isChecked()
func (q *QAbstractButton) IsChecked() bool {
	var __rv bool
	q.Drv(193000,193124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::isDown()
func (q *QAbstractButton) IsDown() bool {
	var __rv bool
	q.Drv(193000,193125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::keyPressEvent(QKeyEvent*)
func (q *QAbstractButton) KeyPressEvent(e *QKeyEvent)  {
	q.Drv(193000,193126,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::keyReleaseEvent(QKeyEvent*)
func (q *QAbstractButton) KeyReleaseEvent(e *QKeyEvent)  {
	q.Drv(193000,193127,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::mouseMoveEvent(QMouseEvent*)
func (q *QAbstractButton) MouseMoveEvent(e *QMouseEvent)  {
	q.Drv(193000,193128,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::mousePressEvent(QMouseEvent*)
func (q *QAbstractButton) MousePressEvent(e *QMouseEvent)  {
	q.Drv(193000,193129,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::mouseReleaseEvent(QMouseEvent*)
func (q *QAbstractButton) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(193000,193130,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::nextCheckState()
func (q *QAbstractButton) NextCheckState()  {
	q.Drv(193000,193131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::paintEvent(QPaintEvent*)
func (q *QAbstractButton) PaintEvent(e *QPaintEvent)  {
	q.Drv(193000,193132,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setAutoExclusive(bool)
func (q *QAbstractButton) SetAutoExclusive(value bool)  {
	q.Drv(193000,193133,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setAutoRepeat(bool)
func (q *QAbstractButton) SetAutoRepeat(value bool)  {
	q.Drv(193000,193134,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setAutoRepeatDelay(int)
func (q *QAbstractButton) SetAutoRepeatDelay(value int)  {
	q.Drv(193000,193135,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setAutoRepeatInterval(int)
func (q *QAbstractButton) SetAutoRepeatInterval(value int)  {
	q.Drv(193000,193136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setCheckable(bool)
func (q *QAbstractButton) SetCheckable(value bool)  {
	q.Drv(193000,193137,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setChecked(bool)
func (q *QAbstractButton) SetChecked(value bool)  {
	q.Drv(193000,193138,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setDown(bool)
func (q *QAbstractButton) SetDown(value bool)  {
	q.Drv(193000,193139,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setIcon(QIcon const&)
func (q *QAbstractButton) SetIcon(icon *QIcon)  {
	q.Drv(193000,193140,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setIconSize(QSize const&)
func (q *QAbstractButton) SetIconSize(size *QSize)  {
	q.Drv(193000,193141,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setShortcut(QKeySequence const&)
func (q *QAbstractButton) SetShortcut(key *QKeySequence)  {
	q.Drv(193000,193142,Native(key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::setText(QString const&)
func (q *QAbstractButton) SetText(text string)  {
	q.Drv(193000,193143,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::shortcut()
func (q *QAbstractButton) Shortcut() *QKeySequence {
	var __rv uintptr
	q.Drv(193000,193144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QAbstractButton::text()
func (q *QAbstractButton) Text() string {
	var __rv string
	q.Drv(193000,193145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractButton::timerEvent(QTimerEvent*)
func (q *QAbstractButton) TimerEvent(e *QTimerEvent)  {
	q.Drv(193000,193146,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractButton::toggle()
func (q *QAbstractButton) Toggle()  {
	q.Drv(193000,193147,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
