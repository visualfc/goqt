// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsDropShadowEffect : QGraphicsDropShadowEffect
type QGraphicsDropShadowEffect struct {
	QGraphicsEffect
}
func (q *QGraphicsDropShadowEffect) OnBlurRadiusChanged(fn func(float64)) uintptr {
	var __rv uintptr
	q.Drv(252000,252102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsDropShadowEffect) OnOffsetChanged(fn func(*QPointF)) uintptr {
	var __rv uintptr
	q.Drv(252000,252103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsDropShadowEffect) OnColorChanged(fn func(*QColor)) uintptr {
	var __rv uintptr
	q.Drv(252000,252104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsDropShadowEffect::QGraphicsDropShadowEffect()	
func NewGraphicsDropShadowEffect() *QGraphicsDropShadowEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),252000,252105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsDropShadowEffect{}
	_p.SetDriver(__rv,252000,false)
	return _p
} 
//QGraphicsDropShadowEffect::QGraphicsDropShadowEffect(QObject*)	
func NewGraphicsDropShadowEffectWithParent(parent QObjectInterface) *QGraphicsDropShadowEffect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),252000,252106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsDropShadowEffect{}
	_p.SetDriver(__rv,252000,false)
	return _p
} 
//QGraphicsDropShadowEffect::blurRadius()
func (q *QGraphicsDropShadowEffect) BlurRadius() float64 {
	var __rv float64
	q.Drv(252000,252107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsDropShadowEffect::boundingRectFor(QRectF const&)
func (q *QGraphicsDropShadowEffect) BoundingRectFor(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(252000,252108,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsDropShadowEffect::color()
func (q *QGraphicsDropShadowEffect) Color() *QColor {
	var __rv uintptr
	q.Drv(252000,252109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QGraphicsDropShadowEffect::draw(QPainter*)
func (q *QGraphicsDropShadowEffect) Draw(painter *QPainter)  {
	q.Drv(252000,252110,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::offset()
func (q *QGraphicsDropShadowEffect) Offset() *QPointF {
	var __rv uintptr
	q.Drv(252000,252111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsDropShadowEffect::setBlurRadius(double)
func (q *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64)  {
	q.Drv(252000,252112,unsafe.Pointer(&blurRadius),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setColor(QColor const&)
func (q *QGraphicsDropShadowEffect) SetColor(color *QColor)  {
	q.Drv(252000,252113,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setOffset(QPointF const&)
func (q *QGraphicsDropShadowEffect) SetOffset(ofs *QPointF)  {
	q.Drv(252000,252114,Native(ofs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setOffset(double)
func (q *QGraphicsDropShadowEffect) SetOffsetFWithFloat64(d float64)  {
	q.Drv(252000,252115,unsafe.Pointer(&d),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setOffset(double,double)
func (q *QGraphicsDropShadowEffect) SetOffsetFWithDxDy(dx float64,dy float64)  {
	q.Drv(252000,252116,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setXOffset(double)
func (q *QGraphicsDropShadowEffect) SetXOffset(dx float64)  {
	q.Drv(252000,252117,unsafe.Pointer(&dx),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::setYOffset(double)
func (q *QGraphicsDropShadowEffect) SetYOffset(dy float64)  {
	q.Drv(252000,252118,unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsDropShadowEffect::xOffset()
func (q *QGraphicsDropShadowEffect) XOffset() float64 {
	var __rv float64
	q.Drv(252000,252119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsDropShadowEffect::yOffset()
func (q *QGraphicsDropShadowEffect) YOffset() float64 {
	var __rv float64
	q.Drv(252000,252120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
