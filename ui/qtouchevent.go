// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTouchEvent_DeviceType - QTouchEvent::DeviceType
type QTouchEvent_DeviceType uint32
const (
	QTouchEvent_TouchScreen QTouchEvent_DeviceType = 0
	QTouchEvent_TouchPad QTouchEvent_DeviceType = 1
)
//struct QTouchEvent : QTouchEvent
type QTouchEvent struct {
	QInputEvent
}
//QTouchEvent::setTouchPointStates(QFlags<Qt::TouchPointState>)
func (q *QTouchEvent) SetTouchPointStates(aTouchPointStates Qt_TouchPointState)  {
	q.Drv(174000,174102,unsafe.Pointer(&aTouchPointStates),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::setTouchPoints(QList<QTouchEvent::TouchPoint> const&)
func (q *QTouchEvent) SetTouchPoints(atouchPoints []*QTouchEventTouchPoint)  {
	q.Drv(174000,174103,unsafe.Pointer(&atouchPoints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTouchEvent::touchPointStates()
func (q *QTouchEvent) TouchPointStates() Qt_TouchPointState {
	var __rv Qt_TouchPointState
	q.Drv(174000,174104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTouchEvent::touchPoints()
func (q *QTouchEvent) TouchPoints() []*QTouchEventTouchPoint {
	var __rv []*QTouchEventTouchPoint
	q.Drv(174000,174105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
