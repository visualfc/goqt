// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QChildEvent : QChildEvent
type QChildEvent struct {
	QEvent
}
//QChildEvent::QChildEvent(QEvent::Type,QObject*)	
func NewChildEvent(_type QEvent_Type,child QObjectInterface) *QChildEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),11000,11102,unsafe.Pointer(&_type),Native(child),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QChildEvent{}
	_p.SetDriver(__rv,11000,true)
	return _p
} 
//QChildEvent::added()
func (q *QChildEvent) Added() bool {
	var __rv bool
	q.Drv(11000,11103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QChildEvent::child()
func (q *QChildEvent) Child() *QObject {
	var __rv uintptr
	q.Drv(11000,11104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QChildEvent::polished()
func (q *QChildEvent) Polished() bool {
	var __rv bool
	q.Drv(11000,11105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QChildEvent::removed()
func (q *QChildEvent) Removed() bool {
	var __rv bool
	q.Drv(11000,11106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
