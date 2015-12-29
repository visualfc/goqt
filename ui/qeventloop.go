// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QEventLoop_ProcessEventsFlag - QEventLoop::ProcessEventsFlag
type QEventLoop_ProcessEventsFlag uint32
const (
	QEventLoop_AllEvents QEventLoop_ProcessEventsFlag = 0x00
	QEventLoop_ExcludeUserInputEvents QEventLoop_ProcessEventsFlag = 0x01
	QEventLoop_ExcludeSocketNotifiers QEventLoop_ProcessEventsFlag = 0x02
	QEventLoop_WaitForMoreEvents QEventLoop_ProcessEventsFlag = 0x04
	QEventLoop_X11ExcludeTimers QEventLoop_ProcessEventsFlag = 0x08
	QEventLoop_DeferredDeletion QEventLoop_ProcessEventsFlag = 0x10
	QEventLoop_EventLoopExec QEventLoop_ProcessEventsFlag = 0x20
	QEventLoop_DialogExec QEventLoop_ProcessEventsFlag = 0x40
)
//struct QEventLoop : QEventLoop
type QEventLoop struct {
	QObject
}
//QEventLoop::QEventLoop()	
func NewEventLoop() *QEventLoop {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),234000,234102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEventLoop{}
	_p.SetDriver(__rv,234000,false)
	return _p
} 
//QEventLoop::QEventLoop(QObject*)	
func NewEventLoopWithParent(parent QObjectInterface) *QEventLoop {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),234000,234103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QEventLoop{}
	_p.SetDriver(__rv,234000,false)
	return _p
} 
//QEventLoop::exec()
func (q *QEventLoop) Exec() int {
	var __rv int
	q.Drv(234000,234104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventLoop::exec(QFlags<QEventLoop::ProcessEventsFlag>)
func (q *QEventLoop) ExecWithFlags(flags QEventLoop_ProcessEventsFlag) int {
	var __rv int
	q.Drv(234000,234105,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventLoop::exit()
func (q *QEventLoop) Exit()  {
	q.Drv(234000,234106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventLoop::exit(int)
func (q *QEventLoop) ExitWithReturncode(returnCode int)  {
	q.Drv(234000,234107,unsafe.Pointer(&returnCode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventLoop::isRunning()
func (q *QEventLoop) IsRunning() bool {
	var __rv bool
	q.Drv(234000,234108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventLoop::processEvents()
func (q *QEventLoop) ProcessEvents() bool {
	var __rv bool
	q.Drv(234000,234109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventLoop::processEvents(QFlags<QEventLoop::ProcessEventsFlag>)
func (q *QEventLoop) ProcessEventsWithFlags(flags QEventLoop_ProcessEventsFlag) bool {
	var __rv bool
	q.Drv(234000,234110,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QEventLoop::processEvents(QFlags<QEventLoop::ProcessEventsFlag>,int)
func (q *QEventLoop) ProcessEventsWithFlagsMaximumtime(flags QEventLoop_ProcessEventsFlag,maximumTime int)  {
	q.Drv(234000,234111,unsafe.Pointer(&flags),unsafe.Pointer(&maximumTime),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventLoop::quit()
func (q *QEventLoop) Quit()  {
	q.Drv(234000,234112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QEventLoop::wakeUp()
func (q *QEventLoop) WakeUp()  {
	q.Drv(234000,234113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
