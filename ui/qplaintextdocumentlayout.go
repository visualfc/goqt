// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPlainTextDocumentLayout : QPlainTextDocumentLayout
type QPlainTextDocumentLayout struct {
	QAbstractTextDocumentLayout
}
//QPlainTextDocumentLayout::QPlainTextDocumentLayout(QTextDocument*)	
func NewPlainTextDocumentLayout(document *QTextDocument) *QPlainTextDocumentLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),320000,320102,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPlainTextDocumentLayout{}
	_p.SetDriver(__rv,320000,false)
	return _p
} 
//QPlainTextDocumentLayout::blockBoundingRect(QTextBlock const&)
func (q *QPlainTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *QRectF {
	var __rv uintptr
	q.Drv(320000,320103,Native(block),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPlainTextDocumentLayout::cursorWidth()
func (q *QPlainTextDocumentLayout) CursorWidth() int {
	var __rv int
	q.Drv(320000,320104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextDocumentLayout::documentChanged(int,int,int)
func (q *QPlainTextDocumentLayout) DocumentChanged(from int,value2 int,charsAdded int)  {
	q.Drv(320000,320105,unsafe.Pointer(&from),unsafe.Pointer(&value2),unsafe.Pointer(&charsAdded),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextDocumentLayout::documentSize()
func (q *QPlainTextDocumentLayout) DocumentSize() *QSizeF {
	var __rv uintptr
	q.Drv(320000,320106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QPlainTextDocumentLayout::draw(QPainter*,QAbstractTextDocumentLayout::PaintContext const&)
func (q *QPlainTextDocumentLayout) Draw(value2 *QPainter,value3 *QAbstractTextDocumentLayoutPaintContext)  {
	q.Drv(320000,320107,Native(value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextDocumentLayout::ensureBlockLayout(QTextBlock const&)
func (q *QPlainTextDocumentLayout) EnsureBlockLayout(block *QTextBlock)  {
	q.Drv(320000,320108,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextDocumentLayout::frameBoundingRect(QTextFrame*)
func (q *QPlainTextDocumentLayout) FrameBoundingRect(value *QTextFrame) *QRectF {
	var __rv uintptr
	q.Drv(320000,320109,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPlainTextDocumentLayout::hitTest(QPointF const&,Qt::HitTestAccuracy)
func (q *QPlainTextDocumentLayout) HitTest(value2 *QPointF,value3 Qt_HitTestAccuracy) int {
	var __rv int
	q.Drv(320000,320110,Native(value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextDocumentLayout::pageCount()
func (q *QPlainTextDocumentLayout) PageCount() int {
	var __rv int
	q.Drv(320000,320111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextDocumentLayout::requestUpdate()
func (q *QPlainTextDocumentLayout) RequestUpdate()  {
	q.Drv(320000,320112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextDocumentLayout::setCursorWidth(int)
func (q *QPlainTextDocumentLayout) SetCursorWidth(width int)  {
	q.Drv(320000,320113,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
