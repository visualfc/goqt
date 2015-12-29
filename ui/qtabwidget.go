// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTabWidget_TabShape - QTabWidget::TabShape
type QTabWidget_TabShape uint32
const (
	QTabWidget_Rounded QTabWidget_TabShape = 0
	QTabWidget_Triangular QTabWidget_TabShape = 1
)
//enum QTabWidget_TabPosition - QTabWidget::TabPosition
type QTabWidget_TabPosition uint32
const (
	QTabWidget_North QTabWidget_TabPosition = 0
	QTabWidget_South QTabWidget_TabPosition = 1
	QTabWidget_West QTabWidget_TabPosition = 2
	QTabWidget_East QTabWidget_TabPosition = 3
)
//struct QTabWidget : QTabWidget
type QTabWidget struct {
	QWidget
}
func (q *QTabWidget) OnCurrentChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(364000,364102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTabWidget) OnTabCloseRequested(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(364000,364103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTabWidget::QTabWidget()	
func NewTabWidget() *QTabWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),364000,364104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTabWidget{}
	_p.SetDriver(__rv,364000,false)
	return _p
} 
//QTabWidget::QTabWidget(QWidget*)	
func NewTabWidgetWithParent(parent QWidgetInterface) *QTabWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),364000,364105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTabWidget{}
	_p.SetDriver(__rv,364000,false)
	return _p
} 
//QTabWidget::addTab(QWidget*,QString const&)
func (q *QTabWidget) AddTabWithWidgetString(widget QWidgetInterface,value2 string) int {
	var __rv int
	q.Drv(364000,364106,Native(widget),unsafe.Pointer(&value2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::addTab(QWidget*,QIcon const&,QString const&)
func (q *QTabWidget) AddTabWithWidgetIconLabel(widget QWidgetInterface,icon *QIcon,label string) int {
	var __rv int
	q.Drv(364000,364107,Native(widget),Native(icon),unsafe.Pointer(&label),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::changeEvent(QEvent*)
func (q *QTabWidget) ChangeEvent(value *QEvent)  {
	q.Drv(364000,364108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::clear()
func (q *QTabWidget) Clear()  {
	q.Drv(364000,364109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::cornerWidget()
func (q *QTabWidget) CornerWidget() *QWidget {
	var __rv uintptr
	q.Drv(364000,364110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTabWidget::cornerWidget(Qt::Corner)
func (q *QTabWidget) CornerWidgetWithCorner(corner Qt_Corner) *QWidget {
	var __rv uintptr
	q.Drv(364000,364111,unsafe.Pointer(&corner),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTabWidget::count()
func (q *QTabWidget) Count() int {
	var __rv int
	q.Drv(364000,364112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::currentIndex()
func (q *QTabWidget) CurrentIndex() int {
	var __rv int
	q.Drv(364000,364113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::currentWidget()
func (q *QTabWidget) CurrentWidget() *QWidget {
	var __rv uintptr
	q.Drv(364000,364114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTabWidget::documentMode()
func (q *QTabWidget) DocumentMode() bool {
	var __rv bool
	q.Drv(364000,364115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::elideMode()
func (q *QTabWidget) ElideMode() Qt_TextElideMode {
	var __rv Qt_TextElideMode
	q.Drv(364000,364116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::event(QEvent*)
func (q *QTabWidget) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(364000,364117,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::iconSize()
func (q *QTabWidget) IconSize() *QSize {
	var __rv uintptr
	q.Drv(364000,364118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabWidget::indexOf(QWidget*)
func (q *QTabWidget) IndexOf(widget QWidgetInterface) int {
	var __rv int
	q.Drv(364000,364119,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::insertTab(int,QWidget*,QString const&)
func (q *QTabWidget) InsertTabWithIndexWidgetString(index int,widget QWidgetInterface,value2 string) int {
	var __rv int
	q.Drv(364000,364120,unsafe.Pointer(&index),Native(widget),unsafe.Pointer(&value2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::insertTab(int,QWidget*,QIcon const&,QString const&)
func (q *QTabWidget) InsertTabWithIndexWidgetIconLabel(index int,widget QWidgetInterface,icon *QIcon,label string) int {
	var __rv int
	q.Drv(364000,364121,unsafe.Pointer(&index),Native(widget),Native(icon),unsafe.Pointer(&label),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::isMovable()
func (q *QTabWidget) IsMovable() bool {
	var __rv bool
	q.Drv(364000,364122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::isTabEnabled(int)
func (q *QTabWidget) IsTabEnabled(index int) bool {
	var __rv bool
	q.Drv(364000,364123,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::keyPressEvent(QKeyEvent*)
func (q *QTabWidget) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(364000,364124,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::minimumSizeHint()
func (q *QTabWidget) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(364000,364125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabWidget::paintEvent(QPaintEvent*)
func (q *QTabWidget) PaintEvent(value *QPaintEvent)  {
	q.Drv(364000,364126,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::removeTab(int)
func (q *QTabWidget) RemoveTab(index int)  {
	q.Drv(364000,364127,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::resizeEvent(QResizeEvent*)
func (q *QTabWidget) ResizeEvent(value *QResizeEvent)  {
	q.Drv(364000,364128,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setCornerWidget(QWidget*)
func (q *QTabWidget) SetCornerWidget(w QWidgetInterface)  {
	q.Drv(364000,364129,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setCornerWidget(QWidget*,Qt::Corner)
func (q *QTabWidget) SetCornerWidgetWithWidgetCorner(w QWidgetInterface,corner Qt_Corner)  {
	q.Drv(364000,364130,Native(w),unsafe.Pointer(&corner),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setCurrentIndex(int)
func (q *QTabWidget) SetCurrentIndex(index int)  {
	q.Drv(364000,364131,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setCurrentWidget(QWidget*)
func (q *QTabWidget) SetCurrentWidget(widget QWidgetInterface)  {
	q.Drv(364000,364132,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setDocumentMode(bool)
func (q *QTabWidget) SetDocumentMode(set bool)  {
	q.Drv(364000,364133,unsafe.Pointer(&set),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setElideMode(Qt::TextElideMode)
func (q *QTabWidget) SetElideMode(value Qt_TextElideMode)  {
	q.Drv(364000,364134,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setIconSize(QSize const&)
func (q *QTabWidget) SetIconSize(size *QSize)  {
	q.Drv(364000,364135,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setMovable(bool)
func (q *QTabWidget) SetMovable(movable bool)  {
	q.Drv(364000,364136,unsafe.Pointer(&movable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabBar(QTabBar*)
func (q *QTabWidget) SetTabBar(value *QTabBar)  {
	q.Drv(364000,364137,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabEnabled(int,bool)
func (q *QTabWidget) SetTabEnabled(index int,value2 bool)  {
	q.Drv(364000,364138,unsafe.Pointer(&index),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabIcon(int,QIcon const&)
func (q *QTabWidget) SetTabIcon(index int,icon *QIcon)  {
	q.Drv(364000,364139,unsafe.Pointer(&index),Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabPosition(QTabWidget::TabPosition)
func (q *QTabWidget) SetTabPosition(value QTabWidget_TabPosition)  {
	q.Drv(364000,364140,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabShape(QTabWidget::TabShape)
func (q *QTabWidget) SetTabShape(s QTabWidget_TabShape)  {
	q.Drv(364000,364141,unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabText(int,QString const&)
func (q *QTabWidget) SetTabText(index int,value2 string)  {
	q.Drv(364000,364142,unsafe.Pointer(&index),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabToolTip(int,QString const&)
func (q *QTabWidget) SetTabToolTip(index int,tip string)  {
	q.Drv(364000,364143,unsafe.Pointer(&index),unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabWhatsThis(int,QString const&)
func (q *QTabWidget) SetTabWhatsThis(index int,text string)  {
	q.Drv(364000,364144,unsafe.Pointer(&index),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setTabsClosable(bool)
func (q *QTabWidget) SetTabsClosable(closeable bool)  {
	q.Drv(364000,364145,unsafe.Pointer(&closeable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::setUsesScrollButtons(bool)
func (q *QTabWidget) SetUsesScrollButtons(useButtons bool)  {
	q.Drv(364000,364146,unsafe.Pointer(&useButtons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::showEvent(QShowEvent*)
func (q *QTabWidget) ShowEvent(value *QShowEvent)  {
	q.Drv(364000,364147,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::sizeHint()
func (q *QTabWidget) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(364000,364148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QTabWidget::tabBar()
func (q *QTabWidget) TabBar() *QTabBar {
	var __rv uintptr
	q.Drv(364000,364149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTabBar{}
	_rp.SetDriver(__rv,363000,false)
	return _rp
}	
//QTabWidget::tabIcon(int)
func (q *QTabWidget) TabIcon(index int) *QIcon {
	var __rv uintptr
	q.Drv(364000,364150,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QTabWidget::tabInserted(int)
func (q *QTabWidget) TabInserted(index int)  {
	q.Drv(364000,364151,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::tabPosition()
func (q *QTabWidget) TabPosition() QTabWidget_TabPosition {
	var __rv QTabWidget_TabPosition
	q.Drv(364000,364152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::tabRemoved(int)
func (q *QTabWidget) TabRemoved(index int)  {
	q.Drv(364000,364153,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTabWidget::tabShape()
func (q *QTabWidget) TabShape() QTabWidget_TabShape {
	var __rv QTabWidget_TabShape
	q.Drv(364000,364154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::tabText(int)
func (q *QTabWidget) TabText(index int) string {
	var __rv string
	q.Drv(364000,364155,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::tabToolTip(int)
func (q *QTabWidget) TabToolTip(index int) string {
	var __rv string
	q.Drv(364000,364156,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::tabWhatsThis(int)
func (q *QTabWidget) TabWhatsThis(index int) string {
	var __rv string
	q.Drv(364000,364157,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::tabsClosable()
func (q *QTabWidget) TabsClosable() bool {
	var __rv bool
	q.Drv(364000,364158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::usesScrollButtons()
func (q *QTabWidget) UsesScrollButtons() bool {
	var __rv bool
	q.Drv(364000,364159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabWidget::widget(int)
func (q *QTabWidget) Widget(index int) *QWidget {
	var __rv uintptr
	q.Drv(364000,364160,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
