// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVector3D : QVector3D
type QVector3D struct {
	BaseDrv
}
//QVector3D::QVector3D()	
func NewVector3D() *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(QPoint const&)	
func NewVector3DWithPoint(point *QPoint) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185103,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(QPointF const&)	
func NewVector3DFWithPoint(point *QPointF) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185104,Native(point),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(QVector2D const&)	
func NewVector3DWithVector2d(vector *QVector2D) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185105,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(QVector4D const&)	
func NewVector3DWithVector4d(vector *QVector4D) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185106,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(QVector2D const&,double)	
func NewVector3DWithVector2dZpos(vector *QVector2D,zpos float64) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185107,Native(vector),unsafe.Pointer(&zpos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::QVector3D(double,double,double)	
func NewVector3DWithXposYposZpos(xpos float64,ypos float64,zpos float64) *QVector3D {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),185000,185108,unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),unsafe.Pointer(&zpos),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QVector3D{}
	_p.SetDriver(__rv,185000,true)
	return _p
} 
//QVector3D::crossProduct(QVector3D const&,QVector3D const&)	
func QVector3DCrossProduct(v1 *QVector3D,v2 *QVector3D) *QVector3D {
	var __rv uintptr
	DirectQtDrv(nil,185000,185109,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::crossProduct(QVector3D const&,QVector3D const&)
func (q *QVector3D) CrossProduct(v1 *QVector3D,v2 *QVector3D) *QVector3D {
	var __rv uintptr
	q.Drv(185000,185109,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::distanceToLine(QVector3D const&,QVector3D const&)
func (q *QVector3D) DistanceToLine(point *QVector3D,direction *QVector3D) float64 {
	var __rv float64
	q.Drv(185000,185110,Native(point),Native(direction),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::distanceToPlane(QVector3D const&,QVector3D const&)
func (q *QVector3D) DistanceToPlaneWithPlaneNormal(plane *QVector3D,normal *QVector3D) float64 {
	var __rv float64
	q.Drv(185000,185111,Native(plane),Native(normal),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::distanceToPlane(QVector3D const&,QVector3D const&,QVector3D const&)
func (q *QVector3D) DistanceToPlaneWithPlane1Plane2Plane3(plane1 *QVector3D,plane2 *QVector3D,plane3 *QVector3D) float64 {
	var __rv float64
	q.Drv(185000,185112,Native(plane1),Native(plane2),Native(plane3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::dotProduct(QVector3D const&,QVector3D const&)	
func QVector3DDotProduct(v1 *QVector3D,v2 *QVector3D) float64 {
	var __rv float64
	DirectQtDrv(nil,185000,185113,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::dotProduct(QVector3D const&,QVector3D const&)
func (q *QVector3D) DotProduct(v1 *QVector3D,v2 *QVector3D) float64 {
	var __rv float64
	q.Drv(185000,185113,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::isNull()
func (q *QVector3D) IsNull() bool {
	var __rv bool
	q.Drv(185000,185114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::length()
func (q *QVector3D) Length() float64 {
	var __rv float64
	q.Drv(185000,185115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::lengthSquared()
func (q *QVector3D) LengthSquared() float64 {
	var __rv float64
	q.Drv(185000,185116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::normal(QVector3D const&,QVector3D const&)	
func QVector3DNormalWithV1V2(v1 *QVector3D,v2 *QVector3D) *QVector3D {
	var __rv uintptr
	DirectQtDrv(nil,185000,185117,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::normal(QVector3D const&,QVector3D const&)
func (q *QVector3D) NormalWithV1V2(v1 *QVector3D,v2 *QVector3D) *QVector3D {
	var __rv uintptr
	q.Drv(185000,185117,Native(v1),Native(v2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::normal(QVector3D const&,QVector3D const&,QVector3D const&)	
func QVector3DNormalWithV1V2V3(v1 *QVector3D,v2 *QVector3D,v3 *QVector3D) *QVector3D {
	var __rv uintptr
	DirectQtDrv(nil,185000,185118,Native(v1),Native(v2),Native(v3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::normal(QVector3D const&,QVector3D const&,QVector3D const&)
func (q *QVector3D) NormalWithV1V2V3(v1 *QVector3D,v2 *QVector3D,v3 *QVector3D) *QVector3D {
	var __rv uintptr
	q.Drv(185000,185118,Native(v1),Native(v2),Native(v3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::normalize()
func (q *QVector3D) Normalize()  {
	q.Drv(185000,185119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector3D::normalized()
func (q *QVector3D) Normalized() *QVector3D {
	var __rv uintptr
	q.Drv(185000,185120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QVector3D::setX(double)
func (q *QVector3D) SetX(x float64)  {
	q.Drv(185000,185121,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector3D::setY(double)
func (q *QVector3D) SetY(y float64)  {
	q.Drv(185000,185122,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector3D::setZ(double)
func (q *QVector3D) SetZ(z float64)  {
	q.Drv(185000,185123,unsafe.Pointer(&z),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVector3D::toPoint()
func (q *QVector3D) ToPoint() *QPoint {
	var __rv uintptr
	q.Drv(185000,185124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QVector3D::toPointF()
func (q *QVector3D) ToPointF() *QPointF {
	var __rv uintptr
	q.Drv(185000,185125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QVector3D::toVector2D()
func (q *QVector3D) ToVector2D() *QVector2D {
	var __rv uintptr
	q.Drv(185000,185126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector2D{}
	_rp.SetDriver(__rv,184000,true)
	return _rp
}	
//QVector3D::toVector4D()
func (q *QVector3D) ToVector4D() *QVector4D {
	var __rv uintptr
	q.Drv(185000,185127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector4D{}
	_rp.SetDriver(__rv,186000,true)
	return _rp
}	
//QVector3D::x()
func (q *QVector3D) X() float64 {
	var __rv float64
	q.Drv(185000,185128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::y()
func (q *QVector3D) Y() float64 {
	var __rv float64
	q.Drv(185000,185129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVector3D::z()
func (q *QVector3D) Z() float64 {
	var __rv float64
	q.Drv(185000,185130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
