// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPropertyAnimation : QPropertyAnimation
type QPropertyAnimation struct {
	QVariantAnimation
}
//QPropertyAnimation::QPropertyAnimation()	
func NewPropertyAnimation() *QPropertyAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),329000,329102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPropertyAnimation{}
	_p.SetDriver(__rv,329000,false)
	return _p
} 
//QPropertyAnimation::QPropertyAnimation(QObject*)	
func NewPropertyAnimationWithParent(parent QObjectInterface) *QPropertyAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),329000,329103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPropertyAnimation{}
	_p.SetDriver(__rv,329000,false)
	return _p
} 
//QPropertyAnimation::QPropertyAnimation(QObject*,QByteArray const&,QObject*)	
func NewPropertyAnimationWithObjectPropertynameParent(target QObjectInterface,propertyName []byte,parent QObjectInterface) *QPropertyAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),329000,329104,Native(target),unsafe.Pointer(&propertyName),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPropertyAnimation{}
	_p.SetDriver(__rv,329000,false)
	return _p
} 
//QPropertyAnimation::event(QEvent*)
func (q *QPropertyAnimation) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(329000,329105,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPropertyAnimation::propertyName()
func (q *QPropertyAnimation) PropertyName() []byte {
	var __rv []byte
	q.Drv(329000,329106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPropertyAnimation::setPropertyName(QByteArray const&)
func (q *QPropertyAnimation) SetPropertyName(propertyName []byte)  {
	q.Drv(329000,329107,unsafe.Pointer(&propertyName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPropertyAnimation::setTargetObject(QObject*)
func (q *QPropertyAnimation) SetTargetObject(target QObjectInterface)  {
	q.Drv(329000,329108,Native(target),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPropertyAnimation::targetObject()
func (q *QPropertyAnimation) TargetObject() *QObject {
	var __rv uintptr
	q.Drv(329000,329109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QPropertyAnimation::updateCurrentValue(QVariant const&)
func (q *QPropertyAnimation) UpdateCurrentValue(value *QVariant)  {
	q.Drv(329000,329110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPropertyAnimation::updateState(QAbstractAnimation::State,QAbstractAnimation::State)
func (q *QPropertyAnimation) UpdateState(newState QAbstractAnimation_State,oldState QAbstractAnimation_State)  {
	q.Drv(329000,329111,unsafe.Pointer(&newState),unsafe.Pointer(&oldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
