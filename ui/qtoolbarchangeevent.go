// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QToolBarChangeEvent : QToolBarChangeEvent
type QToolBarChangeEvent struct {
	QEvent
}
//QToolBarChangeEvent::QToolBarChangeEvent(bool)	
func NewToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),172000,172102,unsafe.Pointer(&t),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBarChangeEvent{}
	_p.SetDriver(__rv,172000,true)
	return _p
} 
//QToolBarChangeEvent::toggle()
func (q *QToolBarChangeEvent) Toggle() bool {
	var __rv bool
	q.Drv(172000,172103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
