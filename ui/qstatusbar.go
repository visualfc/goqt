// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStatusBar : QStatusBar
type QStatusBar struct {
	QWidget
}
func (q *QStatusBar) OnMessageChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(355000,355102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QStatusBar::QStatusBar()	
func NewStatusBar() *QStatusBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),355000,355103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStatusBar{}
	_p.SetDriver(__rv,355000,false)
	return _p
} 
//QStatusBar::QStatusBar(QWidget*)	
func NewStatusBarWithParent(parent QWidgetInterface) *QStatusBar {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),355000,355104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStatusBar{}
	_p.SetDriver(__rv,355000,false)
	return _p
} 
//QStatusBar::addPermanentWidget(QWidget*)
func (q *QStatusBar) AddPermanentWidget(widget QWidgetInterface)  {
	q.Drv(355000,355105,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::addPermanentWidget(QWidget*,int)
func (q *QStatusBar) AddPermanentWidgetWithWidgetStretch(widget QWidgetInterface,stretch int)  {
	q.Drv(355000,355106,Native(widget),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::addWidget(QWidget*)
func (q *QStatusBar) AddWidget(widget QWidgetInterface)  {
	q.Drv(355000,355107,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::addWidget(QWidget*,int)
func (q *QStatusBar) AddWidgetWithWidgetStretch(widget QWidgetInterface,stretch int)  {
	q.Drv(355000,355108,Native(widget),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::clearMessage()
func (q *QStatusBar) ClearMessage()  {
	q.Drv(355000,355109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::currentMessage()
func (q *QStatusBar) CurrentMessage() string {
	var __rv string
	q.Drv(355000,355110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStatusBar::event(QEvent*)
func (q *QStatusBar) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(355000,355111,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStatusBar::hideOrShow()
func (q *QStatusBar) HideOrShow()  {
	q.Drv(355000,355112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::insertPermanentWidget(int,QWidget*,int)
func (q *QStatusBar) InsertPermanentWidget(index int,widget QWidgetInterface,stretch int) int {
	var __rv int
	q.Drv(355000,355113,unsafe.Pointer(&index),Native(widget),unsafe.Pointer(&stretch),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStatusBar::insertWidget(int,QWidget*,int)
func (q *QStatusBar) InsertWidget(index int,widget QWidgetInterface,stretch int) int {
	var __rv int
	q.Drv(355000,355114,unsafe.Pointer(&index),Native(widget),unsafe.Pointer(&stretch),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStatusBar::isSizeGripEnabled()
func (q *QStatusBar) IsSizeGripEnabled() bool {
	var __rv bool
	q.Drv(355000,355115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStatusBar::paintEvent(QPaintEvent*)
func (q *QStatusBar) PaintEvent(value *QPaintEvent)  {
	q.Drv(355000,355116,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::reformat()
func (q *QStatusBar) Reformat()  {
	q.Drv(355000,355117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::removeWidget(QWidget*)
func (q *QStatusBar) RemoveWidget(widget QWidgetInterface)  {
	q.Drv(355000,355118,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::resizeEvent(QResizeEvent*)
func (q *QStatusBar) ResizeEvent(value *QResizeEvent)  {
	q.Drv(355000,355119,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::setSizeGripEnabled(bool)
func (q *QStatusBar) SetSizeGripEnabled(value bool)  {
	q.Drv(355000,355120,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::showEvent(QShowEvent*)
func (q *QStatusBar) ShowEvent(value *QShowEvent)  {
	q.Drv(355000,355121,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::showMessage(QString const&)
func (q *QStatusBar) ShowMessage(text string)  {
	q.Drv(355000,355122,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStatusBar::showMessage(QString const&,int)
func (q *QStatusBar) ShowMessageWithTextTimeout(text string,timeout int)  {
	q.Drv(355000,355123,unsafe.Pointer(&text),unsafe.Pointer(&timeout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
