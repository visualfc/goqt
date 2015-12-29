// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTimer : QTimer
type QTimer struct {
	QObject
}
func (q *QTimer) OnTimeout(fn func()) uintptr {
	var __rv uintptr
	q.Drv(380000,380102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTimer::QTimer()	
func NewTimer() *QTimer {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),380000,380103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimer{}
	_p.SetDriver(__rv,380000,false)
	return _p
} 
//QTimer::QTimer(QObject*)	
func NewTimerWithParent(parent QObjectInterface) *QTimer {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),380000,380104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimer{}
	_p.SetDriver(__rv,380000,false)
	return _p
} 
//QTimer::interval()
func (q *QTimer) Interval() int {
	var __rv int
	q.Drv(380000,380105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimer::isActive()
func (q *QTimer) IsActive() bool {
	var __rv bool
	q.Drv(380000,380106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimer::isSingleShot()
func (q *QTimer) IsSingleShot() bool {
	var __rv bool
	q.Drv(380000,380107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimer::setInterval(int)
func (q *QTimer) SetInterval(msec int)  {
	q.Drv(380000,380108,unsafe.Pointer(&msec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::setSingleShot(bool)
func (q *QTimer) SetSingleShot(singleShot bool)  {
	q.Drv(380000,380109,unsafe.Pointer(&singleShot),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::singleShot(int,QObject*,char const*)	
func QTimerSingleShot(msec int,receiver QObjectInterface,member string)  {
	DirectQtDrv(nil,380000,380110,unsafe.Pointer(&msec),Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::singleShot(int,QObject*,char const*)
func (q *QTimer) SingleShot(msec int,receiver QObjectInterface,member string)  {
	q.Drv(380000,380110,unsafe.Pointer(&msec),Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::start()
func (q *QTimer) Start()  {
	q.Drv(380000,380111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::start(int)
func (q *QTimer) StartWithMsec(msec int)  {
	q.Drv(380000,380112,unsafe.Pointer(&msec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::stop()
func (q *QTimer) Stop()  {
	q.Drv(380000,380113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::timerEvent(QTimerEvent*)
func (q *QTimer) TimerEvent(value *QTimerEvent)  {
	q.Drv(380000,380114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimer::timerId()
func (q *QTimer) TimerId() int {
	var __rv int
	q.Drv(380000,380115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
