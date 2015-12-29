// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QListView_ResizeMode - QListView::ResizeMode
type QListView_ResizeMode uint32
const (
	QListView_Fixed QListView_ResizeMode = 0
	QListView_Adjust QListView_ResizeMode = 1
)
//enum QListView_LayoutMode - QListView::LayoutMode
type QListView_LayoutMode uint32
const (
	QListView_SinglePass QListView_LayoutMode = 0
	QListView_Batched QListView_LayoutMode = 1
)
//enum QListView_ViewMode - QListView::ViewMode
type QListView_ViewMode uint32
const (
	QListView_ListMode QListView_ViewMode = 0
	QListView_IconMode QListView_ViewMode = 1
)
//enum QListView_Flow - QListView::Flow
type QListView_Flow uint32
const (
	QListView_LeftToRight QListView_Flow = 0
	QListView_TopToBottom QListView_Flow = 1
)
//enum QListView_Movement - QListView::Movement
type QListView_Movement uint32
const (
	QListView_Static QListView_Movement = 0
	QListView_Free QListView_Movement = 1
	QListView_Snap QListView_Movement = 2
)
//struct QListView : QListView
type QListView struct {
	QAbstractItemView
}
func (q *QListView) OnIndexesMoved(fn func([]*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(303000,303102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QListView::QListView()	
func NewListView() *QListView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),303000,303103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListView{}
	_p.SetDriver(__rv,303000,false)
	return _p
} 
//QListView::QListView(QWidget*)	
func NewListViewWithParent(parent QWidgetInterface) *QListView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),303000,303104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListView{}
	_p.SetDriver(__rv,303000,false)
	return _p
} 
//QListView::batchSize()
func (q *QListView) BatchSize() int {
	var __rv int
	q.Drv(303000,303105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::clearPropertyFlags()
func (q *QListView) ClearPropertyFlags()  {
	q.Drv(303000,303106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::contentsSize()
func (q *QListView) ContentsSize() *QSize {
	var __rv uintptr
	q.Drv(303000,303107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QListView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QListView) CurrentChanged(current *QModelIndex,previous *QModelIndex)  {
	q.Drv(303000,303108,Native(current),Native(previous),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::dataChanged(QModelIndex const&,QModelIndex const&)
func (q *QListView) DataChanged(topLeft *QModelIndex,bottomRight *QModelIndex)  {
	q.Drv(303000,303109,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::doItemsLayout()
func (q *QListView) DoItemsLayout()  {
	q.Drv(303000,303110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::dragLeaveEvent(QDragLeaveEvent*)
func (q *QListView) DragLeaveEvent(e *QDragLeaveEvent)  {
	q.Drv(303000,303111,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::dragMoveEvent(QDragMoveEvent*)
func (q *QListView) DragMoveEvent(e *QDragMoveEvent)  {
	q.Drv(303000,303112,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::dropEvent(QDropEvent*)
func (q *QListView) DropEvent(e *QDropEvent)  {
	q.Drv(303000,303113,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::event(QEvent*)
func (q *QListView) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(303000,303114,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::flow()
func (q *QListView) Flow() QListView_Flow {
	var __rv QListView_Flow
	q.Drv(303000,303115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::gridSize()
func (q *QListView) GridSize() *QSize {
	var __rv uintptr
	q.Drv(303000,303116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QListView::horizontalOffset()
func (q *QListView) HorizontalOffset() int {
	var __rv int
	q.Drv(303000,303117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::indexAt(QPoint const&)
func (q *QListView) IndexAt(p *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(303000,303118,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QListView::isIndexHidden(QModelIndex const&)
func (q *QListView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(303000,303119,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::isRowHidden(int)
func (q *QListView) IsRowHidden(row int) bool {
	var __rv bool
	q.Drv(303000,303120,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::isSelectionRectVisible()
func (q *QListView) IsSelectionRectVisible() bool {
	var __rv bool
	q.Drv(303000,303121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::isWrapping()
func (q *QListView) IsWrapping() bool {
	var __rv bool
	q.Drv(303000,303122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::layoutMode()
func (q *QListView) LayoutMode() QListView_LayoutMode {
	var __rv QListView_LayoutMode
	q.Drv(303000,303123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::modelColumn()
func (q *QListView) ModelColumn() int {
	var __rv int
	q.Drv(303000,303124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::mouseMoveEvent(QMouseEvent*)
func (q *QListView) MouseMoveEvent(e *QMouseEvent)  {
	q.Drv(303000,303125,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::mouseReleaseEvent(QMouseEvent*)
func (q *QListView) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(303000,303126,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QListView) MoveCursor(cursorAction QAbstractItemView_CursorAction,modifiers Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(303000,303127,unsafe.Pointer(&cursorAction),unsafe.Pointer(&modifiers),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QListView::movement()
func (q *QListView) Movement() QListView_Movement {
	var __rv QListView_Movement
	q.Drv(303000,303128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::paintEvent(QPaintEvent*)
func (q *QListView) PaintEvent(e *QPaintEvent)  {
	q.Drv(303000,303129,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::rectForIndex(QModelIndex const&)
func (q *QListView) RectForIndex(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(303000,303130,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QListView::reset()
func (q *QListView) Reset()  {
	q.Drv(303000,303131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::resizeContents(int,int)
func (q *QListView) ResizeContents(width int,height int)  {
	q.Drv(303000,303132,unsafe.Pointer(&width),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::resizeEvent(QResizeEvent*)
func (q *QListView) ResizeEvent(e *QResizeEvent)  {
	q.Drv(303000,303133,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::resizeMode()
func (q *QListView) ResizeMode() QListView_ResizeMode {
	var __rv QListView_ResizeMode
	q.Drv(303000,303134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::rowsAboutToBeRemoved(QModelIndex const&,int,int)
func (q *QListView) RowsAboutToBeRemoved(parent *QModelIndex,start int,end int)  {
	q.Drv(303000,303135,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::rowsInserted(QModelIndex const&,int,int)
func (q *QListView) RowsInserted(parent *QModelIndex,start int,end int)  {
	q.Drv(303000,303136,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::scrollContentsBy(int,int)
func (q *QListView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(303000,303137,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::scrollTo(QModelIndex const&)
func (q *QListView) ScrollTo(index *QModelIndex)  {
	q.Drv(303000,303138,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QListView) ScrollToWithIndexHint(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(303000,303139,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::selectedIndexes()
func (q *QListView) SelectedIndexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(303000,303140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::selectionChanged(QItemSelection const&,QItemSelection const&)
func (q *QListView) SelectionChanged(selected *QItemSelection,deselected *QItemSelection)  {
	q.Drv(303000,303141,Native(selected),Native(deselected),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setBatchSize(int)
func (q *QListView) SetBatchSize(batchSize int)  {
	q.Drv(303000,303142,unsafe.Pointer(&batchSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setFlow(QListView::Flow)
func (q *QListView) SetFlow(flow QListView_Flow)  {
	q.Drv(303000,303143,unsafe.Pointer(&flow),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setGridSize(QSize const&)
func (q *QListView) SetGridSize(size *QSize)  {
	q.Drv(303000,303144,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setLayoutMode(QListView::LayoutMode)
func (q *QListView) SetLayoutMode(mode QListView_LayoutMode)  {
	q.Drv(303000,303145,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setModelColumn(int)
func (q *QListView) SetModelColumn(column int)  {
	q.Drv(303000,303146,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setMovement(QListView::Movement)
func (q *QListView) SetMovement(movement QListView_Movement)  {
	q.Drv(303000,303147,unsafe.Pointer(&movement),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setPositionForIndex(QPoint const&,QModelIndex const&)
func (q *QListView) SetPositionForIndex(position *QPoint,index *QModelIndex)  {
	q.Drv(303000,303148,Native(position),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setResizeMode(QListView::ResizeMode)
func (q *QListView) SetResizeMode(mode QListView_ResizeMode)  {
	q.Drv(303000,303149,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setRootIndex(QModelIndex const&)
func (q *QListView) SetRootIndex(index *QModelIndex)  {
	q.Drv(303000,303150,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setRowHidden(int,bool)
func (q *QListView) SetRowHidden(row int,hide bool)  {
	q.Drv(303000,303151,unsafe.Pointer(&row),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QListView) SetSelection(rect *QRect,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(303000,303152,Native(rect),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setSelectionRectVisible(bool)
func (q *QListView) SetSelectionRectVisible(show bool)  {
	q.Drv(303000,303153,unsafe.Pointer(&show),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setSpacing(int)
func (q *QListView) SetSpacing(space int)  {
	q.Drv(303000,303154,unsafe.Pointer(&space),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setUniformItemSizes(bool)
func (q *QListView) SetUniformItemSizes(enable bool)  {
	q.Drv(303000,303155,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setViewMode(QListView::ViewMode)
func (q *QListView) SetViewMode(mode QListView_ViewMode)  {
	q.Drv(303000,303156,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setWordWrap(bool)
func (q *QListView) SetWordWrap(on bool)  {
	q.Drv(303000,303157,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::setWrapping(bool)
func (q *QListView) SetWrapping(enable bool)  {
	q.Drv(303000,303158,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::spacing()
func (q *QListView) Spacing() int {
	var __rv int
	q.Drv(303000,303159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::startDrag(QFlags<Qt::DropAction>)
func (q *QListView) StartDrag(supportedActions Qt_DropAction)  {
	q.Drv(303000,303160,unsafe.Pointer(&supportedActions),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::timerEvent(QTimerEvent*)
func (q *QListView) TimerEvent(e *QTimerEvent)  {
	q.Drv(303000,303161,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::uniformItemSizes()
func (q *QListView) UniformItemSizes() bool {
	var __rv bool
	q.Drv(303000,303162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::updateGeometries()
func (q *QListView) UpdateGeometries()  {
	q.Drv(303000,303163,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListView::verticalOffset()
func (q *QListView) VerticalOffset() int {
	var __rv int
	q.Drv(303000,303164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::viewMode()
func (q *QListView) ViewMode() QListView_ViewMode {
	var __rv QListView_ViewMode
	q.Drv(303000,303165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListView::visualRect(QModelIndex const&)
func (q *QListView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(303000,303166,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QListView::visualRegionForSelection(QItemSelection const&)
func (q *QListView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(303000,303167,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QListView::wordWrap()
func (q *QListView) WordWrap() bool {
	var __rv bool
	q.Drv(303000,303168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
