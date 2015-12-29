// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSize : QSize
type QSize struct {
	BaseDrv
}
//QSize::QSize()	
func NewSize() *QSize {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),119000,119102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSize{}
	_p.SetDriver(__rv,119000,true)
	return _p
} 
//QSize::QSize(QSize const&)	
func NewSizeCopy(other *QSize) *QSize {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),119000,119103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSize{}
	_p.SetDriver(__rv,119000,true)
	return _p
} 
//QSize::QSize(int,int)	
func NewSizeWithWidthHeight(w int,h int) *QSize {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),119000,119104,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSize{}
	_p.SetDriver(__rv,119000,true)
	return _p
} 
//QSize::boundedTo(QSize const&)
func (q *QSize) BoundedTo(value *QSize) *QSize {
	var __rv uintptr
	q.Drv(119000,119105,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSize::expandedTo(QSize const&)
func (q *QSize) ExpandedTo(value *QSize) *QSize {
	var __rv uintptr
	q.Drv(119000,119106,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSize::height()
func (q *QSize) Height() int {
	var __rv int
	q.Drv(119000,119107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::isEmpty()
func (q *QSize) IsEmpty() bool {
	var __rv bool
	q.Drv(119000,119108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::isNull()
func (q *QSize) IsNull() bool {
	var __rv bool
	q.Drv(119000,119109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::isValid()
func (q *QSize) IsValid() bool {
	var __rv bool
	q.Drv(119000,119110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::rheight()
func (q *QSize) Rheight() *int {
	var __rv *int
	q.Drv(119000,119111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::rwidth()
func (q *QSize) Rwidth() *int {
	var __rv *int
	q.Drv(119000,119112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSize::scale(QSize const&,Qt::AspectRatioMode)
func (q *QSize) ScaleWithSizeMode(s *QSize,mode Qt_AspectRatioMode)  {
	q.Drv(119000,119113,Native(s),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSize::scale(int,int,Qt::AspectRatioMode)
func (q *QSize) ScaleWithWidthHeightMode(w int,h int,mode Qt_AspectRatioMode)  {
	q.Drv(119000,119114,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSize::setHeight(int)
func (q *QSize) SetHeight(h int)  {
	q.Drv(119000,119115,unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSize::setWidth(int)
func (q *QSize) SetWidth(w int)  {
	q.Drv(119000,119116,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSize::transpose()
func (q *QSize) Transpose()  {
	q.Drv(119000,119117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSize::width()
func (q *QSize) Width() int {
	var __rv int
	q.Drv(119000,119118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
