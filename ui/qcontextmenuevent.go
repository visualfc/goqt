// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QContextMenuEvent_Reason - QContextMenuEvent::Reason
type QContextMenuEvent_Reason uint32
const (
	QContextMenuEvent_Mouse QContextMenuEvent_Reason = 0
	QContextMenuEvent_Keyboard QContextMenuEvent_Reason = 1
	QContextMenuEvent_Other QContextMenuEvent_Reason = 2
)
//struct QContextMenuEvent : QContextMenuEvent
type QContextMenuEvent struct {
	QInputEvent
}
//QContextMenuEvent::QContextMenuEvent(QContextMenuEvent::Reason,QPoint const&)	
func NewContextMenuEvent(reason QContextMenuEvent_Reason,pos *QPoint) *QContextMenuEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),16000,16102,unsafe.Pointer(&reason),Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QContextMenuEvent{}
	_p.SetDriver(__rv,16000,true)
	return _p
} 
//QContextMenuEvent::QContextMenuEvent(QContextMenuEvent::Reason,QPoint const&,QPoint const&)	
func NewContextMenuEventWithReasonPosGlobalpos(reason QContextMenuEvent_Reason,pos *QPoint,globalPos *QPoint) *QContextMenuEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),16000,16103,unsafe.Pointer(&reason),Native(pos),Native(globalPos),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QContextMenuEvent{}
	_p.SetDriver(__rv,16000,true)
	return _p
} 
//QContextMenuEvent::QContextMenuEvent(QContextMenuEvent::Reason,QPoint const&,QPoint const&,QFlags<Qt::KeyboardModifier>)	
func NewContextMenuEventWithReasonPosGlobalposModifiers(reason QContextMenuEvent_Reason,pos *QPoint,globalPos *QPoint,modifiers Qt_KeyboardModifier) *QContextMenuEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),16000,16104,unsafe.Pointer(&reason),Native(pos),Native(globalPos),unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QContextMenuEvent{}
	_p.SetDriver(__rv,16000,true)
	return _p
} 
//QContextMenuEvent::globalPos()
func (q *QContextMenuEvent) GlobalPos() *QPoint {
	var __rv uintptr
	q.Drv(16000,16105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QContextMenuEvent::globalX()
func (q *QContextMenuEvent) GlobalX() int {
	var __rv int
	q.Drv(16000,16106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QContextMenuEvent::globalY()
func (q *QContextMenuEvent) GlobalY() int {
	var __rv int
	q.Drv(16000,16107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QContextMenuEvent::pos()
func (q *QContextMenuEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(16000,16108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QContextMenuEvent::reason()
func (q *QContextMenuEvent) Reason() QContextMenuEvent_Reason {
	var __rv QContextMenuEvent_Reason
	q.Drv(16000,16109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QContextMenuEvent::x()
func (q *QContextMenuEvent) X() int {
	var __rv int
	q.Drv(16000,16110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QContextMenuEvent::y()
func (q *QContextMenuEvent) Y() int {
	var __rv int
	q.Drv(16000,16111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
