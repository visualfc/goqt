// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStateMachine_Error - QStateMachine::Error
type QStateMachine_Error uint32
const (
	QStateMachine_NoError QStateMachine_Error = 0
	QStateMachine_NoInitialStateError QStateMachine_Error = 1
	QStateMachine_NoDefaultStateInHistoryStateError QStateMachine_Error = 2
	QStateMachine_NoCommonAncestorForTransitionError QStateMachine_Error = 3
)
//enum QStateMachine_EventPriority - QStateMachine::EventPriority
type QStateMachine_EventPriority uint32
const (
	QStateMachine_NormalPriority QStateMachine_EventPriority = 0
	QStateMachine_HighPriority QStateMachine_EventPriority = 1
)
//enum QStateMachine_RestorePolicy - QStateMachine::RestorePolicy
type QStateMachine_RestorePolicy uint32
const (
	QStateMachine_DontRestoreProperties QStateMachine_RestorePolicy = 0
	QStateMachine_RestoreProperties QStateMachine_RestorePolicy = 1
)
//struct QStateMachine : QStateMachine
type QStateMachine struct {
	QState
}
func (q *QStateMachine) OnStarted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(354000,354102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QStateMachine) OnStopped(fn func()) uintptr {
	var __rv uintptr
	q.Drv(354000,354103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QStateMachine::QStateMachine()	
func NewStateMachine() *QStateMachine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),354000,354104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStateMachine{}
	_p.SetDriver(__rv,354000,false)
	return _p
} 
//QStateMachine::QStateMachine(QObject*)	
func NewStateMachineWithParent(parent QObjectInterface) *QStateMachine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),354000,354105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStateMachine{}
	_p.SetDriver(__rv,354000,false)
	return _p
} 
//QStateMachine::addDefaultAnimation(QAbstractAnimation*)
func (q *QStateMachine) AddDefaultAnimation(animation *QAbstractAnimation)  {
	q.Drv(354000,354106,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::addState(QAbstractState*)
func (q *QStateMachine) AddState(state *QAbstractState)  {
	q.Drv(354000,354107,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::beginMicrostep(QEvent*)
func (q *QStateMachine) BeginMicrostep(event *QEvent)  {
	q.Drv(354000,354108,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::beginSelectTransitions(QEvent*)
func (q *QStateMachine) BeginSelectTransitions(event *QEvent)  {
	q.Drv(354000,354109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::cancelDelayedEvent(int)
func (q *QStateMachine) CancelDelayedEvent(id int) bool {
	var __rv bool
	q.Drv(354000,354110,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::clearError()
func (q *QStateMachine) ClearError()  {
	q.Drv(354000,354111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::defaultAnimations()
func (q *QStateMachine) DefaultAnimations() []*QAbstractAnimation {
	var __rv []*QAbstractAnimation
	q.Drv(354000,354112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::endMicrostep(QEvent*)
func (q *QStateMachine) EndMicrostep(event *QEvent)  {
	q.Drv(354000,354113,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::endSelectTransitions(QEvent*)
func (q *QStateMachine) EndSelectTransitions(event *QEvent)  {
	q.Drv(354000,354114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::error()
func (q *QStateMachine) Error() QStateMachine_Error {
	var __rv QStateMachine_Error
	q.Drv(354000,354115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::errorString()
func (q *QStateMachine) ErrorString() string {
	var __rv string
	q.Drv(354000,354116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::event(QEvent*)
func (q *QStateMachine) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(354000,354117,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::eventFilter(QObject*,QEvent*)
func (q *QStateMachine) EventFilter(watched QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(354000,354118,Native(watched),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::globalRestorePolicy()
func (q *QStateMachine) GlobalRestorePolicy() QStateMachine_RestorePolicy {
	var __rv QStateMachine_RestorePolicy
	q.Drv(354000,354119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::isAnimated()
func (q *QStateMachine) IsAnimated() bool {
	var __rv bool
	q.Drv(354000,354120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::isRunning()
func (q *QStateMachine) IsRunning() bool {
	var __rv bool
	q.Drv(354000,354121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::onEntry(QEvent*)
func (q *QStateMachine) OnEntry(event *QEvent)  {
	q.Drv(354000,354122,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::onExit(QEvent*)
func (q *QStateMachine) OnExit(event *QEvent)  {
	q.Drv(354000,354123,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::postDelayedEvent(QEvent*,int)
func (q *QStateMachine) PostDelayedEvent(event *QEvent,delay int) int {
	var __rv int
	q.Drv(354000,354124,Native(event),unsafe.Pointer(&delay),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStateMachine::postEvent(QEvent*)
func (q *QStateMachine) PostEvent(event *QEvent)  {
	q.Drv(354000,354125,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::postEvent(QEvent*,QStateMachine::EventPriority)
func (q *QStateMachine) PostEventWithEventPriority(event *QEvent,priority QStateMachine_EventPriority)  {
	q.Drv(354000,354126,Native(event),unsafe.Pointer(&priority),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::removeDefaultAnimation(QAbstractAnimation*)
func (q *QStateMachine) RemoveDefaultAnimation(animation *QAbstractAnimation)  {
	q.Drv(354000,354127,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::removeState(QAbstractState*)
func (q *QStateMachine) RemoveState(state *QAbstractState)  {
	q.Drv(354000,354128,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::setAnimated(bool)
func (q *QStateMachine) SetAnimated(enabled bool)  {
	q.Drv(354000,354129,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::setGlobalRestorePolicy(QStateMachine::RestorePolicy)
func (q *QStateMachine) SetGlobalRestorePolicy(restorePolicy QStateMachine_RestorePolicy)  {
	q.Drv(354000,354130,unsafe.Pointer(&restorePolicy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::start()
func (q *QStateMachine) Start()  {
	q.Drv(354000,354131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStateMachine::stop()
func (q *QStateMachine) Stop()  {
	q.Drv(354000,354132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
