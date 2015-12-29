// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QBoxLayout_Direction - QBoxLayout::Direction
type QBoxLayout_Direction uint32
const (
	QBoxLayout_LeftToRight QBoxLayout_Direction = 0
	QBoxLayout_RightToLeft QBoxLayout_Direction = 1
	QBoxLayout_TopToBottom QBoxLayout_Direction = 2
	QBoxLayout_BottomToTop QBoxLayout_Direction = 3
	QBoxLayout_Down QBoxLayout_Direction = QBoxLayout_TopToBottom
	QBoxLayout_Up QBoxLayout_Direction = QBoxLayout_BottomToTop
)
//struct QBoxLayout : QBoxLayout
type QBoxLayout struct {
	QLayout
}
//QBoxLayout::QBoxLayout(QBoxLayout::Direction,QWidget*)	
func NewBoxLayout(value2 QBoxLayout_Direction,parent QWidgetInterface) *QBoxLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),210000,210102,unsafe.Pointer(&value2),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBoxLayout{}
	_p.SetDriver(__rv,210000,false)
	return _p
} 
//QBoxLayout::addItem(QLayoutItem*)
func (q *QBoxLayout) AddItem(value *QLayoutItem)  {
	q.Drv(210000,210103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addLayout(QLayout*)
func (q *QBoxLayout) AddLayout(layout QLayoutInterface)  {
	q.Drv(210000,210104,Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addLayout(QLayout*,int)
func (q *QBoxLayout) AddLayoutWithLayoutStretch(layout QLayoutInterface,stretch int)  {
	q.Drv(210000,210105,Native(layout),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addSpacerItem(QSpacerItem*)
func (q *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem)  {
	q.Drv(210000,210106,Native(spacerItem),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addSpacing(int)
func (q *QBoxLayout) AddSpacing(size int)  {
	q.Drv(210000,210107,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addStretch()
func (q *QBoxLayout) AddStretch()  {
	q.Drv(210000,210108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addStretch(int)
func (q *QBoxLayout) AddStretchWithStretch(stretch int)  {
	q.Drv(210000,210109,unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addStrut(int)
func (q *QBoxLayout) AddStrut(value int)  {
	q.Drv(210000,210110,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addWidget(QWidget*)
func (q *QBoxLayout) AddWidget(value QWidgetInterface)  {
	q.Drv(210000,210111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::addWidget(QWidget*,int,QFlags<Qt::AlignmentFlag>)
func (q *QBoxLayout) AddWidgetWithWidgetStretchAlignment(value2 QWidgetInterface,stretch int,alignment Qt_AlignmentFlag)  {
	q.Drv(210000,210112,Native(value2),unsafe.Pointer(&stretch),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::count()
func (q *QBoxLayout) Count() int {
	var __rv int
	q.Drv(210000,210113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::direction()
func (q *QBoxLayout) Direction() QBoxLayout_Direction {
	var __rv QBoxLayout_Direction
	q.Drv(210000,210114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::expandingDirections()
func (q *QBoxLayout) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(210000,210115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::hasHeightForWidth()
func (q *QBoxLayout) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(210000,210116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::heightForWidth(int)
func (q *QBoxLayout) HeightForWidth(value int) int {
	var __rv int
	q.Drv(210000,210117,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::insertItem(int,QLayoutItem*)
func (q *QBoxLayout) InsertItem(index int,value2 *QLayoutItem)  {
	q.Drv(210000,210118,unsafe.Pointer(&index),Native(value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertLayout(int,QLayout*,int)
func (q *QBoxLayout) InsertLayout(index int,layout QLayoutInterface,stretch int)  {
	q.Drv(210000,210119,unsafe.Pointer(&index),Native(layout),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertSpacerItem(int,QSpacerItem*)
func (q *QBoxLayout) InsertSpacerItem(index int,spacerItem *QSpacerItem)  {
	q.Drv(210000,210120,unsafe.Pointer(&index),Native(spacerItem),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertSpacing(int,int)
func (q *QBoxLayout) InsertSpacing(index int,size int)  {
	q.Drv(210000,210121,unsafe.Pointer(&index),unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertStretch(int)
func (q *QBoxLayout) InsertStretch(index int)  {
	q.Drv(210000,210122,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertStretch(int,int)
func (q *QBoxLayout) InsertStretchWithIndexStretch(index int,stretch int)  {
	q.Drv(210000,210123,unsafe.Pointer(&index),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::insertWidget(int,QWidget*,int,QFlags<Qt::AlignmentFlag>)
func (q *QBoxLayout) InsertWidget(index int,widget QWidgetInterface,stretch int,alignment Qt_AlignmentFlag)  {
	q.Drv(210000,210124,unsafe.Pointer(&index),Native(widget),unsafe.Pointer(&stretch),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::invalidate()
func (q *QBoxLayout) Invalidate()  {
	q.Drv(210000,210125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::itemAt(int)
func (q *QBoxLayout) ItemAt(value int) *QLayoutItem {
	var __rv uintptr
	q.Drv(210000,210126,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QBoxLayout::maximumSize()
func (q *QBoxLayout) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(210000,210127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QBoxLayout::minimumHeightForWidth(int)
func (q *QBoxLayout) MinimumHeightForWidth(value int) int {
	var __rv int
	q.Drv(210000,210128,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::minimumSize()
func (q *QBoxLayout) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(210000,210129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QBoxLayout::setDirection(QBoxLayout::Direction)
func (q *QBoxLayout) SetDirection(value QBoxLayout_Direction)  {
	q.Drv(210000,210130,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::setGeometry(QRect const&)
func (q *QBoxLayout) SetGeometry(value *QRect)  {
	q.Drv(210000,210131,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::setSpacing(int)
func (q *QBoxLayout) SetSpacing(spacing int)  {
	q.Drv(210000,210132,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::setStretch(int,int)
func (q *QBoxLayout) SetStretch(index int,stretch int)  {
	q.Drv(210000,210133,unsafe.Pointer(&index),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBoxLayout::setStretchFactor(QLayout*,int)
func (q *QBoxLayout) SetStretchFactorWithLayoutStretch(l QLayoutInterface,stretch int) bool {
	var __rv bool
	q.Drv(210000,210134,Native(l),unsafe.Pointer(&stretch),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::setStretchFactor(QWidget*,int)
func (q *QBoxLayout) SetStretchFactorWithWidgetStretch(w QWidgetInterface,stretch int) bool {
	var __rv bool
	q.Drv(210000,210135,Native(w),unsafe.Pointer(&stretch),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::sizeHint()
func (q *QBoxLayout) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(210000,210136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QBoxLayout::spacing()
func (q *QBoxLayout) Spacing() int {
	var __rv int
	q.Drv(210000,210137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::stretch(int)
func (q *QBoxLayout) Stretch(index int) int {
	var __rv int
	q.Drv(210000,210138,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBoxLayout::takeAt(int)
func (q *QBoxLayout) TakeAt(value int) *QLayoutItem {
	var __rv uintptr
	q.Drv(210000,210139,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
