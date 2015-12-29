// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFocusEvent : QFocusEvent
type QFocusEvent struct {
	QEvent
}
//QFocusEvent::QFocusEvent(QEvent::Type,Qt::FocusReason)	
func NewFocusEvent(_type QEvent_Type,reason Qt_FocusReason) *QFocusEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),36000,36102,unsafe.Pointer(&_type),unsafe.Pointer(&reason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFocusEvent{}
	_p.SetDriver(__rv,36000,true)
	return _p
} 
//QFocusEvent::gotFocus()
func (q *QFocusEvent) GotFocus() bool {
	var __rv bool
	q.Drv(36000,36103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFocusEvent::lostFocus()
func (q *QFocusEvent) LostFocus() bool {
	var __rv bool
	q.Drv(36000,36104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFocusEvent::reason()
func (q *QFocusEvent) Reason() Qt_FocusReason {
	var __rv Qt_FocusReason
	q.Drv(36000,36105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
