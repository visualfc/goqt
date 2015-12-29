// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QConicalGradient : QConicalGradient
type QConicalGradient struct {
	QGradient
}
//QConicalGradient::QConicalGradient()	
func NewConicalGradient() *QConicalGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),15000,15102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QConicalGradient{}
	_p.SetDriver(__rv,15000,true)
	return _p
} 
//QConicalGradient::QConicalGradient(QPointF const&,double)	
func NewConicalGradientWithCenterStartangle(center *QPointF,startAngle float64) *QConicalGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),15000,15103,Native(center),unsafe.Pointer(&startAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QConicalGradient{}
	_p.SetDriver(__rv,15000,true)
	return _p
} 
//QConicalGradient::QConicalGradient(double,double,double)	
func NewConicalGradientWithCxCyStartangle(cx float64,cy float64,startAngle float64) *QConicalGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),15000,15104,unsafe.Pointer(&cx),unsafe.Pointer(&cy),unsafe.Pointer(&startAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QConicalGradient{}
	_p.SetDriver(__rv,15000,true)
	return _p
} 
//QConicalGradient::angle()
func (q *QConicalGradient) Angle() float64 {
	var __rv float64
	q.Drv(15000,15105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QConicalGradient::center()
func (q *QConicalGradient) Center() *QPointF {
	var __rv uintptr
	q.Drv(15000,15106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QConicalGradient::setAngle(double)
func (q *QConicalGradient) SetAngle(angle float64)  {
	q.Drv(15000,15107,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QConicalGradient::setCenter(QPointF const&)
func (q *QConicalGradient) SetCenter(center *QPointF)  {
	q.Drv(15000,15108,Native(center),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QConicalGradient::setCenter(double,double)
func (q *QConicalGradient) SetCenterFWithXY(x float64,y float64)  {
	q.Drv(15000,15109,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
