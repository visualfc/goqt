// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QUuid_Variant - QUuid::Variant
type QUuid_Variant int32
const (
	QUuid_VarUnknown QUuid_Variant = -1
	QUuid_NCS QUuid_Variant = 0
	QUuid_DCE QUuid_Variant = 2
	QUuid_Microsoft QUuid_Variant = 6
	QUuid_Reserved QUuid_Variant = 7
)
//enum QUuid_Version - QUuid::Version
type QUuid_Version int32
const (
	QUuid_VerUnknown QUuid_Version = -1
	QUuid_Time QUuid_Version = 1
	QUuid_EmbeddedPOSIX QUuid_Version = 2
	QUuid_Name QUuid_Version = 3
	QUuid_Random QUuid_Version = 4
)
//struct QUuid : QUuid
type QUuid struct {
	BaseDrv
}
//QUuid::QUuid()	
func NewUuid() *QUuid {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),181000,181102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUuid{}
	_p.SetDriver(__rv,181000,true)
	return _p
} 
//QUuid::QUuid(QString const&)	
func NewUuidWithString(value string) *QUuid {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),181000,181103,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUuid{}
	_p.SetDriver(__rv,181000,true)
	return _p
} 
//QUuid::QUuid(unsigned int,unsigned short,unsigned short,unsigned char,unsigned char,unsigned char,unsigned char,unsigned char,unsigned char,unsigned char,unsigned char)	
func NewUuidWithLW1W2B1B2B3B4B5B6B7B8(l uint,w1 uint16,w2 uint16,b1 byte,b2 byte,b3 byte,b4 byte,b5 byte,b6 byte,b7 byte,b8 byte) *QUuid {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),181000,181104,unsafe.Pointer(&l),unsafe.Pointer(&w1),unsafe.Pointer(&w2),unsafe.Pointer(&b1),unsafe.Pointer(&b2),unsafe.Pointer(&b3),unsafe.Pointer(&b4),unsafe.Pointer(&b5),unsafe.Pointer(&b6),unsafe.Pointer(&b7),unsafe.Pointer(&b8),nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUuid{}
	_p.SetDriver(__rv,181000,true)
	return _p
} 
//QUuid::createUuid()	
func QUuidCreateUuid() *QUuid {
	var __rv uintptr
	DirectQtDrv(nil,181000,181105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUuid{}
	_rp.SetDriver(__rv,181000,true)
	return _rp
}	
//QUuid::createUuid()
func (q *QUuid) CreateUuid() *QUuid {
	var __rv uintptr
	q.Drv(181000,181105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUuid{}
	_rp.SetDriver(__rv,181000,true)
	return _rp
}	
//QUuid::isNull()
func (q *QUuid) IsNull() bool {
	var __rv bool
	q.Drv(181000,181106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUuid::toString()
func (q *QUuid) ToString() string {
	var __rv string
	q.Drv(181000,181107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUuid::variant()
func (q *QUuid) Variant() QUuid_Variant {
	var __rv QUuid_Variant
	q.Drv(181000,181108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUuid::version()
func (q *QUuid) Version() QUuid_Version {
	var __rv QUuid_Version
	q.Drv(181000,181109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
