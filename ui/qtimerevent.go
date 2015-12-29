// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTimerEvent : QTimerEvent
type QTimerEvent struct {
	QEvent
}
//QTimerEvent::QTimerEvent(int)	
func NewTimerEvent(timerId int) *QTimerEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),171000,171102,unsafe.Pointer(&timerId),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimerEvent{}
	_p.SetDriver(__rv,171000,true)
	return _p
} 
//QTimerEvent::timerId()
func (q *QTimerEvent) TimerId() int {
	var __rv int
	q.Drv(171000,171103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
