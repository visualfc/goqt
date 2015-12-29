// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QToolButton_ToolButtonPopupMode - QToolButton::ToolButtonPopupMode
type QToolButton_ToolButtonPopupMode uint32
const (
	QToolButton_DelayedPopup QToolButton_ToolButtonPopupMode = 0
	QToolButton_MenuButtonPopup QToolButton_ToolButtonPopupMode = 1
	QToolButton_InstantPopup QToolButton_ToolButtonPopupMode = 2
)
//struct QToolButton : QToolButton
type QToolButton struct {
	QAbstractButton
}
func (q *QToolButton) OnTriggered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(383000,383102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QToolButton::QToolButton()	
func NewToolButton() *QToolButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),383000,383103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolButton{}
	_p.SetDriver(__rv,383000,false)
	return _p
} 
//QToolButton::QToolButton(QWidget*)	
func NewToolButtonWithParent(parent QWidgetInterface) *QToolButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),383000,383104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolButton{}
	_p.SetDriver(__rv,383000,false)
	return _p
} 
//QToolButton::actionEvent(QActionEvent*)
func (q *QToolButton) ActionEvent(value *QActionEvent)  {
	q.Drv(383000,383105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::arrowType()
func (q *QToolButton) ArrowType() Qt_ArrowType {
	var __rv Qt_ArrowType
	q.Drv(383000,383106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolButton::autoRaise()
func (q *QToolButton) AutoRaise() bool {
	var __rv bool
	q.Drv(383000,383107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolButton::changeEvent(QEvent*)
func (q *QToolButton) ChangeEvent(value *QEvent)  {
	q.Drv(383000,383108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::defaultAction()
func (q *QToolButton) DefaultAction() *QAction {
	var __rv uintptr
	q.Drv(383000,383109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolButton::enterEvent(QEvent*)
func (q *QToolButton) EnterEvent(value *QEvent)  {
	q.Drv(383000,383110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::event(QEvent*)
func (q *QToolButton) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(383000,383111,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolButton::hitButton(QPoint const&)
func (q *QToolButton) HitButton(pos *QPoint) bool {
	var __rv bool
	q.Drv(383000,383112,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolButton::leaveEvent(QEvent*)
func (q *QToolButton) LeaveEvent(value *QEvent)  {
	q.Drv(383000,383113,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::menu()
func (q *QToolButton) Menu() *QMenu {
	var __rv uintptr
	q.Drv(383000,383114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QToolButton::minimumSizeHint()
func (q *QToolButton) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(383000,383115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QToolButton::mousePressEvent(QMouseEvent*)
func (q *QToolButton) MousePressEvent(value *QMouseEvent)  {
	q.Drv(383000,383116,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::mouseReleaseEvent(QMouseEvent*)
func (q *QToolButton) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(383000,383117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::nextCheckState()
func (q *QToolButton) NextCheckState()  {
	q.Drv(383000,383118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::paintEvent(QPaintEvent*)
func (q *QToolButton) PaintEvent(value *QPaintEvent)  {
	q.Drv(383000,383119,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::popupMode()
func (q *QToolButton) PopupMode() QToolButton_ToolButtonPopupMode {
	var __rv QToolButton_ToolButtonPopupMode
	q.Drv(383000,383120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolButton::setArrowType(Qt::ArrowType)
func (q *QToolButton) SetArrowType(_type Qt_ArrowType)  {
	q.Drv(383000,383121,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::setAutoRaise(bool)
func (q *QToolButton) SetAutoRaise(enable bool)  {
	q.Drv(383000,383122,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::setDefaultAction(QAction*)
func (q *QToolButton) SetDefaultAction(value *QAction)  {
	q.Drv(383000,383123,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::setMenu(QMenu*)
func (q *QToolButton) SetMenu(menu *QMenu)  {
	q.Drv(383000,383124,Native(menu),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::setPopupMode(QToolButton::ToolButtonPopupMode)
func (q *QToolButton) SetPopupMode(mode QToolButton_ToolButtonPopupMode)  {
	q.Drv(383000,383125,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::setToolButtonStyle(Qt::ToolButtonStyle)
func (q *QToolButton) SetToolButtonStyle(style Qt_ToolButtonStyle)  {
	q.Drv(383000,383126,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::showMenu()
func (q *QToolButton) ShowMenu()  {
	q.Drv(383000,383127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::sizeHint()
func (q *QToolButton) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(383000,383128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QToolButton::timerEvent(QTimerEvent*)
func (q *QToolButton) TimerEvent(value *QTimerEvent)  {
	q.Drv(383000,383129,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolButton::toolButtonStyle()
func (q *QToolButton) ToolButtonStyle() Qt_ToolButtonStyle {
	var __rv Qt_ToolButtonStyle
	q.Drv(383000,383130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
