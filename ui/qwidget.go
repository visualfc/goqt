// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QWidget_RenderFlag - QWidget::RenderFlag
type QWidget_RenderFlag uint32
const (
	QWidget_DrawWindowBackground QWidget_RenderFlag = 0x1
	QWidget_DrawChildren QWidget_RenderFlag = 0x2
	QWidget_IgnoreMask QWidget_RenderFlag = 0x4
)
//struct QWidget : QWidget
type QWidget struct {
	QObject
}
func (q *QWidget) OnCustomContextMenuRequested(fn func(*QPoint)) uintptr {
	var __rv uintptr
	q.Drv(395000,395102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QWidget::QWidget()	
func NewWidget() *QWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),395000,395103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWidget{}
	_p.SetDriver(__rv,395000,false)
	return _p
} 
//QWidget::QWidget(QWidget*,QFlags<Qt::WindowType>)	
func NewWidgetWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),395000,395104,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWidget{}
	_p.SetDriver(__rv,395000,false)
	return _p
} 
//QWidget::acceptDrops()
func (q *QWidget) AcceptDrops() bool {
	var __rv bool
	q.Drv(395000,395105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::accessibleDescription()
func (q *QWidget) AccessibleDescription() string {
	var __rv string
	q.Drv(395000,395106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::accessibleName()
func (q *QWidget) AccessibleName() string {
	var __rv string
	q.Drv(395000,395107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::actionEvent(QActionEvent*)
func (q *QWidget) ActionEvent(value *QActionEvent)  {
	q.Drv(395000,395108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::actions()
func (q *QWidget) Actions() []*QAction {
	var __rv []*QAction
	q.Drv(395000,395109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::activateWindow()
func (q *QWidget) ActivateWindow()  {
	q.Drv(395000,395110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::addAction(QAction*)
func (q *QWidget) AddAction(action *QAction)  {
	q.Drv(395000,395111,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::addActions(QList<QAction*>)
func (q *QWidget) AddActions(actions []*QAction)  {
	q.Drv(395000,395112,unsafe.Pointer(&actions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::adjustSize()
func (q *QWidget) AdjustSize()  {
	q.Drv(395000,395113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::autoFillBackground()
func (q *QWidget) AutoFillBackground() bool {
	var __rv bool
	q.Drv(395000,395114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::backgroundRole()
func (q *QWidget) BackgroundRole() QPalette_ColorRole {
	var __rv QPalette_ColorRole
	q.Drv(395000,395115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::baseSize()
func (q *QWidget) BaseSize() *QSize {
	var __rv uintptr
	q.Drv(395000,395116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::changeEvent(QEvent*)
func (q *QWidget) ChangeEvent(value *QEvent)  {
	q.Drv(395000,395117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::childAt(QPoint const&)
func (q *QWidget) ChildAt(p *QPoint) *QWidget {
	var __rv uintptr
	q.Drv(395000,395118,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::childAt(int,int)
func (q *QWidget) ChildAtWithXY(x int,y int) *QWidget {
	var __rv uintptr
	q.Drv(395000,395119,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::childrenRect()
func (q *QWidget) ChildrenRect() *QRect {
	var __rv uintptr
	q.Drv(395000,395120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::childrenRegion()
func (q *QWidget) ChildrenRegion() *QRegion {
	var __rv uintptr
	q.Drv(395000,395121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QWidget::clearFocus()
func (q *QWidget) ClearFocus()  {
	q.Drv(395000,395122,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::clearMask()
func (q *QWidget) ClearMask()  {
	q.Drv(395000,395123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::close()
func (q *QWidget) Close() bool {
	var __rv bool
	q.Drv(395000,395124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::closeEvent(QCloseEvent*)
func (q *QWidget) CloseEvent(value *QCloseEvent)  {
	q.Drv(395000,395125,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintDevice::colorCount()
func (q *QWidget) ColorCount() int {
	var __rv int
	q.Drv(395000,395126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::contentsMargins()
func (q *QWidget) ContentsMargins() *QMargins {
	var __rv uintptr
	q.Drv(395000,395127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMargins{}
	_rp.SetDriver(__rv,73000,true)
	return _rp
}	
//QWidget::contentsRect()
func (q *QWidget) ContentsRect() *QRect {
	var __rv uintptr
	q.Drv(395000,395128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::contextMenuEvent(QContextMenuEvent*)
func (q *QWidget) ContextMenuEvent(value *QContextMenuEvent)  {
	q.Drv(395000,395129,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::contextMenuPolicy()
func (q *QWidget) ContextMenuPolicy() Qt_ContextMenuPolicy {
	var __rv Qt_ContextMenuPolicy
	q.Drv(395000,395130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::createWinId()
func (q *QWidget) CreateWinId()  {
	q.Drv(395000,395131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::cursor()
func (q *QWidget) Cursor() *QCursor {
	var __rv uintptr
	q.Drv(395000,395132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCursor{}
	_rp.SetDriver(__rv,18000,true)
	return _rp
}	
//QPaintDevice::depth()
func (q *QWidget) Depth() int {
	var __rv int
	q.Drv(395000,395133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::destroy(bool,bool)
func (q *QWidget) Destroy(destroyWindow bool,destroySubWindows bool)  {
	q.Drv(395000,395134,unsafe.Pointer(&destroyWindow),unsafe.Pointer(&destroySubWindows),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::devType()
func (q *QWidget) DevType() int {
	var __rv int
	q.Drv(395000,395135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::dragEnterEvent(QDragEnterEvent*)
func (q *QWidget) DragEnterEvent(value *QDragEnterEvent)  {
	q.Drv(395000,395136,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::dragLeaveEvent(QDragLeaveEvent*)
func (q *QWidget) DragLeaveEvent(value *QDragLeaveEvent)  {
	q.Drv(395000,395137,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::dragMoveEvent(QDragMoveEvent*)
func (q *QWidget) DragMoveEvent(value *QDragMoveEvent)  {
	q.Drv(395000,395138,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::dropEvent(QDropEvent*)
func (q *QWidget) DropEvent(value *QDropEvent)  {
	q.Drv(395000,395139,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::ensurePolished()
func (q *QWidget) EnsurePolished()  {
	q.Drv(395000,395140,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::enterEvent(QEvent*)
func (q *QWidget) EnterEvent(value *QEvent)  {
	q.Drv(395000,395141,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::event(QEvent*)
func (q *QWidget) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(395000,395142,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::focusInEvent(QFocusEvent*)
func (q *QWidget) FocusInEvent(value *QFocusEvent)  {
	q.Drv(395000,395143,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::focusNextChild()
func (q *QWidget) FocusNextChild() bool {
	var __rv bool
	q.Drv(395000,395144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::focusNextPrevChild(bool)
func (q *QWidget) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(395000,395145,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::focusOutEvent(QFocusEvent*)
func (q *QWidget) FocusOutEvent(value *QFocusEvent)  {
	q.Drv(395000,395146,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::focusPolicy()
func (q *QWidget) FocusPolicy() Qt_FocusPolicy {
	var __rv Qt_FocusPolicy
	q.Drv(395000,395147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::focusPreviousChild()
func (q *QWidget) FocusPreviousChild() bool {
	var __rv bool
	q.Drv(395000,395148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::focusProxy()
func (q *QWidget) FocusProxy() *QWidget {
	var __rv uintptr
	q.Drv(395000,395149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::focusWidget()
func (q *QWidget) FocusWidget() *QWidget {
	var __rv uintptr
	q.Drv(395000,395150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::font()
func (q *QWidget) Font() *QFont {
	var __rv uintptr
	q.Drv(395000,395151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QWidget::fontInfo()
func (q *QWidget) FontInfo() *QFontInfo {
	var __rv uintptr
	q.Drv(395000,395152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontInfo{}
	_rp.SetDriver(__rv,39000,true)
	return _rp
}	
//QWidget::fontMetrics()
func (q *QWidget) FontMetrics() *QFontMetrics {
	var __rv uintptr
	q.Drv(395000,395153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontMetrics{}
	_rp.SetDriver(__rv,40000,true)
	return _rp
}	
//QWidget::foregroundRole()
func (q *QWidget) ForegroundRole() QPalette_ColorRole {
	var __rv QPalette_ColorRole
	q.Drv(395000,395154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::frameGeometry()
func (q *QWidget) FrameGeometry() *QRect {
	var __rv uintptr
	q.Drv(395000,395155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::frameSize()
func (q *QWidget) FrameSize() *QSize {
	var __rv uintptr
	q.Drv(395000,395156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::geometry()
func (q *QWidget) Geometry() *QRect {
	var __rv uintptr
	q.Drv(395000,395157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::getContentsMargins(int*,int*,int*,int*)
func (q *QWidget) GetContentsMargins(left *int,top *int,right *int,bottom *int)  {
	q.Drv(395000,395158,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabGesture(Qt::GestureType)
func (q *QWidget) GrabGesture(_type Qt_GestureType)  {
	q.Drv(395000,395159,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabGesture(Qt::GestureType,QFlags<Qt::GestureFlag>)
func (q *QWidget) GrabGestureWithTypeFlags(_type Qt_GestureType,flags Qt_GestureFlag)  {
	q.Drv(395000,395160,unsafe.Pointer(&_type),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabKeyboard()
func (q *QWidget) GrabKeyboard()  {
	q.Drv(395000,395161,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabMouse()
func (q *QWidget) GrabMouse()  {
	q.Drv(395000,395162,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabMouse(QCursor const&)
func (q *QWidget) GrabMouseWithCursor(value *QCursor)  {
	q.Drv(395000,395163,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::grabShortcut(QKeySequence const&)
func (q *QWidget) GrabShortcut(key *QKeySequence) int {
	var __rv int
	q.Drv(395000,395164,Native(key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::grabShortcut(QKeySequence const&,Qt::ShortcutContext)
func (q *QWidget) GrabShortcutWithKeyContext(key *QKeySequence,context Qt_ShortcutContext) int {
	var __rv int
	q.Drv(395000,395165,Native(key),unsafe.Pointer(&context),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::graphicsEffect()
func (q *QWidget) GraphicsEffect() *QGraphicsEffect {
	var __rv uintptr
	q.Drv(395000,395166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsEffect{}
	_rp.SetDriver(__rv,253000,false)
	return _rp
}	
//QWidget::graphicsProxyWidget()
func (q *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	var __rv uintptr
	q.Drv(395000,395167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsProxyWidget{}
	_rp.SetDriver(__rv,268000,false)
	return _rp
}	
//QWidget::hasFocus()
func (q *QWidget) HasFocus() bool {
	var __rv bool
	q.Drv(395000,395168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::hasMouseTracking()
func (q *QWidget) HasMouseTracking() bool {
	var __rv bool
	q.Drv(395000,395169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::height()
func (q *QWidget) Height() int {
	var __rv int
	q.Drv(395000,395170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::heightForWidth(int)
func (q *QWidget) HeightForWidth(value int) int {
	var __rv int
	q.Drv(395000,395171,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::heightMM()
func (q *QWidget) HeightMM() int {
	var __rv int
	q.Drv(395000,395172,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::hide()
func (q *QWidget) Hide()  {
	q.Drv(395000,395173,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::hideEvent(QHideEvent*)
func (q *QWidget) HideEvent(value *QHideEvent)  {
	q.Drv(395000,395174,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::inputMethodEvent(QInputMethodEvent*)
func (q *QWidget) InputMethodEvent(value *QInputMethodEvent)  {
	q.Drv(395000,395175,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::inputMethodHints()
func (q *QWidget) InputMethodHints() Qt_InputMethodHint {
	var __rv Qt_InputMethodHint
	q.Drv(395000,395176,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::inputMethodQuery(Qt::InputMethodQuery)
func (q *QWidget) InputMethodQuery(value Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(395000,395177,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QWidget::insertAction(QAction*,QAction*)
func (q *QWidget) InsertAction(before *QAction,action *QAction)  {
	q.Drv(395000,395178,Native(before),Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::insertActions(QAction*,QList<QAction*>)
func (q *QWidget) InsertActions(before *QAction,actions []*QAction)  {
	q.Drv(395000,395179,Native(before),unsafe.Pointer(&actions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::isActiveWindow()
func (q *QWidget) IsActiveWindow() bool {
	var __rv bool
	q.Drv(395000,395180,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isAncestorOf(QWidget const*)
func (q *QWidget) IsAncestorOf(child QWidgetInterface) bool {
	var __rv bool
	q.Drv(395000,395181,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isEnabled()
func (q *QWidget) IsEnabled() bool {
	var __rv bool
	q.Drv(395000,395182,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isEnabledTo(QWidget*)
func (q *QWidget) IsEnabledTo(value QWidgetInterface) bool {
	var __rv bool
	q.Drv(395000,395183,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isEnabledToTLW()
func (q *QWidget) IsEnabledToTLW() bool {
	var __rv bool
	q.Drv(395000,395184,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isFullScreen()
func (q *QWidget) IsFullScreen() bool {
	var __rv bool
	q.Drv(395000,395185,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isHidden()
func (q *QWidget) IsHidden() bool {
	var __rv bool
	q.Drv(395000,395186,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isLeftToRight()
func (q *QWidget) IsLeftToRight() bool {
	var __rv bool
	q.Drv(395000,395187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isMaximized()
func (q *QWidget) IsMaximized() bool {
	var __rv bool
	q.Drv(395000,395188,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isMinimized()
func (q *QWidget) IsMinimized() bool {
	var __rv bool
	q.Drv(395000,395189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isModal()
func (q *QWidget) IsModal() bool {
	var __rv bool
	q.Drv(395000,395190,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isRightToLeft()
func (q *QWidget) IsRightToLeft() bool {
	var __rv bool
	q.Drv(395000,395191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isTopLevel()
func (q *QWidget) IsTopLevel() bool {
	var __rv bool
	q.Drv(395000,395192,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isVisible()
func (q *QWidget) IsVisible() bool {
	var __rv bool
	q.Drv(395000,395193,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isVisibleTo(QWidget*)
func (q *QWidget) IsVisibleTo(value QWidgetInterface) bool {
	var __rv bool
	q.Drv(395000,395194,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isWindow()
func (q *QWidget) IsWindow() bool {
	var __rv bool
	q.Drv(395000,395195,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::isWindowModified()
func (q *QWidget) IsWindowModified() bool {
	var __rv bool
	q.Drv(395000,395196,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::keyPressEvent(QKeyEvent*)
func (q *QWidget) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(395000,395197,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::keyReleaseEvent(QKeyEvent*)
func (q *QWidget) KeyReleaseEvent(value *QKeyEvent)  {
	q.Drv(395000,395198,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::keyboardGrabber()	
func QWidgetKeyboardGrabber() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,395000,395199,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::keyboardGrabber()
func (q *QWidget) KeyboardGrabber() *QWidget {
	var __rv uintptr
	q.Drv(395000,395199,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::layout()
func (q *QWidget) Layout() *QLayout {
	var __rv uintptr
	q.Drv(395000,395200,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayout{}
	_rp.SetDriver(__rv,300000,false)
	return _rp
}	
//QWidget::layoutDirection()
func (q *QWidget) LayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(395000,395201,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::leaveEvent(QEvent*)
func (q *QWidget) LeaveEvent(value *QEvent)  {
	q.Drv(395000,395202,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::locale()
func (q *QWidget) Locale() *QLocale {
	var __rv uintptr
	q.Drv(395000,395203,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QPaintDevice::logicalDpiX()
func (q *QWidget) LogicalDpiX() int {
	var __rv int
	q.Drv(395000,395204,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::logicalDpiY()
func (q *QWidget) LogicalDpiY() int {
	var __rv int
	q.Drv(395000,395205,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::lower()
func (q *QWidget) Lower()  {
	q.Drv(395000,395206,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::mapFrom(QWidget*,QPoint const&)
func (q *QWidget) MapFrom(value2 QWidgetInterface,value3 *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395207,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mapFromGlobal(QPoint const&)
func (q *QWidget) MapFromGlobal(value *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395208,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mapFromParent(QPoint const&)
func (q *QWidget) MapFromParent(value *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395209,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mapTo(QWidget*,QPoint const&)
func (q *QWidget) MapTo(value2 QWidgetInterface,value3 *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395210,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mapToGlobal(QPoint const&)
func (q *QWidget) MapToGlobal(value *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395211,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mapToParent(QPoint const&)
func (q *QWidget) MapToParent(value *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(395000,395212,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::mask()
func (q *QWidget) Mask() *QRegion {
	var __rv uintptr
	q.Drv(395000,395213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QWidget::maximumHeight()
func (q *QWidget) MaximumHeight() int {
	var __rv int
	q.Drv(395000,395214,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::maximumSize()
func (q *QWidget) MaximumSize() *QSize {
	var __rv uintptr
	q.Drv(395000,395215,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::maximumWidth()
func (q *QWidget) MaximumWidth() int {
	var __rv int
	q.Drv(395000,395216,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::metric(QPaintDevice::PaintDeviceMetric)
func (q *QWidget) Metric(value QPaintDevice_PaintDeviceMetric) int {
	var __rv int
	q.Drv(395000,395217,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::minimumHeight()
func (q *QWidget) MinimumHeight() int {
	var __rv int
	q.Drv(395000,395218,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::minimumSize()
func (q *QWidget) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(395000,395219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::minimumSizeHint()
func (q *QWidget) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(395000,395220,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::minimumWidth()
func (q *QWidget) MinimumWidth() int {
	var __rv int
	q.Drv(395000,395221,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::mouseDoubleClickEvent(QMouseEvent*)
func (q *QWidget) MouseDoubleClickEvent(value *QMouseEvent)  {
	q.Drv(395000,395222,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::mouseGrabber()	
func QWidgetMouseGrabber() *QWidget {
	var __rv uintptr
	DirectQtDrv(nil,395000,395223,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::mouseGrabber()
func (q *QWidget) MouseGrabber() *QWidget {
	var __rv uintptr
	q.Drv(395000,395223,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::mouseMoveEvent(QMouseEvent*)
func (q *QWidget) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(395000,395224,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::mousePressEvent(QMouseEvent*)
func (q *QWidget) MousePressEvent(value *QMouseEvent)  {
	q.Drv(395000,395225,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::mouseReleaseEvent(QMouseEvent*)
func (q *QWidget) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(395000,395226,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::move(QPoint const&)
func (q *QWidget) Move(value *QPoint)  {
	q.Drv(395000,395227,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::move(int,int)
func (q *QWidget) MoveWithXY(x int,y int)  {
	q.Drv(395000,395228,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::moveEvent(QMoveEvent*)
func (q *QWidget) MoveEvent(value *QMoveEvent)  {
	q.Drv(395000,395229,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::nativeParentWidget()
func (q *QWidget) NativeParentWidget() *QWidget {
	var __rv uintptr
	q.Drv(395000,395230,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::nextInFocusChain()
func (q *QWidget) NextInFocusChain() *QWidget {
	var __rv uintptr
	q.Drv(395000,395231,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::normalGeometry()
func (q *QWidget) NormalGeometry() *QRect {
	var __rv uintptr
	q.Drv(395000,395232,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::overrideWindowFlags(QFlags<Qt::WindowType>)
func (q *QWidget) OverrideWindowFlags(_type Qt_WindowType)  {
	q.Drv(395000,395233,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::overrideWindowState(QFlags<Qt::WindowState>)
func (q *QWidget) OverrideWindowState(state Qt_WindowState)  {
	q.Drv(395000,395234,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::paintEngine()
func (q *QWidget) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(395000,395235,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QWidget::paintEvent(QPaintEvent*)
func (q *QWidget) PaintEvent(value *QPaintEvent)  {
	q.Drv(395000,395236,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintDevice::paintingActive()
func (q *QWidget) PaintingActive() bool {
	var __rv bool
	q.Drv(395000,395237,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::palette()
func (q *QWidget) Palette() *QPalette {
	var __rv uintptr
	q.Drv(395000,395238,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QWidget::parentWidget()
func (q *QWidget) ParentWidget() *QWidget {
	var __rv uintptr
	q.Drv(395000,395239,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QPaintDevice::physicalDpiX()
func (q *QWidget) PhysicalDpiX() int {
	var __rv int
	q.Drv(395000,395240,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::physicalDpiY()
func (q *QWidget) PhysicalDpiY() int {
	var __rv int
	q.Drv(395000,395241,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::pos()
func (q *QWidget) Pos() *QPoint {
	var __rv uintptr
	q.Drv(395000,395242,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QWidget::previousInFocusChain()
func (q *QWidget) PreviousInFocusChain() *QWidget {
	var __rv uintptr
	q.Drv(395000,395243,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::raise()
func (q *QWidget) Raise()  {
	q.Drv(395000,395244,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::rect()
func (q *QWidget) Rect() *QRect {
	var __rv uintptr
	q.Drv(395000,395245,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QWidget::releaseKeyboard()
func (q *QWidget) ReleaseKeyboard()  {
	q.Drv(395000,395246,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::releaseMouse()
func (q *QWidget) ReleaseMouse()  {
	q.Drv(395000,395247,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::releaseShortcut(int)
func (q *QWidget) ReleaseShortcut(id int)  {
	q.Drv(395000,395248,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::removeAction(QAction*)
func (q *QWidget) RemoveAction(action *QAction)  {
	q.Drv(395000,395249,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::render(QPaintDevice*)
func (q *QWidget) Render(target QPaintDeviceInterface)  {
	q.Drv(395000,395250,unsafe.Pointer(new_pd_head(target)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::render(QPainter*)
func (q *QWidget) RenderWithPainter(painter *QPainter)  {
	q.Drv(395000,395251,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::render(QPaintDevice*,QPoint const&,QRegion const&,QFlags<QWidget::RenderFlag>)
func (q *QWidget) RenderWithPaintDeviceTargetoffsetSourceregionRenderflags(target QPaintDeviceInterface,targetOffset *QPoint,sourceRegion *QRegion,renderFlags QWidget_RenderFlag)  {
	q.Drv(395000,395252,unsafe.Pointer(new_pd_head(target)),Native(targetOffset),Native(sourceRegion),unsafe.Pointer(&renderFlags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::render(QPainter*,QPoint const&,QRegion const&,QFlags<QWidget::RenderFlag>)
func (q *QWidget) RenderWithPainterTargetoffsetSourceregionRenderflags(painter *QPainter,targetOffset *QPoint,sourceRegion *QRegion,renderFlags QWidget_RenderFlag)  {
	q.Drv(395000,395253,Native(painter),Native(targetOffset),Native(sourceRegion),unsafe.Pointer(&renderFlags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::repaint()
func (q *QWidget) Repaint()  {
	q.Drv(395000,395254,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::repaint(QRect const&)
func (q *QWidget) RepaintWithRect(value *QRect)  {
	q.Drv(395000,395255,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::repaint(QRegion const&)
func (q *QWidget) RepaintWithRegion(value *QRegion)  {
	q.Drv(395000,395256,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::repaint(int,int,int,int)
func (q *QWidget) RepaintWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(395000,395257,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::resize(QSize const&)
func (q *QWidget) Resize(value *QSize)  {
	q.Drv(395000,395258,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::resize(int,int)
func (q *QWidget) ResizeWithWidthHeight(w int,h int)  {
	q.Drv(395000,395259,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::resizeEvent(QResizeEvent*)
func (q *QWidget) ResizeEvent(value *QResizeEvent)  {
	q.Drv(395000,395260,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::restoreGeometry(QByteArray const&)
func (q *QWidget) RestoreGeometry(geometry []byte) bool {
	var __rv bool
	q.Drv(395000,395261,unsafe.Pointer(&geometry),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::saveGeometry()
func (q *QWidget) SaveGeometry() []byte {
	var __rv []byte
	q.Drv(395000,395262,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::scroll(int,int)
func (q *QWidget) ScrollWithDxDy(dx int,dy int)  {
	q.Drv(395000,395263,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::scroll(int,int,QRect const&)
func (q *QWidget) ScrollWithDxDyRect(dx int,dy int,value2 *QRect)  {
	q.Drv(395000,395264,unsafe.Pointer(&dx),unsafe.Pointer(&dy),Native(value2),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAcceptDrops(bool)
func (q *QWidget) SetAcceptDrops(on bool)  {
	q.Drv(395000,395265,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAccessibleDescription(QString const&)
func (q *QWidget) SetAccessibleDescription(description string)  {
	q.Drv(395000,395266,unsafe.Pointer(&description),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAccessibleName(QString const&)
func (q *QWidget) SetAccessibleName(name string)  {
	q.Drv(395000,395267,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAttribute(Qt::WidgetAttribute)
func (q *QWidget) SetAttribute(value Qt_WidgetAttribute)  {
	q.Drv(395000,395268,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAttribute(Qt::WidgetAttribute,bool)
func (q *QWidget) SetAttributeWithWidgetattributeOn(value2 Qt_WidgetAttribute,on bool)  {
	q.Drv(395000,395269,unsafe.Pointer(&value2),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setAutoFillBackground(bool)
func (q *QWidget) SetAutoFillBackground(enabled bool)  {
	q.Drv(395000,395270,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setBackgroundRole(QPalette::ColorRole)
func (q *QWidget) SetBackgroundRole(value QPalette_ColorRole)  {
	q.Drv(395000,395271,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setBaseSize(QSize const&)
func (q *QWidget) SetBaseSize(value *QSize)  {
	q.Drv(395000,395272,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setBaseSize(int,int)
func (q *QWidget) SetBaseSizeWithBasewBaseh(basew int,baseh int)  {
	q.Drv(395000,395273,unsafe.Pointer(&basew),unsafe.Pointer(&baseh),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setContentsMargins(QMargins const&)
func (q *QWidget) SetContentsMargins(margins *QMargins)  {
	q.Drv(395000,395274,Native(margins),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setContentsMargins(int,int,int,int)
func (q *QWidget) SetContentsMarginsWithLeftTopRightBottom(left int,top int,right int,bottom int)  {
	q.Drv(395000,395275,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setContextMenuPolicy(Qt::ContextMenuPolicy)
func (q *QWidget) SetContextMenuPolicy(policy Qt_ContextMenuPolicy)  {
	q.Drv(395000,395276,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setCursor(QCursor const&)
func (q *QWidget) SetCursor(value *QCursor)  {
	q.Drv(395000,395277,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setDisabled(bool)
func (q *QWidget) SetDisabled(value bool)  {
	q.Drv(395000,395278,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setEnabled(bool)
func (q *QWidget) SetEnabled(value bool)  {
	q.Drv(395000,395279,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFixedHeight(int)
func (q *QWidget) SetFixedHeight(h int)  {
	q.Drv(395000,395280,unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFixedSize(QSize const&)
func (q *QWidget) SetFixedSize(value *QSize)  {
	q.Drv(395000,395281,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFixedSize(int,int)
func (q *QWidget) SetFixedSizeWithWidthHeight(w int,h int)  {
	q.Drv(395000,395282,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFixedWidth(int)
func (q *QWidget) SetFixedWidth(w int)  {
	q.Drv(395000,395283,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFocus()
func (q *QWidget) SetFocus()  {
	q.Drv(395000,395284,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFocus(Qt::FocusReason)
func (q *QWidget) SetFocusWithReason(reason Qt_FocusReason)  {
	q.Drv(395000,395285,unsafe.Pointer(&reason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFocusPolicy(Qt::FocusPolicy)
func (q *QWidget) SetFocusPolicy(policy Qt_FocusPolicy)  {
	q.Drv(395000,395286,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFocusProxy(QWidget*)
func (q *QWidget) SetFocusProxy(value QWidgetInterface)  {
	q.Drv(395000,395287,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setFont(QFont const&)
func (q *QWidget) SetFont(value *QFont)  {
	q.Drv(395000,395288,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setForegroundRole(QPalette::ColorRole)
func (q *QWidget) SetForegroundRole(value QPalette_ColorRole)  {
	q.Drv(395000,395289,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setGeometry(QRect const&)
func (q *QWidget) SetGeometry(value *QRect)  {
	q.Drv(395000,395290,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setGeometry(int,int,int,int)
func (q *QWidget) SetGeometryWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(395000,395291,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setGraphicsEffect(QGraphicsEffect*)
func (q *QWidget) SetGraphicsEffect(effect *QGraphicsEffect)  {
	q.Drv(395000,395292,Native(effect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setHidden(bool)
func (q *QWidget) SetHidden(hidden bool)  {
	q.Drv(395000,395293,unsafe.Pointer(&hidden),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setInputMethodHints(QFlags<Qt::InputMethodHint>)
func (q *QWidget) SetInputMethodHints(hints Qt_InputMethodHint)  {
	q.Drv(395000,395294,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setLayout(QLayout*)
func (q *QWidget) SetLayout(value QLayoutInterface)  {
	q.Drv(395000,395295,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setLayoutDirection(Qt::LayoutDirection)
func (q *QWidget) SetLayoutDirection(direction Qt_LayoutDirection)  {
	q.Drv(395000,395296,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setLocale(QLocale const&)
func (q *QWidget) SetLocale(locale *QLocale)  {
	q.Drv(395000,395297,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMask(QBitmap const&)
func (q *QWidget) SetMask(value *QBitmap)  {
	q.Drv(395000,395298,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMask(QRegion const&)
func (q *QWidget) SetMaskWithRegion(value *QRegion)  {
	q.Drv(395000,395299,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMaximumHeight(int)
func (q *QWidget) SetMaximumHeight(maxh int)  {
	q.Drv(395000,395300,unsafe.Pointer(&maxh),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMaximumSize(QSize const&)
func (q *QWidget) SetMaximumSize(value *QSize)  {
	q.Drv(395000,395301,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMaximumSize(int,int)
func (q *QWidget) SetMaximumSizeWithMaxwMaxh(maxw int,maxh int)  {
	q.Drv(395000,395302,unsafe.Pointer(&maxw),unsafe.Pointer(&maxh),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMaximumWidth(int)
func (q *QWidget) SetMaximumWidth(maxw int)  {
	q.Drv(395000,395303,unsafe.Pointer(&maxw),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMinimumHeight(int)
func (q *QWidget) SetMinimumHeight(minh int)  {
	q.Drv(395000,395304,unsafe.Pointer(&minh),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMinimumSize(QSize const&)
func (q *QWidget) SetMinimumSize(value *QSize)  {
	q.Drv(395000,395305,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMinimumSize(int,int)
func (q *QWidget) SetMinimumSizeWithMinwMinh(minw int,minh int)  {
	q.Drv(395000,395306,unsafe.Pointer(&minw),unsafe.Pointer(&minh),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMinimumWidth(int)
func (q *QWidget) SetMinimumWidth(minw int)  {
	q.Drv(395000,395307,unsafe.Pointer(&minw),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setMouseTracking(bool)
func (q *QWidget) SetMouseTracking(enable bool)  {
	q.Drv(395000,395308,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setPalette(QPalette const&)
func (q *QWidget) SetPalette(value *QPalette)  {
	q.Drv(395000,395309,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setParent(QWidget*)
func (q *QWidget) SetParentWidget(parent QWidgetInterface)  {
	q.Drv(395000,395310,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setParent(QWidget*,QFlags<Qt::WindowType>)
func (q *QWidget) SetParentWidgetWithParentFlags(parent QWidgetInterface,f Qt_WindowType)  {
	q.Drv(395000,395311,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setShortcutAutoRepeat(int)
func (q *QWidget) SetShortcutAutoRepeat(id int)  {
	q.Drv(395000,395312,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setShortcutAutoRepeat(int,bool)
func (q *QWidget) SetShortcutAutoRepeatWithIdEnable(id int,enable bool)  {
	q.Drv(395000,395313,unsafe.Pointer(&id),unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setShortcutEnabled(int)
func (q *QWidget) SetShortcutEnabled(id int)  {
	q.Drv(395000,395314,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setShortcutEnabled(int,bool)
func (q *QWidget) SetShortcutEnabledWithIdEnable(id int,enable bool)  {
	q.Drv(395000,395315,unsafe.Pointer(&id),unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setSizeIncrement(QSize const&)
func (q *QWidget) SetSizeIncrement(value *QSize)  {
	q.Drv(395000,395316,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setSizeIncrement(int,int)
func (q *QWidget) SetSizeIncrementWithWidthHeight(w int,h int)  {
	q.Drv(395000,395317,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setSizePolicy(QSizePolicy)
func (q *QWidget) SetSizePolicy(value *QSizePolicy)  {
	q.Drv(395000,395318,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setSizePolicy(QSizePolicy::Policy,QSizePolicy::Policy)
func (q *QWidget) SetSizePolicyWithHorizontalVertical(horizontal QSizePolicy_Policy,vertical QSizePolicy_Policy)  {
	q.Drv(395000,395319,unsafe.Pointer(&horizontal),unsafe.Pointer(&vertical),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setStatusTip(QString const&)
func (q *QWidget) SetStatusTip(value string)  {
	q.Drv(395000,395320,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setStyle(QStyle*)
func (q *QWidget) SetStyle(value *QStyle)  {
	q.Drv(395000,395321,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setStyleSheet(QString const&)
func (q *QWidget) SetStyleSheet(styleSheet string)  {
	q.Drv(395000,395322,unsafe.Pointer(&styleSheet),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setTabOrder(QWidget*,QWidget*)	
func QWidgetSetTabOrder(value2 QWidgetInterface,value3 QWidgetInterface)  {
	DirectQtDrv(nil,395000,395323,Native(value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setTabOrder(QWidget*,QWidget*)
func (q *QWidget) SetTabOrder(value2 QWidgetInterface,value3 QWidgetInterface)  {
	q.Drv(395000,395323,Native(value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setToolTip(QString const&)
func (q *QWidget) SetToolTip(value string)  {
	q.Drv(395000,395324,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setUpdatesEnabled(bool)
func (q *QWidget) SetUpdatesEnabled(enable bool)  {
	q.Drv(395000,395325,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setVisible(bool)
func (q *QWidget) SetVisible(visible bool)  {
	q.Drv(395000,395326,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWhatsThis(QString const&)
func (q *QWidget) SetWhatsThis(value string)  {
	q.Drv(395000,395327,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowFilePath(QString const&)
func (q *QWidget) SetWindowFilePath(filePath string)  {
	q.Drv(395000,395328,unsafe.Pointer(&filePath),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowFlags(QFlags<Qt::WindowType>)
func (q *QWidget) SetWindowFlags(_type Qt_WindowType)  {
	q.Drv(395000,395329,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowIcon(QIcon const&)
func (q *QWidget) SetWindowIcon(icon *QIcon)  {
	q.Drv(395000,395330,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowIconText(QString const&)
func (q *QWidget) SetWindowIconText(value string)  {
	q.Drv(395000,395331,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowModality(Qt::WindowModality)
func (q *QWidget) SetWindowModality(windowModality Qt_WindowModality)  {
	q.Drv(395000,395332,unsafe.Pointer(&windowModality),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowModified(bool)
func (q *QWidget) SetWindowModified(value bool)  {
	q.Drv(395000,395333,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowOpacity(double)
func (q *QWidget) SetWindowOpacity(level float64)  {
	q.Drv(395000,395334,unsafe.Pointer(&level),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowRole(QString const&)
func (q *QWidget) SetWindowRole(value string)  {
	q.Drv(395000,395335,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowState(QFlags<Qt::WindowState>)
func (q *QWidget) SetWindowState(state Qt_WindowState)  {
	q.Drv(395000,395336,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::setWindowTitle(QString const&)
func (q *QWidget) SetWindowTitle(value string)  {
	q.Drv(395000,395337,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::show()
func (q *QWidget) Show()  {
	q.Drv(395000,395338,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::showEvent(QShowEvent*)
func (q *QWidget) ShowEvent(value *QShowEvent)  {
	q.Drv(395000,395339,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::showFullScreen()
func (q *QWidget) ShowFullScreen()  {
	q.Drv(395000,395340,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::showMaximized()
func (q *QWidget) ShowMaximized()  {
	q.Drv(395000,395341,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::showMinimized()
func (q *QWidget) ShowMinimized()  {
	q.Drv(395000,395342,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::showNormal()
func (q *QWidget) ShowNormal()  {
	q.Drv(395000,395343,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::size()
func (q *QWidget) Size() *QSize {
	var __rv uintptr
	q.Drv(395000,395344,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::sizeHint()
func (q *QWidget) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(395000,395345,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::sizeIncrement()
func (q *QWidget) SizeIncrement() *QSize {
	var __rv uintptr
	q.Drv(395000,395346,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWidget::sizePolicy()
func (q *QWidget) SizePolicy() *QSizePolicy {
	var __rv uintptr
	q.Drv(395000,395347,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizePolicy{}
	_rp.SetDriver(__rv,121000,true)
	return _rp
}	
//QWidget::stackUnder(QWidget*)
func (q *QWidget) StackUnder(value QWidgetInterface)  {
	q.Drv(395000,395348,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::statusTip()
func (q *QWidget) StatusTip() string {
	var __rv string
	q.Drv(395000,395349,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::style()
func (q *QWidget) Style() *QStyle {
	var __rv uintptr
	q.Drv(395000,395350,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QWidget::styleSheet()
func (q *QWidget) StyleSheet() string {
	var __rv string
	q.Drv(395000,395351,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::tabletEvent(QTabletEvent*)
func (q *QWidget) TabletEvent(value *QTabletEvent)  {
	q.Drv(395000,395352,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::testAttribute(Qt::WidgetAttribute)
func (q *QWidget) TestAttribute(value Qt_WidgetAttribute) bool {
	var __rv bool
	q.Drv(395000,395353,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::toolTip()
func (q *QWidget) ToolTip() string {
	var __rv string
	q.Drv(395000,395354,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::topLevelWidget()
func (q *QWidget) TopLevelWidget() *QWidget {
	var __rv uintptr
	q.Drv(395000,395355,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::underMouse()
func (q *QWidget) UnderMouse() bool {
	var __rv bool
	q.Drv(395000,395356,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::ungrabGesture(Qt::GestureType)
func (q *QWidget) UngrabGesture(_type Qt_GestureType)  {
	q.Drv(395000,395357,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::unsetCursor()
func (q *QWidget) UnsetCursor()  {
	q.Drv(395000,395358,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::unsetLayoutDirection()
func (q *QWidget) UnsetLayoutDirection()  {
	q.Drv(395000,395359,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::unsetLocale()
func (q *QWidget) UnsetLocale()  {
	q.Drv(395000,395360,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::update()
func (q *QWidget) Update()  {
	q.Drv(395000,395361,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::update(QRect const&)
func (q *QWidget) UpdateWithRect(value *QRect)  {
	q.Drv(395000,395362,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::update(QRegion const&)
func (q *QWidget) UpdateWithRegion(value *QRegion)  {
	q.Drv(395000,395363,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::update(int,int,int,int)
func (q *QWidget) UpdateWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(395000,395364,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::updateGeometry()
func (q *QWidget) UpdateGeometry()  {
	q.Drv(395000,395365,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::updateMicroFocus()
func (q *QWidget) UpdateMicroFocus()  {
	q.Drv(395000,395366,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::updatesEnabled()
func (q *QWidget) UpdatesEnabled() bool {
	var __rv bool
	q.Drv(395000,395367,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::visibleRegion()
func (q *QWidget) VisibleRegion() *QRegion {
	var __rv uintptr
	q.Drv(395000,395368,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QWidget::whatsThis()
func (q *QWidget) WhatsThis() string {
	var __rv string
	q.Drv(395000,395369,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::wheelEvent(QWheelEvent*)
func (q *QWidget) WheelEvent(value *QWheelEvent)  {
	q.Drv(395000,395370,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidget::width()
func (q *QWidget) Width() int {
	var __rv int
	q.Drv(395000,395371,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::widthMM()
func (q *QWidget) WidthMM() int {
	var __rv int
	q.Drv(395000,395372,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::window()
func (q *QWidget) Window() *QWidget {
	var __rv uintptr
	q.Drv(395000,395373,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidget::windowFilePath()
func (q *QWidget) WindowFilePath() string {
	var __rv string
	q.Drv(395000,395374,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowFlags()
func (q *QWidget) WindowFlags() Qt_WindowType {
	var __rv Qt_WindowType
	q.Drv(395000,395375,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowIcon()
func (q *QWidget) WindowIcon() *QIcon {
	var __rv uintptr
	q.Drv(395000,395376,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QWidget::windowIconText()
func (q *QWidget) WindowIconText() string {
	var __rv string
	q.Drv(395000,395377,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowModality()
func (q *QWidget) WindowModality() Qt_WindowModality {
	var __rv Qt_WindowModality
	q.Drv(395000,395378,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowOpacity()
func (q *QWidget) WindowOpacity() float64 {
	var __rv float64
	q.Drv(395000,395379,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowRole()
func (q *QWidget) WindowRole() string {
	var __rv string
	q.Drv(395000,395380,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowState()
func (q *QWidget) WindowState() Qt_WindowState {
	var __rv Qt_WindowState
	q.Drv(395000,395381,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowTitle()
func (q *QWidget) WindowTitle() string {
	var __rv string
	q.Drv(395000,395382,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::windowType()
func (q *QWidget) WindowType() Qt_WindowType {
	var __rv Qt_WindowType
	q.Drv(395000,395383,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::x()
func (q *QWidget) X() int {
	var __rv int
	q.Drv(395000,395384,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidget::y()
func (q *QWidget) Y() int {
	var __rv int
	q.Drv(395000,395385,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
