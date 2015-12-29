// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDialogButtonBox_ButtonLayout - QDialogButtonBox::ButtonLayout
type QDialogButtonBox_ButtonLayout uint32
const (
	QDialogButtonBox_WinLayout QDialogButtonBox_ButtonLayout = 0
	QDialogButtonBox_MacLayout QDialogButtonBox_ButtonLayout = 1
	QDialogButtonBox_KdeLayout QDialogButtonBox_ButtonLayout = 2
	QDialogButtonBox_GnomeLayout QDialogButtonBox_ButtonLayout = 3
)
//enum QDialogButtonBox_ButtonRole - QDialogButtonBox::ButtonRole
type QDialogButtonBox_ButtonRole int32
const (
	QDialogButtonBox_InvalidRole QDialogButtonBox_ButtonRole = -1
	QDialogButtonBox_AcceptRole QDialogButtonBox_ButtonRole = -1+1
	QDialogButtonBox_RejectRole QDialogButtonBox_ButtonRole = -1+1+1
	QDialogButtonBox_DestructiveRole QDialogButtonBox_ButtonRole = -1+1+1+1
	QDialogButtonBox_ActionRole QDialogButtonBox_ButtonRole = -1+1+1+1+1
	QDialogButtonBox_HelpRole QDialogButtonBox_ButtonRole = -1+1+1+1+1+1
	QDialogButtonBox_YesRole QDialogButtonBox_ButtonRole = -1+1+1+1+1+1+1
	QDialogButtonBox_NoRole QDialogButtonBox_ButtonRole = -1+1+1+1+1+1+1+1
	QDialogButtonBox_ResetRole QDialogButtonBox_ButtonRole = -1+1+1+1+1+1+1+1+1
	QDialogButtonBox_ApplyRole QDialogButtonBox_ButtonRole = -1+1+1+1+1+1+1+1+1+1
	QDialogButtonBox_NRoles QDialogButtonBox_ButtonRole = -1+1+1+1+1+1+1+1+1+1+1
)
//enum QDialogButtonBox_StandardButton - QDialogButtonBox::StandardButton
type QDialogButtonBox_StandardButton uint32
const (
	QDialogButtonBox_NoButton QDialogButtonBox_StandardButton = 0x00000000
	QDialogButtonBox_Ok QDialogButtonBox_StandardButton = 0x00000400
	QDialogButtonBox_Save QDialogButtonBox_StandardButton = 0x00000800
	QDialogButtonBox_SaveAll QDialogButtonBox_StandardButton = 0x00001000
	QDialogButtonBox_Open QDialogButtonBox_StandardButton = 0x00002000
	QDialogButtonBox_Yes QDialogButtonBox_StandardButton = 0x00004000
	QDialogButtonBox_YesToAll QDialogButtonBox_StandardButton = 0x00008000
	QDialogButtonBox_No QDialogButtonBox_StandardButton = 0x00010000
	QDialogButtonBox_NoToAll QDialogButtonBox_StandardButton = 0x00020000
	QDialogButtonBox_Abort QDialogButtonBox_StandardButton = 0x00040000
	QDialogButtonBox_Retry QDialogButtonBox_StandardButton = 0x00080000
	QDialogButtonBox_Ignore QDialogButtonBox_StandardButton = 0x00100000
	QDialogButtonBox_Close QDialogButtonBox_StandardButton = 0x00200000
	QDialogButtonBox_Cancel QDialogButtonBox_StandardButton = 0x00400000
	QDialogButtonBox_Discard QDialogButtonBox_StandardButton = 0x00800000
	QDialogButtonBox_Help QDialogButtonBox_StandardButton = 0x01000000
	QDialogButtonBox_Apply QDialogButtonBox_StandardButton = 0x02000000
	QDialogButtonBox_Reset QDialogButtonBox_StandardButton = 0x04000000
	QDialogButtonBox_RestoreDefaults QDialogButtonBox_StandardButton = 0x08000000
	QDialogButtonBox_FirstButton QDialogButtonBox_StandardButton = QDialogButtonBox_Ok
	QDialogButtonBox_LastButton QDialogButtonBox_StandardButton = QDialogButtonBox_RestoreDefaults
)
//struct QDialogButtonBox : QDialogButtonBox
type QDialogButtonBox struct {
	QWidget
}
func (q *QDialogButtonBox) OnAccepted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(227000,227102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDialogButtonBox) OnClicked(fn func(*QAbstractButton)) uintptr {
	var __rv uintptr
	q.Drv(227000,227103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDialogButtonBox) OnHelpRequested(fn func()) uintptr {
	var __rv uintptr
	q.Drv(227000,227104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDialogButtonBox) OnRejected(fn func()) uintptr {
	var __rv uintptr
	q.Drv(227000,227105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDialogButtonBox::QDialogButtonBox()	
func NewDialogButtonBox() *QDialogButtonBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),227000,227106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialogButtonBox{}
	_p.SetDriver(__rv,227000,false)
	return _p
} 
//QDialogButtonBox::QDialogButtonBox(QWidget*)	
func NewDialogButtonBoxWithParent(parent QWidgetInterface) *QDialogButtonBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),227000,227107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialogButtonBox{}
	_p.SetDriver(__rv,227000,false)
	return _p
} 
//QDialogButtonBox::QDialogButtonBox(Qt::Orientation,QWidget*)	
func NewDialogButtonBoxWithOrientationParent(orientation Qt_Orientation,parent QWidgetInterface) *QDialogButtonBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),227000,227108,unsafe.Pointer(&orientation),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialogButtonBox{}
	_p.SetDriver(__rv,227000,false)
	return _p
} 
//QDialogButtonBox::QDialogButtonBox(QFlags<QDialogButtonBox::StandardButton>,Qt::Orientation,QWidget*)	
func NewDialogButtonBoxWithButtonsOrientationParent(buttons QDialogButtonBox_StandardButton,orientation Qt_Orientation,parent QWidgetInterface) *QDialogButtonBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),227000,227109,unsafe.Pointer(&buttons),unsafe.Pointer(&orientation),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDialogButtonBox{}
	_p.SetDriver(__rv,227000,false)
	return _p
} 
//QDialogButtonBox::addButton(QDialogButtonBox::StandardButton)
func (q *QDialogButtonBox) AddButton(button QDialogButtonBox_StandardButton) *QPushButton {
	var __rv uintptr
	q.Drv(227000,227110,unsafe.Pointer(&button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QDialogButtonBox::addButton(QAbstractButton*,QDialogButtonBox::ButtonRole)
func (q *QDialogButtonBox) AddButtonWithButtonRole(button *QAbstractButton,role QDialogButtonBox_ButtonRole)  {
	q.Drv(227000,227111,Native(button),unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::addButton(QString const&,QDialogButtonBox::ButtonRole)
func (q *QDialogButtonBox) AddButtonWithTextRole(text string,role QDialogButtonBox_ButtonRole) *QPushButton {
	var __rv uintptr
	q.Drv(227000,227112,unsafe.Pointer(&text),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QDialogButtonBox::button(QDialogButtonBox::StandardButton)
func (q *QDialogButtonBox) Button(which QDialogButtonBox_StandardButton) *QPushButton {
	var __rv uintptr
	q.Drv(227000,227113,unsafe.Pointer(&which),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPushButton{}
	_rp.SetDriver(__rv,331000,false)
	return _rp
}	
//QDialogButtonBox::buttonRole(QAbstractButton*)
func (q *QDialogButtonBox) ButtonRole(button *QAbstractButton) QDialogButtonBox_ButtonRole {
	var __rv QDialogButtonBox_ButtonRole
	q.Drv(227000,227114,Native(button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::buttons()
func (q *QDialogButtonBox) Buttons() []*QAbstractButton {
	var __rv []*QAbstractButton
	q.Drv(227000,227115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::centerButtons()
func (q *QDialogButtonBox) CenterButtons() bool {
	var __rv bool
	q.Drv(227000,227116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::changeEvent(QEvent*)
func (q *QDialogButtonBox) ChangeEvent(event *QEvent)  {
	q.Drv(227000,227117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::clear()
func (q *QDialogButtonBox) Clear()  {
	q.Drv(227000,227118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::event(QEvent*)
func (q *QDialogButtonBox) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(227000,227119,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::orientation()
func (q *QDialogButtonBox) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(227000,227120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::removeButton(QAbstractButton*)
func (q *QDialogButtonBox) RemoveButton(button *QAbstractButton)  {
	q.Drv(227000,227121,Native(button),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::setCenterButtons(bool)
func (q *QDialogButtonBox) SetCenterButtons(center bool)  {
	q.Drv(227000,227122,unsafe.Pointer(&center),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::setOrientation(Qt::Orientation)
func (q *QDialogButtonBox) SetOrientation(orientation Qt_Orientation)  {
	q.Drv(227000,227123,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::setStandardButtons(QFlags<QDialogButtonBox::StandardButton>)
func (q *QDialogButtonBox) SetStandardButtons(buttons QDialogButtonBox_StandardButton)  {
	q.Drv(227000,227124,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDialogButtonBox::standardButton(QAbstractButton*)
func (q *QDialogButtonBox) StandardButton(button *QAbstractButton) QDialogButtonBox_StandardButton {
	var __rv QDialogButtonBox_StandardButton
	q.Drv(227000,227125,Native(button),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDialogButtonBox::standardButtons()
func (q *QDialogButtonBox) StandardButtons() QDialogButtonBox_StandardButton {
	var __rv QDialogButtonBox_StandardButton
	q.Drv(227000,227126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
