// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QLinearGradient : QLinearGradient
type QLinearGradient struct {
	QGradient
}
//QLinearGradient::QLinearGradient()	
func NewLinearGradient() *QLinearGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),70000,70102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLinearGradient{}
	_p.SetDriver(__rv,70000,true)
	return _p
} 
//QLinearGradient::QLinearGradient(QPointF const&,QPointF const&)	
func NewLinearGradientWithStartFinalstop(start *QPointF,finalStop *QPointF) *QLinearGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),70000,70103,Native(start),Native(finalStop),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLinearGradient{}
	_p.SetDriver(__rv,70000,true)
	return _p
} 
//QLinearGradient::QLinearGradient(double,double,double,double)	
func NewLinearGradientWithXstartYstartXfinalstopYfinalstop(xStart float64,yStart float64,xFinalStop float64,yFinalStop float64) *QLinearGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),70000,70104,unsafe.Pointer(&xStart),unsafe.Pointer(&yStart),unsafe.Pointer(&xFinalStop),unsafe.Pointer(&yFinalStop),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLinearGradient{}
	_p.SetDriver(__rv,70000,true)
	return _p
} 
//QLinearGradient::finalStop()
func (q *QLinearGradient) FinalStop() *QPointF {
	var __rv uintptr
	q.Drv(70000,70105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QLinearGradient::setFinalStop(QPointF const&)
func (q *QLinearGradient) SetFinalStop(stop *QPointF)  {
	q.Drv(70000,70106,Native(stop),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLinearGradient::setFinalStop(double,double)
func (q *QLinearGradient) SetFinalStopFWithXY(x float64,y float64)  {
	q.Drv(70000,70107,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLinearGradient::setStart(QPointF const&)
func (q *QLinearGradient) SetStart(start *QPointF)  {
	q.Drv(70000,70108,Native(start),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLinearGradient::setStart(double,double)
func (q *QLinearGradient) SetStartFWithXY(x float64,y float64)  {
	q.Drv(70000,70109,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLinearGradient::start()
func (q *QLinearGradient) Start() *QPointF {
	var __rv uintptr
	q.Drv(70000,70110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
