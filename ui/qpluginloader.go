// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPluginLoader : QPluginLoader
type QPluginLoader struct {
	QObject
}
//QPluginLoader::QPluginLoader()	
func NewPluginLoader() *QPluginLoader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),322000,322102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPluginLoader{}
	_p.SetDriver(__rv,322000,false)
	return _p
} 
//QPluginLoader::QPluginLoader(QObject*)	
func NewPluginLoaderWithParent(parent QObjectInterface) *QPluginLoader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),322000,322103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPluginLoader{}
	_p.SetDriver(__rv,322000,false)
	return _p
} 
//QPluginLoader::QPluginLoader(QString const&,QObject*)	
func NewPluginLoaderWithFilenameParent(fileName string,parent QObjectInterface) *QPluginLoader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),322000,322104,unsafe.Pointer(&fileName),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPluginLoader{}
	_p.SetDriver(__rv,322000,false)
	return _p
} 
//QPluginLoader::errorString()
func (q *QPluginLoader) ErrorString() string {
	var __rv string
	q.Drv(322000,322105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::fileName()
func (q *QPluginLoader) FileName() string {
	var __rv string
	q.Drv(322000,322106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::instance()
func (q *QPluginLoader) Instance() *QObject {
	var __rv uintptr
	q.Drv(322000,322107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QObject{}
	_rp.SetDriver(__rv,314000,false)
	return _rp
}	
//QPluginLoader::isLoaded()
func (q *QPluginLoader) IsLoaded() bool {
	var __rv bool
	q.Drv(322000,322108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::load()
func (q *QPluginLoader) Load() bool {
	var __rv bool
	q.Drv(322000,322109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::loadHints()
func (q *QPluginLoader) LoadHints() QLibrary_LoadHint {
	var __rv QLibrary_LoadHint
	q.Drv(322000,322110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::setFileName(QString const&)
func (q *QPluginLoader) SetFileName(fileName string)  {
	q.Drv(322000,322111,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPluginLoader::setLoadHints(QFlags<QLibrary::LoadHint>)
func (q *QPluginLoader) SetLoadHints(loadHints QLibrary_LoadHint)  {
	q.Drv(322000,322112,unsafe.Pointer(&loadHints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPluginLoader::staticInstances()	
func QPluginLoaderStaticInstances() []*QObject {
	var __rv []*QObject
	DirectQtDrv(nil,322000,322113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::staticInstances()
func (q *QPluginLoader) StaticInstances() []*QObject {
	var __rv []*QObject
	q.Drv(322000,322113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPluginLoader::unload()
func (q *QPluginLoader) Unload() bool {
	var __rv bool
	q.Drv(322000,322114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
