// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFileInfo : QFileInfo
type QFileInfo struct {
	BaseDrv
}
//QFileInfo::QFileInfo()	
func NewFileInfo() *QFileInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),34000,34102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileInfo{}
	_p.SetDriver(__rv,34000,true)
	return _p
} 
//QFileInfo::QFileInfo(QFile const&)	
func NewFileInfoWithFile(file *QFile) *QFileInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),34000,34103,Native(file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileInfo{}
	_p.SetDriver(__rv,34000,true)
	return _p
} 
//QFileInfo::QFileInfo(QFileInfo const&)	
func NewFileInfoCopy(fileinfo *QFileInfo) *QFileInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),34000,34104,Native(fileinfo),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileInfo{}
	_p.SetDriver(__rv,34000,true)
	return _p
} 
//QFileInfo::QFileInfo(QString const&)	
func NewFileInfoWithFilename(file string) *QFileInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),34000,34105,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileInfo{}
	_p.SetDriver(__rv,34000,true)
	return _p
} 
//QFileInfo::QFileInfo(QDir const&,QString const&)	
func NewFileInfoWithDirFilename(dir *QDir,file string) *QFileInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),34000,34106,Native(dir),unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFileInfo{}
	_p.SetDriver(__rv,34000,true)
	return _p
} 
//QFileInfo::absoluteDir()
func (q *QFileInfo) AbsoluteDir() *QDir {
	var __rv uintptr
	q.Drv(34000,34107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QFileInfo::absoluteFilePath()
func (q *QFileInfo) AbsoluteFilePath() string {
	var __rv string
	q.Drv(34000,34108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::absolutePath()
func (q *QFileInfo) AbsolutePath() string {
	var __rv string
	q.Drv(34000,34109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::baseName()
func (q *QFileInfo) BaseName() string {
	var __rv string
	q.Drv(34000,34110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::bundleName()
func (q *QFileInfo) BundleName() string {
	var __rv string
	q.Drv(34000,34111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::caching()
func (q *QFileInfo) Caching() bool {
	var __rv bool
	q.Drv(34000,34112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::canonicalFilePath()
func (q *QFileInfo) CanonicalFilePath() string {
	var __rv string
	q.Drv(34000,34113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::canonicalPath()
func (q *QFileInfo) CanonicalPath() string {
	var __rv string
	q.Drv(34000,34114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::completeBaseName()
func (q *QFileInfo) CompleteBaseName() string {
	var __rv string
	q.Drv(34000,34115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::completeSuffix()
func (q *QFileInfo) CompleteSuffix() string {
	var __rv string
	q.Drv(34000,34116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::created()
func (q *QFileInfo) Created() *QDateTime {
	var __rv uintptr
	q.Drv(34000,34117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QFileInfo::dir()
func (q *QFileInfo) Dir() *QDir {
	var __rv uintptr
	q.Drv(34000,34118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDir{}
	_rp.SetDriver(__rv,22000,true)
	return _rp
}	
//QFileInfo::exists()
func (q *QFileInfo) Exists() bool {
	var __rv bool
	q.Drv(34000,34119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::fileName()
func (q *QFileInfo) FileName() string {
	var __rv string
	q.Drv(34000,34120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::filePath()
func (q *QFileInfo) FilePath() string {
	var __rv string
	q.Drv(34000,34121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::group()
func (q *QFileInfo) Group() string {
	var __rv string
	q.Drv(34000,34122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::groupId()
func (q *QFileInfo) GroupId() uint {
	var __rv uint
	q.Drv(34000,34123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isAbsolute()
func (q *QFileInfo) IsAbsolute() bool {
	var __rv bool
	q.Drv(34000,34124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isBundle()
func (q *QFileInfo) IsBundle() bool {
	var __rv bool
	q.Drv(34000,34125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isDir()
func (q *QFileInfo) IsDir() bool {
	var __rv bool
	q.Drv(34000,34126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isExecutable()
func (q *QFileInfo) IsExecutable() bool {
	var __rv bool
	q.Drv(34000,34127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isFile()
func (q *QFileInfo) IsFile() bool {
	var __rv bool
	q.Drv(34000,34128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isHidden()
func (q *QFileInfo) IsHidden() bool {
	var __rv bool
	q.Drv(34000,34129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isReadable()
func (q *QFileInfo) IsReadable() bool {
	var __rv bool
	q.Drv(34000,34130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isRelative()
func (q *QFileInfo) IsRelative() bool {
	var __rv bool
	q.Drv(34000,34131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isRoot()
func (q *QFileInfo) IsRoot() bool {
	var __rv bool
	q.Drv(34000,34132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isSymLink()
func (q *QFileInfo) IsSymLink() bool {
	var __rv bool
	q.Drv(34000,34133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::isWritable()
func (q *QFileInfo) IsWritable() bool {
	var __rv bool
	q.Drv(34000,34134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::lastModified()
func (q *QFileInfo) LastModified() *QDateTime {
	var __rv uintptr
	q.Drv(34000,34135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QFileInfo::lastRead()
func (q *QFileInfo) LastRead() *QDateTime {
	var __rv uintptr
	q.Drv(34000,34136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QDateTime{}
	_rp.SetDriver(__rv,20000,true)
	return _rp
}	
//QFileInfo::makeAbsolute()
func (q *QFileInfo) MakeAbsolute() bool {
	var __rv bool
	q.Drv(34000,34137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::owner()
func (q *QFileInfo) Owner() string {
	var __rv string
	q.Drv(34000,34138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::ownerId()
func (q *QFileInfo) OwnerId() uint {
	var __rv uint
	q.Drv(34000,34139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::path()
func (q *QFileInfo) Path() string {
	var __rv string
	q.Drv(34000,34140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::permission(QFlags<QFile::Permission>)
func (q *QFileInfo) Permission(permissions QFile_Permission) bool {
	var __rv bool
	q.Drv(34000,34141,unsafe.Pointer(&permissions),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::permissions()
func (q *QFileInfo) Permissions() QFile_Permission {
	var __rv QFile_Permission
	q.Drv(34000,34142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::readLink()
func (q *QFileInfo) ReadLink() string {
	var __rv string
	q.Drv(34000,34143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::refresh()
func (q *QFileInfo) Refresh()  {
	q.Drv(34000,34144,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileInfo::setCaching(bool)
func (q *QFileInfo) SetCaching(on bool)  {
	q.Drv(34000,34145,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileInfo::setFile(QFile const&)
func (q *QFileInfo) SetFile(file *QFile)  {
	q.Drv(34000,34146,Native(file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileInfo::setFile(QString const&)
func (q *QFileInfo) SetFileWithFilename(file string)  {
	q.Drv(34000,34147,unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileInfo::setFile(QDir const&,QString const&)
func (q *QFileInfo) SetFileWithDirFilename(dir *QDir,file string)  {
	q.Drv(34000,34148,Native(dir),unsafe.Pointer(&file),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFileInfo::size()
func (q *QFileInfo) Size() int64 {
	var __rv int64
	q.Drv(34000,34149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::suffix()
func (q *QFileInfo) Suffix() string {
	var __rv string
	q.Drv(34000,34150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFileInfo::symLinkTarget()
func (q *QFileInfo) SymLinkTarget() string {
	var __rv string
	q.Drv(34000,34151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
