// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMouseEventTransition : QMouseEventTransition
type QMouseEventTransition struct {
	QEventTransition
}
//QMouseEventTransition::QMouseEventTransition()	
func NewMouseEventTransition() *QMouseEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),312000,312102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMouseEventTransition{}
	_p.SetDriver(__rv,312000,false)
	return _p
} 
//QMouseEventTransition::QMouseEventTransition(QState*)	
func NewMouseEventTransitionWithSourcestate(sourceState *QState) *QMouseEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),312000,312103,Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMouseEventTransition{}
	_p.SetDriver(__rv,312000,false)
	return _p
} 
//QMouseEventTransition::QMouseEventTransition(QObject*,QEvent::Type,Qt::MouseButton,QState*)	
func NewMouseEventTransitionWithObjectTypeButtonSourcestate(object QObjectInterface,_type QEvent_Type,button Qt_MouseButton,sourceState *QState) *QMouseEventTransition {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),312000,312104,Native(object),unsafe.Pointer(&_type),unsafe.Pointer(&button),Native(sourceState),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMouseEventTransition{}
	_p.SetDriver(__rv,312000,false)
	return _p
} 
//QMouseEventTransition::button()
func (q *QMouseEventTransition) Button() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(312000,312105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEventTransition::eventTest(QEvent*)
func (q *QMouseEventTransition) EventTest(event *QEvent) bool {
	var __rv bool
	q.Drv(312000,312106,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEventTransition::hitTestPath()
func (q *QMouseEventTransition) HitTestPath() *QPainterPath {
	var __rv uintptr
	q.Drv(312000,312107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QMouseEventTransition::modifierMask()
func (q *QMouseEventTransition) ModifierMask() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(312000,312108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMouseEventTransition::onTransition(QEvent*)
func (q *QMouseEventTransition) OnTransition(event *QEvent)  {
	q.Drv(312000,312109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMouseEventTransition::setButton(Qt::MouseButton)
func (q *QMouseEventTransition) SetButton(button Qt_MouseButton)  {
	q.Drv(312000,312110,unsafe.Pointer(&button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMouseEventTransition::setHitTestPath(QPainterPath const&)
func (q *QMouseEventTransition) SetHitTestPath(path *QPainterPath)  {
	q.Drv(312000,312111,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMouseEventTransition::setModifierMask(QFlags<Qt::KeyboardModifier>)
func (q *QMouseEventTransition) SetModifierMask(modifiers Qt_KeyboardModifier)  {
	q.Drv(312000,312112,unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
