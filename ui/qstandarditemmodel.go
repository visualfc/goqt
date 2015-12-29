// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStandardItemModel : QStandardItemModel
type QStandardItemModel struct {
	QAbstractItemModel
}
func (q *QStandardItemModel) OnItemChanged(fn func(*QStandardItem)) uintptr {
	var __rv uintptr
	q.Drv(352000,352102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QStandardItemModel::QStandardItemModel()	
func NewStandardItemModel() *QStandardItemModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),352000,352103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItemModel{}
	_p.SetDriver(__rv,352000,false)
	return _p
} 
//QStandardItemModel::QStandardItemModel(QObject*)	
func NewStandardItemModelWithParent(parent QObjectInterface) *QStandardItemModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),352000,352104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItemModel{}
	_p.SetDriver(__rv,352000,false)
	return _p
} 
//QStandardItemModel::QStandardItemModel(int,int,QObject*)	
func NewStandardItemModelWithRowsColumnsParent(rows int,columns int,parent QObjectInterface) *QStandardItemModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),352000,352105,unsafe.Pointer(&rows),unsafe.Pointer(&columns),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStandardItemModel{}
	_p.SetDriver(__rv,352000,false)
	return _p
} 
//QStandardItemModel::appendColumn(QList<QStandardItem*> const&)
func (q *QStandardItemModel) AppendColumn(items []*QStandardItem)  {
	q.Drv(352000,352106,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::appendRow(QList<QStandardItem*> const&)
func (q *QStandardItemModel) AppendRow(items []*QStandardItem)  {
	q.Drv(352000,352107,unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::appendRow(QStandardItem*)
func (q *QStandardItemModel) AppendRowWithItem(item *QStandardItem)  {
	q.Drv(352000,352108,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::clear()
func (q *QStandardItemModel) Clear()  {
	q.Drv(352000,352109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::columnCount()
func (q *QStandardItemModel) ColumnCount() int {
	var __rv int
	q.Drv(352000,352110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::columnCount(QModelIndex const&)
func (q *QStandardItemModel) ColumnCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(352000,352111,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::data(QModelIndex const&)
func (q *QStandardItemModel) Data(index *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(352000,352112,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStandardItemModel::data(QModelIndex const&,int)
func (q *QStandardItemModel) DataWithIndexRole(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(352000,352113,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStandardItemModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QStandardItemModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352114,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::findItems(QString const&)
func (q *QStandardItemModel) FindItems(text string) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(352000,352115,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::findItems(QString const&,QFlags<Qt::MatchFlag>,int)
func (q *QStandardItemModel) FindItemsWithTextFlagsColumn(text string,flags Qt_MatchFlag,column int) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(352000,352116,unsafe.Pointer(&text),unsafe.Pointer(&flags),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::flags(QModelIndex const&)
func (q *QStandardItemModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(352000,352117,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::hasChildren()
func (q *QStandardItemModel) HasChildren() bool {
	var __rv bool
	q.Drv(352000,352118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::hasChildren(QModelIndex const&)
func (q *QStandardItemModel) HasChildrenWithParent(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352119,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::headerData(int,Qt::Orientation,int)
func (q *QStandardItemModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(352000,352120,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStandardItemModel::horizontalHeaderItem(int)
func (q *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352121,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::index(int,int,QModelIndex const&)
func (q *QStandardItemModel) Index(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(352000,352122,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QStandardItemModel::indexFromItem(QStandardItem const*)
func (q *QStandardItemModel) IndexFromItem(item *QStandardItem) *QModelIndex {
	var __rv uintptr
	q.Drv(352000,352123,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QStandardItemModel::insertColumn(int)
func (q *QStandardItemModel) InsertColumn(column int) bool {
	var __rv bool
	q.Drv(352000,352124,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::insertColumn(int,QList<QStandardItem*> const&)
func (q *QStandardItemModel) InsertColumnWithColumnItems(column int,items []*QStandardItem)  {
	q.Drv(352000,352125,unsafe.Pointer(&column),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::insertColumn(int,QModelIndex const&)
func (q *QStandardItemModel) InsertColumnWithColumnParent(column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352126,unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::insertColumns(int,int,QModelIndex const&)
func (q *QStandardItemModel) InsertColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352127,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::insertRow(int)
func (q *QStandardItemModel) InsertRow(row int) bool {
	var __rv bool
	q.Drv(352000,352128,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::insertRow(int,QList<QStandardItem*> const&)
func (q *QStandardItemModel) InsertRowWithRowItems(row int,items []*QStandardItem)  {
	q.Drv(352000,352129,unsafe.Pointer(&row),unsafe.Pointer(&items),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::insertRow(int,QModelIndex const&)
func (q *QStandardItemModel) InsertRowWithRowParent(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352130,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::insertRow(int,QStandardItem*)
func (q *QStandardItemModel) InsertRowWithRowItem(row int,item *QStandardItem)  {
	q.Drv(352000,352131,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::insertRows(int,int,QModelIndex const&)
func (q *QStandardItemModel) InsertRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352132,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::invisibleRootItem()
func (q *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::item(int)
func (q *QStandardItemModel) Item(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352134,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::item(int,int)
func (q *QStandardItemModel) ItemWithRowColumn(row int,column int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352135,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::itemData(QModelIndex const&)
func (q *QStandardItemModel) ItemData(index *QModelIndex) map[int]*QVariant {
	__rv := make(map[int]*QVariant)
	q.Drv(352000,352136,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::itemFromIndex(QModelIndex const&)
func (q *QStandardItemModel) ItemFromIndex(index *QModelIndex) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352137,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::itemPrototype()
func (q *QStandardItemModel) ItemPrototype() *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::mimeData(QList<QModelIndex> const&)
func (q *QStandardItemModel) MimeData(indexes []*QModelIndex) *QMimeData {
	var __rv uintptr
	q.Drv(352000,352139,unsafe.Pointer(&indexes),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QStandardItemModel::mimeTypes()
func (q *QStandardItemModel) MimeTypes() []string {
	var __rv []string
	q.Drv(352000,352140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::parent()
func (q *QStandardItemModel) Parent() *QObject {
	var __rv uintptr
	q.Drv(352000,352141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QStandardItemModel::parent(QModelIndex const&)
func (q *QStandardItemModel) ParentWithChild(child *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(352000,352142,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QStandardItemModel::removeColumns(int,int,QModelIndex const&)
func (q *QStandardItemModel) RemoveColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352143,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::removeRows(int,int,QModelIndex const&)
func (q *QStandardItemModel) RemoveRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(352000,352144,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::rowCount()
func (q *QStandardItemModel) RowCount() int {
	var __rv int
	q.Drv(352000,352145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::rowCount(QModelIndex const&)
func (q *QStandardItemModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(352000,352146,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::setColumnCount(int)
func (q *QStandardItemModel) SetColumnCount(columns int)  {
	q.Drv(352000,352147,unsafe.Pointer(&columns),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QStandardItemModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(352000,352148,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::setHeaderData(int,Qt::Orientation,QVariant const&,int)
func (q *QStandardItemModel) SetHeaderData(section int,orientation Qt_Orientation,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(352000,352149,unsafe.Pointer(&section),unsafe.Pointer(&orientation),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::setHorizontalHeaderItem(int,QStandardItem*)
func (q *QStandardItemModel) SetHorizontalHeaderItem(column int,item *QStandardItem)  {
	q.Drv(352000,352150,unsafe.Pointer(&column),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setHorizontalHeaderLabels(QStringList const&)
func (q *QStandardItemModel) SetHorizontalHeaderLabels(labels []string)  {
	q.Drv(352000,352151,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setItem(int,QStandardItem*)
func (q *QStandardItemModel) SetItemWithRowItem(row int,item *QStandardItem)  {
	q.Drv(352000,352152,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setItem(int,int,QStandardItem*)
func (q *QStandardItemModel) SetItemWithRowColumnItem(row int,column int,item *QStandardItem)  {
	q.Drv(352000,352153,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setItemData(QModelIndex const&,QMap<int,QVariant> const&)
func (q *QStandardItemModel) SetItemData(index *QModelIndex,roles map[int]*QVariant) bool {
	var __rv bool
	q.Drv(352000,352154,Native(index),unsafe.Pointer(&roles),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::setItemPrototype(QStandardItem const*)
func (q *QStandardItemModel) SetItemPrototype(item *QStandardItem)  {
	q.Drv(352000,352155,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setRowCount(int)
func (q *QStandardItemModel) SetRowCount(rows int)  {
	q.Drv(352000,352156,unsafe.Pointer(&rows),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setSortRole(int)
func (q *QStandardItemModel) SetSortRole(role int)  {
	q.Drv(352000,352157,unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setVerticalHeaderItem(int,QStandardItem*)
func (q *QStandardItemModel) SetVerticalHeaderItem(row int,item *QStandardItem)  {
	q.Drv(352000,352158,unsafe.Pointer(&row),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::setVerticalHeaderLabels(QStringList const&)
func (q *QStandardItemModel) SetVerticalHeaderLabels(labels []string)  {
	q.Drv(352000,352159,unsafe.Pointer(&labels),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::sort(int)
func (q *QStandardItemModel) Sort(column int)  {
	q.Drv(352000,352160,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::sort(int,Qt::SortOrder)
func (q *QStandardItemModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(352000,352161,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStandardItemModel::sortRole()
func (q *QStandardItemModel) SortRole() int {
	var __rv int
	q.Drv(352000,352162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::supportedDropActions()
func (q *QStandardItemModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(352000,352163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::takeColumn(int)
func (q *QStandardItemModel) TakeColumn(column int) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(352000,352164,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::takeHorizontalHeaderItem(int)
func (q *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352165,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::takeItem(int)
func (q *QStandardItemModel) TakeItem(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352166,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::takeItem(int,int)
func (q *QStandardItemModel) TakeItemWithRowColumn(row int,column int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352167,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::takeRow(int)
func (q *QStandardItemModel) TakeRow(row int) []*QStandardItem {
	var __rv []*QStandardItem
	q.Drv(352000,352168,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStandardItemModel::takeVerticalHeaderItem(int)
func (q *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352169,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
//QStandardItemModel::verticalHeaderItem(int)
func (q *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	var __rv uintptr
	q.Drv(352000,352170,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStandardItem{}
	_rp.SetDriver(__rv,123000,true)
	return _rp
}	
