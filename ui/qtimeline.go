// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTimeLine_Direction - QTimeLine::Direction
type QTimeLine_Direction uint32
const (
	QTimeLine_Forward QTimeLine_Direction = 0
	QTimeLine_Backward QTimeLine_Direction = 1
)
//enum QTimeLine_CurveShape - QTimeLine::CurveShape
type QTimeLine_CurveShape uint32
const (
	QTimeLine_EaseInCurve QTimeLine_CurveShape = 0
	QTimeLine_EaseOutCurve QTimeLine_CurveShape = 1
	QTimeLine_EaseInOutCurve QTimeLine_CurveShape = 2
	QTimeLine_LinearCurve QTimeLine_CurveShape = 3
	QTimeLine_SineCurve QTimeLine_CurveShape = 4
	QTimeLine_CosineCurve QTimeLine_CurveShape = 5
)
//enum QTimeLine_State - QTimeLine::State
type QTimeLine_State uint32
const (
	QTimeLine_NotRunning QTimeLine_State = 0
	QTimeLine_Paused QTimeLine_State = 1
	QTimeLine_Running QTimeLine_State = 2
)
//struct QTimeLine : QTimeLine
type QTimeLine struct {
	QObject
}
func (q *QTimeLine) OnFrameChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(379000,379102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTimeLine) OnStateChanged(fn func(QTimeLine_State)) uintptr {
	var __rv uintptr
	q.Drv(379000,379103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTimeLine) OnFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(379000,379104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTimeLine) OnValueChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(379000,379105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTimeLine::QTimeLine()	
func NewTimeLine() *QTimeLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),379000,379106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimeLine{}
	_p.SetDriver(__rv,379000,false)
	return _p
} 
//QTimeLine::QTimeLine(int,QObject*)	
func NewTimeLineWithDurationParent(duration int,parent QObjectInterface) *QTimeLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),379000,379107,unsafe.Pointer(&duration),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTimeLine{}
	_p.SetDriver(__rv,379000,false)
	return _p
} 
//QTimeLine::currentFrame()
func (q *QTimeLine) CurrentFrame() int {
	var __rv int
	q.Drv(379000,379108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::currentTime()
func (q *QTimeLine) CurrentTime() int {
	var __rv int
	q.Drv(379000,379109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::currentValue()
func (q *QTimeLine) CurrentValue() float64 {
	var __rv float64
	q.Drv(379000,379110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::curveShape()
func (q *QTimeLine) CurveShape() QTimeLine_CurveShape {
	var __rv QTimeLine_CurveShape
	q.Drv(379000,379111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::direction()
func (q *QTimeLine) Direction() QTimeLine_Direction {
	var __rv QTimeLine_Direction
	q.Drv(379000,379112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::duration()
func (q *QTimeLine) Duration() int {
	var __rv int
	q.Drv(379000,379113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::easingCurve()
func (q *QTimeLine) EasingCurve() *QEasingCurve {
	var __rv uintptr
	q.Drv(379000,379114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QEasingCurve{}
	_rp.SetDriver(__rv,29000,true)
	return _rp
}	
//QTimeLine::endFrame()
func (q *QTimeLine) EndFrame() int {
	var __rv int
	q.Drv(379000,379115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::frameForTime(int)
func (q *QTimeLine) FrameForTime(msec int) int {
	var __rv int
	q.Drv(379000,379116,unsafe.Pointer(&msec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::loopCount()
func (q *QTimeLine) LoopCount() int {
	var __rv int
	q.Drv(379000,379117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::resume()
func (q *QTimeLine) Resume()  {
	q.Drv(379000,379118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setCurrentTime(int)
func (q *QTimeLine) SetCurrentTime(msec int)  {
	q.Drv(379000,379119,unsafe.Pointer(&msec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setCurveShape(QTimeLine::CurveShape)
func (q *QTimeLine) SetCurveShape(shape QTimeLine_CurveShape)  {
	q.Drv(379000,379120,unsafe.Pointer(&shape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setDirection(QTimeLine::Direction)
func (q *QTimeLine) SetDirection(direction QTimeLine_Direction)  {
	q.Drv(379000,379121,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setDuration(int)
func (q *QTimeLine) SetDuration(duration int)  {
	q.Drv(379000,379122,unsafe.Pointer(&duration),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setEasingCurve(QEasingCurve const&)
func (q *QTimeLine) SetEasingCurve(curve *QEasingCurve)  {
	q.Drv(379000,379123,Native(curve),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setEndFrame(int)
func (q *QTimeLine) SetEndFrame(frame int)  {
	q.Drv(379000,379124,unsafe.Pointer(&frame),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setFrameRange(int,int)
func (q *QTimeLine) SetFrameRange(startFrame int,endFrame int)  {
	q.Drv(379000,379125,unsafe.Pointer(&startFrame),unsafe.Pointer(&endFrame),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setLoopCount(int)
func (q *QTimeLine) SetLoopCount(count int)  {
	q.Drv(379000,379126,unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setPaused(bool)
func (q *QTimeLine) SetPaused(paused bool)  {
	q.Drv(379000,379127,unsafe.Pointer(&paused),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setStartFrame(int)
func (q *QTimeLine) SetStartFrame(frame int)  {
	q.Drv(379000,379128,unsafe.Pointer(&frame),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::setUpdateInterval(int)
func (q *QTimeLine) SetUpdateInterval(interval int)  {
	q.Drv(379000,379129,unsafe.Pointer(&interval),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::start()
func (q *QTimeLine) Start()  {
	q.Drv(379000,379130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::startFrame()
func (q *QTimeLine) StartFrame() int {
	var __rv int
	q.Drv(379000,379131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::state()
func (q *QTimeLine) State() QTimeLine_State {
	var __rv QTimeLine_State
	q.Drv(379000,379132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::stop()
func (q *QTimeLine) Stop()  {
	q.Drv(379000,379133,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::timerEvent(QTimerEvent*)
func (q *QTimeLine) TimerEvent(event *QTimerEvent)  {
	q.Drv(379000,379134,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::toggleDirection()
func (q *QTimeLine) ToggleDirection()  {
	q.Drv(379000,379135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTimeLine::updateInterval()
func (q *QTimeLine) UpdateInterval() int {
	var __rv int
	q.Drv(379000,379136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTimeLine::valueForTime(int)
func (q *QTimeLine) ValueForTime(msec int) float64 {
	var __rv float64
	q.Drv(379000,379137,unsafe.Pointer(&msec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
