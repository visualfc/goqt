// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsGridLayout : QGraphicsGridLayout
type QGraphicsGridLayout struct {
	QGraphicsLayout
}
//QGraphicsGridLayout::QGraphicsGridLayout()	
func NewGraphicsGridLayout() *QGraphicsGridLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),255000,255102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsGridLayout{}
	_p.SetDriver(__rv,255000,true)
	return _p
} 
//QGraphicsGridLayout::QGraphicsGridLayout(QGraphicsLayoutItem*)	
func NewGraphicsGridLayoutWithParent(parent *QGraphicsLayoutItem) *QGraphicsGridLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),255000,255103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsGridLayout{}
	_p.SetDriver(__rv,255000,true)
	return _p
} 
//QGraphicsGridLayout::addItem(QGraphicsLayoutItem*,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsGridLayout) AddItemWithItemRowColumnAlignment(item *QGraphicsLayoutItem,row int,column int,alignment Qt_AlignmentFlag)  {
	q.Drv(255000,255104,Native(item),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::addItem(QGraphicsLayoutItem*,int,int,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsGridLayout) AddItemWithItemRowColumnRowspanColumnspanAlignment(item *QGraphicsLayoutItem,row int,column int,rowSpan int,columnSpan int,alignment Qt_AlignmentFlag)  {
	q.Drv(255000,255105,Native(item),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::alignment(QGraphicsLayoutItem*)
func (q *QGraphicsGridLayout) Alignment(item *QGraphicsLayoutItem) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(255000,255106,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnAlignment(int)
func (q *QGraphicsGridLayout) ColumnAlignment(column int) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(255000,255107,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnCount()
func (q *QGraphicsGridLayout) ColumnCount() int {
	var __rv int
	q.Drv(255000,255108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnMaximumWidth(int)
func (q *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	var __rv float64
	q.Drv(255000,255109,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnMinimumWidth(int)
func (q *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	var __rv float64
	q.Drv(255000,255110,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnPreferredWidth(int)
func (q *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	var __rv float64
	q.Drv(255000,255111,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnSpacing(int)
func (q *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	var __rv float64
	q.Drv(255000,255112,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::columnStretchFactor(int)
func (q *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	var __rv int
	q.Drv(255000,255113,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::count()
func (q *QGraphicsGridLayout) Count() int {
	var __rv int
	q.Drv(255000,255114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::horizontalSpacing()
func (q *QGraphicsGridLayout) HorizontalSpacing() float64 {
	var __rv float64
	q.Drv(255000,255115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::invalidate()
func (q *QGraphicsGridLayout) Invalidate()  {
	q.Drv(255000,255116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::itemAt(int)
func (q *QGraphicsGridLayout) ItemAt(index int) *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(255000,255117,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsGridLayout::itemAt(int,int)
func (q *QGraphicsGridLayout) ItemAtWithRowColumn(row int,column int) *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(255000,255118,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsGridLayout::removeAt(int)
func (q *QGraphicsGridLayout) RemoveAt(index int)  {
	q.Drv(255000,255119,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::rowAlignment(int)
func (q *QGraphicsGridLayout) RowAlignment(row int) Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(255000,255120,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowCount()
func (q *QGraphicsGridLayout) RowCount() int {
	var __rv int
	q.Drv(255000,255121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowMaximumHeight(int)
func (q *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	var __rv float64
	q.Drv(255000,255122,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowMinimumHeight(int)
func (q *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	var __rv float64
	q.Drv(255000,255123,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowPreferredHeight(int)
func (q *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	var __rv float64
	q.Drv(255000,255124,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowSpacing(int)
func (q *QGraphicsGridLayout) RowSpacing(row int) float64 {
	var __rv float64
	q.Drv(255000,255125,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::rowStretchFactor(int)
func (q *QGraphicsGridLayout) RowStretchFactor(row int) int {
	var __rv int
	q.Drv(255000,255126,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsGridLayout::setAlignment(QGraphicsLayoutItem*,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsGridLayout) SetAlignment(item *QGraphicsLayoutItem,alignment Qt_AlignmentFlag)  {
	q.Drv(255000,255127,Native(item),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnAlignment(int,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsGridLayout) SetColumnAlignment(column int,alignment Qt_AlignmentFlag)  {
	q.Drv(255000,255128,unsafe.Pointer(&column),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnFixedWidth(int,double)
func (q *QGraphicsGridLayout) SetColumnFixedWidth(column int,width float64)  {
	q.Drv(255000,255129,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnMaximumWidth(int,double)
func (q *QGraphicsGridLayout) SetColumnMaximumWidth(column int,width float64)  {
	q.Drv(255000,255130,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnMinimumWidth(int,double)
func (q *QGraphicsGridLayout) SetColumnMinimumWidth(column int,width float64)  {
	q.Drv(255000,255131,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnPreferredWidth(int,double)
func (q *QGraphicsGridLayout) SetColumnPreferredWidth(column int,width float64)  {
	q.Drv(255000,255132,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnSpacing(int,double)
func (q *QGraphicsGridLayout) SetColumnSpacing(column int,spacing float64)  {
	q.Drv(255000,255133,unsafe.Pointer(&column),unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setColumnStretchFactor(int,int)
func (q *QGraphicsGridLayout) SetColumnStretchFactor(column int,stretch int)  {
	q.Drv(255000,255134,unsafe.Pointer(&column),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setGeometry(QRectF const&)
func (q *QGraphicsGridLayout) SetGeometry(rect *QRectF)  {
	q.Drv(255000,255135,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setHorizontalSpacing(double)
func (q *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64)  {
	q.Drv(255000,255136,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowAlignment(int,QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsGridLayout) SetRowAlignment(row int,alignment Qt_AlignmentFlag)  {
	q.Drv(255000,255137,unsafe.Pointer(&row),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowFixedHeight(int,double)
func (q *QGraphicsGridLayout) SetRowFixedHeight(row int,height float64)  {
	q.Drv(255000,255138,unsafe.Pointer(&row),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowMaximumHeight(int,double)
func (q *QGraphicsGridLayout) SetRowMaximumHeight(row int,height float64)  {
	q.Drv(255000,255139,unsafe.Pointer(&row),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowMinimumHeight(int,double)
func (q *QGraphicsGridLayout) SetRowMinimumHeight(row int,height float64)  {
	q.Drv(255000,255140,unsafe.Pointer(&row),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowPreferredHeight(int,double)
func (q *QGraphicsGridLayout) SetRowPreferredHeight(row int,height float64)  {
	q.Drv(255000,255141,unsafe.Pointer(&row),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowSpacing(int,double)
func (q *QGraphicsGridLayout) SetRowSpacing(row int,spacing float64)  {
	q.Drv(255000,255142,unsafe.Pointer(&row),unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setRowStretchFactor(int,int)
func (q *QGraphicsGridLayout) SetRowStretchFactor(row int,stretch int)  {
	q.Drv(255000,255143,unsafe.Pointer(&row),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setSpacing(double)
func (q *QGraphicsGridLayout) SetSpacing(spacing float64)  {
	q.Drv(255000,255144,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::setVerticalSpacing(double)
func (q *QGraphicsGridLayout) SetVerticalSpacing(spacing float64)  {
	q.Drv(255000,255145,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsGridLayout::sizeHint(Qt::SizeHint)
func (q *QGraphicsGridLayout) SizeHint(which Qt_SizeHint) *QSizeF {
	var __rv uintptr
	q.Drv(255000,255146,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsGridLayout::sizeHint(Qt::SizeHint,QSizeF const&)
func (q *QGraphicsGridLayout) SizeHintFWithWhichConstraint(which Qt_SizeHint,constraint *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(255000,255147,unsafe.Pointer(&which),Native(constraint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsGridLayout::verticalSpacing()
func (q *QGraphicsGridLayout) VerticalSpacing() float64 {
	var __rv float64
	q.Drv(255000,255148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
