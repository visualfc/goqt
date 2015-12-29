// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QSharedMemory_SharedMemoryError - QSharedMemory::SharedMemoryError
type QSharedMemory_SharedMemoryError uint32
const (
	QSharedMemory_NoError QSharedMemory_SharedMemoryError = 0
	QSharedMemory_PermissionDenied QSharedMemory_SharedMemoryError = 1
	QSharedMemory_InvalidSize QSharedMemory_SharedMemoryError = 2
	QSharedMemory_KeyError QSharedMemory_SharedMemoryError = 3
	QSharedMemory_AlreadyExists QSharedMemory_SharedMemoryError = 4
	QSharedMemory_NotFound QSharedMemory_SharedMemoryError = 5
	QSharedMemory_LockError QSharedMemory_SharedMemoryError = 6
	QSharedMemory_OutOfResources QSharedMemory_SharedMemoryError = 7
	QSharedMemory_UnknownError QSharedMemory_SharedMemoryError = 8
)
//enum QSharedMemory_AccessMode - QSharedMemory::AccessMode
type QSharedMemory_AccessMode uint32
const (
	QSharedMemory_ReadOnly QSharedMemory_AccessMode = 0
	QSharedMemory_ReadWrite QSharedMemory_AccessMode = 1
)
//struct QSharedMemory : QSharedMemory
type QSharedMemory struct {
	QObject
}
//QSharedMemory::QSharedMemory()	
func NewSharedMemory() *QSharedMemory {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),340000,340102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSharedMemory{}
	_p.SetDriver(__rv,340000,false)
	return _p
} 
//QSharedMemory::QSharedMemory(QObject*)	
func NewSharedMemoryWithParent(parent QObjectInterface) *QSharedMemory {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),340000,340103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSharedMemory{}
	_p.SetDriver(__rv,340000,false)
	return _p
} 
//QSharedMemory::QSharedMemory(QString const&,QObject*)	
func NewSharedMemoryWithKeyParent(key string,parent QObjectInterface) *QSharedMemory {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),340000,340104,unsafe.Pointer(&key),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSharedMemory{}
	_p.SetDriver(__rv,340000,false)
	return _p
} 
//QSharedMemory::attach()
func (q *QSharedMemory) Attach() bool {
	var __rv bool
	q.Drv(340000,340105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::attach(QSharedMemory::AccessMode)
func (q *QSharedMemory) AttachWithMode(mode QSharedMemory_AccessMode) bool {
	var __rv bool
	q.Drv(340000,340106,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::constData()
func (q *QSharedMemory) ConstData() uintptr {
	var __rv uintptr
	q.Drv(340000,340107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::create(int)
func (q *QSharedMemory) Create(size int) bool {
	var __rv bool
	q.Drv(340000,340108,unsafe.Pointer(&size),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::create(int,QSharedMemory::AccessMode)
func (q *QSharedMemory) CreateWithSizeMode(size int,mode QSharedMemory_AccessMode) bool {
	var __rv bool
	q.Drv(340000,340109,unsafe.Pointer(&size),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::data()
func (q *QSharedMemory) Data() uintptr {
	var __rv uintptr
	q.Drv(340000,340110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::detach()
func (q *QSharedMemory) Detach() bool {
	var __rv bool
	q.Drv(340000,340111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::error()
func (q *QSharedMemory) Error() QSharedMemory_SharedMemoryError {
	var __rv QSharedMemory_SharedMemoryError
	q.Drv(340000,340112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::errorString()
func (q *QSharedMemory) ErrorString() string {
	var __rv string
	q.Drv(340000,340113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::isAttached()
func (q *QSharedMemory) IsAttached() bool {
	var __rv bool
	q.Drv(340000,340114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::key()
func (q *QSharedMemory) Key() string {
	var __rv string
	q.Drv(340000,340115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::lock()
func (q *QSharedMemory) Lock() bool {
	var __rv bool
	q.Drv(340000,340116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::setKey(QString const&)
func (q *QSharedMemory) SetKey(key string)  {
	q.Drv(340000,340117,unsafe.Pointer(&key),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSharedMemory::size()
func (q *QSharedMemory) Size() int {
	var __rv int
	q.Drv(340000,340118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSharedMemory::unlock()
func (q *QSharedMemory) Unlock() bool {
	var __rv bool
	q.Drv(340000,340119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
