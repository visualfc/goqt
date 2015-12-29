// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPanGesture : QPanGesture
type QPanGesture struct {
	QGesture
}
//QPanGesture::QPanGesture()	
func NewPanGesture() *QPanGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),316000,316102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPanGesture{}
	_p.SetDriver(__rv,316000,false)
	return _p
} 
//QPanGesture::QPanGesture(QObject*)	
func NewPanGestureWithParent(parent QObjectInterface) *QPanGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),316000,316103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPanGesture{}
	_p.SetDriver(__rv,316000,false)
	return _p
} 
//QPanGesture::acceleration()
func (q *QPanGesture) Acceleration() float64 {
	var __rv float64
	q.Drv(316000,316104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPanGesture::delta()
func (q *QPanGesture) Delta() *QPointF {
	var __rv uintptr
	q.Drv(316000,316105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPanGesture::lastOffset()
func (q *QPanGesture) LastOffset() *QPointF {
	var __rv uintptr
	q.Drv(316000,316106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPanGesture::offset()
func (q *QPanGesture) Offset() *QPointF {
	var __rv uintptr
	q.Drv(316000,316107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPanGesture::setAcceleration(double)
func (q *QPanGesture) SetAcceleration(value float64)  {
	q.Drv(316000,316108,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPanGesture::setLastOffset(QPointF const&)
func (q *QPanGesture) SetLastOffset(value *QPointF)  {
	q.Drv(316000,316109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPanGesture::setOffset(QPointF const&)
func (q *QPanGesture) SetOffset(value *QPointF)  {
	q.Drv(316000,316110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
