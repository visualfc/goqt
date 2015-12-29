// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWidgetItem : QWidgetItem
type QWidgetItem struct {
	QLayoutItem
}
//QWidgetItem::QWidgetItem(QWidget*)	
func NewWidgetItem(w QWidgetInterface) *QWidgetItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),190000,190102,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWidgetItem{}
	_p.SetDriver(__rv,190000,true)
	return _p
} 
//QWidgetItem::expandingDirections()
func (q *QWidgetItem) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(190000,190103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetItem::geometry()
func (q *QWidgetItem) Geometry() *QRect {
	var __rv uintptr
	q.Drv(190000,190104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidgetItem::hasHeightForWidth()
func (q *QWidgetItem) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(190000,190105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetItem::heightForWidth(int)
func (q *QWidgetItem) HeightForWidth(value int) int {
	var __rv int
	q.Drv(190000,190106,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetItem::isEmpty()
func (q *QWidgetItem) IsEmpty() bool {
	var __rv bool
	q.Drv(190000,190107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetItem::maximumSize()
func (q *QWidgetItem) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(190000,190108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidgetItem::minimumSize()
func (q *QWidgetItem) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(190000,190109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidgetItem::setGeometry(QRect const&)
func (q *QWidgetItem) SetGeometry(value *QRect)  {
	q.Drv(190000,190110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidgetItem::sizeHint()
func (q *QWidgetItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(190000,190111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidgetItem::widget()
func (q *QWidgetItem) Widget() *QWidget {
	var __rv uintptr
	q.Drv(190000,190112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
