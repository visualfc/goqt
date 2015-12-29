// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSystemTrayIcon_MessageIcon - QSystemTrayIcon::MessageIcon
type QSystemTrayIcon_MessageIcon uint32
const (
	QSystemTrayIcon_NoIcon QSystemTrayIcon_MessageIcon = 0
	QSystemTrayIcon_Information QSystemTrayIcon_MessageIcon = 1
	QSystemTrayIcon_Warning QSystemTrayIcon_MessageIcon = 2
	QSystemTrayIcon_Critical QSystemTrayIcon_MessageIcon = 3
)
//enum QSystemTrayIcon_ActivationReason - QSystemTrayIcon::ActivationReason
type QSystemTrayIcon_ActivationReason uint32
const (
	QSystemTrayIcon_Unknown QSystemTrayIcon_ActivationReason = 0
	QSystemTrayIcon_Context QSystemTrayIcon_ActivationReason = 1
	QSystemTrayIcon_DoubleClick QSystemTrayIcon_ActivationReason = 2
	QSystemTrayIcon_Trigger QSystemTrayIcon_ActivationReason = 3
	QSystemTrayIcon_MiddleClick QSystemTrayIcon_ActivationReason = 4
)
//struct QSystemTrayIcon : QSystemTrayIcon
type QSystemTrayIcon struct {
	QObject
}
func (q *QSystemTrayIcon) OnMessageClicked(fn func()) uintptr {
	var __rv uintptr
	q.Drv(362000,362102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QSystemTrayIcon) OnActivated(fn func(QSystemTrayIcon_ActivationReason)) uintptr {
	var __rv uintptr
	q.Drv(362000,362103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSystemTrayIcon::QSystemTrayIcon()	
func NewSystemTrayIcon() *QSystemTrayIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),362000,362104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSystemTrayIcon{}
	_p.SetDriver(__rv,362000,false)
	return _p
} 
//QSystemTrayIcon::QSystemTrayIcon(QObject*)	
func NewSystemTrayIconWithParent(parent QObjectInterface) *QSystemTrayIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),362000,362105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSystemTrayIcon{}
	_p.SetDriver(__rv,362000,false)
	return _p
} 
//QSystemTrayIcon::QSystemTrayIcon(QIcon const&,QObject*)	
func NewSystemTrayIconWithIconParent(icon *QIcon,parent QObjectInterface) *QSystemTrayIcon {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),362000,362106,Native(icon),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSystemTrayIcon{}
	_p.SetDriver(__rv,362000,false)
	return _p
} 
//QSystemTrayIcon::contextMenu()
func (q *QSystemTrayIcon) ContextMenu() *QMenu {
	var __rv uintptr
	q.Drv(362000,362107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QSystemTrayIcon::event(QEvent*)
func (q *QSystemTrayIcon) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(362000,362108,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::geometry()
func (q *QSystemTrayIcon) Geometry() *QRect {
	var __rv uintptr
	q.Drv(362000,362109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QSystemTrayIcon::hide()
func (q *QSystemTrayIcon) Hide()  {
	q.Drv(362000,362110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::icon()
func (q *QSystemTrayIcon) Icon() *QIcon {
	var __rv uintptr
	q.Drv(362000,362111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QSystemTrayIcon::isSystemTrayAvailable()	
func QSystemTrayIconIsSystemTrayAvailable() bool {
	var __rv bool
	DirectQtDrv(nil,362000,362112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::isSystemTrayAvailable()
func (q *QSystemTrayIcon) IsSystemTrayAvailable() bool {
	var __rv bool
	q.Drv(362000,362112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::isVisible()
func (q *QSystemTrayIcon) IsVisible() bool {
	var __rv bool
	q.Drv(362000,362113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::setContextMenu(QMenu*)
func (q *QSystemTrayIcon) SetContextMenu(menu *QMenu)  {
	q.Drv(362000,362114,Native(menu),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::setIcon(QIcon const&)
func (q *QSystemTrayIcon) SetIcon(icon *QIcon)  {
	q.Drv(362000,362115,Native(icon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::setToolTip(QString const&)
func (q *QSystemTrayIcon) SetToolTip(tip string)  {
	q.Drv(362000,362116,unsafe.Pointer(&tip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::setVisible(bool)
func (q *QSystemTrayIcon) SetVisible(visible bool)  {
	q.Drv(362000,362117,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::show()
func (q *QSystemTrayIcon) Show()  {
	q.Drv(362000,362118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::showMessage(QString const&,QString const&,QSystemTrayIcon::MessageIcon,int)
func (q *QSystemTrayIcon) ShowMessage(title string,msg string,icon QSystemTrayIcon_MessageIcon,msecs int)  {
	q.Drv(362000,362119,unsafe.Pointer(&title),unsafe.Pointer(&msg),unsafe.Pointer(&icon),unsafe.Pointer(&msecs),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSystemTrayIcon::supportsMessages()	
func QSystemTrayIconSupportsMessages() bool {
	var __rv bool
	DirectQtDrv(nil,362000,362120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::supportsMessages()
func (q *QSystemTrayIcon) SupportsMessages() bool {
	var __rv bool
	q.Drv(362000,362120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSystemTrayIcon::toolTip()
func (q *QSystemTrayIcon) ToolTip() string {
	var __rv string
	q.Drv(362000,362121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
