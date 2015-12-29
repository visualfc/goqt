// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLayout_SizeConstraint - QLayout::SizeConstraint
type QLayout_SizeConstraint uint32
const (
	QLayout_SetDefaultConstraint QLayout_SizeConstraint = 0
	QLayout_SetNoConstraint QLayout_SizeConstraint = 1
	QLayout_SetMinimumSize QLayout_SizeConstraint = 2
	QLayout_SetFixedSize QLayout_SizeConstraint = 3
	QLayout_SetMaximumSize QLayout_SizeConstraint = 4
	QLayout_SetMinAndMaxSize QLayout_SizeConstraint = 5
)
//struct QLayout : QLayout
type QLayout struct {
	QObject
}
//QLayout::activate()
func (q *QLayout) Activate() bool {
	var __rv bool
	q.Drv(300000,300102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::addChildLayout(QLayout*)
func (q *QLayout) AddChildLayout(l QLayoutInterface)  {
	q.Drv(300000,300103,Native(l),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::addChildWidget(QWidget*)
func (q *QLayout) AddChildWidget(w QWidgetInterface)  {
	q.Drv(300000,300104,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::addItem(QLayoutItem*)
func (q *QLayout) AddItem(value *QLayoutItem)  {
	q.Drv(300000,300105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::addWidget(QWidget*)
func (q *QLayout) AddWidget(w QWidgetInterface)  {
	q.Drv(300000,300106,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::alignmentRect(QRect const&)
func (q *QLayout) AlignmentRect(value *QRect) *QRect {
	var __rv uintptr
	q.Drv(300000,300107,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QLayout::childEvent(QChildEvent*)
func (q *QLayout) ChildEvent(e *QChildEvent)  {
	q.Drv(300000,300108,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::closestAcceptableSize(QWidget const*,QSize const&)	
func QLayoutClosestAcceptableSize(w QWidgetInterface,s *QSize) *QSize {
	var __rv uintptr
	DirectQtDrv(nil,300000,300109,Native(w),Native(s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::closestAcceptableSize(QWidget const*,QSize const&)
func (q *QLayout) ClosestAcceptableSize(w QWidgetInterface,s *QSize) *QSize {
	var __rv uintptr
	q.Drv(300000,300109,Native(w),Native(s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::contentsMargins()
func (q *QLayout) ContentsMargins() *QMargins {
	var __rv uintptr
	q.Drv(300000,300110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMargins{}
	_rp.SetDriver(__rv,73000,true)
	return _rp
}	
//QLayout::contentsRect()
func (q *QLayout) ContentsRect() *QRect {
	var __rv uintptr
	q.Drv(300000,300111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QLayout::count()
func (q *QLayout) Count() int {
	var __rv int
	q.Drv(300000,300112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::expandingDirections()
func (q *QLayout) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(300000,300113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::geometry()
func (q *QLayout) Geometry() *QRect {
	var __rv uintptr
	q.Drv(300000,300114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QLayout::getContentsMargins(int*,int*,int*,int*)
func (q *QLayout) GetContentsMargins(left *int,top *int,right *int,bottom *int)  {
	q.Drv(300000,300115,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::indexOf(QWidget*)
func (q *QLayout) IndexOf(value QWidgetInterface) int {
	var __rv int
	q.Drv(300000,300116,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::invalidate()
func (q *QLayout) Invalidate()  {
	q.Drv(300000,300117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::isEmpty()
func (q *QLayout) IsEmpty() bool {
	var __rv bool
	q.Drv(300000,300118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::isEnabled()
func (q *QLayout) IsEnabled() bool {
	var __rv bool
	q.Drv(300000,300119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::itemAt(int)
func (q *QLayout) ItemAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(300000,300120,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QLayout::layout()
func (q *QLayout) Layout() *QLayout {
	var __rv uintptr
	q.Drv(300000,300121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayout{}
	_rp.SetDriver(__rv,300000,false)
	return _rp
}	
//QLayout::margin()
func (q *QLayout) Margin() int {
	var __rv int
	q.Drv(300000,300122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::maximumSize()
func (q *QLayout) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(300000,300123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::menuBar()
func (q *QLayout) MenuBar() *QWidget {
	var __rv uintptr
	q.Drv(300000,300124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QLayout::minimumSize()
func (q *QLayout) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(300000,300125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::parentWidget()
func (q *QLayout) ParentWidget() *QWidget {
	var __rv uintptr
	q.Drv(300000,300126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QLayout::removeItem(QLayoutItem*)
func (q *QLayout) RemoveItem(value *QLayoutItem)  {
	q.Drv(300000,300127,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::removeWidget(QWidget*)
func (q *QLayout) RemoveWidget(w QWidgetInterface)  {
	q.Drv(300000,300128,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QLayout) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(300000,300129,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setAlignment(QLayout*,QFlags<Qt::AlignmentFlag>)
func (q *QLayout) SetAlignmentWithLayoutAlignment(l QLayoutInterface,alignment Qt_AlignmentFlag) bool {
	var __rv bool
	q.Drv(300000,300130,Native(l),unsafe.Pointer(&alignment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::setAlignment(QWidget*,QFlags<Qt::AlignmentFlag>)
func (q *QLayout) SetAlignmentWithWidgetAlignment(w QWidgetInterface,alignment Qt_AlignmentFlag) bool {
	var __rv bool
	q.Drv(300000,300131,Native(w),unsafe.Pointer(&alignment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::setContentsMargins(QMargins const&)
func (q *QLayout) SetContentsMargins(margins *QMargins)  {
	q.Drv(300000,300132,Native(margins),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setContentsMargins(int,int,int,int)
func (q *QLayout) SetContentsMarginsWithLeftTopRightBottom(left int,top int,right int,bottom int)  {
	q.Drv(300000,300133,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setEnabled(bool)
func (q *QLayout) SetEnabled(value bool)  {
	q.Drv(300000,300134,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setGeometry(QRect const&)
func (q *QLayout) SetGeometry(value *QRect)  {
	q.Drv(300000,300135,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setMargin(int)
func (q *QLayout) SetMargin(value int)  {
	q.Drv(300000,300136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setMenuBar(QWidget*)
func (q *QLayout) SetMenuBar(w QWidgetInterface)  {
	q.Drv(300000,300137,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setSizeConstraint(QLayout::SizeConstraint)
func (q *QLayout) SetSizeConstraint(value QLayout_SizeConstraint)  {
	q.Drv(300000,300138,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::setSpacing(int)
func (q *QLayout) SetSpacing(value int)  {
	q.Drv(300000,300139,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::sizeConstraint()
func (q *QLayout) SizeConstraint() QLayout_SizeConstraint {
	var __rv QLayout_SizeConstraint
	q.Drv(300000,300140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::spacing()
func (q *QLayout) Spacing() int {
	var __rv int
	q.Drv(300000,300141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::takeAt(int)
func (q *QLayout) TakeAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(300000,300142,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QLayout::totalHeightForWidth(int)
func (q *QLayout) TotalHeightForWidth(w int) int {
	var __rv int
	q.Drv(300000,300143,unsafe.Pointer(&w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLayout::totalMaximumSize()
func (q *QLayout) TotalMaximumSize() *QSize {
	var __rv uintptr
	q.Drv(300000,300144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::totalMinimumSize()
func (q *QLayout) TotalMinimumSize() *QSize {
	var __rv uintptr
	q.Drv(300000,300145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::totalSizeHint()
func (q *QLayout) TotalSizeHint() *QSize {
	var __rv uintptr
	q.Drv(300000,300146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLayout::update()
func (q *QLayout) Update()  {
	q.Drv(300000,300147,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLayout::widgetEvent(QEvent*)
func (q *QLayout) WidgetEvent(value *QEvent)  {
	q.Drv(300000,300148,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
