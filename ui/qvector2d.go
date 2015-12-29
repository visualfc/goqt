// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVector2D : QVector2D
type QVector2D struct {
	BaseDrv
}
//QVector2D::QVector2D()	
func NewVector2D() *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::QVector2D(QPoint const&)	
func NewVector2DWithPoint(point *QPoint) *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184103,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::QVector2D(QPointF const&)	
func NewVector2DFWithPoint(point *QPointF) *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184104,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::QVector2D(QVector3D const&)	
func NewVector2DWithVector3d(vector *QVector3D) *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184105,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::QVector2D(QVector4D const&)	
func NewVector2DWithVector4d(vector *QVector4D) *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184106,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::QVector2D(double,double)	
func NewVector2DWithXposYpos(xpos float64,ypos float64) *QVector2D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),184000,184107,unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector2D{}
	_p.SetDriver(__rv,184000,true)
	return _p
} 
//QVector2D::dotProduct(QVector2D const&,QVector2D const&)	
func QVector2DDotProduct(v1 *QVector2D,v2 *QVector2D) float64 {
	var __rv float64
	DirectQtDrv(nil,184000,184108,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::dotProduct(QVector2D const&,QVector2D const&)
func (q *QVector2D) DotProduct(v1 *QVector2D,v2 *QVector2D) float64 {
	var __rv float64
	q.Drv(184000,184108,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::isNull()
func (q *QVector2D) IsNull() bool {
	var __rv bool
	q.Drv(184000,184109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::length()
func (q *QVector2D) Length() float64 {
	var __rv float64
	q.Drv(184000,184110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::lengthSquared()
func (q *QVector2D) LengthSquared() float64 {
	var __rv float64
	q.Drv(184000,184111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::normalize()
func (q *QVector2D) Normalize()  {
	q.Drv(184000,184112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector2D::normalized()
func (q *QVector2D) Normalized() *QVector2D {
	var __rv uintptr
	q.Drv(184000,184113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector2D{}
	_rp.SetDriver(__rv,184000,true)
	return _rp
}	
//QVector2D::setX(double)
func (q *QVector2D) SetX(x float64)  {
	q.Drv(184000,184114,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector2D::setY(double)
func (q *QVector2D) SetY(y float64)  {
	q.Drv(184000,184115,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector2D::toPoint()
func (q *QVector2D) ToPoint() *QPoint {
	var __rv uintptr
	q.Drv(184000,184116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QVector2D::toPointF()
func (q *QVector2D) ToPointF() *QPointF {
	var __rv uintptr
	q.Drv(184000,184117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QVector2D::toVector3D()
func (q *QVector2D) ToVector3D() *QVector3D {
	var __rv uintptr
	q.Drv(184000,184118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector2D::toVector4D()
func (q *QVector2D) ToVector4D() *QVector4D {
	var __rv uintptr
	q.Drv(184000,184119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector4D{}
	_rp.SetDriver(__rv,186000,true)
	return _rp
}	
//QVector2D::x()
func (q *QVector2D) X() float64 {
	var __rv float64
	q.Drv(184000,184120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector2D::y()
func (q *QVector2D) Y() float64 {
	var __rv float64
	q.Drv(184000,184121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
