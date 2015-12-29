// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLCDNumber_Mode - QLCDNumber::Mode
type QLCDNumber_Mode uint32
const (
	QLCDNumber_Hex QLCDNumber_Mode = 0
	QLCDNumber_Dec QLCDNumber_Mode = 1
	QLCDNumber_Oct QLCDNumber_Mode = 2
	QLCDNumber_Bin QLCDNumber_Mode = 3
)
//enum QLCDNumber_SegmentStyle - QLCDNumber::SegmentStyle
type QLCDNumber_SegmentStyle uint32
const (
	QLCDNumber_Outline QLCDNumber_SegmentStyle = 0
	QLCDNumber_Filled QLCDNumber_SegmentStyle = 1
	QLCDNumber_Flat QLCDNumber_SegmentStyle = 2
)
//struct QLCDNumber : QLCDNumber
type QLCDNumber struct {
	QFrame
}
func (q *QLCDNumber) OnOverflow(fn func()) uintptr {
	var __rv uintptr
	q.Drv(298000,298102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QLCDNumber::QLCDNumber()	
func NewLCDNumber() *QLCDNumber {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),298000,298103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLCDNumber{}
	_p.SetDriver(__rv,298000,false)
	return _p
} 
//QLCDNumber::QLCDNumber(QWidget*)	
func NewLCDNumberWithParent(parent QWidgetInterface) *QLCDNumber {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),298000,298104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLCDNumber{}
	_p.SetDriver(__rv,298000,false)
	return _p
} 
//QLCDNumber::QLCDNumber(unsigned int,QWidget*)	
func NewLCDNumberWithNumdigitsParent(numDigits uint,parent QWidgetInterface) *QLCDNumber {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),298000,298105,unsafe.Pointer(&numDigits),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLCDNumber{}
	_p.SetDriver(__rv,298000,false)
	return _p
} 
//QLCDNumber::checkOverflow(double)
func (q *QLCDNumber) CheckOverflow(num float64) bool {
	var __rv bool
	q.Drv(298000,298106,unsafe.Pointer(&num),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::checkOverflow(int)
func (q *QLCDNumber) CheckOverflowWithNum(num int) bool {
	var __rv bool
	q.Drv(298000,298107,unsafe.Pointer(&num),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::digitCount()
func (q *QLCDNumber) DigitCount() int {
	var __rv int
	q.Drv(298000,298108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::display(QString const&)
func (q *QLCDNumber) Display(str string)  {
	q.Drv(298000,298109,unsafe.Pointer(&str),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::display(double)
func (q *QLCDNumber) DisplayFWithNum(num float64)  {
	q.Drv(298000,298110,unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::display(int)
func (q *QLCDNumber) DisplayWithNum(num int)  {
	q.Drv(298000,298111,unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::event(QEvent*)
func (q *QLCDNumber) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(298000,298112,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::intValue()
func (q *QLCDNumber) IntValue() int {
	var __rv int
	q.Drv(298000,298113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::mode()
func (q *QLCDNumber) Mode() QLCDNumber_Mode {
	var __rv QLCDNumber_Mode
	q.Drv(298000,298114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::paintEvent(QPaintEvent*)
func (q *QLCDNumber) PaintEvent(value *QPaintEvent)  {
	q.Drv(298000,298115,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::segmentStyle()
func (q *QLCDNumber) SegmentStyle() QLCDNumber_SegmentStyle {
	var __rv QLCDNumber_SegmentStyle
	q.Drv(298000,298116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::setBinMode()
func (q *QLCDNumber) SetBinMode()  {
	q.Drv(298000,298117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setDecMode()
func (q *QLCDNumber) SetDecMode()  {
	q.Drv(298000,298118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setDigitCount(int)
func (q *QLCDNumber) SetDigitCount(nDigits int)  {
	q.Drv(298000,298119,unsafe.Pointer(&nDigits),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setHexMode()
func (q *QLCDNumber) SetHexMode()  {
	q.Drv(298000,298120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setMode(QLCDNumber::Mode)
func (q *QLCDNumber) SetMode(value QLCDNumber_Mode)  {
	q.Drv(298000,298121,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setOctMode()
func (q *QLCDNumber) SetOctMode()  {
	q.Drv(298000,298122,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setSegmentStyle(QLCDNumber::SegmentStyle)
func (q *QLCDNumber) SetSegmentStyle(value QLCDNumber_SegmentStyle)  {
	q.Drv(298000,298123,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::setSmallDecimalPoint(bool)
func (q *QLCDNumber) SetSmallDecimalPoint(value bool)  {
	q.Drv(298000,298124,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLCDNumber::sizeHint()
func (q *QLCDNumber) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(298000,298125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLCDNumber::smallDecimalPoint()
func (q *QLCDNumber) SmallDecimalPoint() bool {
	var __rv bool
	q.Drv(298000,298126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLCDNumber::value()
func (q *QLCDNumber) Value() float64 {
	var __rv float64
	q.Drv(298000,298127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
