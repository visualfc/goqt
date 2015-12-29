// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSortFilterProxyModel : QSortFilterProxyModel
type QSortFilterProxyModel struct {
	QAbstractProxyModel
}
//QSortFilterProxyModel::QSortFilterProxyModel()	
func NewSortFilterProxyModel() *QSortFilterProxyModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),345000,345102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSortFilterProxyModel{}
	_p.SetDriver(__rv,345000,false)
	return _p
} 
//QSortFilterProxyModel::QSortFilterProxyModel(QObject*)	
func NewSortFilterProxyModelWithParent(parent QObjectInterface) *QSortFilterProxyModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),345000,345103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSortFilterProxyModel{}
	_p.SetDriver(__rv,345000,false)
	return _p
} 
//QSortFilterProxyModel::buddy(QModelIndex const&)
func (q *QSortFilterProxyModel) Buddy(index *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(345000,345104,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QSortFilterProxyModel::canFetchMore(QModelIndex const&)
func (q *QSortFilterProxyModel) CanFetchMore(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345105,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::clear()
func (q *QSortFilterProxyModel) Clear()  {
	q.Drv(345000,345106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::columnCount()
func (q *QSortFilterProxyModel) ColumnCount() int {
	var __rv int
	q.Drv(345000,345107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::columnCount(QModelIndex const&)
func (q *QSortFilterProxyModel) ColumnCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(345000,345108,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::data(QModelIndex const&)
func (q *QSortFilterProxyModel) Data(index *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(345000,345109,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QSortFilterProxyModel::data(QModelIndex const&,int)
func (q *QSortFilterProxyModel) DataWithIndexRole(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(345000,345110,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QSortFilterProxyModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345111,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::dynamicSortFilter()
func (q *QSortFilterProxyModel) DynamicSortFilter() bool {
	var __rv bool
	q.Drv(345000,345112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::fetchMore(QModelIndex const&)
func (q *QSortFilterProxyModel) FetchMore(parent *QModelIndex)  {
	q.Drv(345000,345113,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::filterAcceptsColumn(int,QModelIndex const&)
func (q *QSortFilterProxyModel) FilterAcceptsColumn(source_column int,source_parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345114,unsafe.Pointer(&source_column),Native(source_parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::filterAcceptsRow(int,QModelIndex const&)
func (q *QSortFilterProxyModel) FilterAcceptsRow(source_row int,source_parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345115,unsafe.Pointer(&source_row),Native(source_parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::filterCaseSensitivity()
func (q *QSortFilterProxyModel) FilterCaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(345000,345116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::filterChanged()
func (q *QSortFilterProxyModel) FilterChanged()  {
	q.Drv(345000,345117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::filterKeyColumn()
func (q *QSortFilterProxyModel) FilterKeyColumn() int {
	var __rv int
	q.Drv(345000,345118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::filterRegExp()
func (q *QSortFilterProxyModel) FilterRegExp() *QRegExp {
	var __rv uintptr
	q.Drv(345000,345119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegExp{}
	_rp.SetDriver(__rv,112000,true)
	return _rp
}	
//QSortFilterProxyModel::filterRole()
func (q *QSortFilterProxyModel) FilterRole() int {
	var __rv int
	q.Drv(345000,345120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::flags(QModelIndex const&)
func (q *QSortFilterProxyModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(345000,345121,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::hasChildren()
func (q *QSortFilterProxyModel) HasChildren() bool {
	var __rv bool
	q.Drv(345000,345122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::hasChildren(QModelIndex const&)
func (q *QSortFilterProxyModel) HasChildrenWithParent(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345123,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::headerData(int,Qt::Orientation,int)
func (q *QSortFilterProxyModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(345000,345124,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QSortFilterProxyModel::index(int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) Index(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(345000,345125,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QSortFilterProxyModel::insertColumns(int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) InsertColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345126,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::insertRows(int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) InsertRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345127,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::invalidate()
func (q *QSortFilterProxyModel) Invalidate()  {
	q.Drv(345000,345128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::invalidateFilter()
func (q *QSortFilterProxyModel) InvalidateFilter()  {
	q.Drv(345000,345129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::isSortLocaleAware()
func (q *QSortFilterProxyModel) IsSortLocaleAware() bool {
	var __rv bool
	q.Drv(345000,345130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::lessThan(QModelIndex const&,QModelIndex const&)
func (q *QSortFilterProxyModel) LessThan(left *QModelIndex,right *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345131,Native(left),Native(right),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::mapFromSource(QModelIndex const&)
func (q *QSortFilterProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(345000,345132,Native(sourceIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QSortFilterProxyModel::mapSelectionFromSource(QItemSelection const&)
func (q *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection *QItemSelection) *QItemSelection {
	var __rv uintptr
	q.Drv(345000,345133,Native(sourceSelection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelection{}
	_rp.SetDriver(__rv,62000,true)
	return _rp
}	
//QSortFilterProxyModel::mapSelectionToSource(QItemSelection const&)
func (q *QSortFilterProxyModel) MapSelectionToSource(proxySelection *QItemSelection) *QItemSelection {
	var __rv uintptr
	q.Drv(345000,345134,Native(proxySelection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelection{}
	_rp.SetDriver(__rv,62000,true)
	return _rp
}	
//QSortFilterProxyModel::mapToSource(QModelIndex const&)
func (q *QSortFilterProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(345000,345135,Native(proxyIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QSortFilterProxyModel::match(QModelIndex const&,int,QVariant const&,int,QFlags<Qt::MatchFlag>)
func (q *QSortFilterProxyModel) Match(start *QModelIndex,role int,value *QVariant,hits int,flags Qt_MatchFlag) []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(345000,345136,Native(start),unsafe.Pointer(&role),Native(value),unsafe.Pointer(&hits),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::mimeData(QList<QModelIndex> const&)
func (q *QSortFilterProxyModel) MimeData(indexes []*QModelIndex) *QMimeData {
	var __rv uintptr
	q.Drv(345000,345137,unsafe.Pointer(&indexes),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QSortFilterProxyModel::mimeTypes()
func (q *QSortFilterProxyModel) MimeTypes() []string {
	var __rv []string
	q.Drv(345000,345138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::parent()
func (q *QSortFilterProxyModel) Parent() *QObject {
	var __rv uintptr
	q.Drv(345000,345139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QSortFilterProxyModel::parent(QModelIndex const&)
func (q *QSortFilterProxyModel) ParentWithChild(child *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(345000,345140,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QSortFilterProxyModel::removeColumns(int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) RemoveColumns(column int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345141,unsafe.Pointer(&column),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::removeRows(int,int,QModelIndex const&)
func (q *QSortFilterProxyModel) RemoveRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(345000,345142,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::rowCount()
func (q *QSortFilterProxyModel) RowCount() int {
	var __rv int
	q.Drv(345000,345143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::rowCount(QModelIndex const&)
func (q *QSortFilterProxyModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(345000,345144,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QSortFilterProxyModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(345000,345145,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::setDynamicSortFilter(bool)
func (q *QSortFilterProxyModel) SetDynamicSortFilter(enable bool)  {
	q.Drv(345000,345146,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterCaseSensitivity(Qt::CaseSensitivity)
func (q *QSortFilterProxyModel) SetFilterCaseSensitivity(cs Qt_CaseSensitivity)  {
	q.Drv(345000,345147,unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterFixedString(QString const&)
func (q *QSortFilterProxyModel) SetFilterFixedString(pattern string)  {
	q.Drv(345000,345148,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterKeyColumn(int)
func (q *QSortFilterProxyModel) SetFilterKeyColumn(column int)  {
	q.Drv(345000,345149,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterRegExp(QRegExp const&)
func (q *QSortFilterProxyModel) SetFilterRegExp(regExp *QRegExp)  {
	q.Drv(345000,345150,Native(regExp),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterRegExp(QString const&)
func (q *QSortFilterProxyModel) SetFilterRegExpWithPattern(pattern string)  {
	q.Drv(345000,345151,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterRole(int)
func (q *QSortFilterProxyModel) SetFilterRole(role int)  {
	q.Drv(345000,345152,unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setFilterWildcard(QString const&)
func (q *QSortFilterProxyModel) SetFilterWildcard(pattern string)  {
	q.Drv(345000,345153,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setHeaderData(int,Qt::Orientation,QVariant const&,int)
func (q *QSortFilterProxyModel) SetHeaderData(section int,orientation Qt_Orientation,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(345000,345154,unsafe.Pointer(&section),unsafe.Pointer(&orientation),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::setSortCaseSensitivity(Qt::CaseSensitivity)
func (q *QSortFilterProxyModel) SetSortCaseSensitivity(cs Qt_CaseSensitivity)  {
	q.Drv(345000,345155,unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setSortLocaleAware(bool)
func (q *QSortFilterProxyModel) SetSortLocaleAware(on bool)  {
	q.Drv(345000,345156,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setSortRole(int)
func (q *QSortFilterProxyModel) SetSortRole(role int)  {
	q.Drv(345000,345157,unsafe.Pointer(&role),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::setSourceModel(QAbstractItemModel*)
func (q *QSortFilterProxyModel) SetSourceModel(sourceModel QAbstractItemModelInterface)  {
	q.Drv(345000,345158,Native(sourceModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::sort(int)
func (q *QSortFilterProxyModel) Sort(column int)  {
	q.Drv(345000,345159,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::sort(int,Qt::SortOrder)
func (q *QSortFilterProxyModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(345000,345160,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSortFilterProxyModel::sortCaseSensitivity()
func (q *QSortFilterProxyModel) SortCaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(345000,345161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::sortColumn()
func (q *QSortFilterProxyModel) SortColumn() int {
	var __rv int
	q.Drv(345000,345162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::sortOrder()
func (q *QSortFilterProxyModel) SortOrder() Qt_SortOrder {
	var __rv Qt_SortOrder
	q.Drv(345000,345163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::sortRole()
func (q *QSortFilterProxyModel) SortRole() int {
	var __rv int
	q.Drv(345000,345164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSortFilterProxyModel::span(QModelIndex const&)
func (q *QSortFilterProxyModel) Span(index *QModelIndex) *QSize {
	var __rv uintptr
	q.Drv(345000,345165,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSortFilterProxyModel::supportedDropActions()
func (q *QSortFilterProxyModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(345000,345166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
