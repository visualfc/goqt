// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QEvent_Type - QEvent::Type
type QEvent_Type uint32
const (
	QEvent_None QEvent_Type = 0
	QEvent_Timer QEvent_Type = 1
	QEvent_MouseButtonPress QEvent_Type = 2
	QEvent_MouseButtonRelease QEvent_Type = 3
	QEvent_MouseButtonDblClick QEvent_Type = 4
	QEvent_MouseMove QEvent_Type = 5
	QEvent_KeyPress QEvent_Type = 6
	QEvent_KeyRelease QEvent_Type = 7
	QEvent_FocusIn QEvent_Type = 8
	QEvent_FocusOut QEvent_Type = 9
	QEvent_Enter QEvent_Type = 10
	QEvent_Leave QEvent_Type = 11
	QEvent_Paint QEvent_Type = 12
	QEvent_Move QEvent_Type = 13
	QEvent_Resize QEvent_Type = 14
	QEvent_Create QEvent_Type = 15
	QEvent_Destroy QEvent_Type = 16
	QEvent_Show QEvent_Type = 17
	QEvent_Hide QEvent_Type = 18
	QEvent_Close QEvent_Type = 19
	QEvent_Quit QEvent_Type = 20
	QEvent_ParentChange QEvent_Type = 21
	QEvent_ParentAboutToChange QEvent_Type = 131
	QEvent_ThreadChange QEvent_Type = 22
	QEvent_WindowActivate QEvent_Type = 24
	QEvent_WindowDeactivate QEvent_Type = 25
	QEvent_ShowToParent QEvent_Type = 26
	QEvent_HideToParent QEvent_Type = 27
	QEvent_Wheel QEvent_Type = 31
	QEvent_WindowTitleChange QEvent_Type = 33
	QEvent_WindowIconChange QEvent_Type = 34
	QEvent_ApplicationWindowIconChange QEvent_Type = 35
	QEvent_ApplicationFontChange QEvent_Type = 36
	QEvent_ApplicationLayoutDirectionChange QEvent_Type = 37
	QEvent_ApplicationPaletteChange QEvent_Type = 38
	QEvent_PaletteChange QEvent_Type = 39
	QEvent_Clipboard QEvent_Type = 40
	QEvent_Speech QEvent_Type = 42
	QEvent_MetaCall QEvent_Type = 43
	QEvent_SockAct QEvent_Type = 50
	QEvent_WinEventAct QEvent_Type = 132
	QEvent_DeferredDelete QEvent_Type = 52
	QEvent_DragEnter QEvent_Type = 60
	QEvent_DragMove QEvent_Type = 61
	QEvent_DragLeave QEvent_Type = 62
	QEvent_Drop QEvent_Type = 63
	QEvent_DragResponse QEvent_Type = 64
	QEvent_ChildAdded QEvent_Type = 68
	QEvent_ChildPolished QEvent_Type = 69
	QEvent_ChildRemoved QEvent_Type = 71
	QEvent_ShowWindowRequest QEvent_Type = 73
	QEvent_PolishRequest QEvent_Type = 74
	QEvent_Polish QEvent_Type = 75
	QEvent_LayoutRequest QEvent_Type = 76
	QEvent_UpdateRequest QEvent_Type = 77
	QEvent_UpdateLater QEvent_Type = 78
	QEvent_EmbeddingControl QEvent_Type = 79
	QEvent_ActivateControl QEvent_Type = 80
	QEvent_DeactivateControl QEvent_Type = 81
	QEvent_ContextMenu QEvent_Type = 82
	QEvent_InputMethod QEvent_Type = 83
	QEvent_AccessibilityPrepare QEvent_Type = 86
	QEvent_TabletMove QEvent_Type = 87
	QEvent_LocaleChange QEvent_Type = 88
	QEvent_LanguageChange QEvent_Type = 89
	QEvent_LayoutDirectionChange QEvent_Type = 90
	QEvent_Style QEvent_Type = 91
	QEvent_TabletPress QEvent_Type = 92
	QEvent_TabletRelease QEvent_Type = 93
	QEvent_OkRequest QEvent_Type = 94
	QEvent_HelpRequest QEvent_Type = 95
	QEvent_IconDrag QEvent_Type = 96
	QEvent_FontChange QEvent_Type = 97
	QEvent_EnabledChange QEvent_Type = 98
	QEvent_ActivationChange QEvent_Type = 99
	QEvent_StyleChange QEvent_Type = 100
	QEvent_IconTextChange QEvent_Type = 101
	QEvent_ModifiedChange QEvent_Type = 102
	QEvent_MouseTrackingChange QEvent_Type = 109
	QEvent_WindowBlocked QEvent_Type = 103
	QEvent_WindowUnblocked QEvent_Type = 104
	QEvent_WindowStateChange QEvent_Type = 105
	QEvent_ToolTip QEvent_Type = 110
	QEvent_WhatsThis QEvent_Type = 111
	QEvent_StatusTip QEvent_Type = 112
	QEvent_ActionChanged QEvent_Type = 113
	QEvent_ActionAdded QEvent_Type = 114
	QEvent_ActionRemoved QEvent_Type = 115
	QEvent_FileOpen QEvent_Type = 116
	QEvent_Shortcut QEvent_Type = 117
	QEvent_ShortcutOverride QEvent_Type = 51
	QEvent_WhatsThisClicked QEvent_Type = 118
	QEvent_ToolBarChange QEvent_Type = 120
	QEvent_ApplicationActivate QEvent_Type = 121
	QEvent_ApplicationDeactivate QEvent_Type = 122
	QEvent_QueryWhatsThis QEvent_Type = 123
	QEvent_EnterWhatsThisMode QEvent_Type = 124
	QEvent_LeaveWhatsThisMode QEvent_Type = 125
	QEvent_ZOrderChange QEvent_Type = 126
	QEvent_HoverEnter QEvent_Type = 127
	QEvent_HoverLeave QEvent_Type = 128
	QEvent_HoverMove QEvent_Type = 129
	QEvent_AccessibilityHelp QEvent_Type = 119
	QEvent_AccessibilityDescription QEvent_Type = 130
	QEvent_AcceptDropsChange QEvent_Type = 152
	QEvent_MenubarUpdated QEvent_Type = 153
	QEvent_ZeroTimerEvent QEvent_Type = 154
	QEvent_GraphicsSceneMouseMove QEvent_Type = 155
	QEvent_GraphicsSceneMousePress QEvent_Type = 156
	QEvent_GraphicsSceneMouseRelease QEvent_Type = 157
	QEvent_GraphicsSceneMouseDoubleClick QEvent_Type = 158
	QEvent_GraphicsSceneContextMenu QEvent_Type = 159
	QEvent_GraphicsSceneHoverEnter QEvent_Type = 160
	QEvent_GraphicsSceneHoverMove QEvent_Type = 161
	QEvent_GraphicsSceneHoverLeave QEvent_Type = 162
	QEvent_GraphicsSceneHelp QEvent_Type = 163
	QEvent_GraphicsSceneDragEnter QEvent_Type = 164
	QEvent_GraphicsSceneDragMove QEvent_Type = 165
	QEvent_GraphicsSceneDragLeave QEvent_Type = 166
	QEvent_GraphicsSceneDrop QEvent_Type = 167
	QEvent_GraphicsSceneWheel QEvent_Type = 168
	QEvent_KeyboardLayoutChange QEvent_Type = 169
	QEvent_DynamicPropertyChange QEvent_Type = 170
	QEvent_TabletEnterProximity QEvent_Type = 171
	QEvent_TabletLeaveProximity QEvent_Type = 172
	QEvent_NonClientAreaMouseMove QEvent_Type = 173
	QEvent_NonClientAreaMouseButtonPress QEvent_Type = 174
	QEvent_NonClientAreaMouseButtonRelease QEvent_Type = 175
	QEvent_NonClientAreaMouseButtonDblClick QEvent_Type = 176
	QEvent_MacSizeChange QEvent_Type = 177
	QEvent_ContentsRectChange QEvent_Type = 178
	QEvent_MacGLWindowChange QEvent_Type = 179
	QEvent_FutureCallOut QEvent_Type = 180
	QEvent_GraphicsSceneResize QEvent_Type = 181
	QEvent_GraphicsSceneMove QEvent_Type = 182
	QEvent_CursorChange QEvent_Type = 183
	QEvent_ToolTipChange QEvent_Type = 184
	QEvent_NetworkReplyUpdated QEvent_Type = 185
	QEvent_GrabMouse QEvent_Type = 186
	QEvent_UngrabMouse QEvent_Type = 187
	QEvent_GrabKeyboard QEvent_Type = 188
	QEvent_UngrabKeyboard QEvent_Type = 189
	QEvent_MacGLClearDrawable QEvent_Type = 191
	QEvent_StateMachineSignal QEvent_Type = 192
	QEvent_StateMachineWrapped QEvent_Type = 193
	QEvent_TouchBegin QEvent_Type = 194
	QEvent_TouchUpdate QEvent_Type = 195
	QEvent_TouchEnd QEvent_Type = 196
	QEvent_NativeGesture QEvent_Type = 197
	QEvent_RequestSoftwareInputPanel QEvent_Type = 199
	QEvent_CloseSoftwareInputPanel QEvent_Type = 200
	QEvent_UpdateSoftKeys QEvent_Type = 201
	QEvent_WinIdChange QEvent_Type = 203
	QEvent_Gesture QEvent_Type = 198
	QEvent_GestureOverride QEvent_Type = 202
	QEvent_User QEvent_Type = 1000
	QEvent_MaxUser QEvent_Type = 65535
)
//struct QEvent : QEvent
type QEvent struct {
	BaseDrv
}
//QEvent::QEvent(QEvent::Type)	
func NewEvent(_type QEvent_Type) *QEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),31000,31102,unsafe.Pointer(&_type),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEvent{}
	_p.SetDriver(__rv,31000,true)
	return _p
} 
//QEvent::accept()
func (q *QEvent) Accept()  {
	q.Drv(31000,31103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEvent::ignore()
func (q *QEvent) Ignore()  {
	q.Drv(31000,31104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEvent::isAccepted()
func (q *QEvent) IsAccepted() bool {
	var __rv bool
	q.Drv(31000,31105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::registerEventType()	
func QEventRegisterEventType() int {
	var __rv int
	DirectQtDrv(nil,31000,31106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::registerEventType()
func (q *QEvent) RegisterEventType() int {
	var __rv int
	q.Drv(31000,31106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::registerEventType(int)	
func QEventRegisterEventTypeWithHint(hint int) int {
	var __rv int
	DirectQtDrv(nil,31000,31107,unsafe.Pointer(&hint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::registerEventType(int)
func (q *QEvent) RegisterEventTypeWithHint(hint int) int {
	var __rv int
	q.Drv(31000,31107,unsafe.Pointer(&hint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::setAccepted(bool)
func (q *QEvent) SetAccepted(accepted bool)  {
	q.Drv(31000,31108,unsafe.Pointer(&accepted),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEvent::spontaneous()
func (q *QEvent) Spontaneous() bool {
	var __rv bool
	q.Drv(31000,31109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEvent::type()
func (q *QEvent) Type() QEvent_Type {
	var __rv QEvent_Type
	q.Drv(31000,31110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
