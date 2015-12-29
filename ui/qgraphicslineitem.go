// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsLineItem_enum_1 - QGraphicsLineItem::enum_1
type QGraphicsLineItem_enum_1 uint32
const (
	QGraphicsLineItem_Type QGraphicsLineItem_enum_1 = 6
)
//struct QGraphicsLineItem : QGraphicsLineItem
type QGraphicsLineItem struct {
	QGraphicsItem
}
//QGraphicsLineItem::QGraphicsLineItem()	
func NewGraphicsLineItem() *QGraphicsLineItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),261000,261102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLineItem{}
	_p.SetDriver(__rv,261000,true)
	return _p
} 
//QGraphicsLineItem::QGraphicsLineItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsLineItemWithParentScene(parent *QGraphicsItem) *QGraphicsLineItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),261000,261103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLineItem{}
	_p.SetDriver(__rv,261000,true)
	return _p
} 
//QGraphicsLineItem::QGraphicsLineItem(QLineF const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsLineItemWithLineParentScene(line *QLineF,parent *QGraphicsItem) *QGraphicsLineItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),261000,261104,Native(line),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLineItem{}
	_p.SetDriver(__rv,261000,true)
	return _p
} 
//QGraphicsLineItem::QGraphicsLineItem(double,double,double,double,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsLineItemWithX1Y1X2Y2ParentScene(x1 float64,y1 float64,x2 float64,y2 float64,parent *QGraphicsItem) *QGraphicsLineItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),261000,261105,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),Native(parent),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsLineItem{}
	_p.SetDriver(__rv,261000,true)
	return _p
} 
//QGraphicsLineItem::boundingRect()
func (q *QGraphicsLineItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(261000,261106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsLineItem::contains(QPointF const&)
func (q *QGraphicsLineItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(261000,261107,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLineItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsLineItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(261000,261108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLineItem::line()
func (q *QGraphicsLineItem) Line() *QLineF {
	var __rv uintptr
	q.Drv(261000,261109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QGraphicsLineItem::opaqueArea()
func (q *QGraphicsLineItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(261000,261110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsLineItem::pen()
func (q *QGraphicsLineItem) Pen() *QPen {
	var __rv uintptr
	q.Drv(261000,261111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QGraphicsLineItem::setLine(QLineF const&)
func (q *QGraphicsLineItem) SetLine(line *QLineF)  {
	q.Drv(261000,261112,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLineItem::setLine(double,double,double,double)
func (q *QGraphicsLineItem) SetLineFWithX1Y1X2Y2(x1 float64,y1 float64,x2 float64,y2 float64)  {
	q.Drv(261000,261113,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLineItem::setPen(QPen const&)
func (q *QGraphicsLineItem) SetPen(pen *QPen)  {
	q.Drv(261000,261114,Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLineItem::shape()
func (q *QGraphicsLineItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(261000,261115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsLineItem::type()
func (q *QGraphicsLineItem) Type() int {
	var __rv int
	q.Drv(261000,261116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
