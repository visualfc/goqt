// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSwipeGesture_SwipeDirection - QSwipeGesture::SwipeDirection
type QSwipeGesture_SwipeDirection uint32
const (
	QSwipeGesture_NoDirection QSwipeGesture_SwipeDirection = 0
	QSwipeGesture_Left QSwipeGesture_SwipeDirection = 1
	QSwipeGesture_Right QSwipeGesture_SwipeDirection = 2
	QSwipeGesture_Up QSwipeGesture_SwipeDirection = 3
	QSwipeGesture_Down QSwipeGesture_SwipeDirection = 4
)
//struct QSwipeGesture : QSwipeGesture
type QSwipeGesture struct {
	QGesture
}
//QSwipeGesture::QSwipeGesture()	
func NewSwipeGesture() *QSwipeGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),359000,359102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSwipeGesture{}
	_p.SetDriver(__rv,359000,false)
	return _p
} 
//QSwipeGesture::QSwipeGesture(QObject*)	
func NewSwipeGestureWithParent(parent QObjectInterface) *QSwipeGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),359000,359103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSwipeGesture{}
	_p.SetDriver(__rv,359000,false)
	return _p
} 
//QSwipeGesture::horizontalDirection()
func (q *QSwipeGesture) HorizontalDirection() QSwipeGesture_SwipeDirection {
	var __rv QSwipeGesture_SwipeDirection
	q.Drv(359000,359104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSwipeGesture::setSwipeAngle(double)
func (q *QSwipeGesture) SetSwipeAngle(value float64)  {
	q.Drv(359000,359105,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSwipeGesture::swipeAngle()
func (q *QSwipeGesture) SwipeAngle() float64 {
	var __rv float64
	q.Drv(359000,359106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSwipeGesture::verticalDirection()
func (q *QSwipeGesture) VerticalDirection() QSwipeGesture_SwipeDirection {
	var __rv QSwipeGesture_SwipeDirection
	q.Drv(359000,359107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
