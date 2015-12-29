// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QInputEvent : QInputEvent
type QInputEvent struct {
	QEvent
}
//QInputEvent::QInputEvent(QEvent::Type,QFlags<Qt::KeyboardModifier>)	
func NewInputEvent(_type QEvent_Type,modifiers Qt_KeyboardModifier) *QInputEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),57000,57102,unsafe.Pointer(&_type),unsafe.Pointer(&modifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QInputEvent{}
	_p.SetDriver(__rv,57000,true)
	return _p
} 
//QInputEvent::modifiers()
func (q *QInputEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(57000,57103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QInputEvent::setModifiers(QFlags<Qt::KeyboardModifier>)
func (q *QInputEvent) SetModifiers(amodifiers Qt_KeyboardModifier)  {
	q.Drv(57000,57104,unsafe.Pointer(&amodifiers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
