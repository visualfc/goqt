// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUndoStack : QUndoStack
type QUndoStack struct {
	QObject
}
func (q *QUndoStack) OnIndexChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(389000,389102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoStack) OnUndoTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(389000,389103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoStack) OnCanUndoChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(389000,389104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoStack) OnRedoTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(389000,389105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoStack) OnCanRedoChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(389000,389106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoStack) OnCleanChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(389000,389107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QUndoStack::QUndoStack()	
func NewUndoStack() *QUndoStack {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),389000,389108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoStack{}
	_p.SetDriver(__rv,389000,false)
	return _p
} 
//QUndoStack::QUndoStack(QObject*)	
func NewUndoStackWithParent(parent QObjectInterface) *QUndoStack {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),389000,389109,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoStack{}
	_p.SetDriver(__rv,389000,false)
	return _p
} 
//QUndoStack::beginMacro(QString const&)
func (q *QUndoStack) BeginMacro(text string)  {
	q.Drv(389000,389110,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::canRedo()
func (q *QUndoStack) CanRedo() bool {
	var __rv bool
	q.Drv(389000,389111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::canUndo()
func (q *QUndoStack) CanUndo() bool {
	var __rv bool
	q.Drv(389000,389112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::cleanIndex()
func (q *QUndoStack) CleanIndex() int {
	var __rv int
	q.Drv(389000,389113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::clear()
func (q *QUndoStack) Clear()  {
	q.Drv(389000,389114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::command(int)
func (q *QUndoStack) Command(index int) *QUndoCommand {
	var __rv uintptr
	q.Drv(389000,389115,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUndoCommand{}
	_rp.SetDriver(__rv,179000,true)
	return _rp
}	
//QUndoStack::count()
func (q *QUndoStack) Count() int {
	var __rv int
	q.Drv(389000,389116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::createRedoAction(QObject*)
func (q *QUndoStack) CreateRedoAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(389000,389117,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoStack::createRedoAction(QObject*,QString const&)
func (q *QUndoStack) CreateRedoActionWithParentPrefix(parent QObjectInterface,prefix string) *QAction {
	var __rv uintptr
	q.Drv(389000,389118,Native(parent),unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoStack::createUndoAction(QObject*)
func (q *QUndoStack) CreateUndoAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(389000,389119,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoStack::createUndoAction(QObject*,QString const&)
func (q *QUndoStack) CreateUndoActionWithParentPrefix(parent QObjectInterface,prefix string) *QAction {
	var __rv uintptr
	q.Drv(389000,389120,Native(parent),unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoStack::endMacro()
func (q *QUndoStack) EndMacro()  {
	q.Drv(389000,389121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::index()
func (q *QUndoStack) Index() int {
	var __rv int
	q.Drv(389000,389122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::isActive()
func (q *QUndoStack) IsActive() bool {
	var __rv bool
	q.Drv(389000,389123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::isClean()
func (q *QUndoStack) IsClean() bool {
	var __rv bool
	q.Drv(389000,389124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::push(QUndoCommand*)
func (q *QUndoStack) Push(cmd *QUndoCommand)  {
	q.Drv(389000,389125,Native(cmd),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::redo()
func (q *QUndoStack) Redo()  {
	q.Drv(389000,389126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::redoText()
func (q *QUndoStack) RedoText() string {
	var __rv string
	q.Drv(389000,389127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::setActive(bool)
func (q *QUndoStack) SetActive(active bool)  {
	q.Drv(389000,389128,unsafe.Pointer(&active),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::setClean()
func (q *QUndoStack) SetClean()  {
	q.Drv(389000,389129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::setIndex(int)
func (q *QUndoStack) SetIndex(idx int)  {
	q.Drv(389000,389130,unsafe.Pointer(&idx),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::setUndoLimit(int)
func (q *QUndoStack) SetUndoLimit(limit int)  {
	q.Drv(389000,389131,unsafe.Pointer(&limit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::text(int)
func (q *QUndoStack) Text(idx int) string {
	var __rv string
	q.Drv(389000,389132,unsafe.Pointer(&idx),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::undo()
func (q *QUndoStack) Undo()  {
	q.Drv(389000,389133,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoStack::undoLimit()
func (q *QUndoStack) UndoLimit() int {
	var __rv int
	q.Drv(389000,389134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoStack::undoText()
func (q *QUndoStack) UndoText() string {
	var __rv string
	q.Drv(389000,389135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
