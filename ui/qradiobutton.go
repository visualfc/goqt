// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QRadioButton : QRadioButton
type QRadioButton struct {
	QAbstractButton
}
//QRadioButton::QRadioButton()	
func NewRadioButton() *QRadioButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),332000,332102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadioButton{}
	_p.SetDriver(__rv,332000,false)
	return _p
} 
//QRadioButton::QRadioButton(QWidget*)	
func NewRadioButtonWithParent(parent QWidgetInterface) *QRadioButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),332000,332103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadioButton{}
	_p.SetDriver(__rv,332000,false)
	return _p
} 
//QRadioButton::QRadioButton(QString const&,QWidget*)	
func NewRadioButtonWithTextParent(text string,parent QWidgetInterface) *QRadioButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),332000,332104,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadioButton{}
	_p.SetDriver(__rv,332000,false)
	return _p
} 
//QRadioButton::event(QEvent*)
func (q *QRadioButton) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(332000,332105,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRadioButton::hitButton(QPoint const&)
func (q *QRadioButton) HitButton(value *QPoint) bool {
	var __rv bool
	q.Drv(332000,332106,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRadioButton::mouseMoveEvent(QMouseEvent*)
func (q *QRadioButton) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(332000,332107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadioButton::paintEvent(QPaintEvent*)
func (q *QRadioButton) PaintEvent(value *QPaintEvent)  {
	q.Drv(332000,332108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadioButton::sizeHint()
func (q *QRadioButton) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(332000,332109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
