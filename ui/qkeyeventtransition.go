// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QKeyEventTransition : QKeyEventTransition
type QKeyEventTransition struct {
	QEventTransition
}
//QKeyEventTransition::QKeyEventTransition()	
func NewKeyEventTransition() *QKeyEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),297000,297102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeyEventTransition{}
	_p.SetDriver(__rv,297000,false)
	return _p
} 
//QKeyEventTransition::QKeyEventTransition(QState*)	
func NewKeyEventTransitionWithSourcestate(sourceState *QState) *QKeyEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),297000,297103,Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeyEventTransition{}
	_p.SetDriver(__rv,297000,false)
	return _p
} 
//QKeyEventTransition::QKeyEventTransition(QObject*,QEvent::Type,int,QState*)	
func NewKeyEventTransitionWithObjectTypeKeySourcestate(object QObjectInterface,_type QEvent_Type,key int,sourceState *QState) *QKeyEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),297000,297104,Native(object),unsafe.Pointer(&_type),unsafe.Pointer(&key),Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeyEventTransition{}
	_p.SetDriver(__rv,297000,false)
	return _p
} 
//QKeyEventTransition::eventTest(QEvent*)
func (q *QKeyEventTransition) EventTest(event *QEvent) bool {
	var __rv bool
	q.Drv(297000,297105,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEventTransition::key()
func (q *QKeyEventTransition) Key() int {
	var __rv int
	q.Drv(297000,297106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEventTransition::modifierMask()
func (q *QKeyEventTransition) ModifierMask() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(297000,297107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEventTransition::onTransition(QEvent*)
func (q *QKeyEventTransition) OnTransition(event *QEvent)  {
	q.Drv(297000,297108,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QKeyEventTransition::setKey(int)
func (q *QKeyEventTransition) SetKey(key int)  {
	q.Drv(297000,297109,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QKeyEventTransition::setModifierMask(QFlags<Qt::KeyboardModifier>)
func (q *QKeyEventTransition) SetModifierMask(modifiers Qt_KeyboardModifier)  {
	q.Drv(297000,297110,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
