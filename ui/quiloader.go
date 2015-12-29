// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QUiLoader : QUiLoader
type QUiLoader struct {
	QObject
}
//QUiLoader::QUiLoader()	
func NewUiLoader() *QUiLoader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),387000,387102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUiLoader{}
	_p.SetDriver(__rv,387000,false)
	return _p
} 
//QUiLoader::QUiLoader(QObject*)	
func NewUiLoaderWithParent(parent QObjectInterface) *QUiLoader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),387000,387103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QUiLoader{}
	_p.SetDriver(__rv,387000,false)
	return _p
} 
//QUiLoader::addPluginPath(QString const&)
func (q *QUiLoader) AddPluginPath(path string)  {
	q.Drv(387000,387104,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUiLoader::availableLayouts()
func (q *QUiLoader) AvailableLayouts() []string {
	var __rv []string
	q.Drv(387000,387105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUiLoader::availableWidgets()
func (q *QUiLoader) AvailableWidgets() []string {
	var __rv []string
	q.Drv(387000,387106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUiLoader::clearPluginPaths()
func (q *QUiLoader) ClearPluginPaths()  {
	q.Drv(387000,387107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUiLoader::createAction(QObject*)
func (q *QUiLoader) CreateAction(parent QObjectInterface) *QAction {
	var __rv uintptr
	q.Drv(387000,387108,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUiLoader::createAction(QObject*,QString const&)
func (q *QUiLoader) CreateActionWithParentName(parent QObjectInterface,name string) *QAction {
	var __rv uintptr
	q.Drv(387000,387109,Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QUiLoader::createActionGroup()
func (q *QUiLoader) CreateActionGroup() *QActionGroup {
	var __rv uintptr
	q.Drv(387000,387110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QActionGroup{}
	_rp.SetDriver(__rv,208000,false)
	return _rp
}	
//QUiLoader::createActionGroup(QObject*,QString const&)
func (q *QUiLoader) CreateActionGroupWithParentName(parent QObjectInterface,name string) *QActionGroup {
	var __rv uintptr
	q.Drv(387000,387111,Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QActionGroup{}
	_rp.SetDriver(__rv,208000,false)
	return _rp
}	
//QUiLoader::createLayout(QString const&)
func (q *QUiLoader) CreateLayout(className string) *QLayout {
	var __rv uintptr
	q.Drv(387000,387112,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayout{}
	_rp.SetDriver(__rv,300000,false)
	return _rp
}	
//QUiLoader::createLayout(QString const&,QObject*,QString const&)
func (q *QUiLoader) CreateLayoutWithClassnameParentName(className string,parent QObjectInterface,name string) *QLayout {
	var __rv uintptr
	q.Drv(387000,387113,unsafe.Pointer(&className),Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayout{}
	_rp.SetDriver(__rv,300000,false)
	return _rp
}	
//QUiLoader::createWidget(QString const&)
func (q *QUiLoader) CreateWidget(className string) *QWidget {
	var __rv uintptr
	q.Drv(387000,387114,unsafe.Pointer(&className),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QUiLoader::createWidget(QString const&,QWidget*,QString const&)
func (q *QUiLoader) CreateWidgetWithClassnameParentName(className string,parent QWidgetInterface,name string) *QWidget {
	var __rv uintptr
	q.Drv(387000,387115,unsafe.Pointer(&className),Native(parent),unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QUiLoader::isLanguageChangeEnabled()
func (q *QUiLoader) IsLanguageChangeEnabled() bool {
	var __rv bool
	q.Drv(387000,387116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUiLoader::load(QIODevice*)
func (q *QUiLoader) Load(device QIODeviceInterface) *QWidget {
	var __rv uintptr
	q.Drv(387000,387117,Native(device),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QUiLoader::load(QIODevice*,QWidget*)
func (q *QUiLoader) LoadWithDeviceWidget(device QIODeviceInterface,parentWidget QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(387000,387118,Native(device),Native(parentWidget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QUiLoader::pluginPaths()
func (q *QUiLoader) PluginPaths() []string {
	var __rv []string
	q.Drv(387000,387119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QUiLoader::setLanguageChangeEnabled(bool)
func (q *QUiLoader) SetLanguageChangeEnabled(enabled bool)  {
	q.Drv(387000,387120,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUiLoader::setWorkingDirectory(QDir const&)
func (q *QUiLoader) SetWorkingDirectory(dir *QDir)  {
	q.Drv(387000,387121,Native(dir),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QUiLoader::workingDirectory()
func (q *QUiLoader) WorkingDirectory() *QDir {
	var __rv uintptr
	q.Drv(387000,387122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
