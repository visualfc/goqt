// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QElapsedTimer_ClockType - QElapsedTimer::ClockType
type QElapsedTimer_ClockType uint32
const (
	QElapsedTimer_SystemTime QElapsedTimer_ClockType = 0
	QElapsedTimer_MonotonicClock QElapsedTimer_ClockType = 1
	QElapsedTimer_TickCounter QElapsedTimer_ClockType = 2
	QElapsedTimer_MachAbsoluteTime QElapsedTimer_ClockType = 3
)
//struct QElapsedTimer : QElapsedTimer
type QElapsedTimer struct {
	BaseDrv
}
//QElapsedTimer::QElapsedTimer()	
func NewElapsedTimer() *QElapsedTimer {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),30000,30102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QElapsedTimer{}
	_p.SetDriver(__rv,30000,true)
	return _p
} 
//QElapsedTimer::clockType()	
func QElapsedTimerClockType() QElapsedTimer_ClockType {
	var __rv QElapsedTimer_ClockType
	DirectQtDrv(nil,30000,30103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::clockType()
func (q *QElapsedTimer) ClockType() QElapsedTimer_ClockType {
	var __rv QElapsedTimer_ClockType
	q.Drv(30000,30103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::elapsed()
func (q *QElapsedTimer) Elapsed() int64 {
	var __rv int64
	q.Drv(30000,30104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::hasExpired(qint64)
func (q *QElapsedTimer) HasExpired(timeout int64) bool {
	var __rv bool
	q.Drv(30000,30105,unsafe.Pointer(&timeout),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::invalidate()
func (q *QElapsedTimer) Invalidate()  {
	q.Drv(30000,30106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QElapsedTimer::isMonotonic()	
func QElapsedTimerIsMonotonic() bool {
	var __rv bool
	DirectQtDrv(nil,30000,30107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::isMonotonic()
func (q *QElapsedTimer) IsMonotonic() bool {
	var __rv bool
	q.Drv(30000,30107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::isValid()
func (q *QElapsedTimer) IsValid() bool {
	var __rv bool
	q.Drv(30000,30108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::msecsSinceReference()
func (q *QElapsedTimer) MsecsSinceReference() int64 {
	var __rv int64
	q.Drv(30000,30109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::msecsTo(QElapsedTimer const&)
func (q *QElapsedTimer) MsecsTo(other *QElapsedTimer) int64 {
	var __rv int64
	q.Drv(30000,30110,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::restart()
func (q *QElapsedTimer) Restart() int64 {
	var __rv int64
	q.Drv(30000,30111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::secsTo(QElapsedTimer const&)
func (q *QElapsedTimer) SecsTo(other *QElapsedTimer) int64 {
	var __rv int64
	q.Drv(30000,30112,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QElapsedTimer::start()
func (q *QElapsedTimer) Start()  {
	q.Drv(30000,30113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
