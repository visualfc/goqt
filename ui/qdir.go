// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDir_Filter - QDir::Filter
type QDir_Filter int32
const (
	QDir_Dirs QDir_Filter = 0x001
	QDir_Files QDir_Filter = 0x002
	QDir_Drives QDir_Filter = 0x004
	QDir_NoSymLinks QDir_Filter = 0x008
	QDir_AllEntries QDir_Filter = QDir_Dirs|QDir_Files|QDir_Drives
	QDir_TypeMask QDir_Filter = 0x00f
	QDir_Readable QDir_Filter = 0x010
	QDir_Writable QDir_Filter = 0x020
	QDir_Executable QDir_Filter = 0x040
	QDir_PermissionMask QDir_Filter = 0x070
	QDir_Modified QDir_Filter = 0x080
	QDir_Hidden QDir_Filter = 0x100
	QDir_System QDir_Filter = 0x200
	QDir_AccessMask QDir_Filter = 0x3F0
	QDir_AllDirs QDir_Filter = 0x400
	QDir_CaseSensitive QDir_Filter = 0x800
	QDir_NoDotAndDotDot QDir_Filter = 0x1000
	QDir_NoDot QDir_Filter = 0x2000
	QDir_NoDotDot QDir_Filter = 0x4000
	QDir_NoFilter QDir_Filter = -1
)
//enum QDir_SortFlag - QDir::SortFlag
type QDir_SortFlag int32
const (
	QDir_Name QDir_SortFlag = 0x00
	QDir_Time QDir_SortFlag = 0x01
	QDir_Size QDir_SortFlag = 0x02
	QDir_Unsorted QDir_SortFlag = 0x03
	QDir_SortByMask QDir_SortFlag = 0x03
	QDir_DirsFirst QDir_SortFlag = 0x04
	QDir_Reversed QDir_SortFlag = 0x08
	QDir_IgnoreCase QDir_SortFlag = 0x10
	QDir_DirsLast QDir_SortFlag = 0x20
	QDir_LocaleAware QDir_SortFlag = 0x40
	QDir_Type QDir_SortFlag = 0x80
	QDir_NoSort QDir_SortFlag = -1
)
//struct QDir : QDir
type QDir struct {
	BaseDrv
}
//QDir::QDir()	
func NewDir() *QDir {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),22000,22102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDir{}
	_p.SetDriver(__rv,22000,true)
	return _p
} 
//QDir::QDir(QDir const&)	
func NewDirCopy(value *QDir) *QDir {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),22000,22103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDir{}
	_p.SetDriver(__rv,22000,true)
	return _p
} 
//QDir::QDir(QString const&)	
func NewDirWithPath(path string) *QDir {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),22000,22104,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDir{}
	_p.SetDriver(__rv,22000,true)
	return _p
} 
//QDir::QDir(QString const&,QString const&,QFlags<QDir::SortFlag>,QFlags<QDir::Filter>)	
func NewDirWithPathNamefilterSortFilter(path string,nameFilter string,sort QDir_SortFlag,filter QDir_Filter) *QDir {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),22000,22105,unsafe.Pointer(&path),unsafe.Pointer(&nameFilter),unsafe.Pointer(&sort),unsafe.Pointer(&filter),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDir{}
	_p.SetDriver(__rv,22000,true)
	return _p
} 
//QDir::absoluteFilePath(QString const&)
func (q *QDir) AbsoluteFilePath(fileName string) string {
	var __rv string
	q.Drv(22000,22106,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::absolutePath()
func (q *QDir) AbsolutePath() string {
	var __rv string
	q.Drv(22000,22107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::addResourceSearchPath(QString const&)	
func QDirAddResourceSearchPath(path string)  {
	DirectQtDrv(nil,22000,22108,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::addResourceSearchPath(QString const&)
func (q *QDir) AddResourceSearchPath(path string)  {
	q.Drv(22000,22108,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::addSearchPath(QString const&,QString const&)	
func QDirAddSearchPath(prefix string,path string)  {
	DirectQtDrv(nil,22000,22109,unsafe.Pointer(&prefix),unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::addSearchPath(QString const&,QString const&)
func (q *QDir) AddSearchPath(prefix string,path string)  {
	q.Drv(22000,22109,unsafe.Pointer(&prefix),unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::canonicalPath()
func (q *QDir) CanonicalPath() string {
	var __rv string
	q.Drv(22000,22110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::cd(QString const&)
func (q *QDir) Cd(dirName string) bool {
	var __rv bool
	q.Drv(22000,22111,unsafe.Pointer(&dirName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::cdUp()
func (q *QDir) CdUp() bool {
	var __rv bool
	q.Drv(22000,22112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::cleanPath(QString const&)	
func QDirCleanPath(path string) string {
	var __rv string
	DirectQtDrv(nil,22000,22113,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::cleanPath(QString const&)
func (q *QDir) CleanPath(path string) string {
	var __rv string
	q.Drv(22000,22113,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::count()
func (q *QDir) Count() uint {
	var __rv uint
	q.Drv(22000,22114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::current()	
func QDirCurrent() *QDir {
	var __rv uintptr
	DirectQtDrv(nil,22000,22115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::current()
func (q *QDir) Current() *QDir {
	var __rv uintptr
	q.Drv(22000,22115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::currentPath()	
func QDirCurrentPath() string {
	var __rv string
	DirectQtDrv(nil,22000,22116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::currentPath()
func (q *QDir) CurrentPath() string {
	var __rv string
	q.Drv(22000,22116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::dirName()
func (q *QDir) DirName() string {
	var __rv string
	q.Drv(22000,22117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::drives()	
func QDirDrives() []*QFileInfo {
	var __rv []*QFileInfo
	DirectQtDrv(nil,22000,22118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::drives()
func (q *QDir) Drives() []*QFileInfo {
	var __rv []*QFileInfo
	q.Drv(22000,22118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryInfoList()
func (q *QDir) EntryInfoList() []*QFileInfo {
	var __rv []*QFileInfo
	q.Drv(22000,22119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryInfoList(QStringList const&)
func (q *QDir) EntryInfoListWithNamefilters(nameFilters []string) []*QFileInfo {
	var __rv []*QFileInfo
	q.Drv(22000,22120,unsafe.Pointer(&nameFilters),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryInfoList(QFlags<QDir::Filter>,QFlags<QDir::SortFlag>)
func (q *QDir) EntryInfoListWithFiltersSort(filters QDir_Filter,sort QDir_SortFlag) []*QFileInfo {
	var __rv []*QFileInfo
	q.Drv(22000,22121,unsafe.Pointer(&filters),unsafe.Pointer(&sort),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryInfoList(QStringList const&,QFlags<QDir::Filter>,QFlags<QDir::SortFlag>)
func (q *QDir) EntryInfoListWithNamefiltersFiltersSort(nameFilters []string,filters QDir_Filter,sort QDir_SortFlag) []*QFileInfo {
	var __rv []*QFileInfo
	q.Drv(22000,22122,unsafe.Pointer(&nameFilters),unsafe.Pointer(&filters),unsafe.Pointer(&sort),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryList()
func (q *QDir) EntryList() []string {
	var __rv []string
	q.Drv(22000,22123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryList(QStringList const&)
func (q *QDir) EntryListWithNamefilters(nameFilters []string) []string {
	var __rv []string
	q.Drv(22000,22124,unsafe.Pointer(&nameFilters),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryList(QFlags<QDir::Filter>,QFlags<QDir::SortFlag>)
func (q *QDir) EntryListWithFiltersSort(filters QDir_Filter,sort QDir_SortFlag) []string {
	var __rv []string
	q.Drv(22000,22125,unsafe.Pointer(&filters),unsafe.Pointer(&sort),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::entryList(QStringList const&,QFlags<QDir::Filter>,QFlags<QDir::SortFlag>)
func (q *QDir) EntryListWithNamefiltersFiltersSort(nameFilters []string,filters QDir_Filter,sort QDir_SortFlag) []string {
	var __rv []string
	q.Drv(22000,22126,unsafe.Pointer(&nameFilters),unsafe.Pointer(&filters),unsafe.Pointer(&sort),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::exists()
func (q *QDir) Exists() bool {
	var __rv bool
	q.Drv(22000,22127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::exists(QString const&)
func (q *QDir) ExistsWithName(name string) bool {
	var __rv bool
	q.Drv(22000,22128,unsafe.Pointer(&name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::filePath(QString const&)
func (q *QDir) FilePath(fileName string) string {
	var __rv string
	q.Drv(22000,22129,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::filter()
func (q *QDir) Filter() QDir_Filter {
	var __rv QDir_Filter
	q.Drv(22000,22130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::fromNativeSeparators(QString const&)	
func QDirFromNativeSeparators(pathName string) string {
	var __rv string
	DirectQtDrv(nil,22000,22131,unsafe.Pointer(&pathName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::fromNativeSeparators(QString const&)
func (q *QDir) FromNativeSeparators(pathName string) string {
	var __rv string
	q.Drv(22000,22131,unsafe.Pointer(&pathName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::home()	
func QDirHome() *QDir {
	var __rv uintptr
	DirectQtDrv(nil,22000,22132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::home()
func (q *QDir) Home() *QDir {
	var __rv uintptr
	q.Drv(22000,22132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::homePath()	
func QDirHomePath() string {
	var __rv string
	DirectQtDrv(nil,22000,22133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::homePath()
func (q *QDir) HomePath() string {
	var __rv string
	q.Drv(22000,22133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isAbsolute()
func (q *QDir) IsAbsolute() bool {
	var __rv bool
	q.Drv(22000,22134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isAbsolutePath(QString const&)	
func QDirIsAbsolutePath(path string) bool {
	var __rv bool
	DirectQtDrv(nil,22000,22135,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isAbsolutePath(QString const&)
func (q *QDir) IsAbsolutePath(path string) bool {
	var __rv bool
	q.Drv(22000,22135,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isReadable()
func (q *QDir) IsReadable() bool {
	var __rv bool
	q.Drv(22000,22136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isRelative()
func (q *QDir) IsRelative() bool {
	var __rv bool
	q.Drv(22000,22137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isRelativePath(QString const&)	
func QDirIsRelativePath(path string) bool {
	var __rv bool
	DirectQtDrv(nil,22000,22138,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isRelativePath(QString const&)
func (q *QDir) IsRelativePath(path string) bool {
	var __rv bool
	q.Drv(22000,22138,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::isRoot()
func (q *QDir) IsRoot() bool {
	var __rv bool
	q.Drv(22000,22139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::makeAbsolute()
func (q *QDir) MakeAbsolute() bool {
	var __rv bool
	q.Drv(22000,22140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::match(QString const&,QString const&)	
func QDirMatchWithFilterFilename(filter string,fileName string) bool {
	var __rv bool
	DirectQtDrv(nil,22000,22141,unsafe.Pointer(&filter),unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::match(QString const&,QString const&)
func (q *QDir) MatchWithFilterFilename(filter string,fileName string) bool {
	var __rv bool
	q.Drv(22000,22141,unsafe.Pointer(&filter),unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::match(QStringList const&,QString const&)	
func QDirMatchWithFiltersFilename(filters []string,fileName string) bool {
	var __rv bool
	DirectQtDrv(nil,22000,22142,unsafe.Pointer(&filters),unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::match(QStringList const&,QString const&)
func (q *QDir) MatchWithFiltersFilename(filters []string,fileName string) bool {
	var __rv bool
	q.Drv(22000,22142,unsafe.Pointer(&filters),unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::mkdir(QString const&)
func (q *QDir) Mkdir(dirName string) bool {
	var __rv bool
	q.Drv(22000,22143,unsafe.Pointer(&dirName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::mkpath(QString const&)
func (q *QDir) Mkpath(dirPath string) bool {
	var __rv bool
	q.Drv(22000,22144,unsafe.Pointer(&dirPath),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::nameFilters()
func (q *QDir) NameFilters() []string {
	var __rv []string
	q.Drv(22000,22145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::nameFiltersFromString(QString const&)	
func QDirNameFiltersFromString(nameFilter string) []string {
	var __rv []string
	DirectQtDrv(nil,22000,22146,unsafe.Pointer(&nameFilter),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::nameFiltersFromString(QString const&)
func (q *QDir) NameFiltersFromString(nameFilter string) []string {
	var __rv []string
	q.Drv(22000,22146,unsafe.Pointer(&nameFilter),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::path()
func (q *QDir) Path() string {
	var __rv string
	q.Drv(22000,22147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::refresh()
func (q *QDir) Refresh()  {
	q.Drv(22000,22148,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::relativeFilePath(QString const&)
func (q *QDir) RelativeFilePath(fileName string) string {
	var __rv string
	q.Drv(22000,22149,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::remove(QString const&)
func (q *QDir) Remove(fileName string) bool {
	var __rv bool
	q.Drv(22000,22150,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::rename(QString const&,QString const&)
func (q *QDir) Rename(oldName string,newName string) bool {
	var __rv bool
	q.Drv(22000,22151,unsafe.Pointer(&oldName),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::rmdir(QString const&)
func (q *QDir) Rmdir(dirName string) bool {
	var __rv bool
	q.Drv(22000,22152,unsafe.Pointer(&dirName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::rmpath(QString const&)
func (q *QDir) Rmpath(dirPath string) bool {
	var __rv bool
	q.Drv(22000,22153,unsafe.Pointer(&dirPath),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::root()	
func QDirRoot() *QDir {
	var __rv uintptr
	DirectQtDrv(nil,22000,22154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::root()
func (q *QDir) Root() *QDir {
	var __rv uintptr
	q.Drv(22000,22154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::rootPath()	
func QDirRootPath() string {
	var __rv string
	DirectQtDrv(nil,22000,22155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::rootPath()
func (q *QDir) RootPath() string {
	var __rv string
	q.Drv(22000,22155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::searchPaths(QString const&)	
func QDirSearchPaths(prefix string) []string {
	var __rv []string
	DirectQtDrv(nil,22000,22156,unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::searchPaths(QString const&)
func (q *QDir) SearchPaths(prefix string) []string {
	var __rv []string
	q.Drv(22000,22156,unsafe.Pointer(&prefix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::separator()	
func QDirSeparator() rune {
	var __rv rune
	DirectQtDrv(nil,22000,22157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::separator()
func (q *QDir) Separator() rune {
	var __rv rune
	q.Drv(22000,22157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::setCurrent(QString const&)	
func QDirSetCurrent(path string) bool {
	var __rv bool
	DirectQtDrv(nil,22000,22158,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::setCurrent(QString const&)
func (q *QDir) SetCurrent(path string) bool {
	var __rv bool
	q.Drv(22000,22158,unsafe.Pointer(&path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::setFilter(QFlags<QDir::Filter>)
func (q *QDir) SetFilter(filter QDir_Filter)  {
	q.Drv(22000,22159,unsafe.Pointer(&filter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::setNameFilters(QStringList const&)
func (q *QDir) SetNameFilters(nameFilters []string)  {
	q.Drv(22000,22160,unsafe.Pointer(&nameFilters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::setPath(QString const&)
func (q *QDir) SetPath(path string)  {
	q.Drv(22000,22161,unsafe.Pointer(&path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::setSearchPaths(QString const&,QStringList const&)	
func QDirSetSearchPaths(prefix string,searchPaths []string)  {
	DirectQtDrv(nil,22000,22162,unsafe.Pointer(&prefix),unsafe.Pointer(&searchPaths),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::setSearchPaths(QString const&,QStringList const&)
func (q *QDir) SetSearchPaths(prefix string,searchPaths []string)  {
	q.Drv(22000,22162,unsafe.Pointer(&prefix),unsafe.Pointer(&searchPaths),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::setSorting(QFlags<QDir::SortFlag>)
func (q *QDir) SetSorting(sort QDir_SortFlag)  {
	q.Drv(22000,22163,unsafe.Pointer(&sort),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDir::sorting()
func (q *QDir) Sorting() QDir_SortFlag {
	var __rv QDir_SortFlag
	q.Drv(22000,22164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::temp()	
func QDirTemp() *QDir {
	var __rv uintptr
	DirectQtDrv(nil,22000,22165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::temp()
func (q *QDir) Temp() *QDir {
	var __rv uintptr
	q.Drv(22000,22165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QDir::tempPath()	
func QDirTempPath() string {
	var __rv string
	DirectQtDrv(nil,22000,22166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::tempPath()
func (q *QDir) TempPath() string {
	var __rv string
	q.Drv(22000,22166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::toNativeSeparators(QString const&)	
func QDirToNativeSeparators(pathName string) string {
	var __rv string
	DirectQtDrv(nil,22000,22167,unsafe.Pointer(&pathName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDir::toNativeSeparators(QString const&)
func (q *QDir) ToNativeSeparators(pathName string) string {
	var __rv string
	q.Drv(22000,22167,unsafe.Pointer(&pathName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
