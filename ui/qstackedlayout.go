// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStackedLayout_StackingMode - QStackedLayout::StackingMode
type QStackedLayout_StackingMode uint32
const (
	QStackedLayout_StackOne QStackedLayout_StackingMode = 0
	QStackedLayout_StackAll QStackedLayout_StackingMode = 1
)
//struct QStackedLayout : QStackedLayout
type QStackedLayout struct {
	QLayout
}
func (q *QStackedLayout) OnWidgetRemoved(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(350000,350102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QStackedLayout) OnCurrentChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(350000,350103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QStackedLayout::QStackedLayout()	
func NewStackedLayout() *QStackedLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),350000,350104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStackedLayout{}
	_p.SetDriver(__rv,350000,false)
	return _p
} 
//QStackedLayout::QStackedLayout(QLayout*)	
func NewStackedLayoutWithLayout(parentLayout QLayoutInterface) *QStackedLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),350000,350105,Native(parentLayout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStackedLayout{}
	_p.SetDriver(__rv,350000,false)
	return _p
} 
//QStackedLayout::QStackedLayout(QWidget*)	
func NewStackedLayoutWithParent(parent QWidgetInterface) *QStackedLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),350000,350106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStackedLayout{}
	_p.SetDriver(__rv,350000,false)
	return _p
} 
//QStackedLayout::addItem(QLayoutItem*)
func (q *QStackedLayout) AddItem(item *QLayoutItem)  {
	q.Drv(350000,350107,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedLayout::addWidget(QWidget*)
func (q *QStackedLayout) AddWidget(w QWidgetInterface) int {
	var __rv int
	q.Drv(350000,350108,Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedLayout::count()
func (q *QStackedLayout) Count() int {
	var __rv int
	q.Drv(350000,350109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedLayout::currentIndex()
func (q *QStackedLayout) CurrentIndex() int {
	var __rv int
	q.Drv(350000,350110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedLayout::currentWidget()
func (q *QStackedLayout) CurrentWidget() *QWidget {
	var __rv uintptr
	q.Drv(350000,350111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QStackedLayout::insertWidget(int,QWidget*)
func (q *QStackedLayout) InsertWidget(index int,w QWidgetInterface) int {
	var __rv int
	q.Drv(350000,350112,unsafe.Pointer(&index),Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedLayout::itemAt(int)
func (q *QStackedLayout) ItemAt(value int) *QLayoutItem {
	var __rv uintptr
	q.Drv(350000,350113,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QStackedLayout::minimumSize()
func (q *QStackedLayout) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(350000,350114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QStackedLayout::setCurrentIndex(int)
func (q *QStackedLayout) SetCurrentIndex(index int)  {
	q.Drv(350000,350115,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedLayout::setCurrentWidget(QWidget*)
func (q *QStackedLayout) SetCurrentWidget(w QWidgetInterface)  {
	q.Drv(350000,350116,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedLayout::setGeometry(QRect const&)
func (q *QStackedLayout) SetGeometry(rect *QRect)  {
	q.Drv(350000,350117,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedLayout::setStackingMode(QStackedLayout::StackingMode)
func (q *QStackedLayout) SetStackingMode(stackingMode QStackedLayout_StackingMode)  {
	q.Drv(350000,350118,unsafe.Pointer(&stackingMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedLayout::sizeHint()
func (q *QStackedLayout) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(350000,350119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QStackedLayout::stackingMode()
func (q *QStackedLayout) StackingMode() QStackedLayout_StackingMode {
	var __rv QStackedLayout_StackingMode
	q.Drv(350000,350120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedLayout::takeAt(int)
func (q *QStackedLayout) TakeAt(value int) *QLayoutItem {
	var __rv uintptr
	q.Drv(350000,350121,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QStackedLayout::widget()
func (q *QStackedLayout) Widget() *QWidget {
	var __rv uintptr
	q.Drv(350000,350122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QStackedLayout::widget(int)
func (q *QStackedLayout) WidgetWithInt(value int) *QWidget {
	var __rv uintptr
	q.Drv(350000,350123,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
