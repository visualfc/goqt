// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QShortcutEvent : QShortcutEvent
type QShortcutEvent struct {
	QEvent
}
//QShortcutEvent::QShortcutEvent(QKeySequence const&,int,bool)	
func NewShortcutEvent(key *QKeySequence,id int,ambiguous bool) *QShortcutEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),117000,117102,Native(key),unsafe.Pointer(&id),unsafe.Pointer(&ambiguous),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QShortcutEvent{}
	_p.SetDriver(__rv,117000,true)
	return _p
} 
//QShortcutEvent::isAmbiguous()
func (q *QShortcutEvent) IsAmbiguous() bool {
	var __rv bool
	q.Drv(117000,117103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcutEvent::key()
func (q *QShortcutEvent) Key() *QKeySequence {
	var __rv uintptr
	q.Drv(117000,117104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QShortcutEvent::shortcutId()
func (q *QShortcutEvent) ShortcutId() int {
	var __rv int
	q.Drv(117000,117105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
