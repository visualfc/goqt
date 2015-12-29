// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QHideEvent : QHideEvent
type QHideEvent struct {
	QEvent
}
//QHideEvent::QHideEvent()	
func NewHideEvent() *QHideEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),49000,49102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHideEvent{}
	_p.SetDriver(__rv,49000,true)
	return _p
} 
