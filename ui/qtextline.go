// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextLine_Edge - QTextLine::Edge
type QTextLine_Edge uint32
const (
	QTextLine_Leading QTextLine_Edge = 0
	QTextLine_Trailing QTextLine_Edge = 1
)
//enum QTextLine_CursorPosition - QTextLine::CursorPosition
type QTextLine_CursorPosition uint32
const (
	QTextLine_CursorBetweenCharacters QTextLine_CursorPosition = 0
	QTextLine_CursorOnCharacter QTextLine_CursorPosition = 1
)
//struct QTextLine : QTextLine
type QTextLine struct {
	BaseDrv
}
//QTextLine::QTextLine()	
func NewTextLine() *QTextLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),161000,161102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLine{}
	_p.SetDriver(__rv,161000,true)
	return _p
} 
//QTextLine::ascent()
func (q *QTextLine) Ascent() float64 {
	var __rv float64
	q.Drv(161000,161103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::cursorToX(int)
func (q *QTextLine) CursorToX(cursorPos int) float64 {
	var __rv float64
	q.Drv(161000,161104,unsafe.Pointer(&cursorPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::cursorToX(int*)
func (q *QTextLine) CursorToXWithCursorposp(cursorPos *int) float64 {
	var __rv float64
	q.Drv(161000,161105,unsafe.Pointer(&cursorPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::cursorToX(int*,QTextLine::Edge)
func (q *QTextLine) CursorToXWithCursorpospEdge(cursorPos *int,edge QTextLine_Edge) float64 {
	var __rv float64
	q.Drv(161000,161106,unsafe.Pointer(&cursorPos),unsafe.Pointer(&edge),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::cursorToX(int,QTextLine::Edge)
func (q *QTextLine) CursorToXWithCursorposEdge(cursorPos int,edge QTextLine_Edge) float64 {
	var __rv float64
	q.Drv(161000,161107,unsafe.Pointer(&cursorPos),unsafe.Pointer(&edge),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::descent()
func (q *QTextLine) Descent() float64 {
	var __rv float64
	q.Drv(161000,161108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::draw(QPainter*,QPointF const&,QTextLayout::FormatRange const*)
func (q *QTextLine) Draw(p *QPainter,point *QPointF,selection *QTextLayoutFormatRange)  {
	q.Drv(161000,161109,Native(p),Native(point),Native(selection),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::height()
func (q *QTextLine) Height() float64 {
	var __rv float64
	q.Drv(161000,161110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::horizontalAdvance()
func (q *QTextLine) HorizontalAdvance() float64 {
	var __rv float64
	q.Drv(161000,161111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::isValid()
func (q *QTextLine) IsValid() bool {
	var __rv bool
	q.Drv(161000,161112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::leading()
func (q *QTextLine) Leading() float64 {
	var __rv float64
	q.Drv(161000,161113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::leadingIncluded()
func (q *QTextLine) LeadingIncluded() bool {
	var __rv bool
	q.Drv(161000,161114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::lineNumber()
func (q *QTextLine) LineNumber() int {
	var __rv int
	q.Drv(161000,161115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::naturalTextRect()
func (q *QTextLine) NaturalTextRect() *QRectF {
	var __rv uintptr
	q.Drv(161000,161116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTextLine::naturalTextWidth()
func (q *QTextLine) NaturalTextWidth() float64 {
	var __rv float64
	q.Drv(161000,161117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::position()
func (q *QTextLine) Position() *QPointF {
	var __rv uintptr
	q.Drv(161000,161118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTextLine::rect()
func (q *QTextLine) Rect() *QRectF {
	var __rv uintptr
	q.Drv(161000,161119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTextLine::setLeadingIncluded(bool)
func (q *QTextLine) SetLeadingIncluded(included bool)  {
	q.Drv(161000,161120,unsafe.Pointer(&included),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::setLineWidth(double)
func (q *QTextLine) SetLineWidth(width float64)  {
	q.Drv(161000,161121,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::setNumColumns(int)
func (q *QTextLine) SetNumColumns(columns int)  {
	q.Drv(161000,161122,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::setNumColumns(int,double)
func (q *QTextLine) SetNumColumnsFWithColumnsAlignmentwidth(columns int,alignmentWidth float64)  {
	q.Drv(161000,161123,unsafe.Pointer(&columns),unsafe.Pointer(&alignmentWidth),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::setPosition(QPointF const&)
func (q *QTextLine) SetPosition(pos *QPointF)  {
	q.Drv(161000,161124,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLine::textLength()
func (q *QTextLine) TextLength() int {
	var __rv int
	q.Drv(161000,161125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::textStart()
func (q *QTextLine) TextStart() int {
	var __rv int
	q.Drv(161000,161126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::width()
func (q *QTextLine) Width() float64 {
	var __rv float64
	q.Drv(161000,161127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::x()
func (q *QTextLine) X() float64 {
	var __rv float64
	q.Drv(161000,161128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::xToCursor(double)
func (q *QTextLine) XToCursor(x float64) int {
	var __rv int
	q.Drv(161000,161129,unsafe.Pointer(&x),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::xToCursor(double,QTextLine::CursorPosition)
func (q *QTextLine) XToCursorFWithXCursorposition(x float64,value2 QTextLine_CursorPosition) int {
	var __rv int
	q.Drv(161000,161130,unsafe.Pointer(&x),unsafe.Pointer(&value2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLine::y()
func (q *QTextLine) Y() float64 {
	var __rv float64
	q.Drv(161000,161131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
