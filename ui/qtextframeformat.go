// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextFrameFormat_Position - QTextFrameFormat::Position
type QTextFrameFormat_Position uint32
const (
	QTextFrameFormat_InFlow QTextFrameFormat_Position = 0
	QTextFrameFormat_FloatLeft QTextFrameFormat_Position = 1
	QTextFrameFormat_FloatRight QTextFrameFormat_Position = 2
)
//enum QTextFrameFormat_BorderStyle - QTextFrameFormat::BorderStyle
type QTextFrameFormat_BorderStyle uint32
const (
	QTextFrameFormat_BorderStyle_None QTextFrameFormat_BorderStyle = 0
	QTextFrameFormat_BorderStyle_Dotted QTextFrameFormat_BorderStyle = 1
	QTextFrameFormat_BorderStyle_Dashed QTextFrameFormat_BorderStyle = 2
	QTextFrameFormat_BorderStyle_Solid QTextFrameFormat_BorderStyle = 3
	QTextFrameFormat_BorderStyle_Double QTextFrameFormat_BorderStyle = 4
	QTextFrameFormat_BorderStyle_DotDash QTextFrameFormat_BorderStyle = 5
	QTextFrameFormat_BorderStyle_DotDotDash QTextFrameFormat_BorderStyle = 6
	QTextFrameFormat_BorderStyle_Groove QTextFrameFormat_BorderStyle = 7
	QTextFrameFormat_BorderStyle_Ridge QTextFrameFormat_BorderStyle = 8
	QTextFrameFormat_BorderStyle_Inset QTextFrameFormat_BorderStyle = 9
	QTextFrameFormat_BorderStyle_Outset QTextFrameFormat_BorderStyle = 10
)
//struct QTextFrameFormat : QTextFrameFormat
type QTextFrameFormat struct {
	QTextFormat
}
//QTextFrameFormat::QTextFrameFormat()	
func NewTextFrameFormat() *QTextFrameFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),154000,154102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFrameFormat{}
	_p.SetDriver(__rv,154000,true)
	return _p
} 
//QTextFrameFormat::border()
func (q *QTextFrameFormat) Border() float64 {
	var __rv float64
	q.Drv(154000,154103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::borderBrush()
func (q *QTextFrameFormat) BorderBrush() *QBrush {
	var __rv uintptr
	q.Drv(154000,154104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QTextFrameFormat::borderStyle()
func (q *QTextFrameFormat) BorderStyle() QTextFrameFormat_BorderStyle {
	var __rv QTextFrameFormat_BorderStyle
	q.Drv(154000,154105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::bottomMargin()
func (q *QTextFrameFormat) BottomMargin() float64 {
	var __rv float64
	q.Drv(154000,154106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::height()
func (q *QTextFrameFormat) Height() *QTextLength {
	var __rv uintptr
	q.Drv(154000,154107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLength{}
	_rp.SetDriver(__rv,160000,true)
	return _rp
}	
//QTextFrameFormat::isValid()
func (q *QTextFrameFormat) IsValid() bool {
	var __rv bool
	q.Drv(154000,154108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::leftMargin()
func (q *QTextFrameFormat) LeftMargin() float64 {
	var __rv float64
	q.Drv(154000,154109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::margin()
func (q *QTextFrameFormat) Margin() float64 {
	var __rv float64
	q.Drv(154000,154110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::padding()
func (q *QTextFrameFormat) Padding() float64 {
	var __rv float64
	q.Drv(154000,154111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::pageBreakPolicy()
func (q *QTextFrameFormat) PageBreakPolicy() QTextFormat_PageBreakFlag {
	var __rv QTextFormat_PageBreakFlag
	q.Drv(154000,154112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::position()
func (q *QTextFrameFormat) Position() QTextFrameFormat_Position {
	var __rv QTextFrameFormat_Position
	q.Drv(154000,154113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::rightMargin()
func (q *QTextFrameFormat) RightMargin() float64 {
	var __rv float64
	q.Drv(154000,154114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::setBorder(double)
func (q *QTextFrameFormat) SetBorder(border float64)  {
	q.Drv(154000,154115,unsafe.Pointer(&border),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setBorderBrush(QBrush const&)
func (q *QTextFrameFormat) SetBorderBrush(brush *QBrush)  {
	q.Drv(154000,154116,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setBorderStyle(QTextFrameFormat::BorderStyle)
func (q *QTextFrameFormat) SetBorderStyle(style QTextFrameFormat_BorderStyle)  {
	q.Drv(154000,154117,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setBottomMargin(double)
func (q *QTextFrameFormat) SetBottomMargin(margin float64)  {
	q.Drv(154000,154118,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setHeight(QTextLength const&)
func (q *QTextFrameFormat) SetHeight(height *QTextLength)  {
	q.Drv(154000,154119,Native(height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setHeight(double)
func (q *QTextFrameFormat) SetHeightFWithHeight(height float64)  {
	q.Drv(154000,154120,unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setLeftMargin(double)
func (q *QTextFrameFormat) SetLeftMargin(margin float64)  {
	q.Drv(154000,154121,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setMargin(double)
func (q *QTextFrameFormat) SetMargin(margin float64)  {
	q.Drv(154000,154122,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setPadding(double)
func (q *QTextFrameFormat) SetPadding(padding float64)  {
	q.Drv(154000,154123,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setPageBreakPolicy(QFlags<QTextFormat::PageBreakFlag>)
func (q *QTextFrameFormat) SetPageBreakPolicy(flags QTextFormat_PageBreakFlag)  {
	q.Drv(154000,154124,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setPosition(QTextFrameFormat::Position)
func (q *QTextFrameFormat) SetPosition(f QTextFrameFormat_Position)  {
	q.Drv(154000,154125,unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setRightMargin(double)
func (q *QTextFrameFormat) SetRightMargin(margin float64)  {
	q.Drv(154000,154126,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setTopMargin(double)
func (q *QTextFrameFormat) SetTopMargin(margin float64)  {
	q.Drv(154000,154127,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setWidth(QTextLength const&)
func (q *QTextFrameFormat) SetWidth(length *QTextLength)  {
	q.Drv(154000,154128,Native(length),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::setWidth(double)
func (q *QTextFrameFormat) SetWidthFWithWidth(width float64)  {
	q.Drv(154000,154129,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextFrameFormat::topMargin()
func (q *QTextFrameFormat) TopMargin() float64 {
	var __rv float64
	q.Drv(154000,154130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrameFormat::width()
func (q *QTextFrameFormat) Width() *QTextLength {
	var __rv uintptr
	q.Drv(154000,154131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLength{}
	_rp.SetDriver(__rv,160000,true)
	return _rp
}	
