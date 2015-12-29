// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMouseEvent : QMouseEvent
type QMouseEvent struct {
	QInputEvent
}
//QMouseEvent::QMouseEvent(QEvent::Type,QPoint const&,Qt::MouseButton,QFlags<Qt::MouseButton>,QFlags<Qt::KeyboardModifier>)	
func NewMouseEvent(_type QEvent_Type,pos *QPoint,button Qt_MouseButton,buttons Qt_MouseButton,modifiers Qt_KeyboardModifier) *QMouseEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),80000,80102,unsafe.Pointer(&_type),Native(pos),unsafe.Pointer(&button),unsafe.Pointer(&buttons),unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMouseEvent{}
	_p.SetDriver(__rv,80000,true)
	return _p
} 
//QMouseEvent::QMouseEvent(QEvent::Type,QPoint const&,QPoint const&,Qt::MouseButton,QFlags<Qt::MouseButton>,QFlags<Qt::KeyboardModifier>)	
func NewMouseEventWithTypePosGlobalposButtonButtonsModifiers(_type QEvent_Type,pos *QPoint,globalPos *QPoint,button Qt_MouseButton,buttons Qt_MouseButton,modifiers Qt_KeyboardModifier) *QMouseEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),80000,80103,unsafe.Pointer(&_type),Native(pos),Native(globalPos),unsafe.Pointer(&button),unsafe.Pointer(&buttons),unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMouseEvent{}
	_p.SetDriver(__rv,80000,true)
	return _p
} 
//QMouseEvent::button()
func (q *QMouseEvent) Button() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(80000,80104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEvent::buttons()
func (q *QMouseEvent) Buttons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(80000,80105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEvent::globalPos()
func (q *QMouseEvent) GlobalPos() *QPoint {
	var __rv uintptr
	q.Drv(80000,80106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QMouseEvent::globalX()
func (q *QMouseEvent) GlobalX() int {
	var __rv int
	q.Drv(80000,80107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEvent::globalY()
func (q *QMouseEvent) GlobalY() int {
	var __rv int
	q.Drv(80000,80108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEvent::pos()
func (q *QMouseEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(80000,80109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QMouseEvent::posF()
func (q *QMouseEvent) LocalPos() *QPointF {
	var __rv uintptr
	q.Drv(80000,80110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QMouseEvent::x()
func (q *QMouseEvent) X() int {
	var __rv int
	q.Drv(80000,80111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEvent::y()
func (q *QMouseEvent) Y() int {
	var __rv int
	q.Drv(80000,80112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
