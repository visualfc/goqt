// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGesture_GestureCancelPolicy - QGesture::GestureCancelPolicy
type QGesture_GestureCancelPolicy uint32
const (
	QGesture_CancelNone QGesture_GestureCancelPolicy = 0
	QGesture_CancelAllInContext QGesture_GestureCancelPolicy = 0
)
//struct QGesture : QGesture
type QGesture struct {
	QObject
}
//QGesture::QGesture()	
func NewGesture() *QGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),247000,247102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGesture{}
	_p.SetDriver(__rv,247000,false)
	return _p
} 
//QGesture::QGesture(QObject*)	
func NewGestureWithParent(parent QObjectInterface) *QGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),247000,247103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGesture{}
	_p.SetDriver(__rv,247000,false)
	return _p
} 
//QGesture::gestureCancelPolicy()
func (q *QGesture) GestureCancelPolicy() QGesture_GestureCancelPolicy {
	var __rv QGesture_GestureCancelPolicy
	q.Drv(247000,247104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGesture::gestureType()
func (q *QGesture) GestureType() Qt_GestureType {
	var __rv Qt_GestureType
	q.Drv(247000,247105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGesture::hasHotSpot()
func (q *QGesture) HasHotSpot() bool {
	var __rv bool
	q.Drv(247000,247106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGesture::hotSpot()
func (q *QGesture) HotSpot() *QPointF {
	var __rv uintptr
	q.Drv(247000,247107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGesture::setGestureCancelPolicy(QGesture::GestureCancelPolicy)
func (q *QGesture) SetGestureCancelPolicy(policy QGesture_GestureCancelPolicy)  {
	q.Drv(247000,247108,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGesture::setHotSpot(QPointF const&)
func (q *QGesture) SetHotSpot(value *QPointF)  {
	q.Drv(247000,247109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGesture::state()
func (q *QGesture) State() Qt_GestureState {
	var __rv Qt_GestureState
	q.Drv(247000,247110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGesture::unsetHotSpot()
func (q *QGesture) UnsetHotSpot()  {
	q.Drv(247000,247111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
