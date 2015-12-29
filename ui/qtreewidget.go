// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTreeWidget : QTreeWidget
type QTreeWidget struct {
	QTreeView
}
func (q *QTreeWidget) OnItemEntered(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemPressed(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(386000,386104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemCollapsed(fn func(*QTreeWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(386000,386105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemActivated(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnCurrentItemChanged(fn func(*QTreeWidgetItem,*QTreeWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(386000,386107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemDoubleClicked(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemExpanded(fn func(*QTreeWidgetItem)) uintptr {
	var __rv uintptr
	q.Drv(386000,386109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemChanged(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTreeWidget) OnItemClicked(fn func(*QTreeWidgetItem,int)) uintptr {
	var __rv uintptr
	q.Drv(386000,386111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTreeWidget::QTreeWidget()	
func NewTreeWidget() *QTreeWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),386000,386112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidget{}
	_p.SetDriver(__rv,386000,false)
	return _p
} 
//QTreeWidget::QTreeWidget(QWidget*)	
func NewTreeWidgetWithParent(parent QWidgetInterface) *QTreeWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),386000,386113,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidget{}
	_p.SetDriver(__rv,386000,false)
	return _p
} 
//QTreeWidget::addTopLevelItem(QTreeWidgetItem*)
func (q *QTreeWidget) AddTopLevelItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386114,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::addTopLevelItems(QList<QTreeWidgetItem*> const&)
func (q *QTreeWidget) AddTopLevelItems(items []*QTreeWidgetItem)  {
	q.Drv(386000,386115,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::clear()
func (q *QTreeWidget) Clear()  {
	q.Drv(386000,386116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::closePersistentEditor(QTreeWidgetItem*)
func (q *QTreeWidget) ClosePersistentEditor(item *QTreeWidgetItem)  {
	q.Drv(386000,386117,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::closePersistentEditor(QTreeWidgetItem*,int)
func (q *QTreeWidget) ClosePersistentEditorWithItemColumn(item *QTreeWidgetItem,column int)  {
	q.Drv(386000,386118,Native(item),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::collapseItem(QTreeWidgetItem const*)
func (q *QTreeWidget) CollapseItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386119,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::columnCount()
func (q *QTreeWidget) ColumnCount() int {
	var __rv int
	q.Drv(386000,386120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::currentColumn()
func (q *QTreeWidget) CurrentColumn() int {
	var __rv int
	q.Drv(386000,386121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::currentItem()
func (q *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::dropEvent(QDropEvent*)
func (q *QTreeWidget) DropEvent(event *QDropEvent)  {
	q.Drv(386000,386123,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::dropMimeData(QTreeWidgetItem*,int,QMimeData const*,Qt::DropAction)
func (q *QTreeWidget) DropMimeData(parent *QTreeWidgetItem,index int,data *QMimeData,action Qt_DropAction) bool {
	var __rv bool
	q.Drv(386000,386124,Native(parent),unsafe.Pointer(&index),Native(data),unsafe.Pointer(&action),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::editItem(QTreeWidgetItem*)
func (q *QTreeWidget) EditItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386125,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::editItem(QTreeWidgetItem*,int)
func (q *QTreeWidget) EditItemWithItemColumn(item *QTreeWidgetItem,column int)  {
	q.Drv(386000,386126,Native(item),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::event(QEvent*)
func (q *QTreeWidget) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(386000,386127,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::expandItem(QTreeWidgetItem const*)
func (q *QTreeWidget) ExpandItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386128,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::findItems(QString const&,QFlags<Qt::MatchFlag>,int)
func (q *QTreeWidget) FindItems(text string,flags Qt_MatchFlag,column int) []*QTreeWidgetItem {
	var __rv []*QTreeWidgetItem
	q.Drv(386000,386129,unsafe.Pointer(&text),unsafe.Pointer(&flags),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::headerItem()
func (q *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::indexFromItem(QTreeWidgetItem*,int)
func (q *QTreeWidget) IndexFromItem(item *QTreeWidgetItem,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(386000,386131,Native(item),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QTreeWidget::indexOfTopLevelItem(QTreeWidgetItem*)
func (q *QTreeWidget) IndexOfTopLevelItem(item *QTreeWidgetItem) int {
	var __rv int
	q.Drv(386000,386132,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::insertTopLevelItem(int,QTreeWidgetItem*)
func (q *QTreeWidget) InsertTopLevelItem(index int,item *QTreeWidgetItem)  {
	q.Drv(386000,386133,unsafe.Pointer(&index),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::insertTopLevelItems(int,QList<QTreeWidgetItem*> const&)
func (q *QTreeWidget) InsertTopLevelItems(index int,items []*QTreeWidgetItem)  {
	q.Drv(386000,386134,unsafe.Pointer(&index),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::invisibleRootItem()
func (q *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::isFirstItemColumnSpanned(QTreeWidgetItem const*)
func (q *QTreeWidget) IsFirstItemColumnSpanned(item *QTreeWidgetItem) bool {
	var __rv bool
	q.Drv(386000,386136,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::isItemExpanded(QTreeWidgetItem const*)
func (q *QTreeWidget) IsItemExpanded(item *QTreeWidgetItem) bool {
	var __rv bool
	q.Drv(386000,386137,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::isItemHidden(QTreeWidgetItem const*)
func (q *QTreeWidget) IsItemHidden(item *QTreeWidgetItem) bool {
	var __rv bool
	q.Drv(386000,386138,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::isItemSelected(QTreeWidgetItem const*)
func (q *QTreeWidget) IsItemSelected(item *QTreeWidgetItem) bool {
	var __rv bool
	q.Drv(386000,386139,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::isSortingEnabled()
func (q *QTreeWidget) IsSortingEnabled() bool {
	var __rv bool
	q.Drv(386000,386140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::itemAbove(QTreeWidgetItem const*)
func (q *QTreeWidget) ItemAbove(item *QTreeWidgetItem) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386141,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::itemAt(QPoint const&)
func (q *QTreeWidget) ItemAt(p *QPoint) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386142,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::itemAt(int,int)
func (q *QTreeWidget) ItemAtWithXY(x int,y int) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386143,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::itemBelow(QTreeWidgetItem const*)
func (q *QTreeWidget) ItemBelow(item *QTreeWidgetItem) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386144,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::itemFromIndex(QModelIndex const&)
func (q *QTreeWidget) ItemFromIndex(index *QModelIndex) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386145,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::itemWidget(QTreeWidgetItem*,int)
func (q *QTreeWidget) ItemWidget(item *QTreeWidgetItem,column int) *QWidget {
	var __rv uintptr
	q.Drv(386000,386146,Native(item),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QTreeWidget::items(QMimeData const*)
func (q *QTreeWidget) Items(data *QMimeData) []*QTreeWidgetItem {
	var __rv []*QTreeWidgetItem
	q.Drv(386000,386147,Native(data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::mimeData(QList<QTreeWidgetItem*> const)
func (q *QTreeWidget) MimeData(items []*QTreeWidgetItem) *QMimeData {
	var __rv uintptr
	q.Drv(386000,386148,unsafe.Pointer(&items),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QTreeWidget::mimeTypes()
func (q *QTreeWidget) MimeTypes() []string {
	var __rv []string
	q.Drv(386000,386149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::openPersistentEditor(QTreeWidgetItem*)
func (q *QTreeWidget) OpenPersistentEditor(item *QTreeWidgetItem)  {
	q.Drv(386000,386150,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::openPersistentEditor(QTreeWidgetItem*,int)
func (q *QTreeWidget) OpenPersistentEditorWithItemColumn(item *QTreeWidgetItem,column int)  {
	q.Drv(386000,386151,Native(item),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::removeItemWidget(QTreeWidgetItem*,int)
func (q *QTreeWidget) RemoveItemWidget(item *QTreeWidgetItem,column int)  {
	q.Drv(386000,386152,Native(item),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::scrollToItem(QTreeWidgetItem const*)
func (q *QTreeWidget) ScrollToItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386153,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::scrollToItem(QTreeWidgetItem const*,QAbstractItemView::ScrollHint)
func (q *QTreeWidget) ScrollToItemWithItemHint(item *QTreeWidgetItem,hint QAbstractItemView_ScrollHint)  {
	q.Drv(386000,386154,Native(item),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::selectedItems()
func (q *QTreeWidget) SelectedItems() []*QTreeWidgetItem {
	var __rv []*QTreeWidgetItem
	q.Drv(386000,386155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::setColumnCount(int)
func (q *QTreeWidget) SetColumnCount(columns int)  {
	q.Drv(386000,386156,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setCurrentItem(QTreeWidgetItem*)
func (q *QTreeWidget) SetCurrentItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386157,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setCurrentItem(QTreeWidgetItem*,int)
func (q *QTreeWidget) SetCurrentItemWithItemColumn(item *QTreeWidgetItem,column int)  {
	q.Drv(386000,386158,Native(item),unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setCurrentItem(QTreeWidgetItem*,int,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QTreeWidget) SetCurrentItemWithItemColumnCommand(item *QTreeWidgetItem,column int,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(386000,386159,Native(item),unsafe.Pointer(&column),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setFirstItemColumnSpanned(QTreeWidgetItem const*,bool)
func (q *QTreeWidget) SetFirstItemColumnSpanned(item *QTreeWidgetItem,span bool)  {
	q.Drv(386000,386160,Native(item),unsafe.Pointer(&span),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setHeaderItem(QTreeWidgetItem*)
func (q *QTreeWidget) SetHeaderItem(item *QTreeWidgetItem)  {
	q.Drv(386000,386161,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setHeaderLabel(QString const&)
func (q *QTreeWidget) SetHeaderLabel(label string)  {
	q.Drv(386000,386162,unsafe.Pointer(&label),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setHeaderLabels(QStringList const&)
func (q *QTreeWidget) SetHeaderLabels(labels []string)  {
	q.Drv(386000,386163,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setItemExpanded(QTreeWidgetItem const*,bool)
func (q *QTreeWidget) SetItemExpanded(item *QTreeWidgetItem,expand bool)  {
	q.Drv(386000,386164,Native(item),unsafe.Pointer(&expand),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setItemHidden(QTreeWidgetItem const*,bool)
func (q *QTreeWidget) SetItemHidden(item *QTreeWidgetItem,hide bool)  {
	q.Drv(386000,386165,Native(item),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setItemSelected(QTreeWidgetItem const*,bool)
func (q *QTreeWidget) SetItemSelected(item *QTreeWidgetItem,_select bool)  {
	q.Drv(386000,386166,Native(item),unsafe.Pointer(&_select),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setItemWidget(QTreeWidgetItem*,int,QWidget*)
func (q *QTreeWidget) SetItemWidget(item *QTreeWidgetItem,column int,widget QWidgetInterface)  {
	q.Drv(386000,386167,Native(item),unsafe.Pointer(&column),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setSelectionModel(QItemSelectionModel*)
func (q *QTreeWidget) SetSelectionModel(selectionModel *QItemSelectionModel)  {
	q.Drv(386000,386168,Native(selectionModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::setSortingEnabled(bool)
func (q *QTreeWidget) SetSortingEnabled(enable bool)  {
	q.Drv(386000,386169,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::sortColumn()
func (q *QTreeWidget) SortColumn() int {
	var __rv int
	q.Drv(386000,386170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::sortItems(int,Qt::SortOrder)
func (q *QTreeWidget) SortItems(column int,order Qt_SortOrder)  {
	q.Drv(386000,386171,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTreeWidget::supportedDropActions()
func (q *QTreeWidget) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(386000,386172,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::takeTopLevelItem(int)
func (q *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386173,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::topLevelItem(int)
func (q *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	var __rv uintptr
	q.Drv(386000,386174,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTreeWidgetItem{}
	_rp.SetDriver(__rv,177000,true)
	return _rp
}	
//QTreeWidget::topLevelItemCount()
func (q *QTreeWidget) TopLevelItemCount() int {
	var __rv int
	q.Drv(386000,386175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTreeWidget::visualItemRect(QTreeWidgetItem const*)
func (q *QTreeWidget) VisualItemRect(item *QTreeWidgetItem) *QRect {
	var __rv uintptr
	q.Drv(386000,386176,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
