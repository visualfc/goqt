// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWhatsThisClickedEvent : QWhatsThisClickedEvent
type QWhatsThisClickedEvent struct {
	QEvent
}
//QWhatsThisClickedEvent::QWhatsThisClickedEvent(QString const&)	
func NewWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),188000,188102,unsafe.Pointer(&href),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWhatsThisClickedEvent{}
	_p.SetDriver(__rv,188000,true)
	return _p
} 
//QWhatsThisClickedEvent::href()
func (q *QWhatsThisClickedEvent) Href() string {
	var __rv string
	q.Drv(188000,188103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
