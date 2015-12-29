// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextTableCellFormat : QTextTableCellFormat
type QTextTableCellFormat struct {
	QTextCharFormat
}
//QTextTableCellFormat::QTextTableCellFormat()	
func NewTextTableCellFormat() *QTextTableCellFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),167000,167102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextTableCellFormat{}
	_p.SetDriver(__rv,167000,true)
	return _p
} 
//QTextTableCellFormat::bottomPadding()
func (q *QTextTableCellFormat) BottomPadding() float64 {
	var __rv float64
	q.Drv(167000,167103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCellFormat::isValid()
func (q *QTextTableCellFormat) IsValid() bool {
	var __rv bool
	q.Drv(167000,167104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCellFormat::leftPadding()
func (q *QTextTableCellFormat) LeftPadding() float64 {
	var __rv float64
	q.Drv(167000,167105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCellFormat::rightPadding()
func (q *QTextTableCellFormat) RightPadding() float64 {
	var __rv float64
	q.Drv(167000,167106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextTableCellFormat::setBottomPadding(double)
func (q *QTextTableCellFormat) SetBottomPadding(padding float64)  {
	q.Drv(167000,167107,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCellFormat::setLeftPadding(double)
func (q *QTextTableCellFormat) SetLeftPadding(padding float64)  {
	q.Drv(167000,167108,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCellFormat::setPadding(double)
func (q *QTextTableCellFormat) SetPadding(padding float64)  {
	q.Drv(167000,167109,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCellFormat::setRightPadding(double)
func (q *QTextTableCellFormat) SetRightPadding(padding float64)  {
	q.Drv(167000,167110,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCellFormat::setTopPadding(double)
func (q *QTextTableCellFormat) SetTopPadding(padding float64)  {
	q.Drv(167000,167111,unsafe.Pointer(&padding),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextTableCellFormat::topPadding()
func (q *QTextTableCellFormat) TopPadding() float64 {
	var __rv float64
	q.Drv(167000,167112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
