// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBlockFormat : QTextBlockFormat
type QTextBlockFormat struct {
	QTextFormat
}
//QTextBlockFormat::QTextBlockFormat()	
func NewTextBlockFormat() *QTextBlockFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),139000,139102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlockFormat{}
	_p.SetDriver(__rv,139000,true)
	return _p
} 
//QTextBlockFormat::alignment()
func (q *QTextBlockFormat) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(139000,139103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::bottomMargin()
func (q *QTextBlockFormat) BottomMargin() float64 {
	var __rv float64
	q.Drv(139000,139104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::indent()
func (q *QTextBlockFormat) Indent() int {
	var __rv int
	q.Drv(139000,139105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::isValid()
func (q *QTextBlockFormat) IsValid() bool {
	var __rv bool
	q.Drv(139000,139106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::leftMargin()
func (q *QTextBlockFormat) LeftMargin() float64 {
	var __rv float64
	q.Drv(139000,139107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::nonBreakableLines()
func (q *QTextBlockFormat) NonBreakableLines() bool {
	var __rv bool
	q.Drv(139000,139108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::pageBreakPolicy()
func (q *QTextBlockFormat) PageBreakPolicy() QTextFormat_PageBreakFlag {
	var __rv QTextFormat_PageBreakFlag
	q.Drv(139000,139109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::rightMargin()
func (q *QTextBlockFormat) RightMargin() float64 {
	var __rv float64
	q.Drv(139000,139110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QTextBlockFormat) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(139000,139111,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setBottomMargin(double)
func (q *QTextBlockFormat) SetBottomMargin(margin float64)  {
	q.Drv(139000,139112,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setIndent(int)
func (q *QTextBlockFormat) SetIndent(indent int)  {
	q.Drv(139000,139113,unsafe.Pointer(&indent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setLeftMargin(double)
func (q *QTextBlockFormat) SetLeftMargin(margin float64)  {
	q.Drv(139000,139114,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setNonBreakableLines(bool)
func (q *QTextBlockFormat) SetNonBreakableLines(b bool)  {
	q.Drv(139000,139115,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setPageBreakPolicy(QFlags<QTextFormat::PageBreakFlag>)
func (q *QTextBlockFormat) SetPageBreakPolicy(flags QTextFormat_PageBreakFlag)  {
	q.Drv(139000,139116,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setRightMargin(double)
func (q *QTextBlockFormat) SetRightMargin(margin float64)  {
	q.Drv(139000,139117,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setTabPositions(QList<QTextOption::Tab> const&)
func (q *QTextBlockFormat) SetTabPositions(tabs []*QTextOptionTab)  {
	q.Drv(139000,139118,unsafe.Pointer(&tabs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setTextIndent(double)
func (q *QTextBlockFormat) SetTextIndent(aindent float64)  {
	q.Drv(139000,139119,unsafe.Pointer(&aindent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::setTopMargin(double)
func (q *QTextBlockFormat) SetTopMargin(margin float64)  {
	q.Drv(139000,139120,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlockFormat::tabPositions()
func (q *QTextBlockFormat) TabPositions() []*QTextOptionTab {
	var __rv []*QTextOptionTab
	q.Drv(139000,139121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::textIndent()
func (q *QTextBlockFormat) TextIndent() float64 {
	var __rv float64
	q.Drv(139000,139122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlockFormat::topMargin()
func (q *QTextBlockFormat) TopMargin() float64 {
	var __rv float64
	q.Drv(139000,139123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
