// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneHoverEvent : QGraphicsSceneHoverEvent
type QGraphicsSceneHoverEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneHoverEvent::QGraphicsSceneHoverEvent()	
func NewGraphicsSceneHoverEvent() *QGraphicsSceneHoverEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),277000,277102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneHoverEvent{}
	_p.SetDriver(__rv,277000,true)
	return _p
} 
//QGraphicsSceneHoverEvent::QGraphicsSceneHoverEvent(QEvent::Type)	
func NewGraphicsSceneHoverEventWithType(_type QEvent_Type) *QGraphicsSceneHoverEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),277000,277103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneHoverEvent{}
	_p.SetDriver(__rv,277000,true)
	return _p
} 
//QGraphicsSceneHoverEvent::lastPos()
func (q *QGraphicsSceneHoverEvent) LastPos() *QPointF {
	var __rv uintptr
	q.Drv(277000,277104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::lastScenePos()
func (q *QGraphicsSceneHoverEvent) LastScenePos() *QPointF {
	var __rv uintptr
	q.Drv(277000,277105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::lastScreenPos()
func (q *QGraphicsSceneHoverEvent) LastScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(277000,277106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::modifiers()
func (q *QGraphicsSceneHoverEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(277000,277107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneHoverEvent::pos()
func (q *QGraphicsSceneHoverEvent) Pos() *QPointF {
	var __rv uintptr
	q.Drv(277000,277108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::scenePos()
func (q *QGraphicsSceneHoverEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(277000,277109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::screenPos()
func (q *QGraphicsSceneHoverEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(277000,277110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneHoverEvent::setLastPos(QPointF const&)
func (q *QGraphicsSceneHoverEvent) SetLastPos(pos *QPointF)  {
	q.Drv(277000,277111,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setLastScenePos(QPointF const&)
func (q *QGraphicsSceneHoverEvent) SetLastScenePos(pos *QPointF)  {
	q.Drv(277000,277112,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setLastScreenPos(QPoint const&)
func (q *QGraphicsSceneHoverEvent) SetLastScreenPos(pos *QPoint)  {
	q.Drv(277000,277113,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QGraphicsSceneHoverEvent) SetModifiers(modifiers Qt_KeyboardModifier)  {
	q.Drv(277000,277114,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setPos(QPointF const&)
func (q *QGraphicsSceneHoverEvent) SetPos(pos *QPointF)  {
	q.Drv(277000,277115,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneHoverEvent) SetScenePos(pos *QPointF)  {
	q.Drv(277000,277116,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHoverEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneHoverEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(277000,277117,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
