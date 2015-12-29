// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPointF : QPointF
type QPointF struct {
	BaseDrv
}
//QPointF::QPointF()	
func NewPointF() *QPointF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),100000,100102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPointF{}
	_p.SetDriver(__rv,100000,true)
	return _p
} 
//QPointF::QPointF(QPoint const&)	
func NewPointFWithPoint(p *QPoint) *QPointF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),100000,100103,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPointF{}
	_p.SetDriver(__rv,100000,true)
	return _p
} 
//QPointF::QPointF(double,double)	
func NewPointFWithXposYpos(xpos float64,ypos float64) *QPointF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),100000,100104,unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPointF{}
	_p.SetDriver(__rv,100000,true)
	return _p
} 
//QPointF::isNull()
func (q *QPointF) IsNull() bool {
	var __rv bool
	q.Drv(100000,100105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPointF::manhattanLength()
func (q *QPointF) ManhattanLength() float64 {
	var __rv float64
	q.Drv(100000,100106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPointF::rx()
func (q *QPointF) Rx() *float64 {
	var __rv *float64
	q.Drv(100000,100107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPointF::ry()
func (q *QPointF) Ry() *float64 {
	var __rv *float64
	q.Drv(100000,100108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPointF::setX(double)
func (q *QPointF) SetX(x float64)  {
	q.Drv(100000,100109,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPointF::setY(double)
func (q *QPointF) SetY(y float64)  {
	q.Drv(100000,100110,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPointF::toPoint()
func (q *QPointF) ToPoint() *QPoint {
	var __rv uintptr
	q.Drv(100000,100111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QPointF::x()
func (q *QPointF) X() float64 {
	var __rv float64
	q.Drv(100000,100112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPointF::y()
func (q *QPointF) Y() float64 {
	var __rv float64
	q.Drv(100000,100113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
