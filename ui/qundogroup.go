// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUndoGroup : QUndoGroup
type QUndoGroup struct {
	QObject
}
func (q *QUndoGroup) OnIndexChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(388000,388102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnActiveStackChanged(fn func(*QUndoStack)) uintptr {
	var __rv uintptr
	q.Drv(388000,388103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnUndoTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(388000,388104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnCanUndoChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(388000,388105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnRedoTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(388000,388106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnCanRedoChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(388000,388107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QUndoGroup) OnCleanChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(388000,388108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QUndoGroup::QUndoGroup()	
func NewUndoGroup() *QUndoGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),388000,388109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoGroup{}
	_p.SetDriver(__rv,388000,false)
	return _p
} 
//QUndoGroup::QUndoGroup(QObject*)	
func NewUndoGroupWithParent(parent QObjectInterface) *QUndoGroup {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),388000,388110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUndoGroup{}
	_p.SetDriver(__rv,388000,false)
	return _p
} 
//QUndoGroup::activeStack()
func (q *QUndoGroup) ActiveStack() *QUndoStack {
	var __rv uintptr
	q.Drv(388000,388111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUndoStack{}
	_rp.SetDriver(__rv,389000,false)
	return _rp
}	
//QUndoGroup::addStack(QUndoStack*)
func (q *QUndoGroup) AddStack(stack *QUndoStack)  {
	q.Drv(388000,388112,Native(stack),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoGroup::canRedo()
func (q *QUndoGroup) CanRedo() bool {
	var __rv bool
	q.Drv(388000,388113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoGroup::canUndo()
func (q *QUndoGroup) CanUndo() bool {
	var __rv bool
	q.Drv(388000,388114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoGroup::createRedoAction(QObject*)
func (q *QUndoGroup) CreateRedoAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(388000,388115,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoGroup::createRedoAction(QObject*,QString const&)
func (q *QUndoGroup) CreateRedoActionWithParentPrefix(parent QObjectInterface,prefix string) *QAction {
	var __rv uintptr
	q.Drv(388000,388116,Native(parent),unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoGroup::createUndoAction(QObject*)
func (q *QUndoGroup) CreateUndoAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(388000,388117,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoGroup::createUndoAction(QObject*,QString const&)
func (q *QUndoGroup) CreateUndoActionWithParentPrefix(parent QObjectInterface,prefix string) *QAction {
	var __rv uintptr
	q.Drv(388000,388118,Native(parent),unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUndoGroup::isClean()
func (q *QUndoGroup) IsClean() bool {
	var __rv bool
	q.Drv(388000,388119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoGroup::redo()
func (q *QUndoGroup) Redo()  {
	q.Drv(388000,388120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoGroup::redoText()
func (q *QUndoGroup) RedoText() string {
	var __rv string
	q.Drv(388000,388121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoGroup::removeStack(QUndoStack*)
func (q *QUndoGroup) RemoveStack(stack *QUndoStack)  {
	q.Drv(388000,388122,Native(stack),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoGroup::setActiveStack(QUndoStack*)
func (q *QUndoGroup) SetActiveStack(stack *QUndoStack)  {
	q.Drv(388000,388123,Native(stack),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoGroup::stacks()
func (q *QUndoGroup) Stacks() []*QUndoStack {
	var __rv []*QUndoStack
	q.Drv(388000,388124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUndoGroup::undo()
func (q *QUndoGroup) Undo()  {
	q.Drv(388000,388125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUndoGroup::undoText()
func (q *QUndoGroup) UndoText() string {
	var __rv string
	q.Drv(388000,388126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
