// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsPathItem_enum_1 - QGraphicsPathItem::enum_1
type QGraphicsPathItem_enum_1 uint32
const (
	QGraphicsPathItem_Type QGraphicsPathItem_enum_1 = 2
)
//struct QGraphicsPathItem : QGraphicsPathItem
type QGraphicsPathItem struct {
	QAbstractGraphicsShapeItem
}
//QGraphicsPathItem::QGraphicsPathItem()	
func NewGraphicsPathItem() *QGraphicsPathItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),265000,265102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPathItem{}
	_p.SetDriver(__rv,265000,true)
	return _p
} 
//QGraphicsPathItem::QGraphicsPathItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPathItemWithParentScene(parent *QGraphicsItem) *QGraphicsPathItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),265000,265103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPathItem{}
	_p.SetDriver(__rv,265000,true)
	return _p
} 
//QGraphicsPathItem::QGraphicsPathItem(QPainterPath const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPathItemWithPathParentScene(path *QPainterPath,parent *QGraphicsItem) *QGraphicsPathItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),265000,265104,Native(path),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPathItem{}
	_p.SetDriver(__rv,265000,true)
	return _p
} 
//QGraphicsPathItem::boundingRect()
func (q *QGraphicsPathItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(265000,265105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsPathItem::contains(QPointF const&)
func (q *QGraphicsPathItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(265000,265106,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPathItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsPathItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(265000,265107,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPathItem::opaqueArea()
func (q *QGraphicsPathItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(265000,265108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPathItem::path()
func (q *QGraphicsPathItem) Path() *QPainterPath {
	var __rv uintptr
	q.Drv(265000,265109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPathItem::setPath(QPainterPath const&)
func (q *QGraphicsPathItem) SetPath(path *QPainterPath)  {
	q.Drv(265000,265110,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPathItem::shape()
func (q *QGraphicsPathItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(265000,265111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPathItem::type()
func (q *QGraphicsPathItem) Type() int {
	var __rv int
	q.Drv(265000,265112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
