// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLineF_IntersectType - QLineF::IntersectType
type QLineF_IntersectType uint32
const (
	QLineF_NoIntersection QLineF_IntersectType = 0
	QLineF_BoundedIntersection QLineF_IntersectType = 1
	QLineF_UnboundedIntersection QLineF_IntersectType = 2
)
//struct QLineF : QLineF
type QLineF struct {
	BaseDrv
}
//QLineF::QLineF()	
func NewLineF() *QLineF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),69000,69102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineF{}
	_p.SetDriver(__rv,69000,true)
	return _p
} 
//QLineF::QLineF(QLine const&)	
func NewLineFWithLine(line *QLine) *QLineF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),69000,69103,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineF{}
	_p.SetDriver(__rv,69000,true)
	return _p
} 
//QLineF::QLineF(QLineF const&)	
func NewLineFCopy(other *QLineF) *QLineF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),69000,69104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineF{}
	_p.SetDriver(__rv,69000,true)
	return _p
} 
//QLineF::QLineF(QPointF const&,QPointF const&)	
func NewLineFWithPt1Pt2(pt1 *QPointF,pt2 *QPointF) *QLineF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),69000,69105,Native(pt1),Native(pt2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineF{}
	_p.SetDriver(__rv,69000,true)
	return _p
} 
//QLineF::QLineF(double,double,double,double)	
func NewLineFWithX1Y1X2Y2(x1 float64,y1 float64,x2 float64,y2 float64) *QLineF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),69000,69106,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineF{}
	_p.SetDriver(__rv,69000,true)
	return _p
} 
//QLineF::angle()
func (q *QLineF) Angle() float64 {
	var __rv float64
	q.Drv(69000,69107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::angle(QLineF const&)
func (q *QLineF) AngleFWithLinef(l *QLineF) float64 {
	var __rv float64
	q.Drv(69000,69108,Native(l),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::angleTo(QLineF const&)
func (q *QLineF) AngleTo(l *QLineF) float64 {
	var __rv float64
	q.Drv(69000,69109,Native(l),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::dx()
func (q *QLineF) Dx() float64 {
	var __rv float64
	q.Drv(69000,69110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::dy()
func (q *QLineF) Dy() float64 {
	var __rv float64
	q.Drv(69000,69111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::fromPolar(double,double)	
func QLineFFromPolar(length float64,angle float64) *QLineF {
	var __rv uintptr
	DirectQtDrv(nil,69000,69112,unsafe.Pointer(&length),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::fromPolar(double,double)
func (q *QLineF) FromPolar(length float64,angle float64) *QLineF {
	var __rv uintptr
	q.Drv(69000,69112,unsafe.Pointer(&length),unsafe.Pointer(&angle),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::intersect(QLineF const&,QPointF*)
func (q *QLineF) Intersect(l *QLineF,intersectionPoint *QPointF) QLineF_IntersectType {
	var __rv QLineF_IntersectType
	q.Drv(69000,69113,Native(l),Native(intersectionPoint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::isNull()
func (q *QLineF) IsNull() bool {
	var __rv bool
	q.Drv(69000,69114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::length()
func (q *QLineF) Length() float64 {
	var __rv float64
	q.Drv(69000,69115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::normalVector()
func (q *QLineF) NormalVector() *QLineF {
	var __rv uintptr
	q.Drv(69000,69116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::p1()
func (q *QLineF) P1() *QPointF {
	var __rv uintptr
	q.Drv(69000,69117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QLineF::p2()
func (q *QLineF) P2() *QPointF {
	var __rv uintptr
	q.Drv(69000,69118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QLineF::pointAt(double)
func (q *QLineF) PointAt(t float64) *QPointF {
	var __rv uintptr
	q.Drv(69000,69119,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QLineF::setAngle(double)
func (q *QLineF) SetAngle(angle float64)  {
	q.Drv(69000,69120,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::setLength(double)
func (q *QLineF) SetLength(len float64)  {
	q.Drv(69000,69121,unsafe.Pointer(&len),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::setLine(double,double,double,double)
func (q *QLineF) SetLine(x1 float64,y1 float64,x2 float64,y2 float64)  {
	q.Drv(69000,69122,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::setP1(QPointF const&)
func (q *QLineF) SetP1(p1 *QPointF)  {
	q.Drv(69000,69123,Native(p1),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::setP2(QPointF const&)
func (q *QLineF) SetP2(p2 *QPointF)  {
	q.Drv(69000,69124,Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::setPoints(QPointF const&,QPointF const&)
func (q *QLineF) SetPoints(p1 *QPointF,p2 *QPointF)  {
	q.Drv(69000,69125,Native(p1),Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::toLine()
func (q *QLineF) ToLine() *QLine {
	var __rv uintptr
	q.Drv(69000,69126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLine{}
	_rp.SetDriver(__rv,68000,true)
	return _rp
}	
//QLineF::translate(QPointF const&)
func (q *QLineF) Translate(p *QPointF)  {
	q.Drv(69000,69127,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::translate(double,double)
func (q *QLineF) TranslateFWithDxDy(dx float64,dy float64)  {
	q.Drv(69000,69128,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineF::translated(QPointF const&)
func (q *QLineF) Translated(p *QPointF) *QLineF {
	var __rv uintptr
	q.Drv(69000,69129,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::translated(double,double)
func (q *QLineF) TranslatedFWithDxDy(dx float64,dy float64) *QLineF {
	var __rv uintptr
	q.Drv(69000,69130,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::unitVector()
func (q *QLineF) UnitVector() *QLineF {
	var __rv uintptr
	q.Drv(69000,69131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineF{}
	_rp.SetDriver(__rv,69000,true)
	return _rp
}	
//QLineF::x1()
func (q *QLineF) X1() float64 {
	var __rv float64
	q.Drv(69000,69132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::x2()
func (q *QLineF) X2() float64 {
	var __rv float64
	q.Drv(69000,69133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::y1()
func (q *QLineF) Y1() float64 {
	var __rv float64
	q.Drv(69000,69134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineF::y2()
func (q *QLineF) Y2() float64 {
	var __rv float64
	q.Drv(69000,69135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
