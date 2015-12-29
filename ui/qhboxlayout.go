// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QHBoxLayout : QHBoxLayout
type QHBoxLayout struct {
	QBoxLayout
}
//QHBoxLayout::QHBoxLayout()	
func NewHBoxLayout() *QHBoxLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),289000,289102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHBoxLayout{}
	_p.SetDriver(__rv,289000,false)
	return _p
} 
//QHBoxLayout::QHBoxLayout(QWidget*)	
func NewHBoxLayoutWithParent(parent QWidgetInterface) *QHBoxLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),289000,289103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QHBoxLayout{}
	_p.SetDriver(__rv,289000,false)
	return _p
} 
