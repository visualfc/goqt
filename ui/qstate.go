// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QState_ChildMode - QState::ChildMode
type QState_ChildMode uint32
const (
	QState_ExclusiveStates QState_ChildMode = 0
	QState_ParallelStates QState_ChildMode = 1
)
//struct QState : QState
type QState struct {
	QAbstractState
}
func (q *QState) OnFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(353000,353102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QState) OnPropertiesAssigned(fn func()) uintptr {
	var __rv uintptr
	q.Drv(353000,353103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QState::QState()	
func NewState() *QState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),353000,353104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QState{}
	_p.SetDriver(__rv,353000,false)
	return _p
} 
//QState::QState(QState*)	
func NewStateWithParent(parent *QState) *QState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),353000,353105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QState{}
	_p.SetDriver(__rv,353000,false)
	return _p
} 
//QState::QState(QState::ChildMode,QState*)	
func NewStateWithChildmodeParent(childMode QState_ChildMode,parent *QState) *QState {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),353000,353106,unsafe.Pointer(&childMode),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QState{}
	_p.SetDriver(__rv,353000,false)
	return _p
} 
//QState::addTransition(QAbstractState*)
func (q *QState) AddTransition(target *QAbstractState) *QAbstractTransition {
	var __rv uintptr
	q.Drv(353000,353107,Native(target),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractTransition{}
	_rp.SetDriver(__rv,206000,false)
	return _rp
}	
//QState::addTransition(QAbstractTransition*)
func (q *QState) AddTransitionWithTransition(transition *QAbstractTransition)  {
	q.Drv(353000,353108,Native(transition),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::assignProperty(QObject*,char const*,QVariant const&)
func (q *QState) AssignProperty(object QObjectInterface,name string,value *QVariant)  {
	q.Drv(353000,353109,Native(object),unsafe.Pointer(&name),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::childMode()
func (q *QState) ChildMode() QState_ChildMode {
	var __rv QState_ChildMode
	q.Drv(353000,353110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QState::errorState()
func (q *QState) ErrorState() *QAbstractState {
	var __rv uintptr
	q.Drv(353000,353111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractState{}
	_rp.SetDriver(__rv,203000,false)
	return _rp
}	
//QState::event(QEvent*)
func (q *QState) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(353000,353112,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QState::initialState()
func (q *QState) InitialState() *QAbstractState {
	var __rv uintptr
	q.Drv(353000,353113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractState{}
	_rp.SetDriver(__rv,203000,false)
	return _rp
}	
//QState::onEntry(QEvent*)
func (q *QState) OnEntry(event *QEvent)  {
	q.Drv(353000,353114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::onExit(QEvent*)
func (q *QState) OnExit(event *QEvent)  {
	q.Drv(353000,353115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::removeTransition(QAbstractTransition*)
func (q *QState) RemoveTransition(transition *QAbstractTransition)  {
	q.Drv(353000,353116,Native(transition),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::setChildMode(QState::ChildMode)
func (q *QState) SetChildMode(mode QState_ChildMode)  {
	q.Drv(353000,353117,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::setErrorState(QAbstractState*)
func (q *QState) SetErrorState(state *QAbstractState)  {
	q.Drv(353000,353118,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::setInitialState(QAbstractState*)
func (q *QState) SetInitialState(state *QAbstractState)  {
	q.Drv(353000,353119,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QState::transitions()
func (q *QState) Transitions() []*QAbstractTransition {
	var __rv []*QAbstractTransition
	q.Drv(353000,353120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
