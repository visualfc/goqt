// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUndoView : QUndoView
type QUndoView struct {
	QListView
}
//QUndoView::QUndoView()	
func NewUndoView() *QUndoView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),390000,390102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoView{}
	_p.SetDriver(__rv,390000,false)
	return _p
} 
//QUndoView::QUndoView(QWidget*)	
func NewUndoViewWithParent(parent QWidgetInterface) *QUndoView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),390000,390103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoView{}
	_p.SetDriver(__rv,390000,false)
	return _p
} 
//QUndoView::QUndoView(QUndoGroup*,QWidget*)	
func NewUndoViewWithGroupParent(group *QUndoGroup,parent QWidgetInterface) *QUndoView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),390000,390104,Native(group),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoView{}
	_p.SetDriver(__rv,390000,false)
	return _p
} 
//QUndoView::QUndoView(QUndoStack*,QWidget*)	
func NewUndoViewWithStackParent(stack *QUndoStack,parent QWidgetInterface) *QUndoView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),390000,390105,Native(stack),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoView{}
	_p.SetDriver(__rv,390000,false)
	return _p
} 
//QUndoView::cleanIcon()
func (q *QUndoView) CleanIcon() *QIcon {
	var __rv uintptr
	q.Drv(390000,390106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QUndoView::emptyLabel()
func (q *QUndoView) EmptyLabel() string {
	var __rv string
	q.Drv(390000,390107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoView::group()
func (q *QUndoView) Group() *QUndoGroup {
	var __rv uintptr
	q.Drv(390000,390108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUndoGroup{}
	_rp.SetDriver(__rv,388000,false)
	return _rp
}	
//QUndoView::setCleanIcon(QIcon const&)
func (q *QUndoView) SetCleanIcon(icon *QIcon)  {
	q.Drv(390000,390109,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoView::setEmptyLabel(QString const&)
func (q *QUndoView) SetEmptyLabel(label string)  {
	q.Drv(390000,390110,unsafe.Pointer(&label),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoView::setGroup(QUndoGroup*)
func (q *QUndoView) SetGroup(group *QUndoGroup)  {
	q.Drv(390000,390111,Native(group),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoView::setStack(QUndoStack*)
func (q *QUndoView) SetStack(stack *QUndoStack)  {
	q.Drv(390000,390112,Native(stack),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoView::stack()
func (q *QUndoView) Stack() *QUndoStack {
	var __rv uintptr
	q.Drv(390000,390113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUndoStack{}
	_rp.SetDriver(__rv,389000,false)
	return _rp
}	
