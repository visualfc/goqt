// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QModelIndex : QModelIndex
type QModelIndex struct {
	BaseDrv
}
//QModelIndex::QModelIndex()	
func NewModelIndex() *QModelIndex {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),79000,79102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QModelIndex{}
	_p.SetDriver(__rv,79000,true)
	return _p
} 
//QModelIndex::QModelIndex(QModelIndex const&)	
func NewModelIndexCopy(other *QModelIndex) *QModelIndex {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),79000,79103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QModelIndex{}
	_p.SetDriver(__rv,79000,true)
	return _p
} 
//QModelIndex::child(int,int)
func (q *QModelIndex) Child(row int,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(79000,79104,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QModelIndex::column()
func (q *QModelIndex) Column() int {
	var __rv int
	q.Drv(79000,79105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::data()
func (q *QModelIndex) Data() *QVariant {
	var __rv uintptr
	q.Drv(79000,79106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QModelIndex::data(int)
func (q *QModelIndex) DataWithRole(role int) *QVariant {
	var __rv uintptr
	q.Drv(79000,79107,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QModelIndex::flags()
func (q *QModelIndex) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(79000,79108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::internalId()
func (q *QModelIndex) InternalId() int64 {
	var __rv int64
	q.Drv(79000,79109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::internalPointer()
func (q *QModelIndex) InternalPointer() uintptr {
	var __rv uintptr
	q.Drv(79000,79110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::isValid()
func (q *QModelIndex) IsValid() bool {
	var __rv bool
	q.Drv(79000,79111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::model()
func (q *QModelIndex) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(79000,79112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QModelIndex::parent()
func (q *QModelIndex) Parent() *QModelIndex {
	var __rv uintptr
	q.Drv(79000,79113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QModelIndex::row()
func (q *QModelIndex) Row() int {
	var __rv int
	q.Drv(79000,79114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QModelIndex::sibling(int,int)
func (q *QModelIndex) Sibling(row int,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(79000,79115,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
