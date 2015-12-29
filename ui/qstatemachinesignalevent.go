// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStateMachineSignalEvent : QStateMachine::SignalEvent
type QStateMachineSignalEvent struct {
	QEvent
}
//QStateMachine::SignalEvent::SignalEvent(QObject*,int,QList<QVariant> const&)	
func NewStateMachineSignalEvent(sender QObjectInterface,signalIndex int,arguments []*QVariant) *QStateMachineSignalEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),124000,124102,Native(sender),unsafe.Pointer(&signalIndex),unsafe.Pointer(&arguments),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStateMachineSignalEvent{}
	_p.SetDriver(__rv,124000,true)
	return _p
} 
//QStateMachine::SignalEvent::arguments()
func (q *QStateMachineSignalEvent) Arguments() []*QVariant {
	var __rv []*QVariant
	q.Drv(124000,124103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::SignalEvent::sender()
func (q *QStateMachineSignalEvent) Sender() *QObject {
	var __rv uintptr
	q.Drv(124000,124104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QStateMachine::SignalEvent::signalIndex()
func (q *QStateMachineSignalEvent) SignalIndex() int {
	var __rv int
	q.Drv(124000,124105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
