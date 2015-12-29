// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractItemView_DragDropMode - QAbstractItemView::DragDropMode
type QAbstractItemView_DragDropMode uint32
const (
	QAbstractItemView_NoDragDrop QAbstractItemView_DragDropMode = 0
	QAbstractItemView_DragOnly QAbstractItemView_DragDropMode = 1
	QAbstractItemView_DropOnly QAbstractItemView_DragDropMode = 2
	QAbstractItemView_DragDrop QAbstractItemView_DragDropMode = 3
	QAbstractItemView_InternalMove QAbstractItemView_DragDropMode = 4
)
//enum QAbstractItemView_SelectionBehavior - QAbstractItemView::SelectionBehavior
type QAbstractItemView_SelectionBehavior uint32
const (
	QAbstractItemView_SelectItems QAbstractItemView_SelectionBehavior = 0
	QAbstractItemView_SelectRows QAbstractItemView_SelectionBehavior = 1
	QAbstractItemView_SelectColumns QAbstractItemView_SelectionBehavior = 2
)
//enum QAbstractItemView_EditTrigger - QAbstractItemView::EditTrigger
type QAbstractItemView_EditTrigger uint32
const (
	QAbstractItemView_NoEditTriggers QAbstractItemView_EditTrigger = 0
	QAbstractItemView_CurrentChanged QAbstractItemView_EditTrigger = 1
	QAbstractItemView_DoubleClicked QAbstractItemView_EditTrigger = 2
	QAbstractItemView_SelectedClicked QAbstractItemView_EditTrigger = 4
	QAbstractItemView_EditKeyPressed QAbstractItemView_EditTrigger = 8
	QAbstractItemView_AnyKeyPressed QAbstractItemView_EditTrigger = 16
	QAbstractItemView_AllEditTriggers QAbstractItemView_EditTrigger = 31
)
//enum QAbstractItemView_ScrollMode - QAbstractItemView::ScrollMode
type QAbstractItemView_ScrollMode uint32
const (
	QAbstractItemView_ScrollPerItem QAbstractItemView_ScrollMode = 0
	QAbstractItemView_ScrollPerPixel QAbstractItemView_ScrollMode = 1
)
//enum QAbstractItemView_CursorAction - QAbstractItemView::CursorAction
type QAbstractItemView_CursorAction uint32
const (
	QAbstractItemView_MoveUp QAbstractItemView_CursorAction = 0
	QAbstractItemView_MoveDown QAbstractItemView_CursorAction = 1
	QAbstractItemView_MoveLeft QAbstractItemView_CursorAction = 2
	QAbstractItemView_MoveRight QAbstractItemView_CursorAction = 3
	QAbstractItemView_MoveHome QAbstractItemView_CursorAction = 4
	QAbstractItemView_MoveEnd QAbstractItemView_CursorAction = 5
	QAbstractItemView_MovePageUp QAbstractItemView_CursorAction = 6
	QAbstractItemView_MovePageDown QAbstractItemView_CursorAction = 7
	QAbstractItemView_MoveNext QAbstractItemView_CursorAction = 8
	QAbstractItemView_MovePrevious QAbstractItemView_CursorAction = 9
)
//enum QAbstractItemView_DropIndicatorPosition - QAbstractItemView::DropIndicatorPosition
type QAbstractItemView_DropIndicatorPosition uint32
const (
	QAbstractItemView_OnItem QAbstractItemView_DropIndicatorPosition = 0
	QAbstractItemView_AboveItem QAbstractItemView_DropIndicatorPosition = 1
	QAbstractItemView_BelowItem QAbstractItemView_DropIndicatorPosition = 2
	QAbstractItemView_OnViewport QAbstractItemView_DropIndicatorPosition = 3
)
//enum QAbstractItemView_ScrollHint - QAbstractItemView::ScrollHint
type QAbstractItemView_ScrollHint uint32
const (
	QAbstractItemView_EnsureVisible QAbstractItemView_ScrollHint = 0
	QAbstractItemView_PositionAtTop QAbstractItemView_ScrollHint = 1
	QAbstractItemView_PositionAtBottom QAbstractItemView_ScrollHint = 2
	QAbstractItemView_PositionAtCenter QAbstractItemView_ScrollHint = 3
)
//enum QAbstractItemView_SelectionMode - QAbstractItemView::SelectionMode
type QAbstractItemView_SelectionMode uint32
const (
	QAbstractItemView_NoSelection QAbstractItemView_SelectionMode = 0
	QAbstractItemView_SingleSelection QAbstractItemView_SelectionMode = 1
	QAbstractItemView_MultiSelection QAbstractItemView_SelectionMode = 2
	QAbstractItemView_ExtendedSelection QAbstractItemView_SelectionMode = 3
	QAbstractItemView_ContiguousSelection QAbstractItemView_SelectionMode = 4
)
//enum QAbstractItemView_State - QAbstractItemView::State
type QAbstractItemView_State uint32
const (
	QAbstractItemView_NoState QAbstractItemView_State = 0
	QAbstractItemView_DraggingState QAbstractItemView_State = 1
	QAbstractItemView_DragSelectingState QAbstractItemView_State = 2
	QAbstractItemView_EditingState QAbstractItemView_State = 3
	QAbstractItemView_ExpandingState QAbstractItemView_State = 4
	QAbstractItemView_CollapsingState QAbstractItemView_State = 5
	QAbstractItemView_AnimatingState QAbstractItemView_State = 6
)
//struct QAbstractItemView : QAbstractItemView
type QAbstractItemView struct {
	QAbstractScrollArea
}
func (q *QAbstractItemView) OnPressed(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(196000,196102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemView) OnDoubleClicked(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(196000,196103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemView) OnActivated(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(196000,196104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemView) OnClicked(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(196000,196105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemView) OnEntered(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(196000,196106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemView) OnViewportEntered(fn func()) uintptr {
	var __rv uintptr
	q.Drv(196000,196107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractItemView::alternatingRowColors()
func (q *QAbstractItemView) AlternatingRowColors() bool {
	var __rv bool
	q.Drv(196000,196108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::autoScrollMargin()
func (q *QAbstractItemView) AutoScrollMargin() int {
	var __rv int
	q.Drv(196000,196109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::clearSelection()
func (q *QAbstractItemView) ClearSelection()  {
	q.Drv(196000,196110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::closeEditor(QWidget*,QAbstractItemDelegate::EndEditHint)
func (q *QAbstractItemView) CloseEditor(editor QWidgetInterface,hint QAbstractItemDelegate_EndEditHint)  {
	q.Drv(196000,196111,Native(editor),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::closePersistentEditor(QModelIndex const&)
func (q *QAbstractItemView) ClosePersistentEditor(index *QModelIndex)  {
	q.Drv(196000,196112,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::commitData(QWidget*)
func (q *QAbstractItemView) CommitData(editor QWidgetInterface)  {
	q.Drv(196000,196113,Native(editor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QAbstractItemView) CurrentChanged(current *QModelIndex,previous *QModelIndex)  {
	q.Drv(196000,196114,Native(current),Native(previous),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::currentIndex()
func (q *QAbstractItemView) CurrentIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(196000,196115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemView::dataChanged(QModelIndex const&,QModelIndex const&)
func (q *QAbstractItemView) DataChanged(topLeft *QModelIndex,bottomRight *QModelIndex)  {
	q.Drv(196000,196116,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::defaultDropAction()
func (q *QAbstractItemView) DefaultDropAction() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(196000,196117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::dirtyRegionOffset()
func (q *QAbstractItemView) DirtyRegionOffset() *QPoint {
	var __rv uintptr
	q.Drv(196000,196118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QAbstractItemView::doAutoScroll()
func (q *QAbstractItemView) DoAutoScroll()  {
	q.Drv(196000,196119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::doItemsLayout()
func (q *QAbstractItemView) DoItemsLayout()  {
	q.Drv(196000,196120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::dragDropMode()
func (q *QAbstractItemView) DragDropMode() QAbstractItemView_DragDropMode {
	var __rv QAbstractItemView_DragDropMode
	q.Drv(196000,196121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::dragDropOverwriteMode()
func (q *QAbstractItemView) DragDropOverwriteMode() bool {
	var __rv bool
	q.Drv(196000,196122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::dragEnabled()
func (q *QAbstractItemView) DragEnabled() bool {
	var __rv bool
	q.Drv(196000,196123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::dragEnterEvent(QDragEnterEvent*)
func (q *QAbstractItemView) DragEnterEvent(event *QDragEnterEvent)  {
	q.Drv(196000,196124,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::dragLeaveEvent(QDragLeaveEvent*)
func (q *QAbstractItemView) DragLeaveEvent(event *QDragLeaveEvent)  {
	q.Drv(196000,196125,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::dragMoveEvent(QDragMoveEvent*)
func (q *QAbstractItemView) DragMoveEvent(event *QDragMoveEvent)  {
	q.Drv(196000,196126,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::dropEvent(QDropEvent*)
func (q *QAbstractItemView) DropEvent(event *QDropEvent)  {
	q.Drv(196000,196127,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::dropIndicatorPosition()
func (q *QAbstractItemView) DropIndicatorPosition() QAbstractItemView_DropIndicatorPosition {
	var __rv QAbstractItemView_DropIndicatorPosition
	q.Drv(196000,196128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::edit(QModelIndex const&)
func (q *QAbstractItemView) Edit(index *QModelIndex)  {
	q.Drv(196000,196129,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::edit(QModelIndex const&,QAbstractItemView::EditTrigger,QEvent*)
func (q *QAbstractItemView) EditWithIndexTriggerEvent(index *QModelIndex,trigger QAbstractItemView_EditTrigger,event *QEvent) bool {
	var __rv bool
	q.Drv(196000,196130,Native(index),unsafe.Pointer(&trigger),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::editTriggers()
func (q *QAbstractItemView) EditTriggers() QAbstractItemView_EditTrigger {
	var __rv QAbstractItemView_EditTrigger
	q.Drv(196000,196131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::editorDestroyed(QObject*)
func (q *QAbstractItemView) EditorDestroyed(editor QObjectInterface)  {
	q.Drv(196000,196132,Native(editor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::event(QEvent*)
func (q *QAbstractItemView) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(196000,196133,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::executeDelayedItemsLayout()
func (q *QAbstractItemView) ExecuteDelayedItemsLayout()  {
	q.Drv(196000,196134,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::focusInEvent(QFocusEvent*)
func (q *QAbstractItemView) FocusInEvent(event *QFocusEvent)  {
	q.Drv(196000,196135,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::focusNextPrevChild(bool)
func (q *QAbstractItemView) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(196000,196136,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::focusOutEvent(QFocusEvent*)
func (q *QAbstractItemView) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(196000,196137,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::hasAutoScroll()
func (q *QAbstractItemView) HasAutoScroll() bool {
	var __rv bool
	q.Drv(196000,196138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::horizontalOffset()
func (q *QAbstractItemView) HorizontalOffset() int {
	var __rv int
	q.Drv(196000,196139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::horizontalScrollMode()
func (q *QAbstractItemView) HorizontalScrollMode() QAbstractItemView_ScrollMode {
	var __rv QAbstractItemView_ScrollMode
	q.Drv(196000,196140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::horizontalScrollbarAction(int)
func (q *QAbstractItemView) HorizontalScrollbarAction(action int)  {
	q.Drv(196000,196141,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::horizontalScrollbarValueChanged(int)
func (q *QAbstractItemView) HorizontalScrollbarValueChanged(value int)  {
	q.Drv(196000,196142,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::horizontalStepsPerItem()
func (q *QAbstractItemView) HorizontalStepsPerItem() int {
	var __rv int
	q.Drv(196000,196143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::iconSize()
func (q *QAbstractItemView) IconSize() *QSize {
	var __rv uintptr
	q.Drv(196000,196144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractItemView::indexAt(QPoint const&)
func (q *QAbstractItemView) IndexAt(point *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(196000,196145,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemView::indexWidget(QModelIndex const&)
func (q *QAbstractItemView) IndexWidget(index *QModelIndex) *QWidget {
	var __rv uintptr
	q.Drv(196000,196146,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QAbstractItemView::inputMethodEvent(QInputMethodEvent*)
func (q *QAbstractItemView) InputMethodEvent(event *QInputMethodEvent)  {
	q.Drv(196000,196147,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::inputMethodQuery(Qt::InputMethodQuery)
func (q *QAbstractItemView) InputMethodQuery(query Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(196000,196148,unsafe.Pointer(&query),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractItemView::isIndexHidden(QModelIndex const&)
func (q *QAbstractItemView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(196000,196149,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::itemDelegate()
func (q *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(196000,196150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QAbstractItemView::itemDelegate(QModelIndex const&)
func (q *QAbstractItemView) ItemDelegateWithIndex(index *QModelIndex) *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(196000,196151,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QAbstractItemView::itemDelegateForColumn(int)
func (q *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(196000,196152,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QAbstractItemView::itemDelegateForRow(int)
func (q *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(196000,196153,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QAbstractItemView::keyPressEvent(QKeyEvent*)
func (q *QAbstractItemView) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(196000,196154,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::keyboardSearch(QString const&)
func (q *QAbstractItemView) KeyboardSearch(search string)  {
	q.Drv(196000,196155,unsafe.Pointer(&search),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::model()
func (q *QAbstractItemView) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(196000,196156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QAbstractItemView::mouseDoubleClickEvent(QMouseEvent*)
func (q *QAbstractItemView) MouseDoubleClickEvent(event *QMouseEvent)  {
	q.Drv(196000,196157,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::mouseMoveEvent(QMouseEvent*)
func (q *QAbstractItemView) MouseMoveEvent(event *QMouseEvent)  {
	q.Drv(196000,196158,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::mousePressEvent(QMouseEvent*)
func (q *QAbstractItemView) MousePressEvent(event *QMouseEvent)  {
	q.Drv(196000,196159,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::mouseReleaseEvent(QMouseEvent*)
func (q *QAbstractItemView) MouseReleaseEvent(event *QMouseEvent)  {
	q.Drv(196000,196160,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QAbstractItemView) MoveCursor(cursorAction QAbstractItemView_CursorAction,modifiers Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(196000,196161,unsafe.Pointer(&cursorAction),unsafe.Pointer(&modifiers),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemView::openPersistentEditor(QModelIndex const&)
func (q *QAbstractItemView) OpenPersistentEditor(index *QModelIndex)  {
	q.Drv(196000,196162,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::reset()
func (q *QAbstractItemView) Reset()  {
	q.Drv(196000,196163,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::resizeEvent(QResizeEvent*)
func (q *QAbstractItemView) ResizeEvent(event *QResizeEvent)  {
	q.Drv(196000,196164,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::rootIndex()
func (q *QAbstractItemView) RootIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(196000,196165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemView::rowsAboutToBeRemoved(QModelIndex const&,int,int)
func (q *QAbstractItemView) RowsAboutToBeRemoved(parent *QModelIndex,start int,end int)  {
	q.Drv(196000,196166,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::rowsInserted(QModelIndex const&,int,int)
func (q *QAbstractItemView) RowsInserted(parent *QModelIndex,start int,end int)  {
	q.Drv(196000,196167,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scheduleDelayedItemsLayout()
func (q *QAbstractItemView) ScheduleDelayedItemsLayout()  {
	q.Drv(196000,196168,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scrollDirtyRegion(int,int)
func (q *QAbstractItemView) ScrollDirtyRegion(dx int,dy int)  {
	q.Drv(196000,196169,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scrollTo(QModelIndex const&)
func (q *QAbstractItemView) ScrollTo(index *QModelIndex)  {
	q.Drv(196000,196170,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QAbstractItemView) ScrollToWithIndexHint(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(196000,196171,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scrollToBottom()
func (q *QAbstractItemView) ScrollToBottom()  {
	q.Drv(196000,196172,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::scrollToTop()
func (q *QAbstractItemView) ScrollToTop()  {
	q.Drv(196000,196173,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::selectAll()
func (q *QAbstractItemView) SelectAll()  {
	q.Drv(196000,196174,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::selectedIndexes()
func (q *QAbstractItemView) SelectedIndexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(196000,196175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::selectionBehavior()
func (q *QAbstractItemView) SelectionBehavior() QAbstractItemView_SelectionBehavior {
	var __rv QAbstractItemView_SelectionBehavior
	q.Drv(196000,196176,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::selectionChanged(QItemSelection const&,QItemSelection const&)
func (q *QAbstractItemView) SelectionChanged(selected *QItemSelection,deselected *QItemSelection)  {
	q.Drv(196000,196177,Native(selected),Native(deselected),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::selectionCommand(QModelIndex const&,QEvent const*)
func (q *QAbstractItemView) SelectionCommand(index *QModelIndex,event *QEvent) QItemSelectionModel_SelectionFlag {
	var __rv QItemSelectionModel_SelectionFlag
	q.Drv(196000,196178,Native(index),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::selectionMode()
func (q *QAbstractItemView) SelectionMode() QAbstractItemView_SelectionMode {
	var __rv QAbstractItemView_SelectionMode
	q.Drv(196000,196179,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::selectionModel()
func (q *QAbstractItemView) SelectionModel() *QItemSelectionModel {
	var __rv uintptr
	q.Drv(196000,196180,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelectionModel{}
	_rp.SetDriver(__rv,296000,false)
	return _rp
}	
//QAbstractItemView::setAlternatingRowColors(bool)
func (q *QAbstractItemView) SetAlternatingRowColors(enable bool)  {
	q.Drv(196000,196181,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setAutoScroll(bool)
func (q *QAbstractItemView) SetAutoScroll(enable bool)  {
	q.Drv(196000,196182,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setAutoScrollMargin(int)
func (q *QAbstractItemView) SetAutoScrollMargin(margin int)  {
	q.Drv(196000,196183,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setCurrentIndex(QModelIndex const&)
func (q *QAbstractItemView) SetCurrentIndex(index *QModelIndex)  {
	q.Drv(196000,196184,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDefaultDropAction(Qt::DropAction)
func (q *QAbstractItemView) SetDefaultDropAction(dropAction Qt_DropAction)  {
	q.Drv(196000,196185,unsafe.Pointer(&dropAction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDirtyRegion(QRegion const&)
func (q *QAbstractItemView) SetDirtyRegion(region *QRegion)  {
	q.Drv(196000,196186,Native(region),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDragDropMode(QAbstractItemView::DragDropMode)
func (q *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView_DragDropMode)  {
	q.Drv(196000,196187,unsafe.Pointer(&behavior),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDragDropOverwriteMode(bool)
func (q *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool)  {
	q.Drv(196000,196188,unsafe.Pointer(&overwrite),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDragEnabled(bool)
func (q *QAbstractItemView) SetDragEnabled(enable bool)  {
	q.Drv(196000,196189,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setDropIndicatorShown(bool)
func (q *QAbstractItemView) SetDropIndicatorShown(enable bool)  {
	q.Drv(196000,196190,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setEditTriggers(QFlags<QAbstractItemView::EditTrigger>)
func (q *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView_EditTrigger)  {
	q.Drv(196000,196191,unsafe.Pointer(&triggers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setHorizontalScrollMode(QAbstractItemView::ScrollMode)
func (q *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView_ScrollMode)  {
	q.Drv(196000,196192,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setHorizontalStepsPerItem(int)
func (q *QAbstractItemView) SetHorizontalStepsPerItem(steps int)  {
	q.Drv(196000,196193,unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setIconSize(QSize const&)
func (q *QAbstractItemView) SetIconSize(size *QSize)  {
	q.Drv(196000,196194,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setIndexWidget(QModelIndex const&,QWidget*)
func (q *QAbstractItemView) SetIndexWidget(index *QModelIndex,widget QWidgetInterface)  {
	q.Drv(196000,196195,Native(index),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setItemDelegate(QAbstractItemDelegate*)
func (q *QAbstractItemView) SetItemDelegate(delegate *QAbstractItemDelegate)  {
	q.Drv(196000,196196,Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setItemDelegateForColumn(int,QAbstractItemDelegate*)
func (q *QAbstractItemView) SetItemDelegateForColumn(column int,delegate *QAbstractItemDelegate)  {
	q.Drv(196000,196197,unsafe.Pointer(&column),Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setItemDelegateForRow(int,QAbstractItemDelegate*)
func (q *QAbstractItemView) SetItemDelegateForRow(row int,delegate *QAbstractItemDelegate)  {
	q.Drv(196000,196198,unsafe.Pointer(&row),Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setModel(QAbstractItemModel*)
func (q *QAbstractItemView) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(196000,196199,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setRootIndex(QModelIndex const&)
func (q *QAbstractItemView) SetRootIndex(index *QModelIndex)  {
	q.Drv(196000,196200,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QAbstractItemView) SetSelection(rect *QRect,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(196000,196201,Native(rect),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setSelectionBehavior(QAbstractItemView::SelectionBehavior)
func (q *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView_SelectionBehavior)  {
	q.Drv(196000,196202,unsafe.Pointer(&behavior),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setSelectionMode(QAbstractItemView::SelectionMode)
func (q *QAbstractItemView) SetSelectionMode(mode QAbstractItemView_SelectionMode)  {
	q.Drv(196000,196203,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setSelectionModel(QItemSelectionModel*)
func (q *QAbstractItemView) SetSelectionModel(selectionModel *QItemSelectionModel)  {
	q.Drv(196000,196204,Native(selectionModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setState(QAbstractItemView::State)
func (q *QAbstractItemView) SetState(state QAbstractItemView_State)  {
	q.Drv(196000,196205,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setTabKeyNavigation(bool)
func (q *QAbstractItemView) SetTabKeyNavigation(enable bool)  {
	q.Drv(196000,196206,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setTextElideMode(Qt::TextElideMode)
func (q *QAbstractItemView) SetTextElideMode(mode Qt_TextElideMode)  {
	q.Drv(196000,196207,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setVerticalScrollMode(QAbstractItemView::ScrollMode)
func (q *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView_ScrollMode)  {
	q.Drv(196000,196208,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::setVerticalStepsPerItem(int)
func (q *QAbstractItemView) SetVerticalStepsPerItem(steps int)  {
	q.Drv(196000,196209,unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::showDropIndicator()
func (q *QAbstractItemView) ShowDropIndicator() bool {
	var __rv bool
	q.Drv(196000,196210,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::sizeHintForColumn(int)
func (q *QAbstractItemView) SizeHintForColumn(column int) int {
	var __rv int
	q.Drv(196000,196211,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::sizeHintForIndex(QModelIndex const&)
func (q *QAbstractItemView) SizeHintForIndex(index *QModelIndex) *QSize {
	var __rv uintptr
	q.Drv(196000,196212,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractItemView::sizeHintForRow(int)
func (q *QAbstractItemView) SizeHintForRow(row int) int {
	var __rv int
	q.Drv(196000,196213,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::startAutoScroll()
func (q *QAbstractItemView) StartAutoScroll()  {
	q.Drv(196000,196214,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::startDrag(QFlags<Qt::DropAction>)
func (q *QAbstractItemView) StartDrag(supportedActions Qt_DropAction)  {
	q.Drv(196000,196215,unsafe.Pointer(&supportedActions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::state()
func (q *QAbstractItemView) State() QAbstractItemView_State {
	var __rv QAbstractItemView_State
	q.Drv(196000,196216,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::stopAutoScroll()
func (q *QAbstractItemView) StopAutoScroll()  {
	q.Drv(196000,196217,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::tabKeyNavigation()
func (q *QAbstractItemView) TabKeyNavigation() bool {
	var __rv bool
	q.Drv(196000,196218,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::textElideMode()
func (q *QAbstractItemView) TextElideMode() Qt_TextElideMode {
	var __rv Qt_TextElideMode
	q.Drv(196000,196219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::timerEvent(QTimerEvent*)
func (q *QAbstractItemView) TimerEvent(event *QTimerEvent)  {
	q.Drv(196000,196220,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::update()
func (q *QAbstractItemView) Update()  {
	q.Drv(196000,196221,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::update(QModelIndex const&)
func (q *QAbstractItemView) UpdateWithIndex(index *QModelIndex)  {
	q.Drv(196000,196222,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::updateEditorData()
func (q *QAbstractItemView) UpdateEditorData()  {
	q.Drv(196000,196223,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::updateEditorGeometries()
func (q *QAbstractItemView) UpdateEditorGeometries()  {
	q.Drv(196000,196224,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::updateGeometries()
func (q *QAbstractItemView) UpdateGeometries()  {
	q.Drv(196000,196225,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::verticalOffset()
func (q *QAbstractItemView) VerticalOffset() int {
	var __rv int
	q.Drv(196000,196226,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::verticalScrollMode()
func (q *QAbstractItemView) VerticalScrollMode() QAbstractItemView_ScrollMode {
	var __rv QAbstractItemView_ScrollMode
	q.Drv(196000,196227,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::verticalScrollbarAction(int)
func (q *QAbstractItemView) VerticalScrollbarAction(action int)  {
	q.Drv(196000,196228,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::verticalScrollbarValueChanged(int)
func (q *QAbstractItemView) VerticalScrollbarValueChanged(value int)  {
	q.Drv(196000,196229,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemView::verticalStepsPerItem()
func (q *QAbstractItemView) VerticalStepsPerItem() int {
	var __rv int
	q.Drv(196000,196230,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::viewportEvent(QEvent*)
func (q *QAbstractItemView) ViewportEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(196000,196231,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemView::visualRect(QModelIndex const&)
func (q *QAbstractItemView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(196000,196232,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QAbstractItemView::visualRegionForSelection(QItemSelection const&)
func (q *QAbstractItemView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(196000,196233,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
