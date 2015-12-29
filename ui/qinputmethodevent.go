// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QInputMethodEvent_AttributeType - QInputMethodEvent::AttributeType
type QInputMethodEvent_AttributeType uint32
const (
	QInputMethodEvent_TextFormat QInputMethodEvent_AttributeType = 0
	QInputMethodEvent_Cursor QInputMethodEvent_AttributeType = 1
	QInputMethodEvent_Language QInputMethodEvent_AttributeType = 2
	QInputMethodEvent_Ruby QInputMethodEvent_AttributeType = 3
	QInputMethodEvent_Selection QInputMethodEvent_AttributeType = 4
)
//struct QInputMethodEvent : QInputMethodEvent
type QInputMethodEvent struct {
	QEvent
}
//QInputMethodEvent::QInputMethodEvent()	
func NewInputMethodEvent() *QInputMethodEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),58000,58102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputMethodEvent{}
	_p.SetDriver(__rv,58000,true)
	return _p
} 
//QInputMethodEvent::QInputMethodEvent(QInputMethodEvent const&)	
func NewInputMethodEventCopy(other *QInputMethodEvent) *QInputMethodEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),58000,58103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputMethodEvent{}
	_p.SetDriver(__rv,58000,true)
	return _p
} 
//QInputMethodEvent::QInputMethodEvent(QString const&,QList<QInputMethodEvent::Attribute> const&)	
func NewInputMethodEventWithPreedittextAttributes(preeditText string,attributes []*QInputMethodEventAttribute) *QInputMethodEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),58000,58104,unsafe.Pointer(&preeditText),unsafe.Pointer(&attributes),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputMethodEvent{}
	_p.SetDriver(__rv,58000,true)
	return _p
} 
//QInputMethodEvent::attributes()
func (q *QInputMethodEvent) Attributes() []*QInputMethodEventAttribute {
	var __rv []*QInputMethodEventAttribute
	q.Drv(58000,58105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputMethodEvent::commitString()
func (q *QInputMethodEvent) CommitString() string {
	var __rv string
	q.Drv(58000,58106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputMethodEvent::preeditString()
func (q *QInputMethodEvent) PreeditString() string {
	var __rv string
	q.Drv(58000,58107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputMethodEvent::replacementLength()
func (q *QInputMethodEvent) ReplacementLength() int {
	var __rv int
	q.Drv(58000,58108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputMethodEvent::replacementStart()
func (q *QInputMethodEvent) ReplacementStart() int {
	var __rv int
	q.Drv(58000,58109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputMethodEvent::setCommitString(QString const&)
func (q *QInputMethodEvent) SetCommitString(commitString string)  {
	q.Drv(58000,58110,unsafe.Pointer(&commitString),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QInputMethodEvent::setCommitString(QString const&,int,int)
func (q *QInputMethodEvent) SetCommitStringWithCommitstringReplacefromReplacelength(commitString string,replaceFrom int,replaceLength int)  {
	q.Drv(58000,58111,unsafe.Pointer(&commitString),unsafe.Pointer(&replaceFrom),unsafe.Pointer(&replaceLength),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
