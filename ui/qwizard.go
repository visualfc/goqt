// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QWizard_WizardButton - QWizard::WizardButton
type QWizard_WizardButton int32
const (
	QWizard_BackButton QWizard_WizardButton = 0
	QWizard_NextButton QWizard_WizardButton = 1
	QWizard_CommitButton QWizard_WizardButton = 2
	QWizard_FinishButton QWizard_WizardButton = 3
	QWizard_CancelButton QWizard_WizardButton = 4
	QWizard_HelpButton QWizard_WizardButton = 5
	QWizard_CustomButton1 QWizard_WizardButton = 6
	QWizard_CustomButton2 QWizard_WizardButton = 7
	QWizard_CustomButton3 QWizard_WizardButton = 8
	QWizard_Stretch QWizard_WizardButton = 9
	QWizard_NoButton QWizard_WizardButton = -1
	QWizard_NStandardButtons QWizard_WizardButton = 6
	QWizard_NButtons QWizard_WizardButton = 9
)
//enum QWizard_WizardStyle - QWizard::WizardStyle
type QWizard_WizardStyle uint32
const (
	QWizard_ClassicStyle QWizard_WizardStyle = 0
	QWizard_ModernStyle QWizard_WizardStyle = 1
	QWizard_MacStyle QWizard_WizardStyle = 2
	QWizard_AeroStyle QWizard_WizardStyle = 3
	QWizard_NStyles QWizard_WizardStyle = 4
)
//enum QWizard_WizardPixmap - QWizard::WizardPixmap
type QWizard_WizardPixmap uint32
const (
	QWizard_WatermarkPixmap QWizard_WizardPixmap = 0
	QWizard_LogoPixmap QWizard_WizardPixmap = 1
	QWizard_BannerPixmap QWizard_WizardPixmap = 2
	QWizard_BackgroundPixmap QWizard_WizardPixmap = 3
	QWizard_NPixmaps QWizard_WizardPixmap = 4
)
//enum QWizard_WizardOption - QWizard::WizardOption
type QWizard_WizardOption uint32
const (
	QWizard_IndependentPages QWizard_WizardOption = 0x00000001
	QWizard_IgnoreSubTitles QWizard_WizardOption = 0x00000002
	QWizard_ExtendedWatermarkPixmap QWizard_WizardOption = 0x00000004
	QWizard_NoDefaultButton QWizard_WizardOption = 0x00000008
	QWizard_NoBackButtonOnStartPage QWizard_WizardOption = 0x00000010
	QWizard_NoBackButtonOnLastPage QWizard_WizardOption = 0x00000020
	QWizard_DisabledBackButtonOnLastPage QWizard_WizardOption = 0x00000040
	QWizard_HaveNextButtonOnLastPage QWizard_WizardOption = 0x00000080
	QWizard_HaveFinishButtonOnEarlyPages QWizard_WizardOption = 0x00000100
	QWizard_NoCancelButton QWizard_WizardOption = 0x00000200
	QWizard_CancelButtonOnLeft QWizard_WizardOption = 0x00000400
	QWizard_HaveHelpButton QWizard_WizardOption = 0x00000800
	QWizard_HelpButtonOnRight QWizard_WizardOption = 0x00001000
	QWizard_HaveCustomButton1 QWizard_WizardOption = 0x00002000
	QWizard_HaveCustomButton2 QWizard_WizardOption = 0x00004000
	QWizard_HaveCustomButton3 QWizard_WizardOption = 0x00008000
)
//struct QWizard : QWizard
type QWizard struct {
	QDialog
}
func (q *QWizard) OnHelpRequested(fn func()) uintptr {
	var __rv uintptr
	q.Drv(397000,397102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QWizard) OnCurrentIdChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(397000,397103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QWizard) OnCustomButtonClicked(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(397000,397104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QWizard) OnPageAdded(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(397000,397105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QWizard) OnPageRemoved(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(397000,397106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QWizard::QWizard()	
func NewWizard() *QWizard {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),397000,397107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWizard{}
	_p.SetDriver(__rv,397000,false)
	return _p
} 
//QWizard::QWizard(QWidget*,QFlags<Qt::WindowType>)	
func NewWizardWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QWizard {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),397000,397108,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QWizard{}
	_p.SetDriver(__rv,397000,false)
	return _p
} 
//QWizard::addPage(QWizardPage*)
func (q *QWizard) AddPage(page *QWizardPage) int {
	var __rv int
	q.Drv(397000,397109,Native(page),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::back()
func (q *QWizard) Back()  {
	q.Drv(397000,397110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::button(QWizard::WizardButton)
func (q *QWizard) Button(which QWizard_WizardButton) *QAbstractButton {
	var __rv uintptr
	q.Drv(397000,397111,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QWizard::buttonText(QWizard::WizardButton)
func (q *QWizard) ButtonText(which QWizard_WizardButton) string {
	var __rv string
	q.Drv(397000,397112,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::cleanupPage(int)
func (q *QWizard) CleanupPage(id int)  {
	q.Drv(397000,397113,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::currentId()
func (q *QWizard) CurrentId() int {
	var __rv int
	q.Drv(397000,397114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::currentPage()
func (q *QWizard) CurrentPage() *QWizardPage {
	var __rv uintptr
	q.Drv(397000,397115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWizardPage{}
	_rp.SetDriver(__rv,398000,false)
	return _rp
}	
//QWizard::done(int)
func (q *QWizard) Done(result int)  {
	q.Drv(397000,397116,unsafe.Pointer(&result),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::event(QEvent*)
func (q *QWizard) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(397000,397117,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::field(QString const&)
func (q *QWizard) Field(name string) *QVariant {
	var __rv uintptr
	q.Drv(397000,397118,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QWizard::hasVisitedPage(int)
func (q *QWizard) HasVisitedPage(id int) bool {
	var __rv bool
	q.Drv(397000,397119,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::initializePage(int)
func (q *QWizard) InitializePage(id int)  {
	q.Drv(397000,397120,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::next()
func (q *QWizard) Next()  {
	q.Drv(397000,397121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::nextId()
func (q *QWizard) NextId() int {
	var __rv int
	q.Drv(397000,397122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::options()
func (q *QWizard) Options() QWizard_WizardOption {
	var __rv QWizard_WizardOption
	q.Drv(397000,397123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::page(int)
func (q *QWizard) Page(id int) *QWizardPage {
	var __rv uintptr
	q.Drv(397000,397124,unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWizardPage{}
	_rp.SetDriver(__rv,398000,false)
	return _rp
}	
//QWizard::pageIds()
func (q *QWizard) PageIds() []int {
	var __rv []int
	q.Drv(397000,397125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::paintEvent(QPaintEvent*)
func (q *QWizard) PaintEvent(event *QPaintEvent)  {
	q.Drv(397000,397126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::pixmap(QWizard::WizardPixmap)
func (q *QWizard) Pixmap(which QWizard_WizardPixmap) *QPixmap {
	var __rv uintptr
	q.Drv(397000,397127,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QWizard::removePage(int)
func (q *QWizard) RemovePage(id int)  {
	q.Drv(397000,397128,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::resizeEvent(QResizeEvent*)
func (q *QWizard) ResizeEvent(event *QResizeEvent)  {
	q.Drv(397000,397129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::restart()
func (q *QWizard) Restart()  {
	q.Drv(397000,397130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setButton(QWizard::WizardButton,QAbstractButton*)
func (q *QWizard) SetButton(which QWizard_WizardButton,button *QAbstractButton)  {
	q.Drv(397000,397131,unsafe.Pointer(&which),Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setButtonLayout(QList<QWizard::WizardButton> const&)
func (q *QWizard) SetButtonLayout(layout []QWizard_WizardButton)  {
	q.Drv(397000,397132,unsafe.Pointer(&layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setButtonText(QWizard::WizardButton,QString const&)
func (q *QWizard) SetButtonText(which QWizard_WizardButton,text string)  {
	q.Drv(397000,397133,unsafe.Pointer(&which),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setDefaultProperty(char const*,char const*,char const*)
func (q *QWizard) SetDefaultProperty(className string,property string,changedSignal string)  {
	q.Drv(397000,397134,unsafe.Pointer(&className),unsafe.Pointer(&property),unsafe.Pointer(&changedSignal),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setField(QString const&,QVariant const&)
func (q *QWizard) SetField(name string,value *QVariant)  {
	q.Drv(397000,397135,unsafe.Pointer(&name),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setOption(QWizard::WizardOption)
func (q *QWizard) SetOption(option QWizard_WizardOption)  {
	q.Drv(397000,397136,unsafe.Pointer(&option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setOption(QWizard::WizardOption,bool)
func (q *QWizard) SetOptionWithOptionOn(option QWizard_WizardOption,on bool)  {
	q.Drv(397000,397137,unsafe.Pointer(&option),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setOptions(QFlags<QWizard::WizardOption>)
func (q *QWizard) SetOptions(options QWizard_WizardOption)  {
	q.Drv(397000,397138,unsafe.Pointer(&options),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setPage(int,QWizardPage*)
func (q *QWizard) SetPage(id int,page *QWizardPage)  {
	q.Drv(397000,397139,unsafe.Pointer(&id),Native(page),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setPixmap(QWizard::WizardPixmap,QPixmap const&)
func (q *QWizard) SetPixmap(which QWizard_WizardPixmap,pixmap *QPixmap)  {
	q.Drv(397000,397140,unsafe.Pointer(&which),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setSideWidget(QWidget*)
func (q *QWizard) SetSideWidget(widget QWidgetInterface)  {
	q.Drv(397000,397141,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setStartId(int)
func (q *QWizard) SetStartId(id int)  {
	q.Drv(397000,397142,unsafe.Pointer(&id),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setSubTitleFormat(Qt::TextFormat)
func (q *QWizard) SetSubTitleFormat(format Qt_TextFormat)  {
	q.Drv(397000,397143,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setTitleFormat(Qt::TextFormat)
func (q *QWizard) SetTitleFormat(format Qt_TextFormat)  {
	q.Drv(397000,397144,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setVisible(bool)
func (q *QWizard) SetVisible(visible bool)  {
	q.Drv(397000,397145,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::setWizardStyle(QWizard::WizardStyle)
func (q *QWizard) SetWizardStyle(style QWizard_WizardStyle)  {
	q.Drv(397000,397146,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QWizard::sideWidget()
func (q *QWizard) SideWidget() *QWidget {
	var __rv uintptr
	q.Drv(397000,397147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QWizard::sizeHint()
func (q *QWizard) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(397000,397148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QWizard::startId()
func (q *QWizard) StartId() int {
	var __rv int
	q.Drv(397000,397149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::subTitleFormat()
func (q *QWizard) SubTitleFormat() Qt_TextFormat {
	var __rv Qt_TextFormat
	q.Drv(397000,397150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::testOption(QWizard::WizardOption)
func (q *QWizard) TestOption(option QWizard_WizardOption) bool {
	var __rv bool
	q.Drv(397000,397151,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::titleFormat()
func (q *QWizard) TitleFormat() Qt_TextFormat {
	var __rv Qt_TextFormat
	q.Drv(397000,397152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::validateCurrentPage()
func (q *QWizard) ValidateCurrentPage() bool {
	var __rv bool
	q.Drv(397000,397153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::visitedPages()
func (q *QWizard) VisitedPages() []int {
	var __rv []int
	q.Drv(397000,397154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QWizard::wizardStyle()
func (q *QWizard) WizardStyle() QWizard_WizardStyle {
	var __rv QWizard_WizardStyle
	q.Drv(397000,397155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
