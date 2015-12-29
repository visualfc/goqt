// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractAnimation_Direction - QAbstractAnimation::Direction
type QAbstractAnimation_Direction uint32
const (
	QAbstractAnimation_Forward QAbstractAnimation_Direction = 0
	QAbstractAnimation_Backward QAbstractAnimation_Direction = 1
)
//enum QAbstractAnimation_DeletionPolicy - QAbstractAnimation::DeletionPolicy
type QAbstractAnimation_DeletionPolicy uint32
const (
	QAbstractAnimation_KeepWhenStopped QAbstractAnimation_DeletionPolicy = 0
	QAbstractAnimation_DeleteWhenStopped QAbstractAnimation_DeletionPolicy = 0
)
//enum QAbstractAnimation_State - QAbstractAnimation::State
type QAbstractAnimation_State uint32
const (
	QAbstractAnimation_Stopped QAbstractAnimation_State = 0
	QAbstractAnimation_Paused QAbstractAnimation_State = 1
	QAbstractAnimation_Running QAbstractAnimation_State = 2
)
//struct QAbstractAnimation : QAbstractAnimation
type QAbstractAnimation struct {
	QObject
}
func (q *QAbstractAnimation) OnFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(192000,192102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractAnimation) OnCurrentLoopChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(192000,192103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractAnimation) OnStateChanged(fn func(QAbstractAnimation_State,QAbstractAnimation_State)) uintptr {
	var __rv uintptr
	q.Drv(192000,192104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractAnimation) OnDirectionChanged(fn func(QAbstractAnimation_Direction)) uintptr {
	var __rv uintptr
	q.Drv(192000,192105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractAnimation::currentLoop()
func (q *QAbstractAnimation) CurrentLoop() int {
	var __rv int
	q.Drv(192000,192106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::currentLoopTime()
func (q *QAbstractAnimation) CurrentLoopTime() int {
	var __rv int
	q.Drv(192000,192107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::currentTime()
func (q *QAbstractAnimation) CurrentTime() int {
	var __rv int
	q.Drv(192000,192108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::direction()
func (q *QAbstractAnimation) Direction() QAbstractAnimation_Direction {
	var __rv QAbstractAnimation_Direction
	q.Drv(192000,192109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::duration()
func (q *QAbstractAnimation) Duration() int {
	var __rv int
	q.Drv(192000,192110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::event(QEvent*)
func (q *QAbstractAnimation) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(192000,192111,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::group()
func (q *QAbstractAnimation) Group() *QAnimationGroup {
	var __rv uintptr
	q.Drv(192000,192112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAnimationGroup{}
	_rp.SetDriver(__rv,209000,false)
	return _rp
}	
//QAbstractAnimation::loopCount()
func (q *QAbstractAnimation) LoopCount() int {
	var __rv int
	q.Drv(192000,192113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::pause()
func (q *QAbstractAnimation) Pause()  {
	q.Drv(192000,192114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::resume()
func (q *QAbstractAnimation) Resume()  {
	q.Drv(192000,192115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::setCurrentTime(int)
func (q *QAbstractAnimation) SetCurrentTime(msecs int)  {
	q.Drv(192000,192116,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::setDirection(QAbstractAnimation::Direction)
func (q *QAbstractAnimation) SetDirection(direction QAbstractAnimation_Direction)  {
	q.Drv(192000,192117,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::setLoopCount(int)
func (q *QAbstractAnimation) SetLoopCount(loopCount int)  {
	q.Drv(192000,192118,unsafe.Pointer(&loopCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::setPaused(bool)
func (q *QAbstractAnimation) SetPaused(value bool)  {
	q.Drv(192000,192119,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::start()
func (q *QAbstractAnimation) Start()  {
	q.Drv(192000,192120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::start(QAbstractAnimation::DeletionPolicy)
func (q *QAbstractAnimation) StartWithPolicy(policy QAbstractAnimation_DeletionPolicy)  {
	q.Drv(192000,192121,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::state()
func (q *QAbstractAnimation) State() QAbstractAnimation_State {
	var __rv QAbstractAnimation_State
	q.Drv(192000,192122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::stop()
func (q *QAbstractAnimation) Stop()  {
	q.Drv(192000,192123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::totalDuration()
func (q *QAbstractAnimation) TotalDuration() int {
	var __rv int
	q.Drv(192000,192124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractAnimation::updateCurrentTime(int)
func (q *QAbstractAnimation) UpdateCurrentTime(currentTime int)  {
	q.Drv(192000,192125,unsafe.Pointer(&currentTime),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::updateDirection(QAbstractAnimation::Direction)
func (q *QAbstractAnimation) UpdateDirection(direction QAbstractAnimation_Direction)  {
	q.Drv(192000,192126,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractAnimation::updateState(QAbstractAnimation::State,QAbstractAnimation::State)
func (q *QAbstractAnimation) UpdateState(newState QAbstractAnimation_State,oldState QAbstractAnimation_State)  {
	q.Drv(192000,192127,unsafe.Pointer(&newState),unsafe.Pointer(&oldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
