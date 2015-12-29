// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPainterPathElement : QPainterPath::Element
type QPainterPathElement struct {
	BaseDrv
}
//QPainterPath::Element::Element()	
func NewPainterPathElement() *QPainterPathElement {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),89000,89102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPathElement{}
	_p.SetDriver(__rv,89000,true)
	return _p
} 
//QPainterPath::Element::isCurveTo()
func (q *QPainterPathElement) IsCurveTo() bool {
	var __rv bool
	q.Drv(89000,89103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::Element::isLineTo()
func (q *QPainterPathElement) IsLineTo() bool {
	var __rv bool
	q.Drv(89000,89104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::Element::isMoveTo()
func (q *QPainterPathElement) IsMoveTo() bool {
	var __rv bool
	q.Drv(89000,89105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
