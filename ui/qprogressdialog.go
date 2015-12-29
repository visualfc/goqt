// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QProgressDialog : QProgressDialog
type QProgressDialog struct {
	QDialog
}
func (q *QProgressDialog) OnCanceled(fn func()) uintptr {
	var __rv uintptr
	q.Drv(328000,328102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QProgressDialog::QProgressDialog()	
func NewProgressDialog() *QProgressDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),328000,328103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProgressDialog{}
	_p.SetDriver(__rv,328000,false)
	return _p
} 
//QProgressDialog::QProgressDialog(QWidget*,QFlags<Qt::WindowType>)	
func NewProgressDialogWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QProgressDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),328000,328104,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProgressDialog{}
	_p.SetDriver(__rv,328000,false)
	return _p
} 
//QProgressDialog::QProgressDialog(QString const&,QString const&,int,int,QWidget*,QFlags<Qt::WindowType>)	
func NewProgressDialogWithLabeltextCancelbuttontextMinimumMaximumParentFlags(labelText string,cancelButtonText string,minimum int,maximum int,parent QWidgetInterface,flags Qt_WindowType) *QProgressDialog {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),328000,328105,unsafe.Pointer(&labelText),unsafe.Pointer(&cancelButtonText),unsafe.Pointer(&minimum),unsafe.Pointer(&maximum),Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProgressDialog{}
	_p.SetDriver(__rv,328000,false)
	return _p
} 
//QProgressDialog::autoClose()
func (q *QProgressDialog) AutoClose() bool {
	var __rv bool
	q.Drv(328000,328106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::autoReset()
func (q *QProgressDialog) AutoReset() bool {
	var __rv bool
	q.Drv(328000,328107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::cancel()
func (q *QProgressDialog) Cancel()  {
	q.Drv(328000,328108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::changeEvent(QEvent*)
func (q *QProgressDialog) ChangeEvent(event *QEvent)  {
	q.Drv(328000,328109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::closeEvent(QCloseEvent*)
func (q *QProgressDialog) CloseEvent(event *QCloseEvent)  {
	q.Drv(328000,328110,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::forceShow()
func (q *QProgressDialog) ForceShow()  {
	q.Drv(328000,328111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::labelText()
func (q *QProgressDialog) LabelText() string {
	var __rv string
	q.Drv(328000,328112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::maximum()
func (q *QProgressDialog) Maximum() int {
	var __rv int
	q.Drv(328000,328113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::minimum()
func (q *QProgressDialog) Minimum() int {
	var __rv int
	q.Drv(328000,328114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::minimumDuration()
func (q *QProgressDialog) MinimumDuration() int {
	var __rv int
	q.Drv(328000,328115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::open()
func (q *QProgressDialog) Open()  {
	q.Drv(328000,328116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::open(QObject*,char const*)
func (q *QProgressDialog) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(328000,328117,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::reset()
func (q *QProgressDialog) Reset()  {
	q.Drv(328000,328118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::resizeEvent(QResizeEvent*)
func (q *QProgressDialog) ResizeEvent(event *QResizeEvent)  {
	q.Drv(328000,328119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setAutoClose(bool)
func (q *QProgressDialog) SetAutoClose(close bool)  {
	q.Drv(328000,328120,unsafe.Pointer(&close),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setAutoReset(bool)
func (q *QProgressDialog) SetAutoReset(reset bool)  {
	q.Drv(328000,328121,unsafe.Pointer(&reset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setBar(QProgressBar*)
func (q *QProgressDialog) SetBar(bar *QProgressBar)  {
	q.Drv(328000,328122,Native(bar),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setCancelButton(QPushButton*)
func (q *QProgressDialog) SetCancelButton(button *QPushButton)  {
	q.Drv(328000,328123,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setCancelButtonText(QString const&)
func (q *QProgressDialog) SetCancelButtonText(text string)  {
	q.Drv(328000,328124,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setLabel(QLabel*)
func (q *QProgressDialog) SetLabel(label *QLabel)  {
	q.Drv(328000,328125,Native(label),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setLabelText(QString const&)
func (q *QProgressDialog) SetLabelText(text string)  {
	q.Drv(328000,328126,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setMaximum(int)
func (q *QProgressDialog) SetMaximum(maximum int)  {
	q.Drv(328000,328127,unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setMinimum(int)
func (q *QProgressDialog) SetMinimum(minimum int)  {
	q.Drv(328000,328128,unsafe.Pointer(&minimum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setMinimumDuration(int)
func (q *QProgressDialog) SetMinimumDuration(ms int)  {
	q.Drv(328000,328129,unsafe.Pointer(&ms),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setRange(int,int)
func (q *QProgressDialog) SetRange(minimum int,maximum int)  {
	q.Drv(328000,328130,unsafe.Pointer(&minimum),unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::setValue(int)
func (q *QProgressDialog) SetValue(progress int)  {
	q.Drv(328000,328131,unsafe.Pointer(&progress),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::showEvent(QShowEvent*)
func (q *QProgressDialog) ShowEvent(event *QShowEvent)  {
	q.Drv(328000,328132,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProgressDialog::sizeHint()
func (q *QProgressDialog) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(328000,328133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QProgressDialog::value()
func (q *QProgressDialog) Value() int {
	var __rv int
	q.Drv(328000,328134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProgressDialog::wasCanceled()
func (q *QProgressDialog) WasCanceled() bool {
	var __rv bool
	q.Drv(328000,328135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
