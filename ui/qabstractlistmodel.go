// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractListModel : QAbstractListModel
type QAbstractListModel struct {
	QAbstractItemModel
}
//QAbstractListModel::dropMimeData(QMimeData const*,Qt::DropAction,int,int,QModelIndex const&)
func (q *QAbstractListModel) DropMimeData(data *QMimeData,action Qt_DropAction,row int,column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(197000,197102,Native(data),unsafe.Pointer(&action),unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractListModel::index(int)
func (q *QAbstractListModel) Index(row int) *QModelIndex {
	var __rv uintptr
	q.Drv(197000,197103,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QAbstractListModel::index(int,int,QModelIndex const&)
func (q *QAbstractListModel) IndexWithRowColumnParent(row int,column int,parent *QModelIndex) *QModelIndex {
	var __rv uintptr
	q.Drv(197000,197104,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
