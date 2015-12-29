// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVBoxLayout : QVBoxLayout
type QVBoxLayout struct {
	QBoxLayout
}
//QVBoxLayout::QVBoxLayout()	
func NewVBoxLayout() *QVBoxLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),392000,392102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVBoxLayout{}
	_p.SetDriver(__rv,392000,false)
	return _p
} 
//QVBoxLayout::QVBoxLayout(QWidget*)	
func NewVBoxLayoutWithParent(parent QWidgetInterface) *QVBoxLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),392000,392103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVBoxLayout{}
	_p.SetDriver(__rv,392000,false)
	return _p
} 
