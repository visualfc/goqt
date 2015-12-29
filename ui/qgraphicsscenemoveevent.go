// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneMoveEvent : QGraphicsSceneMoveEvent
type QGraphicsSceneMoveEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent()	
func NewGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),279000,279102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneMoveEvent{}
	_p.SetDriver(__rv,279000,true)
	return _p
} 
//QGraphicsSceneMoveEvent::newPos()
func (q *QGraphicsSceneMoveEvent) NewPos() *QPointF {
	var __rv uintptr
	q.Drv(279000,279103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMoveEvent::oldPos()
func (q *QGraphicsSceneMoveEvent) OldPos() *QPointF {
	var __rv uintptr
	q.Drv(279000,279104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsSceneMoveEvent::setNewPos(QPointF const&)
func (q *QGraphicsSceneMoveEvent) SetNewPos(pos *QPointF)  {
	q.Drv(279000,279105,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneMoveEvent::setOldPos(QPointF const&)
func (q *QGraphicsSceneMoveEvent) SetOldPos(pos *QPointF)  {
	q.Drv(279000,279106,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
