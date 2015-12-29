// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QObject : QObject
type QObject struct {
	BaseDrv
}
func (q *QObject) OnDestroyed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(314000,314102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QObject) OnDestroyedEx(fn func(*QObject)) uintptr {
	var __rv uintptr
	q.Drv(314000,314103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QObject::QObject()	
func NewObject() *QObject {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),314000,314104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QObject{}
	_p.SetDriver(__rv,314000,false)
	return _p
} 
//QObject::QObject(QObject*)	
func NewObjectWithParent(parent QObjectInterface) *QObject {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),314000,314105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QObject{}
	_p.SetDriver(__rv,314000,false)
	return _p
} 
//QObject::blockSignals(bool)
func (q *QObject) BlockSignals(b bool) bool {
	var __rv bool
	q.Drv(314000,314106,unsafe.Pointer(&b),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::childEvent(QChildEvent*)
func (q *QObject) ChildEvent(value *QChildEvent)  {
	q.Drv(314000,314107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::children()
func (q *QObject) Children() []*QObject {
	var __rv []*QObject
	q.Drv(314000,314108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::customEvent(QEvent*)
func (q *QObject) CustomEvent(value *QEvent)  {
	q.Drv(314000,314109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::deleteLater()
func (q *QObject) DeleteLater()  {
	q.Drv(314000,314110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::dynamicPropertyNames()
func (q *QObject) DynamicPropertyNames() [][]byte {
	var __rv [][]byte
	q.Drv(314000,314111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::event(QEvent*)
func (q *QObject) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(314000,314112,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::findChild(QString const&)
func (q *QObject) FindChild(name string) *QObject {
	var __rv uintptr
	q.Drv(314000,314113,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QObject::findChildren(QString const&)
func (q *QObject) FindChildren(name string) []*QObject {
	var __rv []*QObject
	q.Drv(314000,314114,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::findChildrenWithRegexp(QRegExp const&)
func (q *QObject) FindChildrenWithRegexp(regexp *QRegExp) []*QObject {
	var __rv []*QObject
	q.Drv(314000,314115,Native(regexp),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::inherits(char const*)
func (q *QObject) Inherits(classname string) bool {
	var __rv bool
	q.Drv(314000,314116,unsafe.Pointer(&classname),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::isWidgetType()
func (q *QObject) IsWidgetType() bool {
	var __rv bool
	q.Drv(314000,314117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::killTimer(int)
func (q *QObject) KillTimer(id int)  {
	q.Drv(314000,314118,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::metaObject()
func (q *QObject) MetaObject() *QMetaObject {
	var __rv uintptr
	q.Drv(314000,314119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMetaObject{}
	_rp.SetDriver(__rv,77000,true)
	return _rp
}	
//QObject::objectName()
func (q *QObject) ObjectName() string {
	var __rv string
	q.Drv(314000,314120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::parent()
func (q *QObject) Parent() *QObject {
	var __rv uintptr
	q.Drv(314000,314121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QObject::property(char const*)
func (q *QObject) Property(name string) *QVariant {
	var __rv uintptr
	q.Drv(314000,314122,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QObject::receivers(char const*)
func (q *QObject) Receivers(signal string) int {
	var __rv int
	q.Drv(314000,314123,unsafe.Pointer(&signal),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::registerUserData()	
func QObjectRegisterUserData() uint {
	var __rv uint
	DirectQtDrv(nil,314000,314124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::registerUserData()
func (q *QObject) RegisterUserData() uint {
	var __rv uint
	q.Drv(314000,314124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::sender()
func (q *QObject) Sender() *QObject {
	var __rv uintptr
	q.Drv(314000,314125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QObject::setObjectName(QString const&)
func (q *QObject) SetObjectName(name string)  {
	q.Drv(314000,314126,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::setParent(QObject*)
func (q *QObject) SetParent(value QObjectInterface)  {
	q.Drv(314000,314127,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::setProperty(char const*,QVariant const&)
func (q *QObject) SetProperty(name string,value *QVariant) bool {
	var __rv bool
	q.Drv(314000,314128,unsafe.Pointer(&name),Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::signalsBlocked()
func (q *QObject) SignalsBlocked() bool {
	var __rv bool
	q.Drv(314000,314129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::startTimer(int)
func (q *QObject) StartTimer(interval int) int {
	var __rv int
	q.Drv(314000,314130,unsafe.Pointer(&interval),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::timerEvent(QTimerEvent*)
func (q *QObject) TimerEvent(value *QTimerEvent)  {
	q.Drv(314000,314131,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QObject::tr(char const*)
func (q *QObject) Tr(sourceText string) string {
	var __rv string
	q.Drv(314000,314132,unsafe.Pointer(&sourceText),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QObject::tr(char const*,char const*)
func (q *QObject) TrWithSourcetextDisambiguation(sourceText string,disambiguation string) string {
	var __rv string
	q.Drv(314000,314133,unsafe.Pointer(&sourceText),unsafe.Pointer(&disambiguation),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//  InstallEventFilter
//
//OnEvent(event *QEvent) bool
//
//OnEvent(watched *QObject, event *QEvent) bool
//
//OnTimerEvent(event *QTimerEvent) bool
//
//OnTimerEvent(watched *QObject, event *QTimerEvent) bool
//
//OnMouseButtonPressEvent(event *QMouseButtonPressEvent) bool
//
//OnMouseButtonPressEvent(watched *QObject, event *QMouseButtonPressEvent) bool
//
//OnMouseButtonReleaseEvent(event *QMouseButtonReleaseEvent) bool
//
//OnMouseButtonReleaseEvent(watched *QObject, event *QMouseButtonReleaseEvent) bool
//
//OnMouseButtonDblClickEvent(event *QMouseButtonDblClickEvent) bool
//
//OnMouseButtonDblClickEvent(watched *QObject, event *QMouseButtonDblClickEvent) bool
//
//OnMouseMoveEvent(event *QMouseMoveEvent) bool
//
//OnMouseMoveEvent(watched *QObject, event *QMouseMoveEvent) bool
//
//OnKeyPressEvent(event *QKeyPressEvent) bool
//
//OnKeyPressEvent(watched *QObject, event *QKeyPressEvent) bool
//
//OnKeyReleaseEvent(event *QKeyReleaseEvent) bool
//
//OnKeyReleaseEvent(watched *QObject, event *QKeyReleaseEvent) bool
//
//OnFocusInEvent(event *QFocusInEvent) bool
//
//OnFocusInEvent(watched *QObject, event *QFocusInEvent) bool
//
//OnFocusOutEvent(event *QFocusOutEvent) bool
//
//OnFocusOutEvent(watched *QObject, event *QFocusOutEvent) bool
//
//OnEnterEvent(event *QEnterEvent) bool
//
//OnEnterEvent(watched *QObject, event *QEnterEvent) bool
//
//OnLeaveEvent(event *QLeaveEvent) bool
//
//OnLeaveEvent(watched *QObject, event *QLeaveEvent) bool
//
//OnPaintEvent(event *QPaintEvent) bool
//
//OnPaintEvent(watched *QObject, event *QPaintEvent) bool
//
//OnMoveEvent(event *QMoveEvent) bool
//
//OnMoveEvent(watched *QObject, event *QMoveEvent) bool
//
//OnResizeEvent(event *QResizeEvent) bool
//
//OnResizeEvent(watched *QObject, event *QResizeEvent) bool
//
//OnCreateEvent(event *QCreateEvent) bool
//
//OnCreateEvent(watched *QObject, event *QCreateEvent) bool
//
//OnDestroyEvent(event *QDestroyEvent) bool
//
//OnDestroyEvent(watched *QObject, event *QDestroyEvent) bool
//
//OnShowEvent(event *QShowEvent) bool
//
//OnShowEvent(watched *QObject, event *QShowEvent) bool
//
//OnHideEvent(event *QHideEvent) bool
//
//OnHideEvent(watched *QObject, event *QHideEvent) bool
//
//OnCloseEvent(event *QCloseEvent) bool
//
//OnCloseEvent(watched *QObject, event *QCloseEvent) bool
//
//OnQuitEvent(event *QQuitEvent) bool
//
//OnQuitEvent(watched *QObject, event *QQuitEvent) bool
//
//OnParentChangeEvent(event *QParentChangeEvent) bool
//
//OnParentChangeEvent(watched *QObject, event *QParentChangeEvent) bool
//
//OnParentAboutToChangeEvent(event *QParentAboutToChangeEvent) bool
//
//OnParentAboutToChangeEvent(watched *QObject, event *QParentAboutToChangeEvent) bool
//
//OnThreadChangeEvent(event *QThreadChangeEvent) bool
//
//OnThreadChangeEvent(watched *QObject, event *QThreadChangeEvent) bool
//
//OnWindowActivateEvent(event *QWindowActivateEvent) bool
//
//OnWindowActivateEvent(watched *QObject, event *QWindowActivateEvent) bool
//
//OnWindowDeactivateEvent(event *QWindowDeactivateEvent) bool
//
//OnWindowDeactivateEvent(watched *QObject, event *QWindowDeactivateEvent) bool
//
//OnShowToParentEvent(event *QShowToParentEvent) bool
//
//OnShowToParentEvent(watched *QObject, event *QShowToParentEvent) bool
//
//OnHideToParentEvent(event *QHideToParentEvent) bool
//
//OnHideToParentEvent(watched *QObject, event *QHideToParentEvent) bool
//
//OnWheelEvent(event *QWheelEvent) bool
//
//OnWheelEvent(watched *QObject, event *QWheelEvent) bool
//
//OnWindowTitleChangeEvent(event *QWindowTitleChangeEvent) bool
//
//OnWindowTitleChangeEvent(watched *QObject, event *QWindowTitleChangeEvent) bool
//
//OnWindowIconChangeEvent(event *QWindowIconChangeEvent) bool
//
//OnWindowIconChangeEvent(watched *QObject, event *QWindowIconChangeEvent) bool
//
//OnApplicationWindowIconChangeEvent(event *QApplicationWindowIconChangeEvent) bool
//
//OnApplicationWindowIconChangeEvent(watched *QObject, event *QApplicationWindowIconChangeEvent) bool
//
//OnApplicationFontChangeEvent(event *QApplicationFontChangeEvent) bool
//
//OnApplicationFontChangeEvent(watched *QObject, event *QApplicationFontChangeEvent) bool
//
//OnApplicationLayoutDirectionChangeEvent(event *QApplicationLayoutDirectionChangeEvent) bool
//
//OnApplicationLayoutDirectionChangeEvent(watched *QObject, event *QApplicationLayoutDirectionChangeEvent) bool
//
//OnApplicationPaletteChangeEvent(event *QApplicationPaletteChangeEvent) bool
//
//OnApplicationPaletteChangeEvent(watched *QObject, event *QApplicationPaletteChangeEvent) bool
//
//OnPaletteChangeEvent(event *QPaletteChangeEvent) bool
//
//OnPaletteChangeEvent(watched *QObject, event *QPaletteChangeEvent) bool
//
//OnClipboardEvent(event *QClipboardEvent) bool
//
//OnClipboardEvent(watched *QObject, event *QClipboardEvent) bool
//
//OnSpeechEvent(event *QSpeechEvent) bool
//
//OnSpeechEvent(watched *QObject, event *QSpeechEvent) bool
//
//OnMetaCallEvent(event *QMetaCallEvent) bool
//
//OnMetaCallEvent(watched *QObject, event *QMetaCallEvent) bool
//
//OnSockActEvent(event *QSockActEvent) bool
//
//OnSockActEvent(watched *QObject, event *QSockActEvent) bool
//
//OnWinEventActEvent(event *QWinEventActEvent) bool
//
//OnWinEventActEvent(watched *QObject, event *QWinEventActEvent) bool
//
//OnDeferredDeleteEvent(event *QDeferredDeleteEvent) bool
//
//OnDeferredDeleteEvent(watched *QObject, event *QDeferredDeleteEvent) bool
//
//OnDragEnterEvent(event *QDragEnterEvent) bool
//
//OnDragEnterEvent(watched *QObject, event *QDragEnterEvent) bool
//
//OnDragMoveEvent(event *QDragMoveEvent) bool
//
//OnDragMoveEvent(watched *QObject, event *QDragMoveEvent) bool
//
//OnDragLeaveEvent(event *QDragLeaveEvent) bool
//
//OnDragLeaveEvent(watched *QObject, event *QDragLeaveEvent) bool
//
//OnDropEvent(event *QDropEvent) bool
//
//OnDropEvent(watched *QObject, event *QDropEvent) bool
//
//OnDragResponseEvent(event *QDragResponseEvent) bool
//
//OnDragResponseEvent(watched *QObject, event *QDragResponseEvent) bool
//
//OnChildAddedEvent(event *QChildAddedEvent) bool
//
//OnChildAddedEvent(watched *QObject, event *QChildAddedEvent) bool
//
//OnChildPolishedEvent(event *QChildPolishedEvent) bool
//
//OnChildPolishedEvent(watched *QObject, event *QChildPolishedEvent) bool
//
//OnChildRemovedEvent(event *QChildRemovedEvent) bool
//
//OnChildRemovedEvent(watched *QObject, event *QChildRemovedEvent) bool
//
//OnShowWindowRequestEvent(event *QShowWindowRequestEvent) bool
//
//OnShowWindowRequestEvent(watched *QObject, event *QShowWindowRequestEvent) bool
//
//OnPolishRequestEvent(event *QPolishRequestEvent) bool
//
//OnPolishRequestEvent(watched *QObject, event *QPolishRequestEvent) bool
//
//OnPolishEvent(event *QPolishEvent) bool
//
//OnPolishEvent(watched *QObject, event *QPolishEvent) bool
//
//OnLayoutRequestEvent(event *QLayoutRequestEvent) bool
//
//OnLayoutRequestEvent(watched *QObject, event *QLayoutRequestEvent) bool
//
//OnUpdateRequestEvent(event *QUpdateRequestEvent) bool
//
//OnUpdateRequestEvent(watched *QObject, event *QUpdateRequestEvent) bool
//
//OnUpdateLaterEvent(event *QUpdateLaterEvent) bool
//
//OnUpdateLaterEvent(watched *QObject, event *QUpdateLaterEvent) bool
//
//OnEmbeddingControlEvent(event *QEmbeddingControlEvent) bool
//
//OnEmbeddingControlEvent(watched *QObject, event *QEmbeddingControlEvent) bool
//
//OnActivateControlEvent(event *QActivateControlEvent) bool
//
//OnActivateControlEvent(watched *QObject, event *QActivateControlEvent) bool
//
//OnDeactivateControlEvent(event *QDeactivateControlEvent) bool
//
//OnDeactivateControlEvent(watched *QObject, event *QDeactivateControlEvent) bool
//
//OnContextMenuEvent(event *QContextMenuEvent) bool
//
//OnContextMenuEvent(watched *QObject, event *QContextMenuEvent) bool
//
//OnInputMethodEvent(event *QInputMethodEvent) bool
//
//OnInputMethodEvent(watched *QObject, event *QInputMethodEvent) bool
//
//OnAccessibilityPrepareEvent(event *QAccessibilityPrepareEvent) bool
//
//OnAccessibilityPrepareEvent(watched *QObject, event *QAccessibilityPrepareEvent) bool
//
//OnTabletMoveEvent(event *QTabletMoveEvent) bool
//
//OnTabletMoveEvent(watched *QObject, event *QTabletMoveEvent) bool
//
//OnLocaleChangeEvent(event *QLocaleChangeEvent) bool
//
//OnLocaleChangeEvent(watched *QObject, event *QLocaleChangeEvent) bool
//
//OnLanguageChangeEvent(event *QLanguageChangeEvent) bool
//
//OnLanguageChangeEvent(watched *QObject, event *QLanguageChangeEvent) bool
//
//OnLayoutDirectionChangeEvent(event *QLayoutDirectionChangeEvent) bool
//
//OnLayoutDirectionChangeEvent(watched *QObject, event *QLayoutDirectionChangeEvent) bool
//
//OnStyleEvent(event *QStyleEvent) bool
//
//OnStyleEvent(watched *QObject, event *QStyleEvent) bool
//
//OnTabletPressEvent(event *QTabletPressEvent) bool
//
//OnTabletPressEvent(watched *QObject, event *QTabletPressEvent) bool
//
//OnTabletReleaseEvent(event *QTabletReleaseEvent) bool
//
//OnTabletReleaseEvent(watched *QObject, event *QTabletReleaseEvent) bool
//
//OnOkRequestEvent(event *QOkRequestEvent) bool
//
//OnOkRequestEvent(watched *QObject, event *QOkRequestEvent) bool
//
//OnHelpRequestEvent(event *QHelpRequestEvent) bool
//
//OnHelpRequestEvent(watched *QObject, event *QHelpRequestEvent) bool
//
//OnIconDragEvent(event *QIconDragEvent) bool
//
//OnIconDragEvent(watched *QObject, event *QIconDragEvent) bool
//
//OnFontChangeEvent(event *QFontChangeEvent) bool
//
//OnFontChangeEvent(watched *QObject, event *QFontChangeEvent) bool
//
//OnEnabledChangeEvent(event *QEnabledChangeEvent) bool
//
//OnEnabledChangeEvent(watched *QObject, event *QEnabledChangeEvent) bool
//
//OnActivationChangeEvent(event *QActivationChangeEvent) bool
//
//OnActivationChangeEvent(watched *QObject, event *QActivationChangeEvent) bool
//
//OnStyleChangeEvent(event *QStyleChangeEvent) bool
//
//OnStyleChangeEvent(watched *QObject, event *QStyleChangeEvent) bool
//
//OnIconTextChangeEvent(event *QIconTextChangeEvent) bool
//
//OnIconTextChangeEvent(watched *QObject, event *QIconTextChangeEvent) bool
//
//OnModifiedChangeEvent(event *QModifiedChangeEvent) bool
//
//OnModifiedChangeEvent(watched *QObject, event *QModifiedChangeEvent) bool
//
//OnMouseTrackingChangeEvent(event *QMouseTrackingChangeEvent) bool
//
//OnMouseTrackingChangeEvent(watched *QObject, event *QMouseTrackingChangeEvent) bool
//
//OnWindowBlockedEvent(event *QWindowBlockedEvent) bool
//
//OnWindowBlockedEvent(watched *QObject, event *QWindowBlockedEvent) bool
//
//OnWindowUnblockedEvent(event *QWindowUnblockedEvent) bool
//
//OnWindowUnblockedEvent(watched *QObject, event *QWindowUnblockedEvent) bool
//
//OnWindowStateChangeEvent(event *QWindowStateChangeEvent) bool
//
//OnWindowStateChangeEvent(watched *QObject, event *QWindowStateChangeEvent) bool
//
//OnToolTipEvent(event *QToolTipEvent) bool
//
//OnToolTipEvent(watched *QObject, event *QToolTipEvent) bool
//
//OnWhatsThisEvent(event *QWhatsThisEvent) bool
//
//OnWhatsThisEvent(watched *QObject, event *QWhatsThisEvent) bool
//
//OnStatusTipEvent(event *QStatusTipEvent) bool
//
//OnStatusTipEvent(watched *QObject, event *QStatusTipEvent) bool
//
//OnActionChangedEvent(event *QActionChangedEvent) bool
//
//OnActionChangedEvent(watched *QObject, event *QActionChangedEvent) bool
//
//OnActionAddedEvent(event *QActionAddedEvent) bool
//
//OnActionAddedEvent(watched *QObject, event *QActionAddedEvent) bool
//
//OnActionRemovedEvent(event *QActionRemovedEvent) bool
//
//OnActionRemovedEvent(watched *QObject, event *QActionRemovedEvent) bool
//
//OnFileOpenEvent(event *QFileOpenEvent) bool
//
//OnFileOpenEvent(watched *QObject, event *QFileOpenEvent) bool
//
//OnShortcutEvent(event *QShortcutEvent) bool
//
//OnShortcutEvent(watched *QObject, event *QShortcutEvent) bool
//
//OnShortcutOverrideEvent(event *QShortcutOverrideEvent) bool
//
//OnShortcutOverrideEvent(watched *QObject, event *QShortcutOverrideEvent) bool
//
//OnWhatsThisClickedEvent(event *QWhatsThisClickedEvent) bool
//
//OnWhatsThisClickedEvent(watched *QObject, event *QWhatsThisClickedEvent) bool
//
//OnToolBarChangeEvent(event *QToolBarChangeEvent) bool
//
//OnToolBarChangeEvent(watched *QObject, event *QToolBarChangeEvent) bool
//
//OnApplicationActivateEvent(event *QApplicationActivateEvent) bool
//
//OnApplicationActivateEvent(watched *QObject, event *QApplicationActivateEvent) bool
//
//OnApplicationDeactivateEvent(event *QApplicationDeactivateEvent) bool
//
//OnApplicationDeactivateEvent(watched *QObject, event *QApplicationDeactivateEvent) bool
//
//OnQueryWhatsThisEvent(event *QQueryWhatsThisEvent) bool
//
//OnQueryWhatsThisEvent(watched *QObject, event *QQueryWhatsThisEvent) bool
//
//OnEnterWhatsThisModeEvent(event *QEnterWhatsThisModeEvent) bool
//
//OnEnterWhatsThisModeEvent(watched *QObject, event *QEnterWhatsThisModeEvent) bool
//
//OnLeaveWhatsThisModeEvent(event *QLeaveWhatsThisModeEvent) bool
//
//OnLeaveWhatsThisModeEvent(watched *QObject, event *QLeaveWhatsThisModeEvent) bool
//
//OnZOrderChangeEvent(event *QZOrderChangeEvent) bool
//
//OnZOrderChangeEvent(watched *QObject, event *QZOrderChangeEvent) bool
//
//OnHoverEnterEvent(event *QHoverEnterEvent) bool
//
//OnHoverEnterEvent(watched *QObject, event *QHoverEnterEvent) bool
//
//OnHoverLeaveEvent(event *QHoverLeaveEvent) bool
//
//OnHoverLeaveEvent(watched *QObject, event *QHoverLeaveEvent) bool
//
//OnHoverMoveEvent(event *QHoverMoveEvent) bool
//
//OnHoverMoveEvent(watched *QObject, event *QHoverMoveEvent) bool
//
//OnAccessibilityHelpEvent(event *QAccessibilityHelpEvent) bool
//
//OnAccessibilityHelpEvent(watched *QObject, event *QAccessibilityHelpEvent) bool
//
//OnAccessibilityDescriptionEvent(event *QAccessibilityDescriptionEvent) bool
//
//OnAccessibilityDescriptionEvent(watched *QObject, event *QAccessibilityDescriptionEvent) bool
//
//OnAcceptDropsChangeEvent(event *QAcceptDropsChangeEvent) bool
//
//OnAcceptDropsChangeEvent(watched *QObject, event *QAcceptDropsChangeEvent) bool
//
//OnMenubarUpdatedEvent(event *QMenubarUpdatedEvent) bool
//
//OnMenubarUpdatedEvent(watched *QObject, event *QMenubarUpdatedEvent) bool
//
//OnZeroTimerEventEvent(event *QZeroTimerEventEvent) bool
//
//OnZeroTimerEventEvent(watched *QObject, event *QZeroTimerEventEvent) bool
//
//OnGraphicsSceneMouseMoveEvent(event *QGraphicsSceneMouseMoveEvent) bool
//
//OnGraphicsSceneMouseMoveEvent(watched *QObject, event *QGraphicsSceneMouseMoveEvent) bool
//
//OnGraphicsSceneMousePressEvent(event *QGraphicsSceneMousePressEvent) bool
//
//OnGraphicsSceneMousePressEvent(watched *QObject, event *QGraphicsSceneMousePressEvent) bool
//
//OnGraphicsSceneMouseReleaseEvent(event *QGraphicsSceneMouseReleaseEvent) bool
//
//OnGraphicsSceneMouseReleaseEvent(watched *QObject, event *QGraphicsSceneMouseReleaseEvent) bool
//
//OnGraphicsSceneMouseDoubleClickEvent(event *QGraphicsSceneMouseDoubleClickEvent) bool
//
//OnGraphicsSceneMouseDoubleClickEvent(watched *QObject, event *QGraphicsSceneMouseDoubleClickEvent) bool
//
//OnGraphicsSceneContextMenuEvent(event *QGraphicsSceneContextMenuEvent) bool
//
//OnGraphicsSceneContextMenuEvent(watched *QObject, event *QGraphicsSceneContextMenuEvent) bool
//
//OnGraphicsSceneHoverEnterEvent(event *QGraphicsSceneHoverEnterEvent) bool
//
//OnGraphicsSceneHoverEnterEvent(watched *QObject, event *QGraphicsSceneHoverEnterEvent) bool
//
//OnGraphicsSceneHoverMoveEvent(event *QGraphicsSceneHoverMoveEvent) bool
//
//OnGraphicsSceneHoverMoveEvent(watched *QObject, event *QGraphicsSceneHoverMoveEvent) bool
//
//OnGraphicsSceneHoverLeaveEvent(event *QGraphicsSceneHoverLeaveEvent) bool
//
//OnGraphicsSceneHoverLeaveEvent(watched *QObject, event *QGraphicsSceneHoverLeaveEvent) bool
//
//OnGraphicsSceneHelpEvent(event *QGraphicsSceneHelpEvent) bool
//
//OnGraphicsSceneHelpEvent(watched *QObject, event *QGraphicsSceneHelpEvent) bool
//
//OnGraphicsSceneDragEnterEvent(event *QGraphicsSceneDragEnterEvent) bool
//
//OnGraphicsSceneDragEnterEvent(watched *QObject, event *QGraphicsSceneDragEnterEvent) bool
//
//OnGraphicsSceneDragMoveEvent(event *QGraphicsSceneDragMoveEvent) bool
//
//OnGraphicsSceneDragMoveEvent(watched *QObject, event *QGraphicsSceneDragMoveEvent) bool
//
//OnGraphicsSceneDragLeaveEvent(event *QGraphicsSceneDragLeaveEvent) bool
//
//OnGraphicsSceneDragLeaveEvent(watched *QObject, event *QGraphicsSceneDragLeaveEvent) bool
//
//OnGraphicsSceneDropEvent(event *QGraphicsSceneDropEvent) bool
//
//OnGraphicsSceneDropEvent(watched *QObject, event *QGraphicsSceneDropEvent) bool
//
//OnGraphicsSceneWheelEvent(event *QGraphicsSceneWheelEvent) bool
//
//OnGraphicsSceneWheelEvent(watched *QObject, event *QGraphicsSceneWheelEvent) bool
//
//OnKeyboardLayoutChangeEvent(event *QKeyboardLayoutChangeEvent) bool
//
//OnKeyboardLayoutChangeEvent(watched *QObject, event *QKeyboardLayoutChangeEvent) bool
//
//OnDynamicPropertyChangeEvent(event *QDynamicPropertyChangeEvent) bool
//
//OnDynamicPropertyChangeEvent(watched *QObject, event *QDynamicPropertyChangeEvent) bool
//
//OnTabletEnterProximityEvent(event *QTabletEnterProximityEvent) bool
//
//OnTabletEnterProximityEvent(watched *QObject, event *QTabletEnterProximityEvent) bool
//
//OnTabletLeaveProximityEvent(event *QTabletLeaveProximityEvent) bool
//
//OnTabletLeaveProximityEvent(watched *QObject, event *QTabletLeaveProximityEvent) bool
//
//OnNonClientAreaMouseMoveEvent(event *QNonClientAreaMouseMoveEvent) bool
//
//OnNonClientAreaMouseMoveEvent(watched *QObject, event *QNonClientAreaMouseMoveEvent) bool
//
//OnNonClientAreaMouseButtonPressEvent(event *QNonClientAreaMouseButtonPressEvent) bool
//
//OnNonClientAreaMouseButtonPressEvent(watched *QObject, event *QNonClientAreaMouseButtonPressEvent) bool
//
//OnNonClientAreaMouseButtonReleaseEvent(event *QNonClientAreaMouseButtonReleaseEvent) bool
//
//OnNonClientAreaMouseButtonReleaseEvent(watched *QObject, event *QNonClientAreaMouseButtonReleaseEvent) bool
//
//OnNonClientAreaMouseButtonDblClickEvent(event *QNonClientAreaMouseButtonDblClickEvent) bool
//
//OnNonClientAreaMouseButtonDblClickEvent(watched *QObject, event *QNonClientAreaMouseButtonDblClickEvent) bool
//
//OnMacSizeChangeEvent(event *QMacSizeChangeEvent) bool
//
//OnMacSizeChangeEvent(watched *QObject, event *QMacSizeChangeEvent) bool
//
//OnContentsRectChangeEvent(event *QContentsRectChangeEvent) bool
//
//OnContentsRectChangeEvent(watched *QObject, event *QContentsRectChangeEvent) bool
//
//OnMacGLWindowChangeEvent(event *QMacGLWindowChangeEvent) bool
//
//OnMacGLWindowChangeEvent(watched *QObject, event *QMacGLWindowChangeEvent) bool
//
//OnFutureCallOutEvent(event *QFutureCallOutEvent) bool
//
//OnFutureCallOutEvent(watched *QObject, event *QFutureCallOutEvent) bool
//
//OnGraphicsSceneResizeEvent(event *QGraphicsSceneResizeEvent) bool
//
//OnGraphicsSceneResizeEvent(watched *QObject, event *QGraphicsSceneResizeEvent) bool
//
//OnGraphicsSceneMoveEvent(event *QGraphicsSceneMoveEvent) bool
//
//OnGraphicsSceneMoveEvent(watched *QObject, event *QGraphicsSceneMoveEvent) bool
//
//OnCursorChangeEvent(event *QCursorChangeEvent) bool
//
//OnCursorChangeEvent(watched *QObject, event *QCursorChangeEvent) bool
//
//OnToolTipChangeEvent(event *QToolTipChangeEvent) bool
//
//OnToolTipChangeEvent(watched *QObject, event *QToolTipChangeEvent) bool
//
//OnNetworkReplyUpdatedEvent(event *QNetworkReplyUpdatedEvent) bool
//
//OnNetworkReplyUpdatedEvent(watched *QObject, event *QNetworkReplyUpdatedEvent) bool
//
//OnGrabMouseEvent(event *QGrabMouseEvent) bool
//
//OnGrabMouseEvent(watched *QObject, event *QGrabMouseEvent) bool
//
//OnUngrabMouseEvent(event *QUngrabMouseEvent) bool
//
//OnUngrabMouseEvent(watched *QObject, event *QUngrabMouseEvent) bool
//
//OnGrabKeyboardEvent(event *QGrabKeyboardEvent) bool
//
//OnGrabKeyboardEvent(watched *QObject, event *QGrabKeyboardEvent) bool
//
//OnUngrabKeyboardEvent(event *QUngrabKeyboardEvent) bool
//
//OnUngrabKeyboardEvent(watched *QObject, event *QUngrabKeyboardEvent) bool
//
//OnMacGLClearDrawableEvent(event *QMacGLClearDrawableEvent) bool
//
//OnMacGLClearDrawableEvent(watched *QObject, event *QMacGLClearDrawableEvent) bool
//
//OnStateMachineSignalEvent(event *QStateMachineSignalEvent) bool
//
//OnStateMachineSignalEvent(watched *QObject, event *QStateMachineSignalEvent) bool
//
//OnStateMachineWrappedEvent(event *QStateMachineWrappedEvent) bool
//
//OnStateMachineWrappedEvent(watched *QObject, event *QStateMachineWrappedEvent) bool
//
//OnTouchBeginEvent(event *QTouchBeginEvent) bool
//
//OnTouchBeginEvent(watched *QObject, event *QTouchBeginEvent) bool
//
//OnTouchUpdateEvent(event *QTouchUpdateEvent) bool
//
//OnTouchUpdateEvent(watched *QObject, event *QTouchUpdateEvent) bool
//
//OnTouchEndEvent(event *QTouchEndEvent) bool
//
//OnTouchEndEvent(watched *QObject, event *QTouchEndEvent) bool
//
//OnNativeGestureEvent(event *QNativeGestureEvent) bool
//
//OnNativeGestureEvent(watched *QObject, event *QNativeGestureEvent) bool
//
//OnRequestSoftwareInputPanelEvent(event *QRequestSoftwareInputPanelEvent) bool
//
//OnRequestSoftwareInputPanelEvent(watched *QObject, event *QRequestSoftwareInputPanelEvent) bool
//
//OnCloseSoftwareInputPanelEvent(event *QCloseSoftwareInputPanelEvent) bool
//
//OnCloseSoftwareInputPanelEvent(watched *QObject, event *QCloseSoftwareInputPanelEvent) bool
//
//OnUpdateSoftKeysEvent(event *QUpdateSoftKeysEvent) bool
//
//OnUpdateSoftKeysEvent(watched *QObject, event *QUpdateSoftKeysEvent) bool
//
//OnWinIdChangeEvent(event *QWinIdChangeEvent) bool
//
//OnWinIdChangeEvent(watched *QObject, event *QWinIdChangeEvent) bool
//
//OnGestureEvent(event *QGestureEvent) bool
//
//OnGestureEvent(watched *QObject, event *QGestureEvent) bool
//
//OnGestureOverrideEvent(event *QGestureOverrideEvent) bool
//
//OnGestureOverrideEvent(watched *QObject, event *QGestureOverrideEvent) bool
//
//OnUserEvent(event *QEvent) bool
//
//OnUserEvent(watched *QObject, event *QEvent) bool
func (q *QObject) InstallEventFilter(filter interface{}) {
	drvInstallEventFilter(q, filter)
}

