// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMainWindow_DockOption - QMainWindow::DockOption
type QMainWindow_DockOption uint32
const (
	QMainWindow_AnimatedDocks QMainWindow_DockOption = 0x01
	QMainWindow_AllowNestedDocks QMainWindow_DockOption = 0x02
	QMainWindow_AllowTabbedDocks QMainWindow_DockOption = 0x04
	QMainWindow_ForceTabbedDocks QMainWindow_DockOption = 0x08
	QMainWindow_VerticalTabs QMainWindow_DockOption = 0x10
)
//struct QMainWindow : QMainWindow
type QMainWindow struct {
	QWidget
}
func (q *QMainWindow) OnToolButtonStyleChanged(fn func(Qt_ToolButtonStyle)) uintptr {
	var __rv uintptr
	q.Drv(305000,305102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMainWindow) OnIconSizeChanged(fn func(*QSize)) uintptr {
	var __rv uintptr
	q.Drv(305000,305103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMainWindow::QMainWindow()	
func NewMainWindow() *QMainWindow {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),305000,305104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMainWindow{}
	_p.SetDriver(__rv,305000,false)
	return _p
} 
//QMainWindow::QMainWindow(QWidget*,QFlags<Qt::WindowType>)	
func NewMainWindowWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QMainWindow {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),305000,305105,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMainWindow{}
	_p.SetDriver(__rv,305000,false)
	return _p
} 
//QMainWindow::addDockWidget(Qt::DockWidgetArea,QDockWidget*)
func (q *QMainWindow) AddDockWidgetWithAreaDockwidget(area Qt_DockWidgetArea,dockwidget *QDockWidget)  {
	q.Drv(305000,305106,unsafe.Pointer(&area),Native(dockwidget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::addDockWidget(Qt::DockWidgetArea,QDockWidget*,Qt::Orientation)
func (q *QMainWindow) AddDockWidgetWithAreaDockwidgetOrientation(area Qt_DockWidgetArea,dockwidget *QDockWidget,orientation Qt_Orientation)  {
	q.Drv(305000,305107,unsafe.Pointer(&area),Native(dockwidget),unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::addToolBar(QString const&)
func (q *QMainWindow) AddToolBar(title string) *QToolBar {
	var __rv uintptr
	q.Drv(305000,305108,unsafe.Pointer(&title),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QToolBar{}
	_rp.SetDriver(__rv,381000,false)
	return _rp
}	
//QMainWindow::addToolBar(QToolBar*)
func (q *QMainWindow) AddToolBarWithToolbar(toolbar *QToolBar)  {
	q.Drv(305000,305109,Native(toolbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::addToolBar(Qt::ToolBarArea,QToolBar*)
func (q *QMainWindow) AddToolBarWithAreaToolbar(area Qt_ToolBarArea,toolbar *QToolBar)  {
	q.Drv(305000,305110,unsafe.Pointer(&area),Native(toolbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::addToolBarBreak()
func (q *QMainWindow) AddToolBarBreak()  {
	q.Drv(305000,305111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::addToolBarBreak(Qt::ToolBarArea)
func (q *QMainWindow) AddToolBarBreakWithArea(area Qt_ToolBarArea)  {
	q.Drv(305000,305112,unsafe.Pointer(&area),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::centralWidget()
func (q *QMainWindow) CentralWidget() *QWidget {
	var __rv uintptr
	q.Drv(305000,305113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMainWindow::contextMenuEvent(QContextMenuEvent*)
func (q *QMainWindow) ContextMenuEvent(event *QContextMenuEvent)  {
	q.Drv(305000,305114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::corner(Qt::Corner)
func (q *QMainWindow) Corner(corner Qt_Corner) Qt_DockWidgetArea {
	var __rv Qt_DockWidgetArea
	q.Drv(305000,305115,unsafe.Pointer(&corner),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::createPopupMenu()
func (q *QMainWindow) CreatePopupMenu() *QMenu {
	var __rv uintptr
	q.Drv(305000,305116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QMainWindow::dockOptions()
func (q *QMainWindow) DockOptions() QMainWindow_DockOption {
	var __rv QMainWindow_DockOption
	q.Drv(305000,305117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::dockWidgetArea(QDockWidget*)
func (q *QMainWindow) DockWidgetArea(dockwidget *QDockWidget) Qt_DockWidgetArea {
	var __rv Qt_DockWidgetArea
	q.Drv(305000,305118,Native(dockwidget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::documentMode()
func (q *QMainWindow) DocumentMode() bool {
	var __rv bool
	q.Drv(305000,305119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::event(QEvent*)
func (q *QMainWindow) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(305000,305120,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::iconSize()
func (q *QMainWindow) IconSize() *QSize {
	var __rv uintptr
	q.Drv(305000,305121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMainWindow::insertToolBar(QToolBar*,QToolBar*)
func (q *QMainWindow) InsertToolBar(before *QToolBar,toolbar *QToolBar)  {
	q.Drv(305000,305122,Native(before),Native(toolbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::insertToolBarBreak(QToolBar*)
func (q *QMainWindow) InsertToolBarBreak(before *QToolBar)  {
	q.Drv(305000,305123,Native(before),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::isAnimated()
func (q *QMainWindow) IsAnimated() bool {
	var __rv bool
	q.Drv(305000,305124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::isDockNestingEnabled()
func (q *QMainWindow) IsDockNestingEnabled() bool {
	var __rv bool
	q.Drv(305000,305125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::isSeparator(QPoint const&)
func (q *QMainWindow) IsSeparator(pos *QPoint) bool {
	var __rv bool
	q.Drv(305000,305126,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::menuBar()
func (q *QMainWindow) MenuBar() *QMenuBar {
	var __rv uintptr
	q.Drv(305000,305127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenuBar{}
	_rp.SetDriver(__rv,309000,false)
	return _rp
}	
//QMainWindow::menuWidget()
func (q *QMainWindow) MenuWidget() *QWidget {
	var __rv uintptr
	q.Drv(305000,305128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QMainWindow::removeDockWidget(QDockWidget*)
func (q *QMainWindow) RemoveDockWidget(dockwidget *QDockWidget)  {
	q.Drv(305000,305129,Native(dockwidget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::removeToolBar(QToolBar*)
func (q *QMainWindow) RemoveToolBar(toolbar *QToolBar)  {
	q.Drv(305000,305130,Native(toolbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::removeToolBarBreak(QToolBar*)
func (q *QMainWindow) RemoveToolBarBreak(before *QToolBar)  {
	q.Drv(305000,305131,Native(before),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::restoreDockWidget(QDockWidget*)
func (q *QMainWindow) RestoreDockWidget(dockwidget *QDockWidget) bool {
	var __rv bool
	q.Drv(305000,305132,Native(dockwidget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::restoreState(QByteArray const&)
func (q *QMainWindow) RestoreState(state []byte) bool {
	var __rv bool
	q.Drv(305000,305133,unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::restoreState(QByteArray const&,int)
func (q *QMainWindow) RestoreStateWithStateVersion(state []byte,version int) bool {
	var __rv bool
	q.Drv(305000,305134,unsafe.Pointer(&state),unsafe.Pointer(&version),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::saveState()
func (q *QMainWindow) SaveState() []byte {
	var __rv []byte
	q.Drv(305000,305135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::saveState(int)
func (q *QMainWindow) SaveStateWithVersion(version int) []byte {
	var __rv []byte
	q.Drv(305000,305136,unsafe.Pointer(&version),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::setAnimated(bool)
func (q *QMainWindow) SetAnimated(enabled bool)  {
	q.Drv(305000,305137,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setCentralWidget(QWidget*)
func (q *QMainWindow) SetCentralWidget(widget QWidgetInterface)  {
	q.Drv(305000,305138,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setCorner(Qt::Corner,Qt::DockWidgetArea)
func (q *QMainWindow) SetCorner(corner Qt_Corner,area Qt_DockWidgetArea)  {
	q.Drv(305000,305139,unsafe.Pointer(&corner),unsafe.Pointer(&area),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setDockNestingEnabled(bool)
func (q *QMainWindow) SetDockNestingEnabled(enabled bool)  {
	q.Drv(305000,305140,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setDockOptions(QFlags<QMainWindow::DockOption>)
func (q *QMainWindow) SetDockOptions(options QMainWindow_DockOption)  {
	q.Drv(305000,305141,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setDocumentMode(bool)
func (q *QMainWindow) SetDocumentMode(enabled bool)  {
	q.Drv(305000,305142,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setIconSize(QSize const&)
func (q *QMainWindow) SetIconSize(iconSize *QSize)  {
	q.Drv(305000,305143,Native(iconSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setMenuBar(QMenuBar*)
func (q *QMainWindow) SetMenuBar(menubar *QMenuBar)  {
	q.Drv(305000,305144,Native(menubar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setMenuWidget(QWidget*)
func (q *QMainWindow) SetMenuWidget(menubar QWidgetInterface)  {
	q.Drv(305000,305145,Native(menubar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setStatusBar(QStatusBar*)
func (q *QMainWindow) SetStatusBar(statusbar *QStatusBar)  {
	q.Drv(305000,305146,Native(statusbar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setTabPosition(QFlags<Qt::DockWidgetArea>,QTabWidget::TabPosition)
func (q *QMainWindow) SetTabPosition(areas Qt_DockWidgetArea,tabPosition QTabWidget_TabPosition)  {
	q.Drv(305000,305147,unsafe.Pointer(&areas),unsafe.Pointer(&tabPosition),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setTabShape(QTabWidget::TabShape)
func (q *QMainWindow) SetTabShape(tabShape QTabWidget_TabShape)  {
	q.Drv(305000,305148,unsafe.Pointer(&tabShape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setToolButtonStyle(Qt::ToolButtonStyle)
func (q *QMainWindow) SetToolButtonStyle(toolButtonStyle Qt_ToolButtonStyle)  {
	q.Drv(305000,305149,unsafe.Pointer(&toolButtonStyle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::setUnifiedTitleAndToolBarOnMac(bool)
func (q *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool)  {
	q.Drv(305000,305150,unsafe.Pointer(&set),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::splitDockWidget(QDockWidget*,QDockWidget*,Qt::Orientation)
func (q *QMainWindow) SplitDockWidget(after *QDockWidget,dockwidget *QDockWidget,orientation Qt_Orientation)  {
	q.Drv(305000,305151,Native(after),Native(dockwidget),unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::statusBar()
func (q *QMainWindow) StatusBar() *QStatusBar {
	var __rv uintptr
	q.Drv(305000,305152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStatusBar{}
	_rp.SetDriver(__rv,355000,false)
	return _rp
}	
//QMainWindow::tabPosition(Qt::DockWidgetArea)
func (q *QMainWindow) TabPosition(area Qt_DockWidgetArea) QTabWidget_TabPosition {
	var __rv QTabWidget_TabPosition
	q.Drv(305000,305153,unsafe.Pointer(&area),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::tabShape()
func (q *QMainWindow) TabShape() QTabWidget_TabShape {
	var __rv QTabWidget_TabShape
	q.Drv(305000,305154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::tabifiedDockWidgets(QDockWidget*)
func (q *QMainWindow) TabifiedDockWidgets(dockwidget *QDockWidget) []*QDockWidget {
	var __rv []*QDockWidget
	q.Drv(305000,305155,Native(dockwidget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::tabifyDockWidget(QDockWidget*,QDockWidget*)
func (q *QMainWindow) TabifyDockWidget(first *QDockWidget,second *QDockWidget)  {
	q.Drv(305000,305156,Native(first),Native(second),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMainWindow::toolBarArea(QToolBar*)
func (q *QMainWindow) ToolBarArea(toolbar *QToolBar) Qt_ToolBarArea {
	var __rv Qt_ToolBarArea
	q.Drv(305000,305157,Native(toolbar),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::toolBarBreak(QToolBar*)
func (q *QMainWindow) ToolBarBreak(toolbar *QToolBar) bool {
	var __rv bool
	q.Drv(305000,305158,Native(toolbar),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::toolButtonStyle()
func (q *QMainWindow) ToolButtonStyle() Qt_ToolButtonStyle {
	var __rv Qt_ToolButtonStyle
	q.Drv(305000,305159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMainWindow::unifiedTitleAndToolBarOnMac()
func (q *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	var __rv bool
	q.Drv(305000,305160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
