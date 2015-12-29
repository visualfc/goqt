// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPen : QPen
type QPen struct {
	BaseDrv
}
//QPen::QPen()	
func NewPen() *QPen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),92000,92102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPen{}
	_p.SetDriver(__rv,92000,true)
	return _p
} 
//QPen::QPen(QColor const&)	
func NewPenWithColor(color *QColor) *QPen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),92000,92103,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPen{}
	_p.SetDriver(__rv,92000,true)
	return _p
} 
//QPen::QPen(QPen const&)	
func NewPenCopy(pen *QPen) *QPen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),92000,92104,Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPen{}
	_p.SetDriver(__rv,92000,true)
	return _p
} 
//QPen::QPen(Qt::PenStyle)	
func NewPenWithPenstyle(value Qt_PenStyle) *QPen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),92000,92105,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPen{}
	_p.SetDriver(__rv,92000,true)
	return _p
} 
//QPen::QPen(QBrush const&,double,Qt::PenStyle,Qt::PenCapStyle,Qt::PenJoinStyle)	
func NewPenWithBrushWidthSCJ(brush *QBrush,width float64,s Qt_PenStyle,c Qt_PenCapStyle,j Qt_PenJoinStyle) *QPen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),92000,92106,Native(brush),unsafe.Pointer(&width),unsafe.Pointer(&s),unsafe.Pointer(&c),unsafe.Pointer(&j),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPen{}
	_p.SetDriver(__rv,92000,true)
	return _p
} 
//QPen::brush()
func (q *QPen) Brush() *QBrush {
	var __rv uintptr
	q.Drv(92000,92107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPen::capStyle()
func (q *QPen) CapStyle() Qt_PenCapStyle {
	var __rv Qt_PenCapStyle
	q.Drv(92000,92108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::color()
func (q *QPen) Color() *QColor {
	var __rv uintptr
	q.Drv(92000,92109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QPen::dashOffset()
func (q *QPen) DashOffset() float64 {
	var __rv float64
	q.Drv(92000,92110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::dashPattern()
func (q *QPen) DashPattern() []float64 {
	var __rv []float64
	q.Drv(92000,92111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::isCosmetic()
func (q *QPen) IsCosmetic() bool {
	var __rv bool
	q.Drv(92000,92112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::isDetached()
func (q *QPen) IsDetached() bool {
	var __rv bool
	q.Drv(92000,92113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::isSolid()
func (q *QPen) IsSolid() bool {
	var __rv bool
	q.Drv(92000,92114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::joinStyle()
func (q *QPen) JoinStyle() Qt_PenJoinStyle {
	var __rv Qt_PenJoinStyle
	q.Drv(92000,92115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::miterLimit()
func (q *QPen) MiterLimit() float64 {
	var __rv float64
	q.Drv(92000,92116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::setBrush(QBrush const&)
func (q *QPen) SetBrush(brush *QBrush)  {
	q.Drv(92000,92117,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setCapStyle(Qt::PenCapStyle)
func (q *QPen) SetCapStyle(pcs Qt_PenCapStyle)  {
	q.Drv(92000,92118,unsafe.Pointer(&pcs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setColor(QColor const&)
func (q *QPen) SetColor(color *QColor)  {
	q.Drv(92000,92119,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setCosmetic(bool)
func (q *QPen) SetCosmetic(cosmetic bool)  {
	q.Drv(92000,92120,unsafe.Pointer(&cosmetic),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setDashOffset(double)
func (q *QPen) SetDashOffset(doffset float64)  {
	q.Drv(92000,92121,unsafe.Pointer(&doffset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setDashPattern(QVector<qreal> const&)
func (q *QPen) SetDashPattern(pattern []float64)  {
	q.Drv(92000,92122,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setJoinStyle(Qt::PenJoinStyle)
func (q *QPen) SetJoinStyle(pcs Qt_PenJoinStyle)  {
	q.Drv(92000,92123,unsafe.Pointer(&pcs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setMiterLimit(double)
func (q *QPen) SetMiterLimit(limit float64)  {
	q.Drv(92000,92124,unsafe.Pointer(&limit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setStyle(Qt::PenStyle)
func (q *QPen) SetStyle(value Qt_PenStyle)  {
	q.Drv(92000,92125,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setWidth(int)
func (q *QPen) SetWidth(width int)  {
	q.Drv(92000,92126,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::setWidthF(double)
func (q *QPen) SetWidthF(width float64)  {
	q.Drv(92000,92127,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPen::style()
func (q *QPen) Style() Qt_PenStyle {
	var __rv Qt_PenStyle
	q.Drv(92000,92128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::width()
func (q *QPen) Width() int {
	var __rv int
	q.Drv(92000,92129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPen::widthF()
func (q *QPen) WidthF() float64 {
	var __rv float64
	q.Drv(92000,92130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
