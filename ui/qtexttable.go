// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextTable : QTextTable
type QTextTable struct {
	QTextFrame
}
//QTextTable::QTextTable(QTextDocument*)	
func NewTextTable(doc *QTextDocument) *QTextTable {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),377000,377102,Native(doc),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextTable{}
	_p.SetDriver(__rv,377000,false)
	return _p
} 
//QTextTable::appendColumns(int)
func (q *QTextTable) AppendColumns(count int)  {
	q.Drv(377000,377103,unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::appendRows(int)
func (q *QTextTable) AppendRows(count int)  {
	q.Drv(377000,377104,unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::cellAt(QTextCursor const&)
func (q *QTextTable) CellAt(c *QTextCursor) *QTextTableCell {
	var __rv uintptr
	q.Drv(377000,377105,Native(c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableCell{}
	_rp.SetDriver(__rv,166000,true)
	return _rp
}	
//QTextTable::cellAt(int)
func (q *QTextTable) CellAtWithPosition(position int) *QTextTableCell {
	var __rv uintptr
	q.Drv(377000,377106,unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableCell{}
	_rp.SetDriver(__rv,166000,true)
	return _rp
}	
//QTextTable::cellAt(int,int)
func (q *QTextTable) CellAtWithRowCol(row int,col int) *QTextTableCell {
	var __rv uintptr
	q.Drv(377000,377107,unsafe.Pointer(&row),unsafe.Pointer(&col),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableCell{}
	_rp.SetDriver(__rv,166000,true)
	return _rp
}	
//QTextTable::columns()
func (q *QTextTable) Columns() int {
	var __rv int
	q.Drv(377000,377108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTable::format()
func (q *QTextTable) Format() *QTextTableFormat {
	var __rv uintptr
	q.Drv(377000,377109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTableFormat{}
	_rp.SetDriver(__rv,168000,true)
	return _rp
}	
//QTextTable::insertColumns(int,int)
func (q *QTextTable) InsertColumns(pos int,num int)  {
	q.Drv(377000,377110,unsafe.Pointer(&pos),unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::insertRows(int,int)
func (q *QTextTable) InsertRows(pos int,num int)  {
	q.Drv(377000,377111,unsafe.Pointer(&pos),unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::mergeCells(QTextCursor const&)
func (q *QTextTable) MergeCells(cursor *QTextCursor)  {
	q.Drv(377000,377112,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::mergeCells(int,int,int,int)
func (q *QTextTable) MergeCellsWithRowColNumrowsNumcols(row int,col int,numRows int,numCols int)  {
	q.Drv(377000,377113,unsafe.Pointer(&row),unsafe.Pointer(&col),unsafe.Pointer(&numRows),unsafe.Pointer(&numCols),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::removeColumns(int,int)
func (q *QTextTable) RemoveColumns(pos int,num int)  {
	q.Drv(377000,377114,unsafe.Pointer(&pos),unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::removeRows(int,int)
func (q *QTextTable) RemoveRows(pos int,num int)  {
	q.Drv(377000,377115,unsafe.Pointer(&pos),unsafe.Pointer(&num),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::resize(int,int)
func (q *QTextTable) Resize(rows int,cols int)  {
	q.Drv(377000,377116,unsafe.Pointer(&rows),unsafe.Pointer(&cols),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::rowEnd(QTextCursor const&)
func (q *QTextTable) RowEnd(c *QTextCursor) *QTextCursor {
	var __rv uintptr
	q.Drv(377000,377117,Native(c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextTable::rowStart(QTextCursor const&)
func (q *QTextTable) RowStart(c *QTextCursor) *QTextCursor {
	var __rv uintptr
	q.Drv(377000,377118,Native(c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextTable::rows()
func (q *QTextTable) Rows() int {
	var __rv int
	q.Drv(377000,377119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTable::setFormat(QTextTableFormat const&)
func (q *QTextTable) SetFormat(format *QTextTableFormat)  {
	q.Drv(377000,377120,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTable::splitCell(int,int,int,int)
func (q *QTextTable) SplitCell(row int,col int,numRows int,numCols int)  {
	q.Drv(377000,377121,unsafe.Pointer(&row),unsafe.Pointer(&col),unsafe.Pointer(&numRows),unsafe.Pointer(&numCols),nil,nil,nil,nil,nil,nil,nil,nil)
}	
