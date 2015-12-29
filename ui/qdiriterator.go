// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDirIterator_IteratorFlag - QDirIterator::IteratorFlag
type QDirIterator_IteratorFlag uint32
const (
	QDirIterator_NoIteratorFlags QDirIterator_IteratorFlag = 0x0
	QDirIterator_FollowSymlinks QDirIterator_IteratorFlag = 0x1
	QDirIterator_Subdirectories QDirIterator_IteratorFlag = 0x2
)
//struct QDirIterator : QDirIterator
type QDirIterator struct {
	BaseDrv
}
//QDirIterator::QDirIterator(QDir const&,QFlags<QDirIterator::IteratorFlag>)	
func NewDirIterator(dir *QDir,flags QDirIterator_IteratorFlag) *QDirIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),23000,23102,Native(dir),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirIterator{}
	_p.SetDriver(__rv,23000,true)
	return _p
} 
//QDirIterator::QDirIterator(QString const&,QFlags<QDirIterator::IteratorFlag>)	
func NewDirIteratorWithPathFlags(path string,flags QDirIterator_IteratorFlag) *QDirIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),23000,23103,unsafe.Pointer(&path),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirIterator{}
	_p.SetDriver(__rv,23000,true)
	return _p
} 
//QDirIterator::QDirIterator(QString const&,QFlags<QDir::Filter>,QFlags<QDirIterator::IteratorFlag>)	
func NewDirIteratorWithPathFilterFlags(path string,filter QDir_Filter,flags QDirIterator_IteratorFlag) *QDirIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),23000,23104,unsafe.Pointer(&path),unsafe.Pointer(&filter),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirIterator{}
	_p.SetDriver(__rv,23000,true)
	return _p
} 
//QDirIterator::QDirIterator(QString const&,QStringList const&,QFlags<QDir::Filter>,QFlags<QDirIterator::IteratorFlag>)	
func NewDirIteratorWithPathNamefiltersFiltersFlags(path string,nameFilters []string,filters QDir_Filter,flags QDirIterator_IteratorFlag) *QDirIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),23000,23105,unsafe.Pointer(&path),unsafe.Pointer(&nameFilters),unsafe.Pointer(&filters),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDirIterator{}
	_p.SetDriver(__rv,23000,true)
	return _p
} 
//QDirIterator::fileInfo()
func (q *QDirIterator) FileInfo() *QFileInfo {
	var __rv uintptr
	q.Drv(23000,23106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFileInfo{}
	_rp.SetDriver(__rv,34000,true)
	return _rp
}	
//QDirIterator::fileName()
func (q *QDirIterator) FileName() string {
	var __rv string
	q.Drv(23000,23107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirIterator::filePath()
func (q *QDirIterator) FilePath() string {
	var __rv string
	q.Drv(23000,23108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirIterator::hasNext()
func (q *QDirIterator) HasNext() bool {
	var __rv bool
	q.Drv(23000,23109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirIterator::next()
func (q *QDirIterator) Next() string {
	var __rv string
	q.Drv(23000,23110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDirIterator::path()
func (q *QDirIterator) Path() string {
	var __rv string
	q.Drv(23000,23111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
