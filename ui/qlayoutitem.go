// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QLayoutItem : QLayoutItem
type QLayoutItem struct {
	BaseDrv
}
//QLayoutItem::alignment()
func (q *QLayoutItem) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(66000,66102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::controlTypes()
func (q *QLayoutItem) ControlTypes() QSizePolicy_ControlType {
	var __rv QSizePolicy_ControlType
	q.Drv(66000,66103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::expandingDirections()
func (q *QLayoutItem) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(66000,66104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::geometry()
func (q *QLayoutItem) Geometry() *QRect {
	var __rv uintptr
	q.Drv(66000,66105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QLayoutItem::hasHeightForWidth()
func (q *QLayoutItem) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(66000,66106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::heightForWidth(int)
func (q *QLayoutItem) HeightForWidth(value int) int {
	var __rv int
	q.Drv(66000,66107,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::invalidate()
func (q *QLayoutItem) Invalidate()  {
	q.Drv(66000,66108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayoutItem::isEmpty()
func (q *QLayoutItem) IsEmpty() bool {
	var __rv bool
	q.Drv(66000,66109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::layout()
func (q *QLayoutItem) Layout() *QLayout {
	var __rv uintptr
	q.Drv(66000,66110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayout{}
	_rp.SetDriver(__rv,300000,false)
	return _rp
}	
//QLayoutItem::maximumSize()
func (q *QLayoutItem) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(66000,66111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayoutItem::minimumHeightForWidth(int)
func (q *QLayoutItem) MinimumHeightForWidth(value int) int {
	var __rv int
	q.Drv(66000,66112,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayoutItem::minimumSize()
func (q *QLayoutItem) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(66000,66113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayoutItem::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QLayoutItem) SetAlignment(a Qt_AlignmentFlag)  {
	q.Drv(66000,66114,unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayoutItem::setGeometry(QRect const&)
func (q *QLayoutItem) SetGeometry(value *QRect)  {
	q.Drv(66000,66115,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayoutItem::sizeHint()
func (q *QLayoutItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(66000,66116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayoutItem::spacerItem()
func (q *QLayoutItem) SpacerItem() *QSpacerItem {
	var __rv uintptr
	q.Drv(66000,66117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSpacerItem{}
	_rp.SetDriver(__rv,122000,true)
	return _rp
}	
//QLayoutItem::widget()
func (q *QLayoutItem) Widget() *QWidget {
	var __rv uintptr
	q.Drv(66000,66118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
