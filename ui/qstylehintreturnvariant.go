// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStyleHintReturnVariant_StyleOptionVersion - QStyleHintReturnVariant::StyleOptionVersion
type QStyleHintReturnVariant_StyleOptionVersion uint32
const (
	QStyleHintReturnVariant_Version QStyleHintReturnVariant_StyleOptionVersion = 1
)
//enum QStyleHintReturnVariant_StyleOptionType - QStyleHintReturnVariant::StyleOptionType
type QStyleHintReturnVariant_StyleOptionType uint32
const (
	QStyleHintReturnVariant_Type QStyleHintReturnVariant_StyleOptionType = QStyleHintReturnVariant_StyleOptionType(QStyleHintReturn_SH_Variant)
)
//struct QStyleHintReturnVariant : QStyleHintReturnVariant
type QStyleHintReturnVariant struct {
	QStyleHintReturn
}
//QStyleHintReturnVariant::QStyleHintReturnVariant()	
func NewStyleHintReturnVariant() *QStyleHintReturnVariant {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),132000,132102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyleHintReturnVariant{}
	_p.SetDriver(__rv,132000,true)
	return _p
} 
