// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QActionEvent : QActionEvent
type QActionEvent struct {
	QEvent
}
//QActionEvent::QActionEvent(int,QAction*,QAction*)	
func NewActionEvent(_type int,action *QAction,before *QAction) *QActionEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),5000,5102,unsafe.Pointer(&_type),Native(action),Native(before),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QActionEvent{}
	_p.SetDriver(__rv,5000,true)
	return _p
} 
//QActionEvent::action()
func (q *QActionEvent) Action() *QAction {
	var __rv uintptr
	q.Drv(5000,5103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QActionEvent::before()
func (q *QActionEvent) Before() *QAction {
	var __rv uintptr
	q.Drv(5000,5104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
