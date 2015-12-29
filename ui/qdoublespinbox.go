// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDoubleSpinBox : QDoubleSpinBox
type QDoubleSpinBox struct {
	QAbstractSpinBox
}
func (q *QDoubleSpinBox) OnValueChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(230000,230102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDoubleSpinBox) OnValueChangedFWithFloat64(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(230000,230103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDoubleSpinBox::QDoubleSpinBox()	
func NewDoubleSpinBox() *QDoubleSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),230000,230104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDoubleSpinBox{}
	_p.SetDriver(__rv,230000,false)
	return _p
} 
//QDoubleSpinBox::QDoubleSpinBox(QWidget*)	
func NewDoubleSpinBoxWithParent(parent QWidgetInterface) *QDoubleSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),230000,230105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDoubleSpinBox{}
	_p.SetDriver(__rv,230000,false)
	return _p
} 
//QDoubleSpinBox::cleanText()
func (q *QDoubleSpinBox) CleanText() string {
	var __rv string
	q.Drv(230000,230106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::decimals()
func (q *QDoubleSpinBox) Decimals() int {
	var __rv int
	q.Drv(230000,230107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::fixup(QString&)
func (q *QDoubleSpinBox) Fixup(str *string)  {
	q.Drv(230000,230108,unsafe.Pointer(&str),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::maximum()
func (q *QDoubleSpinBox) Maximum() float64 {
	var __rv float64
	q.Drv(230000,230109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::minimum()
func (q *QDoubleSpinBox) Minimum() float64 {
	var __rv float64
	q.Drv(230000,230110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::prefix()
func (q *QDoubleSpinBox) Prefix() string {
	var __rv string
	q.Drv(230000,230111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::setDecimals(int)
func (q *QDoubleSpinBox) SetDecimals(prec int)  {
	q.Drv(230000,230112,unsafe.Pointer(&prec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setMaximum(double)
func (q *QDoubleSpinBox) SetMaximum(max float64)  {
	q.Drv(230000,230113,unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setMinimum(double)
func (q *QDoubleSpinBox) SetMinimum(min float64)  {
	q.Drv(230000,230114,unsafe.Pointer(&min),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setPrefix(QString const&)
func (q *QDoubleSpinBox) SetPrefix(prefix string)  {
	q.Drv(230000,230115,unsafe.Pointer(&prefix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setRange(double,double)
func (q *QDoubleSpinBox) SetRange(min float64,max float64)  {
	q.Drv(230000,230116,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setSingleStep(double)
func (q *QDoubleSpinBox) SetSingleStep(val float64)  {
	q.Drv(230000,230117,unsafe.Pointer(&val),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setSuffix(QString const&)
func (q *QDoubleSpinBox) SetSuffix(suffix string)  {
	q.Drv(230000,230118,unsafe.Pointer(&suffix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::setValue(double)
func (q *QDoubleSpinBox) SetValue(val float64)  {
	q.Drv(230000,230119,unsafe.Pointer(&val),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDoubleSpinBox::singleStep()
func (q *QDoubleSpinBox) SingleStep() float64 {
	var __rv float64
	q.Drv(230000,230120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::suffix()
func (q *QDoubleSpinBox) Suffix() string {
	var __rv string
	q.Drv(230000,230121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::textFromValue(double)
func (q *QDoubleSpinBox) TextFromValue(val float64) string {
	var __rv string
	q.Drv(230000,230122,unsafe.Pointer(&val),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::validate(QString&,int&)
func (q *QDoubleSpinBox) Validate(input *string,pos *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(230000,230123,unsafe.Pointer(&input),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::value()
func (q *QDoubleSpinBox) Value() float64 {
	var __rv float64
	q.Drv(230000,230124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDoubleSpinBox::valueFromText(QString const&)
func (q *QDoubleSpinBox) ValueFromText(text string) float64 {
	var __rv float64
	q.Drv(230000,230125,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
