// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QEventTransition : QEventTransition
type QEventTransition struct {
	QAbstractTransition
}
//QEventTransition::QEventTransition()	
func NewEventTransition() *QEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),235000,235102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEventTransition{}
	_p.SetDriver(__rv,235000,false)
	return _p
} 
//QEventTransition::QEventTransition(QState*)	
func NewEventTransitionWithSourcestate(sourceState *QState) *QEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),235000,235103,Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEventTransition{}
	_p.SetDriver(__rv,235000,false)
	return _p
} 
//QEventTransition::QEventTransition(QObject*,QEvent::Type,QState*)	
func NewEventTransitionWithObjectTypeSourcestate(object QObjectInterface,_type QEvent_Type,sourceState *QState) *QEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),235000,235104,Native(object),unsafe.Pointer(&_type),Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEventTransition{}
	_p.SetDriver(__rv,235000,false)
	return _p
} 
//QEventTransition::event(QEvent*)
func (q *QEventTransition) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(235000,235105,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventTransition::eventSource()
func (q *QEventTransition) EventSource() *QObject {
	var __rv uintptr
	q.Drv(235000,235106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QEventTransition::eventTest(QEvent*)
func (q *QEventTransition) EventTest(event *QEvent) bool {
	var __rv bool
	q.Drv(235000,235107,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventTransition::eventType()
func (q *QEventTransition) EventType() QEvent_Type {
	var __rv QEvent_Type
	q.Drv(235000,235108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventTransition::onTransition(QEvent*)
func (q *QEventTransition) OnTransition(event *QEvent)  {
	q.Drv(235000,235109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventTransition::setEventSource(QObject*)
func (q *QEventTransition) SetEventSource(object QObjectInterface)  {
	q.Drv(235000,235110,Native(object),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventTransition::setEventType(QEvent::Type)
func (q *QEventTransition) SetEventType(_type QEvent_Type)  {
	q.Drv(235000,235111,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
