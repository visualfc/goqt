// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractScrollArea : QAbstractScrollArea
type QAbstractScrollArea struct {
	QFrame
}
//QAbstractScrollArea::QAbstractScrollArea()	
func NewAbstractScrollArea() *QAbstractScrollArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),200000,200102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractScrollArea{}
	_p.SetDriver(__rv,200000,false)
	return _p
} 
//QAbstractScrollArea::QAbstractScrollArea(QWidget*)	
func NewAbstractScrollAreaWithParent(parent QWidgetInterface) *QAbstractScrollArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),200000,200103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractScrollArea{}
	_p.SetDriver(__rv,200000,false)
	return _p
} 
//QAbstractScrollArea::addScrollBarWidget(QWidget*,QFlags<Qt::AlignmentFlag>)
func (q *QAbstractScrollArea) AddScrollBarWidget(widget QWidgetInterface,alignment Qt_AlignmentFlag)  {
	q.Drv(200000,200104,Native(widget),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::contextMenuEvent(QContextMenuEvent*)
func (q *QAbstractScrollArea) ContextMenuEvent(value *QContextMenuEvent)  {
	q.Drv(200000,200105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::cornerWidget()
func (q *QAbstractScrollArea) CornerWidget() *QWidget {
	var __rv uintptr
	q.Drv(200000,200106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QAbstractScrollArea::dragEnterEvent(QDragEnterEvent*)
func (q *QAbstractScrollArea) DragEnterEvent(value *QDragEnterEvent)  {
	q.Drv(200000,200107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::dragLeaveEvent(QDragLeaveEvent*)
func (q *QAbstractScrollArea) DragLeaveEvent(value *QDragLeaveEvent)  {
	q.Drv(200000,200108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::dragMoveEvent(QDragMoveEvent*)
func (q *QAbstractScrollArea) DragMoveEvent(value *QDragMoveEvent)  {
	q.Drv(200000,200109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::dropEvent(QDropEvent*)
func (q *QAbstractScrollArea) DropEvent(value *QDropEvent)  {
	q.Drv(200000,200110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::event(QEvent*)
func (q *QAbstractScrollArea) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(200000,200111,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractScrollArea::horizontalScrollBar()
func (q *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	var __rv uintptr
	q.Drv(200000,200112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QScrollBar{}
	_rp.SetDriver(__rv,336000,false)
	return _rp
}	
//QAbstractScrollArea::horizontalScrollBarPolicy()
func (q *QAbstractScrollArea) HorizontalScrollBarPolicy() Qt_ScrollBarPolicy {
	var __rv Qt_ScrollBarPolicy
	q.Drv(200000,200113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractScrollArea::keyPressEvent(QKeyEvent*)
func (q *QAbstractScrollArea) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(200000,200114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::maximumViewportSize()
func (q *QAbstractScrollArea) MaximumViewportSize() *QSize {
	var __rv uintptr
	q.Drv(200000,200115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractScrollArea::minimumSizeHint()
func (q *QAbstractScrollArea) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(200000,200116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractScrollArea::mouseDoubleClickEvent(QMouseEvent*)
func (q *QAbstractScrollArea) MouseDoubleClickEvent(value *QMouseEvent)  {
	q.Drv(200000,200117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::mouseMoveEvent(QMouseEvent*)
func (q *QAbstractScrollArea) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(200000,200118,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::mousePressEvent(QMouseEvent*)
func (q *QAbstractScrollArea) MousePressEvent(value *QMouseEvent)  {
	q.Drv(200000,200119,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::mouseReleaseEvent(QMouseEvent*)
func (q *QAbstractScrollArea) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(200000,200120,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::paintEvent(QPaintEvent*)
func (q *QAbstractScrollArea) PaintEvent(value *QPaintEvent)  {
	q.Drv(200000,200121,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::resizeEvent(QResizeEvent*)
func (q *QAbstractScrollArea) ResizeEvent(value *QResizeEvent)  {
	q.Drv(200000,200122,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::scrollBarWidgets(QFlags<Qt::AlignmentFlag>)
func (q *QAbstractScrollArea) ScrollBarWidgets(alignment Qt_AlignmentFlag) []*QWidget {
	var __rv []*QWidget
	q.Drv(200000,200123,unsafe.Pointer(&alignment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractScrollArea::scrollContentsBy(int,int)
func (q *QAbstractScrollArea) ScrollContentsBy(dx int,dy int)  {
	q.Drv(200000,200124,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setCornerWidget(QWidget*)
func (q *QAbstractScrollArea) SetCornerWidget(widget QWidgetInterface)  {
	q.Drv(200000,200125,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setHorizontalScrollBar(QScrollBar*)
func (q *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar)  {
	q.Drv(200000,200126,Native(scrollbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setHorizontalScrollBarPolicy(Qt::ScrollBarPolicy)
func (q *QAbstractScrollArea) SetHorizontalScrollBarPolicy(value Qt_ScrollBarPolicy)  {
	q.Drv(200000,200127,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setVerticalScrollBar(QScrollBar*)
func (q *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar)  {
	q.Drv(200000,200128,Native(scrollbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setVerticalScrollBarPolicy(Qt::ScrollBarPolicy)
func (q *QAbstractScrollArea) SetVerticalScrollBarPolicy(value Qt_ScrollBarPolicy)  {
	q.Drv(200000,200129,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setViewport(QWidget*)
func (q *QAbstractScrollArea) SetViewport(widget QWidgetInterface)  {
	q.Drv(200000,200130,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setViewportMargins(QMargins const&)
func (q *QAbstractScrollArea) SetViewportMargins(margins *QMargins)  {
	q.Drv(200000,200131,Native(margins),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setViewportMargins(int,int,int,int)
func (q *QAbstractScrollArea) SetViewportMarginsWithLeftTopRightBottom(left int,top int,right int,bottom int)  {
	q.Drv(200000,200132,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::setupViewport(QWidget*)
func (q *QAbstractScrollArea) SetupViewport(viewport QWidgetInterface)  {
	q.Drv(200000,200133,Native(viewport),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractScrollArea::sizeHint()
func (q *QAbstractScrollArea) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(200000,200134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractScrollArea::verticalScrollBar()
func (q *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	var __rv uintptr
	q.Drv(200000,200135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QScrollBar{}
	_rp.SetDriver(__rv,336000,false)
	return _rp
}	
//QAbstractScrollArea::verticalScrollBarPolicy()
func (q *QAbstractScrollArea) VerticalScrollBarPolicy() Qt_ScrollBarPolicy {
	var __rv Qt_ScrollBarPolicy
	q.Drv(200000,200136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractScrollArea::viewport()
func (q *QAbstractScrollArea) Viewport() *QWidget {
	var __rv uintptr
	q.Drv(200000,200137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QAbstractScrollArea::viewportEvent(QEvent*)
func (q *QAbstractScrollArea) ViewportEvent(value *QEvent) bool {
	var __rv bool
	q.Drv(200000,200138,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractScrollArea::wheelEvent(QWheelEvent*)
func (q *QAbstractScrollArea) WheelEvent(value *QWheelEvent)  {
	q.Drv(200000,200139,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
