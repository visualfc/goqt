// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVector4D : QVector4D
type QVector4D struct {
	BaseDrv
}
//QVector4D::QVector4D()	
func NewVector4D() *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QPoint const&)	
func NewVector4DWithPoint(point *QPoint) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186103,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QPointF const&)	
func NewVector4DFWithPoint(point *QPointF) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186104,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QVector2D const&)	
func NewVector4DWithVector2d(vector *QVector2D) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186105,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QVector3D const&)	
func NewVector4DWithVector3d(vector *QVector3D) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186106,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QVector3D const&,double)	
func NewVector4DWithVector3dWpos(vector *QVector3D,wpos float64) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186107,Native(vector),unsafe.Pointer(&wpos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(QVector2D const&,double,double)	
func NewVector4DWithVector2dZposWpos(vector *QVector2D,zpos float64,wpos float64) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186108,Native(vector),unsafe.Pointer(&zpos),unsafe.Pointer(&wpos),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::QVector4D(double,double,double,double)	
func NewVector4DWithXposYposZposWpos(xpos float64,ypos float64,zpos float64,wpos float64) *QVector4D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),186000,186109,unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),unsafe.Pointer(&zpos),unsafe.Pointer(&wpos),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector4D{}
	_p.SetDriver(__rv,186000,true)
	return _p
} 
//QVector4D::dotProduct(QVector4D const&,QVector4D const&)	
func QVector4DDotProduct(v1 *QVector4D,v2 *QVector4D) float64 {
	var __rv float64
	DirectQtDrv(nil,186000,186110,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::dotProduct(QVector4D const&,QVector4D const&)
func (q *QVector4D) DotProduct(v1 *QVector4D,v2 *QVector4D) float64 {
	var __rv float64
	q.Drv(186000,186110,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::isNull()
func (q *QVector4D) IsNull() bool {
	var __rv bool
	q.Drv(186000,186111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::length()
func (q *QVector4D) Length() float64 {
	var __rv float64
	q.Drv(186000,186112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::lengthSquared()
func (q *QVector4D) LengthSquared() float64 {
	var __rv float64
	q.Drv(186000,186113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::normalize()
func (q *QVector4D) Normalize()  {
	q.Drv(186000,186114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector4D::normalized()
func (q *QVector4D) Normalized() *QVector4D {
	var __rv uintptr
	q.Drv(186000,186115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector4D{}
	_rp.SetDriver(__rv,186000,true)
	return _rp
}	
//QVector4D::setW(double)
func (q *QVector4D) SetW(w float64)  {
	q.Drv(186000,186116,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector4D::setX(double)
func (q *QVector4D) SetX(x float64)  {
	q.Drv(186000,186117,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector4D::setY(double)
func (q *QVector4D) SetY(y float64)  {
	q.Drv(186000,186118,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector4D::setZ(double)
func (q *QVector4D) SetZ(z float64)  {
	q.Drv(186000,186119,unsafe.Pointer(&z),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector4D::toPoint()
func (q *QVector4D) ToPoint() *QPoint {
	var __rv uintptr
	q.Drv(186000,186120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QVector4D::toPointF()
func (q *QVector4D) ToPointF() *QPointF {
	var __rv uintptr
	q.Drv(186000,186121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QVector4D::toVector2D()
func (q *QVector4D) ToVector2D() *QVector2D {
	var __rv uintptr
	q.Drv(186000,186122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector2D{}
	_rp.SetDriver(__rv,184000,true)
	return _rp
}	
//QVector4D::toVector2DAffine()
func (q *QVector4D) ToVector2DAffine() *QVector2D {
	var __rv uintptr
	q.Drv(186000,186123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector2D{}
	_rp.SetDriver(__rv,184000,true)
	return _rp
}	
//QVector4D::toVector3D()
func (q *QVector4D) ToVector3D() *QVector3D {
	var __rv uintptr
	q.Drv(186000,186124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector4D::toVector3DAffine()
func (q *QVector4D) ToVector3DAffine() *QVector3D {
	var __rv uintptr
	q.Drv(186000,186125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector4D::w()
func (q *QVector4D) W() float64 {
	var __rv float64
	q.Drv(186000,186126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::x()
func (q *QVector4D) X() float64 {
	var __rv float64
	q.Drv(186000,186127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::y()
func (q *QVector4D) Y() float64 {
	var __rv float64
	q.Drv(186000,186128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector4D::z()
func (q *QVector4D) Z() float64 {
	var __rv float64
	q.Drv(186000,186129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
