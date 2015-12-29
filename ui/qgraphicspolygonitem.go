// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsPolygonItem_enum_1 - QGraphicsPolygonItem::enum_1
type QGraphicsPolygonItem_enum_1 uint32
const (
	QGraphicsPolygonItem_Type QGraphicsPolygonItem_enum_1 = 5
)
//struct QGraphicsPolygonItem : QGraphicsPolygonItem
type QGraphicsPolygonItem struct {
	QAbstractGraphicsShapeItem
}
//QGraphicsPolygonItem::QGraphicsPolygonItem()	
func NewGraphicsPolygonItem() *QGraphicsPolygonItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),267000,267102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPolygonItem{}
	_p.SetDriver(__rv,267000,true)
	return _p
} 
//QGraphicsPolygonItem::QGraphicsPolygonItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPolygonItemWithParentScene(parent *QGraphicsItem) *QGraphicsPolygonItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),267000,267103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPolygonItem{}
	_p.SetDriver(__rv,267000,true)
	return _p
} 
//QGraphicsPolygonItem::QGraphicsPolygonItem(QPolygonF const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPolygonItemWithPolygonParentScene(polygon *QPolygonF,parent *QGraphicsItem) *QGraphicsPolygonItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),267000,267104,Native(polygon),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPolygonItem{}
	_p.SetDriver(__rv,267000,true)
	return _p
} 
//QGraphicsPolygonItem::boundingRect()
func (q *QGraphicsPolygonItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(267000,267105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsPolygonItem::contains(QPointF const&)
func (q *QGraphicsPolygonItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(267000,267106,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPolygonItem::fillRule()
func (q *QGraphicsPolygonItem) FillRule() Qt_FillRule {
	var __rv Qt_FillRule
	q.Drv(267000,267107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPolygonItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsPolygonItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(267000,267108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPolygonItem::opaqueArea()
func (q *QGraphicsPolygonItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(267000,267109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPolygonItem::polygon()
func (q *QGraphicsPolygonItem) Polygon() *QPolygonF {
	var __rv uintptr
	q.Drv(267000,267110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsPolygonItem::setFillRule(Qt::FillRule)
func (q *QGraphicsPolygonItem) SetFillRule(rule Qt_FillRule)  {
	q.Drv(267000,267111,unsafe.Pointer(&rule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPolygonItem::setPolygon(QPolygonF const&)
func (q *QGraphicsPolygonItem) SetPolygon(polygon *QPolygonF)  {
	q.Drv(267000,267112,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPolygonItem::shape()
func (q *QGraphicsPolygonItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(267000,267113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPolygonItem::type()
func (q *QGraphicsPolygonItem) Type() int {
	var __rv int
	q.Drv(267000,267114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
