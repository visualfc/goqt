// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QVariantAnimation : QVariantAnimation
type QVariantAnimation struct {
	QAbstractAnimation
}
func (q *QVariantAnimation) OnValueChanged(fn func(*QVariant)) uintptr {
	var __rv uintptr
	q.Drv(394000,394102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QVariantAnimation::currentValue()
func (q *QVariantAnimation) CurrentValue() *QVariant {
	var __rv uintptr
	q.Drv(394000,394103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QVariantAnimation::duration()
func (q *QVariantAnimation) Duration() int {
	var __rv int
	q.Drv(394000,394104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariantAnimation::easingCurve()
func (q *QVariantAnimation) EasingCurve() *QEasingCurve {
	var __rv uintptr
	q.Drv(394000,394105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QEasingCurve{}
	_rp.SetDriver(__rv,29000,true)
	return _rp
}	
//QVariantAnimation::endValue()
func (q *QVariantAnimation) EndValue() *QVariant {
	var __rv uintptr
	q.Drv(394000,394106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QVariantAnimation::event(QEvent*)
func (q *QVariantAnimation) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(394000,394107,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QVariantAnimation::interpolated(QVariant const&,QVariant const&,double)
func (q *QVariantAnimation) Interpolated(from *QVariant,to *QVariant,progress float64) *QVariant {
	var __rv uintptr
	q.Drv(394000,394108,Native(from),Native(to),unsafe.Pointer(&progress),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QVariantAnimation::keyValueAt(double)
func (q *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	var __rv uintptr
	q.Drv(394000,394109,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QVariantAnimation::setDuration(int)
func (q *QVariantAnimation) SetDuration(msecs int)  {
	q.Drv(394000,394110,unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::setEasingCurve(QEasingCurve const&)
func (q *QVariantAnimation) SetEasingCurve(easing *QEasingCurve)  {
	q.Drv(394000,394111,Native(easing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::setEndValue(QVariant const&)
func (q *QVariantAnimation) SetEndValue(value *QVariant)  {
	q.Drv(394000,394112,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::setKeyValueAt(double,QVariant const&)
func (q *QVariantAnimation) SetKeyValueAt(step float64,value *QVariant)  {
	q.Drv(394000,394113,unsafe.Pointer(&step),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::setStartValue(QVariant const&)
func (q *QVariantAnimation) SetStartValue(value *QVariant)  {
	q.Drv(394000,394114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::startValue()
func (q *QVariantAnimation) StartValue() *QVariant {
	var __rv uintptr
	q.Drv(394000,394115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QVariantAnimation::updateCurrentTime(int)
func (q *QVariantAnimation) UpdateCurrentTime(value int)  {
	q.Drv(394000,394116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::updateCurrentValue(QVariant const&)
func (q *QVariantAnimation) UpdateCurrentValue(value *QVariant)  {
	q.Drv(394000,394117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QVariantAnimation::updateState(QAbstractAnimation::State,QAbstractAnimation::State)
func (q *QVariantAnimation) UpdateState(newState QAbstractAnimation_State,oldState QAbstractAnimation_State)  {
	q.Drv(394000,394118,unsafe.Pointer(&newState),unsafe.Pointer(&oldState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
