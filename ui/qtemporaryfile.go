// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTemporaryFile : QTemporaryFile
type QTemporaryFile struct {
	QFile
}
//QTemporaryFile::QTemporaryFile()	
func NewTemporaryFile() *QTemporaryFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),369000,369102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTemporaryFile{}
	_p.SetDriver(__rv,369000,false)
	return _p
} 
//QTemporaryFile::QTemporaryFile(QObject*)	
func NewTemporaryFileWithParent(parent QObjectInterface) *QTemporaryFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),369000,369103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTemporaryFile{}
	_p.SetDriver(__rv,369000,false)
	return _p
} 
//QTemporaryFile::QTemporaryFile(QString const&)	
func NewTemporaryFileWithTemplatename(templateName string) *QTemporaryFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),369000,369104,unsafe.Pointer(&templateName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTemporaryFile{}
	_p.SetDriver(__rv,369000,false)
	return _p
} 
//QTemporaryFile::QTemporaryFile(QString const&,QObject*)	
func NewTemporaryFileWithTemplatenameParent(templateName string,parent QObjectInterface) *QTemporaryFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),369000,369105,unsafe.Pointer(&templateName),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTemporaryFile{}
	_p.SetDriver(__rv,369000,false)
	return _p
} 
//QTemporaryFile::autoRemove()
func (q *QTemporaryFile) AutoRemove() bool {
	var __rv bool
	q.Drv(369000,369106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTemporaryFile::createLocalFile(QFile&)	
func QTemporaryFileCreateLocalFile(file *QFile) *QTemporaryFile {
	var __rv uintptr
	DirectQtDrv(nil,369000,369107,Native(file),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTemporaryFile{}
	_rp.SetDriver(__rv,369000,false)
	return _rp
}	
//QTemporaryFile::createLocalFile(QFile&)
func (q *QTemporaryFile) CreateLocalFile(file *QFile) *QTemporaryFile {
	var __rv uintptr
	q.Drv(369000,369107,Native(file),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTemporaryFile{}
	_rp.SetDriver(__rv,369000,false)
	return _rp
}	
//QTemporaryFile::createLocalFile(QString const&)	
func QTemporaryFileCreateLocalFileWithFilename(fileName string) *QTemporaryFile {
	var __rv uintptr
	DirectQtDrv(nil,369000,369108,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTemporaryFile{}
	_rp.SetDriver(__rv,369000,false)
	return _rp
}	
//QTemporaryFile::createLocalFile(QString const&)
func (q *QTemporaryFile) CreateLocalFileWithFilename(fileName string) *QTemporaryFile {
	var __rv uintptr
	q.Drv(369000,369108,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTemporaryFile{}
	_rp.SetDriver(__rv,369000,false)
	return _rp
}	
//QTemporaryFile::fileName()
func (q *QTemporaryFile) FileName() string {
	var __rv string
	q.Drv(369000,369109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTemporaryFile::fileTemplate()
func (q *QTemporaryFile) FileTemplate() string {
	var __rv string
	q.Drv(369000,369110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTemporaryFile::open()
func (q *QTemporaryFile) Open() bool {
	var __rv bool
	q.Drv(369000,369111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTemporaryFile::open(QFlags<QIODevice::OpenModeFlag>)
func (q *QTemporaryFile) OpenWithFlags(flags QIODevice_OpenModeFlag) bool {
	var __rv bool
	q.Drv(369000,369112,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTemporaryFile::setAutoRemove(bool)
func (q *QTemporaryFile) SetAutoRemove(b bool)  {
	q.Drv(369000,369113,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTemporaryFile::setFileTemplate(QString const&)
func (q *QTemporaryFile) SetFileTemplate(name string)  {
	q.Drv(369000,369114,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
