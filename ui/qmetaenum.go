// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMetaEnum : QMetaEnum
type QMetaEnum struct {
	BaseDrv
}
//QMetaEnum::QMetaEnum()	
func NewMetaEnum() *QMetaEnum {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),75000,75102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMetaEnum{}
	_p.SetDriver(__rv,75000,true)
	return _p
} 
//QMetaEnum::enclosingMetaObject()
func (q *QMetaEnum) EnclosingMetaObject() *QMetaObject {
	var __rv uintptr
	q.Drv(75000,75103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaObject{}
	_rp.SetDriver(__rv,77000,true)
	return _rp
}	
//QMetaEnum::isFlag()
func (q *QMetaEnum) IsFlag() bool {
	var __rv bool
	q.Drv(75000,75104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::isValid()
func (q *QMetaEnum) IsValid() bool {
	var __rv bool
	q.Drv(75000,75105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::key(int)
func (q *QMetaEnum) Key(index int) string {
	var __rv string
	q.Drv(75000,75106,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::keyCount()
func (q *QMetaEnum) KeyCount() int {
	var __rv int
	q.Drv(75000,75107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::keyToValue(char const*)
func (q *QMetaEnum) KeyToValue(key string) int {
	var __rv int
	q.Drv(75000,75108,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::keysToValue(char const*)
func (q *QMetaEnum) KeysToValue(keys string) int {
	var __rv int
	q.Drv(75000,75109,unsafe.Pointer(&keys),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::name()
func (q *QMetaEnum) Name() string {
	var __rv string
	q.Drv(75000,75110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::scope()
func (q *QMetaEnum) Scope() string {
	var __rv string
	q.Drv(75000,75111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::value(int)
func (q *QMetaEnum) Value(index int) int {
	var __rv int
	q.Drv(75000,75112,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::valueToKey(int)
func (q *QMetaEnum) ValueToKey(value int) string {
	var __rv string
	q.Drv(75000,75113,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMetaEnum::valueToKeys(int)
func (q *QMetaEnum) ValueToKeys(value int) []byte {
	var __rv []byte
	q.Drv(75000,75114,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
