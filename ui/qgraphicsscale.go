// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsScale : QGraphicsScale
type QGraphicsScale struct {
	QGraphicsTransform
}
func (q *QGraphicsScale) OnXScaleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(271000,271102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScale) OnScaleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(271000,271103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScale) OnZScaleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(271000,271104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScale) OnOriginChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(271000,271105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScale) OnYScaleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(271000,271106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsScale::QGraphicsScale()	
func NewGraphicsScale() *QGraphicsScale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),271000,271107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScale{}
	_p.SetDriver(__rv,271000,false)
	return _p
} 
//QGraphicsScale::QGraphicsScale(QObject*)	
func NewGraphicsScaleWithParent(parent QObjectInterface) *QGraphicsScale {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),271000,271108,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScale{}
	_p.SetDriver(__rv,271000,false)
	return _p
} 
//QGraphicsScale::origin()
func (q *QGraphicsScale) Origin() *QVector3D {
	var __rv uintptr
	q.Drv(271000,271109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QGraphicsScale::setOrigin(QVector3D const&)
func (q *QGraphicsScale) SetOrigin(point *QVector3D)  {
	q.Drv(271000,271110,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScale::setXScale(double)
func (q *QGraphicsScale) SetXScale(value float64)  {
	q.Drv(271000,271111,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScale::setYScale(double)
func (q *QGraphicsScale) SetYScale(value float64)  {
	q.Drv(271000,271112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScale::setZScale(double)
func (q *QGraphicsScale) SetZScale(value float64)  {
	q.Drv(271000,271113,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScale::xScale()
func (q *QGraphicsScale) XScale() float64 {
	var __rv float64
	q.Drv(271000,271114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScale::yScale()
func (q *QGraphicsScale) YScale() float64 {
	var __rv float64
	q.Drv(271000,271115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScale::zScale()
func (q *QGraphicsScale) ZScale() float64 {
	var __rv float64
	q.Drv(271000,271116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
