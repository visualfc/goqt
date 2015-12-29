// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractItemModel : QAbstractItemModel
type QAbstractItemModel struct {
	QObject
}
func (q *QAbstractItemModel) OnLayoutChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(195000,195102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemModel) OnHeaderDataChanged(fn func(Qt_Orientation,int,int)) uintptr {
	var __rv uintptr
	q.Drv(195000,195103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemModel) OnDataChanged(fn func(*QModelIndex,*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(195000,195104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemModel) OnLayoutAboutToBeChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(195000,195105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractItemModel::beginInsertColumns(QModelIndex const&,int,int)
func (q *QAbstractItemModel) BeginInsertColumns(parent *QModelIndex,first int,last int)  {
	q.Drv(195000,195106,Native(parent),unsafe.Pointer(&first),unsafe.Pointer(&last),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::beginInsertRows(QModelIndex const&,int,int)
func (q *QAbstractItemModel) BeginInsertRows(parent *QModelIndex,first int,last int)  {
	q.Drv(195000,195107,Native(parent),unsafe.Pointer(&first),unsafe.Pointer(&last),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::beginMoveColumns(QModelIndex const&,int,int,QModelIndex const&,int)
func (q *QAbstractItemModel) BeginMoveColumns(sourceParent *QModelIndex,sourceFirst int,sourceLast int,destinationParent *QModelIndex,destinationColumn int) bool {
	var __rv bool
	q.Drv(195000,195108,Native(sourceParent),unsafe.Pointer(&sourceFirst),unsafe.Pointer(&sourceLast),Native(destinationParent),unsafe.Pointer(&destinationColumn),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::beginMoveRows(QModelIndex const&,int,int,QModelIndex const&,int)
func (q *QAbstractItemModel) BeginMoveRows(sourceParent *QModelIndex,sourceFirst int,sourceLast int,destinationParent *QModelIndex,destinationRow int) bool {
	var __rv bool
	q.Drv(195000,195109,Native(sourceParent),unsafe.Pointer(&sourceFirst),unsafe.Pointer(&sourceLast),Native(destinationParent),unsafe.Pointer(&destinationRow),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::beginRemoveColumns(QModelIndex const&,int,int)
func (q *QAbstractItemModel) BeginRemoveColumns(parent *QModelIndex,first int,last int)  {
	q.Drv(195000,195110,Native(parent),unsafe.Pointer(&first),unsafe.Pointer(&last),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::beginRemoveRows(QModelIndex const&,int,int)
func (q *QAbstractItemModel) BeginRemoveRows(parent *QModelIndex,first int,last int)  {
	q.Drv(195000,195111,Native(parent),unsafe.Pointer(&first),unsafe.Pointer(&last),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::beginResetModel()
func (q *QAbstractItemModel) BeginResetModel()  {
	q.Drv(195000,195112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::buddy(QModelIndex const&)
func (q *QAbstractItemModel) Buddy(index *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195113,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::canFetchMore(QModelIndex const&)
func (q *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195114,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::changePersistentIndex(QModelIndex const&,QModelIndex const&)
func (q *QAbstractItemModel) ChangePersistentIndex(from *QModelIndex,to *QModelIndex)  {
	q.Drv(195000,195115,Native(from),Native(to),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::changePersistentIndexList(QList<QModelIndex> const&,QList<QModelIndex> const&)
func (q *QAbstractItemModel) ChangePersistentIndexList(from []*QModelIndex,to []*QModelIndex)  {
	q.Drv(195000,195116,unsafe.Pointer(&from),unsafe.Pointer(&to),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::columnCount()
func (q *QAbstractItemModel) ColumnCount() int {
	var __rv int
	q.Drv(195000,195117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::columnCount(QModelIndex const&)
func (q *QAbstractItemModel) ColumnCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(195000,195118,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::createIndex(int,int,int)
func (q *QAbstractItemModel) CreateIndex(row int,column int,id int) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195119,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::createIndex(int,int,unsigned int)
func (q *QAbstractItemModel) CreateIndexWithRowColumnId(row int,column int,id uint) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195120,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&id),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::createIndex(int,int,void*)
func (q *QAbstractItemModel) CreateIndexWithRowColumnData(row int,column int,data uintptr) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195121,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::data(QModelIndex const&)
func (q *QAbstractItemModel) Data(index *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(195000,195122,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractItemModel::data(QModelIndex const&,int)
func (q *QAbstractItemModel) DataWithIndexRole(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(195000,195123,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractItemModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QAbstractItemModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195124,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::endInsertColumns()
func (q *QAbstractItemModel) EndInsertColumns()  {
	q.Drv(195000,195125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endInsertRows()
func (q *QAbstractItemModel) EndInsertRows()  {
	q.Drv(195000,195126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endMoveColumns()
func (q *QAbstractItemModel) EndMoveColumns()  {
	q.Drv(195000,195127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endMoveRows()
func (q *QAbstractItemModel) EndMoveRows()  {
	q.Drv(195000,195128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endRemoveColumns()
func (q *QAbstractItemModel) EndRemoveColumns()  {
	q.Drv(195000,195129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endRemoveRows()
func (q *QAbstractItemModel) EndRemoveRows()  {
	q.Drv(195000,195130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::endResetModel()
func (q *QAbstractItemModel) EndResetModel()  {
	q.Drv(195000,195131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::fetchMore(QModelIndex const&)
func (q *QAbstractItemModel) FetchMore(parent *QModelIndex)  {
	q.Drv(195000,195132,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::flags(QModelIndex const&)
func (q *QAbstractItemModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(195000,195133,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::hasChildren()
func (q *QAbstractItemModel) HasChildren() bool {
	var __rv bool
	q.Drv(195000,195134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::hasChildren(QModelIndex const&)
func (q *QAbstractItemModel) HasChildrenWithParent(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195135,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::hasIndex(int,int,QModelIndex const&)
func (q *QAbstractItemModel) HasIndex(row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195136,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::headerData(int,Qt::Orientation,int)
func (q *QAbstractItemModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(195000,195137,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractItemModel::index(int,int,QModelIndex const&)
func (q *QAbstractItemModel) Index(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195138,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::insertColumn(int)
func (q *QAbstractItemModel) InsertColumn(column int) bool {
	var __rv bool
	q.Drv(195000,195139,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::insertColumn(int,QModelIndex const&)
func (q *QAbstractItemModel) InsertColumnWithColumnParent(column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195140,unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::insertColumns(int,int,QModelIndex const&)
func (q *QAbstractItemModel) InsertColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195141,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::insertRow(int)
func (q *QAbstractItemModel) InsertRow(row int) bool {
	var __rv bool
	q.Drv(195000,195142,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::insertRow(int,QModelIndex const&)
func (q *QAbstractItemModel) InsertRowWithRowParent(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195143,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::insertRows(int,int,QModelIndex const&)
func (q *QAbstractItemModel) InsertRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195144,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::itemData(QModelIndex const&)
func (q *QAbstractItemModel) ItemData(index *QModelIndex) map[int]*QVariant {
	__rv := make(map[int]*QVariant)
	q.Drv(195000,195145,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::match(QModelIndex const&,int,QVariant const&,int,QFlags<Qt::MatchFlag>)
func (q *QAbstractItemModel) Match(start *QModelIndex,role int,value *QVariant,hits int,flags Qt_MatchFlag) []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(195000,195146,Native(start),unsafe.Pointer(&role),Native(value),unsafe.Pointer(&hits),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::mimeData(QList<QModelIndex> const&)
func (q *QAbstractItemModel) MimeData(indexes []*QModelIndex) *QMimeData {
	var __rv uintptr
	q.Drv(195000,195147,unsafe.Pointer(&indexes),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QAbstractItemModel::mimeTypes()
func (q *QAbstractItemModel) MimeTypes() []string {
	var __rv []string
	q.Drv(195000,195148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::parent()
func (q *QAbstractItemModel) Parent() *QObject {
	var __rv uintptr
	q.Drv(195000,195149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QAbstractItemModel::parent(QModelIndex const&)
func (q *QAbstractItemModel) ParentWithChild(child *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195150,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::persistentIndexList()
func (q *QAbstractItemModel) PersistentIndexList() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(195000,195151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeColumn(int)
func (q *QAbstractItemModel) RemoveColumn(column int) bool {
	var __rv bool
	q.Drv(195000,195152,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeColumn(int,QModelIndex const&)
func (q *QAbstractItemModel) RemoveColumnWithColumnParent(column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195153,unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeColumns(int,int,QModelIndex const&)
func (q *QAbstractItemModel) RemoveColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195154,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeRow(int)
func (q *QAbstractItemModel) RemoveRow(row int) bool {
	var __rv bool
	q.Drv(195000,195155,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeRow(int,QModelIndex const&)
func (q *QAbstractItemModel) RemoveRowWithRowParent(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195156,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::removeRows(int,int,QModelIndex const&)
func (q *QAbstractItemModel) RemoveRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(195000,195157,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::revert()
func (q *QAbstractItemModel) Revert()  {
	q.Drv(195000,195158,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::rowCount()
func (q *QAbstractItemModel) RowCount() int {
	var __rv int
	q.Drv(195000,195159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::rowCount(QModelIndex const&)
func (q *QAbstractItemModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(195000,195160,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QAbstractItemModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(195000,195161,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::setHeaderData(int,Qt::Orientation,QVariant const&,int)
func (q *QAbstractItemModel) SetHeaderData(section int,orientation Qt_Orientation,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(195000,195162,unsafe.Pointer(&section),unsafe.Pointer(&orientation),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::setItemData(QModelIndex const&,QMap<int,QVariant> const&)
func (q *QAbstractItemModel) SetItemData(index *QModelIndex,roles map[int]*QVariant) bool {
	var __rv bool
	q.Drv(195000,195163,Native(index),unsafe.Pointer(&roles),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::sibling(int,int,QModelIndex const&)
func (q *QAbstractItemModel) Sibling(row int,column int,idx *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(195000,195164,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(idx),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractItemModel::sort(int)
func (q *QAbstractItemModel) Sort(column int)  {
	q.Drv(195000,195165,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::sort(int,Qt::SortOrder)
func (q *QAbstractItemModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(195000,195166,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemModel::span(QModelIndex const&)
func (q *QAbstractItemModel) Span(index *QModelIndex) *QSize {
	var __rv uintptr
	q.Drv(195000,195167,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractItemModel::submit()
func (q *QAbstractItemModel) Submit() bool {
	var __rv bool
	q.Drv(195000,195168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::supportedDragActions()
func (q *QAbstractItemModel) SupportedDragActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(195000,195169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemModel::supportedDropActions()
func (q *QAbstractItemModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(195000,195170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
