// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPushButton : QPushButton
type QPushButton struct {
	QAbstractButton
}
//QPushButton::QPushButton()	
func NewPushButton() *QPushButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),331000,331102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPushButton{}
	_p.SetDriver(__rv,331000,false)
	return _p
} 
//QPushButton::QPushButton(QWidget*)	
func NewPushButtonWithParent(parent QWidgetInterface) *QPushButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),331000,331103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPushButton{}
	_p.SetDriver(__rv,331000,false)
	return _p
} 
//QPushButton::QPushButton(QString const&,QWidget*)	
func NewPushButtonWithTextParent(text string,parent QWidgetInterface) *QPushButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),331000,331104,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPushButton{}
	_p.SetDriver(__rv,331000,false)
	return _p
} 
//QPushButton::QPushButton(QIcon const&,QString const&,QWidget*)	
func NewPushButtonWithIconTextParent(icon *QIcon,text string,parent QWidgetInterface) *QPushButton {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),331000,331105,Native(icon),unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPushButton{}
	_p.SetDriver(__rv,331000,false)
	return _p
} 
//QPushButton::autoDefault()
func (q *QPushButton) AutoDefault() bool {
	var __rv bool
	q.Drv(331000,331106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPushButton::event(QEvent*)
func (q *QPushButton) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(331000,331107,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPushButton::focusInEvent(QFocusEvent*)
func (q *QPushButton) FocusInEvent(value *QFocusEvent)  {
	q.Drv(331000,331108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::focusOutEvent(QFocusEvent*)
func (q *QPushButton) FocusOutEvent(value *QFocusEvent)  {
	q.Drv(331000,331109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::isDefault()
func (q *QPushButton) IsDefault() bool {
	var __rv bool
	q.Drv(331000,331110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPushButton::isFlat()
func (q *QPushButton) IsFlat() bool {
	var __rv bool
	q.Drv(331000,331111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPushButton::keyPressEvent(QKeyEvent*)
func (q *QPushButton) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(331000,331112,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::menu()
func (q *QPushButton) Menu() *QMenu {
	var __rv uintptr
	q.Drv(331000,331113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QPushButton::minimumSizeHint()
func (q *QPushButton) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(331000,331114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QPushButton::paintEvent(QPaintEvent*)
func (q *QPushButton) PaintEvent(value *QPaintEvent)  {
	q.Drv(331000,331115,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::setAutoDefault(bool)
func (q *QPushButton) SetAutoDefault(value bool)  {
	q.Drv(331000,331116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::setDefault(bool)
func (q *QPushButton) SetDefault(value bool)  {
	q.Drv(331000,331117,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::setFlat(bool)
func (q *QPushButton) SetFlat(value bool)  {
	q.Drv(331000,331118,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::setMenu(QMenu*)
func (q *QPushButton) SetMenu(menu *QMenu)  {
	q.Drv(331000,331119,Native(menu),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::showMenu()
func (q *QPushButton) ShowMenu()  {
	q.Drv(331000,331120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPushButton::sizeHint()
func (q *QPushButton) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(331000,331121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
