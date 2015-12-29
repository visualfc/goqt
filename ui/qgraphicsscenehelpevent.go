// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneHelpEvent : QGraphicsSceneHelpEvent
type QGraphicsSceneHelpEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneHelpEvent::QGraphicsSceneHelpEvent()	
func NewGraphicsSceneHelpEvent() *QGraphicsSceneHelpEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),276000,276102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneHelpEvent{}
	_p.SetDriver(__rv,276000,true)
	return _p
} 
//QGraphicsSceneHelpEvent::QGraphicsSceneHelpEvent(QEvent::Type)	
func NewGraphicsSceneHelpEventWithType(_type QEvent_Type) *QGraphicsSceneHelpEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),276000,276103,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneHelpEvent{}
	_p.SetDriver(__rv,276000,true)
	return _p
} 
//QGraphicsSceneHelpEvent::scenePos()
func (q *QGraphicsSceneHelpEvent) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(276000,276104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneHelpEvent::screenPos()
func (q *QGraphicsSceneHelpEvent) ScreenPos() *QPoint {
	var __rv uintptr
	q.Drv(276000,276105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsSceneHelpEvent::setScenePos(QPointF const&)
func (q *QGraphicsSceneHelpEvent) SetScenePos(pos *QPointF)  {
	q.Drv(276000,276106,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneHelpEvent::setScreenPos(QPoint const&)
func (q *QGraphicsSceneHelpEvent) SetScreenPos(pos *QPoint)  {
	q.Drv(276000,276107,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
