// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractGraphicsShapeItem : QAbstractGraphicsShapeItem
type QAbstractGraphicsShapeItem struct {
	QGraphicsItem
}
//QAbstractGraphicsShapeItem::brush()
func (q *QAbstractGraphicsShapeItem) Brush() *QBrush {
	var __rv uintptr
	q.Drv(1000,1102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QAbstractGraphicsShapeItem::isObscuredBy(QGraphicsItem const*)
func (q *QAbstractGraphicsShapeItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(1000,1103,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractGraphicsShapeItem::opaqueArea()
func (q *QAbstractGraphicsShapeItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(1000,1104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QAbstractGraphicsShapeItem::pen()
func (q *QAbstractGraphicsShapeItem) Pen() *QPen {
	var __rv uintptr
	q.Drv(1000,1105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QAbstractGraphicsShapeItem::setBrush(QBrush const&)
func (q *QAbstractGraphicsShapeItem) SetBrush(brush *QBrush)  {
	q.Drv(1000,1106,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractGraphicsShapeItem::setPen(QPen const&)
func (q *QAbstractGraphicsShapeItem) SetPen(pen *QPen)  {
	q.Drv(1000,1107,Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
