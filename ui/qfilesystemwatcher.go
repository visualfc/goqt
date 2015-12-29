// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFileSystemWatcher : QFileSystemWatcher
type QFileSystemWatcher struct {
	QObject
}
func (q *QFileSystemWatcher) OnFileChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(239000,239102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QFileSystemWatcher) OnDirectoryChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(239000,239103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFileSystemWatcher::QFileSystemWatcher()	
func NewFileSystemWatcher() *QFileSystemWatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),239000,239104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileSystemWatcher{}
	_p.SetDriver(__rv,239000,false)
	return _p
} 
//QFileSystemWatcher::QFileSystemWatcher(QObject*)	
func NewFileSystemWatcherWithParent(parent QObjectInterface) *QFileSystemWatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),239000,239105,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileSystemWatcher{}
	_p.SetDriver(__rv,239000,false)
	return _p
} 
//QFileSystemWatcher::QFileSystemWatcher(QStringList const&,QObject*)	
func NewFileSystemWatcherWithPathsParent(paths []string,parent QObjectInterface) *QFileSystemWatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),239000,239106,unsafe.Pointer(&paths),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileSystemWatcher{}
	_p.SetDriver(__rv,239000,false)
	return _p
} 
//QFileSystemWatcher::addPath(QString const&)
func (q *QFileSystemWatcher) AddPath(file string)  {
	q.Drv(239000,239107,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemWatcher::addPaths(QStringList const&)
func (q *QFileSystemWatcher) AddPaths(files []string)  {
	q.Drv(239000,239108,unsafe.Pointer(&files),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemWatcher::directories()
func (q *QFileSystemWatcher) Directories() []string {
	var __rv []string
	q.Drv(239000,239109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemWatcher::files()
func (q *QFileSystemWatcher) Files() []string {
	var __rv []string
	q.Drv(239000,239110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileSystemWatcher::removePath(QString const&)
func (q *QFileSystemWatcher) RemovePath(file string)  {
	q.Drv(239000,239111,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileSystemWatcher::removePaths(QStringList const&)
func (q *QFileSystemWatcher) RemovePaths(files []string)  {
	q.Drv(239000,239112,unsafe.Pointer(&files),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
