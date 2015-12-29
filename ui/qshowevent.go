// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QShowEvent : QShowEvent
type QShowEvent struct {
	QEvent
}
//QShowEvent::QShowEvent()	
func NewShowEvent() *QShowEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),118000,118102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QShowEvent{}
	_p.SetDriver(__rv,118000,true)
	return _p
} 
