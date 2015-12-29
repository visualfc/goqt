// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsLayoutItem : QGraphicsLayoutItem
type QGraphicsLayoutItem struct {
	BaseDrv
}
//QGraphicsLayoutItem::contentsRect()
func (q *QGraphicsLayoutItem) ContentsRect() *QRectF {
	var __rv uintptr
	q.Drv(260000,260102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsLayoutItem::effectiveSizeHint(Qt::SizeHint)
func (q *QGraphicsLayoutItem) EffectiveSizeHint(which Qt_SizeHint) *QSizeF {
	var __rv uintptr
	q.Drv(260000,260103,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLayoutItem::effectiveSizeHint(Qt::SizeHint,QSizeF const&)
func (q *QGraphicsLayoutItem) EffectiveSizeHintFWithWhichConstraint(which Qt_SizeHint,constraint *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(260000,260104,unsafe.Pointer(&which),Native(constraint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLayoutItem::geometry()
func (q *QGraphicsLayoutItem) Geometry() *QRectF {
	var __rv uintptr
	q.Drv(260000,260105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsLayoutItem::getContentsMargins(double*,double*,double*,double*)
func (q *QGraphicsLayoutItem) GetContentsMargins(left *float64,top *float64,right *float64,bottom *float64)  {
	q.Drv(260000,260106,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::graphicsItem()
func (q *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(260000,260107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsLayoutItem::isLayout()
func (q *QGraphicsLayoutItem) IsLayout() bool {
	var __rv bool
	q.Drv(260000,260108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::maximumHeight()
func (q *QGraphicsLayoutItem) MaximumHeight() float64 {
	var __rv float64
	q.Drv(260000,260109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::maximumSize()
func (q *QGraphicsLayoutItem) MaximumSize() *QSizeF {
	var __rv uintptr
	q.Drv(260000,260110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLayoutItem::maximumWidth()
func (q *QGraphicsLayoutItem) MaximumWidth() float64 {
	var __rv float64
	q.Drv(260000,260111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::minimumHeight()
func (q *QGraphicsLayoutItem) MinimumHeight() float64 {
	var __rv float64
	q.Drv(260000,260112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::minimumSize()
func (q *QGraphicsLayoutItem) MinimumSize() *QSizeF {
	var __rv uintptr
	q.Drv(260000,260113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLayoutItem::minimumWidth()
func (q *QGraphicsLayoutItem) MinimumWidth() float64 {
	var __rv float64
	q.Drv(260000,260114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::ownedByLayout()
func (q *QGraphicsLayoutItem) OwnedByLayout() bool {
	var __rv bool
	q.Drv(260000,260115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::parentLayoutItem()
func (q *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(260000,260116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsLayoutItem::preferredHeight()
func (q *QGraphicsLayoutItem) PreferredHeight() float64 {
	var __rv float64
	q.Drv(260000,260117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::preferredSize()
func (q *QGraphicsLayoutItem) PreferredSize() *QSizeF {
	var __rv uintptr
	q.Drv(260000,260118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLayoutItem::preferredWidth()
func (q *QGraphicsLayoutItem) PreferredWidth() float64 {
	var __rv float64
	q.Drv(260000,260119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayoutItem::setGeometry(QRectF const&)
func (q *QGraphicsLayoutItem) SetGeometry(rect *QRectF)  {
	q.Drv(260000,260120,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMaximumHeight(double)
func (q *QGraphicsLayoutItem) SetMaximumHeight(height float64)  {
	q.Drv(260000,260121,unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMaximumSize(QSizeF const&)
func (q *QGraphicsLayoutItem) SetMaximumSize(size *QSizeF)  {
	q.Drv(260000,260122,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMaximumSize(double,double)
func (q *QGraphicsLayoutItem) SetMaximumSizeFWithWidthHeight(w float64,h float64)  {
	q.Drv(260000,260123,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMaximumWidth(double)
func (q *QGraphicsLayoutItem) SetMaximumWidth(width float64)  {
	q.Drv(260000,260124,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMinimumHeight(double)
func (q *QGraphicsLayoutItem) SetMinimumHeight(height float64)  {
	q.Drv(260000,260125,unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMinimumSize(QSizeF const&)
func (q *QGraphicsLayoutItem) SetMinimumSize(size *QSizeF)  {
	q.Drv(260000,260126,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMinimumSize(double,double)
func (q *QGraphicsLayoutItem) SetMinimumSizeFWithWidthHeight(w float64,h float64)  {
	q.Drv(260000,260127,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setMinimumWidth(double)
func (q *QGraphicsLayoutItem) SetMinimumWidth(width float64)  {
	q.Drv(260000,260128,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setParentLayoutItem(QGraphicsLayoutItem*)
func (q *QGraphicsLayoutItem) SetParentLayoutItem(parent *QGraphicsLayoutItem)  {
	q.Drv(260000,260129,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setPreferredHeight(double)
func (q *QGraphicsLayoutItem) SetPreferredHeight(height float64)  {
	q.Drv(260000,260130,unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setPreferredSize(QSizeF const&)
func (q *QGraphicsLayoutItem) SetPreferredSize(size *QSizeF)  {
	q.Drv(260000,260131,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setPreferredSize(double,double)
func (q *QGraphicsLayoutItem) SetPreferredSizeFWithWidthHeight(w float64,h float64)  {
	q.Drv(260000,260132,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setPreferredWidth(double)
func (q *QGraphicsLayoutItem) SetPreferredWidth(width float64)  {
	q.Drv(260000,260133,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setSizePolicy(QSizePolicy const&)
func (q *QGraphicsLayoutItem) SetSizePolicy(policy *QSizePolicy)  {
	q.Drv(260000,260134,Native(policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::setSizePolicy(QSizePolicy::Policy,QSizePolicy::Policy,QSizePolicy::ControlType)
func (q *QGraphicsLayoutItem) SetSizePolicyWithHpolicyVpolicyControltype(hPolicy QSizePolicy_Policy,vPolicy QSizePolicy_Policy,controlType QSizePolicy_ControlType)  {
	q.Drv(260000,260135,unsafe.Pointer(&hPolicy),unsafe.Pointer(&vPolicy),unsafe.Pointer(&controlType),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayoutItem::sizePolicy()
func (q *QGraphicsLayoutItem) SizePolicy() *QSizePolicy {
	var __rv uintptr
	q.Drv(260000,260136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizePolicy{}
	_rp.SetDriver(__rv,121000,true)
	return _rp
}	
//QGraphicsLayoutItem::updateGeometry()
func (q *QGraphicsLayoutItem) UpdateGeometry()  {
	q.Drv(260000,260137,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
