// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStateMachineWrappedEvent : QStateMachine::WrappedEvent
type QStateMachineWrappedEvent struct {
	QEvent
}
//QStateMachine::WrappedEvent::WrappedEvent(QObject*,QEvent*)	
func NewStateMachineWrappedEvent(object QObjectInterface,event *QEvent) *QStateMachineWrappedEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),125000,125102,Native(object),Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStateMachineWrappedEvent{}
	_p.SetDriver(__rv,125000,true)
	return _p
} 
//QStateMachine::WrappedEvent::event()
func (q *QStateMachineWrappedEvent) Event() *QEvent {
	var __rv uintptr
	q.Drv(125000,125103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QEvent{}
	_rp.SetDriver(__rv,31000,true)
	return _rp
}	
//QStateMachine::WrappedEvent::object()
func (q *QStateMachineWrappedEvent) Object() *QObject {
	var __rv uintptr
	q.Drv(125000,125104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
