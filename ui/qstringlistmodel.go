// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStringListModel : QStringListModel
type QStringListModel struct {
	QAbstractListModel
}
//QStringListModel::data(QModelIndex const&,int)
func (q *QStringListModel) Data(index *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(356000,356102,Native(index),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QStringListModel::flags(QModelIndex const&)
func (q *QStringListModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(356000,356103,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::insertRows(int,int,QModelIndex const&)
func (q *QStringListModel) InsertRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(356000,356104,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::removeRows(int,int,QModelIndex const&)
func (q *QStringListModel) RemoveRows(row int,count int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(356000,356105,unsafe.Pointer(&row),unsafe.Pointer(&count),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::rowCount()
func (q *QStringListModel) RowCount() int {
	var __rv int
	q.Drv(356000,356106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::rowCount(QModelIndex const&)
func (q *QStringListModel) RowCountWithParent(parent *QModelIndex) int {
	var __rv int
	q.Drv(356000,356107,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QStringListModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(356000,356108,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::setStringList(QStringList const&)
func (q *QStringListModel) SetStringList(strings []string)  {
	q.Drv(356000,356109,unsafe.Pointer(&strings),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStringListModel::sort(int)
func (q *QStringListModel) Sort(column int)  {
	q.Drv(356000,356110,unsafe.Pointer(&column),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStringListModel::sort(int,Qt::SortOrder)
func (q *QStringListModel) SortWithColumnOrder(column int,order Qt_SortOrder)  {
	q.Drv(356000,356111,unsafe.Pointer(&column),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStringListModel::stringList()
func (q *QStringListModel) StringList() []string {
	var __rv []string
	q.Drv(356000,356112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringListModel::supportedDropActions()
func (q *QStringListModel) SupportedDropActions() Qt_DropAction {
	var __rv Qt_DropAction
	q.Drv(356000,356113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
