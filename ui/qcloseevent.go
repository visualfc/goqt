// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QCloseEvent : QCloseEvent
type QCloseEvent struct {
	QEvent
}
//QCloseEvent::QCloseEvent()	
func NewCloseEvent() *QCloseEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),12000,12102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCloseEvent{}
	_p.SetDriver(__rv,12000,true)
	return _p
} 
