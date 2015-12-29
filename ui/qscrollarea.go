// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QScrollArea : QScrollArea
type QScrollArea struct {
	QAbstractScrollArea
}
//QScrollArea::QScrollArea()	
func NewScrollArea() *QScrollArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),335000,335102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QScrollArea{}
	_p.SetDriver(__rv,335000,false)
	return _p
} 
//QScrollArea::QScrollArea(QWidget*)	
func NewScrollAreaWithParent(parent QWidgetInterface) *QScrollArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),335000,335103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QScrollArea{}
	_p.SetDriver(__rv,335000,false)
	return _p
} 
//QScrollArea::alignment()
func (q *QScrollArea) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(335000,335104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QScrollArea::ensureVisible(int,int,int,int)
func (q *QScrollArea) EnsureVisible(x int,y int,xmargin int,ymargin int)  {
	q.Drv(335000,335105,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::ensureWidgetVisible(QWidget*)
func (q *QScrollArea) EnsureWidgetVisible(childWidget QWidgetInterface)  {
	q.Drv(335000,335106,Native(childWidget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::ensureWidgetVisible(QWidget*,int,int)
func (q *QScrollArea) EnsureWidgetVisibleWithWidgetXmarginYmargin(childWidget QWidgetInterface,xmargin int,ymargin int)  {
	q.Drv(335000,335107,Native(childWidget),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::event(QEvent*)
func (q *QScrollArea) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(335000,335108,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QScrollArea::eventFilter(QObject*,QEvent*)
func (q *QScrollArea) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(335000,335109,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QScrollArea::focusNextPrevChild(bool)
func (q *QScrollArea) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(335000,335110,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QScrollArea::resizeEvent(QResizeEvent*)
func (q *QScrollArea) ResizeEvent(value *QResizeEvent)  {
	q.Drv(335000,335111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::scrollContentsBy(int,int)
func (q *QScrollArea) ScrollContentsBy(dx int,dy int)  {
	q.Drv(335000,335112,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QScrollArea) SetAlignment(value Qt_AlignmentFlag)  {
	q.Drv(335000,335113,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::setWidget(QWidget*)
func (q *QScrollArea) SetWidget(widget QWidgetInterface)  {
	q.Drv(335000,335114,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::setWidgetResizable(bool)
func (q *QScrollArea) SetWidgetResizable(resizable bool)  {
	q.Drv(335000,335115,unsafe.Pointer(&resizable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollArea::sizeHint()
func (q *QScrollArea) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(335000,335116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QScrollArea::takeWidget()
func (q *QScrollArea) TakeWidget() *QWidget {
	var __rv uintptr
	q.Drv(335000,335117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QScrollArea::widget()
func (q *QScrollArea) Widget() *QWidget {
	var __rv uintptr
	q.Drv(335000,335118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QScrollArea::widgetResizable()
func (q *QScrollArea) WidgetResizable() bool {
	var __rv bool
	q.Drv(335000,335119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
