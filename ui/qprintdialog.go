// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPrintDialog : QPrintDialog
type QPrintDialog struct {
	QAbstractPrintDialog
}
func (q *QPrintDialog) OnAccepted(fn func(*QPrinter)) uintptr {
	var __rv uintptr
	q.Drv(323000,323102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QPrintDialog::QPrintDialog()	
func NewPrintDialog() *QPrintDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),323000,323103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintDialog{}
	_p.SetDriver(__rv,323000,false)
	return _p
} 
//QPrintDialog::QPrintDialog(QWidget*)	
func NewPrintDialogWithParent(parent QWidgetInterface) *QPrintDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),323000,323104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintDialog{}
	_p.SetDriver(__rv,323000,false)
	return _p
} 
//QPrintDialog::QPrintDialog(QPrinter*,QWidget*)	
func NewPrintDialogWithPrinterParent(printer *QPrinter,parent QWidgetInterface) *QPrintDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),323000,323105,Native(printer),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintDialog{}
	_p.SetDriver(__rv,323000,false)
	return _p
} 
//QPrintDialog::accepted()
func (q *QPrintDialog) Accepted()  {
	q.Drv(323000,323106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::done(int)
func (q *QPrintDialog) Done(result int)  {
	q.Drv(323000,323107,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::exec()
func (q *QPrintDialog) Exec() int {
	var __rv int
	q.Drv(323000,323108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintDialog::open()
func (q *QPrintDialog) Open()  {
	q.Drv(323000,323109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::open(QObject*,char const*)
func (q *QPrintDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(323000,323110,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::options()
func (q *QPrintDialog) Options() QAbstractPrintDialog_PrintDialogOption {
	var __rv QAbstractPrintDialog_PrintDialogOption
	q.Drv(323000,323111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintDialog::setOption(QAbstractPrintDialog::PrintDialogOption)
func (q *QPrintDialog) SetOption(option QAbstractPrintDialog_PrintDialogOption)  {
	q.Drv(323000,323112,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::setOption(QAbstractPrintDialog::PrintDialogOption,bool)
func (q *QPrintDialog) SetOptionWithOptionOn(option QAbstractPrintDialog_PrintDialogOption,on bool)  {
	q.Drv(323000,323113,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::setOptions(QFlags<QAbstractPrintDialog::PrintDialogOption>)
func (q *QPrintDialog) SetOptions(options QAbstractPrintDialog_PrintDialogOption)  {
	q.Drv(323000,323114,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::setVisible(bool)
func (q *QPrintDialog) SetVisible(visible bool)  {
	q.Drv(323000,323115,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintDialog::testOption(QAbstractPrintDialog::PrintDialogOption)
func (q *QPrintDialog) TestOption(option QAbstractPrintDialog_PrintDialogOption) bool {
	var __rv bool
	q.Drv(323000,323116,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
