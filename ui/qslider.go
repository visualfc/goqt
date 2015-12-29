// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSlider_TickPosition - QSlider::TickPosition
type QSlider_TickPosition uint32
const (
	QSlider_NoTicks QSlider_TickPosition = 0
	QSlider_TicksAbove QSlider_TickPosition = 1
	QSlider_TicksLeft QSlider_TickPosition = QSlider_TicksAbove
	QSlider_TicksBelow QSlider_TickPosition = 2
	QSlider_TicksRight QSlider_TickPosition = QSlider_TicksBelow
	QSlider_TicksBothSides QSlider_TickPosition = 3
)
//struct QSlider : QSlider
type QSlider struct {
	QAbstractSlider
}
//QSlider::QSlider()	
func NewSlider() *QSlider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),344000,344102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSlider{}
	_p.SetDriver(__rv,344000,false)
	return _p
} 
//QSlider::QSlider(QWidget*)	
func NewSliderWithParent(parent QWidgetInterface) *QSlider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),344000,344103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSlider{}
	_p.SetDriver(__rv,344000,false)
	return _p
} 
//QSlider::QSlider(Qt::Orientation,QWidget*)	
func NewSliderWithOrientationParent(orientation Qt_Orientation,parent QWidgetInterface) *QSlider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),344000,344104,unsafe.Pointer(&orientation),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSlider{}
	_p.SetDriver(__rv,344000,false)
	return _p
} 
//QSlider::event(QEvent*)
func (q *QSlider) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(344000,344105,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSlider::minimumSizeHint()
func (q *QSlider) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(344000,344106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSlider::mouseMoveEvent(QMouseEvent*)
func (q *QSlider) MouseMoveEvent(ev *QMouseEvent)  {
	q.Drv(344000,344107,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::mousePressEvent(QMouseEvent*)
func (q *QSlider) MousePressEvent(ev *QMouseEvent)  {
	q.Drv(344000,344108,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::mouseReleaseEvent(QMouseEvent*)
func (q *QSlider) MouseReleaseEvent(ev *QMouseEvent)  {
	q.Drv(344000,344109,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::paintEvent(QPaintEvent*)
func (q *QSlider) PaintEvent(ev *QPaintEvent)  {
	q.Drv(344000,344110,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::setTickInterval(int)
func (q *QSlider) SetTickInterval(ti int)  {
	q.Drv(344000,344111,unsafe.Pointer(&ti),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::setTickPosition(QSlider::TickPosition)
func (q *QSlider) SetTickPosition(position QSlider_TickPosition)  {
	q.Drv(344000,344112,unsafe.Pointer(&position),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSlider::sizeHint()
func (q *QSlider) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(344000,344113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSlider::tickInterval()
func (q *QSlider) TickInterval() int {
	var __rv int
	q.Drv(344000,344114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSlider::tickPosition()
func (q *QSlider) TickPosition() QSlider_TickPosition {
	var __rv QSlider_TickPosition
	q.Drv(344000,344115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
