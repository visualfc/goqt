// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTableWidget : QTableWidget
type QTableWidget struct {
	QTableView
}
func (q *QTableWidget) OnItemEntered(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCurrentCellChanged(fn func(int,int,int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemPressed(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(366000,366105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellActivated(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellClicked(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellDoubleClicked(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemActivated(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCurrentItemChanged(fn func(*QTableWidgetItem,*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemDoubleClicked(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366112,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellPressed(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366113,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnCellEntered(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(366000,366114,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemChanged(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366115,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTableWidget) OnItemClicked(fn func(*QTableWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(366000,366116,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTableWidget::QTableWidget()	
func NewTableWidget() *QTableWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),366000,366117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidget{}
	_p.SetDriver(__rv,366000,false)
	return _p
} 
//QTableWidget::QTableWidget(QWidget*)	
func NewTableWidgetWithParent(parent QWidgetInterface) *QTableWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),366000,366118,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidget{}
	_p.SetDriver(__rv,366000,false)
	return _p
} 
//QTableWidget::QTableWidget(int,int,QWidget*)	
func NewTableWidgetWithRowsColumnsParent(rows int,columns int,parent QWidgetInterface) *QTableWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),366000,366119,unsafe.Pointer(&rows),unsafe.Pointer(&columns),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTableWidget{}
	_p.SetDriver(__rv,366000,false)
	return _p
} 
//QTableWidget::cellWidget(int,int)
func (q *QTableWidget) CellWidget(row int,column int) *QWidget {
	var __rv uintptr
	q.Drv(366000,366120,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTableWidget::clear()
func (q *QTableWidget) Clear()  {
	q.Drv(366000,366121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::clearContents()
func (q *QTableWidget) ClearContents()  {
	q.Drv(366000,366122,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::closePersistentEditor(QTableWidgetItem*)
func (q *QTableWidget) ClosePersistentEditor(item *QTableWidgetItem)  {
	q.Drv(366000,366123,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::column(QTableWidgetItem const*)
func (q *QTableWidget) Column(item *QTableWidgetItem) int {
	var __rv int
	q.Drv(366000,366124,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::columnCount()
func (q *QTableWidget) ColumnCount() int {
	var __rv int
	q.Drv(366000,366125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::currentColumn()
func (q *QTableWidget) CurrentColumn() int {
	var __rv int
	q.Drv(366000,366126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::currentItem()
func (q *QTableWidget) CurrentItem() *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::currentRow()
func (q *QTableWidget) CurrentRow() int {
	var __rv int
	q.Drv(366000,366128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::dropEvent(QDropEvent*)
func (q *QTableWidget) DropEvent(event *QDropEvent)  {
	q.Drv(366000,366129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::dropMimeData(int,int,QMimeData const*,Qt::DropAction)
func (q *QTableWidget) DropMimeData(row int,column int,data *QMimeData,action Qt_DropAction) bool {
	var __rv bool
	q.Drv(366000,366130,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(data),unsafe.Pointer(&action),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::editItem(QTableWidgetItem*)
func (q *QTableWidget) EditItem(item *QTableWidgetItem)  {
	q.Drv(366000,366131,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::event(QEvent*)
func (q *QTableWidget) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(366000,366132,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::findItems(QString const&,QFlags<Qt::MatchFlag>)
func (q *QTableWidget) FindItems(text string,flags Qt_MatchFlag) []*QTableWidgetItem {
	var __rv []*QTableWidgetItem
	q.Drv(366000,366133,unsafe.Pointer(&text),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::horizontalHeaderItem(int)
func (q *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366134,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::indexFromItem(QTableWidgetItem*)
func (q *QTableWidget) IndexFromItem(item *QTableWidgetItem) *QModelIndex {
	var __rv uintptr
	q.Drv(366000,366135,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTableWidget::insertColumn(int)
func (q *QTableWidget) InsertColumn(column int)  {
	q.Drv(366000,366136,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::insertRow(int)
func (q *QTableWidget) InsertRow(row int)  {
	q.Drv(366000,366137,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::isItemSelected(QTableWidgetItem const*)
func (q *QTableWidget) IsItemSelected(item *QTableWidgetItem) bool {
	var __rv bool
	q.Drv(366000,366138,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::isSortingEnabled()
func (q *QTableWidget) IsSortingEnabled() bool {
	var __rv bool
	q.Drv(366000,366139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::item(int,int)
func (q *QTableWidget) Item(row int,column int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366140,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::itemAt(QPoint const&)
func (q *QTableWidget) ItemAt(p *QPoint) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366141,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::itemAt(int,int)
func (q *QTableWidget) ItemAtWithXY(x int,y int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366142,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::itemFromIndex(QModelIndex const&)
func (q *QTableWidget) ItemFromIndex(index *QModelIndex) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366143,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::itemPrototype()
func (q *QTableWidget) ItemPrototype() *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::items(QMimeData const*)
func (q *QTableWidget) Items(data *QMimeData) []*QTableWidgetItem {
	var __rv []*QTableWidgetItem
	q.Drv(366000,366145,Native(data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::mimeData(QList<QTableWidgetItem*> const)
func (q *QTableWidget) MimeData(items []*QTableWidgetItem) *QMimeData {
	var __rv uintptr
	q.Drv(366000,366146,unsafe.Pointer(&items),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QTableWidget::mimeTypes()
func (q *QTableWidget) MimeTypes() []string {
	var __rv []string
	q.Drv(366000,366147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::openPersistentEditor(QTableWidgetItem*)
func (q *QTableWidget) OpenPersistentEditor(item *QTableWidgetItem)  {
	q.Drv(366000,366148,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::removeCellWidget(int,int)
func (q *QTableWidget) RemoveCellWidget(row int,column int)  {
	q.Drv(366000,366149,unsafe.Pointer(&row),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::removeColumn(int)
func (q *QTableWidget) RemoveColumn(column int)  {
	q.Drv(366000,366150,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::removeRow(int)
func (q *QTableWidget) RemoveRow(row int)  {
	q.Drv(366000,366151,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::row(QTableWidgetItem const*)
func (q *QTableWidget) Row(item *QTableWidgetItem) int {
	var __rv int
	q.Drv(366000,366152,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::rowCount()
func (q *QTableWidget) RowCount() int {
	var __rv int
	q.Drv(366000,366153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::scrollToItem(QTableWidgetItem const*)
func (q *QTableWidget) ScrollToItem(item *QTableWidgetItem)  {
	q.Drv(366000,366154,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::scrollToItem(QTableWidgetItem const*,QAbstractItemView::ScrollHint)
func (q *QTableWidget) ScrollToItemWithItemHint(item *QTableWidgetItem,hint QAbstractItemView_ScrollHint)  {
	q.Drv(366000,366155,Native(item),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::selectedItems()
func (q *QTableWidget) SelectedItems() []*QTableWidgetItem {
	var __rv []*QTableWidgetItem
	q.Drv(366000,366156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::selectedRanges()
func (q *QTableWidget) SelectedRanges() []*QTableWidgetSelectionRange {
	var __rv []*QTableWidgetSelectionRange
	q.Drv(366000,366157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::setCellWidget(int,int,QWidget*)
func (q *QTableWidget) SetCellWidget(row int,column int,widget QWidgetInterface)  {
	q.Drv(366000,366158,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setColumnCount(int)
func (q *QTableWidget) SetColumnCount(columns int)  {
	q.Drv(366000,366159,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setCurrentCell(int,int)
func (q *QTableWidget) SetCurrentCellWithRowColumn(row int,column int)  {
	q.Drv(366000,366160,unsafe.Pointer(&row),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setCurrentCell(int,int,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QTableWidget) SetCurrentCellWithRowColumnCommand(row int,column int,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(366000,366161,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setCurrentItem(QTableWidgetItem*)
func (q *QTableWidget) SetCurrentItem(item *QTableWidgetItem)  {
	q.Drv(366000,366162,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setCurrentItem(QTableWidgetItem*,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QTableWidget) SetCurrentItemWithItemCommand(item *QTableWidgetItem,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(366000,366163,Native(item),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setHorizontalHeaderItem(int,QTableWidgetItem*)
func (q *QTableWidget) SetHorizontalHeaderItem(column int,item *QTableWidgetItem)  {
	q.Drv(366000,366164,unsafe.Pointer(&column),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setHorizontalHeaderLabels(QStringList const&)
func (q *QTableWidget) SetHorizontalHeaderLabels(labels []string)  {
	q.Drv(366000,366165,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setItem(int,int,QTableWidgetItem*)
func (q *QTableWidget) SetItem(row int,column int,item *QTableWidgetItem)  {
	q.Drv(366000,366166,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setItemPrototype(QTableWidgetItem const*)
func (q *QTableWidget) SetItemPrototype(item *QTableWidgetItem)  {
	q.Drv(366000,366167,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setItemSelected(QTableWidgetItem const*,bool)
func (q *QTableWidget) SetItemSelected(item *QTableWidgetItem,_select bool)  {
	q.Drv(366000,366168,Native(item),unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setRangeSelected(QTableWidgetSelectionRange const&,bool)
func (q *QTableWidget) SetRangeSelected(_range *QTableWidgetSelectionRange,_select bool)  {
	q.Drv(366000,366169,Native(_range),unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setRowCount(int)
func (q *QTableWidget) SetRowCount(rows int)  {
	q.Drv(366000,366170,unsafe.Pointer(&rows),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setSortingEnabled(bool)
func (q *QTableWidget) SetSortingEnabled(enable bool)  {
	q.Drv(366000,366171,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setVerticalHeaderItem(int,QTableWidgetItem*)
func (q *QTableWidget) SetVerticalHeaderItem(row int,item *QTableWidgetItem)  {
	q.Drv(366000,366172,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::setVerticalHeaderLabels(QStringList const&)
func (q *QTableWidget) SetVerticalHeaderLabels(labels []string)  {
	q.Drv(366000,366173,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::sortItems(int)
func (q *QTableWidget) SortItems(column int)  {
	q.Drv(366000,366174,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::sortItems(int,Qt::SortOrder)
func (q *QTableWidget) SortItemsWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(366000,366175,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTableWidget::supportedDropActions()
func (q *QTableWidget) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(366000,366176,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::takeHorizontalHeaderItem(int)
func (q *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366177,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::takeItem(int,int)
func (q *QTableWidget) TakeItem(row int,column int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366178,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::takeVerticalHeaderItem(int)
func (q *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366179,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::verticalHeaderItem(int)
func (q *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	var __rv uintptr
	q.Drv(366000,366180,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTableWidgetItem{}
	_rp.SetDriver(__rv,134000,true)
	return _rp
}	
//QTableWidget::visualColumn(int)
func (q *QTableWidget) VisualColumn(logicalColumn int) int {
	var __rv int
	q.Drv(366000,366181,unsafe.Pointer(&logicalColumn),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTableWidget::visualItemRect(QTableWidgetItem const*)
func (q *QTableWidget) VisualItemRect(item *QTableWidgetItem) *QRect {
	var __rv uintptr
	q.Drv(366000,366182,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTableWidget::visualRow(int)
func (q *QTableWidget) VisualRow(logicalRow int) int {
	var __rv int
	q.Drv(366000,366183,unsafe.Pointer(&logicalRow),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
