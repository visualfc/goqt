// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsWidget_enum_1 - QGraphicsWidget::enum_1
type QGraphicsWidget_enum_1 uint32
const (
	QGraphicsWidget_Type QGraphicsWidget_enum_1 = 11
)
//struct QGraphicsWidget : QGraphicsWidget
type QGraphicsWidget struct {
	QGraphicsObject
}
func (q *QGraphicsWidget) OnLayoutChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(286000,286102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsWidget) OnGeometryChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(286000,286103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsWidget::QGraphicsWidget()	
func NewGraphicsWidget() *QGraphicsWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),286000,286104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsWidget{}
	_p.SetDriver(__rv,286000,false)
	return _p
} 
//QGraphicsWidget::QGraphicsWidget(QGraphicsItem*,QFlags<Qt::WindowType>)	
func NewGraphicsWidgetWithParentFlags(parent *QGraphicsItem,wFlags Qt_WindowType) *QGraphicsWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),286000,286105,Native(parent),unsafe.Pointer(&wFlags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsWidget{}
	_p.SetDriver(__rv,286000,false)
	return _p
} 
//QGraphicsWidget::actions()
func (q *QGraphicsWidget) Actions() []*QAction {
	var __rv []*QAction
	q.Drv(286000,286106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::addAction(QAction*)
func (q *QGraphicsWidget) AddAction(action *QAction)  {
	q.Drv(286000,286107,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::addActions(QList<QAction*>)
func (q *QGraphicsWidget) AddActions(actions []*QAction)  {
	q.Drv(286000,286108,unsafe.Pointer(&actions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::adjustSize()
func (q *QGraphicsWidget) AdjustSize()  {
	q.Drv(286000,286109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::autoFillBackground()
func (q *QGraphicsWidget) AutoFillBackground() bool {
	var __rv bool
	q.Drv(286000,286110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::boundingRect()
func (q *QGraphicsWidget) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(286000,286111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsWidget::changeEvent(QEvent*)
func (q *QGraphicsWidget) ChangeEvent(event *QEvent)  {
	q.Drv(286000,286112,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::children()
func (q *QGraphicsWidget) Children() []*QObject {
	var __rv []*QObject
	q.Drv(286000,286113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::close()
func (q *QGraphicsWidget) Close() bool {
	var __rv bool
	q.Drv(286000,286114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::closeEvent(QCloseEvent*)
func (q *QGraphicsWidget) CloseEvent(event *QCloseEvent)  {
	q.Drv(286000,286115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::event(QEvent*)
func (q *QGraphicsWidget) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(286000,286116,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::focusInEvent(QFocusEvent*)
func (q *QGraphicsWidget) FocusInEvent(event *QFocusEvent)  {
	q.Drv(286000,286117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::focusNextPrevChild(bool)
func (q *QGraphicsWidget) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(286000,286118,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::focusOutEvent(QFocusEvent*)
func (q *QGraphicsWidget) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(286000,286119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::focusPolicy()
func (q *QGraphicsWidget) FocusPolicy() Qt_FocusPolicy {
	var __rv Qt_FocusPolicy
	q.Drv(286000,286120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::focusWidget()
func (q *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	var __rv uintptr
	q.Drv(286000,286121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsWidget{}
	_rp.SetDriver(__rv,286000,false)
	return _rp
}	
//QGraphicsWidget::font()
func (q *QGraphicsWidget) Font() *QFont {
	var __rv uintptr
	q.Drv(286000,286122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QGraphicsWidget::getContentsMargins(double*,double*,double*,double*)
func (q *QGraphicsWidget) GetContentsMargins(left *float64,top *float64,right *float64,bottom *float64)  {
	q.Drv(286000,286123,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::getWindowFrameMargins(double*,double*,double*,double*)
func (q *QGraphicsWidget) GetWindowFrameMargins(left *float64,top *float64,right *float64,bottom *float64)  {
	q.Drv(286000,286124,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::grabKeyboardEvent(QEvent*)
func (q *QGraphicsWidget) GrabKeyboardEvent(event *QEvent)  {
	q.Drv(286000,286125,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::grabMouseEvent(QEvent*)
func (q *QGraphicsWidget) GrabMouseEvent(event *QEvent)  {
	q.Drv(286000,286126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::grabShortcut(QKeySequence const&)
func (q *QGraphicsWidget) GrabShortcut(sequence *QKeySequence) int {
	var __rv int
	q.Drv(286000,286127,Native(sequence),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::grabShortcut(QKeySequence const&,Qt::ShortcutContext)
func (q *QGraphicsWidget) GrabShortcutWithSequenceContext(sequence *QKeySequence,context Qt_ShortcutContext) int {
	var __rv int
	q.Drv(286000,286128,Native(sequence),unsafe.Pointer(&context),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::hideEvent(QHideEvent*)
func (q *QGraphicsWidget) HideEvent(event *QHideEvent)  {
	q.Drv(286000,286129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::hoverLeaveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsWidget) HoverLeaveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(286000,286130,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::hoverMoveEvent(QGraphicsSceneHoverEvent*)
func (q *QGraphicsWidget) HoverMoveEvent(event *QGraphicsSceneHoverEvent)  {
	q.Drv(286000,286131,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::insertAction(QAction*,QAction*)
func (q *QGraphicsWidget) InsertAction(before *QAction,action *QAction)  {
	q.Drv(286000,286132,Native(before),Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::insertActions(QAction*,QList<QAction*>)
func (q *QGraphicsWidget) InsertActions(before *QAction,actions []*QAction)  {
	q.Drv(286000,286133,Native(before),unsafe.Pointer(&actions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::isActiveWindow()
func (q *QGraphicsWidget) IsActiveWindow() bool {
	var __rv bool
	q.Drv(286000,286134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::itemChange(QGraphicsItem::GraphicsItemChange,QVariant const&)
func (q *QGraphicsWidget) ItemChange(change QGraphicsItem_GraphicsItemChange,value *QVariant) *QVariant {
	var __rv uintptr
	q.Drv(286000,286135,unsafe.Pointer(&change),Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsWidget::layout()
func (q *QGraphicsWidget) Layout() *QGraphicsLayout {
	var __rv uintptr
	q.Drv(286000,286136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayout{}
	_rp.SetDriver(__rv,259000,true)
	return _rp
}	
//QGraphicsWidget::layoutDirection()
func (q *QGraphicsWidget) LayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(286000,286137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::moveEvent(QGraphicsSceneMoveEvent*)
func (q *QGraphicsWidget) MoveEvent(event *QGraphicsSceneMoveEvent)  {
	q.Drv(286000,286138,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::palette()
func (q *QGraphicsWidget) Palette() *QPalette {
	var __rv uintptr
	q.Drv(286000,286139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QGraphicsWidget::polishEvent()
func (q *QGraphicsWidget) PolishEvent()  {
	q.Drv(286000,286140,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::propertyChange(QString const&,QVariant const&)
func (q *QGraphicsWidget) PropertyChange(propertyName string,value *QVariant) *QVariant {
	var __rv uintptr
	q.Drv(286000,286141,unsafe.Pointer(&propertyName),Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsWidget::rect()
func (q *QGraphicsWidget) Rect() *QRectF {
	var __rv uintptr
	q.Drv(286000,286142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsWidget::releaseShortcut(int)
func (q *QGraphicsWidget) ReleaseShortcut(id int)  {
	q.Drv(286000,286143,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::removeAction(QAction*)
func (q *QGraphicsWidget) RemoveAction(action *QAction)  {
	q.Drv(286000,286144,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::resize(QSizeF const&)
func (q *QGraphicsWidget) Resize(size *QSizeF)  {
	q.Drv(286000,286145,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::resize(double,double)
func (q *QGraphicsWidget) ResizeFWithWidthHeight(w float64,h float64)  {
	q.Drv(286000,286146,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::resizeEvent(QGraphicsSceneResizeEvent*)
func (q *QGraphicsWidget) ResizeEvent(event *QGraphicsSceneResizeEvent)  {
	q.Drv(286000,286147,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::sceneEvent(QEvent*)
func (q *QGraphicsWidget) SceneEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(286000,286148,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::setAttribute(Qt::WidgetAttribute)
func (q *QGraphicsWidget) SetAttribute(attribute Qt_WidgetAttribute)  {
	q.Drv(286000,286149,unsafe.Pointer(&attribute),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setAttribute(Qt::WidgetAttribute,bool)
func (q *QGraphicsWidget) SetAttributeWithAttributeOn(attribute Qt_WidgetAttribute,on bool)  {
	q.Drv(286000,286150,unsafe.Pointer(&attribute),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setAutoFillBackground(bool)
func (q *QGraphicsWidget) SetAutoFillBackground(enabled bool)  {
	q.Drv(286000,286151,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setContentsMargins(double,double,double,double)
func (q *QGraphicsWidget) SetContentsMargins(left float64,top float64,right float64,bottom float64)  {
	q.Drv(286000,286152,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setFocusPolicy(Qt::FocusPolicy)
func (q *QGraphicsWidget) SetFocusPolicy(policy Qt_FocusPolicy)  {
	q.Drv(286000,286153,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setFont(QFont const&)
func (q *QGraphicsWidget) SetFont(font *QFont)  {
	q.Drv(286000,286154,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setGeometry(QRectF const&)
func (q *QGraphicsWidget) SetGeometry(rect *QRectF)  {
	q.Drv(286000,286155,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setGeometry(double,double,double,double)
func (q *QGraphicsWidget) SetGeometryFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(286000,286156,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setLayout(QGraphicsLayout*)
func (q *QGraphicsWidget) SetLayout(layout *QGraphicsLayout)  {
	q.Drv(286000,286157,Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setLayoutDirection(Qt::LayoutDirection)
func (q *QGraphicsWidget) SetLayoutDirection(direction Qt_LayoutDirection)  {
	q.Drv(286000,286158,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setPalette(QPalette const&)
func (q *QGraphicsWidget) SetPalette(palette *QPalette)  {
	q.Drv(286000,286159,Native(palette),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setShortcutAutoRepeat(int)
func (q *QGraphicsWidget) SetShortcutAutoRepeat(id int)  {
	q.Drv(286000,286160,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setShortcutAutoRepeat(int,bool)
func (q *QGraphicsWidget) SetShortcutAutoRepeatWithIdEnabled(id int,enabled bool)  {
	q.Drv(286000,286161,unsafe.Pointer(&id),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setShortcutEnabled(int)
func (q *QGraphicsWidget) SetShortcutEnabled(id int)  {
	q.Drv(286000,286162,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setShortcutEnabled(int,bool)
func (q *QGraphicsWidget) SetShortcutEnabledWithIdEnabled(id int,enabled bool)  {
	q.Drv(286000,286163,unsafe.Pointer(&id),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setStyle(QStyle*)
func (q *QGraphicsWidget) SetStyle(style *QStyle)  {
	q.Drv(286000,286164,Native(style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setTabOrder(QGraphicsWidget*,QGraphicsWidget*)	
func QGraphicsWidgetSetTabOrder(first *QGraphicsWidget,second *QGraphicsWidget)  {
	DirectQtDrv(nil,286000,286165,Native(first),Native(second),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setTabOrder(QGraphicsWidget*,QGraphicsWidget*)
func (q *QGraphicsWidget) SetTabOrder(first *QGraphicsWidget,second *QGraphicsWidget)  {
	q.Drv(286000,286165,Native(first),Native(second),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setWindowFlags(QFlags<Qt::WindowType>)
func (q *QGraphicsWidget) SetWindowFlags(wFlags Qt_WindowType)  {
	q.Drv(286000,286166,unsafe.Pointer(&wFlags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setWindowFrameMargins(double,double,double,double)
func (q *QGraphicsWidget) SetWindowFrameMargins(left float64,top float64,right float64,bottom float64)  {
	q.Drv(286000,286167,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::setWindowTitle(QString const&)
func (q *QGraphicsWidget) SetWindowTitle(title string)  {
	q.Drv(286000,286168,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::shape()
func (q *QGraphicsWidget) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(286000,286169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsWidget::showEvent(QShowEvent*)
func (q *QGraphicsWidget) ShowEvent(event *QShowEvent)  {
	q.Drv(286000,286170,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::size()
func (q *QGraphicsWidget) Size() *QSizeF {
	var __rv uintptr
	q.Drv(286000,286171,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsWidget::sizeHint(Qt::SizeHint,QSizeF const&)
func (q *QGraphicsWidget) SizeHint(which Qt_SizeHint,constraint *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(286000,286172,unsafe.Pointer(&which),Native(constraint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QGraphicsWidget::style()
func (q *QGraphicsWidget) Style() *QStyle {
	var __rv uintptr
	q.Drv(286000,286173,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QGraphicsWidget::testAttribute(Qt::WidgetAttribute)
func (q *QGraphicsWidget) TestAttribute(attribute Qt_WidgetAttribute) bool {
	var __rv bool
	q.Drv(286000,286174,unsafe.Pointer(&attribute),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::type()
func (q *QGraphicsWidget) Type() int {
	var __rv int
	q.Drv(286000,286175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::ungrabKeyboardEvent(QEvent*)
func (q *QGraphicsWidget) UngrabKeyboardEvent(event *QEvent)  {
	q.Drv(286000,286176,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::ungrabMouseEvent(QEvent*)
func (q *QGraphicsWidget) UngrabMouseEvent(event *QEvent)  {
	q.Drv(286000,286177,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::unsetLayoutDirection()
func (q *QGraphicsWidget) UnsetLayoutDirection()  {
	q.Drv(286000,286178,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::unsetWindowFrameMargins()
func (q *QGraphicsWidget) UnsetWindowFrameMargins()  {
	q.Drv(286000,286179,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::updateGeometry()
func (q *QGraphicsWidget) UpdateGeometry()  {
	q.Drv(286000,286180,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsWidget::windowFlags()
func (q *QGraphicsWidget) WindowFlags() Qt_WindowType {
	var __rv Qt_WindowType
	q.Drv(286000,286181,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::windowFrameEvent(QEvent*)
func (q *QGraphicsWidget) WindowFrameEvent(e *QEvent) bool {
	var __rv bool
	q.Drv(286000,286182,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::windowFrameGeometry()
func (q *QGraphicsWidget) WindowFrameGeometry() *QRectF {
	var __rv uintptr
	q.Drv(286000,286183,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsWidget::windowFrameRect()
func (q *QGraphicsWidget) WindowFrameRect() *QRectF {
	var __rv uintptr
	q.Drv(286000,286184,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsWidget::windowFrameSectionAt(QPointF const&)
func (q *QGraphicsWidget) WindowFrameSectionAt(pos *QPointF) Qt_WindowFrameSection {
	var __rv Qt_WindowFrameSection
	q.Drv(286000,286185,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::windowTitle()
func (q *QGraphicsWidget) WindowTitle() string {
	var __rv string
	q.Drv(286000,286186,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsWidget::windowType()
func (q *QGraphicsWidget) WindowType() Qt_WindowType {
	var __rv Qt_WindowType
	q.Drv(286000,286187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
