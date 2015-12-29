// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDirModel_Roles - QDirModel::Roles
type QDirModel_Roles uint32
const (
	QDirModel_FileIconRole QDirModel_Roles = QDirModel_Roles(Qt_DecorationRole)
	QDirModel_FilePathRole QDirModel_Roles = QDirModel_Roles(Qt_UserRole+1)
	QDirModel_FileNameRole QDirModel_Roles = QDirModel_Roles(Qt_UserRole+1+1)
)
//struct QDirModel : QDirModel
type QDirModel struct {
	QAbstractItemModel
}
//QDirModel::QDirModel()	
func NewDirModel() *QDirModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),228000,228102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirModel{}
	_p.SetDriver(__rv,228000,false)
	return _p
} 
//QDirModel::QDirModel(QObject*)	
func NewDirModelWithParent(parent QObjectInterface) *QDirModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),228000,228103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirModel{}
	_p.SetDriver(__rv,228000,false)
	return _p
} 
//QDirModel::QDirModel(QStringList const&,QFlags<QDir::Filter>,QFlags<QDir::SortFlag>,QObject*)	
func NewDirModelWithNamefiltersFiltersSortParent(nameFilters []string,filters QDir_Filter,sort QDir_SortFlag,parent QObjectInterface) *QDirModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),228000,228104,unsafe.Pointer(&nameFilters),unsafe.Pointer(&filters),unsafe.Pointer(&sort),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirModel{}
	_p.SetDriver(__rv,228000,false)
	return _p
} 
//QDirModel::columnCount()
func (q *QDirModel) ColumnCount() int {
	var __rv int
	q.Drv(228000,228105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::columnCount(QModelIndex const&)
func (q *QDirModel) ColumnCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(228000,228106,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::data(QModelIndex const&)
func (q *QDirModel) Data(index *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(228000,228107,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QDirModel::data(QModelIndex const&,int)
func (q *QDirModel) DataWithIndexRole(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(228000,228108,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QDirModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QDirModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(228000,228109,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::fileIcon(QModelIndex const&)
func (q *QDirModel) FileIcon(index *QModelIndex) *QIcon {
	var __rv uintptr
	q.Drv(228000,228110,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QDirModel::fileInfo(QModelIndex const&)
func (q *QDirModel) FileInfo(index *QModelIndex) *QFileInfo {
	var __rv uintptr
	q.Drv(228000,228111,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileInfo{}
	_rp.SetDriver(__rv,34000,true)
	return _rp
}	
//QDirModel::fileName(QModelIndex const&)
func (q *QDirModel) FileName(index *QModelIndex) string {
	var __rv string
	q.Drv(228000,228112,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::filePath(QModelIndex const&)
func (q *QDirModel) FilePath(index *QModelIndex) string {
	var __rv string
	q.Drv(228000,228113,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::filter()
func (q *QDirModel) Filter() QDir_Filter {
	var __rv QDir_Filter
	q.Drv(228000,228114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::flags(QModelIndex const&)
func (q *QDirModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(228000,228115,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::hasChildren()
func (q *QDirModel) HasChildren() bool {
	var __rv bool
	q.Drv(228000,228116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::hasChildren(QModelIndex const&)
func (q *QDirModel) HasChildrenWithIndex(index *QModelIndex) bool {
	var __rv bool
	q.Drv(228000,228117,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::headerData(int,Qt::Orientation,int)
func (q *QDirModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(228000,228118,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QDirModel::iconProvider()
func (q *QDirModel) IconProvider() *QFileIconProvider {
	var __rv uintptr
	q.Drv(228000,228119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileIconProvider{}
	_rp.SetDriver(__rv,33000,true)
	return _rp
}	
//QDirModel::index(QString const&)
func (q *QDirModel) Index(path string) *QModelIndex {
	var __rv uintptr
	q.Drv(228000,228120,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDirModel::index(QString const&,int)
func (q *QDirModel) IndexWithPathColumn(path string,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(228000,228121,unsafe.Pointer(&path),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDirModel::index(int,int,QModelIndex const&)
func (q *QDirModel) IndexWithRowColumnParent(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(228000,228122,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDirModel::isDir(QModelIndex const&)
func (q *QDirModel) IsDir(index *QModelIndex) bool {
	var __rv bool
	q.Drv(228000,228123,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::isReadOnly()
func (q *QDirModel) IsReadOnly() bool {
	var __rv bool
	q.Drv(228000,228124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::lazyChildCount()
func (q *QDirModel) LazyChildCount() bool {
	var __rv bool
	q.Drv(228000,228125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::mimeData(QList<QModelIndex> const&)
func (q *QDirModel) MimeData(indexes []*QModelIndex) *QMimeData {
	var __rv uintptr
	q.Drv(228000,228126,unsafe.Pointer(&indexes),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QDirModel::mimeTypes()
func (q *QDirModel) MimeTypes() []string {
	var __rv []string
	q.Drv(228000,228127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::mkdir(QModelIndex const&,QString const&)
func (q *QDirModel) Mkdir(parent *QModelIndex,name string) *QModelIndex {
	var __rv uintptr
	q.Drv(228000,228128,Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDirModel::nameFilters()
func (q *QDirModel) NameFilters() []string {
	var __rv []string
	q.Drv(228000,228129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::parent()
func (q *QDirModel) Parent() *QObject {
	var __rv uintptr
	q.Drv(228000,228130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QDirModel::parent(QModelIndex const&)
func (q *QDirModel) ParentWithChild(child *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(228000,228131,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDirModel::refresh()
func (q *QDirModel) Refresh()  {
	q.Drv(228000,228132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::refresh(QModelIndex const&)
func (q *QDirModel) RefreshWithParent(parent *QModelIndex)  {
	q.Drv(228000,228133,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::remove(QModelIndex const&)
func (q *QDirModel) Remove(index *QModelIndex) bool {
	var __rv bool
	q.Drv(228000,228134,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::resolveSymlinks()
func (q *QDirModel) ResolveSymlinks() bool {
	var __rv bool
	q.Drv(228000,228135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::rmdir(QModelIndex const&)
func (q *QDirModel) Rmdir(index *QModelIndex) bool {
	var __rv bool
	q.Drv(228000,228136,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::rowCount()
func (q *QDirModel) RowCount() int {
	var __rv int
	q.Drv(228000,228137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::rowCount(QModelIndex const&)
func (q *QDirModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(228000,228138,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QDirModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(228000,228139,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::setFilter(QFlags<QDir::Filter>)
func (q *QDirModel) SetFilter(filters QDir_Filter)  {
	q.Drv(228000,228140,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setIconProvider(QFileIconProvider*)
func (q *QDirModel) SetIconProvider(provider *QFileIconProvider)  {
	q.Drv(228000,228141,Native(provider),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setLazyChildCount(bool)
func (q *QDirModel) SetLazyChildCount(enable bool)  {
	q.Drv(228000,228142,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setNameFilters(QStringList const&)
func (q *QDirModel) SetNameFilters(filters []string)  {
	q.Drv(228000,228143,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setReadOnly(bool)
func (q *QDirModel) SetReadOnly(enable bool)  {
	q.Drv(228000,228144,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setResolveSymlinks(bool)
func (q *QDirModel) SetResolveSymlinks(enable bool)  {
	q.Drv(228000,228145,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::setSorting(QFlags<QDir::SortFlag>)
func (q *QDirModel) SetSorting(sort QDir_SortFlag)  {
	q.Drv(228000,228146,unsafe.Pointer(&sort),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::sort(int)
func (q *QDirModel) Sort(column int)  {
	q.Drv(228000,228147,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::sort(int,Qt::SortOrder)
func (q *QDirModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(228000,228148,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDirModel::sorting()
func (q *QDirModel) Sorting() QDir_SortFlag {
	var __rv QDir_SortFlag
	q.Drv(228000,228149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirModel::supportedDropActions()
func (q *QDirModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(228000,228150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
