// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFinalState : QFinalState
type QFinalState struct {
	QAbstractState
}
//QFinalState::QFinalState()	
func NewFinalState() *QFinalState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),240000,240102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFinalState{}
	_p.SetDriver(__rv,240000,false)
	return _p
} 
//QFinalState::QFinalState(QState*)	
func NewFinalStateWithParent(parent *QState) *QFinalState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),240000,240103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFinalState{}
	_p.SetDriver(__rv,240000,false)
	return _p
} 
//QFinalState::event(QEvent*)
func (q *QFinalState) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(240000,240104,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFinalState::onEntry(QEvent*)
func (q *QFinalState) OnEntry(event *QEvent)  {
	q.Drv(240000,240105,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFinalState::onExit(QEvent*)
func (q *QFinalState) OnExit(event *QEvent)  {
	q.Drv(240000,240106,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
