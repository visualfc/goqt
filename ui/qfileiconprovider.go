// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFileIconProvider_IconType - QFileIconProvider::IconType
type QFileIconProvider_IconType uint32
const (
	QFileIconProvider_Computer QFileIconProvider_IconType = 0
	QFileIconProvider_Desktop QFileIconProvider_IconType = 1
	QFileIconProvider_Trashcan QFileIconProvider_IconType = 2
	QFileIconProvider_Network QFileIconProvider_IconType = 3
	QFileIconProvider_Drive QFileIconProvider_IconType = 4
	QFileIconProvider_Folder QFileIconProvider_IconType = 5
	QFileIconProvider_File QFileIconProvider_IconType = 6
)
//struct QFileIconProvider : QFileIconProvider
type QFileIconProvider struct {
	BaseDrv
}
//QFileIconProvider::QFileIconProvider()	
func NewFileIconProvider() *QFileIconProvider {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),33000,33102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileIconProvider{}
	_p.SetDriver(__rv,33000,true)
	return _p
} 
//QFileIconProvider::icon(QFileIconProvider::IconType)
func (q *QFileIconProvider) Icon(_type QFileIconProvider_IconType) *QIcon {
	var __rv uintptr
	q.Drv(33000,33103,unsafe.Pointer(&_type),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QFileIconProvider::icon(QFileInfo const&)
func (q *QFileIconProvider) IconWithInfo(info *QFileInfo) *QIcon {
	var __rv uintptr
	q.Drv(33000,33104,Native(info),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIcon{}
	_rp.SetDriver(__rv,51000,true)
	return _rp
}	
//QFileIconProvider::type(QFileInfo const&)
func (q *QFileIconProvider) Type(info *QFileInfo) string {
	var __rv string
	q.Drv(33000,33105,Native(info),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
