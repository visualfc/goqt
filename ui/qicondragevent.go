// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QIconDragEvent : QIconDragEvent
type QIconDragEvent struct {
	QEvent
}
//QIconDragEvent::QIconDragEvent()	
func NewIconDragEvent() *QIconDragEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),52000,52102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIconDragEvent{}
	_p.SetDriver(__rv,52000,true)
	return _p
} 
