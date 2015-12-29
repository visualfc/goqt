// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextTableFormat : QTextTableFormat
type QTextTableFormat struct {
	QTextFrameFormat
}
//QTextTableFormat::QTextTableFormat()	
func NewTextTableFormat() *QTextTableFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),168000,168102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextTableFormat{}
	_p.SetDriver(__rv,168000,true)
	return _p
} 
//QTextTableFormat::alignment()
func (q *QTextTableFormat) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(168000,168103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::cellPadding()
func (q *QTextTableFormat) CellPadding() float64 {
	var __rv float64
	q.Drv(168000,168104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::cellSpacing()
func (q *QTextTableFormat) CellSpacing() float64 {
	var __rv float64
	q.Drv(168000,168105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::clearColumnWidthConstraints()
func (q *QTextTableFormat) ClearColumnWidthConstraints()  {
	q.Drv(168000,168106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::columnWidthConstraints()
func (q *QTextTableFormat) ColumnWidthConstraints() []*QTextLength {
	var __rv []*QTextLength
	q.Drv(168000,168107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::columns()
func (q *QTextTableFormat) Columns() int {
	var __rv int
	q.Drv(168000,168108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::headerRowCount()
func (q *QTextTableFormat) HeaderRowCount() int {
	var __rv int
	q.Drv(168000,168109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::isValid()
func (q *QTextTableFormat) IsValid() bool {
	var __rv bool
	q.Drv(168000,168110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableFormat::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QTextTableFormat) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(168000,168111,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::setCellPadding(double)
func (q *QTextTableFormat) SetCellPadding(padding float64)  {
	q.Drv(168000,168112,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::setCellSpacing(double)
func (q *QTextTableFormat) SetCellSpacing(spacing float64)  {
	q.Drv(168000,168113,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::setColumnWidthConstraints(QVector<QTextLength> const&)
func (q *QTextTableFormat) SetColumnWidthConstraints(constraints []*QTextLength)  {
	q.Drv(168000,168114,unsafe.Pointer(&constraints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::setColumns(int)
func (q *QTextTableFormat) SetColumns(columns int)  {
	q.Drv(168000,168115,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableFormat::setHeaderRowCount(int)
func (q *QTextTableFormat) SetHeaderRowCount(count int)  {
	q.Drv(168000,168116,unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
