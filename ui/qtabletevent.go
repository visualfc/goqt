// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTabletEvent_TabletDevice - QTabletEvent::TabletDevice
type QTabletEvent_TabletDevice uint32
const (
	QTabletEvent_NoDevice QTabletEvent_TabletDevice = 0
	QTabletEvent_Puck QTabletEvent_TabletDevice = 1
	QTabletEvent_Stylus QTabletEvent_TabletDevice = 2
	QTabletEvent_Airbrush QTabletEvent_TabletDevice = 3
	QTabletEvent_FourDMouse QTabletEvent_TabletDevice = 4
	QTabletEvent_XFreeEraser QTabletEvent_TabletDevice = 5
	QTabletEvent_RotationStylus QTabletEvent_TabletDevice = 6
)
//enum QTabletEvent_PointerType - QTabletEvent::PointerType
type QTabletEvent_PointerType uint32
const (
	QTabletEvent_UnknownPointer QTabletEvent_PointerType = 0
	QTabletEvent_Pen QTabletEvent_PointerType = 1
	QTabletEvent_Cursor QTabletEvent_PointerType = 2
	QTabletEvent_Eraser QTabletEvent_PointerType = 3
)
//struct QTabletEvent : QTabletEvent
type QTabletEvent struct {
	QInputEvent
}
//QTabletEvent::device()
func (q *QTabletEvent) Device() QTabletEvent_TabletDevice {
	var __rv QTabletEvent_TabletDevice
	q.Drv(136000,136102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::globalPos()
func (q *QTabletEvent) GlobalPos() *QPoint {
	var __rv uintptr
	q.Drv(136000,136103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QTabletEvent::globalX()
func (q *QTabletEvent) GlobalX() int {
	var __rv int
	q.Drv(136000,136104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::globalY()
func (q *QTabletEvent) GlobalY() int {
	var __rv int
	q.Drv(136000,136105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::hiResGlobalX()
func (q *QTabletEvent) HiResGlobalX() float64 {
	var __rv float64
	q.Drv(136000,136106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::hiResGlobalY()
func (q *QTabletEvent) HiResGlobalY() float64 {
	var __rv float64
	q.Drv(136000,136107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::pointerType()
func (q *QTabletEvent) PointerType() QTabletEvent_PointerType {
	var __rv QTabletEvent_PointerType
	q.Drv(136000,136108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::pos()
func (q *QTabletEvent) Pos() *QPoint {
	var __rv uintptr
	q.Drv(136000,136109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QTabletEvent::pressure()
func (q *QTabletEvent) Pressure() float64 {
	var __rv float64
	q.Drv(136000,136110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::rotation()
func (q *QTabletEvent) Rotation() float64 {
	var __rv float64
	q.Drv(136000,136111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::tangentialPressure()
func (q *QTabletEvent) TangentialPressure() float64 {
	var __rv float64
	q.Drv(136000,136112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::uniqueId()
func (q *QTabletEvent) UniqueId() int64 {
	var __rv int64
	q.Drv(136000,136113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::x()
func (q *QTabletEvent) X() int {
	var __rv int
	q.Drv(136000,136114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::xTilt()
func (q *QTabletEvent) XTilt() int {
	var __rv int
	q.Drv(136000,136115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::y()
func (q *QTabletEvent) Y() int {
	var __rv int
	q.Drv(136000,136116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::yTilt()
func (q *QTabletEvent) YTilt() int {
	var __rv int
	q.Drv(136000,136117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTabletEvent::z()
func (q *QTabletEvent) Z() int {
	var __rv int
	q.Drv(136000,136118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
