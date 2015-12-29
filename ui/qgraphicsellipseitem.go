// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsEllipseItem_enum_1 - QGraphicsEllipseItem::enum_1
type QGraphicsEllipseItem_enum_1 uint32
const (
	QGraphicsEllipseItem_Type QGraphicsEllipseItem_enum_1 = 4
)
//struct QGraphicsEllipseItem : QGraphicsEllipseItem
type QGraphicsEllipseItem struct {
	QAbstractGraphicsShapeItem
}
//QGraphicsEllipseItem::QGraphicsEllipseItem()	
func NewGraphicsEllipseItem() *QGraphicsEllipseItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),254000,254102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsEllipseItem{}
	_p.SetDriver(__rv,254000,true)
	return _p
} 
//QGraphicsEllipseItem::QGraphicsEllipseItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsEllipseItemWithParentScene(parent *QGraphicsItem) *QGraphicsEllipseItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),254000,254103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsEllipseItem{}
	_p.SetDriver(__rv,254000,true)
	return _p
} 
//QGraphicsEllipseItem::QGraphicsEllipseItem(QRectF const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsEllipseItemWithRectParentScene(rect *QRectF,parent *QGraphicsItem) *QGraphicsEllipseItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),254000,254104,Native(rect),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsEllipseItem{}
	_p.SetDriver(__rv,254000,true)
	return _p
} 
//QGraphicsEllipseItem::QGraphicsEllipseItem(double,double,double,double,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsEllipseItemWithXYWidthHeightParentScene(x float64,y float64,w float64,h float64,parent *QGraphicsItem) *QGraphicsEllipseItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),254000,254105,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(parent),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsEllipseItem{}
	_p.SetDriver(__rv,254000,true)
	return _p
} 
//QGraphicsEllipseItem::boundingRect()
func (q *QGraphicsEllipseItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(254000,254106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsEllipseItem::contains(QPointF const&)
func (q *QGraphicsEllipseItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(254000,254107,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEllipseItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsEllipseItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(254000,254108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEllipseItem::opaqueArea()
func (q *QGraphicsEllipseItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(254000,254109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsEllipseItem::rect()
func (q *QGraphicsEllipseItem) Rect() *QRectF {
	var __rv uintptr
	q.Drv(254000,254110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsEllipseItem::setRect(QRectF const&)
func (q *QGraphicsEllipseItem) SetRect(rect *QRectF)  {
	q.Drv(254000,254111,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEllipseItem::setRect(double,double,double,double)
func (q *QGraphicsEllipseItem) SetRectFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(254000,254112,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEllipseItem::setSpanAngle(int)
func (q *QGraphicsEllipseItem) SetSpanAngle(angle int)  {
	q.Drv(254000,254113,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEllipseItem::setStartAngle(int)
func (q *QGraphicsEllipseItem) SetStartAngle(angle int)  {
	q.Drv(254000,254114,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEllipseItem::shape()
func (q *QGraphicsEllipseItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(254000,254115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsEllipseItem::spanAngle()
func (q *QGraphicsEllipseItem) SpanAngle() int {
	var __rv int
	q.Drv(254000,254116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEllipseItem::startAngle()
func (q *QGraphicsEllipseItem) StartAngle() int {
	var __rv int
	q.Drv(254000,254117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEllipseItem::type()
func (q *QGraphicsEllipseItem) Type() int {
	var __rv int
	q.Drv(254000,254118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
