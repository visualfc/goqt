// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWhatsThis : QWhatsThis
type QWhatsThis struct {
	BaseDrv
}
//QWhatsThis::createAction()	
func QWhatsThisCreateAction() *QAction {
	var __rv uintptr
	DirectQtDrv(nil,187000,187102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QWhatsThis::createAction()
func (q *QWhatsThis) CreateAction() *QAction {
	var __rv uintptr
	q.Drv(187000,187102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QWhatsThis::createAction(QObject*)	
func QWhatsThisCreateActionWithParent(parent QObjectInterface) *QAction {
	var __rv uintptr
	DirectQtDrv(nil,187000,187103,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QWhatsThis::createAction(QObject*)
func (q *QWhatsThis) CreateActionWithParent(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(187000,187103,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QWhatsThis::enterWhatsThisMode()	
func QWhatsThisEnterWhatsThisMode()  {
	DirectQtDrv(nil,187000,187104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::enterWhatsThisMode()
func (q *QWhatsThis) EnterWhatsThisMode()  {
	q.Drv(187000,187104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::hideText()	
func QWhatsThisHideText()  {
	DirectQtDrv(nil,187000,187105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::hideText()
func (q *QWhatsThis) HideText()  {
	q.Drv(187000,187105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::inWhatsThisMode()	
func QWhatsThisInWhatsThisMode() bool {
	var __rv bool
	DirectQtDrv(nil,187000,187106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWhatsThis::inWhatsThisMode()
func (q *QWhatsThis) InWhatsThisMode() bool {
	var __rv bool
	q.Drv(187000,187106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWhatsThis::leaveWhatsThisMode()	
func QWhatsThisLeaveWhatsThisMode()  {
	DirectQtDrv(nil,187000,187107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::leaveWhatsThisMode()
func (q *QWhatsThis) LeaveWhatsThisMode()  {
	q.Drv(187000,187107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::showText(QPoint const&,QString const&,QWidget*)	
func QWhatsThisShowText(pos *QPoint,text string,w QWidgetInterface)  {
	DirectQtDrv(nil,187000,187108,Native(pos),unsafe.Pointer(&text),Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWhatsThis::showText(QPoint const&,QString const&,QWidget*)
func (q *QWhatsThis) ShowText(pos *QPoint,text string,w QWidgetInterface)  {
	q.Drv(187000,187108,Native(pos),unsafe.Pointer(&text),Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
