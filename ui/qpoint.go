// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPoint : QPoint
type QPoint struct {
	BaseDrv
}
//QPoint::QPoint()	
func NewPoint() *QPoint {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),99000,99102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPoint{}
	_p.SetDriver(__rv,99000,true)
	return _p
} 
//QPoint::QPoint(int,int)	
func NewPointWithXposYpos(xpos int,ypos int) *QPoint {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),99000,99103,unsafe.Pointer(&xpos),unsafe.Pointer(&ypos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPoint{}
	_p.SetDriver(__rv,99000,true)
	return _p
} 
//QPoint::isNull()
func (q *QPoint) IsNull() bool {
	var __rv bool
	q.Drv(99000,99104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPoint::manhattanLength()
func (q *QPoint) ManhattanLength() int {
	var __rv int
	q.Drv(99000,99105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPoint::rx()
func (q *QPoint) Rx() *int {
	var __rv *int
	q.Drv(99000,99106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPoint::ry()
func (q *QPoint) Ry() *int {
	var __rv *int
	q.Drv(99000,99107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPoint::setX(int)
func (q *QPoint) SetX(x int)  {
	q.Drv(99000,99108,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPoint::setY(int)
func (q *QPoint) SetY(y int)  {
	q.Drv(99000,99109,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPoint::x()
func (q *QPoint) X() int {
	var __rv int
	q.Drv(99000,99110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPoint::y()
func (q *QPoint) Y() int {
	var __rv int
	q.Drv(99000,99111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
