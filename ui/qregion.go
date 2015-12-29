// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QRegion_RegionType - QRegion::RegionType
type QRegion_RegionType uint32
const (
	QRegion_Rectangle QRegion_RegionType = 0
	QRegion_Ellipse QRegion_RegionType = 1
)
//struct QRegion : QRegion
type QRegion struct {
	BaseDrv
}
//QRegion::QRegion()	
func NewRegion() *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::QRegion(QBitmap const&)	
func NewRegionWithBitmap(bitmap *QBitmap) *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113103,Native(bitmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::QRegion(QRegion const&)	
func NewRegionCopy(region *QRegion) *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113104,Native(region),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::QRegion(QPolygon const&,Qt::FillRule)	
func NewRegionWithPolygonFillrule(pa *QPolygon,fillRule Qt_FillRule) *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113105,Native(pa),unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::QRegion(QRect const&,QRegion::RegionType)	
func NewRegionWithRectType(r *QRect,t QRegion_RegionType) *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113106,Native(r),unsafe.Pointer(&t),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::QRegion(int,int,int,int,QRegion::RegionType)	
func NewRegionWithXYWidthHeightType(x int,y int,w int,h int,t QRegion_RegionType) *QRegion {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),113000,113107,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&t),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegion{}
	_p.SetDriver(__rv,113000,true)
	return _p
} 
//QRegion::boundingRect()
func (q *QRegion) BoundingRect() *QRect {
	var __rv uintptr
	q.Drv(113000,113108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRegion::contains(QPoint const&)
func (q *QRegion) Contains(p *QPoint) bool {
	var __rv bool
	q.Drv(113000,113109,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::contains(QRect const&)
func (q *QRegion) ContainsWithRect(r *QRect) bool {
	var __rv bool
	q.Drv(113000,113110,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::intersected(QRect const&)
func (q *QRegion) Intersected(r *QRect) *QRegion {
	var __rv uintptr
	q.Drv(113000,113111,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::intersected(QRegion const&)
func (q *QRegion) IntersectedWithRegion(r *QRegion) *QRegion {
	var __rv uintptr
	q.Drv(113000,113112,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::intersects(QRect const&)
func (q *QRegion) Intersects(r *QRect) bool {
	var __rv bool
	q.Drv(113000,113113,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::intersects(QRegion const&)
func (q *QRegion) IntersectsWithRegion(r *QRegion) bool {
	var __rv bool
	q.Drv(113000,113114,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::isEmpty()
func (q *QRegion) IsEmpty() bool {
	var __rv bool
	q.Drv(113000,113115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::rectCount()
func (q *QRegion) RectCount() int {
	var __rv int
	q.Drv(113000,113116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::rects()
func (q *QRegion) Rects() []*QRect {
	var __rv []*QRect
	q.Drv(113000,113117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegion::setRects(QRect const*,int)
func (q *QRegion) SetRects(ar []*QRect)  {
	q.Drv(113000,113118,unsafe.Pointer(&ar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegion::subtracted(QRegion const&)
func (q *QRegion) Subtracted(r *QRegion) *QRegion {
	var __rv uintptr
	q.Drv(113000,113119,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::translate(QPoint const&)
func (q *QRegion) Translate(p *QPoint)  {
	q.Drv(113000,113120,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegion::translate(int,int)
func (q *QRegion) TranslateWithDxDy(dx int,dy int)  {
	q.Drv(113000,113121,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegion::translated(QPoint const&)
func (q *QRegion) Translated(p *QPoint) *QRegion {
	var __rv uintptr
	q.Drv(113000,113122,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::translated(int,int)
func (q *QRegion) TranslatedWithDxDy(dx int,dy int) *QRegion {
	var __rv uintptr
	q.Drv(113000,113123,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::united(QRect const&)
func (q *QRegion) United(r *QRect) *QRegion {
	var __rv uintptr
	q.Drv(113000,113124,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::united(QRegion const&)
func (q *QRegion) UnitedWithRegion(r *QRegion) *QRegion {
	var __rv uintptr
	q.Drv(113000,113125,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QRegion::xored(QRegion const&)
func (q *QRegion) Xored(r *QRegion) *QRegion {
	var __rv uintptr
	q.Drv(113000,113126,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
