// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsRotation : QGraphicsRotation
type QGraphicsRotation struct {
	QGraphicsTransform
}
func (q *QGraphicsRotation) OnAxisChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(270000,270102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsRotation) OnAngleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(270000,270103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsRotation) OnOriginChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(270000,270104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsRotation::QGraphicsRotation()	
func NewGraphicsRotation() *QGraphicsRotation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),270000,270105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRotation{}
	_p.SetDriver(__rv,270000,false)
	return _p
} 
//QGraphicsRotation::QGraphicsRotation(QObject*)	
func NewGraphicsRotationWithParent(parent QObjectInterface) *QGraphicsRotation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),270000,270106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsRotation{}
	_p.SetDriver(__rv,270000,false)
	return _p
} 
//QGraphicsRotation::angle()
func (q *QGraphicsRotation) Angle() float64 {
	var __rv float64
	q.Drv(270000,270107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsRotation::axis()
func (q *QGraphicsRotation) Axis() *QVector3D {
	var __rv uintptr
	q.Drv(270000,270108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QGraphicsRotation::origin()
func (q *QGraphicsRotation) Origin() *QVector3D {
	var __rv uintptr
	q.Drv(270000,270109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QGraphicsRotation::setAngle(double)
func (q *QGraphicsRotation) SetAngle(value float64)  {
	q.Drv(270000,270110,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsRotation::setAxis(QVector3D const&)
func (q *QGraphicsRotation) SetAxis(axis *QVector3D)  {
	q.Drv(270000,270111,Native(axis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsRotation::setAxis(Qt::Axis)
func (q *QGraphicsRotation) SetAxisWithAxis(axis Qt_Axis)  {
	q.Drv(270000,270112,unsafe.Pointer(&axis),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsRotation::setOrigin(QVector3D const&)
func (q *QGraphicsRotation) SetOrigin(point *QVector3D)  {
	q.Drv(270000,270113,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
