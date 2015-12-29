// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsRectItem_enum_1 - QGraphicsRectItem::enum_1
type QGraphicsRectItem_enum_1 uint32
const (
	QGraphicsRectItem_Type QGraphicsRectItem_enum_1 = 3
)
//struct QGraphicsRectItem : QGraphicsRectItem
type QGraphicsRectItem struct {
	QAbstractGraphicsShapeItem
}
//QGraphicsRectItem::QGraphicsRectItem()	
func NewGraphicsRectItem() *QGraphicsRectItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),269000,269102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRectItem{}
	_p.SetDriver(__rv,269000,true)
	return _p
} 
//QGraphicsRectItem::QGraphicsRectItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsRectItemWithParentScene(parent *QGraphicsItem) *QGraphicsRectItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),269000,269103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRectItem{}
	_p.SetDriver(__rv,269000,true)
	return _p
} 
//QGraphicsRectItem::QGraphicsRectItem(QRectF const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsRectItemWithRectParentScene(rect *QRectF,parent *QGraphicsItem) *QGraphicsRectItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),269000,269104,Native(rect),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRectItem{}
	_p.SetDriver(__rv,269000,true)
	return _p
} 
//QGraphicsRectItem::QGraphicsRectItem(double,double,double,double,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsRectItemWithXYWidthHeightParentScene(x float64,y float64,w float64,h float64,parent *QGraphicsItem) *QGraphicsRectItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),269000,269105,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(parent),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRectItem{}
	_p.SetDriver(__rv,269000,true)
	return _p
} 
//QGraphicsRectItem::boundingRect()
func (q *QGraphicsRectItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(269000,269106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsRectItem::contains(QPointF const&)
func (q *QGraphicsRectItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(269000,269107,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsRectItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsRectItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(269000,269108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsRectItem::opaqueArea()
func (q *QGraphicsRectItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(269000,269109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsRectItem::rect()
func (q *QGraphicsRectItem) Rect() *QRectF {
	var __rv uintptr
	q.Drv(269000,269110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsRectItem::setRect(QRectF const&)
func (q *QGraphicsRectItem) SetRect(rect *QRectF)  {
	q.Drv(269000,269111,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsRectItem::setRect(double,double,double,double)
func (q *QGraphicsRectItem) SetRectFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(269000,269112,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsRectItem::shape()
func (q *QGraphicsRectItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(269000,269113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsRectItem::type()
func (q *QGraphicsRectItem) Type() int {
	var __rv int
	q.Drv(269000,269114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
