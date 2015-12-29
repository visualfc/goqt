// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFocusFrame : QFocusFrame
type QFocusFrame struct {
	QWidget
}
//QFocusFrame::QFocusFrame()	
func NewFocusFrame() *QFocusFrame {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),241000,241102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFocusFrame{}
	_p.SetDriver(__rv,241000,false)
	return _p
} 
//QFocusFrame::QFocusFrame(QWidget*)	
func NewFocusFrameWithParent(parent QWidgetInterface) *QFocusFrame {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),241000,241103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFocusFrame{}
	_p.SetDriver(__rv,241000,false)
	return _p
} 
//QFocusFrame::event(QEvent*)
func (q *QFocusFrame) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(241000,241104,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFocusFrame::eventFilter(QObject*,QEvent*)
func (q *QFocusFrame) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(241000,241105,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFocusFrame::paintEvent(QPaintEvent*)
func (q *QFocusFrame) PaintEvent(value *QPaintEvent)  {
	q.Drv(241000,241106,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFocusFrame::setWidget(QWidget*)
func (q *QFocusFrame) SetWidget(widget QWidgetInterface)  {
	q.Drv(241000,241107,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFocusFrame::widget()
func (q *QFocusFrame) Widget() *QWidget {
	var __rv uintptr
	q.Drv(241000,241108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
