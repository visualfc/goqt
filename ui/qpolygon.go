// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPolygon : QPolygon
type QPolygon struct {
	BaseDrv
}
//QPolygon::QPolygon()	
func NewPolygon() *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::QPolygon(QPolygon const&)	
func NewPolygonCopy(a *QPolygon) *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101103,Native(a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::QPolygon(QVector<QPoint> const&)	
func NewPolygonWithPointarray(v []*QPoint) *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101104,unsafe.Pointer(&v),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::QPolygon(int)	
func NewPolygonWithSize(size int) *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101105,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::QPolygon(QRect const&,bool)	
func NewPolygonWithRectClosed(r *QRect,closed bool) *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101106,Native(r),unsafe.Pointer(&closed),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::QPolygon(int,int const*)	
func NewPolygonWithNpointsPoints(nPoints int,points *int) *QPolygon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),101000,101107,unsafe.Pointer(&nPoints),unsafe.Pointer(&points),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygon{}
	_p.SetDriver(__rv,101000,true)
	return _p
} 
//QPolygon::boundingRect()
func (q *QPolygon) BoundingRect() *QRect {
	var __rv uintptr
	q.Drv(101000,101108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPolygon::containsPoint(QPoint const&,Qt::FillRule)
func (q *QPolygon) ContainsPoint(pt *QPoint,fillRule Qt_FillRule) bool {
	var __rv bool
	q.Drv(101000,101109,Native(pt),unsafe.Pointer(&fillRule),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPolygon::intersected(QPolygon const&)
func (q *QPolygon) Intersected(r *QPolygon) *QPolygon {
	var __rv uintptr
	q.Drv(101000,101110,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QPolygon::point(int)
func (q *QPolygon) Point(i int) *QPoint {
	var __rv uintptr
	q.Drv(101000,101111,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QPolygon::point(int,int*,int*)
func (q *QPolygon) PointWithIXY(i int,x *int,y *int)  {
	q.Drv(101000,101112,unsafe.Pointer(&i),unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygon::setPoint(int,QPoint const&)
func (q *QPolygon) SetPointWithIndexPoint(index int,p *QPoint)  {
	q.Drv(101000,101113,unsafe.Pointer(&index),Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygon::setPoint(int,int,int)
func (q *QPolygon) SetPointWithIndexXY(index int,x int,y int)  {
	q.Drv(101000,101114,unsafe.Pointer(&index),unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygon::subtracted(QPolygon const&)
func (q *QPolygon) Subtracted(r *QPolygon) *QPolygon {
	var __rv uintptr
	q.Drv(101000,101115,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QPolygon::translate(QPoint const&)
func (q *QPolygon) Translate(offset *QPoint)  {
	q.Drv(101000,101116,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygon::translate(int,int)
func (q *QPolygon) TranslateWithDxDy(dx int,dy int)  {
	q.Drv(101000,101117,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygon::translated(QPoint const&)
func (q *QPolygon) Translated(offset *QPoint) *QPolygon {
	var __rv uintptr
	q.Drv(101000,101118,Native(offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QPolygon::translated(int,int)
func (q *QPolygon) TranslatedWithDxDy(dx int,dy int) *QPolygon {
	var __rv uintptr
	q.Drv(101000,101119,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QPolygon::united(QPolygon const&)
func (q *QPolygon) United(r *QPolygon) *QPolygon {
	var __rv uintptr
	q.Drv(101000,101120,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
