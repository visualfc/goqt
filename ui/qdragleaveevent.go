// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDragLeaveEvent : QDragLeaveEvent
type QDragLeaveEvent struct {
	QEvent
}
//QDragLeaveEvent::QDragLeaveEvent()	
func NewDragLeaveEvent() *QDragLeaveEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),25000,25102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDragLeaveEvent{}
	_p.SetDriver(__rv,25000,true)
	return _p
} 
