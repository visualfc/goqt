// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QColumnView : QColumnView
type QColumnView struct {
	QAbstractItemView
}
func (q *QColumnView) OnUpdatePreviewWidget(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(216000,216102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QColumnView::QColumnView()	
func NewColumnView() *QColumnView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),216000,216103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColumnView{}
	_p.SetDriver(__rv,216000,false)
	return _p
} 
//QColumnView::QColumnView(QWidget*)	
func NewColumnViewWithParent(parent QWidgetInterface) *QColumnView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),216000,216104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColumnView{}
	_p.SetDriver(__rv,216000,false)
	return _p
} 
//QColumnView::columnWidths()
func (q *QColumnView) ColumnWidths() []int {
	var __rv []int
	q.Drv(216000,216105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColumnView::createColumn(QModelIndex const&)
func (q *QColumnView) CreateColumn(rootIndex *QModelIndex) *QAbstractItemView {
	var __rv uintptr
	q.Drv(216000,216106,Native(rootIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemView{}
	_rp.SetDriver(__rv,196000,false)
	return _rp
}	
//QColumnView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QColumnView) CurrentChanged(current *QModelIndex,previous *QModelIndex)  {
	q.Drv(216000,216107,Native(current),Native(previous),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::horizontalOffset()
func (q *QColumnView) HorizontalOffset() int {
	var __rv int
	q.Drv(216000,216108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColumnView::indexAt(QPoint const&)
func (q *QColumnView) IndexAt(point *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(216000,216109,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QColumnView::initializeColumn(QAbstractItemView*)
func (q *QColumnView) InitializeColumn(column *QAbstractItemView)  {
	q.Drv(216000,216110,Native(column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::isIndexHidden(QModelIndex const&)
func (q *QColumnView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(216000,216111,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColumnView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QColumnView) MoveCursor(cursorAction QAbstractItemView_CursorAction,modifiers Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(216000,216112,unsafe.Pointer(&cursorAction),unsafe.Pointer(&modifiers),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QColumnView::previewWidget()
func (q *QColumnView) PreviewWidget() *QWidget {
	var __rv uintptr
	q.Drv(216000,216113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QColumnView::resizeEvent(QResizeEvent*)
func (q *QColumnView) ResizeEvent(event *QResizeEvent)  {
	q.Drv(216000,216114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::resizeGripsVisible()
func (q *QColumnView) ResizeGripsVisible() bool {
	var __rv bool
	q.Drv(216000,216115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColumnView::rowsInserted(QModelIndex const&,int,int)
func (q *QColumnView) RowsInserted(parent *QModelIndex,start int,end int)  {
	q.Drv(216000,216116,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::scrollContentsBy(int,int)
func (q *QColumnView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(216000,216117,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::scrollTo(QModelIndex const&)
func (q *QColumnView) ScrollTo(index *QModelIndex)  {
	q.Drv(216000,216118,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QColumnView) ScrollToWithIndexHint(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(216000,216119,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::selectAll()
func (q *QColumnView) SelectAll()  {
	q.Drv(216000,216120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setColumnWidths(QList<int> const&)
func (q *QColumnView) SetColumnWidths(list []int)  {
	q.Drv(216000,216121,unsafe.Pointer(&list),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setModel(QAbstractItemModel*)
func (q *QColumnView) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(216000,216122,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setPreviewWidget(QWidget*)
func (q *QColumnView) SetPreviewWidget(widget QWidgetInterface)  {
	q.Drv(216000,216123,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setResizeGripsVisible(bool)
func (q *QColumnView) SetResizeGripsVisible(visible bool)  {
	q.Drv(216000,216124,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setRootIndex(QModelIndex const&)
func (q *QColumnView) SetRootIndex(index *QModelIndex)  {
	q.Drv(216000,216125,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QColumnView) SetSelection(rect *QRect,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(216000,216126,Native(rect),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::setSelectionModel(QItemSelectionModel*)
func (q *QColumnView) SetSelectionModel(selectionModel *QItemSelectionModel)  {
	q.Drv(216000,216127,Native(selectionModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColumnView::sizeHint()
func (q *QColumnView) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(216000,216128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QColumnView::verticalOffset()
func (q *QColumnView) VerticalOffset() int {
	var __rv int
	q.Drv(216000,216129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColumnView::visualRect(QModelIndex const&)
func (q *QColumnView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(216000,216130,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QColumnView::visualRegionForSelection(QItemSelection const&)
func (q *QColumnView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(216000,216131,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
