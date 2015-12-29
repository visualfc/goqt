// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QRadialGradient : QRadialGradient
type QRadialGradient struct {
	QGradient
}
//QRadialGradient::QRadialGradient()	
func NewRadialGradient() *QRadialGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),108000,108102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadialGradient{}
	_p.SetDriver(__rv,108000,true)
	return _p
} 
//QRadialGradient::QRadialGradient(QPointF const&,double)	
func NewRadialGradientWithCenterRadius(center *QPointF,radius float64) *QRadialGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),108000,108103,Native(center),unsafe.Pointer(&radius),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadialGradient{}
	_p.SetDriver(__rv,108000,true)
	return _p
} 
//QRadialGradient::QRadialGradient(QPointF const&,double,QPointF const&)	
func NewRadialGradientWithCenterRadiusFocalpoint(center *QPointF,radius float64,focalPoint *QPointF) *QRadialGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),108000,108104,Native(center),unsafe.Pointer(&radius),Native(focalPoint),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadialGradient{}
	_p.SetDriver(__rv,108000,true)
	return _p
} 
//QRadialGradient::QRadialGradient(double,double,double)	
func NewRadialGradientWithCxCyRadius(cx float64,cy float64,radius float64) *QRadialGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),108000,108105,unsafe.Pointer(&cx),unsafe.Pointer(&cy),unsafe.Pointer(&radius),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadialGradient{}
	_p.SetDriver(__rv,108000,true)
	return _p
} 
//QRadialGradient::QRadialGradient(double,double,double,double,double)	
func NewRadialGradientWithCxCyRadiusFxFy(cx float64,cy float64,radius float64,fx float64,fy float64) *QRadialGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),108000,108106,unsafe.Pointer(&cx),unsafe.Pointer(&cy),unsafe.Pointer(&radius),unsafe.Pointer(&fx),unsafe.Pointer(&fy),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRadialGradient{}
	_p.SetDriver(__rv,108000,true)
	return _p
} 
//QRadialGradient::center()
func (q *QRadialGradient) Center() *QPointF {
	var __rv uintptr
	q.Drv(108000,108107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRadialGradient::focalPoint()
func (q *QRadialGradient) FocalPoint() *QPointF {
	var __rv uintptr
	q.Drv(108000,108108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRadialGradient::radius()
func (q *QRadialGradient) Radius() float64 {
	var __rv float64
	q.Drv(108000,108109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRadialGradient::setCenter(QPointF const&)
func (q *QRadialGradient) SetCenter(center *QPointF)  {
	q.Drv(108000,108110,Native(center),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadialGradient::setCenter(double,double)
func (q *QRadialGradient) SetCenterFWithXY(x float64,y float64)  {
	q.Drv(108000,108111,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadialGradient::setFocalPoint(QPointF const&)
func (q *QRadialGradient) SetFocalPoint(focalPoint *QPointF)  {
	q.Drv(108000,108112,Native(focalPoint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadialGradient::setFocalPoint(double,double)
func (q *QRadialGradient) SetFocalPointFWithXY(x float64,y float64)  {
	q.Drv(108000,108113,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRadialGradient::setRadius(double)
func (q *QRadialGradient) SetRadius(radius float64)  {
	q.Drv(108000,108114,unsafe.Pointer(&radius),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
