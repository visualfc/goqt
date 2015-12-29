// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QListWidget : QListWidget
type QListWidget struct {
	QListView
}
func (q *QListWidget) OnItemActivated(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemPressed(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnCurrentRowChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(304000,304104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemClicked(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemChanged(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemDoubleClicked(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemEntered(fn func(*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnItemSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(304000,304109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnCurrentTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(304000,304110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QListWidget) OnCurrentItemChanged(fn func(*QListWidgetItem,*QListWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(304000,304111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QListWidget::QListWidget()	
func NewListWidget() *QListWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),304000,304112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidget{}
	_p.SetDriver(__rv,304000,false)
	return _p
} 
//QListWidget::QListWidget(QWidget*)	
func NewListWidgetWithParent(parent QWidgetInterface) *QListWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),304000,304113,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QListWidget{}
	_p.SetDriver(__rv,304000,false)
	return _p
} 
//QListWidget::addItem(QListWidgetItem*)
func (q *QListWidget) AddItem(item *QListWidgetItem)  {
	q.Drv(304000,304114,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::addItem(QString const&)
func (q *QListWidget) AddItemWithLabel(label string)  {
	q.Drv(304000,304115,unsafe.Pointer(&label),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::addItems(QStringList const&)
func (q *QListWidget) AddItems(labels []string)  {
	q.Drv(304000,304116,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::clear()
func (q *QListWidget) Clear()  {
	q.Drv(304000,304117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::closePersistentEditor(QListWidgetItem*)
func (q *QListWidget) ClosePersistentEditor(item *QListWidgetItem)  {
	q.Drv(304000,304118,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::count()
func (q *QListWidget) Count() int {
	var __rv int
	q.Drv(304000,304119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::currentItem()
func (q *QListWidget) CurrentItem() *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::currentRow()
func (q *QListWidget) CurrentRow() int {
	var __rv int
	q.Drv(304000,304121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::dropEvent(QDropEvent*)
func (q *QListWidget) DropEvent(event *QDropEvent)  {
	q.Drv(304000,304122,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::dropMimeData(int,QMimeData const*,Qt::DropAction)
func (q *QListWidget) DropMimeData(index int,data *QMimeData,action Qt_DropAction) bool {
	var __rv bool
	q.Drv(304000,304123,unsafe.Pointer(&index),Native(data),unsafe.Pointer(&action),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::editItem(QListWidgetItem*)
func (q *QListWidget) EditItem(item *QListWidgetItem)  {
	q.Drv(304000,304124,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::event(QEvent*)
func (q *QListWidget) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(304000,304125,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::findItems(QString const&,QFlags<Qt::MatchFlag>)
func (q *QListWidget) FindItems(text string,flags Qt_MatchFlag) []*QListWidgetItem {
	var __rv []*QListWidgetItem
	q.Drv(304000,304126,unsafe.Pointer(&text),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::indexFromItem(QListWidgetItem*)
func (q *QListWidget) IndexFromItem(item *QListWidgetItem) *QModelIndex {
	var __rv uintptr
	q.Drv(304000,304127,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QListWidget::insertItem(int,QListWidgetItem*)
func (q *QListWidget) InsertItemWithRowItem(row int,item *QListWidgetItem)  {
	q.Drv(304000,304128,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::insertItem(int,QString const&)
func (q *QListWidget) InsertItemWithRowLabel(row int,label string)  {
	q.Drv(304000,304129,unsafe.Pointer(&row),unsafe.Pointer(&label),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::insertItems(int,QStringList const&)
func (q *QListWidget) InsertItems(row int,labels []string)  {
	q.Drv(304000,304130,unsafe.Pointer(&row),unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::isItemHidden(QListWidgetItem const*)
func (q *QListWidget) IsItemHidden(item *QListWidgetItem) bool {
	var __rv bool
	q.Drv(304000,304131,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::isItemSelected(QListWidgetItem const*)
func (q *QListWidget) IsItemSelected(item *QListWidgetItem) bool {
	var __rv bool
	q.Drv(304000,304132,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::isSortingEnabled()
func (q *QListWidget) IsSortingEnabled() bool {
	var __rv bool
	q.Drv(304000,304133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::item(int)
func (q *QListWidget) Item(row int) *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304134,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::itemAt(QPoint const&)
func (q *QListWidget) ItemAt(p *QPoint) *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304135,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::itemAt(int,int)
func (q *QListWidget) ItemAtWithXY(x int,y int) *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304136,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::itemFromIndex(QModelIndex const&)
func (q *QListWidget) ItemFromIndex(index *QModelIndex) *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304137,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::itemWidget(QListWidgetItem*)
func (q *QListWidget) ItemWidget(item *QListWidgetItem) *QWidget {
	var __rv uintptr
	q.Drv(304000,304138,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QListWidget::items(QMimeData const*)
func (q *QListWidget) Items(data *QMimeData) []*QListWidgetItem {
	var __rv []*QListWidgetItem
	q.Drv(304000,304139,Native(data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::mimeData(QList<QListWidgetItem*> const)
func (q *QListWidget) MimeData(items []*QListWidgetItem) *QMimeData {
	var __rv uintptr
	q.Drv(304000,304140,unsafe.Pointer(&items),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QListWidget::mimeTypes()
func (q *QListWidget) MimeTypes() []string {
	var __rv []string
	q.Drv(304000,304141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::openPersistentEditor(QListWidgetItem*)
func (q *QListWidget) OpenPersistentEditor(item *QListWidgetItem)  {
	q.Drv(304000,304142,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::removeItemWidget(QListWidgetItem*)
func (q *QListWidget) RemoveItemWidget(item *QListWidgetItem)  {
	q.Drv(304000,304143,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::row(QListWidgetItem const*)
func (q *QListWidget) Row(item *QListWidgetItem) int {
	var __rv int
	q.Drv(304000,304144,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::scrollToItem(QListWidgetItem const*)
func (q *QListWidget) ScrollToItem(item *QListWidgetItem)  {
	q.Drv(304000,304145,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::scrollToItem(QListWidgetItem const*,QAbstractItemView::ScrollHint)
func (q *QListWidget) ScrollToItemWithItemHint(item *QListWidgetItem,hint QAbstractItemView_ScrollHint)  {
	q.Drv(304000,304146,Native(item),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::selectedItems()
func (q *QListWidget) SelectedItems() []*QListWidgetItem {
	var __rv []*QListWidgetItem
	q.Drv(304000,304147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::setCurrentItem(QListWidgetItem*)
func (q *QListWidget) SetCurrentItem(item *QListWidgetItem)  {
	q.Drv(304000,304148,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setCurrentItem(QListWidgetItem*,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QListWidget) SetCurrentItemWithItemCommand(item *QListWidgetItem,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(304000,304149,Native(item),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setCurrentRow(int)
func (q *QListWidget) SetCurrentRow(row int)  {
	q.Drv(304000,304150,unsafe.Pointer(&row),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setCurrentRow(int,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QListWidget) SetCurrentRowWithRowCommand(row int,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(304000,304151,unsafe.Pointer(&row),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setItemHidden(QListWidgetItem const*,bool)
func (q *QListWidget) SetItemHidden(item *QListWidgetItem,hide bool)  {
	q.Drv(304000,304152,Native(item),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setItemSelected(QListWidgetItem const*,bool)
func (q *QListWidget) SetItemSelected(item *QListWidgetItem,_select bool)  {
	q.Drv(304000,304153,Native(item),unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setItemWidget(QListWidgetItem*,QWidget*)
func (q *QListWidget) SetItemWidget(item *QListWidgetItem,widget QWidgetInterface)  {
	q.Drv(304000,304154,Native(item),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::setSortingEnabled(bool)
func (q *QListWidget) SetSortingEnabled(enable bool)  {
	q.Drv(304000,304155,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::sortItems()
func (q *QListWidget) SortItems()  {
	q.Drv(304000,304156,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::sortItems(Qt::SortOrder)
func (q *QListWidget) SortItemsWithOrder(order Qt_SortOrder)  {
	q.Drv(304000,304157,unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QListWidget::supportedDropActions()
func (q *QListWidget) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(304000,304158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QListWidget::takeItem(int)
func (q *QListWidget) TakeItem(row int) *QListWidgetItem {
	var __rv uintptr
	q.Drv(304000,304159,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QListWidgetItem{}
	_rp.SetDriver(__rv,71000,true)
	return _rp
}	
//QListWidget::visualItemRect(QListWidgetItem const*)
func (q *QListWidget) VisualItemRect(item *QListWidgetItem) *QRect {
	var __rv uintptr
	q.Drv(304000,304160,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
