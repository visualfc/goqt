// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTimeEdit : QTimeEdit
type QTimeEdit struct {
	QDateTimeEdit
}
//QTimeEdit::QTimeEdit()	
func NewTimeEdit() *QTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),378000,378102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimeEdit{}
	_p.SetDriver(__rv,378000,false)
	return _p
} 
//QTimeEdit::QTimeEdit(QWidget*)	
func NewTimeEditWithParent(parent QWidgetInterface) *QTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),378000,378103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimeEdit{}
	_p.SetDriver(__rv,378000,false)
	return _p
} 
//QTimeEdit::QTimeEdit(QTime const&,QWidget*)	
func NewTimeEditWithTimeParent(time *QTime,parent QWidgetInterface) *QTimeEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),378000,378104,Native(time),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimeEdit{}
	_p.SetDriver(__rv,378000,false)
	return _p
} 
