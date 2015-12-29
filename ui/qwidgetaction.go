// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWidgetAction : QWidgetAction
type QWidgetAction struct {
	QAction
}
//QWidgetAction::QWidgetAction(QObject*)	
func NewWidgetAction(parent QObjectInterface) *QWidgetAction {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),396000,396102,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWidgetAction{}
	_p.SetDriver(__rv,396000,false)
	return _p
} 
//QWidgetAction::createWidget(QWidget*)
func (q *QWidgetAction) CreateWidget(parent QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(396000,396103,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidgetAction::createdWidgets()
func (q *QWidgetAction) CreatedWidgets() []*QWidget {
	var __rv []*QWidget
	q.Drv(396000,396104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetAction::defaultWidget()
func (q *QWidgetAction) DefaultWidget() *QWidget {
	var __rv uintptr
	q.Drv(396000,396105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidgetAction::deleteWidget(QWidget*)
func (q *QWidgetAction) DeleteWidget(widget QWidgetInterface)  {
	q.Drv(396000,396106,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidgetAction::event(QEvent*)
func (q *QWidgetAction) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(396000,396107,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetAction::eventFilter(QObject*,QEvent*)
func (q *QWidgetAction) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(396000,396108,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWidgetAction::releaseWidget(QWidget*)
func (q *QWidgetAction) ReleaseWidget(widget QWidgetInterface)  {
	q.Drv(396000,396109,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWidgetAction::requestWidget(QWidget*)
func (q *QWidgetAction) RequestWidget(parent QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(396000,396110,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWidgetAction::setDefaultWidget(QWidget*)
func (q *QWidgetAction) SetDefaultWidget(w QWidgetInterface)  {
	q.Drv(396000,396111,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
