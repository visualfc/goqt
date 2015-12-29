// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractState : QAbstractState
type QAbstractState struct {
	QObject
}
//QAbstractState::event(QEvent*)
func (q *QAbstractState) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(203000,203102,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractState::machine()
func (q *QAbstractState) Machine() *QStateMachine {
	var __rv uintptr
	q.Drv(203000,203103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStateMachine{}
	_rp.SetDriver(__rv,354000,false)
	return _rp
}	
//QAbstractState::onEntry(QEvent*)
func (q *QAbstractState) OnEntry(event *QEvent)  {
	q.Drv(203000,203104,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractState::onExit(QEvent*)
func (q *QAbstractState) OnExit(event *QEvent)  {
	q.Drv(203000,203105,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractState::parentState()
func (q *QAbstractState) ParentState() *QState {
	var __rv uintptr
	q.Drv(203000,203106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QState{}
	_rp.SetDriver(__rv,353000,false)
	return _rp
}	
