// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QErrorMessage : QErrorMessage
type QErrorMessage struct {
	QDialog
}
//QErrorMessage::QErrorMessage()	
func NewErrorMessage() *QErrorMessage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),233000,233102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QErrorMessage{}
	_p.SetDriver(__rv,233000,false)
	return _p
} 
//QErrorMessage::QErrorMessage(QWidget*)	
func NewErrorMessageWithParent(parent QWidgetInterface) *QErrorMessage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),233000,233103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QErrorMessage{}
	_p.SetDriver(__rv,233000,false)
	return _p
} 
//QErrorMessage::changeEvent(QEvent*)
func (q *QErrorMessage) ChangeEvent(e *QEvent)  {
	q.Drv(233000,233104,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QErrorMessage::done(int)
func (q *QErrorMessage) Done(value int)  {
	q.Drv(233000,233105,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QErrorMessage::qtHandler()	
func QErrorMessageQtHandler() *QErrorMessage {
	var __rv uintptr
	DirectQtDrv(nil,233000,233106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QErrorMessage{}
	_rp.SetDriver(__rv,233000,false)
	return _rp
}	
//QErrorMessage::qtHandler()
func (q *QErrorMessage) QtHandler() *QErrorMessage {
	var __rv uintptr
	q.Drv(233000,233106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QErrorMessage{}
	_rp.SetDriver(__rv,233000,false)
	return _rp
}	
//QErrorMessage::showMessage(QString const&)
func (q *QErrorMessage) ShowMessage(message string)  {
	q.Drv(233000,233107,unsafe.Pointer(&message),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QErrorMessage::showMessage(QString const&,QString const&)
func (q *QErrorMessage) ShowMessageWithMessageType(message string,_type string)  {
	q.Drv(233000,233108,unsafe.Pointer(&message),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
