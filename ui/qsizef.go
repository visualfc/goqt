// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSizeF : QSizeF
type QSizeF struct {
	BaseDrv
}
//QSizeF::QSizeF()	
func NewSizeF() *QSizeF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),120000,120102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizeF{}
	_p.SetDriver(__rv,120000,true)
	return _p
} 
//QSizeF::QSizeF(QSize const&)	
func NewSizeFWithSize(sz *QSize) *QSizeF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),120000,120103,Native(sz),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizeF{}
	_p.SetDriver(__rv,120000,true)
	return _p
} 
//QSizeF::QSizeF(QSizeF const&)	
func NewSizeFCopy(other *QSizeF) *QSizeF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),120000,120104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizeF{}
	_p.SetDriver(__rv,120000,true)
	return _p
} 
//QSizeF::QSizeF(double,double)	
func NewSizeFWithWidthHeight(w float64,h float64) *QSizeF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),120000,120105,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSizeF{}
	_p.SetDriver(__rv,120000,true)
	return _p
} 
//QSizeF::boundedTo(QSizeF const&)
func (q *QSizeF) BoundedTo(value *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(120000,120106,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QSizeF::expandedTo(QSizeF const&)
func (q *QSizeF) ExpandedTo(value *QSizeF) *QSizeF {
	var __rv uintptr
	q.Drv(120000,120107,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QSizeF::height()
func (q *QSizeF) Height() float64 {
	var __rv float64
	q.Drv(120000,120108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::isEmpty()
func (q *QSizeF) IsEmpty() bool {
	var __rv bool
	q.Drv(120000,120109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::isNull()
func (q *QSizeF) IsNull() bool {
	var __rv bool
	q.Drv(120000,120110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::isValid()
func (q *QSizeF) IsValid() bool {
	var __rv bool
	q.Drv(120000,120111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::rheight()
func (q *QSizeF) Rheight() *float64 {
	var __rv *float64
	q.Drv(120000,120112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::rwidth()
func (q *QSizeF) Rwidth() *float64 {
	var __rv *float64
	q.Drv(120000,120113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSizeF::scale(QSizeF const&,Qt::AspectRatioMode)
func (q *QSizeF) ScaleFWithSizefMode(s *QSizeF,mode Qt_AspectRatioMode)  {
	q.Drv(120000,120114,Native(s),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeF::scale(double,double,Qt::AspectRatioMode)
func (q *QSizeF) ScaleFWithWidthHeightMode(w float64,h float64,mode Qt_AspectRatioMode)  {
	q.Drv(120000,120115,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeF::setHeight(double)
func (q *QSizeF) SetHeight(h float64)  {
	q.Drv(120000,120116,unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeF::setWidth(double)
func (q *QSizeF) SetWidth(w float64)  {
	q.Drv(120000,120117,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeF::toSize()
func (q *QSizeF) ToSize() *QSize {
	var __rv uintptr
	q.Drv(120000,120118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSizeF::transpose()
func (q *QSizeF) Transpose()  {
	q.Drv(120000,120119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSizeF::width()
func (q *QSizeF) Width() float64 {
	var __rv float64
	q.Drv(120000,120120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
