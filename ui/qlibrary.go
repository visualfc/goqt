// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLibrary_LoadHint - QLibrary::LoadHint
type QLibrary_LoadHint uint32
const (
	QLibrary_ResolveAllSymbolsHint QLibrary_LoadHint = 0x01
	QLibrary_ExportExternalSymbolsHint QLibrary_LoadHint = 0x02
	QLibrary_LoadArchiveMemberHint QLibrary_LoadHint = 0x04
)
//struct QLibrary : QLibrary
type QLibrary struct {
	QObject
}
//QLibrary::QLibrary()	
func NewLibrary() *QLibrary {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),301000,301102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLibrary{}
	_p.SetDriver(__rv,301000,false)
	return _p
} 
//QLibrary::QLibrary(QObject*)	
func NewLibraryWithParent(parent QObjectInterface) *QLibrary {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),301000,301103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLibrary{}
	_p.SetDriver(__rv,301000,false)
	return _p
} 
//QLibrary::QLibrary(QString const&,QObject*)	
func NewLibraryWithFilenameParent(fileName string,parent QObjectInterface) *QLibrary {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),301000,301104,unsafe.Pointer(&fileName),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLibrary{}
	_p.SetDriver(__rv,301000,false)
	return _p
} 
//QLibrary::QLibrary(QString const&,QString const&,QObject*)	
func NewLibraryWithFilenameVersionParent(fileName string,version string,parent QObjectInterface) *QLibrary {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),301000,301105,unsafe.Pointer(&fileName),unsafe.Pointer(&version),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLibrary{}
	_p.SetDriver(__rv,301000,false)
	return _p
} 
//QLibrary::QLibrary(QString const&,int,QObject*)	
func NewLibraryWithFilenameVernumParent(fileName string,verNum int,parent QObjectInterface) *QLibrary {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),301000,301106,unsafe.Pointer(&fileName),unsafe.Pointer(&verNum),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLibrary{}
	_p.SetDriver(__rv,301000,false)
	return _p
} 
//QLibrary::errorString()
func (q *QLibrary) ErrorString() string {
	var __rv string
	q.Drv(301000,301107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::fileName()
func (q *QLibrary) FileName() string {
	var __rv string
	q.Drv(301000,301108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::isLibrary(QString const&)	
func QLibraryIsLibrary(fileName string) bool {
	var __rv bool
	DirectQtDrv(nil,301000,301109,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::isLibrary(QString const&)
func (q *QLibrary) IsLibrary(fileName string) bool {
	var __rv bool
	q.Drv(301000,301109,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::isLoaded()
func (q *QLibrary) IsLoaded() bool {
	var __rv bool
	q.Drv(301000,301110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::load()
func (q *QLibrary) Load() bool {
	var __rv bool
	q.Drv(301000,301111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::loadHints()
func (q *QLibrary) LoadHints() QLibrary_LoadHint {
	var __rv QLibrary_LoadHint
	q.Drv(301000,301112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(char const*)
func (q *QLibrary) Resolve(symbol string) uintptr {
	var __rv uintptr
	q.Drv(301000,301113,unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,char const*)	
func QLibraryResolveWithFilenameSymbol(fileName string,symbol string) uintptr {
	var __rv uintptr
	DirectQtDrv(nil,301000,301114,unsafe.Pointer(&fileName),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,char const*)
func (q *QLibrary) ResolveWithFilenameSymbol(fileName string,symbol string) uintptr {
	var __rv uintptr
	q.Drv(301000,301114,unsafe.Pointer(&fileName),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,QString const&,char const*)	
func QLibraryResolveWithFilenameVersionSymbol(fileName string,version string,symbol string) uintptr {
	var __rv uintptr
	DirectQtDrv(nil,301000,301115,unsafe.Pointer(&fileName),unsafe.Pointer(&version),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,QString const&,char const*)
func (q *QLibrary) ResolveWithFilenameVersionSymbol(fileName string,version string,symbol string) uintptr {
	var __rv uintptr
	q.Drv(301000,301115,unsafe.Pointer(&fileName),unsafe.Pointer(&version),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,int,char const*)	
func QLibraryResolveWithFilenameVernumSymbol(fileName string,verNum int,symbol string) uintptr {
	var __rv uintptr
	DirectQtDrv(nil,301000,301116,unsafe.Pointer(&fileName),unsafe.Pointer(&verNum),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::resolve(QString const&,int,char const*)
func (q *QLibrary) ResolveWithFilenameVernumSymbol(fileName string,verNum int,symbol string) uintptr {
	var __rv uintptr
	q.Drv(301000,301116,unsafe.Pointer(&fileName),unsafe.Pointer(&verNum),unsafe.Pointer(&symbol),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibrary::setFileName(QString const&)
func (q *QLibrary) SetFileName(fileName string)  {
	q.Drv(301000,301117,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLibrary::setFileNameAndVersion(QString const&,QString const&)
func (q *QLibrary) SetFileNameAndVersionWithFilenameVersion(fileName string,version string)  {
	q.Drv(301000,301118,unsafe.Pointer(&fileName),unsafe.Pointer(&version),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLibrary::setFileNameAndVersion(QString const&,int)
func (q *QLibrary) SetFileNameAndVersionWithFilenameVernum(fileName string,verNum int)  {
	q.Drv(301000,301119,unsafe.Pointer(&fileName),unsafe.Pointer(&verNum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLibrary::setLoadHints(QFlags<QLibrary::LoadHint>)
func (q *QLibrary) SetLoadHints(hints QLibrary_LoadHint)  {
	q.Drv(301000,301120,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLibrary::unload()
func (q *QLibrary) Unload() bool {
	var __rv bool
	q.Drv(301000,301121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
