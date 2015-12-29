// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsSimpleTextItem_enum_1 - QGraphicsSimpleTextItem::enum_1
type QGraphicsSimpleTextItem_enum_1 uint32
const (
	QGraphicsSimpleTextItem_Type QGraphicsSimpleTextItem_enum_1 = 9
)
//struct QGraphicsSimpleTextItem : QGraphicsSimpleTextItem
type QGraphicsSimpleTextItem struct {
	QAbstractGraphicsShapeItem
}
//QGraphicsSimpleTextItem::QGraphicsSimpleTextItem()	
func NewGraphicsSimpleTextItem() *QGraphicsSimpleTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),282000,282102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSimpleTextItem{}
	_p.SetDriver(__rv,282000,true)
	return _p
} 
//QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsSimpleTextItemWithParentScene(parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),282000,282103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSimpleTextItem{}
	_p.SetDriver(__rv,282000,true)
	return _p
} 
//QGraphicsSimpleTextItem::QGraphicsSimpleTextItem(QString const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsSimpleTextItemWithTextParentScene(text string,parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),282000,282104,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsSimpleTextItem{}
	_p.SetDriver(__rv,282000,true)
	return _p
} 
//QGraphicsSimpleTextItem::boundingRect()
func (q *QGraphicsSimpleTextItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(282000,282105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsSimpleTextItem::contains(QPointF const&)
func (q *QGraphicsSimpleTextItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(282000,282106,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSimpleTextItem::font()
func (q *QGraphicsSimpleTextItem) Font() *QFont {
	var __rv uintptr
	q.Drv(282000,282107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QGraphicsSimpleTextItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsSimpleTextItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(282000,282108,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSimpleTextItem::opaqueArea()
func (q *QGraphicsSimpleTextItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(282000,282109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsSimpleTextItem::setFont(QFont const&)
func (q *QGraphicsSimpleTextItem) SetFont(font *QFont)  {
	q.Drv(282000,282110,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSimpleTextItem::setText(QString const&)
func (q *QGraphicsSimpleTextItem) SetText(text string)  {
	q.Drv(282000,282111,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsSimpleTextItem::shape()
func (q *QGraphicsSimpleTextItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(282000,282112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsSimpleTextItem::text()
func (q *QGraphicsSimpleTextItem) Text() string {
	var __rv string
	q.Drv(282000,282113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsSimpleTextItem::type()
func (q *QGraphicsSimpleTextItem) Type() int {
	var __rv int
	q.Drv(282000,282114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
