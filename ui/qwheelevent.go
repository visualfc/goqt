// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWheelEvent : QWheelEvent
type QWheelEvent struct {
	QInputEvent
}
//QWheelEvent::QWheelEvent(QPoint const&,int,QFlags<Qt::MouseButton>,QFlags<Qt::KeyboardModifier>,Qt::Orientation)	
func NewWheelEvent(pos *QPoint,delta int,buttons Qt_MouseButton,modifiers Qt_KeyboardModifier,orient Qt_Orientation) *QWheelEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),189000,189102,Native(pos),unsafe.Pointer(&delta),unsafe.Pointer(&buttons),unsafe.Pointer(&modifiers),unsafe.Pointer(&orient),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWheelEvent{}
	_p.SetDriver(__rv,189000,true)
	return _p
} 
//QWheelEvent::QWheelEvent(QPoint const&,QPoint const&,int,QFlags<Qt::MouseButton>,QFlags<Qt::KeyboardModifier>,Qt::Orientation)	
func NewWheelEventWithPosGlobalposDeltaButtonsModifiersOrient(pos *QPoint,globalPos *QPoint,delta int,buttons Qt_MouseButton,modifiers Qt_KeyboardModifier,orient Qt_Orientation) *QWheelEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),189000,189103,Native(pos),Native(globalPos),unsafe.Pointer(&delta),unsafe.Pointer(&buttons),unsafe.Pointer(&modifiers),unsafe.Pointer(&orient),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWheelEvent{}
	_p.SetDriver(__rv,189000,true)
	return _p
} 
//QWheelEvent::buttons()
func (q *QWheelEvent) Buttons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(189000,189104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::delta()
func (q *QWheelEvent) Delta() int {
	var __rv int
	q.Drv(189000,189105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::globalPos()
func (q *QWheelEvent) GlobalPos() *QPoint {
	var __rv uintptr
	q.Drv(189000,189106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWheelEvent::globalX()
func (q *QWheelEvent) GlobalX() int {
	var __rv int
	q.Drv(189000,189107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::globalY()
func (q *QWheelEvent) GlobalY() int {
	var __rv int
	q.Drv(189000,189108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::orientation()
func (q *QWheelEvent) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(189000,189109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::pos()
func (q *QWheelEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(189000,189110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWheelEvent::x()
func (q *QWheelEvent) X() int {
	var __rv int
	q.Drv(189000,189111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWheelEvent::y()
func (q *QWheelEvent) Y() int {
	var __rv int
	q.Drv(189000,189112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
