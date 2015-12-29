// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QRegExpValidator : QRegExpValidator
type QRegExpValidator struct {
	QValidator
}
//QRegExpValidator::QRegExpValidator()	
func NewRegExpValidator() *QRegExpValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),333000,333102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExpValidator{}
	_p.SetDriver(__rv,333000,false)
	return _p
} 
//QRegExpValidator::QRegExpValidator(QObject*)	
func NewRegExpValidatorWithParent(parent QObjectInterface) *QRegExpValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),333000,333103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExpValidator{}
	_p.SetDriver(__rv,333000,false)
	return _p
} 
//QRegExpValidator::QRegExpValidator(QRegExp const&,QObject*)	
func NewRegExpValidatorWithRxParent(rx *QRegExp,parent QObjectInterface) *QRegExpValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),333000,333104,Native(rx),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExpValidator{}
	_p.SetDriver(__rv,333000,false)
	return _p
} 
//QRegExpValidator::regExp()
func (q *QRegExpValidator) RegExp() *QRegExp {
	var __rv uintptr
	q.Drv(333000,333105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegExp{}
	_rp.SetDriver(__rv,112000,true)
	return _rp
}	
//QRegExpValidator::setRegExp(QRegExp const&)
func (q *QRegExpValidator) SetRegExp(rx *QRegExp)  {
	q.Drv(333000,333106,Native(rx),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegExpValidator::validate(QString&,int&)
func (q *QRegExpValidator) Validate(input *string,pos *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(333000,333107,unsafe.Pointer(&input),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
