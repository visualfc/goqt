// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsSceneContextMenuEvent_Reason - QGraphicsSceneContextMenuEvent::Reason
type QGraphicsSceneContextMenuEvent_Reason uint32
const (
	QGraphicsSceneContextMenuEvent_Mouse QGraphicsSceneContextMenuEvent_Reason = 0
	QGraphicsSceneContextMenuEvent_Keyboard QGraphicsSceneContextMenuEvent_Reason = 1
	QGraphicsSceneContextMenuEvent_Other QGraphicsSceneContextMenuEvent_Reason = 2
)
//struct QGraphicsSceneContextMenuEvent : QGraphicsSceneContextMenuEvent
type QGraphicsSceneContextMenuEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneContextMenuEvent::QGraphicsSceneContextMenuEvent()	
func NewGraphicsSceneContextMenuEvent() *QGraphicsSceneContextMenuEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),273000,273102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneContextMenuEvent{}
	_p.SetDriver(__rv,273000,true)
	return _p
} 
//QGraphicsSceneContextMenuEvent::QGraphicsSceneContextMenuEvent(QEvent::Type)	
func NewGraphicsSceneContextMenuEventWithType(_type QEvent_Type) *QGraphicsSceneContextMenuEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),273000,273103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneContextMenuEvent{}
	_p.SetDriver(__rv,273000,true)
	return _p
} 
//QGraphicsSceneContextMenuEvent::modifiers()
func (q *QGraphicsSceneContextMenuEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(273000,273104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneContextMenuEvent::pos()
func (q *QGraphicsSceneContextMenuEvent) Pos() *QPointF {
	var __rv uintptr
	q.Drv(273000,273105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneContextMenuEvent::reason()
func (q *QGraphicsSceneContextMenuEvent) Reason() QGraphicsSceneContextMenuEvent_Reason {
	var __rv QGraphicsSceneContextMenuEvent_Reason
	q.Drv(273000,273106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSceneContextMenuEvent::scenePos()
func (q *QGraphicsSceneContextMenuEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(273000,273107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneContextMenuEvent::screenPos()
func (q *QGraphicsSceneContextMenuEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(273000,273108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneContextMenuEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QGraphicsSceneContextMenuEvent) SetModifiers(modifiers Qt_KeyboardModifier)  {
	q.Drv(273000,273109,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneContextMenuEvent::setPos(QPointF const&)
func (q *QGraphicsSceneContextMenuEvent) SetPos(pos *QPointF)  {
	q.Drv(273000,273110,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneContextMenuEvent::setReason(QGraphicsSceneContextMenuEvent::Reason)
func (q *QGraphicsSceneContextMenuEvent) SetReason(reason QGraphicsSceneContextMenuEvent_Reason)  {
	q.Drv(273000,273111,unsafe.Pointer(&reason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneContextMenuEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneContextMenuEvent) SetScenePos(pos *QPointF)  {
	q.Drv(273000,273112,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneContextMenuEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneContextMenuEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(273000,273113,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
