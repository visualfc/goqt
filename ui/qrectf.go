// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QRectF : QRectF
type QRectF struct {
	BaseDrv
}
//QRectF::QRectF()	
func NewRectF() *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::QRectF(QRect const&)	
func NewRectFWithRect(rect *QRect) *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111103,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::QRectF(QRectF const&)	
func NewRectFCopy(other *QRectF) *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::QRectF(QPointF const&,QPointF const&)	
func NewRectFWithTopleftBottomright(topleft *QPointF,bottomRight *QPointF) *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111105,Native(topleft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::QRectF(QPointF const&,QSizeF const&)	
func NewRectFWithTopleftSize(topleft *QPointF,size *QSizeF) *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111106,Native(topleft),Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::QRectF(double,double,double,double)	
func NewRectFWithLeftTopWidthHeight(left float64,top float64,width float64,height float64) *QRectF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),111000,111107,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&width),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRectF{}
	_p.SetDriver(__rv,111000,true)
	return _p
} 
//QRectF::adjust(double,double,double,double)
func (q *QRectF) Adjust(x1 float64,y1 float64,x2 float64,y2 float64)  {
	q.Drv(111000,111108,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::adjusted(double,double,double,double)
func (q *QRectF) Adjusted(x1 float64,y1 float64,x2 float64,y2 float64) *QRectF {
	var __rv uintptr
	q.Drv(111000,111109,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::bottom()
func (q *QRectF) Bottom() float64 {
	var __rv float64
	q.Drv(111000,111110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::bottomLeft()
func (q *QRectF) BottomLeft() *QPointF {
	var __rv uintptr
	q.Drv(111000,111111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRectF::bottomRight()
func (q *QRectF) BottomRight() *QPointF {
	var __rv uintptr
	q.Drv(111000,111112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRectF::center()
func (q *QRectF) Center() *QPointF {
	var __rv uintptr
	q.Drv(111000,111113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRectF::contains(QPointF const&)
func (q *QRectF) Contains(p *QPointF) bool {
	var __rv bool
	q.Drv(111000,111114,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::contains(QRectF const&)
func (q *QRectF) ContainsFWithRectf(r *QRectF) bool {
	var __rv bool
	q.Drv(111000,111115,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::contains(double,double)
func (q *QRectF) ContainsFWithXY(x float64,y float64) bool {
	var __rv bool
	q.Drv(111000,111116,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::getCoords(double*,double*,double*,double*)
func (q *QRectF) GetCoords(x1 *float64,y1 *float64,x2 *float64,y2 *float64)  {
	q.Drv(111000,111117,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::getRect(double*,double*,double*,double*)
func (q *QRectF) GetRect(x *float64,y *float64,w *float64,h *float64)  {
	q.Drv(111000,111118,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::height()
func (q *QRectF) Height() float64 {
	var __rv float64
	q.Drv(111000,111119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::intersected(QRectF const&)
func (q *QRectF) Intersected(other *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(111000,111120,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::intersects(QRectF const&)
func (q *QRectF) Intersects(r *QRectF) bool {
	var __rv bool
	q.Drv(111000,111121,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::isEmpty()
func (q *QRectF) IsEmpty() bool {
	var __rv bool
	q.Drv(111000,111122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::isNull()
func (q *QRectF) IsNull() bool {
	var __rv bool
	q.Drv(111000,111123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::isValid()
func (q *QRectF) IsValid() bool {
	var __rv bool
	q.Drv(111000,111124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::left()
func (q *QRectF) Left() float64 {
	var __rv float64
	q.Drv(111000,111125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::moveBottom(double)
func (q *QRectF) MoveBottom(pos float64)  {
	q.Drv(111000,111126,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveBottomLeft(QPointF const&)
func (q *QRectF) MoveBottomLeft(p *QPointF)  {
	q.Drv(111000,111127,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveBottomRight(QPointF const&)
func (q *QRectF) MoveBottomRight(p *QPointF)  {
	q.Drv(111000,111128,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveCenter(QPointF const&)
func (q *QRectF) MoveCenter(p *QPointF)  {
	q.Drv(111000,111129,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveLeft(double)
func (q *QRectF) MoveLeft(pos float64)  {
	q.Drv(111000,111130,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveRight(double)
func (q *QRectF) MoveRight(pos float64)  {
	q.Drv(111000,111131,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveTo(QPointF const&)
func (q *QRectF) MoveTo(p *QPointF)  {
	q.Drv(111000,111132,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveTo(double,double)
func (q *QRectF) MoveToFWithXFloat64(x float64,t float64)  {
	q.Drv(111000,111133,unsafe.Pointer(&x),unsafe.Pointer(&t),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveTop(double)
func (q *QRectF) MoveTop(pos float64)  {
	q.Drv(111000,111134,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveTopLeft(QPointF const&)
func (q *QRectF) MoveTopLeft(p *QPointF)  {
	q.Drv(111000,111135,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::moveTopRight(QPointF const&)
func (q *QRectF) MoveTopRight(p *QPointF)  {
	q.Drv(111000,111136,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::normalized()
func (q *QRectF) Normalized() *QRectF {
	var __rv uintptr
	q.Drv(111000,111137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::right()
func (q *QRectF) Right() float64 {
	var __rv float64
	q.Drv(111000,111138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::setBottom(double)
func (q *QRectF) SetBottom(pos float64)  {
	q.Drv(111000,111139,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setBottomLeft(QPointF const&)
func (q *QRectF) SetBottomLeft(p *QPointF)  {
	q.Drv(111000,111140,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setBottomRight(QPointF const&)
func (q *QRectF) SetBottomRight(p *QPointF)  {
	q.Drv(111000,111141,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setCoords(double,double,double,double)
func (q *QRectF) SetCoords(x1 float64,y1 float64,x2 float64,y2 float64)  {
	q.Drv(111000,111142,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setHeight(double)
func (q *QRectF) SetHeight(h float64)  {
	q.Drv(111000,111143,unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setLeft(double)
func (q *QRectF) SetLeft(pos float64)  {
	q.Drv(111000,111144,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setRect(double,double,double,double)
func (q *QRectF) SetRect(x float64,y float64,w float64,h float64)  {
	q.Drv(111000,111145,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setRight(double)
func (q *QRectF) SetRight(pos float64)  {
	q.Drv(111000,111146,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setSize(QSizeF const&)
func (q *QRectF) SetSize(s *QSizeF)  {
	q.Drv(111000,111147,Native(s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setTop(double)
func (q *QRectF) SetTop(pos float64)  {
	q.Drv(111000,111148,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setTopLeft(QPointF const&)
func (q *QRectF) SetTopLeft(p *QPointF)  {
	q.Drv(111000,111149,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setTopRight(QPointF const&)
func (q *QRectF) SetTopRight(p *QPointF)  {
	q.Drv(111000,111150,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setWidth(double)
func (q *QRectF) SetWidth(w float64)  {
	q.Drv(111000,111151,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setX(double)
func (q *QRectF) SetX(pos float64)  {
	q.Drv(111000,111152,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::setY(double)
func (q *QRectF) SetY(pos float64)  {
	q.Drv(111000,111153,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::size()
func (q *QRectF) Size() *QSizeF {
	var __rv uintptr
	q.Drv(111000,111154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QRectF::toAlignedRect()
func (q *QRectF) ToAlignedRect() *QRect {
	var __rv uintptr
	q.Drv(111000,111155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRectF::toRect()
func (q *QRectF) ToRect() *QRect {
	var __rv uintptr
	q.Drv(111000,111156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRectF::top()
func (q *QRectF) Top() float64 {
	var __rv float64
	q.Drv(111000,111157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::topLeft()
func (q *QRectF) TopLeft() *QPointF {
	var __rv uintptr
	q.Drv(111000,111158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRectF::topRight()
func (q *QRectF) TopRight() *QPointF {
	var __rv uintptr
	q.Drv(111000,111159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QRectF::translate(QPointF const&)
func (q *QRectF) Translate(p *QPointF)  {
	q.Drv(111000,111160,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::translate(double,double)
func (q *QRectF) TranslateFWithDxDy(dx float64,dy float64)  {
	q.Drv(111000,111161,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRectF::translated(QPointF const&)
func (q *QRectF) Translated(p *QPointF) *QRectF {
	var __rv uintptr
	q.Drv(111000,111162,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::translated(double,double)
func (q *QRectF) TranslatedFWithDxDy(dx float64,dy float64) *QRectF {
	var __rv uintptr
	q.Drv(111000,111163,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::united(QRectF const&)
func (q *QRectF) United(other *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(111000,111164,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QRectF::width()
func (q *QRectF) Width() float64 {
	var __rv float64
	q.Drv(111000,111165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::x()
func (q *QRectF) X() float64 {
	var __rv float64
	q.Drv(111000,111166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRectF::y()
func (q *QRectF) Y() float64 {
	var __rv float64
	q.Drv(111000,111167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
