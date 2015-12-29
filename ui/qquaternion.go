// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QQuaternion : QQuaternion
type QQuaternion struct {
	BaseDrv
}
//QQuaternion::QQuaternion()	
func NewQuaternion() *QQuaternion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),107000,107102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QQuaternion{}
	_p.SetDriver(__rv,107000,true)
	return _p
} 
//QQuaternion::QQuaternion(QVector4D const&)	
func NewQuaternionWithVector4d(vector *QVector4D) *QQuaternion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),107000,107103,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QQuaternion{}
	_p.SetDriver(__rv,107000,true)
	return _p
} 
//QQuaternion::QQuaternion(double,QVector3D const&)	
func NewQuaternionWithScalarVector3d(scalar float64,vector *QVector3D) *QQuaternion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),107000,107104,unsafe.Pointer(&scalar),Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QQuaternion{}
	_p.SetDriver(__rv,107000,true)
	return _p
} 
//QQuaternion::QQuaternion(double,double,double,double)	
func NewQuaternionWithScalarXposYposZpos(scalar float64,xpos float64,ypos float64,zpos float64) *QQuaternion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),107000,107105,unsafe.Pointer(&scalar),unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),unsafe.Pointer(&zpos),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QQuaternion{}
	_p.SetDriver(__rv,107000,true)
	return _p
} 
//QQuaternion::conjugate()
func (q *QQuaternion) Conjugate() *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::fromAxisAndAngle(QVector3D const&,double)	
func QQuaternionFromAxisAndAngleFWithVector3dAngle(axis *QVector3D,angle float64) *QQuaternion {
	var __rv uintptr
	DirectQtDrv(nil,107000,107107,Native(axis),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::fromAxisAndAngle(QVector3D const&,double)
func (q *QQuaternion) FromAxisAndAngleFWithVector3dAngle(axis *QVector3D,angle float64) *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107107,Native(axis),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::fromAxisAndAngle(double,double,double,double)	
func QQuaternionFromAxisAndAngleFWithXYZAngle(x float64,y float64,z float64,angle float64) *QQuaternion {
	var __rv uintptr
	DirectQtDrv(nil,107000,107108,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&z),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::fromAxisAndAngle(double,double,double,double)
func (q *QQuaternion) FromAxisAndAngleFWithXYZAngle(x float64,y float64,z float64,angle float64) *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107108,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&z),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::isIdentity()
func (q *QQuaternion) IsIdentity() bool {
	var __rv bool
	q.Drv(107000,107109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::isNull()
func (q *QQuaternion) IsNull() bool {
	var __rv bool
	q.Drv(107000,107110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::length()
func (q *QQuaternion) Length() float64 {
	var __rv float64
	q.Drv(107000,107111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::lengthSquared()
func (q *QQuaternion) LengthSquared() float64 {
	var __rv float64
	q.Drv(107000,107112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::nlerp(QQuaternion const&,QQuaternion const&,double)	
func QQuaternionNlerp(q1 *QQuaternion,q2 *QQuaternion,t float64) *QQuaternion {
	var __rv uintptr
	DirectQtDrv(nil,107000,107113,Native(q1),Native(q2),unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::nlerp(QQuaternion const&,QQuaternion const&,double)
func (q *QQuaternion) Nlerp(q1 *QQuaternion,q2 *QQuaternion,t float64) *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107113,Native(q1),Native(q2),unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::normalize()
func (q *QQuaternion) Normalize()  {
	q.Drv(107000,107114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::normalized()
func (q *QQuaternion) Normalized() *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::rotatedVector(QVector3D const&)
func (q *QQuaternion) RotatedVector(vector *QVector3D) *QVector3D {
	var __rv uintptr
	q.Drv(107000,107116,Native(vector),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QQuaternion::scalar()
func (q *QQuaternion) Scalar() float64 {
	var __rv float64
	q.Drv(107000,107117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::setScalar(double)
func (q *QQuaternion) SetScalar(scalar float64)  {
	q.Drv(107000,107118,unsafe.Pointer(&scalar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::setVector(QVector3D const&)
func (q *QQuaternion) SetVector(vector *QVector3D)  {
	q.Drv(107000,107119,Native(vector),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::setVector(double,double,double)
func (q *QQuaternion) SetVectorFWithXYZ(x float64,y float64,z float64)  {
	q.Drv(107000,107120,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&z),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::setX(double)
func (q *QQuaternion) SetX(x float64)  {
	q.Drv(107000,107121,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::setY(double)
func (q *QQuaternion) SetY(y float64)  {
	q.Drv(107000,107122,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::setZ(double)
func (q *QQuaternion) SetZ(z float64)  {
	q.Drv(107000,107123,unsafe.Pointer(&z),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QQuaternion::slerp(QQuaternion const&,QQuaternion const&,double)	
func QQuaternionSlerp(q1 *QQuaternion,q2 *QQuaternion,t float64) *QQuaternion {
	var __rv uintptr
	DirectQtDrv(nil,107000,107124,Native(q1),Native(q2),unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::slerp(QQuaternion const&,QQuaternion const&,double)
func (q *QQuaternion) Slerp(q1 *QQuaternion,q2 *QQuaternion,t float64) *QQuaternion {
	var __rv uintptr
	q.Drv(107000,107124,Native(q1),Native(q2),unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QQuaternion{}
	_rp.SetDriver(__rv,107000,true)
	return _rp
}	
//QQuaternion::toVector4D()
func (q *QQuaternion) ToVector4D() *QVector4D {
	var __rv uintptr
	q.Drv(107000,107125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector4D{}
	_rp.SetDriver(__rv,186000,true)
	return _rp
}	
//QQuaternion::vector()
func (q *QQuaternion) Vector() *QVector3D {
	var __rv uintptr
	q.Drv(107000,107126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVector3D{}
	_rp.SetDriver(__rv,185000,true)
	return _rp
}	
//QQuaternion::x()
func (q *QQuaternion) X() float64 {
	var __rv float64
	q.Drv(107000,107127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::y()
func (q *QQuaternion) Y() float64 {
	var __rv float64
	q.Drv(107000,107128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QQuaternion::z()
func (q *QQuaternion) Z() float64 {
	var __rv float64
	q.Drv(107000,107129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
