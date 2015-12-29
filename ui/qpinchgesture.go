// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPinchGesture_ChangeFlag - QPinchGesture::ChangeFlag
type QPinchGesture_ChangeFlag uint32
const (
	QPinchGesture_ScaleFactorChanged QPinchGesture_ChangeFlag = 0x1
	QPinchGesture_RotationAngleChanged QPinchGesture_ChangeFlag = 0x2
	QPinchGesture_CenterPointChanged QPinchGesture_ChangeFlag = 0x4
)
//struct QPinchGesture : QPinchGesture
type QPinchGesture struct {
	QGesture
}
//QPinchGesture::QPinchGesture()	
func NewPinchGesture() *QPinchGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),319000,319102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPinchGesture{}
	_p.SetDriver(__rv,319000,false)
	return _p
} 
//QPinchGesture::QPinchGesture(QObject*)	
func NewPinchGestureWithParent(parent QObjectInterface) *QPinchGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),319000,319103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPinchGesture{}
	_p.SetDriver(__rv,319000,false)
	return _p
} 
//QPinchGesture::centerPoint()
func (q *QPinchGesture) CenterPoint() *QPointF {
	var __rv uintptr
	q.Drv(319000,319104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPinchGesture::changeFlags()
func (q *QPinchGesture) ChangeFlags() QPinchGesture_ChangeFlag {
	var __rv QPinchGesture_ChangeFlag
	q.Drv(319000,319105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::lastCenterPoint()
func (q *QPinchGesture) LastCenterPoint() *QPointF {
	var __rv uintptr
	q.Drv(319000,319106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPinchGesture::lastRotationAngle()
func (q *QPinchGesture) LastRotationAngle() float64 {
	var __rv float64
	q.Drv(319000,319107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::lastScaleFactor()
func (q *QPinchGesture) LastScaleFactor() float64 {
	var __rv float64
	q.Drv(319000,319108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::rotationAngle()
func (q *QPinchGesture) RotationAngle() float64 {
	var __rv float64
	q.Drv(319000,319109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::scaleFactor()
func (q *QPinchGesture) ScaleFactor() float64 {
	var __rv float64
	q.Drv(319000,319110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::setCenterPoint(QPointF const&)
func (q *QPinchGesture) SetCenterPoint(value *QPointF)  {
	q.Drv(319000,319111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setChangeFlags(QFlags<QPinchGesture::ChangeFlag>)
func (q *QPinchGesture) SetChangeFlags(value QPinchGesture_ChangeFlag)  {
	q.Drv(319000,319112,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setLastCenterPoint(QPointF const&)
func (q *QPinchGesture) SetLastCenterPoint(value *QPointF)  {
	q.Drv(319000,319113,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setLastRotationAngle(double)
func (q *QPinchGesture) SetLastRotationAngle(value float64)  {
	q.Drv(319000,319114,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setLastScaleFactor(double)
func (q *QPinchGesture) SetLastScaleFactor(value float64)  {
	q.Drv(319000,319115,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setRotationAngle(double)
func (q *QPinchGesture) SetRotationAngle(value float64)  {
	q.Drv(319000,319116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setScaleFactor(double)
func (q *QPinchGesture) SetScaleFactor(value float64)  {
	q.Drv(319000,319117,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setStartCenterPoint(QPointF const&)
func (q *QPinchGesture) SetStartCenterPoint(value *QPointF)  {
	q.Drv(319000,319118,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setTotalChangeFlags(QFlags<QPinchGesture::ChangeFlag>)
func (q *QPinchGesture) SetTotalChangeFlags(value QPinchGesture_ChangeFlag)  {
	q.Drv(319000,319119,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setTotalRotationAngle(double)
func (q *QPinchGesture) SetTotalRotationAngle(value float64)  {
	q.Drv(319000,319120,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::setTotalScaleFactor(double)
func (q *QPinchGesture) SetTotalScaleFactor(value float64)  {
	q.Drv(319000,319121,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPinchGesture::startCenterPoint()
func (q *QPinchGesture) StartCenterPoint() *QPointF {
	var __rv uintptr
	q.Drv(319000,319122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPinchGesture::totalChangeFlags()
func (q *QPinchGesture) TotalChangeFlags() QPinchGesture_ChangeFlag {
	var __rv QPinchGesture_ChangeFlag
	q.Drv(319000,319123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::totalRotationAngle()
func (q *QPinchGesture) TotalRotationAngle() float64 {
	var __rv float64
	q.Drv(319000,319124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPinchGesture::totalScaleFactor()
func (q *QPinchGesture) TotalScaleFactor() float64 {
	var __rv float64
	q.Drv(319000,319125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
