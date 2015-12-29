// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPersistentModelIndex : QPersistentModelIndex
type QPersistentModelIndex struct {
	BaseDrv
}
//QPersistentModelIndex::QPersistentModelIndex()	
func NewPersistentModelIndex() *QPersistentModelIndex {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),93000,93102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPersistentModelIndex{}
	_p.SetDriver(__rv,93000,true)
	return _p
} 
//QPersistentModelIndex::QPersistentModelIndex(QModelIndex const&)	
func NewPersistentModelIndexWithIndex(index *QModelIndex) *QPersistentModelIndex {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),93000,93103,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPersistentModelIndex{}
	_p.SetDriver(__rv,93000,true)
	return _p
} 
//QPersistentModelIndex::QPersistentModelIndex(QPersistentModelIndex const&)	
func NewPersistentModelIndexCopy(other *QPersistentModelIndex) *QPersistentModelIndex {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),93000,93104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPersistentModelIndex{}
	_p.SetDriver(__rv,93000,true)
	return _p
} 
//QPersistentModelIndex::child(int,int)
func (q *QPersistentModelIndex) Child(row int,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(93000,93105,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QPersistentModelIndex::column()
func (q *QPersistentModelIndex) Column() int {
	var __rv int
	q.Drv(93000,93106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::data()
func (q *QPersistentModelIndex) Data() *QVariant {
	var __rv uintptr
	q.Drv(93000,93107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QPersistentModelIndex::data(int)
func (q *QPersistentModelIndex) DataWithRole(role int) *QVariant {
	var __rv uintptr
	q.Drv(93000,93108,unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QPersistentModelIndex::flags()
func (q *QPersistentModelIndex) Flags() Qt_ItemFlag {
	var __rv Qt_ItemFlag
	q.Drv(93000,93109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::internalId()
func (q *QPersistentModelIndex) InternalId() int64 {
	var __rv int64
	q.Drv(93000,93110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::internalPointer()
func (q *QPersistentModelIndex) InternalPointer() uintptr {
	var __rv uintptr
	q.Drv(93000,93111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::isValid()
func (q *QPersistentModelIndex) IsValid() bool {
	var __rv bool
	q.Drv(93000,93112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::model()
func (q *QPersistentModelIndex) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(93000,93113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QPersistentModelIndex::parent()
func (q *QPersistentModelIndex) Parent() *QModelIndex {
	var __rv uintptr
	q.Drv(93000,93114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QPersistentModelIndex::row()
func (q *QPersistentModelIndex) Row() int {
	var __rv int
	q.Drv(93000,93115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPersistentModelIndex::sibling(int,int)
func (q *QPersistentModelIndex) Sibling(row int,column int) *QModelIndex {
	var __rv uintptr
	q.Drv(93000,93116,unsafe.Pointer(&row),unsafe.Pointer(&column),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
