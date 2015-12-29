// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGestureEvent : QGestureEvent
type QGestureEvent struct {
	QEvent
}
//QGestureEvent::QGestureEvent(QList<QGesture*> const&)	
func NewGestureEvent(gestures []*QGesture) *QGestureEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),44000,44102,unsafe.Pointer(&gestures),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGestureEvent{}
	_p.SetDriver(__rv,44000,true)
	return _p
} 
//QGestureEvent::accept()
func (q *QGestureEvent) Accept()  {
	q.Drv(44000,44103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::accept(QGesture*)
func (q *QGestureEvent) AcceptWithGesture(value *QGesture)  {
	q.Drv(44000,44104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::accept(Qt::GestureType)
func (q *QGestureEvent) AcceptWithGesturetype(value Qt_GestureType)  {
	q.Drv(44000,44105,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::activeGestures()
func (q *QGestureEvent) ActiveGestures() []*QGesture {
	var __rv []*QGesture
	q.Drv(44000,44106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::canceledGestures()
func (q *QGestureEvent) CanceledGestures() []*QGesture {
	var __rv []*QGesture
	q.Drv(44000,44107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::gesture(Qt::GestureType)
func (q *QGestureEvent) Gesture(_type Qt_GestureType) *QGesture {
	var __rv uintptr
	q.Drv(44000,44108,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGesture{}
	_rp.SetDriver(__rv,247000,false)
	return _rp
}	
//QGestureEvent::gestures()
func (q *QGestureEvent) Gestures() []*QGesture {
	var __rv []*QGesture
	q.Drv(44000,44109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::ignore()
func (q *QGestureEvent) Ignore()  {
	q.Drv(44000,44110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::ignore(QGesture*)
func (q *QGestureEvent) IgnoreWithGesture(value *QGesture)  {
	q.Drv(44000,44111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::ignore(Qt::GestureType)
func (q *QGestureEvent) IgnoreWithGesturetype(value Qt_GestureType)  {
	q.Drv(44000,44112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::isAccepted()
func (q *QGestureEvent) IsAccepted() bool {
	var __rv bool
	q.Drv(44000,44113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::isAccepted(QGesture*)
func (q *QGestureEvent) IsAcceptedWithGesture(value *QGesture) bool {
	var __rv bool
	q.Drv(44000,44114,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::isAccepted(Qt::GestureType)
func (q *QGestureEvent) IsAcceptedWithGesturetype(value Qt_GestureType) bool {
	var __rv bool
	q.Drv(44000,44115,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureEvent::mapToGraphicsScene(QPointF const&)
func (q *QGestureEvent) MapToGraphicsScene(gesturePoint *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(44000,44116,Native(gesturePoint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGestureEvent::setAccepted(bool)
func (q *QGestureEvent) SetAccepted(accepted bool)  {
	q.Drv(44000,44117,unsafe.Pointer(&accepted),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::setAccepted(QGesture*,bool)
func (q *QGestureEvent) SetAcceptedWithGestureBool(value2 *QGesture,value3 bool)  {
	q.Drv(44000,44118,Native(value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::setAccepted(Qt::GestureType,bool)
func (q *QGestureEvent) SetAcceptedWithGesturetypeBool(value2 Qt_GestureType,value3 bool)  {
	q.Drv(44000,44119,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::setWidget(QWidget*)
func (q *QGestureEvent) SetWidget(widget QWidgetInterface)  {
	q.Drv(44000,44120,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureEvent::widget()
func (q *QGestureEvent) Widget() *QWidget {
	var __rv uintptr
	q.Drv(44000,44121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
