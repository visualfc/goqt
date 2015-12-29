// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QCheckBox : QCheckBox
type QCheckBox struct {
	QAbstractButton
}
func (q *QCheckBox) OnStateChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(213000,213102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QCheckBox::QCheckBox()	
func NewCheckBox() *QCheckBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),213000,213103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCheckBox{}
	_p.SetDriver(__rv,213000,false)
	return _p
} 
//QCheckBox::QCheckBox(QWidget*)	
func NewCheckBoxWithParent(parent QWidgetInterface) *QCheckBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),213000,213104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCheckBox{}
	_p.SetDriver(__rv,213000,false)
	return _p
} 
//QCheckBox::QCheckBox(QString const&,QWidget*)	
func NewCheckBoxWithTextParent(text string,parent QWidgetInterface) *QCheckBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),213000,213105,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCheckBox{}
	_p.SetDriver(__rv,213000,false)
	return _p
} 
//QCheckBox::checkState()
func (q *QCheckBox) CheckState() Qt_CheckState {
	var __rv Qt_CheckState
	q.Drv(213000,213106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCheckBox::checkStateSet()
func (q *QCheckBox) CheckStateSet()  {
	q.Drv(213000,213107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::event(QEvent*)
func (q *QCheckBox) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(213000,213108,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCheckBox::hitButton(QPoint const&)
func (q *QCheckBox) HitButton(pos *QPoint) bool {
	var __rv bool
	q.Drv(213000,213109,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCheckBox::isTristate()
func (q *QCheckBox) IsTristate() bool {
	var __rv bool
	q.Drv(213000,213110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCheckBox::mouseMoveEvent(QMouseEvent*)
func (q *QCheckBox) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(213000,213111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::nextCheckState()
func (q *QCheckBox) NextCheckState()  {
	q.Drv(213000,213112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::paintEvent(QPaintEvent*)
func (q *QCheckBox) PaintEvent(value *QPaintEvent)  {
	q.Drv(213000,213113,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::setCheckState(Qt::CheckState)
func (q *QCheckBox) SetCheckState(state Qt_CheckState)  {
	q.Drv(213000,213114,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::setTristate(bool)
func (q *QCheckBox) SetTristate(y bool)  {
	q.Drv(213000,213115,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCheckBox::sizeHint()
func (q *QCheckBox) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(213000,213116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
