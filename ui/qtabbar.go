// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTabBar_SelectionBehavior - QTabBar::SelectionBehavior
type QTabBar_SelectionBehavior uint32
const (
	QTabBar_SelectLeftTab QTabBar_SelectionBehavior = 0
	QTabBar_SelectRightTab QTabBar_SelectionBehavior = 1
	QTabBar_SelectPreviousTab QTabBar_SelectionBehavior = 2
)
//enum QTabBar_ButtonPosition - QTabBar::ButtonPosition
type QTabBar_ButtonPosition uint32
const (
	QTabBar_LeftSide QTabBar_ButtonPosition = 0
	QTabBar_RightSide QTabBar_ButtonPosition = 1
)
//enum QTabBar_Shape - QTabBar::Shape
type QTabBar_Shape uint32
const (
	QTabBar_RoundedNorth QTabBar_Shape = 0
	QTabBar_RoundedSouth QTabBar_Shape = 1
	QTabBar_RoundedWest QTabBar_Shape = 2
	QTabBar_RoundedEast QTabBar_Shape = 3
	QTabBar_TriangularNorth QTabBar_Shape = 4
	QTabBar_TriangularSouth QTabBar_Shape = 5
	QTabBar_TriangularWest QTabBar_Shape = 6
	QTabBar_TriangularEast QTabBar_Shape = 7
)
//struct QTabBar : QTabBar
type QTabBar struct {
	QWidget
}
func (q *QTabBar) OnTabMoved(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(363000,363102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTabBar) OnCurrentChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(363000,363103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTabBar) OnTabCloseRequested(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(363000,363104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTabBar::QTabBar()	
func NewTabBar() *QTabBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),363000,363105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTabBar{}
	_p.SetDriver(__rv,363000,false)
	return _p
} 
//QTabBar::QTabBar(QWidget*)	
func NewTabBarWithParent(parent QWidgetInterface) *QTabBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),363000,363106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTabBar{}
	_p.SetDriver(__rv,363000,false)
	return _p
} 
//QTabBar::addTab(QString const&)
func (q *QTabBar) AddTab(text string) int {
	var __rv int
	q.Drv(363000,363107,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::addTab(QIcon const&,QString const&)
func (q *QTabBar) AddTabWithIconText(icon *QIcon,text string) int {
	var __rv int
	q.Drv(363000,363108,Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::changeEvent(QEvent*)
func (q *QTabBar) ChangeEvent(value *QEvent)  {
	q.Drv(363000,363109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::count()
func (q *QTabBar) Count() int {
	var __rv int
	q.Drv(363000,363110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::currentIndex()
func (q *QTabBar) CurrentIndex() int {
	var __rv int
	q.Drv(363000,363111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::documentMode()
func (q *QTabBar) DocumentMode() bool {
	var __rv bool
	q.Drv(363000,363112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::drawBase()
func (q *QTabBar) DrawBase() bool {
	var __rv bool
	q.Drv(363000,363113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::elideMode()
func (q *QTabBar) ElideMode() Qt_TextElideMode {
	var __rv Qt_TextElideMode
	q.Drv(363000,363114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::event(QEvent*)
func (q *QTabBar) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(363000,363115,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::expanding()
func (q *QTabBar) Expanding() bool {
	var __rv bool
	q.Drv(363000,363116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::hideEvent(QHideEvent*)
func (q *QTabBar) HideEvent(value *QHideEvent)  {
	q.Drv(363000,363117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::iconSize()
func (q *QTabBar) IconSize() *QSize {
	var __rv uintptr
	q.Drv(363000,363118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabBar::insertTab(int,QString const&)
func (q *QTabBar) InsertTabWithIndexText(index int,text string) int {
	var __rv int
	q.Drv(363000,363119,unsafe.Pointer(&index),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::insertTab(int,QIcon const&,QString const&)
func (q *QTabBar) InsertTabWithIndexIconText(index int,icon *QIcon,text string) int {
	var __rv int
	q.Drv(363000,363120,unsafe.Pointer(&index),Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::isMovable()
func (q *QTabBar) IsMovable() bool {
	var __rv bool
	q.Drv(363000,363121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::isTabEnabled(int)
func (q *QTabBar) IsTabEnabled(index int) bool {
	var __rv bool
	q.Drv(363000,363122,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::keyPressEvent(QKeyEvent*)
func (q *QTabBar) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(363000,363123,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::minimumSizeHint()
func (q *QTabBar) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(363000,363124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabBar::mouseMoveEvent(QMouseEvent*)
func (q *QTabBar) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(363000,363125,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::mousePressEvent(QMouseEvent*)
func (q *QTabBar) MousePressEvent(value *QMouseEvent)  {
	q.Drv(363000,363126,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::mouseReleaseEvent(QMouseEvent*)
func (q *QTabBar) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(363000,363127,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::moveTab(int,int)
func (q *QTabBar) MoveTab(from int,to int)  {
	q.Drv(363000,363128,unsafe.Pointer(&from),unsafe.Pointer(&to),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::paintEvent(QPaintEvent*)
func (q *QTabBar) PaintEvent(value *QPaintEvent)  {
	q.Drv(363000,363129,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::removeTab(int)
func (q *QTabBar) RemoveTab(index int)  {
	q.Drv(363000,363130,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::resizeEvent(QResizeEvent*)
func (q *QTabBar) ResizeEvent(value *QResizeEvent)  {
	q.Drv(363000,363131,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::selectionBehaviorOnRemove()
func (q *QTabBar) SelectionBehaviorOnRemove() QTabBar_SelectionBehavior {
	var __rv QTabBar_SelectionBehavior
	q.Drv(363000,363132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::setCurrentIndex(int)
func (q *QTabBar) SetCurrentIndex(index int)  {
	q.Drv(363000,363133,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setDocumentMode(bool)
func (q *QTabBar) SetDocumentMode(set bool)  {
	q.Drv(363000,363134,unsafe.Pointer(&set),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setDrawBase(bool)
func (q *QTabBar) SetDrawBase(drawTheBase bool)  {
	q.Drv(363000,363135,unsafe.Pointer(&drawTheBase),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setElideMode(Qt::TextElideMode)
func (q *QTabBar) SetElideMode(value Qt_TextElideMode)  {
	q.Drv(363000,363136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setExpanding(bool)
func (q *QTabBar) SetExpanding(enabled bool)  {
	q.Drv(363000,363137,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setIconSize(QSize const&)
func (q *QTabBar) SetIconSize(size *QSize)  {
	q.Drv(363000,363138,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setMovable(bool)
func (q *QTabBar) SetMovable(movable bool)  {
	q.Drv(363000,363139,unsafe.Pointer(&movable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setSelectionBehaviorOnRemove(QTabBar::SelectionBehavior)
func (q *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar_SelectionBehavior)  {
	q.Drv(363000,363140,unsafe.Pointer(&behavior),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setShape(QTabBar::Shape)
func (q *QTabBar) SetShape(shape QTabBar_Shape)  {
	q.Drv(363000,363141,unsafe.Pointer(&shape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabButton(int,QTabBar::ButtonPosition,QWidget*)
func (q *QTabBar) SetTabButton(index int,position QTabBar_ButtonPosition,widget QWidgetInterface)  {
	q.Drv(363000,363142,unsafe.Pointer(&index),unsafe.Pointer(&position),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabData(int,QVariant const&)
func (q *QTabBar) SetTabData(index int,data *QVariant)  {
	q.Drv(363000,363143,unsafe.Pointer(&index),Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabEnabled(int,bool)
func (q *QTabBar) SetTabEnabled(index int,value2 bool)  {
	q.Drv(363000,363144,unsafe.Pointer(&index),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabIcon(int,QIcon const&)
func (q *QTabBar) SetTabIcon(index int,icon *QIcon)  {
	q.Drv(363000,363145,unsafe.Pointer(&index),Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabText(int,QString const&)
func (q *QTabBar) SetTabText(index int,text string)  {
	q.Drv(363000,363146,unsafe.Pointer(&index),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabTextColor(int,QColor const&)
func (q *QTabBar) SetTabTextColor(index int,color *QColor)  {
	q.Drv(363000,363147,unsafe.Pointer(&index),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabToolTip(int,QString const&)
func (q *QTabBar) SetTabToolTip(index int,tip string)  {
	q.Drv(363000,363148,unsafe.Pointer(&index),unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabWhatsThis(int,QString const&)
func (q *QTabBar) SetTabWhatsThis(index int,text string)  {
	q.Drv(363000,363149,unsafe.Pointer(&index),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setTabsClosable(bool)
func (q *QTabBar) SetTabsClosable(closable bool)  {
	q.Drv(363000,363150,unsafe.Pointer(&closable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::setUsesScrollButtons(bool)
func (q *QTabBar) SetUsesScrollButtons(useButtons bool)  {
	q.Drv(363000,363151,unsafe.Pointer(&useButtons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::shape()
func (q *QTabBar) Shape() QTabBar_Shape {
	var __rv QTabBar_Shape
	q.Drv(363000,363152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::showEvent(QShowEvent*)
func (q *QTabBar) ShowEvent(value *QShowEvent)  {
	q.Drv(363000,363153,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::sizeHint()
func (q *QTabBar) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(363000,363154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabBar::tabAt(QPoint const&)
func (q *QTabBar) TabAt(pos *QPoint) int {
	var __rv int
	q.Drv(363000,363155,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::tabButton(int,QTabBar::ButtonPosition)
func (q *QTabBar) TabButton(index int,position QTabBar_ButtonPosition) *QWidget {
	var __rv uintptr
	q.Drv(363000,363156,unsafe.Pointer(&index),unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTabBar::tabData(int)
func (q *QTabBar) TabData(index int) *QVariant {
	var __rv uintptr
	q.Drv(363000,363157,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTabBar::tabIcon(int)
func (q *QTabBar) TabIcon(index int) *QIcon {
	var __rv uintptr
	q.Drv(363000,363158,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QTabBar::tabInserted(int)
func (q *QTabBar) TabInserted(index int)  {
	q.Drv(363000,363159,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::tabLayoutChange()
func (q *QTabBar) TabLayoutChange()  {
	q.Drv(363000,363160,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::tabRect(int)
func (q *QTabBar) TabRect(index int) *QRect {
	var __rv uintptr
	q.Drv(363000,363161,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTabBar::tabRemoved(int)
func (q *QTabBar) TabRemoved(index int)  {
	q.Drv(363000,363162,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabBar::tabSizeHint(int)
func (q *QTabBar) TabSizeHint(index int) *QSize {
	var __rv uintptr
	q.Drv(363000,363163,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabBar::tabText(int)
func (q *QTabBar) TabText(index int) string {
	var __rv string
	q.Drv(363000,363164,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::tabTextColor(int)
func (q *QTabBar) TabTextColor(index int) *QColor {
	var __rv uintptr
	q.Drv(363000,363165,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTabBar::tabToolTip(int)
func (q *QTabBar) TabToolTip(index int) string {
	var __rv string
	q.Drv(363000,363166,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::tabWhatsThis(int)
func (q *QTabBar) TabWhatsThis(index int) string {
	var __rv string
	q.Drv(363000,363167,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::tabsClosable()
func (q *QTabBar) TabsClosable() bool {
	var __rv bool
	q.Drv(363000,363168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::usesScrollButtons()
func (q *QTabBar) UsesScrollButtons() bool {
	var __rv bool
	q.Drv(363000,363169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabBar::wheelEvent(QWheelEvent*)
func (q *QTabBar) WheelEvent(event *QWheelEvent)  {
	q.Drv(363000,363170,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
