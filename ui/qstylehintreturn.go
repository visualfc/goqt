// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStyleHintReturn_HintReturnType - QStyleHintReturn::HintReturnType
type QStyleHintReturn_HintReturnType uint32
const (
	QStyleHintReturn_SH_Default QStyleHintReturn_HintReturnType = 0xf000
	QStyleHintReturn_SH_Mask QStyleHintReturn_HintReturnType = 0xf000+1
	QStyleHintReturn_SH_Variant QStyleHintReturn_HintReturnType = 0xf000+1+1
)
//enum QStyleHintReturn_StyleOptionVersion - QStyleHintReturn::StyleOptionVersion
type QStyleHintReturn_StyleOptionVersion uint32
const (
	QStyleHintReturn_Version QStyleHintReturn_StyleOptionVersion = 1
)
//enum QStyleHintReturn_StyleOptionType - QStyleHintReturn::StyleOptionType
type QStyleHintReturn_StyleOptionType uint32
const (
	QStyleHintReturn_Type QStyleHintReturn_StyleOptionType = QStyleHintReturn_StyleOptionType(QStyleHintReturn_SH_Default)
)
//struct QStyleHintReturn : QStyleHintReturn
type QStyleHintReturn struct {
	BaseDrv
}
//QStyleHintReturn::QStyleHintReturn()	
func NewStyleHintReturn() *QStyleHintReturn {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),130000,130102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyleHintReturn{}
	_p.SetDriver(__rv,130000,true)
	return _p
} 
//QStyleHintReturn::QStyleHintReturn(int,int)	
func NewStyleHintReturnWithVersionType(version int,_type int) *QStyleHintReturn {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),130000,130103,unsafe.Pointer(&version),unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyleHintReturn{}
	_p.SetDriver(__rv,130000,true)
	return _p
} 
