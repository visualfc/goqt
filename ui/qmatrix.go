// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMatrix : QMatrix
type QMatrix struct {
	BaseDrv
}
//QMatrix::QMatrix()	
func NewMatrix() *QMatrix {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),74000,74102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMatrix{}
	_p.SetDriver(__rv,74000,true)
	return _p
} 
//QMatrix::QMatrix(QMatrix const&)	
func NewMatrixCopy(matrix *QMatrix) *QMatrix {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),74000,74103,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMatrix{}
	_p.SetDriver(__rv,74000,true)
	return _p
} 
//QMatrix::QMatrix(Qt::Initialization)	
func NewMatrixWithInitialization(value Qt_Initialization) *QMatrix {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),74000,74104,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMatrix{}
	_p.SetDriver(__rv,74000,true)
	return _p
} 
//QMatrix::QMatrix(double,double,double,double,double,double)	
func NewMatrixWithM11M12M21M22DxDy(m11 float64,m12 float64,m21 float64,m22 float64,dx float64,dy float64) *QMatrix {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),74000,74105,unsafe.Pointer(&m11),unsafe.Pointer(&m12),unsafe.Pointer(&m21),unsafe.Pointer(&m22),unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMatrix{}
	_p.SetDriver(__rv,74000,true)
	return _p
} 
//QMatrix::determinant()
func (q *QMatrix) Determinant() float64 {
	var __rv float64
	q.Drv(74000,74106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::dx()
func (q *QMatrix) Dx() float64 {
	var __rv float64
	q.Drv(74000,74107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::dy()
func (q *QMatrix) Dy() float64 {
	var __rv float64
	q.Drv(74000,74108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::inverted(bool*)
func (q *QMatrix) Inverted(invertible *bool) *QMatrix {
	var __rv uintptr
	q.Drv(74000,74109,unsafe.Pointer(&invertible),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QMatrix::isIdentity()
func (q *QMatrix) IsIdentity() bool {
	var __rv bool
	q.Drv(74000,74110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::isInvertible()
func (q *QMatrix) IsInvertible() bool {
	var __rv bool
	q.Drv(74000,74111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::m11()
func (q *QMatrix) M11() float64 {
	var __rv float64
	q.Drv(74000,74112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::m12()
func (q *QMatrix) M12() float64 {
	var __rv float64
	q.Drv(74000,74113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::m21()
func (q *QMatrix) M21() float64 {
	var __rv float64
	q.Drv(74000,74114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::m22()
func (q *QMatrix) M22() float64 {
	var __rv float64
	q.Drv(74000,74115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMatrix::mapRect(QRect const&)
func (q *QMatrix) MapRect(value *QRect) *QRect {
	var __rv uintptr
	q.Drv(74000,74116,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QMatrix::mapRect(QRectF const&)
func (q *QMatrix) MapRectFWithRectf(value *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(74000,74117,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QMatrix::mapToPolygon(QRect const&)
func (q *QMatrix) MapToPolygon(r *QRect) *QPolygon {
	var __rv uintptr
	q.Drv(74000,74118,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QMatrix::reset()
func (q *QMatrix) Reset()  {
	q.Drv(74000,74119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMatrix::rotate(double)
func (q *QMatrix) Rotate(a float64) *QMatrix {
	var __rv uintptr
	q.Drv(74000,74120,unsafe.Pointer(&a),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QMatrix::scale(double,double)
func (q *QMatrix) Scale(sx float64,sy float64) *QMatrix {
	var __rv uintptr
	q.Drv(74000,74121,unsafe.Pointer(&sx),unsafe.Pointer(&sy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QMatrix::setMatrix(double,double,double,double,double,double)
func (q *QMatrix) SetMatrix(m11 float64,m12 float64,m21 float64,m22 float64,dx float64,dy float64)  {
	q.Drv(74000,74122,unsafe.Pointer(&m11),unsafe.Pointer(&m12),unsafe.Pointer(&m21),unsafe.Pointer(&m22),unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil)
}	
//QMatrix::shear(double,double)
func (q *QMatrix) Shear(sh float64,sv float64) *QMatrix {
	var __rv uintptr
	q.Drv(74000,74123,unsafe.Pointer(&sh),unsafe.Pointer(&sv),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QMatrix::translate(double,double)
func (q *QMatrix) Translate(dx float64,dy float64) *QMatrix {
	var __rv uintptr
	q.Drv(74000,74124,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
