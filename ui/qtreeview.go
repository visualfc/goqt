// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTreeView : QTreeView
type QTreeView struct {
	QAbstractItemView
}
func (q *QTreeView) OnCollapsed(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(385000,385102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeView) OnExpanded(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(385000,385103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTreeView::QTreeView()	
func NewTreeView() *QTreeView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),385000,385104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeView{}
	_p.SetDriver(__rv,385000,false)
	return _p
} 
//QTreeView::QTreeView(QWidget*)	
func NewTreeViewWithParent(parent QWidgetInterface) *QTreeView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),385000,385105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeView{}
	_p.SetDriver(__rv,385000,false)
	return _p
} 
//QTreeView::allColumnsShowFocus()
func (q *QTreeView) AllColumnsShowFocus() bool {
	var __rv bool
	q.Drv(385000,385106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::autoExpandDelay()
func (q *QTreeView) AutoExpandDelay() int {
	var __rv int
	q.Drv(385000,385107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::collapse(QModelIndex const&)
func (q *QTreeView) Collapse(index *QModelIndex)  {
	q.Drv(385000,385108,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::collapseAll()
func (q *QTreeView) CollapseAll()  {
	q.Drv(385000,385109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::columnAt(int)
func (q *QTreeView) ColumnAt(x int) int {
	var __rv int
	q.Drv(385000,385110,unsafe.Pointer(&x),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::columnCountChanged(int,int)
func (q *QTreeView) ColumnCountChanged(oldCount int,newCount int)  {
	q.Drv(385000,385111,unsafe.Pointer(&oldCount),unsafe.Pointer(&newCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::columnMoved()
func (q *QTreeView) ColumnMoved()  {
	q.Drv(385000,385112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::columnResized(int,int,int)
func (q *QTreeView) ColumnResized(column int,oldSize int,newSize int)  {
	q.Drv(385000,385113,unsafe.Pointer(&column),unsafe.Pointer(&oldSize),unsafe.Pointer(&newSize),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::columnViewportPosition(int)
func (q *QTreeView) ColumnViewportPosition(column int) int {
	var __rv int
	q.Drv(385000,385114,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::columnWidth(int)
func (q *QTreeView) ColumnWidth(column int) int {
	var __rv int
	q.Drv(385000,385115,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QTreeView) CurrentChanged(current *QModelIndex,previous *QModelIndex)  {
	q.Drv(385000,385116,Native(current),Native(previous),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::dataChanged(QModelIndex const&,QModelIndex const&)
func (q *QTreeView) DataChanged(topLeft *QModelIndex,bottomRight *QModelIndex)  {
	q.Drv(385000,385117,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::doItemsLayout()
func (q *QTreeView) DoItemsLayout()  {
	q.Drv(385000,385118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::dragMoveEvent(QDragMoveEvent*)
func (q *QTreeView) DragMoveEvent(event *QDragMoveEvent)  {
	q.Drv(385000,385119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::drawBranches(QPainter*,QRect const&,QModelIndex const&)
func (q *QTreeView) DrawBranches(painter *QPainter,rect *QRect,index *QModelIndex)  {
	q.Drv(385000,385120,Native(painter),Native(rect),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::drawTree(QPainter*,QRegion const&)
func (q *QTreeView) DrawTree(painter *QPainter,region *QRegion)  {
	q.Drv(385000,385121,Native(painter),Native(region),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::expand(QModelIndex const&)
func (q *QTreeView) Expand(index *QModelIndex)  {
	q.Drv(385000,385122,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::expandAll()
func (q *QTreeView) ExpandAll()  {
	q.Drv(385000,385123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::expandToDepth(int)
func (q *QTreeView) ExpandToDepth(depth int)  {
	q.Drv(385000,385124,unsafe.Pointer(&depth),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::expandsOnDoubleClick()
func (q *QTreeView) ExpandsOnDoubleClick() bool {
	var __rv bool
	q.Drv(385000,385125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::header()
func (q *QTreeView) Header() *QHeaderView {
	var __rv uintptr
	q.Drv(385000,385126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QHeaderView{}
	_rp.SetDriver(__rv,290000,false)
	return _rp
}	
//QTreeView::hideColumn(int)
func (q *QTreeView) HideColumn(column int)  {
	q.Drv(385000,385127,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::horizontalOffset()
func (q *QTreeView) HorizontalOffset() int {
	var __rv int
	q.Drv(385000,385128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::horizontalScrollbarAction(int)
func (q *QTreeView) HorizontalScrollbarAction(action int)  {
	q.Drv(385000,385129,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::indentation()
func (q *QTreeView) Indentation() int {
	var __rv int
	q.Drv(385000,385130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::indexAbove(QModelIndex const&)
func (q *QTreeView) IndexAbove(index *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(385000,385131,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTreeView::indexAt(QPoint const&)
func (q *QTreeView) IndexAt(p *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(385000,385132,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTreeView::indexBelow(QModelIndex const&)
func (q *QTreeView) IndexBelow(index *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(385000,385133,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTreeView::indexRowSizeHint(QModelIndex const&)
func (q *QTreeView) IndexRowSizeHint(index *QModelIndex) int {
	var __rv int
	q.Drv(385000,385134,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isAnimated()
func (q *QTreeView) IsAnimated() bool {
	var __rv bool
	q.Drv(385000,385135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isColumnHidden(int)
func (q *QTreeView) IsColumnHidden(column int) bool {
	var __rv bool
	q.Drv(385000,385136,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isExpanded(QModelIndex const&)
func (q *QTreeView) IsExpanded(index *QModelIndex) bool {
	var __rv bool
	q.Drv(385000,385137,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isFirstColumnSpanned(int,QModelIndex const&)
func (q *QTreeView) IsFirstColumnSpanned(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(385000,385138,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isHeaderHidden()
func (q *QTreeView) IsHeaderHidden() bool {
	var __rv bool
	q.Drv(385000,385139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isIndexHidden(QModelIndex const&)
func (q *QTreeView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(385000,385140,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isRowHidden(int,QModelIndex const&)
func (q *QTreeView) IsRowHidden(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(385000,385141,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::isSortingEnabled()
func (q *QTreeView) IsSortingEnabled() bool {
	var __rv bool
	q.Drv(385000,385142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::itemsExpandable()
func (q *QTreeView) ItemsExpandable() bool {
	var __rv bool
	q.Drv(385000,385143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::keyPressEvent(QKeyEvent*)
func (q *QTreeView) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(385000,385144,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::keyboardSearch(QString const&)
func (q *QTreeView) KeyboardSearch(search string)  {
	q.Drv(385000,385145,unsafe.Pointer(&search),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::mouseDoubleClickEvent(QMouseEvent*)
func (q *QTreeView) MouseDoubleClickEvent(event *QMouseEvent)  {
	q.Drv(385000,385146,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::mouseMoveEvent(QMouseEvent*)
func (q *QTreeView) MouseMoveEvent(event *QMouseEvent)  {
	q.Drv(385000,385147,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::mousePressEvent(QMouseEvent*)
func (q *QTreeView) MousePressEvent(event *QMouseEvent)  {
	q.Drv(385000,385148,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::mouseReleaseEvent(QMouseEvent*)
func (q *QTreeView) MouseReleaseEvent(event *QMouseEvent)  {
	q.Drv(385000,385149,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QTreeView) MoveCursor(cursorAction QAbstractItemView_CursorAction,modifiers Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(385000,385150,unsafe.Pointer(&cursorAction),unsafe.Pointer(&modifiers),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTreeView::paintEvent(QPaintEvent*)
func (q *QTreeView) PaintEvent(event *QPaintEvent)  {
	q.Drv(385000,385151,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::reexpand()
func (q *QTreeView) Reexpand()  {
	q.Drv(385000,385152,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::reset()
func (q *QTreeView) Reset()  {
	q.Drv(385000,385153,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::resizeColumnToContents(int)
func (q *QTreeView) ResizeColumnToContents(column int)  {
	q.Drv(385000,385154,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::rootIsDecorated()
func (q *QTreeView) RootIsDecorated() bool {
	var __rv bool
	q.Drv(385000,385155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::rowHeight(QModelIndex const&)
func (q *QTreeView) RowHeight(index *QModelIndex) int {
	var __rv int
	q.Drv(385000,385156,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::rowsAboutToBeRemoved(QModelIndex const&,int,int)
func (q *QTreeView) RowsAboutToBeRemoved(parent *QModelIndex,start int,end int)  {
	q.Drv(385000,385157,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::rowsInserted(QModelIndex const&,int,int)
func (q *QTreeView) RowsInserted(parent *QModelIndex,start int,end int)  {
	q.Drv(385000,385158,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::rowsRemoved(QModelIndex const&,int,int)
func (q *QTreeView) RowsRemoved(parent *QModelIndex,first int,last int)  {
	q.Drv(385000,385159,Native(parent),unsafe.Pointer(&first),unsafe.Pointer(&last),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::scrollContentsBy(int,int)
func (q *QTreeView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(385000,385160,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::scrollTo(QModelIndex const&)
func (q *QTreeView) ScrollTo(index *QModelIndex)  {
	q.Drv(385000,385161,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QTreeView) ScrollToWithIndexHint(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(385000,385162,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::selectAll()
func (q *QTreeView) SelectAll()  {
	q.Drv(385000,385163,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::selectedIndexes()
func (q *QTreeView) SelectedIndexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(385000,385164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::selectionChanged(QItemSelection const&,QItemSelection const&)
func (q *QTreeView) SelectionChanged(selected *QItemSelection,deselected *QItemSelection)  {
	q.Drv(385000,385165,Native(selected),Native(deselected),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setAllColumnsShowFocus(bool)
func (q *QTreeView) SetAllColumnsShowFocus(enable bool)  {
	q.Drv(385000,385166,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setAnimated(bool)
func (q *QTreeView) SetAnimated(enable bool)  {
	q.Drv(385000,385167,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setAutoExpandDelay(int)
func (q *QTreeView) SetAutoExpandDelay(delay int)  {
	q.Drv(385000,385168,unsafe.Pointer(&delay),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setColumnHidden(int,bool)
func (q *QTreeView) SetColumnHidden(column int,hide bool)  {
	q.Drv(385000,385169,unsafe.Pointer(&column),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setColumnWidth(int,int)
func (q *QTreeView) SetColumnWidth(column int,width int)  {
	q.Drv(385000,385170,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setExpanded(QModelIndex const&,bool)
func (q *QTreeView) SetExpanded(index *QModelIndex,expand bool)  {
	q.Drv(385000,385171,Native(index),unsafe.Pointer(&expand),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setExpandsOnDoubleClick(bool)
func (q *QTreeView) SetExpandsOnDoubleClick(enable bool)  {
	q.Drv(385000,385172,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setFirstColumnSpanned(int,QModelIndex const&,bool)
func (q *QTreeView) SetFirstColumnSpanned(row int,parent *QModelIndex,span bool)  {
	q.Drv(385000,385173,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&span),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setHeader(QHeaderView*)
func (q *QTreeView) SetHeader(header *QHeaderView)  {
	q.Drv(385000,385174,Native(header),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setHeaderHidden(bool)
func (q *QTreeView) SetHeaderHidden(hide bool)  {
	q.Drv(385000,385175,unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setIndentation(int)
func (q *QTreeView) SetIndentation(i int)  {
	q.Drv(385000,385176,unsafe.Pointer(&i),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setItemsExpandable(bool)
func (q *QTreeView) SetItemsExpandable(enable bool)  {
	q.Drv(385000,385177,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setModel(QAbstractItemModel*)
func (q *QTreeView) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(385000,385178,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setRootIndex(QModelIndex const&)
func (q *QTreeView) SetRootIndex(index *QModelIndex)  {
	q.Drv(385000,385179,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setRootIsDecorated(bool)
func (q *QTreeView) SetRootIsDecorated(show bool)  {
	q.Drv(385000,385180,unsafe.Pointer(&show),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setRowHidden(int,QModelIndex const&,bool)
func (q *QTreeView) SetRowHidden(row int,parent *QModelIndex,hide bool)  {
	q.Drv(385000,385181,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QTreeView) SetSelection(rect *QRect,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(385000,385182,Native(rect),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setSelectionModel(QItemSelectionModel*)
func (q *QTreeView) SetSelectionModel(selectionModel *QItemSelectionModel)  {
	q.Drv(385000,385183,Native(selectionModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setSortingEnabled(bool)
func (q *QTreeView) SetSortingEnabled(enable bool)  {
	q.Drv(385000,385184,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setUniformRowHeights(bool)
func (q *QTreeView) SetUniformRowHeights(uniform bool)  {
	q.Drv(385000,385185,unsafe.Pointer(&uniform),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::setWordWrap(bool)
func (q *QTreeView) SetWordWrap(on bool)  {
	q.Drv(385000,385186,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::showColumn(int)
func (q *QTreeView) ShowColumn(column int)  {
	q.Drv(385000,385187,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::sizeHintForColumn(int)
func (q *QTreeView) SizeHintForColumn(column int) int {
	var __rv int
	q.Drv(385000,385188,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::sortByColumn(int)
func (q *QTreeView) SortByColumn(column int)  {
	q.Drv(385000,385189,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::sortByColumn(int,Qt::SortOrder)
func (q *QTreeView) SortByColumnWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(385000,385190,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::timerEvent(QTimerEvent*)
func (q *QTreeView) TimerEvent(event *QTimerEvent)  {
	q.Drv(385000,385191,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::uniformRowHeights()
func (q *QTreeView) UniformRowHeights() bool {
	var __rv bool
	q.Drv(385000,385192,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::updateGeometries()
func (q *QTreeView) UpdateGeometries()  {
	q.Drv(385000,385193,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeView::verticalOffset()
func (q *QTreeView) VerticalOffset() int {
	var __rv int
	q.Drv(385000,385194,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::viewportEvent(QEvent*)
func (q *QTreeView) ViewportEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(385000,385195,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeView::visualRect(QModelIndex const&)
func (q *QTreeView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(385000,385196,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTreeView::visualRegionForSelection(QItemSelection const&)
func (q *QTreeView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(385000,385197,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QTreeView::wordWrap()
func (q *QTreeView) WordWrap() bool {
	var __rv bool
	q.Drv(385000,385198,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
