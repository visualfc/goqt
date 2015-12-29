// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWindowStateChangeEvent : QWindowStateChangeEvent
type QWindowStateChangeEvent struct {
	QEvent
}
//QWindowStateChangeEvent::QWindowStateChangeEvent(QFlags<Qt::WindowState>)	
func NewWindowStateChangeEvent(aOldState Qt_WindowState) *QWindowStateChangeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),191000,191102,unsafe.Pointer(&aOldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWindowStateChangeEvent{}
	_p.SetDriver(__rv,191000,true)
	return _p
} 
//QWindowStateChangeEvent::QWindowStateChangeEvent(QFlags<Qt::WindowState>,bool)	
func NewWindowStateChangeEventWithAoldstateIsoverride(aOldState Qt_WindowState,isOverride bool) *QWindowStateChangeEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),191000,191103,unsafe.Pointer(&aOldState),unsafe.Pointer(&isOverride),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWindowStateChangeEvent{}
	_p.SetDriver(__rv,191000,true)
	return _p
} 
//QWindowStateChangeEvent::isOverride()
func (q *QWindowStateChangeEvent) IsOverride() bool {
	var __rv bool
	q.Drv(191000,191104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWindowStateChangeEvent::oldState()
func (q *QWindowStateChangeEvent) OldState() Qt_WindowState {
	var __rv Qt_WindowState
	q.Drv(191000,191105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
