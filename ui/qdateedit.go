// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDateEdit : QDateEdit
type QDateEdit struct {
	QDateTimeEdit
}
//QDateEdit::QDateEdit()	
func NewDateEdit() *QDateEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),222000,222102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateEdit{}
	_p.SetDriver(__rv,222000,false)
	return _p
} 
//QDateEdit::QDateEdit(QWidget*)	
func NewDateEditWithParent(parent QWidgetInterface) *QDateEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),222000,222103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateEdit{}
	_p.SetDriver(__rv,222000,false)
	return _p
} 
//QDateEdit::QDateEdit(QDate const&,QWidget*)	
func NewDateEditWithDateParent(date *QDate,parent QWidgetInterface) *QDateEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),222000,222104,Native(date),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDateEdit{}
	_p.SetDriver(__rv,222000,false)
	return _p
} 
