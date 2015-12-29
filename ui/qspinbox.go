// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSpinBox : QSpinBox
type QSpinBox struct {
	QAbstractSpinBox
}
func (q *QSpinBox) OnValueChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(346000,346102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QSpinBox) OnValueChangedWithInt(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(346000,346103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSpinBox::QSpinBox()	
func NewSpinBox() *QSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),346000,346104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSpinBox{}
	_p.SetDriver(__rv,346000,false)
	return _p
} 
//QSpinBox::QSpinBox(QWidget*)	
func NewSpinBoxWithParent(parent QWidgetInterface) *QSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),346000,346105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSpinBox{}
	_p.SetDriver(__rv,346000,false)
	return _p
} 
//QSpinBox::cleanText()
func (q *QSpinBox) CleanText() string {
	var __rv string
	q.Drv(346000,346106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::event(QEvent*)
func (q *QSpinBox) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(346000,346107,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::fixup(QString&)
func (q *QSpinBox) Fixup(str *string)  {
	q.Drv(346000,346108,unsafe.Pointer(&str),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::maximum()
func (q *QSpinBox) Maximum() int {
	var __rv int
	q.Drv(346000,346109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::minimum()
func (q *QSpinBox) Minimum() int {
	var __rv int
	q.Drv(346000,346110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::prefix()
func (q *QSpinBox) Prefix() string {
	var __rv string
	q.Drv(346000,346111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::setMaximum(int)
func (q *QSpinBox) SetMaximum(max int)  {
	q.Drv(346000,346112,unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setMinimum(int)
func (q *QSpinBox) SetMinimum(min int)  {
	q.Drv(346000,346113,unsafe.Pointer(&min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setPrefix(QString const&)
func (q *QSpinBox) SetPrefix(prefix string)  {
	q.Drv(346000,346114,unsafe.Pointer(&prefix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setRange(int,int)
func (q *QSpinBox) SetRange(min int,max int)  {
	q.Drv(346000,346115,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setSingleStep(int)
func (q *QSpinBox) SetSingleStep(val int)  {
	q.Drv(346000,346116,unsafe.Pointer(&val),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setSuffix(QString const&)
func (q *QSpinBox) SetSuffix(suffix string)  {
	q.Drv(346000,346117,unsafe.Pointer(&suffix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::setValue(int)
func (q *QSpinBox) SetValue(val int)  {
	q.Drv(346000,346118,unsafe.Pointer(&val),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSpinBox::singleStep()
func (q *QSpinBox) SingleStep() int {
	var __rv int
	q.Drv(346000,346119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::suffix()
func (q *QSpinBox) Suffix() string {
	var __rv string
	q.Drv(346000,346120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::textFromValue(int)
func (q *QSpinBox) TextFromValue(val int) string {
	var __rv string
	q.Drv(346000,346121,unsafe.Pointer(&val),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::validate(QString&,int&)
func (q *QSpinBox) Validate(input *string,pos *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(346000,346122,unsafe.Pointer(&input),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::value()
func (q *QSpinBox) Value() int {
	var __rv int
	q.Drv(346000,346123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSpinBox::valueFromText(QString const&)
func (q *QSpinBox) ValueFromText(text string) int {
	var __rv int
	q.Drv(346000,346124,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
