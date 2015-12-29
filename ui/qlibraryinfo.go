// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLibraryInfo_LibraryLocation - QLibraryInfo::LibraryLocation
type QLibraryInfo_LibraryLocation uint32
const (
	QLibraryInfo_PrefixPath QLibraryInfo_LibraryLocation = 0
	QLibraryInfo_DocumentationPath QLibraryInfo_LibraryLocation = 1
	QLibraryInfo_HeadersPath QLibraryInfo_LibraryLocation = 2
	QLibraryInfo_LibrariesPath QLibraryInfo_LibraryLocation = 3
	QLibraryInfo_BinariesPath QLibraryInfo_LibraryLocation = 4
	QLibraryInfo_PluginsPath QLibraryInfo_LibraryLocation = 5
	QLibraryInfo_DataPath QLibraryInfo_LibraryLocation = 6
	QLibraryInfo_TranslationsPath QLibraryInfo_LibraryLocation = 7
	QLibraryInfo_SettingsPath QLibraryInfo_LibraryLocation = 8
	QLibraryInfo_DemosPath QLibraryInfo_LibraryLocation = 9
	QLibraryInfo_ExamplesPath QLibraryInfo_LibraryLocation = 10
	QLibraryInfo_ImportsPath QLibraryInfo_LibraryLocation = 11
)
//struct QLibraryInfo : QLibraryInfo
type QLibraryInfo struct {
	BaseDrv
}
//QLibraryInfo::buildDate()	
func QLibraryInfoBuildDate() *QDate {
	var __rv uintptr
	DirectQtDrv(nil,67000,67102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QLibraryInfo::buildDate()
func (q *QLibraryInfo) BuildDate() *QDate {
	var __rv uintptr
	q.Drv(67000,67102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDate{}
	_rp.SetDriver(__rv,19000,true)
	return _rp
}	
//QLibraryInfo::licensedProducts()	
func QLibraryInfoLicensedProducts() string {
	var __rv string
	DirectQtDrv(nil,67000,67103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibraryInfo::licensedProducts()
func (q *QLibraryInfo) LicensedProducts() string {
	var __rv string
	q.Drv(67000,67103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibraryInfo::licensee()	
func QLibraryInfoLicensee() string {
	var __rv string
	DirectQtDrv(nil,67000,67104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibraryInfo::licensee()
func (q *QLibraryInfo) Licensee() string {
	var __rv string
	q.Drv(67000,67104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibraryInfo::location(QLibraryInfo::LibraryLocation)	
func QLibraryInfoLocation(value QLibraryInfo_LibraryLocation) string {
	var __rv string
	DirectQtDrv(nil,67000,67105,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLibraryInfo::location(QLibraryInfo::LibraryLocation)
func (q *QLibraryInfo) Location(value QLibraryInfo_LibraryLocation) string {
	var __rv string
	q.Drv(67000,67105,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
