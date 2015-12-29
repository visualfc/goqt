// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsProxyWidget_enum_1 - QGraphicsProxyWidget::enum_1
type QGraphicsProxyWidget_enum_1 uint32
const (
	QGraphicsProxyWidget_Type QGraphicsProxyWidget_enum_1 = 12
)
//struct QGraphicsProxyWidget : QGraphicsProxyWidget
type QGraphicsProxyWidget struct {
	QGraphicsWidget
}
//QGraphicsProxyWidget::QGraphicsProxyWidget()	
func NewGraphicsProxyWidget() *QGraphicsProxyWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),268000,268102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsProxyWidget{}
	_p.SetDriver(__rv,268000,false)
	return _p
} 
//QGraphicsProxyWidget::QGraphicsProxyWidget(QGraphicsItem*,QFlags<Qt::WindowType>)	
func NewGraphicsProxyWidgetWithParentFlags(parent *QGraphicsItem,wFlags Qt_WindowType) *QGraphicsProxyWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),268000,268103,Native(parent),unsafe.Pointer(&wFlags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsProxyWidget{}
	_p.SetDriver(__rv,268000,false)
	return _p
} 
//QGraphicsProxyWidget::contextMenuEvent(QGraphicsSceneContextMenuEvent*)
func (q *QGraphicsProxyWidget) ContextMenuEvent(event *QGraphicsSceneContextMenuEvent)  {
	q.Drv(268000,268104,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::createProxyForChildWidget(QWidget*)
func (q *QGraphicsProxyWidget) CreateProxyForChildWidget(child QWidgetInterface) *QGraphicsProxyWidget {
	var __rv uintptr
	q.Drv(268000,268105,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsProxyWidget{}
	_rp.SetDriver(__rv,268000,false)
	return _rp
}	
//QGraphicsProxyWidget::dragEnterEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsProxyWidget) DragEnterEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(268000,268106,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::dragLeaveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsProxyWidget) DragLeaveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(268000,268107,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::dragMoveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsProxyWidget) DragMoveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(268000,268108,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::dropEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsProxyWidget) DropEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(268000,268109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::event(QEvent*)
func (q *QGraphicsProxyWidget) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(268000,268110,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsProxyWidget::eventFilter(QObject*,QEvent*)
func (q *QGraphicsProxyWidget) EventFilter(object QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(268000,268111,Native(object),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsProxyWidget::focusInEvent(QFocusEvent*)
func (q *QGraphicsProxyWidget) FocusInEvent(event *QFocusEvent)  {
	q.Drv(268000,268112,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::focusNextPrevChild(bool)
func (q *QGraphicsProxyWidget) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(268000,268113,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsProxyWidget::focusOutEvent(QFocusEvent*)
func (q *QGraphicsProxyWidget) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(268000,268114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::grabMouseEvent(QEvent*)
func (q *QGraphicsProxyWidget) GrabMouseEvent(event *QEvent)  {
	q.Drv(268000,268115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::hideEvent(QHideEvent*)
func (q *QGraphicsProxyWidget) HideEvent(event *QHideEvent)  {
	q.Drv(268000,268116,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::hoverEnterEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsProxyWidget) HoverEnterEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(268000,268117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::hoverLeaveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsProxyWidget) HoverLeaveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(268000,268118,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::hoverMoveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsProxyWidget) HoverMoveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(268000,268119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::itemChange(QGraphicsItem::GraphicsItemChange,QVariant const&)
func (q *QGraphicsProxyWidget) ItemChange(change QGraphicsItem_GraphicsItemChange,value *QVariant) *QVariant {
	var __rv uintptr
	q.Drv(268000,268120,unsafe.Pointer(&change),Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsProxyWidget::keyPressEvent(QKeyEvent*)
func (q *QGraphicsProxyWidget) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(268000,268121,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::keyReleaseEvent(QKeyEvent*)
func (q *QGraphicsProxyWidget) KeyReleaseEvent(event *QKeyEvent)  {
	q.Drv(268000,268122,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::mouseDoubleClickEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsProxyWidget) MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(268000,268123,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::mouseMoveEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsProxyWidget) MouseMoveEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(268000,268124,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::mousePressEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsProxyWidget) MousePressEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(268000,268125,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::mouseReleaseEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsProxyWidget) MouseReleaseEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(268000,268126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::newProxyWidget(QWidget const*)
func (q *QGraphicsProxyWidget) NewProxyWidget(value QWidgetInterface) *QGraphicsProxyWidget {
	var __rv uintptr
	q.Drv(268000,268127,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsProxyWidget{}
	_rp.SetDriver(__rv,268000,false)
	return _rp
}	
//QGraphicsProxyWidget::resizeEvent(QGraphicsSceneResizeEvent*)
func (q *QGraphicsProxyWidget) ResizeEvent(event *QGraphicsSceneResizeEvent)  {
	q.Drv(268000,268128,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::setGeometry(QRectF const&)
func (q *QGraphicsProxyWidget) SetGeometry(rect *QRectF)  {
	q.Drv(268000,268129,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::setWidget(QWidget*)
func (q *QGraphicsProxyWidget) SetWidget(widget QWidgetInterface)  {
	q.Drv(268000,268130,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::showEvent(QShowEvent*)
func (q *QGraphicsProxyWidget) ShowEvent(event *QShowEvent)  {
	q.Drv(268000,268131,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::sizeHint(Qt::SizeHint,QSizeF const&)
func (q *QGraphicsProxyWidget) SizeHint(which Qt_SizeHint,constraint *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(268000,268132,unsafe.Pointer(&which),Native(constraint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsProxyWidget::subWidgetRect(QWidget const*)
func (q *QGraphicsProxyWidget) SubWidgetRect(widget QWidgetInterface) *QRectF {
	var __rv uintptr
	q.Drv(268000,268133,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsProxyWidget::type()
func (q *QGraphicsProxyWidget) Type() int {
	var __rv int
	q.Drv(268000,268134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsProxyWidget::ungrabMouseEvent(QEvent*)
func (q *QGraphicsProxyWidget) UngrabMouseEvent(event *QEvent)  {
	q.Drv(268000,268135,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::wheelEvent(QGraphicsSceneWheelEvent*)
func (q *QGraphicsProxyWidget) WheelEvent(event *QGraphicsSceneWheelEvent)  {
	q.Drv(268000,268136,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsProxyWidget::widget()
func (q *QGraphicsProxyWidget) Widget() *QWidget {
	var __rv uintptr
	q.Drv(268000,268137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
