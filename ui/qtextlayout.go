// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextLayout_CursorMode - QTextLayout::CursorMode
type QTextLayout_CursorMode uint32
const (
	QTextLayout_SkipCharacters QTextLayout_CursorMode = 0
	QTextLayout_SkipWords QTextLayout_CursorMode = 1
)
//struct QTextLayout : QTextLayout
type QTextLayout struct {
	BaseDrv
}
//QTextLayout::QTextLayout()	
func NewTextLayout() *QTextLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),158000,158102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLayout{}
	_p.SetDriver(__rv,158000,true)
	return _p
} 
//QTextLayout::QTextLayout(QString const&)	
func NewTextLayoutWithText(text string) *QTextLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),158000,158103,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLayout{}
	_p.SetDriver(__rv,158000,true)
	return _p
} 
//QTextLayout::QTextLayout(QTextBlock const&)	
func NewTextLayoutWithTextblock(b *QTextBlock) *QTextLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),158000,158104,Native(b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLayout{}
	_p.SetDriver(__rv,158000,true)
	return _p
} 
//QTextLayout::QTextLayout(QString const&,QFont const&,QPaintDevice*)	
func NewTextLayoutWithTextFontPaintDevice(text string,font *QFont,paintdevice QPaintDeviceInterface) *QTextLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),158000,158105,unsafe.Pointer(&text),Native(font),unsafe.Pointer(new_pd_head(paintdevice)),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextLayout{}
	_p.SetDriver(__rv,158000,true)
	return _p
} 
//QTextLayout::additionalFormats()
func (q *QTextLayout) AdditionalFormats() []*QTextLayoutFormatRange {
	var __rv []*QTextLayoutFormatRange
	q.Drv(158000,158106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::beginLayout()
func (q *QTextLayout) BeginLayout()  {
	q.Drv(158000,158107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::boundingRect()
func (q *QTextLayout) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(158000,158108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QTextLayout::cacheEnabled()
func (q *QTextLayout) CacheEnabled() bool {
	var __rv bool
	q.Drv(158000,158109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::clearAdditionalFormats()
func (q *QTextLayout) ClearAdditionalFormats()  {
	q.Drv(158000,158110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::clearLayout()
func (q *QTextLayout) ClearLayout()  {
	q.Drv(158000,158111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::createLine()
func (q *QTextLayout) CreateLine() *QTextLine {
	var __rv uintptr
	q.Drv(158000,158112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLine{}
	_rp.SetDriver(__rv,161000,true)
	return _rp
}	
//QTextLayout::draw(QPainter*,QPointF const&,QVector<QTextLayout::FormatRange> const&,QRectF const&)
func (q *QTextLayout) Draw(p *QPainter,pos *QPointF,selections []*QTextLayoutFormatRange,clip *QRectF)  {
	q.Drv(158000,158113,Native(p),Native(pos),unsafe.Pointer(&selections),Native(clip),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::drawCursor(QPainter*,QPointF const&,int)
func (q *QTextLayout) DrawCursorFWithPPosCursorposition(p *QPainter,pos *QPointF,cursorPosition int)  {
	q.Drv(158000,158114,Native(p),Native(pos),unsafe.Pointer(&cursorPosition),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::drawCursor(QPainter*,QPointF const&,int,int)
func (q *QTextLayout) DrawCursorFWithPPosCursorpositionWidth(p *QPainter,pos *QPointF,cursorPosition int,width int)  {
	q.Drv(158000,158115,Native(p),Native(pos),unsafe.Pointer(&cursorPosition),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::endLayout()
func (q *QTextLayout) EndLayout()  {
	q.Drv(158000,158116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::font()
func (q *QTextLayout) Font() *QFont {
	var __rv uintptr
	q.Drv(158000,158117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTextLayout::isValidCursorPosition(int)
func (q *QTextLayout) IsValidCursorPosition(pos int) bool {
	var __rv bool
	q.Drv(158000,158118,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::lineAt(int)
func (q *QTextLayout) LineAt(i int) *QTextLine {
	var __rv uintptr
	q.Drv(158000,158119,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLine{}
	_rp.SetDriver(__rv,161000,true)
	return _rp
}	
//QTextLayout::lineCount()
func (q *QTextLayout) LineCount() int {
	var __rv int
	q.Drv(158000,158120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::lineForTextPosition(int)
func (q *QTextLayout) LineForTextPosition(pos int) *QTextLine {
	var __rv uintptr
	q.Drv(158000,158121,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLine{}
	_rp.SetDriver(__rv,161000,true)
	return _rp
}	
//QTextLayout::maximumWidth()
func (q *QTextLayout) MaximumWidth() float64 {
	var __rv float64
	q.Drv(158000,158122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::minimumWidth()
func (q *QTextLayout) MinimumWidth() float64 {
	var __rv float64
	q.Drv(158000,158123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::nextCursorPosition(int)
func (q *QTextLayout) NextCursorPosition(oldPos int) int {
	var __rv int
	q.Drv(158000,158124,unsafe.Pointer(&oldPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::nextCursorPosition(int,QTextLayout::CursorMode)
func (q *QTextLayout) NextCursorPositionWithOldposMode(oldPos int,mode QTextLayout_CursorMode) int {
	var __rv int
	q.Drv(158000,158125,unsafe.Pointer(&oldPos),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::position()
func (q *QTextLayout) Position() *QPointF {
	var __rv uintptr
	q.Drv(158000,158126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTextLayout::preeditAreaPosition()
func (q *QTextLayout) PreeditAreaPosition() int {
	var __rv int
	q.Drv(158000,158127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::preeditAreaText()
func (q *QTextLayout) PreeditAreaText() string {
	var __rv string
	q.Drv(158000,158128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::previousCursorPosition(int)
func (q *QTextLayout) PreviousCursorPosition(oldPos int) int {
	var __rv int
	q.Drv(158000,158129,unsafe.Pointer(&oldPos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::previousCursorPosition(int,QTextLayout::CursorMode)
func (q *QTextLayout) PreviousCursorPositionWithOldposMode(oldPos int,mode QTextLayout_CursorMode) int {
	var __rv int
	q.Drv(158000,158130,unsafe.Pointer(&oldPos),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::setAdditionalFormats(QList<QTextLayout::FormatRange> const&)
func (q *QTextLayout) SetAdditionalFormats(overrides []*QTextLayoutFormatRange)  {
	q.Drv(158000,158131,unsafe.Pointer(&overrides),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setCacheEnabled(bool)
func (q *QTextLayout) SetCacheEnabled(enable bool)  {
	q.Drv(158000,158132,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setFlags(int)
func (q *QTextLayout) SetFlags(flags int)  {
	q.Drv(158000,158133,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setFont(QFont const&)
func (q *QTextLayout) SetFont(f *QFont)  {
	q.Drv(158000,158134,Native(f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setPosition(QPointF const&)
func (q *QTextLayout) SetPosition(p *QPointF)  {
	q.Drv(158000,158135,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setPreeditArea(int,QString const&)
func (q *QTextLayout) SetPreeditArea(position int,text string)  {
	q.Drv(158000,158136,unsafe.Pointer(&position),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setText(QString const&)
func (q *QTextLayout) SetText(string string)  {
	q.Drv(158000,158137,unsafe.Pointer(&string),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::setTextOption(QTextOption const&)
func (q *QTextLayout) SetTextOption(option *QTextOption)  {
	q.Drv(158000,158138,Native(option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextLayout::text()
func (q *QTextLayout) Text() string {
	var __rv string
	q.Drv(158000,158139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextLayout::textOption()
func (q *QTextLayout) TextOption() *QTextOption {
	var __rv uintptr
	q.Drv(158000,158140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextOption{}
	_rp.SetDriver(__rv,164000,true)
	return _rp
}	
