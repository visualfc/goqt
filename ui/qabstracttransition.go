// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractTransition : QAbstractTransition
type QAbstractTransition struct {
	QObject
}
//QAbstractTransition::addAnimation(QAbstractAnimation*)
func (q *QAbstractTransition) AddAnimation(animation *QAbstractAnimation)  {
	q.Drv(206000,206102,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTransition::animations()
func (q *QAbstractTransition) Animations() []*QAbstractAnimation {
	var __rv []*QAbstractAnimation
	q.Drv(206000,206103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTransition::event(QEvent*)
func (q *QAbstractTransition) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(206000,206104,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTransition::eventTest(QEvent*)
func (q *QAbstractTransition) EventTest(event *QEvent) bool {
	var __rv bool
	q.Drv(206000,206105,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTransition::machine()
func (q *QAbstractTransition) Machine() *QStateMachine {
	var __rv uintptr
	q.Drv(206000,206106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStateMachine{}
	_rp.SetDriver(__rv,354000,false)
	return _rp
}	
//QAbstractTransition::onTransition(QEvent*)
func (q *QAbstractTransition) OnTransition(event *QEvent)  {
	q.Drv(206000,206107,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTransition::removeAnimation(QAbstractAnimation*)
func (q *QAbstractTransition) RemoveAnimation(animation *QAbstractAnimation)  {
	q.Drv(206000,206108,Native(animation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTransition::setTargetState(QAbstractState*)
func (q *QAbstractTransition) SetTargetState(target *QAbstractState)  {
	q.Drv(206000,206109,Native(target),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTransition::setTargetStates(QList<QAbstractState*> const&)
func (q *QAbstractTransition) SetTargetStates(targets []*QAbstractState)  {
	q.Drv(206000,206110,unsafe.Pointer(&targets),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTransition::sourceState()
func (q *QAbstractTransition) SourceState() *QState {
	var __rv uintptr
	q.Drv(206000,206111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QState{}
	_rp.SetDriver(__rv,353000,false)
	return _rp
}	
//QAbstractTransition::targetState()
func (q *QAbstractTransition) TargetState() *QAbstractState {
	var __rv uintptr
	q.Drv(206000,206112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractState{}
	_rp.SetDriver(__rv,203000,false)
	return _rp
}	
//QAbstractTransition::targetStates()
func (q *QAbstractTransition) TargetStates() []*QAbstractState {
	var __rv []*QAbstractState
	q.Drv(206000,206113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
