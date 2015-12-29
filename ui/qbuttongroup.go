// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QButtonGroup : QButtonGroup
type QButtonGroup struct {
	QObject
}
func (q *QButtonGroup) OnButtonClicked(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(211000,211102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QButtonGroup) OnButtonClickedWithAbstractbutton(fn func(*QAbstractButton)) uintptr {
	var __rv uintptr
	q.Drv(211000,211103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QButtonGroup) OnButtonPressed(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(211000,211104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QButtonGroup) OnButtonPressedWithAbstractbutton(fn func(*QAbstractButton)) uintptr {
	var __rv uintptr
	q.Drv(211000,211105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QButtonGroup) OnButtonReleased(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(211000,211106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QButtonGroup) OnButtonReleasedWithAbstractbutton(fn func(*QAbstractButton)) uintptr {
	var __rv uintptr
	q.Drv(211000,211107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QButtonGroup::QButtonGroup()	
func NewButtonGroup() *QButtonGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),211000,211108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QButtonGroup{}
	_p.SetDriver(__rv,211000,false)
	return _p
} 
//QButtonGroup::QButtonGroup(QObject*)	
func NewButtonGroupWithParent(parent QObjectInterface) *QButtonGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),211000,211109,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QButtonGroup{}
	_p.SetDriver(__rv,211000,false)
	return _p
} 
//QButtonGroup::addButton(QAbstractButton*)
func (q *QButtonGroup) AddButton(value *QAbstractButton)  {
	q.Drv(211000,211110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QButtonGroup::addButton(QAbstractButton*,int)
func (q *QButtonGroup) AddButtonWithAbstractbuttonId(value2 *QAbstractButton,id int)  {
	q.Drv(211000,211111,Native(value2),unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QButtonGroup::button(int)
func (q *QButtonGroup) Button(id int) *QAbstractButton {
	var __rv uintptr
	q.Drv(211000,211112,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QButtonGroup::buttons()
func (q *QButtonGroup) Buttons() []*QAbstractButton {
	var __rv []*QAbstractButton
	q.Drv(211000,211113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QButtonGroup::checkedButton()
func (q *QButtonGroup) CheckedButton() *QAbstractButton {
	var __rv uintptr
	q.Drv(211000,211114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QButtonGroup::checkedId()
func (q *QButtonGroup) CheckedId() int {
	var __rv int
	q.Drv(211000,211115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QButtonGroup::exclusive()
func (q *QButtonGroup) Exclusive() bool {
	var __rv bool
	q.Drv(211000,211116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QButtonGroup::id(QAbstractButton*)
func (q *QButtonGroup) Id(button *QAbstractButton) int {
	var __rv int
	q.Drv(211000,211117,Native(button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QButtonGroup::removeButton(QAbstractButton*)
func (q *QButtonGroup) RemoveButton(value *QAbstractButton)  {
	q.Drv(211000,211118,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QButtonGroup::setExclusive(bool)
func (q *QButtonGroup) SetExclusive(value bool)  {
	q.Drv(211000,211119,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QButtonGroup::setId(QAbstractButton*,int)
func (q *QButtonGroup) SetId(button *QAbstractButton,id int)  {
	q.Drv(211000,211120,Native(button),unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
