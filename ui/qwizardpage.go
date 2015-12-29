// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QWizardPage : QWizardPage
type QWizardPage struct {
	QWidget
}
func (q *QWizardPage) OnCompleteChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(398000,398102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QWizardPage::QWizardPage()	
func NewWizardPage() *QWizardPage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),398000,398103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWizardPage{}
	_p.SetDriver(__rv,398000,false)
	return _p
} 
//QWizardPage::QWizardPage(QWidget*)	
func NewWizardPageWithParent(parent QWidgetInterface) *QWizardPage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),398000,398104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWizardPage{}
	_p.SetDriver(__rv,398000,false)
	return _p
} 
//QWizardPage::buttonText(QWizard::WizardButton)
func (q *QWizardPage) ButtonText(which QWizard_WizardButton) string {
	var __rv string
	q.Drv(398000,398105,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::cleanupPage()
func (q *QWizardPage) CleanupPage()  {
	q.Drv(398000,398106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::field(QString const&)
func (q *QWizardPage) Field(name string) *QVariant {
	var __rv uintptr
	q.Drv(398000,398107,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QWizardPage::initializePage()
func (q *QWizardPage) InitializePage()  {
	q.Drv(398000,398108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::isCommitPage()
func (q *QWizardPage) IsCommitPage() bool {
	var __rv bool
	q.Drv(398000,398109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::isComplete()
func (q *QWizardPage) IsComplete() bool {
	var __rv bool
	q.Drv(398000,398110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::isFinalPage()
func (q *QWizardPage) IsFinalPage() bool {
	var __rv bool
	q.Drv(398000,398111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::nextId()
func (q *QWizardPage) NextId() int {
	var __rv int
	q.Drv(398000,398112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::pixmap(QWizard::WizardPixmap)
func (q *QWizardPage) Pixmap(which QWizard_WizardPixmap) *QPixmap {
	var __rv uintptr
	q.Drv(398000,398113,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QWizardPage::registerField(QString const&,QWidget*,char const*,char const*)
func (q *QWizardPage) RegisterField(name string,widget QWidgetInterface,property string,changedSignal string)  {
	q.Drv(398000,398114,unsafe.Pointer(&name),Native(widget),unsafe.Pointer(&property),unsafe.Pointer(&changedSignal),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setButtonText(QWizard::WizardButton,QString const&)
func (q *QWizardPage) SetButtonText(which QWizard_WizardButton,text string)  {
	q.Drv(398000,398115,unsafe.Pointer(&which),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setCommitPage(bool)
func (q *QWizardPage) SetCommitPage(commitPage bool)  {
	q.Drv(398000,398116,unsafe.Pointer(&commitPage),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setField(QString const&,QVariant const&)
func (q *QWizardPage) SetField(name string,value *QVariant)  {
	q.Drv(398000,398117,unsafe.Pointer(&name),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setFinalPage(bool)
func (q *QWizardPage) SetFinalPage(finalPage bool)  {
	q.Drv(398000,398118,unsafe.Pointer(&finalPage),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setPixmap(QWizard::WizardPixmap,QPixmap const&)
func (q *QWizardPage) SetPixmap(which QWizard_WizardPixmap,pixmap *QPixmap)  {
	q.Drv(398000,398119,unsafe.Pointer(&which),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setSubTitle(QString const&)
func (q *QWizardPage) SetSubTitle(subTitle string)  {
	q.Drv(398000,398120,unsafe.Pointer(&subTitle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::setTitle(QString const&)
func (q *QWizardPage) SetTitle(title string)  {
	q.Drv(398000,398121,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizardPage::subTitle()
func (q *QWizardPage) SubTitle() string {
	var __rv string
	q.Drv(398000,398122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::title()
func (q *QWizardPage) Title() string {
	var __rv string
	q.Drv(398000,398123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::validatePage()
func (q *QWizardPage) ValidatePage() bool {
	var __rv bool
	q.Drv(398000,398124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizardPage::wizard()
func (q *QWizardPage) Wizard() *QWizard {
	var __rv uintptr
	q.Drv(398000,398125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWizard{}
	_rp.SetDriver(__rv,397000,false)
	return _rp
}	
