// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPolygonF : QPolygonF
type QPolygonF struct {
	BaseDrv
}
//QPolygonF::QPolygonF()	
func NewPolygonF() *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::QPolygonF(QPolygon const&)	
func NewPolygonFWithPolygon(a *QPolygon) *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102103,Native(a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::QPolygonF(QPolygonF const&)	
func NewPolygonFCopy(a *QPolygonF) *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102104,Native(a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::QPolygonF(QRectF const&)	
func NewPolygonFWithRectf(r *QRectF) *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102105,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::QPolygonF(QVector<QPointF> const&)	
func NewPolygonFWithPointfarray(v []*QPointF) *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102106,unsafe.Pointer(&v),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::QPolygonF(int)	
func NewPolygonFWithSize(size int) *QPolygonF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),102000,102107,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPolygonF{}
	_p.SetDriver(__rv,102000,true)
	return _p
} 
//QPolygonF::boundingRect()
func (q *QPolygonF) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(102000,102108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPolygonF::containsPoint(QPointF const&,Qt::FillRule)
func (q *QPolygonF) ContainsPoint(pt *QPointF,fillRule Qt_FillRule) bool {
	var __rv bool
	q.Drv(102000,102109,Native(pt),unsafe.Pointer(&fillRule),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPolygonF::intersected(QPolygonF const&)
func (q *QPolygonF) Intersected(r *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(102000,102110,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPolygonF::isClosed()
func (q *QPolygonF) IsClosed() bool {
	var __rv bool
	q.Drv(102000,102111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPolygonF::subtracted(QPolygonF const&)
func (q *QPolygonF) Subtracted(r *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(102000,102112,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPolygonF::toPolygon()
func (q *QPolygonF) ToPolygon() *QPolygon {
	var __rv uintptr
	q.Drv(102000,102113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QPolygonF::translate(QPointF const&)
func (q *QPolygonF) Translate(offset *QPointF)  {
	q.Drv(102000,102114,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygonF::translate(double,double)
func (q *QPolygonF) TranslateFWithDxDy(dx float64,dy float64)  {
	q.Drv(102000,102115,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPolygonF::translated(QPointF const&)
func (q *QPolygonF) Translated(offset *QPointF) *QPolygonF {
	var __rv uintptr
	q.Drv(102000,102116,Native(offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPolygonF::translated(double,double)
func (q *QPolygonF) TranslatedFWithDxDy(dx float64,dy float64) *QPolygonF {
	var __rv uintptr
	q.Drv(102000,102117,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPolygonF::united(QPolygonF const&)
func (q *QPolygonF) United(r *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(102000,102118,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
