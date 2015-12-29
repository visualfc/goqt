// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QResizeEvent : QResizeEvent
type QResizeEvent struct {
	QEvent
}
//QResizeEvent::QResizeEvent(QSize const&,QSize const&)	
func NewResizeEvent(size *QSize,oldSize *QSize) *QResizeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),115000,115102,Native(size),Native(oldSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QResizeEvent{}
	_p.SetDriver(__rv,115000,true)
	return _p
} 
//QResizeEvent::oldSize()
func (q *QResizeEvent) OldSize() *QSize {
	var __rv uintptr
	q.Drv(115000,115103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QResizeEvent::size()
func (q *QResizeEvent) Size() *QSize {
	var __rv uintptr
	q.Drv(115000,115104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
