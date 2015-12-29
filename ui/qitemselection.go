// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QItemSelection : QItemSelection
type QItemSelection struct {
	BaseDrv
}
//QItemSelection::QItemSelection()	
func NewItemSelection() *QItemSelection {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),62000,62102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelection{}
	_p.SetDriver(__rv,62000,true)
	return _p
} 
//QItemSelection::QItemSelection(QModelIndex const&,QModelIndex const&)	
func NewItemSelectionWithTopleftBottomright(topLeft *QModelIndex,bottomRight *QModelIndex) *QItemSelection {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),62000,62103,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelection{}
	_p.SetDriver(__rv,62000,true)
	return _p
} 
//QItemSelection::contains(QModelIndex const&)
func (q *QItemSelection) Contains(index *QModelIndex) bool {
	var __rv bool
	q.Drv(62000,62104,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelection::indexes()
func (q *QItemSelection) Indexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(62000,62105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelection::merge(QItemSelection const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QItemSelection) Merge(other *QItemSelection,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(62000,62106,Native(other),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelection::select(QModelIndex const&,QModelIndex const&)
func (q *QItemSelection) Select(topLeft *QModelIndex,bottomRight *QModelIndex)  {
	q.Drv(62000,62107,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelection::split(QItemSelectionRange const&,QItemSelectionRange const&,QItemSelection*)	
func QItemSelectionSplit(_range *QItemSelectionRange,other *QItemSelectionRange,result *QItemSelection)  {
	DirectQtDrv(nil,62000,62108,Native(_range),Native(other),Native(result),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelection::split(QItemSelectionRange const&,QItemSelectionRange const&,QItemSelection*)
func (q *QItemSelection) Split(_range *QItemSelectionRange,other *QItemSelectionRange,result *QItemSelection)  {
	q.Drv(62000,62108,Native(_range),Native(other),Native(result),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
