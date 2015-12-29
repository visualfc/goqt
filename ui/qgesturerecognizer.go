// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGestureRecognizer_ResultFlag - QGestureRecognizer::ResultFlag
type QGestureRecognizer_ResultFlag uint32
const (
	QGestureRecognizer_Ignore QGestureRecognizer_ResultFlag = 0x0001
	QGestureRecognizer_MayBeGesture QGestureRecognizer_ResultFlag = 0x0002
	QGestureRecognizer_TriggerGesture QGestureRecognizer_ResultFlag = 0x0004
	QGestureRecognizer_FinishGesture QGestureRecognizer_ResultFlag = 0x0008
	QGestureRecognizer_CancelGesture QGestureRecognizer_ResultFlag = 0x0010
	QGestureRecognizer_ResultState_Mask QGestureRecognizer_ResultFlag = 0x00ff
	QGestureRecognizer_ConsumeEventHint QGestureRecognizer_ResultFlag = 0x0100
	QGestureRecognizer_ResultHint_Mask QGestureRecognizer_ResultFlag = 0xff00
)
//struct QGestureRecognizer : QGestureRecognizer
type QGestureRecognizer struct {
	BaseDrv
}
//QGestureRecognizer::create(QObject*)
func (q *QGestureRecognizer) Create(target QObjectInterface) *QGesture {
	var __rv uintptr
	q.Drv(45000,45102,Native(target),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGesture{}
	_rp.SetDriver(__rv,247000,false)
	return _rp
}	
//QGestureRecognizer::recognize(QGesture*,QObject*,QEvent*)
func (q *QGestureRecognizer) Recognize(state *QGesture,watched QObjectInterface,event *QEvent) QGestureRecognizer_ResultFlag {
	var __rv QGestureRecognizer_ResultFlag
	q.Drv(45000,45103,Native(state),Native(watched),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureRecognizer::registerRecognizer(QGestureRecognizer*)	
func QGestureRecognizerRegisterRecognizer(recognizer *QGestureRecognizer) Qt_GestureType {
	var __rv Qt_GestureType
	DirectQtDrv(nil,45000,45104,Native(recognizer),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureRecognizer::registerRecognizer(QGestureRecognizer*)
func (q *QGestureRecognizer) RegisterRecognizer(recognizer *QGestureRecognizer) Qt_GestureType {
	var __rv Qt_GestureType
	q.Drv(45000,45104,Native(recognizer),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGestureRecognizer::reset(QGesture*)
func (q *QGestureRecognizer) Reset(state *QGesture)  {
	q.Drv(45000,45105,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureRecognizer::unregisterRecognizer(Qt::GestureType)	
func QGestureRecognizerUnregisterRecognizer(_type Qt_GestureType)  {
	DirectQtDrv(nil,45000,45106,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGestureRecognizer::unregisterRecognizer(Qt::GestureType)
func (q *QGestureRecognizer) UnregisterRecognizer(_type Qt_GestureType)  {
	q.Drv(45000,45106,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
