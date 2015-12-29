// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractProxyModel : QAbstractProxyModel
type QAbstractProxyModel struct {
	QAbstractItemModel
}
//QAbstractProxyModel::data(QModelIndex const&)
func (q *QAbstractProxyModel) Data(proxyIndex *QModelIndex) *QVariant {
	var __rv uintptr
	q.Drv(199000,199102,Native(proxyIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractProxyModel::data(QModelIndex const&,int)
func (q *QAbstractProxyModel) DataWithProxyindexRole(proxyIndex *QModelIndex,role int) *QVariant {
	var __rv uintptr
	q.Drv(199000,199103,Native(proxyIndex),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractProxyModel::flags(QModelIndex const&)
func (q *QAbstractProxyModel) Flags(index *QModelIndex) Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(199000,199104,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractProxyModel::headerData(int,Qt::Orientation,int)
func (q *QAbstractProxyModel) HeaderData(section int,orientation Qt_Orientation,role int) *QVariant {
	var __rv uintptr
	q.Drv(199000,199105,unsafe.Pointer(&section),unsafe.Pointer(&orientation),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractProxyModel::itemData(QModelIndex const&)
func (q *QAbstractProxyModel) ItemData(index *QModelIndex) map[int]*QVariant {
	__rv := make(map[int]*QVariant)
	q.Drv(199000,199106,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractProxyModel::mapFromSource(QModelIndex const&)
func (q *QAbstractProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(199000,199107,Native(sourceIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractProxyModel::mapSelectionFromSource(QItemSelection const&)
func (q *QAbstractProxyModel) MapSelectionFromSource(selection *QItemSelection) *QItemSelection {
	var __rv uintptr
	q.Drv(199000,199108,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelection{}
	_rp.SetDriver(__rv,62000,true)
	return _rp
}	
//QAbstractProxyModel::mapSelectionToSource(QItemSelection const&)
func (q *QAbstractProxyModel) MapSelectionToSource(selection *QItemSelection) *QItemSelection {
	var __rv uintptr
	q.Drv(199000,199109,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelection{}
	_rp.SetDriver(__rv,62000,true)
	return _rp
}	
//QAbstractProxyModel::mapToSource(QModelIndex const&)
func (q *QAbstractProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(199000,199110,Native(proxyIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractProxyModel::revert()
func (q *QAbstractProxyModel) Revert()  {
	q.Drv(199000,199111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractProxyModel::setData(QModelIndex const&,QVariant const&,int)
func (q *QAbstractProxyModel) SetData(index *QModelIndex,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(199000,199112,Native(index),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractProxyModel::setHeaderData(int,Qt::Orientation,QVariant const&,int)
func (q *QAbstractProxyModel) SetHeaderData(section int,orientation Qt_Orientation,value *QVariant,role int) bool {
	var __rv bool
	q.Drv(199000,199113,unsafe.Pointer(&section),unsafe.Pointer(&orientation),Native(value),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractProxyModel::setSourceModel(QAbstractItemModel*)
func (q *QAbstractProxyModel) SetSourceModel(sourceModel QAbstractItemModelInterface)  {
	q.Drv(199000,199114,Native(sourceModel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractProxyModel::sourceModel()
func (q *QAbstractProxyModel) SourceModel() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(199000,199115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QAbstractProxyModel::submit()
func (q *QAbstractProxyModel) Submit() bool {
	var __rv bool
	q.Drv(199000,199116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
