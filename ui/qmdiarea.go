// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMdiArea_WindowOrder - QMdiArea::WindowOrder
type QMdiArea_WindowOrder uint32
const (
	QMdiArea_CreationOrder QMdiArea_WindowOrder = 0
	QMdiArea_StackingOrder QMdiArea_WindowOrder = 1
	QMdiArea_ActivationHistoryOrder QMdiArea_WindowOrder = 2
)
//enum QMdiArea_ViewMode - QMdiArea::ViewMode
type QMdiArea_ViewMode uint32
const (
	QMdiArea_SubWindowView QMdiArea_ViewMode = 0
	QMdiArea_TabbedView QMdiArea_ViewMode = 1
)
//enum QMdiArea_AreaOption - QMdiArea::AreaOption
type QMdiArea_AreaOption uint32
const (
	QMdiArea_DontMaximizeSubWindowOnActivation QMdiArea_AreaOption = 0x1
)
//struct QMdiArea : QMdiArea
type QMdiArea struct {
	QAbstractScrollArea
}
func (q *QMdiArea) OnSubWindowActivated(fn func(*QMdiSubWindow)) uintptr {
	var __rv uintptr
	q.Drv(306000,306102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMdiArea::QMdiArea()	
func NewMdiArea() *QMdiArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),306000,306103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMdiArea{}
	_p.SetDriver(__rv,306000,false)
	return _p
} 
//QMdiArea::QMdiArea(QWidget*)	
func NewMdiAreaWithParent(parent QWidgetInterface) *QMdiArea {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),306000,306104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMdiArea{}
	_p.SetDriver(__rv,306000,false)
	return _p
} 
//QMdiArea::activateNextSubWindow()
func (q *QMdiArea) ActivateNextSubWindow()  {
	q.Drv(306000,306105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::activatePreviousSubWindow()
func (q *QMdiArea) ActivatePreviousSubWindow()  {
	q.Drv(306000,306106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::activationOrder()
func (q *QMdiArea) ActivationOrder() QMdiArea_WindowOrder {
	var __rv QMdiArea_WindowOrder
	q.Drv(306000,306107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::activeSubWindow()
func (q *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	var __rv uintptr
	q.Drv(306000,306108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMdiSubWindow{}
	_rp.SetDriver(__rv,307000,false)
	return _rp
}	
//QMdiArea::addSubWindow(QWidget*)
func (q *QMdiArea) AddSubWindow(widget QWidgetInterface) *QMdiSubWindow {
	var __rv uintptr
	q.Drv(306000,306109,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMdiSubWindow{}
	_rp.SetDriver(__rv,307000,false)
	return _rp
}	
//QMdiArea::addSubWindow(QWidget*,QFlags<Qt::WindowType>)
func (q *QMdiArea) AddSubWindowWithWidgetFlags(widget QWidgetInterface,flags Qt_WindowType) *QMdiSubWindow {
	var __rv uintptr
	q.Drv(306000,306110,Native(widget),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMdiSubWindow{}
	_rp.SetDriver(__rv,307000,false)
	return _rp
}	
//QMdiArea::background()
func (q *QMdiArea) Background() *QBrush {
	var __rv uintptr
	q.Drv(306000,306111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QMdiArea::cascadeSubWindows()
func (q *QMdiArea) CascadeSubWindows()  {
	q.Drv(306000,306112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::childEvent(QChildEvent*)
func (q *QMdiArea) ChildEvent(childEvent *QChildEvent)  {
	q.Drv(306000,306113,Native(childEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::closeActiveSubWindow()
func (q *QMdiArea) CloseActiveSubWindow()  {
	q.Drv(306000,306114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::closeAllSubWindows()
func (q *QMdiArea) CloseAllSubWindows()  {
	q.Drv(306000,306115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::currentSubWindow()
func (q *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	var __rv uintptr
	q.Drv(306000,306116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMdiSubWindow{}
	_rp.SetDriver(__rv,307000,false)
	return _rp
}	
//QMdiArea::documentMode()
func (q *QMdiArea) DocumentMode() bool {
	var __rv bool
	q.Drv(306000,306117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::event(QEvent*)
func (q *QMdiArea) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(306000,306118,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::eventFilter(QObject*,QEvent*)
func (q *QMdiArea) EventFilter(object QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(306000,306119,Native(object),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::minimumSizeHint()
func (q *QMdiArea) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(306000,306120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMdiArea::paintEvent(QPaintEvent*)
func (q *QMdiArea) PaintEvent(paintEvent *QPaintEvent)  {
	q.Drv(306000,306121,Native(paintEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::removeSubWindow(QWidget*)
func (q *QMdiArea) RemoveSubWindow(widget QWidgetInterface)  {
	q.Drv(306000,306122,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::resizeEvent(QResizeEvent*)
func (q *QMdiArea) ResizeEvent(resizeEvent *QResizeEvent)  {
	q.Drv(306000,306123,Native(resizeEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::scrollContentsBy(int,int)
func (q *QMdiArea) ScrollContentsBy(dx int,dy int)  {
	q.Drv(306000,306124,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setActivationOrder(QMdiArea::WindowOrder)
func (q *QMdiArea) SetActivationOrder(order QMdiArea_WindowOrder)  {
	q.Drv(306000,306125,unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setActiveSubWindow(QMdiSubWindow*)
func (q *QMdiArea) SetActiveSubWindow(window *QMdiSubWindow)  {
	q.Drv(306000,306126,Native(window),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setBackground(QBrush const&)
func (q *QMdiArea) SetBackground(background *QBrush)  {
	q.Drv(306000,306127,Native(background),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setDocumentMode(bool)
func (q *QMdiArea) SetDocumentMode(enabled bool)  {
	q.Drv(306000,306128,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setOption(QMdiArea::AreaOption)
func (q *QMdiArea) SetOption(option QMdiArea_AreaOption)  {
	q.Drv(306000,306129,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setOption(QMdiArea::AreaOption,bool)
func (q *QMdiArea) SetOptionWithOptionOn(option QMdiArea_AreaOption,on bool)  {
	q.Drv(306000,306130,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setTabPosition(QTabWidget::TabPosition)
func (q *QMdiArea) SetTabPosition(position QTabWidget_TabPosition)  {
	q.Drv(306000,306131,unsafe.Pointer(&position),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setTabShape(QTabWidget::TabShape)
func (q *QMdiArea) SetTabShape(shape QTabWidget_TabShape)  {
	q.Drv(306000,306132,unsafe.Pointer(&shape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setViewMode(QMdiArea::ViewMode)
func (q *QMdiArea) SetViewMode(mode QMdiArea_ViewMode)  {
	q.Drv(306000,306133,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::setupViewport(QWidget*)
func (q *QMdiArea) SetupViewport(viewport QWidgetInterface)  {
	q.Drv(306000,306134,Native(viewport),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::showEvent(QShowEvent*)
func (q *QMdiArea) ShowEvent(showEvent *QShowEvent)  {
	q.Drv(306000,306135,Native(showEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::sizeHint()
func (q *QMdiArea) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(306000,306136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMdiArea::subWindowList()
func (q *QMdiArea) SubWindowList() []*QMdiSubWindow {
	var __rv []*QMdiSubWindow
	q.Drv(306000,306137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::subWindowList(QMdiArea::WindowOrder)
func (q *QMdiArea) SubWindowListWithOrder(order QMdiArea_WindowOrder) []*QMdiSubWindow {
	var __rv []*QMdiSubWindow
	q.Drv(306000,306138,unsafe.Pointer(&order),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::tabPosition()
func (q *QMdiArea) TabPosition() QTabWidget_TabPosition {
	var __rv QTabWidget_TabPosition
	q.Drv(306000,306139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::tabShape()
func (q *QMdiArea) TabShape() QTabWidget_TabShape {
	var __rv QTabWidget_TabShape
	q.Drv(306000,306140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::testOption(QMdiArea::AreaOption)
func (q *QMdiArea) TestOption(opton QMdiArea_AreaOption) bool {
	var __rv bool
	q.Drv(306000,306141,unsafe.Pointer(&opton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::tileSubWindows()
func (q *QMdiArea) TileSubWindows()  {
	q.Drv(306000,306142,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::timerEvent(QTimerEvent*)
func (q *QMdiArea) TimerEvent(timerEvent *QTimerEvent)  {
	q.Drv(306000,306143,Native(timerEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMdiArea::viewMode()
func (q *QMdiArea) ViewMode() QMdiArea_ViewMode {
	var __rv QMdiArea_ViewMode
	q.Drv(306000,306144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMdiArea::viewportEvent(QEvent*)
func (q *QMdiArea) ViewportEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(306000,306145,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
