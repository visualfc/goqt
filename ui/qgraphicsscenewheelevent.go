// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneWheelEvent : QGraphicsSceneWheelEvent
type QGraphicsSceneWheelEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneWheelEvent::QGraphicsSceneWheelEvent()	
func NewGraphicsSceneWheelEvent() *QGraphicsSceneWheelEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),281000,281102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneWheelEvent{}
	_p.SetDriver(__rv,281000,true)
	return _p
} 
//QGraphicsSceneWheelEvent::QGraphicsSceneWheelEvent(QEvent::Type)	
func NewGraphicsSceneWheelEventWithType(_type QEvent_Type) *QGraphicsSceneWheelEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),281000,281103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneWheelEvent{}
	_p.SetDriver(__rv,281000,true)
	return _p
} 
//QGraphicsSceneWheelEvent::buttons()
func (q *QGraphicsSceneWheelEvent) Buttons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(281000,281104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneWheelEvent::delta()
func (q *QGraphicsSceneWheelEvent) Delta() int {
	var __rv int
	q.Drv(281000,281105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneWheelEvent::modifiers()
func (q *QGraphicsSceneWheelEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(281000,281106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneWheelEvent::orientation()
func (q *QGraphicsSceneWheelEvent) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(281000,281107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneWheelEvent::pos()
func (q *QGraphicsSceneWheelEvent) Pos() *QPointF {
	var __rv uintptr
	q.Drv(281000,281108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneWheelEvent::scenePos()
func (q *QGraphicsSceneWheelEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(281000,281109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneWheelEvent::screenPos()
func (q *QGraphicsSceneWheelEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(281000,281110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneWheelEvent::setButtons(QFlags<Qt::MouseButton>)
func (q *QGraphicsSceneWheelEvent) SetButtons(buttons Qt_MouseButton)  {
	q.Drv(281000,281111,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setDelta(int)
func (q *QGraphicsSceneWheelEvent) SetDelta(delta int)  {
	q.Drv(281000,281112,unsafe.Pointer(&delta),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QGraphicsSceneWheelEvent) SetModifiers(modifiers Qt_KeyboardModifier)  {
	q.Drv(281000,281113,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setOrientation(Qt::Orientation)
func (q *QGraphicsSceneWheelEvent) SetOrientation(orientation Qt_Orientation)  {
	q.Drv(281000,281114,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setPos(QPointF const&)
func (q *QGraphicsSceneWheelEvent) SetPos(pos *QPointF)  {
	q.Drv(281000,281115,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneWheelEvent) SetScenePos(pos *QPointF)  {
	q.Drv(281000,281116,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneWheelEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneWheelEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(281000,281117,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
