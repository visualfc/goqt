// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGenericReturnArgument : QGenericReturnArgument
type QGenericReturnArgument struct {
	QGenericArgument
}
//QGenericReturnArgument::QGenericReturnArgument()	
func NewGenericReturnArgument() *QGenericReturnArgument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),43000,43102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGenericReturnArgument{}
	_p.SetDriver(__rv,43000,true)
	return _p
} 
//QGenericReturnArgument::QGenericReturnArgument(char const*,void*)	
func NewGenericReturnArgumentWithAnameAdata(aName string,aData uintptr) *QGenericReturnArgument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),43000,43103,unsafe.Pointer(&aName),unsafe.Pointer(&aData),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGenericReturnArgument{}
	_p.SetDriver(__rv,43000,true)
	return _p
} 
