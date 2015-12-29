// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTapAndHoldGesture : QTapAndHoldGesture
type QTapAndHoldGesture struct {
	QGesture
}
//QTapAndHoldGesture::QTapAndHoldGesture()	
func NewTapAndHoldGesture() *QTapAndHoldGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),367000,367102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTapAndHoldGesture{}
	_p.SetDriver(__rv,367000,false)
	return _p
} 
//QTapAndHoldGesture::QTapAndHoldGesture(QObject*)	
func NewTapAndHoldGestureWithParent(parent QObjectInterface) *QTapAndHoldGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),367000,367103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTapAndHoldGesture{}
	_p.SetDriver(__rv,367000,false)
	return _p
} 
//QTapAndHoldGesture::position()
func (q *QTapAndHoldGesture) Position() *QPointF {
	var __rv uintptr
	q.Drv(367000,367104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTapAndHoldGesture::setPosition(QPointF const&)
func (q *QTapAndHoldGesture) SetPosition(pos *QPointF)  {
	q.Drv(367000,367105,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTapAndHoldGesture::setTimeout(int)	
func QTapAndHoldGestureSetTimeout(msecs int)  {
	DirectQtDrv(nil,367000,367106,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTapAndHoldGesture::setTimeout(int)
func (q *QTapAndHoldGesture) SetTimeout(msecs int)  {
	q.Drv(367000,367106,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTapAndHoldGesture::timeout()	
func QTapAndHoldGestureTimeout() int {
	var __rv int
	DirectQtDrv(nil,367000,367107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTapAndHoldGesture::timeout()
func (q *QTapAndHoldGesture) Timeout() int {
	var __rv int
	q.Drv(367000,367107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
