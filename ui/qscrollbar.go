// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QScrollBar : QScrollBar
type QScrollBar struct {
	QAbstractSlider
}
//QScrollBar::QScrollBar()	
func NewScrollBar() *QScrollBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),336000,336102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QScrollBar{}
	_p.SetDriver(__rv,336000,false)
	return _p
} 
//QScrollBar::QScrollBar(QWidget*)	
func NewScrollBarWithParent(parent QWidgetInterface) *QScrollBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),336000,336103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QScrollBar{}
	_p.SetDriver(__rv,336000,false)
	return _p
} 
//QScrollBar::QScrollBar(Qt::Orientation,QWidget*)	
func NewScrollBarWithOrientationParent(value2 Qt_Orientation,parent QWidgetInterface) *QScrollBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),336000,336104,unsafe.Pointer(&value2),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QScrollBar{}
	_p.SetDriver(__rv,336000,false)
	return _p
} 
//QScrollBar::contextMenuEvent(QContextMenuEvent*)
func (q *QScrollBar) ContextMenuEvent(value *QContextMenuEvent)  {
	q.Drv(336000,336105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::event(QEvent*)
func (q *QScrollBar) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(336000,336106,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QScrollBar::hideEvent(QHideEvent*)
func (q *QScrollBar) HideEvent(value *QHideEvent)  {
	q.Drv(336000,336107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::mouseMoveEvent(QMouseEvent*)
func (q *QScrollBar) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(336000,336108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::mousePressEvent(QMouseEvent*)
func (q *QScrollBar) MousePressEvent(value *QMouseEvent)  {
	q.Drv(336000,336109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::mouseReleaseEvent(QMouseEvent*)
func (q *QScrollBar) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(336000,336110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::paintEvent(QPaintEvent*)
func (q *QScrollBar) PaintEvent(value *QPaintEvent)  {
	q.Drv(336000,336111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QScrollBar::sizeHint()
func (q *QScrollBar) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(336000,336112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QScrollBar::sliderChange(QAbstractSlider::SliderChange)
func (q *QScrollBar) SliderChange(change QAbstractSlider_SliderChange)  {
	q.Drv(336000,336113,unsafe.Pointer(&change),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
