// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractSlider_SliderAction - QAbstractSlider::SliderAction
type QAbstractSlider_SliderAction uint32
const (
	QAbstractSlider_SliderNoAction QAbstractSlider_SliderAction = 0
	QAbstractSlider_SliderSingleStepAdd QAbstractSlider_SliderAction = 1
	QAbstractSlider_SliderSingleStepSub QAbstractSlider_SliderAction = 2
	QAbstractSlider_SliderPageStepAdd QAbstractSlider_SliderAction = 3
	QAbstractSlider_SliderPageStepSub QAbstractSlider_SliderAction = 4
	QAbstractSlider_SliderToMinimum QAbstractSlider_SliderAction = 5
	QAbstractSlider_SliderToMaximum QAbstractSlider_SliderAction = 6
	QAbstractSlider_SliderMove QAbstractSlider_SliderAction = 7
)
//enum QAbstractSlider_SliderChange - QAbstractSlider::SliderChange
type QAbstractSlider_SliderChange uint32
const (
	QAbstractSlider_SliderRangeChange QAbstractSlider_SliderChange = 0
	QAbstractSlider_SliderOrientationChange QAbstractSlider_SliderChange = 1
	QAbstractSlider_SliderStepsChange QAbstractSlider_SliderChange = 2
	QAbstractSlider_SliderValueChange QAbstractSlider_SliderChange = 3
)
//struct QAbstractSlider : QAbstractSlider
type QAbstractSlider struct {
	QWidget
}
func (q *QAbstractSlider) OnRangeChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(201000,201102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractSlider) OnSliderMoved(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(201000,201103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractSlider) OnActionTriggered(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(201000,201104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractSlider) OnSliderReleased(fn func()) uintptr {
	var __rv uintptr
	q.Drv(201000,201105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractSlider) OnSliderPressed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(201000,201106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractSlider) OnValueChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(201000,201107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractSlider::QAbstractSlider()	
func NewAbstractSlider() *QAbstractSlider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),201000,201108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractSlider{}
	_p.SetDriver(__rv,201000,false)
	return _p
} 
//QAbstractSlider::QAbstractSlider(QWidget*)	
func NewAbstractSliderWithParent(parent QWidgetInterface) *QAbstractSlider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),201000,201109,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractSlider{}
	_p.SetDriver(__rv,201000,false)
	return _p
} 
//QAbstractSlider::changeEvent(QEvent*)
func (q *QAbstractSlider) ChangeEvent(e *QEvent)  {
	q.Drv(201000,201110,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::event(QEvent*)
func (q *QAbstractSlider) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(201000,201111,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::hasTracking()
func (q *QAbstractSlider) HasTracking() bool {
	var __rv bool
	q.Drv(201000,201112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::invertedAppearance()
func (q *QAbstractSlider) InvertedAppearance() bool {
	var __rv bool
	q.Drv(201000,201113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::invertedControls()
func (q *QAbstractSlider) InvertedControls() bool {
	var __rv bool
	q.Drv(201000,201114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::isSliderDown()
func (q *QAbstractSlider) IsSliderDown() bool {
	var __rv bool
	q.Drv(201000,201115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::keyPressEvent(QKeyEvent*)
func (q *QAbstractSlider) KeyPressEvent(ev *QKeyEvent)  {
	q.Drv(201000,201116,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::maximum()
func (q *QAbstractSlider) Maximum() int {
	var __rv int
	q.Drv(201000,201117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::minimum()
func (q *QAbstractSlider) Minimum() int {
	var __rv int
	q.Drv(201000,201118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::orientation()
func (q *QAbstractSlider) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(201000,201119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::pageStep()
func (q *QAbstractSlider) PageStep() int {
	var __rv int
	q.Drv(201000,201120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::repeatAction()
func (q *QAbstractSlider) RepeatAction() QAbstractSlider_SliderAction {
	var __rv QAbstractSlider_SliderAction
	q.Drv(201000,201121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::setInvertedAppearance(bool)
func (q *QAbstractSlider) SetInvertedAppearance(value bool)  {
	q.Drv(201000,201122,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setInvertedControls(bool)
func (q *QAbstractSlider) SetInvertedControls(value bool)  {
	q.Drv(201000,201123,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setMaximum(int)
func (q *QAbstractSlider) SetMaximum(value int)  {
	q.Drv(201000,201124,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setMinimum(int)
func (q *QAbstractSlider) SetMinimum(value int)  {
	q.Drv(201000,201125,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setOrientation(Qt::Orientation)
func (q *QAbstractSlider) SetOrientation(value Qt_Orientation)  {
	q.Drv(201000,201126,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setPageStep(int)
func (q *QAbstractSlider) SetPageStep(value int)  {
	q.Drv(201000,201127,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setRange(int,int)
func (q *QAbstractSlider) SetRange(min int,max int)  {
	q.Drv(201000,201128,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setRepeatAction(QAbstractSlider::SliderAction,int,int)
func (q *QAbstractSlider) SetRepeatAction(action QAbstractSlider_SliderAction,thresholdTime int,repeatTime int)  {
	q.Drv(201000,201129,unsafe.Pointer(&action),unsafe.Pointer(&thresholdTime),unsafe.Pointer(&repeatTime),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setSingleStep(int)
func (q *QAbstractSlider) SetSingleStep(value int)  {
	q.Drv(201000,201130,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setSliderDown(bool)
func (q *QAbstractSlider) SetSliderDown(value bool)  {
	q.Drv(201000,201131,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setSliderPosition(int)
func (q *QAbstractSlider) SetSliderPosition(value int)  {
	q.Drv(201000,201132,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setTracking(bool)
func (q *QAbstractSlider) SetTracking(enable bool)  {
	q.Drv(201000,201133,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::setValue(int)
func (q *QAbstractSlider) SetValue(value int)  {
	q.Drv(201000,201134,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::singleStep()
func (q *QAbstractSlider) SingleStep() int {
	var __rv int
	q.Drv(201000,201135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::sliderChange(QAbstractSlider::SliderChange)
func (q *QAbstractSlider) SliderChange(change QAbstractSlider_SliderChange)  {
	q.Drv(201000,201136,unsafe.Pointer(&change),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::sliderPosition()
func (q *QAbstractSlider) SliderPosition() int {
	var __rv int
	q.Drv(201000,201137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::timerEvent(QTimerEvent*)
func (q *QAbstractSlider) TimerEvent(value *QTimerEvent)  {
	q.Drv(201000,201138,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::triggerAction(QAbstractSlider::SliderAction)
func (q *QAbstractSlider) TriggerAction(action QAbstractSlider_SliderAction)  {
	q.Drv(201000,201139,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSlider::value()
func (q *QAbstractSlider) Value() int {
	var __rv int
	q.Drv(201000,201140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSlider::wheelEvent(QWheelEvent*)
func (q *QAbstractSlider) WheelEvent(e *QWheelEvent)  {
	q.Drv(201000,201141,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
