// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMetaMethod_Access - QMetaMethod::Access
type QMetaMethod_Access uint32
const (
	QMetaMethod_Private QMetaMethod_Access = 0
	QMetaMethod_Protected QMetaMethod_Access = 1
	QMetaMethod_Public QMetaMethod_Access = 2
)
//enum QMetaMethod_Attributes - QMetaMethod::Attributes
type QMetaMethod_Attributes uint32
const (
	QMetaMethod_Compatibility QMetaMethod_Attributes = 0x1
	QMetaMethod_Cloned QMetaMethod_Attributes = 0x2
	QMetaMethod_Scriptable QMetaMethod_Attributes = 0x4
)
//enum QMetaMethod_MethodType - QMetaMethod::MethodType
type QMetaMethod_MethodType uint32
const (
	QMetaMethod_Method QMetaMethod_MethodType = 0
	QMetaMethod_Signal QMetaMethod_MethodType = 1
	QMetaMethod_Slot QMetaMethod_MethodType = 2
	QMetaMethod_Constructor QMetaMethod_MethodType = 3
)
//struct QMetaMethod : QMetaMethod
type QMetaMethod struct {
	BaseDrv
}
//QMetaMethod::QMetaMethod()	
func NewMetaMethod() *QMetaMethod {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),76000,76102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMetaMethod{}
	_p.SetDriver(__rv,76000,true)
	return _p
} 
//QMetaMethod::access()
func (q *QMetaMethod) Access() QMetaMethod_Access {
	var __rv QMetaMethod_Access
	q.Drv(76000,76103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::attributes()
func (q *QMetaMethod) Attributes() int {
	var __rv int
	q.Drv(76000,76104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::enclosingMetaObject()
func (q *QMetaMethod) EnclosingMetaObject() *QMetaObject {
	var __rv uintptr
	q.Drv(76000,76105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaObject{}
	_rp.SetDriver(__rv,77000,true)
	return _rp
}	
//QMetaMethod::methodIndex()
func (q *QMetaMethod) MethodIndex() int {
	var __rv int
	q.Drv(76000,76106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::methodType()
func (q *QMetaMethod) MethodType() QMetaMethod_MethodType {
	var __rv QMetaMethod_MethodType
	q.Drv(76000,76107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::parameterNames()
func (q *QMetaMethod) ParameterNames() [][]byte {
	var __rv [][]byte
	q.Drv(76000,76108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::parameterTypes()
func (q *QMetaMethod) ParameterTypes() [][]byte {
	var __rv [][]byte
	q.Drv(76000,76109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::signature()
func (q *QMetaMethod) MethodSignature() []byte {
	var __rv []byte
	q.Drv(76000,76110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::tag()
func (q *QMetaMethod) Tag() string {
	var __rv string
	q.Drv(76000,76111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaMethod::typeName()
func (q *QMetaMethod) TypeName() string {
	var __rv string
	q.Drv(76000,76112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
