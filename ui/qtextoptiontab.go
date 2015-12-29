// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextOptionTab : QTextOption::Tab
type QTextOptionTab struct {
	BaseDrv
}
//QTextOption::Tab::Tab()	
func NewTextOptionTab() *QTextOptionTab {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),165000,165102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextOptionTab{}
	_p.SetDriver(__rv,165000,true)
	return _p
} 
//QTextOption::Tab::Tab(double,QTextOption::TabType,QChar)	
func NewTextOptionTabWithPosTabtypeDelim(pos float64,tabType QTextOption_TabType,delim rune) *QTextOptionTab {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),165000,165103,unsafe.Pointer(&pos),unsafe.Pointer(&tabType),unsafe.Pointer(&delim),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextOptionTab{}
	_p.SetDriver(__rv,165000,true)
	return _p
} 
