// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSizeGrip : QSizeGrip
type QSizeGrip struct {
	QWidget
}
//QSizeGrip::QSizeGrip(QWidget*)	
func NewSizeGrip(parent QWidgetInterface) *QSizeGrip {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),343000,343102,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizeGrip{}
	_p.SetDriver(__rv,343000,false)
	return _p
} 
//QSizeGrip::event(QEvent*)
func (q *QSizeGrip) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(343000,343103,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeGrip::eventFilter(QObject*,QEvent*)
func (q *QSizeGrip) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(343000,343104,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeGrip::hideEvent(QHideEvent*)
func (q *QSizeGrip) HideEvent(hideEvent *QHideEvent)  {
	q.Drv(343000,343105,Native(hideEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::mouseMoveEvent(QMouseEvent*)
func (q *QSizeGrip) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(343000,343106,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::mousePressEvent(QMouseEvent*)
func (q *QSizeGrip) MousePressEvent(value *QMouseEvent)  {
	q.Drv(343000,343107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::mouseReleaseEvent(QMouseEvent*)
func (q *QSizeGrip) MouseReleaseEvent(mouseEvent *QMouseEvent)  {
	q.Drv(343000,343108,Native(mouseEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::moveEvent(QMoveEvent*)
func (q *QSizeGrip) MoveEvent(moveEvent *QMoveEvent)  {
	q.Drv(343000,343109,Native(moveEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::paintEvent(QPaintEvent*)
func (q *QSizeGrip) PaintEvent(value *QPaintEvent)  {
	q.Drv(343000,343110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::setVisible(bool)
func (q *QSizeGrip) SetVisible(value bool)  {
	q.Drv(343000,343111,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::showEvent(QShowEvent*)
func (q *QSizeGrip) ShowEvent(showEvent *QShowEvent)  {
	q.Drv(343000,343112,Native(showEvent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeGrip::sizeHint()
func (q *QSizeGrip) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(343000,343113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
