// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QActionGroup : QActionGroup
type QActionGroup struct {
	QObject
}
func (q *QActionGroup) OnHovered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(208000,208102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QActionGroup) OnSelected(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(208000,208103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QActionGroup) OnTriggered(fn func(*QAction)) uintptr {
	var __rv uintptr
	q.Drv(208000,208104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QActionGroup::QActionGroup(QObject*)	
func NewActionGroup(parent QObjectInterface) *QActionGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),208000,208105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QActionGroup{}
	_p.SetDriver(__rv,208000,false)
	return _p
} 
//QActionGroup::actions()
func (q *QActionGroup) Actions() []*QAction {
	var __rv []*QAction
	q.Drv(208000,208106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QActionGroup::addAction(QAction*)
func (q *QActionGroup) AddAction(a *QAction) *QAction {
	var __rv uintptr
	q.Drv(208000,208107,Native(a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QActionGroup::addAction(QString const&)
func (q *QActionGroup) AddActionWithText(text string) *QAction {
	var __rv uintptr
	q.Drv(208000,208108,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QActionGroup::addAction(QIcon const&,QString const&)
func (q *QActionGroup) AddActionWithIconText(icon *QIcon,text string) *QAction {
	var __rv uintptr
	q.Drv(208000,208109,Native(icon),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QActionGroup::checkedAction()
func (q *QActionGroup) CheckedAction() *QAction {
	var __rv uintptr
	q.Drv(208000,208110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QActionGroup::isEnabled()
func (q *QActionGroup) IsEnabled() bool {
	var __rv bool
	q.Drv(208000,208111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QActionGroup::isExclusive()
func (q *QActionGroup) IsExclusive() bool {
	var __rv bool
	q.Drv(208000,208112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QActionGroup::isVisible()
func (q *QActionGroup) IsVisible() bool {
	var __rv bool
	q.Drv(208000,208113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QActionGroup::removeAction(QAction*)
func (q *QActionGroup) RemoveAction(a *QAction)  {
	q.Drv(208000,208114,Native(a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QActionGroup::setDisabled(bool)
func (q *QActionGroup) SetDisabled(b bool)  {
	q.Drv(208000,208115,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QActionGroup::setEnabled(bool)
func (q *QActionGroup) SetEnabled(value bool)  {
	q.Drv(208000,208116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QActionGroup::setExclusive(bool)
func (q *QActionGroup) SetExclusive(value bool)  {
	q.Drv(208000,208117,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QActionGroup::setVisible(bool)
func (q *QActionGroup) SetVisible(value bool)  {
	q.Drv(208000,208118,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
