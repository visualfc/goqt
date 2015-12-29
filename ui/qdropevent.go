// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDropEvent : QDropEvent
type QDropEvent struct {
	QEvent
}
//QDropEvent::acceptProposedAction()
func (q *QDropEvent) AcceptProposedAction()  {
	q.Drv(27000,27102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDropEvent::dropAction()
func (q *QDropEvent) DropAction() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(27000,27103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDropEvent::keyboardModifiers()
func (q *QDropEvent) KeyboardModifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(27000,27104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDropEvent::mimeData()
func (q *QDropEvent) MimeData() *QMimeData {
	var __rv uintptr
	q.Drv(27000,27105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QDropEvent::mouseButtons()
func (q *QDropEvent) MouseButtons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(27000,27106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDropEvent::pos()
func (q *QDropEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(27000,27107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QDropEvent::possibleActions()
func (q *QDropEvent) PossibleActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(27000,27108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDropEvent::proposedAction()
func (q *QDropEvent) ProposedAction() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(27000,27109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDropEvent::setDropAction(Qt::DropAction)
func (q *QDropEvent) SetDropAction(action Qt_DropAction)  {
	q.Drv(27000,27110,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDropEvent::source()
func (q *QDropEvent) Source() *QWidget {
	var __rv uintptr
	q.Drv(27000,27111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
