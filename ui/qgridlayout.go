// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGridLayout : QGridLayout
type QGridLayout struct {
	QLayout
}
//QGridLayout::QGridLayout()	
func NewGridLayout() *QGridLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),287000,287102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGridLayout{}
	_p.SetDriver(__rv,287000,false)
	return _p
} 
//QGridLayout::QGridLayout(QWidget*)	
func NewGridLayoutWithParent(parent QWidgetInterface) *QGridLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),287000,287103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGridLayout{}
	_p.SetDriver(__rv,287000,false)
	return _p
} 
//QGridLayout::addItem(QLayoutItem*)
func (q *QGridLayout) AddItem(value *QLayoutItem)  {
	q.Drv(287000,287104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addItem(QLayoutItem*,int,int,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGridLayout) AddItemWithItemRowColumnRowspanColumnspanAlignmentflag(item *QLayoutItem,row int,column int,rowSpan int,columnSpan int,value2 Qt_AlignmentFlag)  {
	q.Drv(287000,287105,Native(item),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addLayout(QLayout*,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGridLayout) AddLayoutWithLayoutRowColumnAlignmentflag(value2 QLayoutInterface,row int,column int,value3 Qt_AlignmentFlag)  {
	q.Drv(287000,287106,Native(value2),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addLayout(QLayout*,int,int,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGridLayout) AddLayoutWithLayoutRowColumnRowspanColumnspanAlignmentflag(value2 QLayoutInterface,row int,column int,rowSpan int,columnSpan int,value3 Qt_AlignmentFlag)  {
	q.Drv(287000,287107,Native(value2),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addWidget(QWidget*)
func (q *QGridLayout) AddWidget(w QWidgetInterface)  {
	q.Drv(287000,287108,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addWidget(QWidget*,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGridLayout) AddWidgetWithWidgetRowColumnAlignmentflag(value2 QWidgetInterface,row int,column int,value3 Qt_AlignmentFlag)  {
	q.Drv(287000,287109,Native(value2),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::addWidget(QWidget*,int,int,int,int,QFlags<Qt::AlignmentFlag>)
func (q *QGridLayout) AddWidgetWithWidgetRowColumnRowspanColumnspanAlignmentflag(value2 QWidgetInterface,row int,column int,rowSpan int,columnSpan int,value3 Qt_AlignmentFlag)  {
	q.Drv(287000,287110,Native(value2),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::cellRect(int,int)
func (q *QGridLayout) CellRect(row int,column int) *QRect {
	var __rv uintptr
	q.Drv(287000,287111,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QGridLayout::columnCount()
func (q *QGridLayout) ColumnCount() int {
	var __rv int
	q.Drv(287000,287112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::columnMinimumWidth(int)
func (q *QGridLayout) ColumnMinimumWidth(column int) int {
	var __rv int
	q.Drv(287000,287113,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::columnStretch(int)
func (q *QGridLayout) ColumnStretch(column int) int {
	var __rv int
	q.Drv(287000,287114,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::count()
func (q *QGridLayout) Count() int {
	var __rv int
	q.Drv(287000,287115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::expandingDirections()
func (q *QGridLayout) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(287000,287116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::getItemPosition(int,int*,int*,int*,int*)
func (q *QGridLayout) GetItemPosition(idx int,row *int,column *int,rowSpan *int,columnSpan *int)  {
	q.Drv(287000,287117,unsafe.Pointer(&idx),unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::hasHeightForWidth()
func (q *QGridLayout) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(287000,287118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::heightForWidth(int)
func (q *QGridLayout) HeightForWidth(value int) int {
	var __rv int
	q.Drv(287000,287119,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::horizontalSpacing()
func (q *QGridLayout) HorizontalSpacing() int {
	var __rv int
	q.Drv(287000,287120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::invalidate()
func (q *QGridLayout) Invalidate()  {
	q.Drv(287000,287121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::itemAt(int)
func (q *QGridLayout) ItemAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(287000,287122,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QGridLayout::itemAtPosition(int,int)
func (q *QGridLayout) ItemAtPosition(row int,column int) *QLayoutItem {
	var __rv uintptr
	q.Drv(287000,287123,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QGridLayout::maximumSize()
func (q *QGridLayout) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(287000,287124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QGridLayout::minimumHeightForWidth(int)
func (q *QGridLayout) MinimumHeightForWidth(value int) int {
	var __rv int
	q.Drv(287000,287125,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::minimumSize()
func (q *QGridLayout) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(287000,287126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QGridLayout::originCorner()
func (q *QGridLayout) OriginCorner() Qt_Corner {
	var __rv Qt_Corner
	q.Drv(287000,287127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::rowCount()
func (q *QGridLayout) RowCount() int {
	var __rv int
	q.Drv(287000,287128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::rowMinimumHeight(int)
func (q *QGridLayout) RowMinimumHeight(row int) int {
	var __rv int
	q.Drv(287000,287129,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::rowStretch(int)
func (q *QGridLayout) RowStretch(row int) int {
	var __rv int
	q.Drv(287000,287130,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::setColumnMinimumWidth(int,int)
func (q *QGridLayout) SetColumnMinimumWidth(column int,minSize int)  {
	q.Drv(287000,287131,unsafe.Pointer(&column),unsafe.Pointer(&minSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setColumnStretch(int,int)
func (q *QGridLayout) SetColumnStretch(column int,stretch int)  {
	q.Drv(287000,287132,unsafe.Pointer(&column),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setDefaultPositioning(int,Qt::Orientation)
func (q *QGridLayout) SetDefaultPositioning(n int,orient Qt_Orientation)  {
	q.Drv(287000,287133,unsafe.Pointer(&n),unsafe.Pointer(&orient),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setGeometry(QRect const&)
func (q *QGridLayout) SetGeometry(value *QRect)  {
	q.Drv(287000,287134,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setHorizontalSpacing(int)
func (q *QGridLayout) SetHorizontalSpacing(spacing int)  {
	q.Drv(287000,287135,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setOriginCorner(Qt::Corner)
func (q *QGridLayout) SetOriginCorner(value Qt_Corner)  {
	q.Drv(287000,287136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setRowMinimumHeight(int,int)
func (q *QGridLayout) SetRowMinimumHeight(row int,minSize int)  {
	q.Drv(287000,287137,unsafe.Pointer(&row),unsafe.Pointer(&minSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setRowStretch(int,int)
func (q *QGridLayout) SetRowStretch(row int,stretch int)  {
	q.Drv(287000,287138,unsafe.Pointer(&row),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setSpacing(int)
func (q *QGridLayout) SetSpacing(spacing int)  {
	q.Drv(287000,287139,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::setVerticalSpacing(int)
func (q *QGridLayout) SetVerticalSpacing(spacing int)  {
	q.Drv(287000,287140,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGridLayout::sizeHint()
func (q *QGridLayout) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(287000,287141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QGridLayout::spacing()
func (q *QGridLayout) Spacing() int {
	var __rv int
	q.Drv(287000,287142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGridLayout::takeAt(int)
func (q *QGridLayout) TakeAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(287000,287143,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QGridLayout::verticalSpacing()
func (q *QGridLayout) VerticalSpacing() int {
	var __rv int
	q.Drv(287000,287144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
