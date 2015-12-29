// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAction_Priority - QAction::Priority
type QAction_Priority uint32
const (
	QAction_LowPriority QAction_Priority = 0
	QAction_NormalPriority QAction_Priority = 128
	QAction_HighPriority QAction_Priority = 256
)
//enum QAction_ActionEvent - QAction::ActionEvent
type QAction_ActionEvent uint32
const (
	QAction_Trigger QAction_ActionEvent = 0
	QAction_Hover QAction_ActionEvent = 1
)
//enum QAction_SoftKeyRole - QAction::SoftKeyRole
type QAction_SoftKeyRole uint32
const (
	QAction_NoSoftKey QAction_SoftKeyRole = 0
	QAction_PositiveSoftKey QAction_SoftKeyRole = 1
	QAction_NegativeSoftKey QAction_SoftKeyRole = 2
	QAction_SelectSoftKey QAction_SoftKeyRole = 3
)
//enum QAction_MenuRole - QAction::MenuRole
type QAction_MenuRole uint32
const (
	QAction_NoRole QAction_MenuRole = 0
	QAction_TextHeuristicRole QAction_MenuRole = 1
	QAction_ApplicationSpecificRole QAction_MenuRole = 2
	QAction_AboutQtRole QAction_MenuRole = 3
	QAction_AboutRole QAction_MenuRole = 4
	QAction_PreferencesRole QAction_MenuRole = 5
	QAction_QuitRole QAction_MenuRole = 6
)
//struct QAction : QAction
type QAction struct {
	QObject
}
func (q *QAction) OnChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(207000,207102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAction) OnTriggered(fn func()) uintptr {
	var __rv uintptr
	q.Drv(207000,207103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAction) OnTriggeredEx(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(207000,207104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAction) OnHovered(fn func()) uintptr {
	var __rv uintptr
	q.Drv(207000,207105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAction) OnToggled(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(207000,207106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAction::QAction(QObject*)	
func NewAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),207000,207107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAction{}
	_p.SetDriver(__rv,207000,false)
	return _p
} 
//QAction::QAction(QString const&,QObject*)	
func NewActionWithTextParent(text string,parent QObjectInterface) *QAction {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),207000,207108,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAction{}
	_p.SetDriver(__rv,207000,false)
	return _p
} 
//QAction::QAction(QIcon const&,QString const&,QObject*)	
func NewActionWithIconTextParent(icon *QIcon,text string,parent QObjectInterface) *QAction {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),207000,207109,Native(icon),unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAction{}
	_p.SetDriver(__rv,207000,false)
	return _p
} 
//QAction::actionGroup()
func (q *QAction) ActionGroup() *QActionGroup {
	var __rv uintptr
	q.Drv(207000,207110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QActionGroup{}
	_rp.SetDriver(__rv,208000,false)
	return _rp
}	
//QAction::activate(QAction::ActionEvent)
func (q *QAction) Activate(event QAction_ActionEvent)  {
	q.Drv(207000,207111,unsafe.Pointer(&event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::associatedGraphicsWidgets()
func (q *QAction) AssociatedGraphicsWidgets() []*QGraphicsWidget {
	var __rv []*QGraphicsWidget
	q.Drv(207000,207112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::associatedWidgets()
func (q *QAction) AssociatedWidgets() []*QWidget {
	var __rv []*QWidget
	q.Drv(207000,207113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::autoRepeat()
func (q *QAction) AutoRepeat() bool {
	var __rv bool
	q.Drv(207000,207114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::data()
func (q *QAction) Data() *QVariant {
	var __rv uintptr
	q.Drv(207000,207115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAction::event(QEvent*)
func (q *QAction) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(207000,207116,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::font()
func (q *QAction) Font() *QFont {
	var __rv uintptr
	q.Drv(207000,207117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QAction::hover()
func (q *QAction) Hover()  {
	q.Drv(207000,207118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::icon()
func (q *QAction) Icon() *QIcon {
	var __rv uintptr
	q.Drv(207000,207119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QAction::iconText()
func (q *QAction) IconText() string {
	var __rv string
	q.Drv(207000,207120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isCheckable()
func (q *QAction) IsCheckable() bool {
	var __rv bool
	q.Drv(207000,207121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isChecked()
func (q *QAction) IsChecked() bool {
	var __rv bool
	q.Drv(207000,207122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isEnabled()
func (q *QAction) IsEnabled() bool {
	var __rv bool
	q.Drv(207000,207123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isIconVisibleInMenu()
func (q *QAction) IsIconVisibleInMenu() bool {
	var __rv bool
	q.Drv(207000,207124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isSeparator()
func (q *QAction) IsSeparator() bool {
	var __rv bool
	q.Drv(207000,207125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::isVisible()
func (q *QAction) IsVisible() bool {
	var __rv bool
	q.Drv(207000,207126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::menu()
func (q *QAction) Menu() *QMenu {
	var __rv uintptr
	q.Drv(207000,207127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QAction::menuRole()
func (q *QAction) MenuRole() QAction_MenuRole {
	var __rv QAction_MenuRole
	q.Drv(207000,207128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::parentWidget()
func (q *QAction) ParentWidget() *QWidget {
	var __rv uintptr
	q.Drv(207000,207129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QAction::priority()
func (q *QAction) Priority() QAction_Priority {
	var __rv QAction_Priority
	q.Drv(207000,207130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::setActionGroup(QActionGroup*)
func (q *QAction) SetActionGroup(group *QActionGroup)  {
	q.Drv(207000,207131,Native(group),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setAutoRepeat(bool)
func (q *QAction) SetAutoRepeat(value bool)  {
	q.Drv(207000,207132,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setCheckable(bool)
func (q *QAction) SetCheckable(value bool)  {
	q.Drv(207000,207133,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setChecked(bool)
func (q *QAction) SetChecked(value bool)  {
	q.Drv(207000,207134,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setData(QVariant const&)
func (q *QAction) SetData(_var *QVariant)  {
	q.Drv(207000,207135,Native(_var),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setDisabled(bool)
func (q *QAction) SetDisabled(b bool)  {
	q.Drv(207000,207136,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setEnabled(bool)
func (q *QAction) SetEnabled(value bool)  {
	q.Drv(207000,207137,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setFont(QFont const&)
func (q *QAction) SetFont(font *QFont)  {
	q.Drv(207000,207138,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setIcon(QIcon const&)
func (q *QAction) SetIcon(icon *QIcon)  {
	q.Drv(207000,207139,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setIconText(QString const&)
func (q *QAction) SetIconText(text string)  {
	q.Drv(207000,207140,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setIconVisibleInMenu(bool)
func (q *QAction) SetIconVisibleInMenu(visible bool)  {
	q.Drv(207000,207141,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setMenu(QMenu*)
func (q *QAction) SetMenu(menu *QMenu)  {
	q.Drv(207000,207142,Native(menu),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setMenuRole(QAction::MenuRole)
func (q *QAction) SetMenuRole(menuRole QAction_MenuRole)  {
	q.Drv(207000,207143,unsafe.Pointer(&menuRole),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setPriority(QAction::Priority)
func (q *QAction) SetPriority(priority QAction_Priority)  {
	q.Drv(207000,207144,unsafe.Pointer(&priority),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setSeparator(bool)
func (q *QAction) SetSeparator(b bool)  {
	q.Drv(207000,207145,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setShortcut(QKeySequence const&)
func (q *QAction) SetShortcut(shortcut *QKeySequence)  {
	q.Drv(207000,207146,Native(shortcut),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setShortcutContext(Qt::ShortcutContext)
func (q *QAction) SetShortcutContext(context Qt_ShortcutContext)  {
	q.Drv(207000,207147,unsafe.Pointer(&context),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setShortcuts(QKeySequence::StandardKey)
func (q *QAction) SetShortcuts(value QKeySequence_StandardKey)  {
	q.Drv(207000,207148,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setShortcuts(QList<QKeySequence> const&)
func (q *QAction) SetShortcutsWithShortcuts(shortcuts []*QKeySequence)  {
	q.Drv(207000,207149,unsafe.Pointer(&shortcuts),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setStatusTip(QString const&)
func (q *QAction) SetStatusTip(statusTip string)  {
	q.Drv(207000,207150,unsafe.Pointer(&statusTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setText(QString const&)
func (q *QAction) SetText(text string)  {
	q.Drv(207000,207151,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setToolTip(QString const&)
func (q *QAction) SetToolTip(tip string)  {
	q.Drv(207000,207152,unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setVisible(bool)
func (q *QAction) SetVisible(value bool)  {
	q.Drv(207000,207153,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::setWhatsThis(QString const&)
func (q *QAction) SetWhatsThis(what string)  {
	q.Drv(207000,207154,unsafe.Pointer(&what),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::shortcut()
func (q *QAction) Shortcut() *QKeySequence {
	var __rv uintptr
	q.Drv(207000,207155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QAction::shortcutContext()
func (q *QAction) ShortcutContext() Qt_ShortcutContext {
	var __rv Qt_ShortcutContext
	q.Drv(207000,207156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::shortcuts()
func (q *QAction) Shortcuts() []*QKeySequence {
	var __rv []*QKeySequence
	q.Drv(207000,207157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::showStatusText()
func (q *QAction) ShowStatusText() bool {
	var __rv bool
	q.Drv(207000,207158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::showStatusText(QWidget*)
func (q *QAction) ShowStatusTextWithWidget(widget QWidgetInterface) bool {
	var __rv bool
	q.Drv(207000,207159,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::statusTip()
func (q *QAction) StatusTip() string {
	var __rv string
	q.Drv(207000,207160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::text()
func (q *QAction) Text() string {
	var __rv string
	q.Drv(207000,207161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::toggle()
func (q *QAction) Toggle()  {
	q.Drv(207000,207162,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::toolTip()
func (q *QAction) ToolTip() string {
	var __rv string
	q.Drv(207000,207163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAction::trigger()
func (q *QAction) Trigger()  {
	q.Drv(207000,207164,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAction::whatsThis()
func (q *QAction) WhatsThis() string {
	var __rv string
	q.Drv(207000,207165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
