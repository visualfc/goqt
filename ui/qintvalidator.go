// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QIntValidator : QIntValidator
type QIntValidator struct {
	QValidator
}
//QIntValidator::QIntValidator()	
func NewIntValidator() *QIntValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),294000,294102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIntValidator{}
	_p.SetDriver(__rv,294000,false)
	return _p
} 
//QIntValidator::QIntValidator(QObject*)	
func NewIntValidatorWithParent(parent QObjectInterface) *QIntValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),294000,294103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIntValidator{}
	_p.SetDriver(__rv,294000,false)
	return _p
} 
//QIntValidator::QIntValidator(int,int,QObject*)	
func NewIntValidatorWithBottomTopParent(bottom int,top int,parent QObjectInterface) *QIntValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),294000,294104,unsafe.Pointer(&bottom),unsafe.Pointer(&top),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QIntValidator{}
	_p.SetDriver(__rv,294000,false)
	return _p
} 
//QIntValidator::bottom()
func (q *QIntValidator) Bottom() int {
	var __rv int
	q.Drv(294000,294105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIntValidator::fixup(QString&)
func (q *QIntValidator) Fixup(input *string)  {
	q.Drv(294000,294106,unsafe.Pointer(&input),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIntValidator::setBottom(int)
func (q *QIntValidator) SetBottom(value int)  {
	q.Drv(294000,294107,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIntValidator::setRange(int,int)
func (q *QIntValidator) SetRange(bottom int,top int)  {
	q.Drv(294000,294108,unsafe.Pointer(&bottom),unsafe.Pointer(&top),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIntValidator::setTop(int)
func (q *QIntValidator) SetTop(value int)  {
	q.Drv(294000,294109,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIntValidator::top()
func (q *QIntValidator) Top() int {
	var __rv int
	q.Drv(294000,294110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIntValidator::validate(QString&,int&)
func (q *QIntValidator) Validate(value2 *string,value3 *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(294000,294111,unsafe.Pointer(&value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
