// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGenericArgument : QGenericArgument
type QGenericArgument struct {
	BaseDrv
}
//QGenericArgument::QGenericArgument()	
func NewGenericArgument() *QGenericArgument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),42000,42102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGenericArgument{}
	_p.SetDriver(__rv,42000,true)
	return _p
} 
//QGenericArgument::QGenericArgument(char const*,void const*)	
func NewGenericArgumentWithAnameAdata(aName string,aData uintptr) *QGenericArgument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),42000,42103,unsafe.Pointer(&aName),unsafe.Pointer(&aData),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGenericArgument{}
	_p.SetDriver(__rv,42000,true)
	return _p
} 
//QGenericArgument::data()
func (q *QGenericArgument) Data() uintptr {
	var __rv uintptr
	q.Drv(42000,42104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGenericArgument::name()
func (q *QGenericArgument) Name() string {
	var __rv string
	q.Drv(42000,42105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
