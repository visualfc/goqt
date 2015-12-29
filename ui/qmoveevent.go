// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMoveEvent : QMoveEvent
type QMoveEvent struct {
	QEvent
}
//QMoveEvent::QMoveEvent(QPoint const&,QPoint const&)	
func NewMoveEvent(pos *QPoint,oldPos *QPoint) *QMoveEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),81000,81102,Native(pos),Native(oldPos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMoveEvent{}
	_p.SetDriver(__rv,81000,true)
	return _p
} 
//QMoveEvent::oldPos()
func (q *QMoveEvent) OldPos() *QPoint {
	var __rv uintptr
	q.Drv(81000,81103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QMoveEvent::pos()
func (q *QMoveEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(81000,81104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
