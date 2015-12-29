// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFileSystemModel_Roles - QFileSystemModel::Roles
type QFileSystemModel_Roles uint32
const (
	QFileSystemModel_FileIconRole QFileSystemModel_Roles = QFileSystemModel_Roles(Qt_DecorationRole)
	QFileSystemModel_FilePathRole QFileSystemModel_Roles = QFileSystemModel_Roles(Qt_UserRole+1)
	QFileSystemModel_FileNameRole QFileSystemModel_Roles = QFileSystemModel_Roles(Qt_UserRole+2)
	QFileSystemModel_FilePermissions QFileSystemModel_Roles = QFileSystemModel_Roles(Qt_UserRole+3)
)
//struct QFileSystemModel : QFileSystemModel
type QFileSystemModel struct {
	QAbstractItemModel
}
func (q *QFileSystemModel) OnDirectoryLoaded(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(238000,238102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileSystemModel) OnRootPathChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(238000,238103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileSystemModel) OnFileRenamed(fn func(string,string,string)) uintptr {
	var __rv uintptr
	q.Drv(238000,238104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFileSystemModel::QFileSystemModel()	
func NewFileSystemModel() *QFileSystemModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),238000,238105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileSystemModel{}
	_p.SetDriver(__rv,238000,false)
	return _p
} 
//QFileSystemModel::QFileSystemModel(QObject*)	
func NewFileSystemModelWithParent(parent QObjectInterface) *QFileSystemModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),238000,238106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileSystemModel{}
	_p.SetDriver(__rv,238000,false)
	return _p
} 
//QFileSystemModel::canFetchMore(QModelIndex const&)
func (q *QFileSystemModel) CanFetchMore(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238107,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::columnCount()
func (q *QFileSystemModel) ColumnCount() int {
	var __rv int
	q.Drv(238000,238108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::columnCount(QModelIndex const&)
func (q *QFileSystemModel) ColumnCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(238000,238109,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::data(QModelIndex const&)
func (q *QFileSystemModel) Data(index *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(238000,238110,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QFileSystemModel::data(QModelIndex const&,int)
func (q *QFileSystemModel) DataWithIndexRole(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(238000,238111,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QFileSystemModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QFileSystemModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238112,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::event(QEvent*)
func (q *QFileSystemModel) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(238000,238113,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::fetchMore(QModelIndex const&)
func (q *QFileSystemModel) FetchMore(parent *QModelIndex)  {
	q.Drv(238000,238114,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::fileIcon(QModelIndex const&)
func (q *QFileSystemModel) FileIcon(index *QModelIndex) *QIcon {
	var __rv uintptr
	q.Drv(238000,238115,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QFileSystemModel::fileInfo(QModelIndex const&)
func (q *QFileSystemModel) FileInfo(index *QModelIndex) *QFileInfo {
	var __rv uintptr
	q.Drv(238000,238116,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileInfo{}
	_rp.SetDriver(__rv,34000,true)
	return _rp
}	
//QFileSystemModel::fileName(QModelIndex const&)
func (q *QFileSystemModel) FileName(index *QModelIndex) string {
	var __rv string
	q.Drv(238000,238117,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::filePath(QModelIndex const&)
func (q *QFileSystemModel) FilePath(index *QModelIndex) string {
	var __rv string
	q.Drv(238000,238118,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::filter()
func (q *QFileSystemModel) Filter() QDir_Filter {
	var __rv QDir_Filter
	q.Drv(238000,238119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::flags(QModelIndex const&)
func (q *QFileSystemModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(238000,238120,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::hasChildren()
func (q *QFileSystemModel) HasChildren() bool {
	var __rv bool
	q.Drv(238000,238121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::hasChildren(QModelIndex const&)
func (q *QFileSystemModel) HasChildrenWithParent(parent *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238122,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::headerData(int,Qt::Orientation,int)
func (q *QFileSystemModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(238000,238123,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QFileSystemModel::iconProvider()
func (q *QFileSystemModel) IconProvider() *QFileIconProvider {
	var __rv uintptr
	q.Drv(238000,238124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileIconProvider{}
	_rp.SetDriver(__rv,33000,true)
	return _rp
}	
//QFileSystemModel::index(QString const&)
func (q *QFileSystemModel) Index(path string) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238125,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::index(QString const&,int)
func (q *QFileSystemModel) IndexWithPathColumn(path string,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238126,unsafe.Pointer(&path),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::index(int,int,QModelIndex const&)
func (q *QFileSystemModel) IndexWithRowColumnParent(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238127,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::isDir(QModelIndex const&)
func (q *QFileSystemModel) IsDir(index *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238128,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::isReadOnly()
func (q *QFileSystemModel) IsReadOnly() bool {
	var __rv bool
	q.Drv(238000,238129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::lastModified(QModelIndex const&)
func (q *QFileSystemModel) LastModified(index *QModelIndex) *QDateTime {
	var __rv uintptr
	q.Drv(238000,238130,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QFileSystemModel::mimeData(QList<QModelIndex> const&)
func (q *QFileSystemModel) MimeData(indexes []*QModelIndex) *QMimeData {
	var __rv uintptr
	q.Drv(238000,238131,unsafe.Pointer(&indexes),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QFileSystemModel::mimeTypes()
func (q *QFileSystemModel) MimeTypes() []string {
	var __rv []string
	q.Drv(238000,238132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::mkdir(QModelIndex const&,QString const&)
func (q *QFileSystemModel) Mkdir(parent *QModelIndex,name string) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238133,Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::myComputer()
func (q *QFileSystemModel) MyComputer() *QVariant {
	var __rv uintptr
	q.Drv(238000,238134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QFileSystemModel::myComputer(int)
func (q *QFileSystemModel) MyComputerWithRole(role int) *QVariant {
	var __rv uintptr
	q.Drv(238000,238135,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QFileSystemModel::nameFilterDisables()
func (q *QFileSystemModel) NameFilterDisables() bool {
	var __rv bool
	q.Drv(238000,238136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::nameFilters()
func (q *QFileSystemModel) NameFilters() []string {
	var __rv []string
	q.Drv(238000,238137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::parent(QModelIndex const&)
func (q *QFileSystemModel) Parent(child *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238138,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::permissions(QModelIndex const&)
func (q *QFileSystemModel) Permissions(index *QModelIndex) QFile_Permission {
	var __rv QFile_Permission
	q.Drv(238000,238139,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::remove(QModelIndex const&)
func (q *QFileSystemModel) Remove(index *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238140,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::resolveSymlinks()
func (q *QFileSystemModel) ResolveSymlinks() bool {
	var __rv bool
	q.Drv(238000,238141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::rmdir(QModelIndex const&)
func (q *QFileSystemModel) Rmdir(index *QModelIndex) bool {
	var __rv bool
	q.Drv(238000,238142,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::rootDirectory()
func (q *QFileSystemModel) RootDirectory() *QDir {
	var __rv uintptr
	q.Drv(238000,238143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QFileSystemModel::rootPath()
func (q *QFileSystemModel) RootPath() string {
	var __rv string
	q.Drv(238000,238144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::rowCount()
func (q *QFileSystemModel) RowCount() int {
	var __rv int
	q.Drv(238000,238145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::rowCount(QModelIndex const&)
func (q *QFileSystemModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(238000,238146,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QFileSystemModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(238000,238147,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::setFilter(QFlags<QDir::Filter>)
func (q *QFileSystemModel) SetFilter(filters QDir_Filter)  {
	q.Drv(238000,238148,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setIconProvider(QFileIconProvider*)
func (q *QFileSystemModel) SetIconProvider(provider *QFileIconProvider)  {
	q.Drv(238000,238149,Native(provider),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setNameFilterDisables(bool)
func (q *QFileSystemModel) SetNameFilterDisables(enable bool)  {
	q.Drv(238000,238150,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setNameFilters(QStringList const&)
func (q *QFileSystemModel) SetNameFilters(filters []string)  {
	q.Drv(238000,238151,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setReadOnly(bool)
func (q *QFileSystemModel) SetReadOnly(enable bool)  {
	q.Drv(238000,238152,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setResolveSymlinks(bool)
func (q *QFileSystemModel) SetResolveSymlinks(enable bool)  {
	q.Drv(238000,238153,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::setRootPath(QString const&)
func (q *QFileSystemModel) SetRootPath(path string) *QModelIndex {
	var __rv uintptr
	q.Drv(238000,238154,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QFileSystemModel::size(QModelIndex const&)
func (q *QFileSystemModel) Size(index *QModelIndex) int64 {
	var __rv int64
	q.Drv(238000,238155,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::sort(int)
func (q *QFileSystemModel) Sort(column int)  {
	q.Drv(238000,238156,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::sort(int,Qt::SortOrder)
func (q *QFileSystemModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(238000,238157,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::supportedDropActions()
func (q *QFileSystemModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(238000,238158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemModel::timerEvent(QTimerEvent*)
func (q *QFileSystemModel) TimerEvent(event *QTimerEvent)  {
	q.Drv(238000,238159,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemModel::type(QModelIndex const&)
func (q *QFileSystemModel) Type(index *QModelIndex) string {
	var __rv string
	q.Drv(238000,238160,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
