// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QEasingCurve_Type - QEasingCurve::Type
type QEasingCurve_Type uint32
const (
	QEasingCurve_Linear QEasingCurve_Type = 0
	QEasingCurve_InQuad QEasingCurve_Type = 1
	QEasingCurve_OutQuad QEasingCurve_Type = 2
	QEasingCurve_InOutQuad QEasingCurve_Type = 3
	QEasingCurve_OutInQuad QEasingCurve_Type = 4
	QEasingCurve_InCubic QEasingCurve_Type = 5
	QEasingCurve_OutCubic QEasingCurve_Type = 6
	QEasingCurve_InOutCubic QEasingCurve_Type = 7
	QEasingCurve_OutInCubic QEasingCurve_Type = 8
	QEasingCurve_InQuart QEasingCurve_Type = 9
	QEasingCurve_OutQuart QEasingCurve_Type = 10
	QEasingCurve_InOutQuart QEasingCurve_Type = 11
	QEasingCurve_OutInQuart QEasingCurve_Type = 12
	QEasingCurve_InQuint QEasingCurve_Type = 13
	QEasingCurve_OutQuint QEasingCurve_Type = 14
	QEasingCurve_InOutQuint QEasingCurve_Type = 15
	QEasingCurve_OutInQuint QEasingCurve_Type = 16
	QEasingCurve_InSine QEasingCurve_Type = 17
	QEasingCurve_OutSine QEasingCurve_Type = 18
	QEasingCurve_InOutSine QEasingCurve_Type = 19
	QEasingCurve_OutInSine QEasingCurve_Type = 20
	QEasingCurve_InExpo QEasingCurve_Type = 21
	QEasingCurve_OutExpo QEasingCurve_Type = 22
	QEasingCurve_InOutExpo QEasingCurve_Type = 23
	QEasingCurve_OutInExpo QEasingCurve_Type = 24
	QEasingCurve_InCirc QEasingCurve_Type = 25
	QEasingCurve_OutCirc QEasingCurve_Type = 26
	QEasingCurve_InOutCirc QEasingCurve_Type = 27
	QEasingCurve_OutInCirc QEasingCurve_Type = 28
	QEasingCurve_InElastic QEasingCurve_Type = 29
	QEasingCurve_OutElastic QEasingCurve_Type = 30
	QEasingCurve_InOutElastic QEasingCurve_Type = 31
	QEasingCurve_OutInElastic QEasingCurve_Type = 32
	QEasingCurve_InBack QEasingCurve_Type = 33
	QEasingCurve_OutBack QEasingCurve_Type = 34
	QEasingCurve_InOutBack QEasingCurve_Type = 35
	QEasingCurve_OutInBack QEasingCurve_Type = 36
	QEasingCurve_InBounce QEasingCurve_Type = 37
	QEasingCurve_OutBounce QEasingCurve_Type = 38
	QEasingCurve_InOutBounce QEasingCurve_Type = 39
	QEasingCurve_OutInBounce QEasingCurve_Type = 40
	QEasingCurve_InCurve QEasingCurve_Type = 41
	QEasingCurve_OutCurve QEasingCurve_Type = 42
	QEasingCurve_SineCurve QEasingCurve_Type = 43
	QEasingCurve_CosineCurve QEasingCurve_Type = 44
	QEasingCurve_Custom QEasingCurve_Type = 45
	QEasingCurve_NCurveTypes QEasingCurve_Type = 46
)
//struct QEasingCurve : QEasingCurve
type QEasingCurve struct {
	BaseDrv
}
//QEasingCurve::QEasingCurve()	
func NewEasingCurve() *QEasingCurve {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),29000,29102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEasingCurve{}
	_p.SetDriver(__rv,29000,true)
	return _p
} 
//QEasingCurve::QEasingCurve(QEasingCurve const&)	
func NewEasingCurveCopy(other *QEasingCurve) *QEasingCurve {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),29000,29103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEasingCurve{}
	_p.SetDriver(__rv,29000,true)
	return _p
} 
//QEasingCurve::QEasingCurve(QEasingCurve::Type)	
func NewEasingCurveWithType(_type QEasingCurve_Type) *QEasingCurve {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),29000,29104,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEasingCurve{}
	_p.SetDriver(__rv,29000,true)
	return _p
} 
//QEasingCurve::amplitude()
func (q *QEasingCurve) Amplitude() float64 {
	var __rv float64
	q.Drv(29000,29105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEasingCurve::overshoot()
func (q *QEasingCurve) Overshoot() float64 {
	var __rv float64
	q.Drv(29000,29106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEasingCurve::period()
func (q *QEasingCurve) Period() float64 {
	var __rv float64
	q.Drv(29000,29107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEasingCurve::setAmplitude(double)
func (q *QEasingCurve) SetAmplitude(amplitude float64)  {
	q.Drv(29000,29108,unsafe.Pointer(&amplitude),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEasingCurve::setOvershoot(double)
func (q *QEasingCurve) SetOvershoot(overshoot float64)  {
	q.Drv(29000,29109,unsafe.Pointer(&overshoot),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEasingCurve::setPeriod(double)
func (q *QEasingCurve) SetPeriod(period float64)  {
	q.Drv(29000,29110,unsafe.Pointer(&period),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEasingCurve::setType(QEasingCurve::Type)
func (q *QEasingCurve) SetType(_type QEasingCurve_Type)  {
	q.Drv(29000,29111,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEasingCurve::type()
func (q *QEasingCurve) Type() QEasingCurve_Type {
	var __rv QEasingCurve_Type
	q.Drv(29000,29112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEasingCurve::valueForProgress(double)
func (q *QEasingCurve) ValueForProgress(progress float64) float64 {
	var __rv float64
	q.Drv(29000,29113,unsafe.Pointer(&progress),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
