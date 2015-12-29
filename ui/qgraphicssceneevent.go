// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsSceneEvent : QGraphicsSceneEvent
type QGraphicsSceneEvent struct {
	QEvent
}
//QGraphicsSceneEvent::QGraphicsSceneEvent(QEvent::Type)	
func NewGraphicsSceneEvent(_type QEvent_Type) *QGraphicsSceneEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),275000,275102,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSceneEvent{}
	_p.SetDriver(__rv,275000,true)
	return _p
} 
//QGraphicsSceneEvent::setWidget(QWidget*)
func (q *QGraphicsSceneEvent) SetWidget(widget QWidgetInterface)  {
	q.Drv(275000,275103,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSceneEvent::widget()
func (q *QGraphicsSceneEvent) Widget() *QWidget {
	var __rv uintptr
	q.Drv(275000,275104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
