// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMdiSubWindow_SubWindowOption - QMdiSubWindow::SubWindowOption
type QMdiSubWindow_SubWindowOption uint32
const (
	QMdiSubWindow_AllowOutsideAreaHorizontally QMdiSubWindow_SubWindowOption = 0x1
	QMdiSubWindow_AllowOutsideAreaVertically QMdiSubWindow_SubWindowOption = 0x2
	QMdiSubWindow_RubberBandResize QMdiSubWindow_SubWindowOption = 0x4
	QMdiSubWindow_RubberBandMove QMdiSubWindow_SubWindowOption = 0x8
)
//struct QMdiSubWindow : QMdiSubWindow
type QMdiSubWindow struct {
	QWidget
}
func (q *QMdiSubWindow) OnAboutToActivate(fn func()) uintptr {
	var __rv uintptr
	q.Drv(307000,307102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMdiSubWindow) OnWindowStateChanged(fn func(Qt_WindowState,Qt_WindowState)) uintptr {
	var __rv uintptr
	q.Drv(307000,307103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMdiSubWindow::QMdiSubWindow()	
func NewMdiSubWindow() *QMdiSubWindow {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),307000,307104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMdiSubWindow{}
	_p.SetDriver(__rv,307000,false)
	return _p
} 
//QMdiSubWindow::QMdiSubWindow(QWidget*,QFlags<Qt::WindowType>)	
func NewMdiSubWindowWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QMdiSubWindow {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),307000,307105,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMdiSubWindow{}
	_p.SetDriver(__rv,307000,false)
	return _p
} 
//QMdiSubWindow::changeEvent(QEvent*)
func (q *QMdiSubWindow) ChangeEvent(changeEvent *QEvent)  {
	q.Drv(307000,307106,Native(changeEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::childEvent(QChildEvent*)
func (q *QMdiSubWindow) ChildEvent(childEvent *QChildEvent)  {
	q.Drv(307000,307107,Native(childEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::closeEvent(QCloseEvent*)
func (q *QMdiSubWindow) CloseEvent(closeEvent *QCloseEvent)  {
	q.Drv(307000,307108,Native(closeEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::contextMenuEvent(QContextMenuEvent*)
func (q *QMdiSubWindow) ContextMenuEvent(contextMenuEvent *QContextMenuEvent)  {
	q.Drv(307000,307109,Native(contextMenuEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::event(QEvent*)
func (q *QMdiSubWindow) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(307000,307110,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::eventFilter(QObject*,QEvent*)
func (q *QMdiSubWindow) EventFilter(object QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(307000,307111,Native(object),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::focusInEvent(QFocusEvent*)
func (q *QMdiSubWindow) FocusInEvent(focusInEvent *QFocusEvent)  {
	q.Drv(307000,307112,Native(focusInEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::focusOutEvent(QFocusEvent*)
func (q *QMdiSubWindow) FocusOutEvent(focusOutEvent *QFocusEvent)  {
	q.Drv(307000,307113,Native(focusOutEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::hideEvent(QHideEvent*)
func (q *QMdiSubWindow) HideEvent(hideEvent *QHideEvent)  {
	q.Drv(307000,307114,Native(hideEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::isShaded()
func (q *QMdiSubWindow) IsShaded() bool {
	var __rv bool
	q.Drv(307000,307115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::keyPressEvent(QKeyEvent*)
func (q *QMdiSubWindow) KeyPressEvent(keyEvent *QKeyEvent)  {
	q.Drv(307000,307116,Native(keyEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::keyboardPageStep()
func (q *QMdiSubWindow) KeyboardPageStep() int {
	var __rv int
	q.Drv(307000,307117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::keyboardSingleStep()
func (q *QMdiSubWindow) KeyboardSingleStep() int {
	var __rv int
	q.Drv(307000,307118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::leaveEvent(QEvent*)
func (q *QMdiSubWindow) LeaveEvent(leaveEvent *QEvent)  {
	q.Drv(307000,307119,Native(leaveEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::maximizedButtonsWidget()
func (q *QMdiSubWindow) MaximizedButtonsWidget() *QWidget {
	var __rv uintptr
	q.Drv(307000,307120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMdiSubWindow::maximizedSystemMenuIconWidget()
func (q *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget {
	var __rv uintptr
	q.Drv(307000,307121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMdiSubWindow::mdiArea()
func (q *QMdiSubWindow) MdiArea() *QMdiArea {
	var __rv uintptr
	q.Drv(307000,307122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMdiArea{}
	_rp.SetDriver(__rv,306000,false)
	return _rp
}	
//QMdiSubWindow::minimumSizeHint()
func (q *QMdiSubWindow) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(307000,307123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMdiSubWindow::mouseDoubleClickEvent(QMouseEvent*)
func (q *QMdiSubWindow) MouseDoubleClickEvent(mouseEvent *QMouseEvent)  {
	q.Drv(307000,307124,Native(mouseEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::mouseMoveEvent(QMouseEvent*)
func (q *QMdiSubWindow) MouseMoveEvent(mouseEvent *QMouseEvent)  {
	q.Drv(307000,307125,Native(mouseEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::mousePressEvent(QMouseEvent*)
func (q *QMdiSubWindow) MousePressEvent(mouseEvent *QMouseEvent)  {
	q.Drv(307000,307126,Native(mouseEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::mouseReleaseEvent(QMouseEvent*)
func (q *QMdiSubWindow) MouseReleaseEvent(mouseEvent *QMouseEvent)  {
	q.Drv(307000,307127,Native(mouseEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::moveEvent(QMoveEvent*)
func (q *QMdiSubWindow) MoveEvent(moveEvent *QMoveEvent)  {
	q.Drv(307000,307128,Native(moveEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::paintEvent(QPaintEvent*)
func (q *QMdiSubWindow) PaintEvent(paintEvent *QPaintEvent)  {
	q.Drv(307000,307129,Native(paintEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::resizeEvent(QResizeEvent*)
func (q *QMdiSubWindow) ResizeEvent(resizeEvent *QResizeEvent)  {
	q.Drv(307000,307130,Native(resizeEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setKeyboardPageStep(int)
func (q *QMdiSubWindow) SetKeyboardPageStep(step int)  {
	q.Drv(307000,307131,unsafe.Pointer(&step),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setKeyboardSingleStep(int)
func (q *QMdiSubWindow) SetKeyboardSingleStep(step int)  {
	q.Drv(307000,307132,unsafe.Pointer(&step),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setOption(QMdiSubWindow::SubWindowOption)
func (q *QMdiSubWindow) SetOption(option QMdiSubWindow_SubWindowOption)  {
	q.Drv(307000,307133,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setOption(QMdiSubWindow::SubWindowOption,bool)
func (q *QMdiSubWindow) SetOptionWithOptionOn(option QMdiSubWindow_SubWindowOption,on bool)  {
	q.Drv(307000,307134,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setSystemMenu(QMenu*)
func (q *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu)  {
	q.Drv(307000,307135,Native(systemMenu),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::setWidget(QWidget*)
func (q *QMdiSubWindow) SetWidget(widget QWidgetInterface)  {
	q.Drv(307000,307136,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::showEvent(QShowEvent*)
func (q *QMdiSubWindow) ShowEvent(showEvent *QShowEvent)  {
	q.Drv(307000,307137,Native(showEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::showShaded()
func (q *QMdiSubWindow) ShowShaded()  {
	q.Drv(307000,307138,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::showSystemMenu()
func (q *QMdiSubWindow) ShowSystemMenu()  {
	q.Drv(307000,307139,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::sizeHint()
func (q *QMdiSubWindow) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(307000,307140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMdiSubWindow::systemMenu()
func (q *QMdiSubWindow) SystemMenu() *QMenu {
	var __rv uintptr
	q.Drv(307000,307141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMdiSubWindow::testOption(QMdiSubWindow::SubWindowOption)
func (q *QMdiSubWindow) TestOption(value QMdiSubWindow_SubWindowOption) bool {
	var __rv bool
	q.Drv(307000,307142,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiSubWindow::timerEvent(QTimerEvent*)
func (q *QMdiSubWindow) TimerEvent(timerEvent *QTimerEvent)  {
	q.Drv(307000,307143,Native(timerEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiSubWindow::widget()
func (q *QMdiSubWindow) Widget() *QWidget {
	var __rv uintptr
	q.Drv(307000,307144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
