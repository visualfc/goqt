// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsOpacityEffect : QGraphicsOpacityEffect
type QGraphicsOpacityEffect struct {
	QGraphicsEffect
}
func (q *QGraphicsOpacityEffect) OnOpacityChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(264000,264102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsOpacityEffect) OnOpacityMaskChanged(fn func(*QBrush)) uintptr {
	var __rv uintptr
	q.Drv(264000,264103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsOpacityEffect::QGraphicsOpacityEffect()	
func NewGraphicsOpacityEffect() *QGraphicsOpacityEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),264000,264104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsOpacityEffect{}
	_p.SetDriver(__rv,264000,false)
	return _p
} 
//QGraphicsOpacityEffect::QGraphicsOpacityEffect(QObject*)	
func NewGraphicsOpacityEffectWithParent(parent QObjectInterface) *QGraphicsOpacityEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),264000,264105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsOpacityEffect{}
	_p.SetDriver(__rv,264000,false)
	return _p
} 
//QGraphicsOpacityEffect::draw(QPainter*)
func (q *QGraphicsOpacityEffect) Draw(painter *QPainter)  {
	q.Drv(264000,264106,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsOpacityEffect::opacity()
func (q *QGraphicsOpacityEffect) Opacity() float64 {
	var __rv float64
	q.Drv(264000,264107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsOpacityEffect::opacityMask()
func (q *QGraphicsOpacityEffect) OpacityMask() *QBrush {
	var __rv uintptr
	q.Drv(264000,264108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QGraphicsOpacityEffect::setOpacity(double)
func (q *QGraphicsOpacityEffect) SetOpacity(opacity float64)  {
	q.Drv(264000,264109,unsafe.Pointer(&opacity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsOpacityEffect::setOpacityMask(QBrush const&)
func (q *QGraphicsOpacityEffect) SetOpacityMask(mask *QBrush)  {
	q.Drv(264000,264110,Native(mask),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
