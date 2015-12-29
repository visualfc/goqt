// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsObject : QGraphicsObject
type QGraphicsObject struct {
	QObject
}
func (q *QGraphicsObject) OnRotationChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnYChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnScaleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnEnabledChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnOpacityChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnVisibleChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnWidthChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnHeightChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnParentChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnZChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnChildrenChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263112,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsObject) OnXChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(263000,263113,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsObject::children()
func (q *QGraphicsObject) Children() []*QObject {
	var __rv []*QObject
	q.Drv(263000,263114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsObject::grabGesture(Qt::GestureType)
func (q *QGraphicsObject) GrabGesture(_type Qt_GestureType)  {
	q.Drv(263000,263115,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsObject::grabGesture(Qt::GestureType,QFlags<Qt::GestureFlag>)
func (q *QGraphicsObject) GrabGestureWithTypeFlags(_type Qt_GestureType,flags Qt_GestureFlag)  {
	q.Drv(263000,263116,unsafe.Pointer(&_type),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsObject::ungrabGesture(Qt::GestureType)
func (q *QGraphicsObject) UngrabGesture(_type Qt_GestureType)  {
	q.Drv(263000,263117,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsObject::updateMicroFocus()
func (q *QGraphicsObject) UpdateMicroFocus()  {
	q.Drv(263000,263118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
