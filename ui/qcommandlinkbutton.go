// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QCommandLinkButton : QCommandLinkButton
type QCommandLinkButton struct {
	QPushButton
}
//QCommandLinkButton::QCommandLinkButton()	
func NewCommandLinkButton() *QCommandLinkButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),218000,218102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCommandLinkButton{}
	_p.SetDriver(__rv,218000,false)
	return _p
} 
//QCommandLinkButton::QCommandLinkButton(QWidget*)	
func NewCommandLinkButtonWithParent(parent QWidgetInterface) *QCommandLinkButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),218000,218103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCommandLinkButton{}
	_p.SetDriver(__rv,218000,false)
	return _p
} 
//QCommandLinkButton::QCommandLinkButton(QString const&,QWidget*)	
func NewCommandLinkButtonWithTextParent(text string,parent QWidgetInterface) *QCommandLinkButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),218000,218104,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCommandLinkButton{}
	_p.SetDriver(__rv,218000,false)
	return _p
} 
//QCommandLinkButton::QCommandLinkButton(QString const&,QString const&,QWidget*)	
func NewCommandLinkButtonWithTextDescriptionParent(text string,description string,parent QWidgetInterface) *QCommandLinkButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),218000,218105,unsafe.Pointer(&text),unsafe.Pointer(&description),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCommandLinkButton{}
	_p.SetDriver(__rv,218000,false)
	return _p
} 
//QCommandLinkButton::description()
func (q *QCommandLinkButton) Description() string {
	var __rv string
	q.Drv(218000,218106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCommandLinkButton::event(QEvent*)
func (q *QCommandLinkButton) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(218000,218107,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCommandLinkButton::heightForWidth(int)
func (q *QCommandLinkButton) HeightForWidth(value int) int {
	var __rv int
	q.Drv(218000,218108,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCommandLinkButton::minimumSizeHint()
func (q *QCommandLinkButton) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(218000,218109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QCommandLinkButton::paintEvent(QPaintEvent*)
func (q *QCommandLinkButton) PaintEvent(value *QPaintEvent)  {
	q.Drv(218000,218110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommandLinkButton::setDescription(QString const&)
func (q *QCommandLinkButton) SetDescription(description string)  {
	q.Drv(218000,218111,unsafe.Pointer(&description),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommandLinkButton::sizeHint()
func (q *QCommandLinkButton) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(218000,218112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
