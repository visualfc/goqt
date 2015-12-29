// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QMargins : QMargins
type QMargins struct {
	BaseDrv
}
//QMargins::QMargins()	
func NewMargins() *QMargins {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),73000,73102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMargins{}
	_p.SetDriver(__rv,73000,true)
	return _p
} 
//QMargins::QMargins(int,int,int,int)	
func NewMarginsWithLeftTopRightBottom(left int,top int,right int,bottom int) *QMargins {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),73000,73103,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMargins{}
	_p.SetDriver(__rv,73000,true)
	return _p
} 
//QMargins::bottom()
func (q *QMargins) Bottom() int {
	var __rv int
	q.Drv(73000,73104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMargins::isNull()
func (q *QMargins) IsNull() bool {
	var __rv bool
	q.Drv(73000,73105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMargins::left()
func (q *QMargins) Left() int {
	var __rv int
	q.Drv(73000,73106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMargins::right()
func (q *QMargins) Right() int {
	var __rv int
	q.Drv(73000,73107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMargins::setBottom(int)
func (q *QMargins) SetBottom(bottom int)  {
	q.Drv(73000,73108,unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMargins::setLeft(int)
func (q *QMargins) SetLeft(left int)  {
	q.Drv(73000,73109,unsafe.Pointer(&left),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMargins::setRight(int)
func (q *QMargins) SetRight(right int)  {
	q.Drv(73000,73110,unsafe.Pointer(&right),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMargins::setTop(int)
func (q *QMargins) SetTop(top int)  {
	q.Drv(73000,73111,unsafe.Pointer(&top),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMargins::top()
func (q *QMargins) Top() int {
	var __rv int
	q.Drv(73000,73112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
