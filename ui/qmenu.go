// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMenu : QMenu
type QMenu struct {
	QWidget
}
func (q *QMenu) OnHovered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(308000,308102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMenu) OnAboutToHide(fn func()) uintptr {
	var __rv uintptr
	q.Drv(308000,308103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMenu) OnTriggered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(308000,308104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMenu) OnAboutToShow(fn func()) uintptr {
	var __rv uintptr
	q.Drv(308000,308105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMenu::QMenu()	
func NewMenu() *QMenu {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),308000,308106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMenu{}
	_p.SetDriver(__rv,308000,false)
	return _p
} 
//QMenu::QMenu(QWidget*)	
func NewMenuWithParent(parent QWidgetInterface) *QMenu {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),308000,308107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMenu{}
	_p.SetDriver(__rv,308000,false)
	return _p
} 
//QMenu::QMenu(QString const&,QWidget*)	
func NewMenuWithTitleParent(title string,parent QWidgetInterface) *QMenu {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),308000,308108,unsafe.Pointer(&title),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMenu{}
	_p.SetDriver(__rv,308000,false)
	return _p
} 
//QMenu::actionAt(QPoint const&)
func (q *QMenu) ActionAt(value *QPoint) *QAction {
	var __rv uintptr
	q.Drv(308000,308109,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::actionEvent(QActionEvent*)
func (q *QMenu) ActionEvent(value *QActionEvent)  {
	q.Drv(308000,308110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::actionGeometry(QAction*)
func (q *QMenu) ActionGeometry(value *QAction) *QRect {
	var __rv uintptr
	q.Drv(308000,308111,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QMenu::activeAction()
func (q *QMenu) ActiveAction() *QAction {
	var __rv uintptr
	q.Drv(308000,308112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addAction(QAction*)
func (q *QMenu) AddAction(action *QAction)  {
	q.Drv(308000,308113,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::addAction(QString const&)
func (q *QMenu) AddActionWithText(text string) *QAction {
	var __rv uintptr
	q.Drv(308000,308114,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addAction(QIcon const&,QString const&)
func (q *QMenu) AddActionWithIconText(icon *QIcon,text string) *QAction {
	var __rv uintptr
	q.Drv(308000,308115,Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addAction(QString const&,QObject const*,char const*,QKeySequence const&)
func (q *QMenu) AddActionWithTextObjectMemberShortcut(text string,receiver QObjectInterface,member string,shortcut *QKeySequence) *QAction {
	var __rv uintptr
	q.Drv(308000,308116,unsafe.Pointer(&text),Native(receiver),unsafe.Pointer(&member),Native(shortcut),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addAction(QIcon const&,QString const&,QObject const*,char const*,QKeySequence const&)
func (q *QMenu) AddActionWithIconTextObjectMemberShortcut(icon *QIcon,text string,receiver QObjectInterface,member string,shortcut *QKeySequence) *QAction {
	var __rv uintptr
	q.Drv(308000,308117,Native(icon),unsafe.Pointer(&text),Native(receiver),unsafe.Pointer(&member),Native(shortcut),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addMenu(QMenu*)
func (q *QMenu) AddMenu(menu *QMenu) *QAction {
	var __rv uintptr
	q.Drv(308000,308118,Native(menu),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::addMenu(QString const&)
func (q *QMenu) AddMenuWithTitle(title string) *QMenu {
	var __rv uintptr
	q.Drv(308000,308119,unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMenu::addMenu(QIcon const&,QString const&)
func (q *QMenu) AddMenuWithIconTitle(icon *QIcon,title string) *QMenu {
	var __rv uintptr
	q.Drv(308000,308120,Native(icon),unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMenu::addSeparator()
func (q *QMenu) AddSeparator() *QAction {
	var __rv uintptr
	q.Drv(308000,308121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::changeEvent(QEvent*)
func (q *QMenu) ChangeEvent(value *QEvent)  {
	q.Drv(308000,308122,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::clear()
func (q *QMenu) Clear()  {
	q.Drv(308000,308123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::columnCount()
func (q *QMenu) ColumnCount() int {
	var __rv int
	q.Drv(308000,308124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::defaultAction()
func (q *QMenu) DefaultAction() *QAction {
	var __rv uintptr
	q.Drv(308000,308125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::enterEvent(QEvent*)
func (q *QMenu) EnterEvent(value *QEvent)  {
	q.Drv(308000,308126,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::event(QEvent*)
func (q *QMenu) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(308000,308127,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::exec()
func (q *QMenu) Exec() *QAction {
	var __rv uintptr
	q.Drv(308000,308128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QPoint const&)
func (q *QMenu) ExecWithPos(pos *QPoint) *QAction {
	var __rv uintptr
	q.Drv(308000,308129,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QPoint const&,QAction*)
func (q *QMenu) ExecWithPosAt(pos *QPoint,at *QAction) *QAction {
	var __rv uintptr
	q.Drv(308000,308130,Native(pos),Native(at),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QList<QAction*>,QPoint const&,QAction*)	
func QMenuExecWithActionsPosAt(actions []*QAction,pos *QPoint,at *QAction) *QAction {
	var __rv uintptr
	DirectQtDrv(nil,308000,308131,unsafe.Pointer(&actions),Native(pos),Native(at),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QList<QAction*>,QPoint const&,QAction*)
func (q *QMenu) ExecWithActionsPosAt(actions []*QAction,pos *QPoint,at *QAction) *QAction {
	var __rv uintptr
	q.Drv(308000,308131,unsafe.Pointer(&actions),Native(pos),Native(at),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QList<QAction*>,QPoint const&,QAction*,QWidget*)	
func QMenuExecWithActionsPosAtParent(actions []*QAction,pos *QPoint,at *QAction,parent QWidgetInterface) *QAction {
	var __rv uintptr
	DirectQtDrv(nil,308000,308132,unsafe.Pointer(&actions),Native(pos),Native(at),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::exec(QList<QAction*>,QPoint const&,QAction*,QWidget*)
func (q *QMenu) ExecWithActionsPosAtParent(actions []*QAction,pos *QPoint,at *QAction,parent QWidgetInterface) *QAction {
	var __rv uintptr
	q.Drv(308000,308132,unsafe.Pointer(&actions),Native(pos),Native(at),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::focusNextPrevChild(bool)
func (q *QMenu) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(308000,308133,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::hideEvent(QHideEvent*)
func (q *QMenu) HideEvent(value *QHideEvent)  {
	q.Drv(308000,308134,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::hideTearOffMenu()
func (q *QMenu) HideTearOffMenu()  {
	q.Drv(308000,308135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::icon()
func (q *QMenu) Icon() *QIcon {
	var __rv uintptr
	q.Drv(308000,308136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QMenu::insertMenu(QAction*,QMenu*)
func (q *QMenu) InsertMenu(before *QAction,menu *QMenu) *QAction {
	var __rv uintptr
	q.Drv(308000,308137,Native(before),Native(menu),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::insertSeparator(QAction*)
func (q *QMenu) InsertSeparator(before *QAction) *QAction {
	var __rv uintptr
	q.Drv(308000,308138,Native(before),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::isEmpty()
func (q *QMenu) IsEmpty() bool {
	var __rv bool
	q.Drv(308000,308139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::isTearOffEnabled()
func (q *QMenu) IsTearOffEnabled() bool {
	var __rv bool
	q.Drv(308000,308140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::isTearOffMenuVisible()
func (q *QMenu) IsTearOffMenuVisible() bool {
	var __rv bool
	q.Drv(308000,308141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::keyPressEvent(QKeyEvent*)
func (q *QMenu) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(308000,308142,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::leaveEvent(QEvent*)
func (q *QMenu) LeaveEvent(value *QEvent)  {
	q.Drv(308000,308143,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::menuAction()
func (q *QMenu) MenuAction() *QAction {
	var __rv uintptr
	q.Drv(308000,308144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QMenu::mouseMoveEvent(QMouseEvent*)
func (q *QMenu) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(308000,308145,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::mousePressEvent(QMouseEvent*)
func (q *QMenu) MousePressEvent(value *QMouseEvent)  {
	q.Drv(308000,308146,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::mouseReleaseEvent(QMouseEvent*)
func (q *QMenu) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(308000,308147,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::paintEvent(QPaintEvent*)
func (q *QMenu) PaintEvent(value *QPaintEvent)  {
	q.Drv(308000,308148,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::popup(QPoint const&)
func (q *QMenu) Popup(pos *QPoint)  {
	q.Drv(308000,308149,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::popup(QPoint const&,QAction*)
func (q *QMenu) PopupWithPosAt(pos *QPoint,at *QAction)  {
	q.Drv(308000,308150,Native(pos),Native(at),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::separatorsCollapsible()
func (q *QMenu) SeparatorsCollapsible() bool {
	var __rv bool
	q.Drv(308000,308151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::setActiveAction(QAction*)
func (q *QMenu) SetActiveAction(act *QAction)  {
	q.Drv(308000,308152,Native(act),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setDefaultAction(QAction*)
func (q *QMenu) SetDefaultAction(value *QAction)  {
	q.Drv(308000,308153,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setIcon(QIcon const&)
func (q *QMenu) SetIcon(icon *QIcon)  {
	q.Drv(308000,308154,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setNoReplayFor(QWidget*)
func (q *QMenu) SetNoReplayFor(widget QWidgetInterface)  {
	q.Drv(308000,308155,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setSeparatorsCollapsible(bool)
func (q *QMenu) SetSeparatorsCollapsible(collapse bool)  {
	q.Drv(308000,308156,unsafe.Pointer(&collapse),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setTearOffEnabled(bool)
func (q *QMenu) SetTearOffEnabled(value bool)  {
	q.Drv(308000,308157,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::setTitle(QString const&)
func (q *QMenu) SetTitle(title string)  {
	q.Drv(308000,308158,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::sizeHint()
func (q *QMenu) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(308000,308159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMenu::timerEvent(QTimerEvent*)
func (q *QMenu) TimerEvent(value *QTimerEvent)  {
	q.Drv(308000,308160,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMenu::title()
func (q *QMenu) Title() string {
	var __rv string
	q.Drv(308000,308161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMenu::wheelEvent(QWheelEvent*)
func (q *QMenu) WheelEvent(value *QWheelEvent)  {
	q.Drv(308000,308162,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
