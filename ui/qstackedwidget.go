// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStackedWidget : QStackedWidget
type QStackedWidget struct {
	QFrame
}
func (q *QStackedWidget) OnWidgetRemoved(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(351000,351102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QStackedWidget) OnCurrentChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(351000,351103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QStackedWidget::QStackedWidget()	
func NewStackedWidget() *QStackedWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),351000,351104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStackedWidget{}
	_p.SetDriver(__rv,351000,false)
	return _p
} 
//QStackedWidget::QStackedWidget(QWidget*)	
func NewStackedWidgetWithParent(parent QWidgetInterface) *QStackedWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),351000,351105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStackedWidget{}
	_p.SetDriver(__rv,351000,false)
	return _p
} 
//QStackedWidget::addWidget(QWidget*)
func (q *QStackedWidget) AddWidget(w QWidgetInterface) int {
	var __rv int
	q.Drv(351000,351106,Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::count()
func (q *QStackedWidget) Count() int {
	var __rv int
	q.Drv(351000,351107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::currentIndex()
func (q *QStackedWidget) CurrentIndex() int {
	var __rv int
	q.Drv(351000,351108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::currentWidget()
func (q *QStackedWidget) CurrentWidget() *QWidget {
	var __rv uintptr
	q.Drv(351000,351109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QStackedWidget::event(QEvent*)
func (q *QStackedWidget) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(351000,351110,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::indexOf(QWidget*)
func (q *QStackedWidget) IndexOf(value QWidgetInterface) int {
	var __rv int
	q.Drv(351000,351111,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::insertWidget(int,QWidget*)
func (q *QStackedWidget) InsertWidget(index int,w QWidgetInterface) int {
	var __rv int
	q.Drv(351000,351112,unsafe.Pointer(&index),Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStackedWidget::removeWidget(QWidget*)
func (q *QStackedWidget) RemoveWidget(w QWidgetInterface)  {
	q.Drv(351000,351113,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedWidget::setCurrentIndex(int)
func (q *QStackedWidget) SetCurrentIndex(index int)  {
	q.Drv(351000,351114,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedWidget::setCurrentWidget(QWidget*)
func (q *QStackedWidget) SetCurrentWidget(w QWidgetInterface)  {
	q.Drv(351000,351115,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStackedWidget::widget(int)
func (q *QStackedWidget) Widget(value int) *QWidget {
	var __rv uintptr
	q.Drv(351000,351116,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
