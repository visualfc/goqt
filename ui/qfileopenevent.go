// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFileOpenEvent : QFileOpenEvent
type QFileOpenEvent struct {
	QEvent
}
//QFileOpenEvent::QFileOpenEvent(QString const&)	
func NewFileOpenEvent(file string) *QFileOpenEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),35000,35102,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileOpenEvent{}
	_p.SetDriver(__rv,35000,true)
	return _p
} 
//QFileOpenEvent::QFileOpenEvent(QUrl const&)	
func NewFileOpenEventWithUrl(url *QUrl) *QFileOpenEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),35000,35103,Native(url),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileOpenEvent{}
	_p.SetDriver(__rv,35000,true)
	return _p
} 
//QFileOpenEvent::file()
func (q *QFileOpenEvent) File() string {
	var __rv string
	q.Drv(35000,35104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileOpenEvent::url()
func (q *QFileOpenEvent) Url() *QUrl {
	var __rv uintptr
	q.Drv(35000,35105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
