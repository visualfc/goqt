// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsColorizeEffect : QGraphicsColorizeEffect
type QGraphicsColorizeEffect struct {
	QGraphicsEffect
}
func (q *QGraphicsColorizeEffect) OnColorChanged(fn func(*QColor)) uintptr {
	var __rv uintptr
	q.Drv(251000,251102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsColorizeEffect) OnStrengthChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(251000,251103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsColorizeEffect::QGraphicsColorizeEffect()	
func NewGraphicsColorizeEffect() *QGraphicsColorizeEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),251000,251104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsColorizeEffect{}
	_p.SetDriver(__rv,251000,false)
	return _p
} 
//QGraphicsColorizeEffect::QGraphicsColorizeEffect(QObject*)	
func NewGraphicsColorizeEffectWithParent(parent QObjectInterface) *QGraphicsColorizeEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),251000,251105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsColorizeEffect{}
	_p.SetDriver(__rv,251000,false)
	return _p
} 
//QGraphicsColorizeEffect::color()
func (q *QGraphicsColorizeEffect) Color() *QColor {
	var __rv uintptr
	q.Drv(251000,251106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QGraphicsColorizeEffect::draw(QPainter*)
func (q *QGraphicsColorizeEffect) Draw(painter *QPainter)  {
	q.Drv(251000,251107,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsColorizeEffect::setColor(QColor const&)
func (q *QGraphicsColorizeEffect) SetColor(c *QColor)  {
	q.Drv(251000,251108,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsColorizeEffect::setStrength(double)
func (q *QGraphicsColorizeEffect) SetStrength(strength float64)  {
	q.Drv(251000,251109,unsafe.Pointer(&strength),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsColorizeEffect::strength()
func (q *QGraphicsColorizeEffect) Strength() float64 {
	var __rv float64
	q.Drv(251000,251110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
