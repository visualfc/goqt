// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsPixmapItem_ShapeMode - QGraphicsPixmapItem::ShapeMode
type QGraphicsPixmapItem_ShapeMode uint32
const (
	QGraphicsPixmapItem_MaskShape QGraphicsPixmapItem_ShapeMode = 0
	QGraphicsPixmapItem_BoundingRectShape QGraphicsPixmapItem_ShapeMode = 1
	QGraphicsPixmapItem_HeuristicMaskShape QGraphicsPixmapItem_ShapeMode = 2
)
//enum QGraphicsPixmapItem_enum_1 - QGraphicsPixmapItem::enum_1
type QGraphicsPixmapItem_enum_1 uint32
const (
	QGraphicsPixmapItem_Type QGraphicsPixmapItem_enum_1 = 7
)
//struct QGraphicsPixmapItem : QGraphicsPixmapItem
type QGraphicsPixmapItem struct {
	QGraphicsItem
}
//QGraphicsPixmapItem::QGraphicsPixmapItem()	
func NewGraphicsPixmapItem() *QGraphicsPixmapItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),266000,266102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPixmapItem{}
	_p.SetDriver(__rv,266000,true)
	return _p
} 
//QGraphicsPixmapItem::QGraphicsPixmapItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPixmapItemWithParentScene(parent *QGraphicsItem) *QGraphicsPixmapItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),266000,266103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPixmapItem{}
	_p.SetDriver(__rv,266000,true)
	return _p
} 
//QGraphicsPixmapItem::QGraphicsPixmapItem(QPixmap const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsPixmapItemWithPixmapParentScene(pixmap *QPixmap,parent *QGraphicsItem) *QGraphicsPixmapItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),266000,266104,Native(pixmap),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsPixmapItem{}
	_p.SetDriver(__rv,266000,true)
	return _p
} 
//QGraphicsPixmapItem::boundingRect()
func (q *QGraphicsPixmapItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(266000,266105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsPixmapItem::contains(QPointF const&)
func (q *QGraphicsPixmapItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(266000,266106,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPixmapItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsPixmapItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(266000,266107,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPixmapItem::offset()
func (q *QGraphicsPixmapItem) Offset() *QPointF {
	var __rv uintptr
	q.Drv(266000,266108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsPixmapItem::opaqueArea()
func (q *QGraphicsPixmapItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(266000,266109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPixmapItem::pixmap()
func (q *QGraphicsPixmapItem) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(266000,266110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QGraphicsPixmapItem::setOffset(QPointF const&)
func (q *QGraphicsPixmapItem) SetOffset(offset *QPointF)  {
	q.Drv(266000,266111,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPixmapItem::setOffset(double,double)
func (q *QGraphicsPixmapItem) SetOffsetFWithXY(x float64,y float64)  {
	q.Drv(266000,266112,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPixmapItem::setPixmap(QPixmap const&)
func (q *QGraphicsPixmapItem) SetPixmap(pixmap *QPixmap)  {
	q.Drv(266000,266113,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPixmapItem::setShapeMode(QGraphicsPixmapItem::ShapeMode)
func (q *QGraphicsPixmapItem) SetShapeMode(mode QGraphicsPixmapItem_ShapeMode)  {
	q.Drv(266000,266114,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPixmapItem::setTransformationMode(Qt::TransformationMode)
func (q *QGraphicsPixmapItem) SetTransformationMode(mode Qt_TransformationMode)  {
	q.Drv(266000,266115,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsPixmapItem::shape()
func (q *QGraphicsPixmapItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(266000,266116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsPixmapItem::shapeMode()
func (q *QGraphicsPixmapItem) ShapeMode() QGraphicsPixmapItem_ShapeMode {
	var __rv QGraphicsPixmapItem_ShapeMode
	q.Drv(266000,266117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPixmapItem::transformationMode()
func (q *QGraphicsPixmapItem) TransformationMode() Qt_TransformationMode {
	var __rv Qt_TransformationMode
	q.Drv(266000,266118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsPixmapItem::type()
func (q *QGraphicsPixmapItem) Type() int {
	var __rv int
	q.Drv(266000,266119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
