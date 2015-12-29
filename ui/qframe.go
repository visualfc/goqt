// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFrame_StyleMask - QFrame::StyleMask
type QFrame_StyleMask uint32
const (
	QFrame_Shadow_Mask QFrame_StyleMask = 0x00f0
	QFrame_Shape_Mask QFrame_StyleMask = 0x000f
)
//enum QFrame_Shape - QFrame::Shape
type QFrame_Shape uint32
const (
	QFrame_NoFrame QFrame_Shape = 0
	QFrame_Box QFrame_Shape = 0x0001
	QFrame_Panel QFrame_Shape = 0x0002
	QFrame_WinPanel QFrame_Shape = 0x0003
	QFrame_HLine QFrame_Shape = 0x0004
	QFrame_VLine QFrame_Shape = 0x0005
	QFrame_StyledPanel QFrame_Shape = 0x0006
)
//enum QFrame_Shadow - QFrame::Shadow
type QFrame_Shadow uint32
const (
	QFrame_Plain QFrame_Shadow = 0x0010
	QFrame_Raised QFrame_Shadow = 0x0020
	QFrame_Sunken QFrame_Shadow = 0x0030
)
//struct QFrame : QFrame
type QFrame struct {
	QWidget
}
//QFrame::QFrame()	
func NewFrame() *QFrame {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),245000,245102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFrame{}
	_p.SetDriver(__rv,245000,false)
	return _p
} 
//QFrame::QFrame(QWidget*,QFlags<Qt::WindowType>)	
func NewFrameWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QFrame {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),245000,245103,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFrame{}
	_p.SetDriver(__rv,245000,false)
	return _p
} 
//QFrame::changeEvent(QEvent*)
func (q *QFrame) ChangeEvent(value *QEvent)  {
	q.Drv(245000,245104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::drawFrame(QPainter*)
func (q *QFrame) DrawFrame(value *QPainter)  {
	q.Drv(245000,245105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::event(QEvent*)
func (q *QFrame) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(245000,245106,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::frameRect()
func (q *QFrame) FrameRect() *QRect {
	var __rv uintptr
	q.Drv(245000,245107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFrame::frameShadow()
func (q *QFrame) FrameShadow() QFrame_Shadow {
	var __rv QFrame_Shadow
	q.Drv(245000,245108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::frameShape()
func (q *QFrame) FrameShape() QFrame_Shape {
	var __rv QFrame_Shape
	q.Drv(245000,245109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::frameStyle()
func (q *QFrame) FrameStyle() int {
	var __rv int
	q.Drv(245000,245110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::frameWidth()
func (q *QFrame) FrameWidth() int {
	var __rv int
	q.Drv(245000,245111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::lineWidth()
func (q *QFrame) LineWidth() int {
	var __rv int
	q.Drv(245000,245112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::midLineWidth()
func (q *QFrame) MidLineWidth() int {
	var __rv int
	q.Drv(245000,245113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFrame::paintEvent(QPaintEvent*)
func (q *QFrame) PaintEvent(value *QPaintEvent)  {
	q.Drv(245000,245114,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setFrameRect(QRect const&)
func (q *QFrame) SetFrameRect(value *QRect)  {
	q.Drv(245000,245115,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setFrameShadow(QFrame::Shadow)
func (q *QFrame) SetFrameShadow(value QFrame_Shadow)  {
	q.Drv(245000,245116,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setFrameShape(QFrame::Shape)
func (q *QFrame) SetFrameShape(value QFrame_Shape)  {
	q.Drv(245000,245117,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setFrameStyle(int)
func (q *QFrame) SetFrameStyle(value int)  {
	q.Drv(245000,245118,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setLineWidth(int)
func (q *QFrame) SetLineWidth(value int)  {
	q.Drv(245000,245119,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::setMidLineWidth(int)
func (q *QFrame) SetMidLineWidth(value int)  {
	q.Drv(245000,245120,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFrame::sizeHint()
func (q *QFrame) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(245000,245121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
