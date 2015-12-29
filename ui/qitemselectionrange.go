// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QItemSelectionRange : QItemSelectionRange
type QItemSelectionRange struct {
	BaseDrv
}
//QItemSelectionRange::QItemSelectionRange()	
func NewItemSelectionRange() *QItemSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),63000,63102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionRange{}
	_p.SetDriver(__rv,63000,true)
	return _p
} 
//QItemSelectionRange::QItemSelectionRange(QItemSelectionRange const&)	
func NewItemSelectionRangeCopy(other *QItemSelectionRange) *QItemSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),63000,63103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionRange{}
	_p.SetDriver(__rv,63000,true)
	return _p
} 
//QItemSelectionRange::QItemSelectionRange(QModelIndex const&)	
func NewItemSelectionRangeWithIndex(index *QModelIndex) *QItemSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),63000,63104,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionRange{}
	_p.SetDriver(__rv,63000,true)
	return _p
} 
//QItemSelectionRange::QItemSelectionRange(QModelIndex const&,QModelIndex const&)	
func NewItemSelectionRangeWithTopleftBottomright(topLeft *QModelIndex,bottomRight *QModelIndex) *QItemSelectionRange {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),63000,63105,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemSelectionRange{}
	_p.SetDriver(__rv,63000,true)
	return _p
} 
//QItemSelectionRange::bottom()
func (q *QItemSelectionRange) Bottom() int {
	var __rv int
	q.Drv(63000,63106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::bottomRight()
func (q *QItemSelectionRange) BottomRight() *QModelIndex {
	var __rv uintptr
	q.Drv(63000,63107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QItemSelectionRange::contains(QModelIndex const&)
func (q *QItemSelectionRange) Contains(index *QModelIndex) bool {
	var __rv bool
	q.Drv(63000,63108,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::contains(int,int,QModelIndex const&)
func (q *QItemSelectionRange) ContainsWithRowColumnParentindex(row int,column int,parentIndex *QModelIndex) bool {
	var __rv bool
	q.Drv(63000,63109,unsafe.Pointer(&row),unsafe.Pointer(&column),Native(parentIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::height()
func (q *QItemSelectionRange) Height() int {
	var __rv int
	q.Drv(63000,63110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::indexes()
func (q *QItemSelectionRange) Indexes() []*QModelIndex {
	var __rv []*QModelIndex
	q.Drv(63000,63111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::intersected(QItemSelectionRange const&)
func (q *QItemSelectionRange) Intersected(other *QItemSelectionRange) *QItemSelectionRange {
	var __rv uintptr
	q.Drv(63000,63112,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemSelectionRange{}
	_rp.SetDriver(__rv,63000,true)
	return _rp
}	
//QItemSelectionRange::intersects(QItemSelectionRange const&)
func (q *QItemSelectionRange) Intersects(other *QItemSelectionRange) bool {
	var __rv bool
	q.Drv(63000,63113,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::isEmpty()
func (q *QItemSelectionRange) IsEmpty() bool {
	var __rv bool
	q.Drv(63000,63114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::isValid()
func (q *QItemSelectionRange) IsValid() bool {
	var __rv bool
	q.Drv(63000,63115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::left()
func (q *QItemSelectionRange) Left() int {
	var __rv int
	q.Drv(63000,63116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::model()
func (q *QItemSelectionRange) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(63000,63117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QItemSelectionRange::parent()
func (q *QItemSelectionRange) Parent() *QModelIndex {
	var __rv uintptr
	q.Drv(63000,63118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QItemSelectionRange::right()
func (q *QItemSelectionRange) Right() int {
	var __rv int
	q.Drv(63000,63119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::top()
func (q *QItemSelectionRange) Top() int {
	var __rv int
	q.Drv(63000,63120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemSelectionRange::topLeft()
func (q *QItemSelectionRange) TopLeft() *QModelIndex {
	var __rv uintptr
	q.Drv(63000,63121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QItemSelectionRange::width()
func (q *QItemSelectionRange) Width() int {
	var __rv int
	q.Drv(63000,63122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
