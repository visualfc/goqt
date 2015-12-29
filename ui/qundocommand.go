// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUndoCommand : QUndoCommand
type QUndoCommand struct {
	BaseDrv
}
//QUndoCommand::QUndoCommand()	
func NewUndoCommand() *QUndoCommand {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),179000,179102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoCommand{}
	_p.SetDriver(__rv,179000,true)
	return _p
} 
//QUndoCommand::QUndoCommand(QUndoCommand*)	
func NewUndoCommandWithParent(parent *QUndoCommand) *QUndoCommand {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),179000,179103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoCommand{}
	_p.SetDriver(__rv,179000,true)
	return _p
} 
//QUndoCommand::QUndoCommand(QString const&,QUndoCommand*)	
func NewUndoCommandWithTextParent(text string,parent *QUndoCommand) *QUndoCommand {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),179000,179104,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoCommand{}
	_p.SetDriver(__rv,179000,true)
	return _p
} 
//QUndoCommand::child(int)
func (q *QUndoCommand) Child(index int) *QUndoCommand {
	var __rv uintptr
	q.Drv(179000,179105,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUndoCommand{}
	_rp.SetDriver(__rv,179000,true)
	return _rp
}	
//QUndoCommand::childCount()
func (q *QUndoCommand) ChildCount() int {
	var __rv int
	q.Drv(179000,179106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoCommand::id()
func (q *QUndoCommand) Id() int {
	var __rv int
	q.Drv(179000,179107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoCommand::mergeWith(QUndoCommand const*)
func (q *QUndoCommand) MergeWith(other *QUndoCommand) bool {
	var __rv bool
	q.Drv(179000,179108,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoCommand::redo()
func (q *QUndoCommand) Redo()  {
	q.Drv(179000,179109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoCommand::setText(QString const&)
func (q *QUndoCommand) SetText(text string)  {
	q.Drv(179000,179110,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoCommand::text()
func (q *QUndoCommand) Text() string {
	var __rv string
	q.Drv(179000,179111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoCommand::undo()
func (q *QUndoCommand) Undo()  {
	q.Drv(179000,179112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
