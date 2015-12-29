// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDoubleValidator_Notation - QDoubleValidator::Notation
type QDoubleValidator_Notation uint32
const (
	QDoubleValidator_StandardNotation QDoubleValidator_Notation = 0
	QDoubleValidator_ScientificNotation QDoubleValidator_Notation = 1
)
//struct QDoubleValidator : QDoubleValidator
type QDoubleValidator struct {
	QValidator
}
//QDoubleValidator::QDoubleValidator()	
func NewDoubleValidator() *QDoubleValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),231000,231102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDoubleValidator{}
	_p.SetDriver(__rv,231000,false)
	return _p
} 
//QDoubleValidator::QDoubleValidator(QObject*)	
func NewDoubleValidatorWithParent(parent QObjectInterface) *QDoubleValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),231000,231103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDoubleValidator{}
	_p.SetDriver(__rv,231000,false)
	return _p
} 
//QDoubleValidator::QDoubleValidator(double,double,int,QObject*)	
func NewDoubleValidatorWithBottomTopDecimalsParent(bottom float64,top float64,decimals int,parent QObjectInterface) *QDoubleValidator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),231000,231104,unsafe.Pointer(&bottom),unsafe.Pointer(&top),unsafe.Pointer(&decimals),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDoubleValidator{}
	_p.SetDriver(__rv,231000,false)
	return _p
} 
//QDoubleValidator::bottom()
func (q *QDoubleValidator) Bottom() float64 {
	var __rv float64
	q.Drv(231000,231105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleValidator::decimals()
func (q *QDoubleValidator) Decimals() int {
	var __rv int
	q.Drv(231000,231106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleValidator::notation()
func (q *QDoubleValidator) Notation() QDoubleValidator_Notation {
	var __rv QDoubleValidator_Notation
	q.Drv(231000,231107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleValidator::setBottom(double)
func (q *QDoubleValidator) SetBottom(value float64)  {
	q.Drv(231000,231108,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleValidator::setDecimals(int)
func (q *QDoubleValidator) SetDecimals(value int)  {
	q.Drv(231000,231109,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleValidator::setNotation(QDoubleValidator::Notation)
func (q *QDoubleValidator) SetNotation(value QDoubleValidator_Notation)  {
	q.Drv(231000,231110,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleValidator::setRange(double,double,int)
func (q *QDoubleValidator) SetRange(bottom float64,top float64,decimals int)  {
	q.Drv(231000,231111,unsafe.Pointer(&bottom),unsafe.Pointer(&top),unsafe.Pointer(&decimals),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleValidator::setTop(double)
func (q *QDoubleValidator) SetTop(value float64)  {
	q.Drv(231000,231112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleValidator::top()
func (q *QDoubleValidator) Top() float64 {
	var __rv float64
	q.Drv(231000,231113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleValidator::validate(QString&,int&)
func (q *QDoubleValidator) Validate(value2 *string,value3 *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(231000,231114,unsafe.Pointer(&value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
