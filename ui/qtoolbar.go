// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QToolBar : QToolBar
type QToolBar struct {
	QWidget
}
func (q *QToolBar) OnMovableChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(381000,381102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnToolButtonStyleChanged(fn func(Qt_ToolButtonStyle)) uintptr {
	var __rv uintptr
	q.Drv(381000,381103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnActionTriggered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(381000,381104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnTopLevelChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(381000,381105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnVisibilityChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(381000,381106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnOrientationChanged(fn func(Qt_Orientation)) uintptr {
	var __rv uintptr
	q.Drv(381000,381107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnAllowedAreasChanged(fn func(Qt_ToolBarArea)) uintptr {
	var __rv uintptr
	q.Drv(381000,381108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QToolBar) OnIconSizeChanged(fn func(*QSize)) uintptr {
	var __rv uintptr
	q.Drv(381000,381109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QToolBar::QToolBar()	
func NewToolBar() *QToolBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),381000,381110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBar{}
	_p.SetDriver(__rv,381000,false)
	return _p
} 
//QToolBar::QToolBar(QWidget*)	
func NewToolBarWithParent(parent QWidgetInterface) *QToolBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),381000,381111,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBar{}
	_p.SetDriver(__rv,381000,false)
	return _p
} 
//QToolBar::QToolBar(QString const&,QWidget*)	
func NewToolBarWithTitleParent(title string,parent QWidgetInterface) *QToolBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),381000,381112,unsafe.Pointer(&title),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QToolBar{}
	_p.SetDriver(__rv,381000,false)
	return _p
} 
//QToolBar::actionAt(QPoint const&)
func (q *QToolBar) ActionAt(p *QPoint) *QAction {
	var __rv uintptr
	q.Drv(381000,381113,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::actionAt(int,int)
func (q *QToolBar) ActionAtWithXY(x int,y int) *QAction {
	var __rv uintptr
	q.Drv(381000,381114,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::actionEvent(QActionEvent*)
func (q *QToolBar) ActionEvent(event *QActionEvent)  {
	q.Drv(381000,381115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::actionGeometry(QAction*)
func (q *QToolBar) ActionGeometry(action *QAction) *QRect {
	var __rv uintptr
	q.Drv(381000,381116,Native(action),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QToolBar::addAction(QAction*)
func (q *QToolBar) AddAction(action *QAction)  {
	q.Drv(381000,381117,Native(action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::addAction(QString const&)
func (q *QToolBar) AddActionWithText(text string) *QAction {
	var __rv uintptr
	q.Drv(381000,381118,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::addAction(QIcon const&,QString const&)
func (q *QToolBar) AddActionWithIconText(icon *QIcon,text string) *QAction {
	var __rv uintptr
	q.Drv(381000,381119,Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::addAction(QString const&,QObject const*,char const*)
func (q *QToolBar) AddActionWithTextObjectMember(text string,receiver QObjectInterface,member string) *QAction {
	var __rv uintptr
	q.Drv(381000,381120,unsafe.Pointer(&text),Native(receiver),unsafe.Pointer(&member),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::addAction(QIcon const&,QString const&,QObject const*,char const*)
func (q *QToolBar) AddActionWithIconTextObjectMember(icon *QIcon,text string,receiver QObjectInterface,member string) *QAction {
	var __rv uintptr
	q.Drv(381000,381121,Native(icon),unsafe.Pointer(&text),Native(receiver),unsafe.Pointer(&member),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::addSeparator()
func (q *QToolBar) AddSeparator() *QAction {
	var __rv uintptr
	q.Drv(381000,381122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::addWidget(QWidget*)
func (q *QToolBar) AddWidget(widget QWidgetInterface) *QAction {
	var __rv uintptr
	q.Drv(381000,381123,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::allowedAreas()
func (q *QToolBar) AllowedAreas() Qt_ToolBarArea {
	var __rv Qt_ToolBarArea
	q.Drv(381000,381124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::changeEvent(QEvent*)
func (q *QToolBar) ChangeEvent(event *QEvent)  {
	q.Drv(381000,381125,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::childEvent(QChildEvent*)
func (q *QToolBar) ChildEvent(event *QChildEvent)  {
	q.Drv(381000,381126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::clear()
func (q *QToolBar) Clear()  {
	q.Drv(381000,381127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::event(QEvent*)
func (q *QToolBar) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(381000,381128,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::iconSize()
func (q *QToolBar) IconSize() *QSize {
	var __rv uintptr
	q.Drv(381000,381129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QToolBar::insertSeparator(QAction*)
func (q *QToolBar) InsertSeparator(before *QAction) *QAction {
	var __rv uintptr
	q.Drv(381000,381130,Native(before),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::insertWidget(QAction*,QWidget*)
func (q *QToolBar) InsertWidget(before *QAction,widget QWidgetInterface) *QAction {
	var __rv uintptr
	q.Drv(381000,381131,Native(before),Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::isAreaAllowed(Qt::ToolBarArea)
func (q *QToolBar) IsAreaAllowed(area Qt_ToolBarArea) bool {
	var __rv bool
	q.Drv(381000,381132,unsafe.Pointer(&area),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::isFloatable()
func (q *QToolBar) IsFloatable() bool {
	var __rv bool
	q.Drv(381000,381133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::isFloating()
func (q *QToolBar) IsFloating() bool {
	var __rv bool
	q.Drv(381000,381134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::isMovable()
func (q *QToolBar) IsMovable() bool {
	var __rv bool
	q.Drv(381000,381135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::orientation()
func (q *QToolBar) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(381000,381136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::paintEvent(QPaintEvent*)
func (q *QToolBar) PaintEvent(event *QPaintEvent)  {
	q.Drv(381000,381137,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::resizeEvent(QResizeEvent*)
func (q *QToolBar) ResizeEvent(event *QResizeEvent)  {
	q.Drv(381000,381138,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setAllowedAreas(QFlags<Qt::ToolBarArea>)
func (q *QToolBar) SetAllowedAreas(areas Qt_ToolBarArea)  {
	q.Drv(381000,381139,unsafe.Pointer(&areas),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setFloatable(bool)
func (q *QToolBar) SetFloatable(floatable bool)  {
	q.Drv(381000,381140,unsafe.Pointer(&floatable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setIconSize(QSize const&)
func (q *QToolBar) SetIconSize(iconSize *QSize)  {
	q.Drv(381000,381141,Native(iconSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setMovable(bool)
func (q *QToolBar) SetMovable(movable bool)  {
	q.Drv(381000,381142,unsafe.Pointer(&movable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setOrientation(Qt::Orientation)
func (q *QToolBar) SetOrientation(orientation Qt_Orientation)  {
	q.Drv(381000,381143,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::setToolButtonStyle(Qt::ToolButtonStyle)
func (q *QToolBar) SetToolButtonStyle(toolButtonStyle Qt_ToolButtonStyle)  {
	q.Drv(381000,381144,unsafe.Pointer(&toolButtonStyle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolBar::toggleViewAction()
func (q *QToolBar) ToggleViewAction() *QAction {
	var __rv uintptr
	q.Drv(381000,381145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QToolBar::toolButtonStyle()
func (q *QToolBar) ToolButtonStyle() Qt_ToolButtonStyle {
	var __rv Qt_ToolButtonStyle
	q.Drv(381000,381146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolBar::widgetForAction(QAction*)
func (q *QToolBar) WidgetForAction(action *QAction) *QWidget {
	var __rv uintptr
	q.Drv(381000,381147,Native(action),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
