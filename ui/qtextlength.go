// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextLength_Type - QTextLength::Type
type QTextLength_Type uint32
const (
	QTextLength_VariableLength QTextLength_Type = 0
	QTextLength_FixedLength QTextLength_Type = 0
	QTextLength_PercentageLength QTextLength_Type = 1
)
//struct QTextLength : QTextLength
type QTextLength struct {
	BaseDrv
}
//QTextLength::QTextLength()	
func NewTextLength() *QTextLength {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),160000,160102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLength{}
	_p.SetDriver(__rv,160000,true)
	return _p
} 
//QTextLength::QTextLength(QTextLength::Type,double)	
func NewTextLengthWithTypeValue(_type QTextLength_Type,value float64) *QTextLength {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),160000,160103,unsafe.Pointer(&_type),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLength{}
	_p.SetDriver(__rv,160000,true)
	return _p
} 
//QTextLength::rawValue()
func (q *QTextLength) RawValue() float64 {
	var __rv float64
	q.Drv(160000,160104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLength::type()
func (q *QTextLength) Type() QTextLength_Type {
	var __rv QTextLength_Type
	q.Drv(160000,160105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLength::value(double)
func (q *QTextLength) Value(maximumLength float64) float64 {
	var __rv float64
	q.Drv(160000,160106,unsafe.Pointer(&maximumLength),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
