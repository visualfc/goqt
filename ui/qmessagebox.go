// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMessageBox_ButtonRole - QMessageBox::ButtonRole
type QMessageBox_ButtonRole int32
const (
	QMessageBox_InvalidRole QMessageBox_ButtonRole = -1
	QMessageBox_AcceptRole QMessageBox_ButtonRole = -1+1
	QMessageBox_RejectRole QMessageBox_ButtonRole = -1+1+1
	QMessageBox_DestructiveRole QMessageBox_ButtonRole = -1+1+1+1
	QMessageBox_ActionRole QMessageBox_ButtonRole = -1+1+1+1+1
	QMessageBox_HelpRole QMessageBox_ButtonRole = -1+1+1+1+1+1
	QMessageBox_YesRole QMessageBox_ButtonRole = -1+1+1+1+1+1+1
	QMessageBox_NoRole QMessageBox_ButtonRole = -1+1+1+1+1+1+1+1
	QMessageBox_ResetRole QMessageBox_ButtonRole = -1+1+1+1+1+1+1+1+1
	QMessageBox_ApplyRole QMessageBox_ButtonRole = -1+1+1+1+1+1+1+1+1+1
	QMessageBox_NRoles QMessageBox_ButtonRole = -1+1+1+1+1+1+1+1+1+1+1
)
//enum QMessageBox_StandardButton - QMessageBox::StandardButton
type QMessageBox_StandardButton uint32
const (
	QMessageBox_NoButton QMessageBox_StandardButton = 0x00000000
	QMessageBox_Ok QMessageBox_StandardButton = 0x00000400
	QMessageBox_Save QMessageBox_StandardButton = 0x00000800
	QMessageBox_SaveAll QMessageBox_StandardButton = 0x00001000
	QMessageBox_Open QMessageBox_StandardButton = 0x00002000
	QMessageBox_Yes QMessageBox_StandardButton = 0x00004000
	QMessageBox_YesToAll QMessageBox_StandardButton = 0x00008000
	QMessageBox_No QMessageBox_StandardButton = 0x00010000
	QMessageBox_NoToAll QMessageBox_StandardButton = 0x00020000
	QMessageBox_Abort QMessageBox_StandardButton = 0x00040000
	QMessageBox_Retry QMessageBox_StandardButton = 0x00080000
	QMessageBox_Ignore QMessageBox_StandardButton = 0x00100000
	QMessageBox_Close QMessageBox_StandardButton = 0x00200000
	QMessageBox_Cancel QMessageBox_StandardButton = 0x00400000
	QMessageBox_Discard QMessageBox_StandardButton = 0x00800000
	QMessageBox_Help QMessageBox_StandardButton = 0x01000000
	QMessageBox_Apply QMessageBox_StandardButton = 0x02000000
	QMessageBox_Reset QMessageBox_StandardButton = 0x04000000
	QMessageBox_RestoreDefaults QMessageBox_StandardButton = 0x08000000
	QMessageBox_FirstButton QMessageBox_StandardButton = QMessageBox_Ok
	QMessageBox_LastButton QMessageBox_StandardButton = QMessageBox_RestoreDefaults
	QMessageBox_YesAll QMessageBox_StandardButton = QMessageBox_YesToAll
	QMessageBox_NoAll QMessageBox_StandardButton = QMessageBox_NoToAll
	QMessageBox_Default QMessageBox_StandardButton = 0x00000100
	QMessageBox_Escape QMessageBox_StandardButton = 0x00000200
	QMessageBox_FlagMask QMessageBox_StandardButton = 0x00000300
	QMessageBox_ButtonMask QMessageBox_StandardButton = ^QMessageBox_FlagMask
)
//enum QMessageBox_Icon - QMessageBox::Icon
type QMessageBox_Icon uint32
const (
	QMessageBox_NoIcon QMessageBox_Icon = 0
	QMessageBox_Information QMessageBox_Icon = 1
	QMessageBox_Warning QMessageBox_Icon = 2
	QMessageBox_Critical QMessageBox_Icon = 3
	QMessageBox_Question QMessageBox_Icon = 4
)
//enum QMessageBox_Button - QMessageBox::Button
type QMessageBox_Button uint32
const (
	QMessageBox_Button_NoButton QMessageBox_Button = 0x00000000
	QMessageBox_Button_Ok QMessageBox_Button = 0x00000400
	QMessageBox_Button_Save QMessageBox_Button = 0x00000800
	QMessageBox_Button_SaveAll QMessageBox_Button = 0x00001000
	QMessageBox_Button_Open QMessageBox_Button = 0x00002000
	QMessageBox_Button_Yes QMessageBox_Button = 0x00004000
	QMessageBox_Button_YesToAll QMessageBox_Button = 0x00008000
	QMessageBox_Button_No QMessageBox_Button = 0x00010000
	QMessageBox_Button_NoToAll QMessageBox_Button = 0x00020000
	QMessageBox_Button_Abort QMessageBox_Button = 0x00040000
	QMessageBox_Button_Retry QMessageBox_Button = 0x00080000
	QMessageBox_Button_Ignore QMessageBox_Button = 0x00100000
	QMessageBox_Button_Close QMessageBox_Button = 0x00200000
	QMessageBox_Button_Cancel QMessageBox_Button = 0x00400000
	QMessageBox_Button_Discard QMessageBox_Button = 0x00800000
	QMessageBox_Button_Help QMessageBox_Button = 0x01000000
	QMessageBox_Button_Apply QMessageBox_Button = 0x02000000
	QMessageBox_Button_Reset QMessageBox_Button = 0x04000000
	QMessageBox_Button_RestoreDefaults QMessageBox_Button = 0x08000000
	QMessageBox_Button_FirstButton QMessageBox_Button = QMessageBox_Button(QMessageBox_Ok)
	QMessageBox_Button_LastButton QMessageBox_Button = QMessageBox_Button(QMessageBox_RestoreDefaults)
	QMessageBox_Button_YesAll QMessageBox_Button = QMessageBox_Button(QMessageBox_YesToAll)
	QMessageBox_Button_NoAll QMessageBox_Button = QMessageBox_Button(QMessageBox_NoToAll)
	QMessageBox_Button_Default QMessageBox_Button = 0x00000100
	QMessageBox_Button_Escape QMessageBox_Button = 0x00000200
	QMessageBox_Button_FlagMask QMessageBox_Button = 0x00000300
	QMessageBox_Button_ButtonMask QMessageBox_Button = ^QMessageBox_Button(QMessageBox_FlagMask)
)
//struct QMessageBox : QMessageBox
type QMessageBox struct {
	QDialog
}
func (q *QMessageBox) OnButtonClicked(fn func(*QAbstractButton)) uintptr {
	var __rv uintptr
	q.Drv(310000,310102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMessageBox::QMessageBox()	
func NewMessageBox() *QMessageBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),310000,310103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMessageBox{}
	_p.SetDriver(__rv,310000,false)
	return _p
} 
//QMessageBox::QMessageBox(QWidget*)	
func NewMessageBoxWithParent(parent QWidgetInterface) *QMessageBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),310000,310104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMessageBox{}
	_p.SetDriver(__rv,310000,false)
	return _p
} 
//QMessageBox::QMessageBox(QMessageBox::Icon,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QWidget*,QFlags<Qt::WindowType>)	
func NewMessageBoxWithIconTitleTextButtonsParentFlags(icon QMessageBox_Icon,title string,text string,buttons QMessageBox_StandardButton,parent QWidgetInterface,flags Qt_WindowType) *QMessageBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),310000,310105,unsafe.Pointer(&icon),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMessageBox{}
	_p.SetDriver(__rv,310000,false)
	return _p
} 
//QMessageBox::QMessageBox(QString const&,QString const&,QMessageBox::Icon,int,int,int,QWidget*,QFlags<Qt::WindowType>)	
func NewMessageBoxWithTitleTextIconButton0Button1Button2ParentFlags(title string,text string,icon QMessageBox_Icon,button0 int,button1 int,button2 int,parent QWidgetInterface,f Qt_WindowType) *QMessageBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),310000,310106,unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&icon),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMessageBox{}
	_p.SetDriver(__rv,310000,false)
	return _p
} 
//QMessageBox::about(QWidget*,QString const&,QString const&)	
func QMessageBoxAbout(parent QWidgetInterface,title string,text string)  {
	DirectQtDrv(nil,310000,310107,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::about(QWidget*,QString const&,QString const&)
func (q *QMessageBox) About(parent QWidgetInterface,title string,text string)  {
	q.Drv(310000,310107,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::aboutQt(QWidget*)	
func QMessageBoxAboutQt(parent QWidgetInterface)  {
	DirectQtDrv(nil,310000,310108,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::aboutQt(QWidget*)
func (q *QMessageBox) AboutQt(parent QWidgetInterface)  {
	q.Drv(310000,310108,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::aboutQt(QWidget*,QString const&)	
func QMessageBoxAboutQtWithParentTitle(parent QWidgetInterface,title string)  {
	DirectQtDrv(nil,310000,310109,Native(parent),unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::aboutQt(QWidget*,QString const&)
func (q *QMessageBox) AboutQtWithParentTitle(parent QWidgetInterface,title string)  {
	q.Drv(310000,310109,Native(parent),unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::addButton(QMessageBox::StandardButton)
func (q *QMessageBox) AddButton(button QMessageBox_StandardButton) *QPushButton {
	var __rv uintptr
	q.Drv(310000,310110,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QMessageBox::addButton(QAbstractButton*,QMessageBox::ButtonRole)
func (q *QMessageBox) AddButtonWithButtonRole(button *QAbstractButton,role QMessageBox_ButtonRole)  {
	q.Drv(310000,310111,Native(button),unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::addButton(QString const&,QMessageBox::ButtonRole)
func (q *QMessageBox) AddButtonWithTextRole(text string,role QMessageBox_ButtonRole) *QPushButton {
	var __rv uintptr
	q.Drv(310000,310112,unsafe.Pointer(&text),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QMessageBox::button(QMessageBox::StandardButton)
func (q *QMessageBox) Button(which QMessageBox_StandardButton) *QAbstractButton {
	var __rv uintptr
	q.Drv(310000,310113,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QMessageBox::buttonRole(QAbstractButton*)
func (q *QMessageBox) ButtonRole(button *QAbstractButton) QMessageBox_ButtonRole {
	var __rv QMessageBox_ButtonRole
	q.Drv(310000,310114,Native(button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::buttonText(int)
func (q *QMessageBox) ButtonText(button int) string {
	var __rv string
	q.Drv(310000,310115,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::buttons()
func (q *QMessageBox) Buttons() []*QAbstractButton {
	var __rv []*QAbstractButton
	q.Drv(310000,310116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::changeEvent(QEvent*)
func (q *QMessageBox) ChangeEvent(event *QEvent)  {
	q.Drv(310000,310117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::clickedButton()
func (q *QMessageBox) ClickedButton() *QAbstractButton {
	var __rv uintptr
	q.Drv(310000,310118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QMessageBox::closeEvent(QCloseEvent*)
func (q *QMessageBox) CloseEvent(event *QCloseEvent)  {
	q.Drv(310000,310119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)	
func QMessageBoxCriticalWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	DirectQtDrv(nil,310000,310120,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)
func (q *QMessageBox) CriticalWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310120,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,int,int,int)	
func QMessageBoxCriticalWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	DirectQtDrv(nil,310000,310121,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,int,int,int)
func (q *QMessageBox) CriticalWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	q.Drv(310000,310121,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)	
func QMessageBoxCriticalWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	DirectQtDrv(nil,310000,310122,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::critical(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)
func (q *QMessageBox) CriticalWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	q.Drv(310000,310122,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::defaultButton()
func (q *QMessageBox) DefaultButton() *QPushButton {
	var __rv uintptr
	q.Drv(310000,310123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QMessageBox::detailedText()
func (q *QMessageBox) DetailedText() string {
	var __rv string
	q.Drv(310000,310124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::escapeButton()
func (q *QMessageBox) EscapeButton() *QAbstractButton {
	var __rv uintptr
	q.Drv(310000,310125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractButton{}
	_rp.SetDriver(__rv,193000,false)
	return _rp
}	
//QMessageBox::event(QEvent*)
func (q *QMessageBox) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(310000,310126,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::icon()
func (q *QMessageBox) Icon() QMessageBox_Icon {
	var __rv QMessageBox_Icon
	q.Drv(310000,310127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::iconPixmap()
func (q *QMessageBox) IconPixmap() *QPixmap {
	var __rv uintptr
	q.Drv(310000,310128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)	
func QMessageBoxInformationWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	DirectQtDrv(nil,310000,310129,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)
func (q *QMessageBox) InformationWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310129,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,int,int,int)	
func QMessageBoxInformationWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	DirectQtDrv(nil,310000,310130,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,int,int,int)
func (q *QMessageBox) InformationWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	q.Drv(310000,310130,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)	
func QMessageBoxInformationWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	DirectQtDrv(nil,310000,310131,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::information(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)
func (q *QMessageBox) InformationWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	q.Drv(310000,310131,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::informativeText()
func (q *QMessageBox) InformativeText() string {
	var __rv string
	q.Drv(310000,310132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::keyPressEvent(QKeyEvent*)
func (q *QMessageBox) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(310000,310133,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::open()
func (q *QMessageBox) Open()  {
	q.Drv(310000,310134,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::open(QObject*,char const*)
func (q *QMessageBox) OpenWithObjectMember(receiver QObjectInterface,member string)  {
	q.Drv(310000,310135,Native(receiver),unsafe.Pointer(&member),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)	
func QMessageBoxQuestionWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	DirectQtDrv(nil,310000,310136,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)
func (q *QMessageBox) QuestionWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310136,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,int,int,int)	
func QMessageBoxQuestionWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	DirectQtDrv(nil,310000,310137,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,int,int,int)
func (q *QMessageBox) QuestionWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	q.Drv(310000,310137,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)	
func QMessageBoxQuestionWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	DirectQtDrv(nil,310000,310138,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::question(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)
func (q *QMessageBox) QuestionWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	q.Drv(310000,310138,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::removeButton(QAbstractButton*)
func (q *QMessageBox) RemoveButton(button *QAbstractButton)  {
	q.Drv(310000,310139,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::resizeEvent(QResizeEvent*)
func (q *QMessageBox) ResizeEvent(event *QResizeEvent)  {
	q.Drv(310000,310140,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setButtonText(int,QString const&)
func (q *QMessageBox) SetButtonText(button int,text string)  {
	q.Drv(310000,310141,unsafe.Pointer(&button),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setDefaultButton(QMessageBox::StandardButton)
func (q *QMessageBox) SetDefaultButton(button QMessageBox_StandardButton)  {
	q.Drv(310000,310142,unsafe.Pointer(&button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setDefaultButton(QPushButton*)
func (q *QMessageBox) SetDefaultButtonWithButton(button *QPushButton)  {
	q.Drv(310000,310143,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setDetailedText(QString const&)
func (q *QMessageBox) SetDetailedText(text string)  {
	q.Drv(310000,310144,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setEscapeButton(QAbstractButton*)
func (q *QMessageBox) SetEscapeButton(button *QAbstractButton)  {
	q.Drv(310000,310145,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setEscapeButton(QMessageBox::StandardButton)
func (q *QMessageBox) SetEscapeButtonWithStandardbutton(button QMessageBox_StandardButton)  {
	q.Drv(310000,310146,unsafe.Pointer(&button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setIcon(QMessageBox::Icon)
func (q *QMessageBox) SetIcon(value QMessageBox_Icon)  {
	q.Drv(310000,310147,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setIconPixmap(QPixmap const&)
func (q *QMessageBox) SetIconPixmap(pixmap *QPixmap)  {
	q.Drv(310000,310148,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setInformativeText(QString const&)
func (q *QMessageBox) SetInformativeText(text string)  {
	q.Drv(310000,310149,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setStandardButtons(QFlags<QMessageBox::StandardButton>)
func (q *QMessageBox) SetStandardButtons(buttons QMessageBox_StandardButton)  {
	q.Drv(310000,310150,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setText(QString const&)
func (q *QMessageBox) SetText(text string)  {
	q.Drv(310000,310151,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setTextFormat(Qt::TextFormat)
func (q *QMessageBox) SetTextFormat(format Qt_TextFormat)  {
	q.Drv(310000,310152,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setWindowModality(Qt::WindowModality)
func (q *QMessageBox) SetWindowModality(windowModality Qt_WindowModality)  {
	q.Drv(310000,310153,unsafe.Pointer(&windowModality),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::setWindowTitle(QString const&)
func (q *QMessageBox) SetWindowTitle(title string)  {
	q.Drv(310000,310154,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::showEvent(QShowEvent*)
func (q *QMessageBox) ShowEvent(event *QShowEvent)  {
	q.Drv(310000,310155,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMessageBox::sizeHint()
func (q *QMessageBox) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(310000,310156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMessageBox::standardButton(QAbstractButton*)
func (q *QMessageBox) StandardButton(button *QAbstractButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310157,Native(button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::standardButtons()
func (q *QMessageBox) StandardButtons() QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::standardIcon(QMessageBox::Icon)	
func QMessageBoxStandardIcon(icon QMessageBox_Icon) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,310000,310159,unsafe.Pointer(&icon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QMessageBox::standardIcon(QMessageBox::Icon)
func (q *QMessageBox) StandardIcon(icon QMessageBox_Icon) *QPixmap {
	var __rv uintptr
	q.Drv(310000,310159,unsafe.Pointer(&icon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QMessageBox::text()
func (q *QMessageBox) Text() string {
	var __rv string
	q.Drv(310000,310160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::textFormat()
func (q *QMessageBox) TextFormat() Qt_TextFormat {
	var __rv Qt_TextFormat
	q.Drv(310000,310161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)	
func QMessageBoxWarningWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	DirectQtDrv(nil,310000,310162,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,QFlags<QMessageBox::StandardButton>,QMessageBox::StandardButton)
func (q *QMessageBox) WarningWithParentTitleTextButtonsStandardbutton(parent QWidgetInterface,title string,text string,buttons QMessageBox_StandardButton,defaultButton QMessageBox_StandardButton) QMessageBox_StandardButton {
	var __rv QMessageBox_StandardButton
	q.Drv(310000,310162,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&buttons),unsafe.Pointer(&defaultButton),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,int,int,int)	
func QMessageBoxWarningWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	DirectQtDrv(nil,310000,310163,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,int,int,int)
func (q *QMessageBox) WarningWithParentTitleTextButton0Button1Button2(parent QWidgetInterface,title string,text string,button0 int,button1 int,button2 int) int {
	var __rv int
	q.Drv(310000,310163,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0),unsafe.Pointer(&button1),unsafe.Pointer(&button2),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)	
func QMessageBoxWarningWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	DirectQtDrv(nil,310000,310164,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
//QMessageBox::warning(QWidget*,QString const&,QString const&,QString const&,QString const&,QString const&,int,int)
func (q *QMessageBox) WarningWithParentTitleTextButton0textButton1textButton2textDefaultbuttonnumberEscapebuttonnumber(parent QWidgetInterface,title string,text string,button0Text string,button1Text string,button2Text string,defaultButtonNumber int,escapeButtonNumber int) int {
	var __rv int
	q.Drv(310000,310164,Native(parent),unsafe.Pointer(&title),unsafe.Pointer(&text),unsafe.Pointer(&button0Text),unsafe.Pointer(&button1Text),unsafe.Pointer(&button2Text),unsafe.Pointer(&defaultButtonNumber),unsafe.Pointer(&escapeButtonNumber),unsafe.Pointer(&__rv),nil,nil,nil)
	return __rv
}	
