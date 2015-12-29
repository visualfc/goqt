// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMenuBar : QMenuBar
type QMenuBar struct {
	QWidget
}
func (q *QMenuBar) OnHovered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(309000,309102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMenuBar) OnTriggered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(309000,309103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMenuBar::QMenuBar()	
func NewMenuBar() *QMenuBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),309000,309104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMenuBar{}
	_p.SetDriver(__rv,309000,false)
	return _p
} 
//QMenuBar::QMenuBar(QWidget*)	
func NewMenuBarWithParent(parent QWidgetInterface) *QMenuBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),309000,309105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMenuBar{}
	_p.SetDriver(__rv,309000,false)
	return _p
} 
//QMenuBar::actionAt(QPoint const&)
func (q *QMenuBar) ActionAt(value *QPoint) *QAction {
	var __rv uintptr
	q.Drv(309000,309106,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::actionEvent(QActionEvent*)
func (q *QMenuBar) ActionEvent(value *QActionEvent)  {
	q.Drv(309000,309107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::actionGeometry(QAction*)
func (q *QMenuBar) ActionGeometry(value *QAction) *QRect {
	var __rv uintptr
	q.Drv(309000,309108,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QMenuBar::activeAction()
func (q *QMenuBar) ActiveAction() *QAction {
	var __rv uintptr
	q.Drv(309000,309109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::addAction(QAction*)
func (q *QMenuBar) AddAction(action *QAction)  {
	q.Drv(309000,309110,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::addAction(QString const&)
func (q *QMenuBar) AddActionWithText(text string) *QAction {
	var __rv uintptr
	q.Drv(309000,309111,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::addAction(QString const&,QObject const*,char const*)
func (q *QMenuBar) AddActionWithTextObjectMember(text string,receiver QObjectInterface,member string) *QAction {
	var __rv uintptr
	q.Drv(309000,309112,unsafe.Pointer(&text),Native(receiver),unsafe.Pointer(&member),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::addMenu(QMenu*)
func (q *QMenuBar) AddMenu(menu *QMenu) *QAction {
	var __rv uintptr
	q.Drv(309000,309113,Native(menu),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::addMenu(QString const&)
func (q *QMenuBar) AddMenuWithTitle(title string) *QMenu {
	var __rv uintptr
	q.Drv(309000,309114,unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMenuBar::addMenu(QIcon const&,QString const&)
func (q *QMenuBar) AddMenuWithIconTitle(icon *QIcon,title string) *QMenu {
	var __rv uintptr
	q.Drv(309000,309115,Native(icon),unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMenuBar::addSeparator()
func (q *QMenuBar) AddSeparator() *QAction {
	var __rv uintptr
	q.Drv(309000,309116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::changeEvent(QEvent*)
func (q *QMenuBar) ChangeEvent(value *QEvent)  {
	q.Drv(309000,309117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::clear()
func (q *QMenuBar) Clear()  {
	q.Drv(309000,309118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::cornerWidget()
func (q *QMenuBar) CornerWidget() *QWidget {
	var __rv uintptr
	q.Drv(309000,309119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMenuBar::cornerWidget(Qt::Corner)
func (q *QMenuBar) CornerWidgetWithCorner(corner Qt_Corner) *QWidget {
	var __rv uintptr
	q.Drv(309000,309120,unsafe.Pointer(&corner),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMenuBar::event(QEvent*)
func (q *QMenuBar) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(309000,309121,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenuBar::eventFilter(QObject*,QEvent*)
func (q *QMenuBar) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(309000,309122,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenuBar::focusInEvent(QFocusEvent*)
func (q *QMenuBar) FocusInEvent(value *QFocusEvent)  {
	q.Drv(309000,309123,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::focusOutEvent(QFocusEvent*)
func (q *QMenuBar) FocusOutEvent(value *QFocusEvent)  {
	q.Drv(309000,309124,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::heightForWidth(int)
func (q *QMenuBar) HeightForWidth(value int) int {
	var __rv int
	q.Drv(309000,309125,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenuBar::insertMenu(QAction*,QMenu*)
func (q *QMenuBar) InsertMenu(before *QAction,menu *QMenu) *QAction {
	var __rv uintptr
	q.Drv(309000,309126,Native(before),Native(menu),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::insertSeparator(QAction*)
func (q *QMenuBar) InsertSeparator(before *QAction) *QAction {
	var __rv uintptr
	q.Drv(309000,309127,Native(before),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenuBar::isDefaultUp()
func (q *QMenuBar) IsDefaultUp() bool {
	var __rv bool
	q.Drv(309000,309128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenuBar::isNativeMenuBar()
func (q *QMenuBar) IsNativeMenuBar() bool {
	var __rv bool
	q.Drv(309000,309129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenuBar::keyPressEvent(QKeyEvent*)
func (q *QMenuBar) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(309000,309130,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::leaveEvent(QEvent*)
func (q *QMenuBar) LeaveEvent(value *QEvent)  {
	q.Drv(309000,309131,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::minimumSizeHint()
func (q *QMenuBar) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(309000,309132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMenuBar::mouseMoveEvent(QMouseEvent*)
func (q *QMenuBar) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(309000,309133,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::mousePressEvent(QMouseEvent*)
func (q *QMenuBar) MousePressEvent(value *QMouseEvent)  {
	q.Drv(309000,309134,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::mouseReleaseEvent(QMouseEvent*)
func (q *QMenuBar) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(309000,309135,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::paintEvent(QPaintEvent*)
func (q *QMenuBar) PaintEvent(value *QPaintEvent)  {
	q.Drv(309000,309136,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::resizeEvent(QResizeEvent*)
func (q *QMenuBar) ResizeEvent(value *QResizeEvent)  {
	q.Drv(309000,309137,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setActiveAction(QAction*)
func (q *QMenuBar) SetActiveAction(action *QAction)  {
	q.Drv(309000,309138,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setCornerWidget(QWidget*)
func (q *QMenuBar) SetCornerWidget(w QWidgetInterface)  {
	q.Drv(309000,309139,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setCornerWidget(QWidget*,Qt::Corner)
func (q *QMenuBar) SetCornerWidgetWithWidgetCorner(w QWidgetInterface,corner Qt_Corner)  {
	q.Drv(309000,309140,Native(w),unsafe.Pointer(&corner),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setDefaultUp(bool)
func (q *QMenuBar) SetDefaultUp(value bool)  {
	q.Drv(309000,309141,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setNativeMenuBar(bool)
func (q *QMenuBar) SetNativeMenuBar(nativeMenuBar bool)  {
	q.Drv(309000,309142,unsafe.Pointer(&nativeMenuBar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::setVisible(bool)
func (q *QMenuBar) SetVisible(visible bool)  {
	q.Drv(309000,309143,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenuBar::sizeHint()
func (q *QMenuBar) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(309000,309144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMenuBar::timerEvent(QTimerEvent*)
func (q *QMenuBar) TimerEvent(value *QTimerEvent)  {
	q.Drv(309000,309145,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
