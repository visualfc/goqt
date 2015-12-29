// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsLinearLayout : QGraphicsLinearLayout
type QGraphicsLinearLayout struct {
	QGraphicsLayout
}
//QGraphicsLinearLayout::QGraphicsLinearLayout()	
func NewGraphicsLinearLayout() *QGraphicsLinearLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),262000,262102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLinearLayout{}
	_p.SetDriver(__rv,262000,true)
	return _p
} 
//QGraphicsLinearLayout::QGraphicsLinearLayout(QGraphicsLayoutItem*)	
func NewGraphicsLinearLayoutWithParent(parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),262000,262103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLinearLayout{}
	_p.SetDriver(__rv,262000,true)
	return _p
} 
//QGraphicsLinearLayout::QGraphicsLinearLayout(Qt::Orientation,QGraphicsLayoutItem*)	
func NewGraphicsLinearLayoutWithOrientationParent(orientation Qt_Orientation,parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),262000,262104,unsafe.Pointer(&orientation),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLinearLayout{}
	_p.SetDriver(__rv,262000,true)
	return _p
} 
//QGraphicsLinearLayout::addItem(QGraphicsLayoutItem*)
func (q *QGraphicsLinearLayout) AddItem(item *QGraphicsLayoutItem)  {
	q.Drv(262000,262105,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::addStretch()
func (q *QGraphicsLinearLayout) AddStretch()  {
	q.Drv(262000,262106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::addStretch(int)
func (q *QGraphicsLinearLayout) AddStretchWithStretch(stretch int)  {
	q.Drv(262000,262107,unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::alignment(QGraphicsLayoutItem*)
func (q *QGraphicsLinearLayout) Alignment(item *QGraphicsLayoutItem) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(262000,262108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLinearLayout::count()
func (q *QGraphicsLinearLayout) Count() int {
	var __rv int
	q.Drv(262000,262109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLinearLayout::dump()
func (q *QGraphicsLinearLayout) Dump()  {
	q.Drv(262000,262110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::dump(int)
func (q *QGraphicsLinearLayout) DumpWithIndent(indent int)  {
	q.Drv(262000,262111,unsafe.Pointer(&indent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::insertItem(int,QGraphicsLayoutItem*)
func (q *QGraphicsLinearLayout) InsertItem(index int,item *QGraphicsLayoutItem)  {
	q.Drv(262000,262112,unsafe.Pointer(&index),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::insertStretch(int)
func (q *QGraphicsLinearLayout) InsertStretch(index int)  {
	q.Drv(262000,262113,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::insertStretch(int,int)
func (q *QGraphicsLinearLayout) InsertStretchWithIndexStretch(index int,stretch int)  {
	q.Drv(262000,262114,unsafe.Pointer(&index),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::invalidate()
func (q *QGraphicsLinearLayout) Invalidate()  {
	q.Drv(262000,262115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::itemAt(int)
func (q *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(262000,262116,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsLinearLayout::itemSpacing(int)
func (q *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	var __rv float64
	q.Drv(262000,262117,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLinearLayout::orientation()
func (q *QGraphicsLinearLayout) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(262000,262118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLinearLayout::removeAt(int)
func (q *QGraphicsLinearLayout) RemoveAt(index int)  {
	q.Drv(262000,262119,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::removeItem(QGraphicsLayoutItem*)
func (q *QGraphicsLinearLayout) RemoveItem(item *QGraphicsLayoutItem)  {
	q.Drv(262000,262120,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setAlignment(QGraphicsLayoutItem*,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsLinearLayout) SetAlignment(item *QGraphicsLayoutItem,alignment Qt_AlignmentFlag)  {
	q.Drv(262000,262121,Native(item),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setGeometry(QRectF const&)
func (q *QGraphicsLinearLayout) SetGeometry(rect *QRectF)  {
	q.Drv(262000,262122,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setItemSpacing(int,double)
func (q *QGraphicsLinearLayout) SetItemSpacing(index int,spacing float64)  {
	q.Drv(262000,262123,unsafe.Pointer(&index),unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setOrientation(Qt::Orientation)
func (q *QGraphicsLinearLayout) SetOrientation(orientation Qt_Orientation)  {
	q.Drv(262000,262124,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setSpacing(double)
func (q *QGraphicsLinearLayout) SetSpacing(spacing float64)  {
	q.Drv(262000,262125,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::setStretchFactor(QGraphicsLayoutItem*,int)
func (q *QGraphicsLinearLayout) SetStretchFactor(item *QGraphicsLayoutItem,stretch int)  {
	q.Drv(262000,262126,Native(item),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLinearLayout::sizeHint(Qt::SizeHint)
func (q *QGraphicsLinearLayout) SizeHint(which Qt_SizeHint) *QSizeF {
	var __rv uintptr
	q.Drv(262000,262127,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLinearLayout::sizeHint(Qt::SizeHint,QSizeF const&)
func (q *QGraphicsLinearLayout) SizeHintFWithWhichConstraint(which Qt_SizeHint,constraint *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(262000,262128,unsafe.Pointer(&which),Native(constraint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsLinearLayout::spacing()
func (q *QGraphicsLinearLayout) Spacing() float64 {
	var __rv float64
	q.Drv(262000,262129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLinearLayout::stretchFactor(QGraphicsLayoutItem*)
func (q *QGraphicsLinearLayout) StretchFactor(item *QGraphicsLayoutItem) int {
	var __rv int
	q.Drv(262000,262130,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
