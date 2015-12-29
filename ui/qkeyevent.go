// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QKeyEvent : QKeyEvent
type QKeyEvent struct {
	QInputEvent
}
//QKeyEvent::QKeyEvent(QEvent::Type,int,QFlags<Qt::KeyboardModifier>,QString const&,bool,unsigned short)	
func NewKeyEvent(_type QEvent_Type,key int,modifiers Qt_KeyboardModifier,text string,autorep bool,count uint16) *QKeyEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),64000,64102,unsafe.Pointer(&_type),unsafe.Pointer(&key),unsafe.Pointer(&modifiers),unsafe.Pointer(&text),unsafe.Pointer(&autorep),unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QKeyEvent{}
	_p.SetDriver(__rv,64000,true)
	return _p
} 
//QKeyEvent::count()
func (q *QKeyEvent) Count() int {
	var __rv int
	q.Drv(64000,64103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::isAutoRepeat()
func (q *QKeyEvent) IsAutoRepeat() bool {
	var __rv bool
	q.Drv(64000,64104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::key()
func (q *QKeyEvent) Key() int {
	var __rv int
	q.Drv(64000,64105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::matches(QKeySequence::StandardKey)
func (q *QKeyEvent) Matches(key QKeySequence_StandardKey) bool {
	var __rv bool
	q.Drv(64000,64106,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::modifiers()
func (q *QKeyEvent) Modifiers() Qt_KeyboardModifier {
	var __rv Qt_KeyboardModifier
	q.Drv(64000,64107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::nativeModifiers()
func (q *QKeyEvent) NativeModifiers() uint {
	var __rv uint
	q.Drv(64000,64108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::nativeScanCode()
func (q *QKeyEvent) NativeScanCode() uint {
	var __rv uint
	q.Drv(64000,64109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::nativeVirtualKey()
func (q *QKeyEvent) NativeVirtualKey() uint {
	var __rv uint
	q.Drv(64000,64110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QKeyEvent::text()
func (q *QKeyEvent) Text() string {
	var __rv string
	q.Drv(64000,64111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
