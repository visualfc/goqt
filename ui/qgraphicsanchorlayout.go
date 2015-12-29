// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsAnchorLayout : QGraphicsAnchorLayout
type QGraphicsAnchorLayout struct {
	QGraphicsLayout
}
//QGraphicsAnchorLayout::QGraphicsAnchorLayout()	
func NewGraphicsAnchorLayout() *QGraphicsAnchorLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),249000,249102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsAnchorLayout{}
	_p.SetDriver(__rv,249000,true)
	return _p
} 
//QGraphicsAnchorLayout::QGraphicsAnchorLayout(QGraphicsLayoutItem*)	
func NewGraphicsAnchorLayoutWithParent(parent *QGraphicsLayoutItem) *QGraphicsAnchorLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),249000,249103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsAnchorLayout{}
	_p.SetDriver(__rv,249000,true)
	return _p
} 
//QGraphicsAnchorLayout::addAnchor(QGraphicsLayoutItem*,Qt::AnchorPoint,QGraphicsLayoutItem*,Qt::AnchorPoint)
func (q *QGraphicsAnchorLayout) AddAnchor(firstItem *QGraphicsLayoutItem,firstEdge Qt_AnchorPoint,secondItem *QGraphicsLayoutItem,secondEdge Qt_AnchorPoint) *QGraphicsAnchor {
	var __rv uintptr
	q.Drv(249000,249104,Native(firstItem),unsafe.Pointer(&firstEdge),Native(secondItem),unsafe.Pointer(&secondEdge),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsAnchor{}
	_rp.SetDriver(__rv,248000,false)
	return _rp
}	
//QGraphicsAnchorLayout::addAnchors(QGraphicsLayoutItem*,QGraphicsLayoutItem*,QFlags<Qt::Orientation>)
func (q *QGraphicsAnchorLayout) AddAnchors(firstItem *QGraphicsLayoutItem,secondItem *QGraphicsLayoutItem,orientations Qt_Orientation)  {
	q.Drv(249000,249105,Native(firstItem),Native(secondItem),unsafe.Pointer(&orientations),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::addCornerAnchors(QGraphicsLayoutItem*,Qt::Corner,QGraphicsLayoutItem*,Qt::Corner)
func (q *QGraphicsAnchorLayout) AddCornerAnchors(firstItem *QGraphicsLayoutItem,firstCorner Qt_Corner,secondItem *QGraphicsLayoutItem,secondCorner Qt_Corner)  {
	q.Drv(249000,249106,Native(firstItem),unsafe.Pointer(&firstCorner),Native(secondItem),unsafe.Pointer(&secondCorner),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::anchor(QGraphicsLayoutItem*,Qt::AnchorPoint,QGraphicsLayoutItem*,Qt::AnchorPoint)
func (q *QGraphicsAnchorLayout) Anchor(firstItem *QGraphicsLayoutItem,firstEdge Qt_AnchorPoint,secondItem *QGraphicsLayoutItem,secondEdge Qt_AnchorPoint) *QGraphicsAnchor {
	var __rv uintptr
	q.Drv(249000,249107,Native(firstItem),unsafe.Pointer(&firstEdge),Native(secondItem),unsafe.Pointer(&secondEdge),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsAnchor{}
	_rp.SetDriver(__rv,248000,false)
	return _rp
}	
//QGraphicsAnchorLayout::count()
func (q *QGraphicsAnchorLayout) Count() int {
	var __rv int
	q.Drv(249000,249108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsAnchorLayout::horizontalSpacing()
func (q *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	var __rv float64
	q.Drv(249000,249109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsAnchorLayout::invalidate()
func (q *QGraphicsAnchorLayout) Invalidate()  {
	q.Drv(249000,249110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::itemAt(int)
func (q *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(249000,249111,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsAnchorLayout::removeAt(int)
func (q *QGraphicsAnchorLayout) RemoveAt(index int)  {
	q.Drv(249000,249112,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::setGeometry(QRectF const&)
func (q *QGraphicsAnchorLayout) SetGeometry(rect *QRectF)  {
	q.Drv(249000,249113,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::setHorizontalSpacing(double)
func (q *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64)  {
	q.Drv(249000,249114,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::setSpacing(double)
func (q *QGraphicsAnchorLayout) SetSpacing(spacing float64)  {
	q.Drv(249000,249115,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::setVerticalSpacing(double)
func (q *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64)  {
	q.Drv(249000,249116,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsAnchorLayout::verticalSpacing()
func (q *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	var __rv float64
	q.Drv(249000,249117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
