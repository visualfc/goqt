// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractPrintDialog_PrintRange - QAbstractPrintDialog::PrintRange
type QAbstractPrintDialog_PrintRange uint32
const (
	QAbstractPrintDialog_AllPages QAbstractPrintDialog_PrintRange = 0
	QAbstractPrintDialog_Selection QAbstractPrintDialog_PrintRange = 1
	QAbstractPrintDialog_PageRange QAbstractPrintDialog_PrintRange = 2
	QAbstractPrintDialog_CurrentPage QAbstractPrintDialog_PrintRange = 3
)
//enum QAbstractPrintDialog_PrintDialogOption - QAbstractPrintDialog::PrintDialogOption
type QAbstractPrintDialog_PrintDialogOption uint32
const (
	QAbstractPrintDialog_None QAbstractPrintDialog_PrintDialogOption = 0x0000
	QAbstractPrintDialog_PrintToFile QAbstractPrintDialog_PrintDialogOption = 0x0001
	QAbstractPrintDialog_PrintSelection QAbstractPrintDialog_PrintDialogOption = 0x0002
	QAbstractPrintDialog_PrintPageRange QAbstractPrintDialog_PrintDialogOption = 0x0004
	QAbstractPrintDialog_PrintShowPageSize QAbstractPrintDialog_PrintDialogOption = 0x0008
	QAbstractPrintDialog_PrintCollateCopies QAbstractPrintDialog_PrintDialogOption = 0x0010
	QAbstractPrintDialog_DontUseSheet QAbstractPrintDialog_PrintDialogOption = 0x0020
	QAbstractPrintDialog_PrintCurrentPage QAbstractPrintDialog_PrintDialogOption = 0x0040
)
//struct QAbstractPrintDialog : QAbstractPrintDialog
type QAbstractPrintDialog struct {
	QDialog
}
//QAbstractPrintDialog::addEnabledOption(QAbstractPrintDialog::PrintDialogOption)
func (q *QAbstractPrintDialog) AddEnabledOption(option QAbstractPrintDialog_PrintDialogOption)  {
	q.Drv(198000,198102,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::enabledOptions()
func (q *QAbstractPrintDialog) EnabledOptions() QAbstractPrintDialog_PrintDialogOption {
	var __rv QAbstractPrintDialog_PrintDialogOption
	q.Drv(198000,198103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::exec()
func (q *QAbstractPrintDialog) Exec() int {
	var __rv int
	q.Drv(198000,198104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::fromPage()
func (q *QAbstractPrintDialog) FromPage() int {
	var __rv int
	q.Drv(198000,198105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::isOptionEnabled(QAbstractPrintDialog::PrintDialogOption)
func (q *QAbstractPrintDialog) IsOptionEnabled(option QAbstractPrintDialog_PrintDialogOption) bool {
	var __rv bool
	q.Drv(198000,198106,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::maxPage()
func (q *QAbstractPrintDialog) MaxPage() int {
	var __rv int
	q.Drv(198000,198107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::minPage()
func (q *QAbstractPrintDialog) MinPage() int {
	var __rv int
	q.Drv(198000,198108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::printRange()
func (q *QAbstractPrintDialog) PrintRange() QAbstractPrintDialog_PrintRange {
	var __rv QAbstractPrintDialog_PrintRange
	q.Drv(198000,198109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractPrintDialog::printer()
func (q *QAbstractPrintDialog) Printer() *QPrinter {
	var __rv uintptr
	q.Drv(198000,198110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPrinter{}
	_rp.SetDriver(__rv,104000,true)
	return _rp
}	
//QAbstractPrintDialog::setEnabledOptions(QFlags<QAbstractPrintDialog::PrintDialogOption>)
func (q *QAbstractPrintDialog) SetEnabledOptions(options QAbstractPrintDialog_PrintDialogOption)  {
	q.Drv(198000,198111,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::setFromTo(int,int)
func (q *QAbstractPrintDialog) SetFromTo(fromPage int,toPage int)  {
	q.Drv(198000,198112,unsafe.Pointer(&fromPage),unsafe.Pointer(&toPage),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::setMinMax(int,int)
func (q *QAbstractPrintDialog) SetMinMax(min int,max int)  {
	q.Drv(198000,198113,unsafe.Pointer(&min),unsafe.Pointer(&max),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::setOptionTabs(QList<QWidget*> const&)
func (q *QAbstractPrintDialog) SetOptionTabs(tabs []QWidgetInterface)  {
	q.Drv(198000,198114,unsafe.Pointer(&tabs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::setPrintRange(QAbstractPrintDialog::PrintRange)
func (q *QAbstractPrintDialog) SetPrintRange(_range QAbstractPrintDialog_PrintRange)  {
	q.Drv(198000,198115,unsafe.Pointer(&_range),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractPrintDialog::toPage()
func (q *QAbstractPrintDialog) ToPage() int {
	var __rv int
	q.Drv(198000,198116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
