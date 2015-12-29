// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QRect : QRect
type QRect struct {
	BaseDrv
}
//QRect::QRect()	
func NewRect() *QRect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),110000,110102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRect{}
	_p.SetDriver(__rv,110000,true)
	return _p
} 
//QRect::QRect(QRect const&)	
func NewRectCopy(other *QRect) *QRect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),110000,110103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRect{}
	_p.SetDriver(__rv,110000,true)
	return _p
} 
//QRect::QRect(QPoint const&,QPoint const&)	
func NewRectWithTopleftBottomright(topleft *QPoint,bottomright *QPoint) *QRect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),110000,110104,Native(topleft),Native(bottomright),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRect{}
	_p.SetDriver(__rv,110000,true)
	return _p
} 
//QRect::QRect(QPoint const&,QSize const&)	
func NewRectWithTopleftSize(topleft *QPoint,size *QSize) *QRect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),110000,110105,Native(topleft),Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRect{}
	_p.SetDriver(__rv,110000,true)
	return _p
} 
//QRect::QRect(int,int,int,int)	
func NewRectWithLeftTopWidthHeight(left int,top int,width int,height int) *QRect {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),110000,110106,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&width),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRect{}
	_p.SetDriver(__rv,110000,true)
	return _p
} 
//QRect::adjust(int,int,int,int)
func (q *QRect) Adjust(x1 int,y1 int,x2 int,y2 int)  {
	q.Drv(110000,110107,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::adjusted(int,int,int,int)
func (q *QRect) Adjusted(x1 int,y1 int,x2 int,y2 int) *QRect {
	var __rv uintptr
	q.Drv(110000,110108,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::bottom()
func (q *QRect) Bottom() int {
	var __rv int
	q.Drv(110000,110109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::bottomLeft()
func (q *QRect) BottomLeft() *QPoint {
	var __rv uintptr
	q.Drv(110000,110110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QRect::bottomRight()
func (q *QRect) BottomRight() *QPoint {
	var __rv uintptr
	q.Drv(110000,110111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QRect::center()
func (q *QRect) Center() *QPoint {
	var __rv uintptr
	q.Drv(110000,110112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QRect::contains(QPoint const&)
func (q *QRect) Contains(p *QPoint) bool {
	var __rv bool
	q.Drv(110000,110113,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::contains(QRect const&)
func (q *QRect) ContainsWithRect(r *QRect) bool {
	var __rv bool
	q.Drv(110000,110114,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::contains(QPoint const&,bool)
func (q *QRect) ContainsWithPointProper(p *QPoint,proper bool) bool {
	var __rv bool
	q.Drv(110000,110115,Native(p),unsafe.Pointer(&proper),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::contains(QRect const&,bool)
func (q *QRect) ContainsWithRectProper(r *QRect,proper bool) bool {
	var __rv bool
	q.Drv(110000,110116,Native(r),unsafe.Pointer(&proper),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::contains(int,int)
func (q *QRect) ContainsWithXY(x int,y int) bool {
	var __rv bool
	q.Drv(110000,110117,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::contains(int,int,bool)
func (q *QRect) ContainsWithXYProper(x int,y int,proper bool) bool {
	var __rv bool
	q.Drv(110000,110118,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&proper),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::getCoords(int*,int*,int*,int*)
func (q *QRect) GetCoords(x1 *int,y1 *int,x2 *int,y2 *int)  {
	q.Drv(110000,110119,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::getRect(int*,int*,int*,int*)
func (q *QRect) GetRect(x *int,y *int,w *int,h *int)  {
	q.Drv(110000,110120,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::height()
func (q *QRect) Height() int {
	var __rv int
	q.Drv(110000,110121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::intersected(QRect const&)
func (q *QRect) Intersected(other *QRect) *QRect {
	var __rv uintptr
	q.Drv(110000,110122,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::intersects(QRect const&)
func (q *QRect) Intersects(r *QRect) bool {
	var __rv bool
	q.Drv(110000,110123,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::isEmpty()
func (q *QRect) IsEmpty() bool {
	var __rv bool
	q.Drv(110000,110124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::isNull()
func (q *QRect) IsNull() bool {
	var __rv bool
	q.Drv(110000,110125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::isValid()
func (q *QRect) IsValid() bool {
	var __rv bool
	q.Drv(110000,110126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::left()
func (q *QRect) Left() int {
	var __rv int
	q.Drv(110000,110127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::moveBottom(int)
func (q *QRect) MoveBottom(pos int)  {
	q.Drv(110000,110128,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveBottomLeft(QPoint const&)
func (q *QRect) MoveBottomLeft(p *QPoint)  {
	q.Drv(110000,110129,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveBottomRight(QPoint const&)
func (q *QRect) MoveBottomRight(p *QPoint)  {
	q.Drv(110000,110130,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveCenter(QPoint const&)
func (q *QRect) MoveCenter(p *QPoint)  {
	q.Drv(110000,110131,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveLeft(int)
func (q *QRect) MoveLeft(pos int)  {
	q.Drv(110000,110132,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveRight(int)
func (q *QRect) MoveRight(pos int)  {
	q.Drv(110000,110133,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveTo(QPoint const&)
func (q *QRect) MoveTo(p *QPoint)  {
	q.Drv(110000,110134,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveTo(int,int)
func (q *QRect) MoveToWithXY(x int,y int)  {
	q.Drv(110000,110135,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveTop(int)
func (q *QRect) MoveTop(pos int)  {
	q.Drv(110000,110136,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveTopLeft(QPoint const&)
func (q *QRect) MoveTopLeft(p *QPoint)  {
	q.Drv(110000,110137,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::moveTopRight(QPoint const&)
func (q *QRect) MoveTopRight(p *QPoint)  {
	q.Drv(110000,110138,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::normalized()
func (q *QRect) Normalized() *QRect {
	var __rv uintptr
	q.Drv(110000,110139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::right()
func (q *QRect) Right() int {
	var __rv int
	q.Drv(110000,110140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::setBottom(int)
func (q *QRect) SetBottom(pos int)  {
	q.Drv(110000,110141,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setBottomLeft(QPoint const&)
func (q *QRect) SetBottomLeft(p *QPoint)  {
	q.Drv(110000,110142,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setBottomRight(QPoint const&)
func (q *QRect) SetBottomRight(p *QPoint)  {
	q.Drv(110000,110143,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setCoords(int,int,int,int)
func (q *QRect) SetCoords(x1 int,y1 int,x2 int,y2 int)  {
	q.Drv(110000,110144,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setHeight(int)
func (q *QRect) SetHeight(h int)  {
	q.Drv(110000,110145,unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setLeft(int)
func (q *QRect) SetLeft(pos int)  {
	q.Drv(110000,110146,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setRect(int,int,int,int)
func (q *QRect) SetRect(x int,y int,w int,h int)  {
	q.Drv(110000,110147,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setRight(int)
func (q *QRect) SetRight(pos int)  {
	q.Drv(110000,110148,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setSize(QSize const&)
func (q *QRect) SetSize(s *QSize)  {
	q.Drv(110000,110149,Native(s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setTop(int)
func (q *QRect) SetTop(pos int)  {
	q.Drv(110000,110150,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setTopLeft(QPoint const&)
func (q *QRect) SetTopLeft(p *QPoint)  {
	q.Drv(110000,110151,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setTopRight(QPoint const&)
func (q *QRect) SetTopRight(p *QPoint)  {
	q.Drv(110000,110152,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setWidth(int)
func (q *QRect) SetWidth(w int)  {
	q.Drv(110000,110153,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setX(int)
func (q *QRect) SetX(x int)  {
	q.Drv(110000,110154,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::setY(int)
func (q *QRect) SetY(y int)  {
	q.Drv(110000,110155,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::size()
func (q *QRect) Size() *QSize {
	var __rv uintptr
	q.Drv(110000,110156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QRect::top()
func (q *QRect) Top() int {
	var __rv int
	q.Drv(110000,110157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::topLeft()
func (q *QRect) TopLeft() *QPoint {
	var __rv uintptr
	q.Drv(110000,110158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QRect::topRight()
func (q *QRect) TopRight() *QPoint {
	var __rv uintptr
	q.Drv(110000,110159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QRect::translate(QPoint const&)
func (q *QRect) Translate(p *QPoint)  {
	q.Drv(110000,110160,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::translate(int,int)
func (q *QRect) TranslateWithDxDy(dx int,dy int)  {
	q.Drv(110000,110161,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRect::translated(QPoint const&)
func (q *QRect) Translated(p *QPoint) *QRect {
	var __rv uintptr
	q.Drv(110000,110162,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::translated(int,int)
func (q *QRect) TranslatedWithDxDy(dx int,dy int) *QRect {
	var __rv uintptr
	q.Drv(110000,110163,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::united(QRect const&)
func (q *QRect) United(other *QRect) *QRect {
	var __rv uintptr
	q.Drv(110000,110164,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QRect::width()
func (q *QRect) Width() int {
	var __rv int
	q.Drv(110000,110165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::x()
func (q *QRect) X() int {
	var __rv int
	q.Drv(110000,110166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRect::y()
func (q *QRect) Y() int {
	var __rv int
	q.Drv(110000,110167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
