// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QHoverEvent : QHoverEvent
type QHoverEvent struct {
	QEvent
}
//QHoverEvent::QHoverEvent(QEvent::Type,QPoint const&,QPoint const&)	
func NewHoverEvent(_type QEvent_Type,pos *QPoint,oldPos *QPoint) *QHoverEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),50000,50102,unsafe.Pointer(&_type),Native(pos),Native(oldPos),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHoverEvent{}
	_p.SetDriver(__rv,50000,true)
	return _p
} 
//QHoverEvent::oldPos()
func (q *QHoverEvent) OldPos() *QPoint {
	var __rv uintptr
	q.Drv(50000,50103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QHoverEvent::pos()
func (q *QHoverEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(50000,50104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
