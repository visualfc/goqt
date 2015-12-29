// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStyleHintReturnMask_StyleOptionVersion - QStyleHintReturnMask::StyleOptionVersion
type QStyleHintReturnMask_StyleOptionVersion uint32
const (
	QStyleHintReturnMask_Version QStyleHintReturnMask_StyleOptionVersion = 1
)
//enum QStyleHintReturnMask_StyleOptionType - QStyleHintReturnMask::StyleOptionType
type QStyleHintReturnMask_StyleOptionType uint32
const (
	QStyleHintReturnMask_Type QStyleHintReturnMask_StyleOptionType = QStyleHintReturnMask_StyleOptionType(QStyleHintReturn_SH_Mask)
)
//struct QStyleHintReturnMask : QStyleHintReturnMask
type QStyleHintReturnMask struct {
	QStyleHintReturn
}
//QStyleHintReturnMask::QStyleHintReturnMask()	
func NewStyleHintReturnMask() *QStyleHintReturnMask {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),131000,131102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyleHintReturnMask{}
	_p.SetDriver(__rv,131000,true)
	return _p
} 
