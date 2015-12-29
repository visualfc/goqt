// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QItemSelectionModel_SelectionFlag - QItemSelectionModel::SelectionFlag
type QItemSelectionModel_SelectionFlag uint32
const (
	QItemSelectionModel_NoUpdate QItemSelectionModel_SelectionFlag = 0x0000
	QItemSelectionModel_Clear QItemSelectionModel_SelectionFlag = 0x0001
	QItemSelectionModel_Select QItemSelectionModel_SelectionFlag = 0x0002
	QItemSelectionModel_Deselect QItemSelectionModel_SelectionFlag = 0x0004
	QItemSelectionModel_Toggle QItemSelectionModel_SelectionFlag = 0x0008
	QItemSelectionModel_Current QItemSelectionModel_SelectionFlag = 0x0010
	QItemSelectionModel_Rows QItemSelectionModel_SelectionFlag = 0x0020
	QItemSelectionModel_Columns QItemSelectionModel_SelectionFlag = 0x0040
	QItemSelectionModel_SelectCurrent QItemSelectionModel_SelectionFlag = QItemSelectionModel_Select|QItemSelectionModel_Current
	QItemSelectionModel_ToggleCurrent QItemSelectionModel_SelectionFlag = QItemSelectionModel_Toggle|QItemSelectionModel_Current
	QItemSelectionModel_ClearAndSelect QItemSelectionModel_SelectionFlag = QItemSelectionModel_Clear|QItemSelectionModel_Select
)
//struct QItemSelectionModel : QItemSelectionModel
type QItemSelectionModel struct {
	QObject
}
func (q *QItemSelectionModel) OnCurrentRowChanged(fn func(*QModelIndex,*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(296000,296102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QItemSelectionModel) OnSelectionChanged(fn func(*QItemSelection,*QItemSelection)) uintptr {
	var __rv uintptr
	q.Drv(296000,296103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QItemSelectionModel) OnCurrentChanged(fn func(*QModelIndex,*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(296000,296104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QItemSelectionModel) OnCurrentColumnChanged(fn func(*QModelIndex,*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(296000,296105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QItemSelectionModel::QItemSelectionModel(QAbstractItemModel*)	
func NewItemSelectionModel(model QAbstractItemModelInterface) *QItemSelectionModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),296000,296106,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionModel{}
	_p.SetDriver(__rv,296000,false)
	return _p
} 
//QItemSelectionModel::QItemSelectionModel(QAbstractItemModel*,QObject*)	
func NewItemSelectionModelWithModelParent(model QAbstractItemModelInterface,parent QObjectInterface) *QItemSelectionModel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),296000,296107,Native(model),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionModel{}
	_p.SetDriver(__rv,296000,false)
	return _p
} 
//QItemSelectionModel::clear()
func (q *QItemSelectionModel) Clear()  {
	q.Drv(296000,296108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::clearSelection()
func (q *QItemSelectionModel) ClearSelection()  {
	q.Drv(296000,296109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::columnIntersectsSelection(int,QModelIndex const&)
func (q *QItemSelectionModel) ColumnIntersectsSelection(column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(296000,296110,unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::currentIndex()
func (q *QItemSelectionModel) CurrentIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(296000,296111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QItemSelectionModel::emitSelectionChanged(QItemSelection const&,QItemSelection const&)
func (q *QItemSelectionModel) EmitSelectionChanged(newSelection *QItemSelection,oldSelection *QItemSelection)  {
	q.Drv(296000,296112,Native(newSelection),Native(oldSelection),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::hasSelection()
func (q *QItemSelectionModel) HasSelection() bool {
	var __rv bool
	q.Drv(296000,296113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::isColumnSelected(int,QModelIndex const&)
func (q *QItemSelectionModel) IsColumnSelected(column int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(296000,296114,unsafe.Pointer(&column),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::isRowSelected(int,QModelIndex const&)
func (q *QItemSelectionModel) IsRowSelected(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(296000,296115,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::isSelected(QModelIndex const&)
func (q *QItemSelectionModel) IsSelected(index *QModelIndex) bool {
	var __rv bool
	q.Drv(296000,296116,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::model()
func (q *QItemSelectionModel) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(296000,296117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QItemSelectionModel::reset()
func (q *QItemSelectionModel) Reset()  {
	q.Drv(296000,296118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::rowIntersectsSelection(int,QModelIndex const&)
func (q *QItemSelectionModel) RowIntersectsSelection(row int,parent *QModelIndex) bool {
	var __rv bool
	q.Drv(296000,296119,unsafe.Pointer(&row),Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::select(QItemSelection const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QItemSelectionModel) SelectWithSelectionCommand(selection *QItemSelection,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(296000,296120,Native(selection),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::select(QModelIndex const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QItemSelectionModel) SelectWithIndexCommand(index *QModelIndex,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(296000,296121,Native(index),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemSelectionModel::selectedColumns()
func (q *QItemSelectionModel) SelectedColumns() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(296000,296122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::selectedColumns(int)
func (q *QItemSelectionModel) SelectedColumnsWithRow(row int) []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(296000,296123,unsafe.Pointer(&row),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::selectedIndexes()
func (q *QItemSelectionModel) SelectedIndexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(296000,296124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::selectedRows()
func (q *QItemSelectionModel) SelectedRows() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(296000,296125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::selectedRows(int)
func (q *QItemSelectionModel) SelectedRowsWithColumn(column int) []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(296000,296126,unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionModel::selection()
func (q *QItemSelectionModel) Selection() *QItemSelection {
	var __rv uintptr
	q.Drv(296000,296127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelection{}
	_rp.SetDriver(__rv,62000,true)
	return _rp
}	
//QItemSelectionModel::setCurrentIndex(QModelIndex const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QItemSelectionModel) SetCurrentIndex(index *QModelIndex,command QItemSelectionModel_SelectionFlag)  {
	q.Drv(296000,296128,Native(index),unsafe.Pointer(&command),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
