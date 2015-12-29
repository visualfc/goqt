// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDialog_DialogCode - QDialog::DialogCode
type QDialog_DialogCode uint32
const (
	QDialog_Rejected QDialog_DialogCode = 0
	QDialog_Accepted QDialog_DialogCode = 1
)
//struct QDialog : QDialog
type QDialog struct {
	QWidget
}
func (q *QDialog) OnFinished(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(226000,226102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDialog) OnAccepted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(226000,226103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDialog) OnRejected(fn func()) uintptr {
	var __rv uintptr
	q.Drv(226000,226104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDialog::QDialog()	
func NewDialog() *QDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),226000,226105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialog{}
	_p.SetDriver(__rv,226000,false)
	return _p
} 
//QDialog::QDialog(QWidget*,QFlags<Qt::WindowType>)	
func NewDialogWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),226000,226106,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialog{}
	_p.SetDriver(__rv,226000,false)
	return _p
} 
//QDialog::accept()
func (q *QDialog) Accept()  {
	q.Drv(226000,226107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::adjustPosition(QWidget*)
func (q *QDialog) AdjustPosition(value QWidgetInterface)  {
	q.Drv(226000,226108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::closeEvent(QCloseEvent*)
func (q *QDialog) CloseEvent(value *QCloseEvent)  {
	q.Drv(226000,226109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::contextMenuEvent(QContextMenuEvent*)
func (q *QDialog) ContextMenuEvent(value *QContextMenuEvent)  {
	q.Drv(226000,226110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::done(int)
func (q *QDialog) Done(value int)  {
	q.Drv(226000,226111,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::eventFilter(QObject*,QEvent*)
func (q *QDialog) EventFilter(value2 QObjectInterface,value3 *QEvent) bool {
	var __rv bool
	q.Drv(226000,226112,Native(value2),Native(value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialog::exec()
func (q *QDialog) Exec() int {
	var __rv int
	q.Drv(226000,226113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialog::extension()
func (q *QDialog) Extension() *QWidget {
	var __rv uintptr
	q.Drv(226000,226114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDialog::isSizeGripEnabled()
func (q *QDialog) IsSizeGripEnabled() bool {
	var __rv bool
	q.Drv(226000,226115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialog::keyPressEvent(QKeyEvent*)
func (q *QDialog) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(226000,226116,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::minimumSizeHint()
func (q *QDialog) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(226000,226117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QDialog::open()
func (q *QDialog) Open()  {
	q.Drv(226000,226118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::orientation()
func (q *QDialog) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(226000,226119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialog::reject()
func (q *QDialog) Reject()  {
	q.Drv(226000,226120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::resizeEvent(QResizeEvent*)
func (q *QDialog) ResizeEvent(value *QResizeEvent)  {
	q.Drv(226000,226121,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::result()
func (q *QDialog) Result() int {
	var __rv int
	q.Drv(226000,226122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialog::setExtension(QWidget*)
func (q *QDialog) SetExtension(extension QWidgetInterface)  {
	q.Drv(226000,226123,Native(extension),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::setModal(bool)
func (q *QDialog) SetModal(modal bool)  {
	q.Drv(226000,226124,unsafe.Pointer(&modal),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::setOrientation(Qt::Orientation)
func (q *QDialog) SetOrientation(orientation Qt_Orientation)  {
	q.Drv(226000,226125,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::setResult(int)
func (q *QDialog) SetResult(r int)  {
	q.Drv(226000,226126,unsafe.Pointer(&r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::setSizeGripEnabled(bool)
func (q *QDialog) SetSizeGripEnabled(value bool)  {
	q.Drv(226000,226127,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::setVisible(bool)
func (q *QDialog) SetVisible(visible bool)  {
	q.Drv(226000,226128,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::showEvent(QShowEvent*)
func (q *QDialog) ShowEvent(value *QShowEvent)  {
	q.Drv(226000,226129,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::showExtension(bool)
func (q *QDialog) ShowExtension(value bool)  {
	q.Drv(226000,226130,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialog::sizeHint()
func (q *QDialog) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(226000,226131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
