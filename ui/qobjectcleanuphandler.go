// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QObjectCleanupHandler : QObjectCleanupHandler
type QObjectCleanupHandler struct {
	QObject
}
//QObjectCleanupHandler::QObjectCleanupHandler()	
func NewObjectCleanupHandler() *QObjectCleanupHandler {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),315000,315102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QObjectCleanupHandler{}
	_p.SetDriver(__rv,315000,false)
	return _p
} 
//QObjectCleanupHandler::add(QObject*)
func (q *QObjectCleanupHandler) Add(object QObjectInterface) *QObject {
	var __rv uintptr
	q.Drv(315000,315103,Native(object),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QObjectCleanupHandler::clear()
func (q *QObjectCleanupHandler) Clear()  {
	q.Drv(315000,315104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObjectCleanupHandler::isEmpty()
func (q *QObjectCleanupHandler) IsEmpty() bool {
	var __rv bool
	q.Drv(315000,315105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObjectCleanupHandler::remove(QObject*)
func (q *QObjectCleanupHandler) Remove(object QObjectInterface)  {
	q.Drv(315000,315106,Native(object),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
