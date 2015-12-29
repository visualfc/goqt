// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFutureWatcherBase : QFutureWatcherBase
type QFutureWatcherBase struct {
	QObject
}
func (q *QFutureWatcherBase) OnFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(246000,246102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnProgressTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(246000,246103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnCanceled(fn func()) uintptr {
	var __rv uintptr
	q.Drv(246000,246104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnProgressValueChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(246000,246105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnStarted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(246000,246106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnProgressRangeChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(246000,246107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnResultsReadyAt(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(246000,246108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnPaused(fn func()) uintptr {
	var __rv uintptr
	q.Drv(246000,246109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnResumed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(246000,246110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFutureWatcherBase) OnResultReadyAt(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(246000,246111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFutureWatcherBase::cancel()
func (q *QFutureWatcherBase) Cancel()  {
	q.Drv(246000,246112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::connectOutputInterface()
func (q *QFutureWatcherBase) ConnectOutputInterface()  {
	q.Drv(246000,246113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::disconnectOutputInterface(bool)
func (q *QFutureWatcherBase) DisconnectOutputInterface(pendingAssignment bool)  {
	q.Drv(246000,246114,unsafe.Pointer(&pendingAssignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::event(QEvent*)
func (q *QFutureWatcherBase) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(246000,246115,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::isCanceled()
func (q *QFutureWatcherBase) IsCanceled() bool {
	var __rv bool
	q.Drv(246000,246116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::isFinished()
func (q *QFutureWatcherBase) IsFinished() bool {
	var __rv bool
	q.Drv(246000,246117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::isPaused()
func (q *QFutureWatcherBase) IsPaused() bool {
	var __rv bool
	q.Drv(246000,246118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::isRunning()
func (q *QFutureWatcherBase) IsRunning() bool {
	var __rv bool
	q.Drv(246000,246119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::isStarted()
func (q *QFutureWatcherBase) IsStarted() bool {
	var __rv bool
	q.Drv(246000,246120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::pause()
func (q *QFutureWatcherBase) Pause()  {
	q.Drv(246000,246121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::progressMaximum()
func (q *QFutureWatcherBase) ProgressMaximum() int {
	var __rv int
	q.Drv(246000,246122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::progressMinimum()
func (q *QFutureWatcherBase) ProgressMinimum() int {
	var __rv int
	q.Drv(246000,246123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::progressText()
func (q *QFutureWatcherBase) ProgressText() string {
	var __rv string
	q.Drv(246000,246124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::progressValue()
func (q *QFutureWatcherBase) ProgressValue() int {
	var __rv int
	q.Drv(246000,246125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFutureWatcherBase::resume()
func (q *QFutureWatcherBase) Resume()  {
	q.Drv(246000,246126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::setPaused(bool)
func (q *QFutureWatcherBase) SetPaused(paused bool)  {
	q.Drv(246000,246127,unsafe.Pointer(&paused),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::setPendingResultsLimit(int)
func (q *QFutureWatcherBase) SetPendingResultsLimit(limit int)  {
	q.Drv(246000,246128,unsafe.Pointer(&limit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::togglePaused()
func (q *QFutureWatcherBase) TogglePaused()  {
	q.Drv(246000,246129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFutureWatcherBase::waitForFinished()
func (q *QFutureWatcherBase) WaitForFinished()  {
	q.Drv(246000,246130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
