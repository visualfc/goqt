// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDesktopServices_StandardLocation - QDesktopServices::StandardLocation
type QDesktopServices_StandardLocation uint32
const (
	QDesktopServices_DesktopLocation QDesktopServices_StandardLocation = 0
	QDesktopServices_DocumentsLocation QDesktopServices_StandardLocation = 1
	QDesktopServices_FontsLocation QDesktopServices_StandardLocation = 2
	QDesktopServices_ApplicationsLocation QDesktopServices_StandardLocation = 3
	QDesktopServices_MusicLocation QDesktopServices_StandardLocation = 4
	QDesktopServices_MoviesLocation QDesktopServices_StandardLocation = 5
	QDesktopServices_PicturesLocation QDesktopServices_StandardLocation = 6
	QDesktopServices_TempLocation QDesktopServices_StandardLocation = 7
	QDesktopServices_HomeLocation QDesktopServices_StandardLocation = 8
	QDesktopServices_DataLocation QDesktopServices_StandardLocation = 9
	QDesktopServices_CacheLocation QDesktopServices_StandardLocation = 10
)
//struct QDesktopServices : QDesktopServices
type QDesktopServices struct {
	BaseDrv
}
//QDesktopServices::QDesktopServices()	
func NewDesktopServices() *QDesktopServices {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),21000,21102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDesktopServices{}
	_p.SetDriver(__rv,21000,true)
	return _p
} 
//QDesktopServices::displayName(QDesktopServices::StandardLocation)	
func QDesktopServicesDisplayName(_type QDesktopServices_StandardLocation) string {
	var __rv string
	DirectQtDrv(nil,21000,21103,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::displayName(QDesktopServices::StandardLocation)
func (q *QDesktopServices) DisplayName(_type QDesktopServices_StandardLocation) string {
	var __rv string
	q.Drv(21000,21103,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::openUrl(QUrl const&)	
func QDesktopServicesOpenUrl(url *QUrl) bool {
	var __rv bool
	DirectQtDrv(nil,21000,21104,Native(url),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::openUrl(QUrl const&)
func (q *QDesktopServices) OpenUrl(url *QUrl) bool {
	var __rv bool
	q.Drv(21000,21104,Native(url),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::setUrlHandler(QString const&,QObject*,char const*)	
func QDesktopServicesSetUrlHandler(scheme string,receiver QObjectInterface,method string)  {
	DirectQtDrv(nil,21000,21105,unsafe.Pointer(&scheme),Native(receiver),unsafe.Pointer(&method),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDesktopServices::setUrlHandler(QString const&,QObject*,char const*)
func (q *QDesktopServices) SetUrlHandler(scheme string,receiver QObjectInterface,method string)  {
	q.Drv(21000,21105,unsafe.Pointer(&scheme),Native(receiver),unsafe.Pointer(&method),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDesktopServices::storageLocation(QDesktopServices::StandardLocation)	
func QDesktopServicesStorageLocation(_type QDesktopServices_StandardLocation) string {
	var __rv string
	DirectQtDrv(nil,21000,21106,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::storageLocation(QDesktopServices::StandardLocation)
func (q *QDesktopServices) StorageLocation(_type QDesktopServices_StandardLocation) string {
	var __rv string
	q.Drv(21000,21106,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopServices::unsetUrlHandler(QString const&)	
func QDesktopServicesUnsetUrlHandler(scheme string)  {
	DirectQtDrv(nil,21000,21107,unsafe.Pointer(&scheme),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDesktopServices::unsetUrlHandler(QString const&)
func (q *QDesktopServices) UnsetUrlHandler(scheme string)  {
	q.Drv(21000,21107,unsafe.Pointer(&scheme),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
