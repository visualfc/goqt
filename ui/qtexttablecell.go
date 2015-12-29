// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextTableCell : QTextTableCell
type QTextTableCell struct {
	BaseDrv
}
//QTextTableCell::QTextTableCell()	
func NewTextTableCell() *QTextTableCell {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),166000,166102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextTableCell{}
	_p.SetDriver(__rv,166000,true)
	return _p
} 
//QTextTableCell::QTextTableCell(QTextTableCell const&)	
func NewTextTableCellCopy(o *QTextTableCell) *QTextTableCell {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),166000,166103,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextTableCell{}
	_p.SetDriver(__rv,166000,true)
	return _p
} 
//QTextTableCell::begin()
func (q *QTextTableCell) Begin() *QTextFrameiterator {
	var __rv uintptr
	q.Drv(166000,166104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameiterator{}
	_rp.SetDriver(__rv,153000,true)
	return _rp
}	
//QTextTableCell::column()
func (q *QTextTableCell) Column() int {
	var __rv int
	q.Drv(166000,166105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::columnSpan()
func (q *QTextTableCell) ColumnSpan() int {
	var __rv int
	q.Drv(166000,166106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::end()
func (q *QTextTableCell) End() *QTextFrameiterator {
	var __rv uintptr
	q.Drv(166000,166107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameiterator{}
	_rp.SetDriver(__rv,153000,true)
	return _rp
}	
//QTextTableCell::firstCursorPosition()
func (q *QTextTableCell) FirstCursorPosition() *QTextCursor {
	var __rv uintptr
	q.Drv(166000,166108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextTableCell::firstPosition()
func (q *QTextTableCell) FirstPosition() int {
	var __rv int
	q.Drv(166000,166109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::format()
func (q *QTextTableCell) Format() *QTextCharFormat {
	var __rv uintptr
	q.Drv(166000,166110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextTableCell::isValid()
func (q *QTextTableCell) IsValid() bool {
	var __rv bool
	q.Drv(166000,166111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::lastCursorPosition()
func (q *QTextTableCell) LastCursorPosition() *QTextCursor {
	var __rv uintptr
	q.Drv(166000,166112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextTableCell::lastPosition()
func (q *QTextTableCell) LastPosition() int {
	var __rv int
	q.Drv(166000,166113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::row()
func (q *QTextTableCell) Row() int {
	var __rv int
	q.Drv(166000,166114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::rowSpan()
func (q *QTextTableCell) RowSpan() int {
	var __rv int
	q.Drv(166000,166115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCell::setFormat(QTextCharFormat const&)
func (q *QTextTableCell) SetFormat(format *QTextCharFormat)  {
	q.Drv(166000,166116,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCell::tableCellFormatIndex()
func (q *QTextTableCell) TableCellFormatIndex() int {
	var __rv int
	q.Drv(166000,166117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
