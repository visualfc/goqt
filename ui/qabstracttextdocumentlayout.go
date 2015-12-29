// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractTextDocumentLayout : QAbstractTextDocumentLayout
type QAbstractTextDocumentLayout struct {
	QObject
}
func (q *QAbstractTextDocumentLayout) OnPageCountChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(205000,205102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractTextDocumentLayout) OnUpdateBlock(fn func(*QTextBlock)) uintptr {
	var __rv uintptr
	q.Drv(205000,205103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractTextDocumentLayout) OnDocumentSizeChanged(fn func(*QSizeF)) uintptr {
	var __rv uintptr
	q.Drv(205000,205104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractTextDocumentLayout) OnUpdate(fn func()) uintptr {
	var __rv uintptr
	q.Drv(205000,205105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractTextDocumentLayout) OnUpdateEx(fn func(*QRectF)) uintptr {
	var __rv uintptr
	q.Drv(205000,205106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractTextDocumentLayout::anchorAt(QPointF const&)
func (q *QAbstractTextDocumentLayout) AnchorAt(pos *QPointF) string {
	var __rv string
	q.Drv(205000,205107,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTextDocumentLayout::blockBoundingRect(QTextBlock const&)
func (q *QAbstractTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *QRectF {
	var __rv uintptr
	q.Drv(205000,205108,Native(block),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::document()
func (q *QAbstractTextDocumentLayout) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(205000,205109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QAbstractTextDocumentLayout::documentChanged(int,int,int)
func (q *QAbstractTextDocumentLayout) DocumentChanged(from int,charsRemoved int,charsAdded int)  {
	q.Drv(205000,205110,unsafe.Pointer(&from),unsafe.Pointer(&charsRemoved),unsafe.Pointer(&charsAdded),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::documentSize()
func (q *QAbstractTextDocumentLayout) DocumentSize() *QSizeF {
	var __rv uintptr
	q.Drv(205000,205111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::draw(QPainter*,QAbstractTextDocumentLayout::PaintContext const&)
func (q *QAbstractTextDocumentLayout) Draw(painter *QPainter,context *QAbstractTextDocumentLayoutPaintContext)  {
	q.Drv(205000,205112,Native(painter),Native(context),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::drawInlineObject(QPainter*,QRectF const&,QTextInlineObject,int,QTextFormat const&)
func (q *QAbstractTextDocumentLayout) DrawInlineObject(painter *QPainter,rect *QRectF,object *QTextInlineObject,posInDocument int,format *QTextFormat)  {
	q.Drv(205000,205113,Native(painter),Native(rect),Native(object),unsafe.Pointer(&posInDocument),Native(format),nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::format(int)
func (q *QAbstractTextDocumentLayout) Format(pos int) *QTextCharFormat {
	var __rv uintptr
	q.Drv(205000,205114,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::formatIndex(int)
func (q *QAbstractTextDocumentLayout) FormatIndex(pos int) int {
	var __rv int
	q.Drv(205000,205115,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTextDocumentLayout::frameBoundingRect(QTextFrame*)
func (q *QAbstractTextDocumentLayout) FrameBoundingRect(frame *QTextFrame) *QRectF {
	var __rv uintptr
	q.Drv(205000,205116,Native(frame),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::handlerForObject(int)
func (q *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface {
	var __rv uintptr
	q.Drv(205000,205117,unsafe.Pointer(&objectType),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextObjectInterface{}
	_rp.SetDriver(__rv,163000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::hitTest(QPointF const&,Qt::HitTestAccuracy)
func (q *QAbstractTextDocumentLayout) HitTest(point *QPointF,accuracy Qt_HitTestAccuracy) int {
	var __rv int
	q.Drv(205000,205118,Native(point),unsafe.Pointer(&accuracy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTextDocumentLayout::pageCount()
func (q *QAbstractTextDocumentLayout) PageCount() int {
	var __rv int
	q.Drv(205000,205119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractTextDocumentLayout::paintDevice()
func (q *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice {
	var __rv uintptr
	q.Drv(205000,205120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QAbstractTextDocumentLayout::positionInlineObject(QTextInlineObject,int,QTextFormat const&)
func (q *QAbstractTextDocumentLayout) PositionInlineObject(item *QTextInlineObject,posInDocument int,format *QTextFormat)  {
	q.Drv(205000,205121,Native(item),unsafe.Pointer(&posInDocument),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::registerHandler(int,QObject*)
func (q *QAbstractTextDocumentLayout) RegisterHandler(objectType int,component QObjectInterface)  {
	q.Drv(205000,205122,unsafe.Pointer(&objectType),Native(component),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::resizeInlineObject(QTextInlineObject,int,QTextFormat const&)
func (q *QAbstractTextDocumentLayout) ResizeInlineObject(item *QTextInlineObject,posInDocument int,format *QTextFormat)  {
	q.Drv(205000,205123,Native(item),unsafe.Pointer(&posInDocument),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractTextDocumentLayout::setPaintDevice(QPaintDevice*)
func (q *QAbstractTextDocumentLayout) SetPaintDevice(device QPaintDeviceInterface)  {
	q.Drv(205000,205124,unsafe.Pointer(new_pd_head(device)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
