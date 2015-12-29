// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QCompleter_ModelSorting - QCompleter::ModelSorting
type QCompleter_ModelSorting uint32
const (
	QCompleter_UnsortedModel QCompleter_ModelSorting = 0
	QCompleter_CaseSensitivelySortedModel QCompleter_ModelSorting = 0
	QCompleter_CaseInsensitivelySortedModel QCompleter_ModelSorting = 1
)
//enum QCompleter_CompletionMode - QCompleter::CompletionMode
type QCompleter_CompletionMode uint32
const (
	QCompleter_PopupCompletion QCompleter_CompletionMode = 0
	QCompleter_UnfilteredPopupCompletion QCompleter_CompletionMode = 1
	QCompleter_InlineCompletion QCompleter_CompletionMode = 2
)
//struct QCompleter : QCompleter
type QCompleter struct {
	QObject
}
func (q *QCompleter) OnHighlighted(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(220000,220102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCompleter) OnHighlightedWithText(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(220000,220103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCompleter) OnActivated(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(220000,220104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QCompleter) OnActivatedWithText(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(220000,220105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QCompleter::QCompleter()	
func NewCompleter() *QCompleter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),220000,220106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCompleter{}
	_p.SetDriver(__rv,220000,false)
	return _p
} 
//QCompleter::QCompleter(QObject*)	
func NewCompleterWithParent(parent QObjectInterface) *QCompleter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),220000,220107,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCompleter{}
	_p.SetDriver(__rv,220000,false)
	return _p
} 
//QCompleter::QCompleter(QAbstractItemModel*,QObject*)	
func NewCompleterWithModelParent(model QAbstractItemModelInterface,parent QObjectInterface) *QCompleter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),220000,220108,Native(model),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCompleter{}
	_p.SetDriver(__rv,220000,false)
	return _p
} 
//QCompleter::QCompleter(QStringList const&,QObject*)	
func NewCompleterWithCompletionsParent(completions []string,parent QObjectInterface) *QCompleter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),220000,220109,unsafe.Pointer(&completions),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCompleter{}
	_p.SetDriver(__rv,220000,false)
	return _p
} 
//QCompleter::caseSensitivity()
func (q *QCompleter) CaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(220000,220110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::complete()
func (q *QCompleter) Complete()  {
	q.Drv(220000,220111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::complete(QRect const&)
func (q *QCompleter) CompleteWithRect(rect *QRect)  {
	q.Drv(220000,220112,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::completionColumn()
func (q *QCompleter) CompletionColumn() int {
	var __rv int
	q.Drv(220000,220113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::completionCount()
func (q *QCompleter) CompletionCount() int {
	var __rv int
	q.Drv(220000,220114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::completionMode()
func (q *QCompleter) CompletionMode() QCompleter_CompletionMode {
	var __rv QCompleter_CompletionMode
	q.Drv(220000,220115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::completionModel()
func (q *QCompleter) CompletionModel() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(220000,220116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QCompleter::completionPrefix()
func (q *QCompleter) CompletionPrefix() string {
	var __rv string
	q.Drv(220000,220117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::completionRole()
func (q *QCompleter) CompletionRole() int {
	var __rv int
	q.Drv(220000,220118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::currentCompletion()
func (q *QCompleter) CurrentCompletion() string {
	var __rv string
	q.Drv(220000,220119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::currentIndex()
func (q *QCompleter) CurrentIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(220000,220120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QCompleter::currentRow()
func (q *QCompleter) CurrentRow() int {
	var __rv int
	q.Drv(220000,220121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::event(QEvent*)
func (q *QCompleter) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(220000,220122,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::eventFilter(QObject*,QEvent*)
func (q *QCompleter) EventFilter(o QObjectInterface,e *QEvent) bool {
	var __rv bool
	q.Drv(220000,220123,Native(o),Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::maxVisibleItems()
func (q *QCompleter) MaxVisibleItems() int {
	var __rv int
	q.Drv(220000,220124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::model()
func (q *QCompleter) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(220000,220125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QCompleter::modelSorting()
func (q *QCompleter) ModelSorting() QCompleter_ModelSorting {
	var __rv QCompleter_ModelSorting
	q.Drv(220000,220126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::pathFromIndex(QModelIndex const&)
func (q *QCompleter) PathFromIndex(index *QModelIndex) string {
	var __rv string
	q.Drv(220000,220127,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::popup()
func (q *QCompleter) Popup() *QAbstractItemView {
	var __rv uintptr
	q.Drv(220000,220128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemView{}
	_rp.SetDriver(__rv,196000,false)
	return _rp
}	
//QCompleter::setCaseSensitivity(Qt::CaseSensitivity)
func (q *QCompleter) SetCaseSensitivity(caseSensitivity Qt_CaseSensitivity)  {
	q.Drv(220000,220129,unsafe.Pointer(&caseSensitivity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setCompletionColumn(int)
func (q *QCompleter) SetCompletionColumn(column int)  {
	q.Drv(220000,220130,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setCompletionMode(QCompleter::CompletionMode)
func (q *QCompleter) SetCompletionMode(mode QCompleter_CompletionMode)  {
	q.Drv(220000,220131,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setCompletionPrefix(QString const&)
func (q *QCompleter) SetCompletionPrefix(prefix string)  {
	q.Drv(220000,220132,unsafe.Pointer(&prefix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setCompletionRole(int)
func (q *QCompleter) SetCompletionRole(role int)  {
	q.Drv(220000,220133,unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setCurrentRow(int)
func (q *QCompleter) SetCurrentRow(row int) bool {
	var __rv bool
	q.Drv(220000,220134,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::setMaxVisibleItems(int)
func (q *QCompleter) SetMaxVisibleItems(maxItems int)  {
	q.Drv(220000,220135,unsafe.Pointer(&maxItems),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setModel(QAbstractItemModel*)
func (q *QCompleter) SetModel(c QAbstractItemModelInterface)  {
	q.Drv(220000,220136,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setModelSorting(QCompleter::ModelSorting)
func (q *QCompleter) SetModelSorting(sorting QCompleter_ModelSorting)  {
	q.Drv(220000,220137,unsafe.Pointer(&sorting),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setPopup(QAbstractItemView*)
func (q *QCompleter) SetPopup(popup *QAbstractItemView)  {
	q.Drv(220000,220138,Native(popup),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setWidget(QWidget*)
func (q *QCompleter) SetWidget(widget QWidgetInterface)  {
	q.Drv(220000,220139,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::setWrapAround(bool)
func (q *QCompleter) SetWrapAround(wrap bool)  {
	q.Drv(220000,220140,unsafe.Pointer(&wrap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCompleter::splitPath(QString const&)
func (q *QCompleter) SplitPath(path string) []string {
	var __rv []string
	q.Drv(220000,220141,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCompleter::widget()
func (q *QCompleter) Widget() *QWidget {
	var __rv uintptr
	q.Drv(220000,220142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QCompleter::wrapAround()
func (q *QCompleter) WrapAround() bool {
	var __rv bool
	q.Drv(220000,220143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
