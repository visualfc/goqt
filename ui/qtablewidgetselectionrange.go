// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTableWidgetSelectionRange : QTableWidgetSelectionRange
type QTableWidgetSelectionRange struct {
	BaseDrv
}
//QTableWidgetSelectionRange::QTableWidgetSelectionRange()	
func NewTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),135000,135102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetSelectionRange{}
	_p.SetDriver(__rv,135000,true)
	return _p
} 
//QTableWidgetSelectionRange::QTableWidgetSelectionRange(QTableWidgetSelectionRange const&)	
func NewTableWidgetSelectionRangeCopy(other *QTableWidgetSelectionRange) *QTableWidgetSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),135000,135103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetSelectionRange{}
	_p.SetDriver(__rv,135000,true)
	return _p
} 
//QTableWidgetSelectionRange::QTableWidgetSelectionRange(int,int,int,int)	
func NewTableWidgetSelectionRangeWithTopLeftBottomRight(top int,left int,bottom int,right int) *QTableWidgetSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),135000,135104,unsafe.Pointer(&top),unsafe.Pointer(&left),unsafe.Pointer(&bottom),unsafe.Pointer(&right),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidgetSelectionRange{}
	_p.SetDriver(__rv,135000,true)
	return _p
} 
//QTableWidgetSelectionRange::bottomRow()
func (q *QTableWidgetSelectionRange) BottomRow() int {
	var __rv int
	q.Drv(135000,135105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetSelectionRange::columnCount()
func (q *QTableWidgetSelectionRange) ColumnCount() int {
	var __rv int
	q.Drv(135000,135106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetSelectionRange::leftColumn()
func (q *QTableWidgetSelectionRange) LeftColumn() int {
	var __rv int
	q.Drv(135000,135107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetSelectionRange::rightColumn()
func (q *QTableWidgetSelectionRange) RightColumn() int {
	var __rv int
	q.Drv(135000,135108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetSelectionRange::rowCount()
func (q *QTableWidgetSelectionRange) RowCount() int {
	var __rv int
	q.Drv(135000,135109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidgetSelectionRange::topRow()
func (q *QTableWidgetSelectionRange) TopRow() int {
	var __rv int
	q.Drv(135000,135110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
