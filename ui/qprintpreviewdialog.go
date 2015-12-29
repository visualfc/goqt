// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPrintPreviewDialog : QPrintPreviewDialog
type QPrintPreviewDialog struct {
	QDialog
}
func (q *QPrintPreviewDialog) OnPaintRequested(fn func(*QPrinter)) uintptr {
	var __rv uintptr
	q.Drv(324000,324102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QPrintPreviewDialog::QPrintPreviewDialog()	
func NewPrintPreviewDialog() *QPrintPreviewDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),324000,324103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewDialog{}
	_p.SetDriver(__rv,324000,false)
	return _p
} 
//QPrintPreviewDialog::QPrintPreviewDialog(QWidget*,QFlags<Qt::WindowType>)	
func NewPrintPreviewDialogWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QPrintPreviewDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),324000,324104,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewDialog{}
	_p.SetDriver(__rv,324000,false)
	return _p
} 
//QPrintPreviewDialog::QPrintPreviewDialog(QPrinter*,QWidget*,QFlags<Qt::WindowType>)	
func NewPrintPreviewDialogWithPrinterParentFlags(printer *QPrinter,parent QWidgetInterface,flags Qt_WindowType) *QPrintPreviewDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),324000,324105,Native(printer),Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewDialog{}
	_p.SetDriver(__rv,324000,false)
	return _p
} 
//QPrintPreviewDialog::done(int)
func (q *QPrintPreviewDialog) Done(result int)  {
	q.Drv(324000,324106,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewDialog::open()
func (q *QPrintPreviewDialog) Open()  {
	q.Drv(324000,324107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewDialog::open(QObject*,char const*)
func (q *QPrintPreviewDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(324000,324108,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewDialog::printer()
func (q *QPrintPreviewDialog) Printer() *QPrinter {
	var __rv uintptr
	q.Drv(324000,324109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPrinter{}
	_rp.SetDriver(__rv,104000,true)
	return _rp
}	
//QPrintPreviewDialog::setVisible(bool)
func (q *QPrintPreviewDialog) SetVisible(visible bool)  {
	q.Drv(324000,324110,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
