// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsTextItem_enum_1 - QGraphicsTextItem::enum_1
type QGraphicsTextItem_enum_1 uint32
const (
	QGraphicsTextItem_Type QGraphicsTextItem_enum_1 = 8
)
//struct QGraphicsTextItem : QGraphicsTextItem
type QGraphicsTextItem struct {
	QGraphicsObject
}
func (q *QGraphicsTextItem) OnLinkActivated(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(283000,283102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsTextItem) OnLinkHovered(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(283000,283103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsTextItem::QGraphicsTextItem()	
func NewGraphicsTextItem() *QGraphicsTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),283000,283104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsTextItem{}
	_p.SetDriver(__rv,283000,false)
	return _p
} 
//QGraphicsTextItem::QGraphicsTextItem(QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsTextItemWithParentScene(parent *QGraphicsItem) *QGraphicsTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),283000,283105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsTextItem{}
	_p.SetDriver(__rv,283000,false)
	return _p
} 
//QGraphicsTextItem::QGraphicsTextItem(QString const&,QGraphicsItem*,QGraphicsScene*)	
func NewGraphicsTextItemWithTextParentScene(text string,parent *QGraphicsItem) *QGraphicsTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),283000,283106,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsTextItem{}
	_p.SetDriver(__rv,283000,false)
	return _p
} 
//QGraphicsTextItem::adjustSize()
func (q *QGraphicsTextItem) AdjustSize()  {
	q.Drv(283000,283107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::boundingRect()
func (q *QGraphicsTextItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(283000,283108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsTextItem::contains(QPointF const&)
func (q *QGraphicsTextItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(283000,283109,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::contextMenuEvent(QGraphicsSceneContextMenuEvent*)
func (q *QGraphicsTextItem) ContextMenuEvent(event *QGraphicsSceneContextMenuEvent)  {
	q.Drv(283000,283110,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::defaultTextColor()
func (q *QGraphicsTextItem) DefaultTextColor() *QColor {
	var __rv uintptr
	q.Drv(283000,283111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QGraphicsTextItem::document()
func (q *QGraphicsTextItem) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(283000,283112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QGraphicsTextItem::dragEnterEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsTextItem) DragEnterEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(283000,283113,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::dragLeaveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsTextItem) DragLeaveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(283000,283114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::dragMoveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsTextItem) DragMoveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(283000,283115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::dropEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsTextItem) DropEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(283000,283116,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::extension(QVariant const&)
func (q *QGraphicsTextItem) Extension(variant *QVariant) *QVariant {
	var __rv uintptr
	q.Drv(283000,283117,Native(variant),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsTextItem::focusInEvent(QFocusEvent*)
func (q *QGraphicsTextItem) FocusInEvent(event *QFocusEvent)  {
	q.Drv(283000,283118,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::focusOutEvent(QFocusEvent*)
func (q *QGraphicsTextItem) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(283000,283119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::font()
func (q *QGraphicsTextItem) Font() *QFont {
	var __rv uintptr
	q.Drv(283000,283120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QGraphicsTextItem::hoverEnterEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsTextItem) HoverEnterEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(283000,283121,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::hoverLeaveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsTextItem) HoverLeaveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(283000,283122,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::hoverMoveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsTextItem) HoverMoveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(283000,283123,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::inputMethodEvent(QInputMethodEvent*)
func (q *QGraphicsTextItem) InputMethodEvent(event *QInputMethodEvent)  {
	q.Drv(283000,283124,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::inputMethodQuery(Qt::InputMethodQuery)
func (q *QGraphicsTextItem) InputMethodQuery(query Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(283000,283125,unsafe.Pointer(&query),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsTextItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsTextItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(283000,283126,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::keyPressEvent(QKeyEvent*)
func (q *QGraphicsTextItem) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(283000,283127,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::keyReleaseEvent(QKeyEvent*)
func (q *QGraphicsTextItem) KeyReleaseEvent(event *QKeyEvent)  {
	q.Drv(283000,283128,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::mouseDoubleClickEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsTextItem) MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(283000,283129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::mouseMoveEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsTextItem) MouseMoveEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(283000,283130,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::mousePressEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsTextItem) MousePressEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(283000,283131,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::mouseReleaseEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsTextItem) MouseReleaseEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(283000,283132,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::opaqueArea()
func (q *QGraphicsTextItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(283000,283133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsTextItem::openExternalLinks()
func (q *QGraphicsTextItem) OpenExternalLinks() bool {
	var __rv bool
	q.Drv(283000,283134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::sceneEvent(QEvent*)
func (q *QGraphicsTextItem) SceneEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(283000,283135,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::setDefaultTextColor(QColor const&)
func (q *QGraphicsTextItem) SetDefaultTextColor(c *QColor)  {
	q.Drv(283000,283136,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setDocument(QTextDocument*)
func (q *QGraphicsTextItem) SetDocument(document *QTextDocument)  {
	q.Drv(283000,283137,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setExtension(QGraphicsItem::Extension,QVariant const&)
func (q *QGraphicsTextItem) SetExtension(extension QGraphicsItem_Extension,variant *QVariant)  {
	q.Drv(283000,283138,unsafe.Pointer(&extension),Native(variant),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setFont(QFont const&)
func (q *QGraphicsTextItem) SetFont(font *QFont)  {
	q.Drv(283000,283139,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setHtml(QString const&)
func (q *QGraphicsTextItem) SetHtml(html string)  {
	q.Drv(283000,283140,unsafe.Pointer(&html),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setOpenExternalLinks(bool)
func (q *QGraphicsTextItem) SetOpenExternalLinks(open bool)  {
	q.Drv(283000,283141,unsafe.Pointer(&open),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setPlainText(QString const&)
func (q *QGraphicsTextItem) SetPlainText(text string)  {
	q.Drv(283000,283142,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setTabChangesFocus(bool)
func (q *QGraphicsTextItem) SetTabChangesFocus(b bool)  {
	q.Drv(283000,283143,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setTextCursor(QTextCursor const&)
func (q *QGraphicsTextItem) SetTextCursor(cursor *QTextCursor)  {
	q.Drv(283000,283144,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>)
func (q *QGraphicsTextItem) SetTextInteractionFlags(flags Qt_TextInteractionFlag)  {
	q.Drv(283000,283145,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::setTextWidth(double)
func (q *QGraphicsTextItem) SetTextWidth(width float64)  {
	q.Drv(283000,283146,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsTextItem::shape()
func (q *QGraphicsTextItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(283000,283147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsTextItem::supportsExtension(QGraphicsItem::Extension)
func (q *QGraphicsTextItem) SupportsExtension(extension QGraphicsItem_Extension) bool {
	var __rv bool
	q.Drv(283000,283148,unsafe.Pointer(&extension),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::tabChangesFocus()
func (q *QGraphicsTextItem) TabChangesFocus() bool {
	var __rv bool
	q.Drv(283000,283149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::textCursor()
func (q *QGraphicsTextItem) TextCursor() *QTextCursor {
	var __rv uintptr
	q.Drv(283000,283150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QGraphicsTextItem::textInteractionFlags()
func (q *QGraphicsTextItem) TextInteractionFlags() Qt_TextInteractionFlag {
	var __rv Qt_TextInteractionFlag
	q.Drv(283000,283151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::textWidth()
func (q *QGraphicsTextItem) TextWidth() float64 {
	var __rv float64
	q.Drv(283000,283152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::toHtml()
func (q *QGraphicsTextItem) ToHtml() string {
	var __rv string
	q.Drv(283000,283153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::toPlainText()
func (q *QGraphicsTextItem) ToPlainText() string {
	var __rv string
	q.Drv(283000,283154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsTextItem::type()
func (q *QGraphicsTextItem) Type() int {
	var __rv int
	q.Drv(283000,283155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
