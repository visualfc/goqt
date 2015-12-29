// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QShortcut : QShortcut
type QShortcut struct {
	QObject
}
func (q *QShortcut) OnActivatedAmbiguously(fn func()) uintptr {
	var __rv uintptr
	q.Drv(341000,341102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QShortcut) OnActivated(fn func()) uintptr {
	var __rv uintptr
	q.Drv(341000,341103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QShortcut::QShortcut(QWidget*)	
func NewShortcut(parent QWidgetInterface) *QShortcut {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),341000,341104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QShortcut{}
	_p.SetDriver(__rv,341000,false)
	return _p
} 
//QShortcut::QShortcut(QKeySequence const&,QWidget*,char const*,char const*,Qt::ShortcutContext)	
func NewShortcutWithKeyParentMemberAmbiguousmemberContext(key *QKeySequence,parent QWidgetInterface,member string,ambiguousMember string,context Qt_ShortcutContext) *QShortcut {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),341000,341105,Native(key),Native(parent),unsafe.Pointer(&member),unsafe.Pointer(&ambiguousMember),unsafe.Pointer(&context),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QShortcut{}
	_p.SetDriver(__rv,341000,false)
	return _p
} 
//QShortcut::autoRepeat()
func (q *QShortcut) AutoRepeat() bool {
	var __rv bool
	q.Drv(341000,341106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcut::context()
func (q *QShortcut) Context() Qt_ShortcutContext {
	var __rv Qt_ShortcutContext
	q.Drv(341000,341107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcut::event(QEvent*)
func (q *QShortcut) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(341000,341108,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcut::id()
func (q *QShortcut) Id() int {
	var __rv int
	q.Drv(341000,341109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcut::isEnabled()
func (q *QShortcut) IsEnabled() bool {
	var __rv bool
	q.Drv(341000,341110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QShortcut::key()
func (q *QShortcut) Key() *QKeySequence {
	var __rv uintptr
	q.Drv(341000,341111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QKeySequence{}
	_rp.SetDriver(__rv,65000,true)
	return _rp
}	
//QShortcut::parentWidget()
func (q *QShortcut) ParentWidget() *QWidget {
	var __rv uintptr
	q.Drv(341000,341112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QShortcut::setAutoRepeat(bool)
func (q *QShortcut) SetAutoRepeat(on bool)  {
	q.Drv(341000,341113,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QShortcut::setContext(Qt::ShortcutContext)
func (q *QShortcut) SetContext(context Qt_ShortcutContext)  {
	q.Drv(341000,341114,unsafe.Pointer(&context),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QShortcut::setEnabled(bool)
func (q *QShortcut) SetEnabled(enable bool)  {
	q.Drv(341000,341115,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QShortcut::setKey(QKeySequence const&)
func (q *QShortcut) SetKey(key *QKeySequence)  {
	q.Drv(341000,341116,Native(key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QShortcut::setWhatsThis(QString const&)
func (q *QShortcut) SetWhatsThis(text string)  {
	q.Drv(341000,341117,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QShortcut::whatsThis()
func (q *QShortcut) WhatsThis() string {
	var __rv string
	q.Drv(341000,341118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
