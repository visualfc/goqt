// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneResizeEvent : QGraphicsSceneResizeEvent
type QGraphicsSceneResizeEvent struct {
	QGraphicsSceneEvent
}
//QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent()	
func NewGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),280000,280102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneResizeEvent{}
	_p.SetDriver(__rv,280000,true)
	return _p
} 
//QGraphicsSceneResizeEvent::newSize()
func (q *QGraphicsSceneResizeEvent) NewSize() *QSizeF {
	var __rv uintptr
	q.Drv(280000,280103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsSceneResizeEvent::oldSize()
func (q *QGraphicsSceneResizeEvent) OldSize() *QSizeF {
	var __rv uintptr
	q.Drv(280000,280104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsSceneResizeEvent::setNewSize(QSizeF const&)
func (q *QGraphicsSceneResizeEvent) SetNewSize(size *QSizeF)  {
	q.Drv(280000,280105,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneResizeEvent::setOldSize(QSizeF const&)
func (q *QGraphicsSceneResizeEvent) SetOldSize(size *QSizeF)  {
	q.Drv(280000,280106,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
