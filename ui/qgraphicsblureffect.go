// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsBlurEffect_BlurHint - QGraphicsBlurEffect::BlurHint
type QGraphicsBlurEffect_BlurHint uint32
const (
	QGraphicsBlurEffect_PerformanceHint QGraphicsBlurEffect_BlurHint = 0x00
	QGraphicsBlurEffect_QualityHint QGraphicsBlurEffect_BlurHint = 0x01
	QGraphicsBlurEffect_AnimationHint QGraphicsBlurEffect_BlurHint = 0x02
)
//struct QGraphicsBlurEffect : QGraphicsBlurEffect
type QGraphicsBlurEffect struct {
	QGraphicsEffect
}
func (q *QGraphicsBlurEffect) OnBlurRadiusChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(250000,250102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsBlurEffect) OnBlurHintsChanged(fn func(QGraphicsBlurEffect_BlurHint)) uintptr {
	var __rv uintptr
	q.Drv(250000,250103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsBlurEffect::QGraphicsBlurEffect()	
func NewGraphicsBlurEffect() *QGraphicsBlurEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),250000,250104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsBlurEffect{}
	_p.SetDriver(__rv,250000,false)
	return _p
} 
//QGraphicsBlurEffect::QGraphicsBlurEffect(QObject*)	
func NewGraphicsBlurEffectWithParent(parent QObjectInterface) *QGraphicsBlurEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),250000,250105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsBlurEffect{}
	_p.SetDriver(__rv,250000,false)
	return _p
} 
//QGraphicsBlurEffect::blurHints()
func (q *QGraphicsBlurEffect) BlurHints() QGraphicsBlurEffect_BlurHint {
	var __rv QGraphicsBlurEffect_BlurHint
	q.Drv(250000,250106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsBlurEffect::blurRadius()
func (q *QGraphicsBlurEffect) BlurRadius() float64 {
	var __rv float64
	q.Drv(250000,250107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsBlurEffect::boundingRectFor(QRectF const&)
func (q *QGraphicsBlurEffect) BoundingRectFor(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(250000,250108,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsBlurEffect::draw(QPainter*)
func (q *QGraphicsBlurEffect) Draw(painter *QPainter)  {
	q.Drv(250000,250109,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsBlurEffect::setBlurHints(QFlags<QGraphicsBlurEffect::BlurHint>)
func (q *QGraphicsBlurEffect) SetBlurHints(hints QGraphicsBlurEffect_BlurHint)  {
	q.Drv(250000,250110,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsBlurEffect::setBlurRadius(double)
func (q *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64)  {
	q.Drv(250000,250111,unsafe.Pointer(&blurRadius),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
