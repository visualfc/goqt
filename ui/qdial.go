// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDial : QDial
type QDial struct {
	QAbstractSlider
}
//QDial::QDial()	
func NewDial() *QDial {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),225000,225102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDial{}
	_p.SetDriver(__rv,225000,false)
	return _p
} 
//QDial::QDial(QWidget*)	
func NewDialWithParent(parent QWidgetInterface) *QDial {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),225000,225103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDial{}
	_p.SetDriver(__rv,225000,false)
	return _p
} 
//QDial::event(QEvent*)
func (q *QDial) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(225000,225104,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDial::minimumSizeHint()
func (q *QDial) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(225000,225105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QDial::mouseMoveEvent(QMouseEvent*)
func (q *QDial) MouseMoveEvent(me *QMouseEvent)  {
	q.Drv(225000,225106,Native(me),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::mousePressEvent(QMouseEvent*)
func (q *QDial) MousePressEvent(me *QMouseEvent)  {
	q.Drv(225000,225107,Native(me),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::mouseReleaseEvent(QMouseEvent*)
func (q *QDial) MouseReleaseEvent(me *QMouseEvent)  {
	q.Drv(225000,225108,Native(me),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::notchSize()
func (q *QDial) NotchSize() int {
	var __rv int
	q.Drv(225000,225109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDial::notchTarget()
func (q *QDial) NotchTarget() float64 {
	var __rv float64
	q.Drv(225000,225110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDial::notchesVisible()
func (q *QDial) NotchesVisible() bool {
	var __rv bool
	q.Drv(225000,225111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDial::paintEvent(QPaintEvent*)
func (q *QDial) PaintEvent(pe *QPaintEvent)  {
	q.Drv(225000,225112,Native(pe),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::resizeEvent(QResizeEvent*)
func (q *QDial) ResizeEvent(re *QResizeEvent)  {
	q.Drv(225000,225113,Native(re),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::setNotchTarget(double)
func (q *QDial) SetNotchTarget(target float64)  {
	q.Drv(225000,225114,unsafe.Pointer(&target),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::setNotchesVisible(bool)
func (q *QDial) SetNotchesVisible(visible bool)  {
	q.Drv(225000,225115,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::setWrapping(bool)
func (q *QDial) SetWrapping(on bool)  {
	q.Drv(225000,225116,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::sizeHint()
func (q *QDial) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(225000,225117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QDial::sliderChange(QAbstractSlider::SliderChange)
func (q *QDial) SliderChange(change QAbstractSlider_SliderChange)  {
	q.Drv(225000,225118,unsafe.Pointer(&change),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDial::wrapping()
func (q *QDial) Wrapping() bool {
	var __rv bool
	q.Drv(225000,225119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
