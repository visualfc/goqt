// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QHelpEvent : QHelpEvent
type QHelpEvent struct {
	QEvent
}
//QHelpEvent::QHelpEvent(QEvent::Type,QPoint const&,QPoint const&)	
func NewHelpEvent(_type QEvent_Type,pos *QPoint,globalPos *QPoint) *QHelpEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),48000,48102,unsafe.Pointer(&_type),Native(pos),Native(globalPos),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHelpEvent{}
	_p.SetDriver(__rv,48000,true)
	return _p
} 
//QHelpEvent::globalPos()
func (q *QHelpEvent) GlobalPos() *QPoint {
	var __rv uintptr
	q.Drv(48000,48103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QHelpEvent::globalX()
func (q *QHelpEvent) GlobalX() int {
	var __rv int
	q.Drv(48000,48104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHelpEvent::globalY()
func (q *QHelpEvent) GlobalY() int {
	var __rv int
	q.Drv(48000,48105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHelpEvent::pos()
func (q *QHelpEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(48000,48106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QHelpEvent::x()
func (q *QHelpEvent) X() int {
	var __rv int
	q.Drv(48000,48107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHelpEvent::y()
func (q *QHelpEvent) Y() int {
	var __rv int
	q.Drv(48000,48108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
