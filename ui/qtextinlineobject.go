// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextInlineObject : QTextInlineObject
type QTextInlineObject struct {
	BaseDrv
}
//QTextInlineObject::QTextInlineObject()	
func NewTextInlineObject() *QTextInlineObject {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),156000,156102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextInlineObject{}
	_p.SetDriver(__rv,156000,true)
	return _p
} 
//QTextInlineObject::ascent()
func (q *QTextInlineObject) Ascent() float64 {
	var __rv float64
	q.Drv(156000,156103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::descent()
func (q *QTextInlineObject) Descent() float64 {
	var __rv float64
	q.Drv(156000,156104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::format()
func (q *QTextInlineObject) Format() *QTextFormat {
	var __rv uintptr
	q.Drv(156000,156105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFormat{}
	_rp.SetDriver(__rv,151000,true)
	return _rp
}	
//QTextInlineObject::formatIndex()
func (q *QTextInlineObject) FormatIndex() int {
	var __rv int
	q.Drv(156000,156106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::height()
func (q *QTextInlineObject) Height() float64 {
	var __rv float64
	q.Drv(156000,156107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::isValid()
func (q *QTextInlineObject) IsValid() bool {
	var __rv bool
	q.Drv(156000,156108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::rect()
func (q *QTextInlineObject) Rect() *QRectF {
	var __rv uintptr
	q.Drv(156000,156109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTextInlineObject::setAscent(double)
func (q *QTextInlineObject) SetAscent(a float64)  {
	q.Drv(156000,156110,unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextInlineObject::setDescent(double)
func (q *QTextInlineObject) SetDescent(d float64)  {
	q.Drv(156000,156111,unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextInlineObject::setWidth(double)
func (q *QTextInlineObject) SetWidth(w float64)  {
	q.Drv(156000,156112,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextInlineObject::textDirection()
func (q *QTextInlineObject) TextDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(156000,156113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::textPosition()
func (q *QTextInlineObject) TextPosition() int {
	var __rv int
	q.Drv(156000,156114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextInlineObject::width()
func (q *QTextInlineObject) Width() float64 {
	var __rv float64
	q.Drv(156000,156115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
