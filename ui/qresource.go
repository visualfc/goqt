// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QResource : QResource
type QResource struct {
	BaseDrv
}
//QResource::QResource()	
func NewResource() *QResource {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),116000,116102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QResource{}
	_p.SetDriver(__rv,116000,true)
	return _p
} 
//QResource::QResource(QString const&,QLocale const&)	
func NewResourceWithFilenameLocale(file string,locale *QLocale) *QResource {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),116000,116103,unsafe.Pointer(&file),Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QResource{}
	_p.SetDriver(__rv,116000,true)
	return _p
} 
//QResource::absoluteFilePath()
func (q *QResource) AbsoluteFilePath() string {
	var __rv string
	q.Drv(116000,116104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::addSearchPath(QString const&)	
func QResourceAddSearchPath(path string)  {
	DirectQtDrv(nil,116000,116105,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QResource::addSearchPath(QString const&)
func (q *QResource) AddSearchPath(path string)  {
	q.Drv(116000,116105,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QResource::data()
func (q *QResource) Data() *byte {
	var __rv *byte
	q.Drv(116000,116106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::fileName()
func (q *QResource) FileName() string {
	var __rv string
	q.Drv(116000,116107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::isCompressed()
func (q *QResource) IsCompressed() bool {
	var __rv bool
	q.Drv(116000,116108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::isValid()
func (q *QResource) IsValid() bool {
	var __rv bool
	q.Drv(116000,116109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::locale()
func (q *QResource) Locale() *QLocale {
	var __rv uintptr
	q.Drv(116000,116110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLocale{}
	_rp.SetDriver(__rv,72000,true)
	return _rp
}	
//QResource::registerResource(QString const&)	
func QResourceRegisterResource(rccFilename string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116111,unsafe.Pointer(&rccFilename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(QString const&)
func (q *QResource) RegisterResource(rccFilename string) bool {
	var __rv bool
	q.Drv(116000,116111,unsafe.Pointer(&rccFilename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(unsigned char const*)	
func QResourceRegisterResourceWithRccdata(rccData *byte) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116112,unsafe.Pointer(&rccData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(unsigned char const*)
func (q *QResource) RegisterResourceWithRccdata(rccData *byte) bool {
	var __rv bool
	q.Drv(116000,116112,unsafe.Pointer(&rccData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(QString const&,QString const&)	
func QResourceRegisterResourceWithRccfilenameResourceroot(rccFilename string,resourceRoot string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116113,unsafe.Pointer(&rccFilename),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(QString const&,QString const&)
func (q *QResource) RegisterResourceWithRccfilenameResourceroot(rccFilename string,resourceRoot string) bool {
	var __rv bool
	q.Drv(116000,116113,unsafe.Pointer(&rccFilename),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(unsigned char const*,QString const&)	
func QResourceRegisterResourceWithRccdataResourceroot(rccData *byte,resourceRoot string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116114,unsafe.Pointer(&rccData),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::registerResource(unsigned char const*,QString const&)
func (q *QResource) RegisterResourceWithRccdataResourceroot(rccData *byte,resourceRoot string) bool {
	var __rv bool
	q.Drv(116000,116114,unsafe.Pointer(&rccData),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::searchPaths()	
func QResourceSearchPaths() []string {
	var __rv []string
	DirectQtDrv(nil,116000,116115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::searchPaths()
func (q *QResource) SearchPaths() []string {
	var __rv []string
	q.Drv(116000,116115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::setFileName(QString const&)
func (q *QResource) SetFileName(file string)  {
	q.Drv(116000,116116,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QResource::setLocale(QLocale const&)
func (q *QResource) SetLocale(locale *QLocale)  {
	q.Drv(116000,116117,Native(locale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QResource::size()
func (q *QResource) Size() int64 {
	var __rv int64
	q.Drv(116000,116118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(QString const&)	
func QResourceUnregisterResource(rccFilename string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116119,unsafe.Pointer(&rccFilename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(QString const&)
func (q *QResource) UnregisterResource(rccFilename string) bool {
	var __rv bool
	q.Drv(116000,116119,unsafe.Pointer(&rccFilename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(unsigned char const*)	
func QResourceUnregisterResourceWithRccdata(rccData *byte) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116120,unsafe.Pointer(&rccData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(unsigned char const*)
func (q *QResource) UnregisterResourceWithRccdata(rccData *byte) bool {
	var __rv bool
	q.Drv(116000,116120,unsafe.Pointer(&rccData),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(QString const&,QString const&)	
func QResourceUnregisterResourceWithRccfilenameResourceroot(rccFilename string,resourceRoot string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116121,unsafe.Pointer(&rccFilename),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(QString const&,QString const&)
func (q *QResource) UnregisterResourceWithRccfilenameResourceroot(rccFilename string,resourceRoot string) bool {
	var __rv bool
	q.Drv(116000,116121,unsafe.Pointer(&rccFilename),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(unsigned char const*,QString const&)	
func QResourceUnregisterResourceWithRccdataResourceroot(rccData *byte,resourceRoot string) bool {
	var __rv bool
	DirectQtDrv(nil,116000,116122,unsafe.Pointer(&rccData),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QResource::unregisterResource(unsigned char const*,QString const&)
func (q *QResource) UnregisterResourceWithRccdataResourceroot(rccData *byte,resourceRoot string) bool {
	var __rv bool
	q.Drv(116000,116122,unsafe.Pointer(&rccData),unsafe.Pointer(&resourceRoot),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
