// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTableView : QTableView
type QTableView struct {
	QAbstractItemView
}
//QTableView::QTableView()	
func NewTableView() *QTableView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),365000,365102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableView{}
	_p.SetDriver(__rv,365000,false)
	return _p
} 
//QTableView::QTableView(QWidget*)	
func NewTableViewWithParent(parent QWidgetInterface) *QTableView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),365000,365103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableView{}
	_p.SetDriver(__rv,365000,false)
	return _p
} 
//QTableView::clearSpans()
func (q *QTableView) ClearSpans()  {
	q.Drv(365000,365104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::columnAt(int)
func (q *QTableView) ColumnAt(x int) int {
	var __rv int
	q.Drv(365000,365105,unsafe.Pointer(&x),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::columnCountChanged(int,int)
func (q *QTableView) ColumnCountChanged(oldCount int,newCount int)  {
	q.Drv(365000,365106,unsafe.Pointer(&oldCount),unsafe.Pointer(&newCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::columnMoved(int,int,int)
func (q *QTableView) ColumnMoved(column int,oldIndex int,newIndex int)  {
	q.Drv(365000,365107,unsafe.Pointer(&column),unsafe.Pointer(&oldIndex),unsafe.Pointer(&newIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::columnResized(int,int,int)
func (q *QTableView) ColumnResized(column int,oldWidth int,newWidth int)  {
	q.Drv(365000,365108,unsafe.Pointer(&column),unsafe.Pointer(&oldWidth),unsafe.Pointer(&newWidth),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::columnSpan(int,int)
func (q *QTableView) ColumnSpan(row int,column int) int {
	var __rv int
	q.Drv(365000,365109,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::columnViewportPosition(int)
func (q *QTableView) ColumnViewportPosition(column int) int {
	var __rv int
	q.Drv(365000,365110,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::columnWidth(int)
func (q *QTableView) ColumnWidth(column int) int {
	var __rv int
	q.Drv(365000,365111,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QTableView) CurrentChanged(current *QModelIndex,previous *QModelIndex)  {
	q.Drv(365000,365112,Native(current),Native(previous),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::gridStyle()
func (q *QTableView) GridStyle() Qt_PenStyle {
	var __rv Qt_PenStyle
	q.Drv(365000,365113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::hideColumn(int)
func (q *QTableView) HideColumn(column int)  {
	q.Drv(365000,365114,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::hideRow(int)
func (q *QTableView) HideRow(row int)  {
	q.Drv(365000,365115,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::horizontalHeader()
func (q *QTableView) HorizontalHeader() *QHeaderView {
	var __rv uintptr
	q.Drv(365000,365116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QHeaderView{}
	_rp.SetDriver(__rv,290000,false)
	return _rp
}	
//QTableView::horizontalOffset()
func (q *QTableView) HorizontalOffset() int {
	var __rv int
	q.Drv(365000,365117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::horizontalScrollbarAction(int)
func (q *QTableView) HorizontalScrollbarAction(action int)  {
	q.Drv(365000,365118,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::indexAt(QPoint const&)
func (q *QTableView) IndexAt(p *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(365000,365119,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTableView::isColumnHidden(int)
func (q *QTableView) IsColumnHidden(column int) bool {
	var __rv bool
	q.Drv(365000,365120,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::isCornerButtonEnabled()
func (q *QTableView) IsCornerButtonEnabled() bool {
	var __rv bool
	q.Drv(365000,365121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::isIndexHidden(QModelIndex const&)
func (q *QTableView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(365000,365122,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::isRowHidden(int)
func (q *QTableView) IsRowHidden(row int) bool {
	var __rv bool
	q.Drv(365000,365123,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::isSortingEnabled()
func (q *QTableView) IsSortingEnabled() bool {
	var __rv bool
	q.Drv(365000,365124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QTableView) MoveCursor(cursorAction QAbstractItemView_CursorAction,modifiers Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(365000,365125,unsafe.Pointer(&cursorAction),unsafe.Pointer(&modifiers),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTableView::paintEvent(QPaintEvent*)
func (q *QTableView) PaintEvent(e *QPaintEvent)  {
	q.Drv(365000,365126,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::resizeColumnToContents(int)
func (q *QTableView) ResizeColumnToContents(column int)  {
	q.Drv(365000,365127,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::resizeColumnsToContents()
func (q *QTableView) ResizeColumnsToContents()  {
	q.Drv(365000,365128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::resizeRowToContents(int)
func (q *QTableView) ResizeRowToContents(row int)  {
	q.Drv(365000,365129,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::resizeRowsToContents()
func (q *QTableView) ResizeRowsToContents()  {
	q.Drv(365000,365130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::rowAt(int)
func (q *QTableView) RowAt(y int) int {
	var __rv int
	q.Drv(365000,365131,unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::rowCountChanged(int,int)
func (q *QTableView) RowCountChanged(oldCount int,newCount int)  {
	q.Drv(365000,365132,unsafe.Pointer(&oldCount),unsafe.Pointer(&newCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::rowHeight(int)
func (q *QTableView) RowHeight(row int) int {
	var __rv int
	q.Drv(365000,365133,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::rowMoved(int,int,int)
func (q *QTableView) RowMoved(row int,oldIndex int,newIndex int)  {
	q.Drv(365000,365134,unsafe.Pointer(&row),unsafe.Pointer(&oldIndex),unsafe.Pointer(&newIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::rowResized(int,int,int)
func (q *QTableView) RowResized(row int,oldHeight int,newHeight int)  {
	q.Drv(365000,365135,unsafe.Pointer(&row),unsafe.Pointer(&oldHeight),unsafe.Pointer(&newHeight),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::rowSpan(int,int)
func (q *QTableView) RowSpan(row int,column int) int {
	var __rv int
	q.Drv(365000,365136,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::rowViewportPosition(int)
func (q *QTableView) RowViewportPosition(row int) int {
	var __rv int
	q.Drv(365000,365137,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::scrollContentsBy(int,int)
func (q *QTableView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(365000,365138,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::scrollTo(QModelIndex const&)
func (q *QTableView) ScrollTo(index *QModelIndex)  {
	q.Drv(365000,365139,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QTableView) ScrollToWithIndexHint(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(365000,365140,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::selectColumn(int)
func (q *QTableView) SelectColumn(column int)  {
	q.Drv(365000,365141,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::selectRow(int)
func (q *QTableView) SelectRow(row int)  {
	q.Drv(365000,365142,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::selectedIndexes()
func (q *QTableView) SelectedIndexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(365000,365143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::selectionChanged(QItemSelection const&,QItemSelection const&)
func (q *QTableView) SelectionChanged(selected *QItemSelection,deselected *QItemSelection)  {
	q.Drv(365000,365144,Native(selected),Native(deselected),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setColumnHidden(int,bool)
func (q *QTableView) SetColumnHidden(column int,hide bool)  {
	q.Drv(365000,365145,unsafe.Pointer(&column),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setColumnWidth(int,int)
func (q *QTableView) SetColumnWidth(column int,width int)  {
	q.Drv(365000,365146,unsafe.Pointer(&column),unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setCornerButtonEnabled(bool)
func (q *QTableView) SetCornerButtonEnabled(enable bool)  {
	q.Drv(365000,365147,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setGridStyle(Qt::PenStyle)
func (q *QTableView) SetGridStyle(style Qt_PenStyle)  {
	q.Drv(365000,365148,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setHorizontalHeader(QHeaderView*)
func (q *QTableView) SetHorizontalHeader(header *QHeaderView)  {
	q.Drv(365000,365149,Native(header),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setModel(QAbstractItemModel*)
func (q *QTableView) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(365000,365150,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setRootIndex(QModelIndex const&)
func (q *QTableView) SetRootIndex(index *QModelIndex)  {
	q.Drv(365000,365151,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setRowHeight(int,int)
func (q *QTableView) SetRowHeight(row int,height int)  {
	q.Drv(365000,365152,unsafe.Pointer(&row),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setRowHidden(int,bool)
func (q *QTableView) SetRowHidden(row int,hide bool)  {
	q.Drv(365000,365153,unsafe.Pointer(&row),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QTableView) SetSelection(rect *QRect,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(365000,365154,Native(rect),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setSelectionModel(QItemSelectionModel*)
func (q *QTableView) SetSelectionModel(selectionModel *QItemSelectionModel)  {
	q.Drv(365000,365155,Native(selectionModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setShowGrid(bool)
func (q *QTableView) SetShowGrid(show bool)  {
	q.Drv(365000,365156,unsafe.Pointer(&show),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setSortingEnabled(bool)
func (q *QTableView) SetSortingEnabled(enable bool)  {
	q.Drv(365000,365157,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setSpan(int,int,int,int)
func (q *QTableView) SetSpan(row int,column int,rowSpan int,columnSpan int)  {
	q.Drv(365000,365158,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&rowSpan),unsafe.Pointer(&columnSpan),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setVerticalHeader(QHeaderView*)
func (q *QTableView) SetVerticalHeader(header *QHeaderView)  {
	q.Drv(365000,365159,Native(header),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::setWordWrap(bool)
func (q *QTableView) SetWordWrap(on bool)  {
	q.Drv(365000,365160,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::showColumn(int)
func (q *QTableView) ShowColumn(column int)  {
	q.Drv(365000,365161,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::showGrid()
func (q *QTableView) ShowGrid() bool {
	var __rv bool
	q.Drv(365000,365162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::showRow(int)
func (q *QTableView) ShowRow(row int)  {
	q.Drv(365000,365163,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::sizeHintForColumn(int)
func (q *QTableView) SizeHintForColumn(column int) int {
	var __rv int
	q.Drv(365000,365164,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::sizeHintForRow(int)
func (q *QTableView) SizeHintForRow(row int) int {
	var __rv int
	q.Drv(365000,365165,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::sortByColumn(int)
func (q *QTableView) SortByColumn(column int)  {
	q.Drv(365000,365166,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::sortByColumn(int,Qt::SortOrder)
func (q *QTableView) SortByColumnWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(365000,365167,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::timerEvent(QTimerEvent*)
func (q *QTableView) TimerEvent(event *QTimerEvent)  {
	q.Drv(365000,365168,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::updateGeometries()
func (q *QTableView) UpdateGeometries()  {
	q.Drv(365000,365169,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::verticalHeader()
func (q *QTableView) VerticalHeader() *QHeaderView {
	var __rv uintptr
	q.Drv(365000,365170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QHeaderView{}
	_rp.SetDriver(__rv,290000,false)
	return _rp
}	
//QTableView::verticalOffset()
func (q *QTableView) VerticalOffset() int {
	var __rv int
	q.Drv(365000,365171,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableView::verticalScrollbarAction(int)
func (q *QTableView) VerticalScrollbarAction(action int)  {
	q.Drv(365000,365172,unsafe.Pointer(&action),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableView::visualRect(QModelIndex const&)
func (q *QTableView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(365000,365173,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTableView::visualRegionForSelection(QItemSelection const&)
func (q *QTableView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(365000,365174,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QTableView::wordWrap()
func (q *QTableView) WordWrap() bool {
	var __rv bool
	q.Drv(365000,365175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
