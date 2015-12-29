// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStatusTipEvent : QStatusTipEvent
type QStatusTipEvent struct {
	QEvent
}
//QStatusTipEvent::QStatusTipEvent(QString const&)	
func NewStatusTipEvent(tip string) *QStatusTipEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),127000,127102,unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStatusTipEvent{}
	_p.SetDriver(__rv,127000,true)
	return _p
} 
//QStatusTipEvent::tip()
func (q *QStatusTipEvent) Tip() string {
	var __rv string
	q.Drv(127000,127103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
