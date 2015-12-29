// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTransform_TransformationType - QTransform::TransformationType
type QTransform_TransformationType uint32
const (
	QTransform_TxNone QTransform_TransformationType = 0x00
	QTransform_TxTranslate QTransform_TransformationType = 0x01
	QTransform_TxScale QTransform_TransformationType = 0x02
	QTransform_TxRotate QTransform_TransformationType = 0x04
	QTransform_TxShear QTransform_TransformationType = 0x08
	QTransform_TxProject QTransform_TransformationType = 0x10
)
//struct QTransform : QTransform
type QTransform struct {
	BaseDrv
}
//QTransform::QTransform()	
func NewTransform() *QTransform {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),176000,176102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTransform{}
	_p.SetDriver(__rv,176000,true)
	return _p
} 
//QTransform::QTransform(QMatrix const&)	
func NewTransformWithMtx(mtx *QMatrix) *QTransform {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),176000,176103,Native(mtx),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTransform{}
	_p.SetDriver(__rv,176000,true)
	return _p
} 
//QTransform::QTransform(Qt::Initialization)	
func NewTransformWithInitialization(value Qt_Initialization) *QTransform {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),176000,176104,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTransform{}
	_p.SetDriver(__rv,176000,true)
	return _p
} 
//QTransform::QTransform(double,double,double,double,double,double)	
func NewTransformWithH11H12H21H22DxDy(h11 float64,h12 float64,h21 float64,h22 float64,dx float64,dy float64) *QTransform {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),176000,176105,unsafe.Pointer(&h11),unsafe.Pointer(&h12),unsafe.Pointer(&h21),unsafe.Pointer(&h22),unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTransform{}
	_p.SetDriver(__rv,176000,true)
	return _p
} 
//QTransform::QTransform(double,double,double,double,double,double,double,double,double)	
func NewTransformWithH11H12H13H21H22H23H31H32H33(h11 float64,h12 float64,h13 float64,h21 float64,h22 float64,h23 float64,h31 float64,h32 float64,h33 float64) *QTransform {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),176000,176106,unsafe.Pointer(&h11),unsafe.Pointer(&h12),unsafe.Pointer(&h13),unsafe.Pointer(&h21),unsafe.Pointer(&h22),unsafe.Pointer(&h23),unsafe.Pointer(&h31),unsafe.Pointer(&h32),unsafe.Pointer(&h33),nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTransform{}
	_p.SetDriver(__rv,176000,true)
	return _p
} 
//QTransform::adjoint()
func (q *QTransform) Adjoint() *QTransform {
	var __rv uintptr
	q.Drv(176000,176107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::determinant()
func (q *QTransform) Determinant() float64 {
	var __rv float64
	q.Drv(176000,176108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::dx()
func (q *QTransform) Dx() float64 {
	var __rv float64
	q.Drv(176000,176109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::dy()
func (q *QTransform) Dy() float64 {
	var __rv float64
	q.Drv(176000,176110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::fromScale(double,double)	
func QTransformFromScale(dx float64,dy float64) *QTransform {
	var __rv uintptr
	DirectQtDrv(nil,176000,176111,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::fromScale(double,double)
func (q *QTransform) FromScale(dx float64,dy float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176111,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::fromTranslate(double,double)	
func QTransformFromTranslate(dx float64,dy float64) *QTransform {
	var __rv uintptr
	DirectQtDrv(nil,176000,176112,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::fromTranslate(double,double)
func (q *QTransform) FromTranslate(dx float64,dy float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176112,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::inverted(bool*)
func (q *QTransform) Inverted(invertible *bool) *QTransform {
	var __rv uintptr
	q.Drv(176000,176113,unsafe.Pointer(&invertible),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::isAffine()
func (q *QTransform) IsAffine() bool {
	var __rv bool
	q.Drv(176000,176114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::isIdentity()
func (q *QTransform) IsIdentity() bool {
	var __rv bool
	q.Drv(176000,176115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::isInvertible()
func (q *QTransform) IsInvertible() bool {
	var __rv bool
	q.Drv(176000,176116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::isRotating()
func (q *QTransform) IsRotating() bool {
	var __rv bool
	q.Drv(176000,176117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::isScaling()
func (q *QTransform) IsScaling() bool {
	var __rv bool
	q.Drv(176000,176118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::isTranslating()
func (q *QTransform) IsTranslating() bool {
	var __rv bool
	q.Drv(176000,176119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m11()
func (q *QTransform) M11() float64 {
	var __rv float64
	q.Drv(176000,176120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m12()
func (q *QTransform) M12() float64 {
	var __rv float64
	q.Drv(176000,176121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m13()
func (q *QTransform) M13() float64 {
	var __rv float64
	q.Drv(176000,176122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m21()
func (q *QTransform) M21() float64 {
	var __rv float64
	q.Drv(176000,176123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m22()
func (q *QTransform) M22() float64 {
	var __rv float64
	q.Drv(176000,176124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m23()
func (q *QTransform) M23() float64 {
	var __rv float64
	q.Drv(176000,176125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m31()
func (q *QTransform) M31() float64 {
	var __rv float64
	q.Drv(176000,176126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m32()
func (q *QTransform) M32() float64 {
	var __rv float64
	q.Drv(176000,176127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::m33()
func (q *QTransform) M33() float64 {
	var __rv float64
	q.Drv(176000,176128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::map(QLine const&)
func (q *QTransform) Map(l *QLine) *QLine {
	var __rv uintptr
	q.Drv(176000,176129,Native(l),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLine{}
	_rp.SetDriver(__rv,68000,true)
	return _rp
}	
//QTransform::map(QLineF const&)
func (q *QTransform) MapFWithLinef(l *QLineF) *QLineF {
	var __rv uintptr
	q.Drv(176000,176130,Native(l),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QTransform::map(QPainterPath const&)
func (q *QTransform) MapWithPainterpath(p *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(176000,176131,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QTransform::map(QPoint const&)
func (q *QTransform) MapWithPoint(p *QPoint) *QPoint {
	var __rv uintptr
	q.Drv(176000,176132,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QTransform::map(QPointF const&)
func (q *QTransform) MapFWithPointf(p *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(176000,176133,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTransform::map(QPolygon const&)
func (q *QTransform) MapWithPolygon(a *QPolygon) *QPolygon {
	var __rv uintptr
	q.Drv(176000,176134,Native(a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QTransform::map(QPolygonF const&)
func (q *QTransform) MapFWithPolygonf(a *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(176000,176135,Native(a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QTransform::map(QRegion const&)
func (q *QTransform) MapWithRegion(r *QRegion) *QRegion {
	var __rv uintptr
	q.Drv(176000,176136,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QTransform::map(double,double,double*,double*)
func (q *QTransform) MapFWithXYTxTy(x float64,y float64,tx *float64,ty *float64)  {
	q.Drv(176000,176137,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&tx),unsafe.Pointer(&ty),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTransform::map(int,int,int*,int*)
func (q *QTransform) MapWithXYTxTy(x int,y int,tx *int,ty *int)  {
	q.Drv(176000,176138,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&tx),unsafe.Pointer(&ty),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTransform::mapRect(QRect const&)
func (q *QTransform) MapRect(value *QRect) *QRect {
	var __rv uintptr
	q.Drv(176000,176139,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTransform::mapRect(QRectF const&)
func (q *QTransform) MapRectFWithRectf(value *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(176000,176140,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTransform::mapToPolygon(QRect const&)
func (q *QTransform) MapToPolygon(r *QRect) *QPolygon {
	var __rv uintptr
	q.Drv(176000,176141,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QTransform::quadToQuad(QPolygonF const&,QPolygonF const&,QTransform&)	
func QTransformQuadToQuad(one *QPolygonF,two *QPolygonF,result *QTransform) bool {
	var __rv bool
	DirectQtDrv(nil,176000,176142,Native(one),Native(two),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::quadToQuad(QPolygonF const&,QPolygonF const&,QTransform&)
func (q *QTransform) QuadToQuad(one *QPolygonF,two *QPolygonF,result *QTransform) bool {
	var __rv bool
	q.Drv(176000,176142,Native(one),Native(two),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::quadToSquare(QPolygonF const&,QTransform&)	
func QTransformQuadToSquare(quad *QPolygonF,result *QTransform) bool {
	var __rv bool
	DirectQtDrv(nil,176000,176143,Native(quad),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::quadToSquare(QPolygonF const&,QTransform&)
func (q *QTransform) QuadToSquare(quad *QPolygonF,result *QTransform) bool {
	var __rv bool
	q.Drv(176000,176143,Native(quad),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::reset()
func (q *QTransform) Reset()  {
	q.Drv(176000,176144,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTransform::rotate(double)
func (q *QTransform) Rotate(a float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176145,unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::rotate(double,Qt::Axis)
func (q *QTransform) RotateFWithFloat64Axis(a float64,axis Qt_Axis) *QTransform {
	var __rv uintptr
	q.Drv(176000,176146,unsafe.Pointer(&a),unsafe.Pointer(&axis),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::rotateRadians(double)
func (q *QTransform) RotateRadians(a float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176147,unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::rotateRadians(double,Qt::Axis)
func (q *QTransform) RotateRadiansFWithFloat64Axis(a float64,axis Qt_Axis) *QTransform {
	var __rv uintptr
	q.Drv(176000,176148,unsafe.Pointer(&a),unsafe.Pointer(&axis),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::scale(double,double)
func (q *QTransform) Scale(sx float64,sy float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176149,unsafe.Pointer(&sx),unsafe.Pointer(&sy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::setMatrix(double,double,double,double,double,double,double,double,double)
func (q *QTransform) SetMatrix(m11 float64,m12 float64,m13 float64,m21 float64,m22 float64,m23 float64,m31 float64,m32 float64,m33 float64)  {
	q.Drv(176000,176150,unsafe.Pointer(&m11),unsafe.Pointer(&m12),unsafe.Pointer(&m13),unsafe.Pointer(&m21),unsafe.Pointer(&m22),unsafe.Pointer(&m23),unsafe.Pointer(&m31),unsafe.Pointer(&m32),unsafe.Pointer(&m33),nil,nil,nil)
}	
//QTransform::shear(double,double)
func (q *QTransform) Shear(sh float64,sv float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176151,unsafe.Pointer(&sh),unsafe.Pointer(&sv),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::squareToQuad(QPolygonF const&,QTransform&)	
func QTransformSquareToQuad(square *QPolygonF,result *QTransform) bool {
	var __rv bool
	DirectQtDrv(nil,176000,176152,Native(square),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::squareToQuad(QPolygonF const&,QTransform&)
func (q *QTransform) SquareToQuad(square *QPolygonF,result *QTransform) bool {
	var __rv bool
	q.Drv(176000,176152,Native(square),Native(result),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTransform::toAffine()
func (q *QTransform) ToAffine() *QMatrix {
	var __rv uintptr
	q.Drv(176000,176153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QTransform::translate(double,double)
func (q *QTransform) Translate(dx float64,dy float64) *QTransform {
	var __rv uintptr
	q.Drv(176000,176154,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::transposed()
func (q *QTransform) Transposed() *QTransform {
	var __rv uintptr
	q.Drv(176000,176155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QTransform::type()
func (q *QTransform) Type() QTransform_TransformationType {
	var __rv QTransform_TransformationType
	q.Drv(176000,176156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
