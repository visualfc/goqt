// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsItemGroup_enum_1 - QGraphicsItemGroup::enum_1
type QGraphicsItemGroup_enum_1 uint32
const (
	QGraphicsItemGroup_Type QGraphicsItemGroup_enum_1 = 10
)
//struct QGraphicsItemGroup : QGraphicsItemGroup
type QGraphicsItemGroup struct {
	QGraphicsItem
}
//QGraphicsItemGroup::QGraphicsItemGroup()	
func NewGraphicsItemGroup() *QGraphicsItemGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),258000,258102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsItemGroup{}
	_p.SetDriver(__rv,258000,true)
	return _p
} 
//QGraphicsItemGroup::QGraphicsItemGroup(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsItemGroupWithParentScene(parent *QGraphicsItem) *QGraphicsItemGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),258000,258103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsItemGroup{}
	_p.SetDriver(__rv,258000,true)
	return _p
} 
//QGraphicsItemGroup::addToGroup(QGraphicsItem*)
func (q *QGraphicsItemGroup) AddToGroup(item *QGraphicsItem)  {
	q.Drv(258000,258104,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemGroup::boundingRect()
func (q *QGraphicsItemGroup) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(258000,258105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItemGroup::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsItemGroup) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(258000,258106,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemGroup::opaqueArea()
func (q *QGraphicsItemGroup) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(258000,258107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItemGroup::removeFromGroup(QGraphicsItem*)
func (q *QGraphicsItemGroup) RemoveFromGroup(item *QGraphicsItem)  {
	q.Drv(258000,258108,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemGroup::type()
func (q *QGraphicsItemGroup) Type() int {
	var __rv int
	q.Drv(258000,258109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
