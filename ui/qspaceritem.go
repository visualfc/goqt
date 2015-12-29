// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSpacerItem : QSpacerItem
type QSpacerItem struct {
	QLayoutItem
}
//QSpacerItem::QSpacerItem(int,int,QSizePolicy::Policy,QSizePolicy::Policy)	
func NewSpacerItem(w int,h int,hData QSizePolicy_Policy,vData QSizePolicy_Policy) *QSpacerItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),122000,122102,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&hData),unsafe.Pointer(&vData),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSpacerItem{}
	_p.SetDriver(__rv,122000,true)
	return _p
} 
//QSpacerItem::changeSize(int,int,QSizePolicy::Policy,QSizePolicy::Policy)
func (q *QSpacerItem) ChangeSize(w int,h int,hData QSizePolicy_Policy,vData QSizePolicy_Policy)  {
	q.Drv(122000,122103,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&hData),unsafe.Pointer(&vData),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpacerItem::expandingDirections()
func (q *QSpacerItem) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(122000,122104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpacerItem::geometry()
func (q *QSpacerItem) Geometry() *QRect {
	var __rv uintptr
	q.Drv(122000,122105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QSpacerItem::isEmpty()
func (q *QSpacerItem) IsEmpty() bool {
	var __rv bool
	q.Drv(122000,122106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpacerItem::maximumSize()
func (q *QSpacerItem) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(122000,122107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSpacerItem::minimumSize()
func (q *QSpacerItem) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(122000,122108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSpacerItem::setGeometry(QRect const&)
func (q *QSpacerItem) SetGeometry(value *QRect)  {
	q.Drv(122000,122109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpacerItem::sizeHint()
func (q *QSpacerItem) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(122000,122110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSpacerItem::spacerItem()
func (q *QSpacerItem) SpacerItem() *QSpacerItem {
	var __rv uintptr
	q.Drv(122000,122111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSpacerItem{}
	_rp.SetDriver(__rv,122000,true)
	return _rp
}	
