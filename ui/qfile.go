// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFile_FileError - QFile::FileError
type QFile_FileError uint32
const (
	QFile_NoError QFile_FileError = 0
	QFile_ReadError QFile_FileError = 1
	QFile_WriteError QFile_FileError = 2
	QFile_FatalError QFile_FileError = 3
	QFile_ResourceError QFile_FileError = 4
	QFile_OpenError QFile_FileError = 5
	QFile_AbortError QFile_FileError = 6
	QFile_TimeOutError QFile_FileError = 7
	QFile_UnspecifiedError QFile_FileError = 8
	QFile_RemoveError QFile_FileError = 9
	QFile_RenameError QFile_FileError = 10
	QFile_PositionError QFile_FileError = 11
	QFile_ResizeError QFile_FileError = 12
	QFile_PermissionsError QFile_FileError = 13
	QFile_CopyError QFile_FileError = 14
)
//enum QFile_MemoryMapFlags - QFile::MemoryMapFlags
type QFile_MemoryMapFlags uint32
const (
	QFile_NoOptions QFile_MemoryMapFlags = 0
)
//enum QFile_Permission - QFile::Permission
type QFile_Permission uint32
const (
	QFile_ReadOwner QFile_Permission = 0x4000
	QFile_WriteOwner QFile_Permission = 0x2000
	QFile_ExeOwner QFile_Permission = 0x1000
	QFile_ReadUser QFile_Permission = 0x0400
	QFile_WriteUser QFile_Permission = 0x0200
	QFile_ExeUser QFile_Permission = 0x0100
	QFile_ReadGroup QFile_Permission = 0x0040
	QFile_WriteGroup QFile_Permission = 0x0020
	QFile_ExeGroup QFile_Permission = 0x0010
	QFile_ReadOther QFile_Permission = 0x0004
	QFile_WriteOther QFile_Permission = 0x0002
	QFile_ExeOther QFile_Permission = 0x0001
)
//struct QFile : QFile
type QFile struct {
	QIODevice
}
//QFile::QFile()	
func NewFile() *QFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),236000,236102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFile{}
	_p.SetDriver(__rv,236000,false)
	return _p
} 
//QFile::QFile(QObject*)	
func NewFileWithParent(parent QObjectInterface) *QFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),236000,236103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFile{}
	_p.SetDriver(__rv,236000,false)
	return _p
} 
//QFile::QFile(QString const&)	
func NewFileWithName(name string) *QFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),236000,236104,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFile{}
	_p.SetDriver(__rv,236000,false)
	return _p
} 
//QFile::QFile(QString const&,QObject*)	
func NewFileWithNameParent(name string,parent QObjectInterface) *QFile {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),236000,236105,unsafe.Pointer(&name),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFile{}
	_p.SetDriver(__rv,236000,false)
	return _p
} 
//QFile::atEnd()
func (q *QFile) AtEnd() bool {
	var __rv bool
	q.Drv(236000,236106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::close()
func (q *QFile) Close()  {
	q.Drv(236000,236107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFile::copy(QString const&)
func (q *QFile) Copy(newName string) bool {
	var __rv bool
	q.Drv(236000,236108,unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::copy(QString const&,QString const&)	
func QFileCopyWithFilenameNewname(fileName string,newName string) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236109,unsafe.Pointer(&fileName),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::copy(QString const&,QString const&)
func (q *QFile) CopyWithFilenameNewname(fileName string,newName string) bool {
	var __rv bool
	q.Drv(236000,236109,unsafe.Pointer(&fileName),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::decodeName(QByteArray const&)	
func QFileDecodeName(localFileName []byte) string {
	var __rv string
	DirectQtDrv(nil,236000,236110,unsafe.Pointer(&localFileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::decodeName(QByteArray const&)
func (q *QFile) DecodeName(localFileName []byte) string {
	var __rv string
	q.Drv(236000,236110,unsafe.Pointer(&localFileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::encodeName(QString const&)	
func QFileEncodeName(fileName string) []byte {
	var __rv []byte
	DirectQtDrv(nil,236000,236111,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::encodeName(QString const&)
func (q *QFile) EncodeName(fileName string) []byte {
	var __rv []byte
	q.Drv(236000,236111,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::error()
func (q *QFile) Error() QFile_FileError {
	var __rv QFile_FileError
	q.Drv(236000,236112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::exists()
func (q *QFile) Exists() bool {
	var __rv bool
	q.Drv(236000,236113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::exists(QString const&)	
func QFileExistsWithFilename(fileName string) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236114,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::exists(QString const&)
func (q *QFile) ExistsWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(236000,236114,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::fileName()
func (q *QFile) FileName() string {
	var __rv string
	q.Drv(236000,236115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::flush()
func (q *QFile) Flush() bool {
	var __rv bool
	q.Drv(236000,236116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::handle()
func (q *QFile) Handle() int {
	var __rv int
	q.Drv(236000,236117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::isSequential()
func (q *QFile) IsSequential() bool {
	var __rv bool
	q.Drv(236000,236118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::link(QString const&)
func (q *QFile) Link(newName string) bool {
	var __rv bool
	q.Drv(236000,236119,unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::link(QString const&,QString const&)	
func QFileLinkWithOldnameNewname(oldname string,newName string) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236120,unsafe.Pointer(&oldname),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::link(QString const&,QString const&)
func (q *QFile) LinkWithOldnameNewname(oldname string,newName string) bool {
	var __rv bool
	q.Drv(236000,236120,unsafe.Pointer(&oldname),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::map(qint64,qint64,QFile::MemoryMapFlags)
func (q *QFile) Map(offset int64,size int64,flags QFile_MemoryMapFlags) *byte {
	var __rv *byte
	q.Drv(236000,236121,unsafe.Pointer(&offset),unsafe.Pointer(&size),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::open(QFlags<QIODevice::OpenModeFlag>)
func (q *QFile) Open(flags QIODevice_OpenModeFlag) bool {
	var __rv bool
	q.Drv(236000,236122,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::open(FILE*,QFlags<QIODevice::OpenModeFlag>)
func (q *QFile) OpenWithFFlags(f uintptr,flags QIODevice_OpenModeFlag) bool {
	var __rv bool
	q.Drv(236000,236123,unsafe.Pointer(&f),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::open(int,QFlags<QIODevice::OpenModeFlag>)
func (q *QFile) OpenWithFdFlags(fd int,flags QIODevice_OpenModeFlag) bool {
	var __rv bool
	q.Drv(236000,236124,unsafe.Pointer(&fd),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::permissions()
func (q *QFile) Permissions() QFile_Permission {
	var __rv QFile_Permission
	q.Drv(236000,236125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::permissions(QString const&)	
func QFilePermissionsWithFilename(filename string) QFile_Permission {
	var __rv QFile_Permission
	DirectQtDrv(nil,236000,236126,unsafe.Pointer(&filename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::permissions(QString const&)
func (q *QFile) PermissionsWithFilename(filename string) QFile_Permission {
	var __rv QFile_Permission
	q.Drv(236000,236126,unsafe.Pointer(&filename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::pos()
func (q *QFile) Pos() int64 {
	var __rv int64
	q.Drv(236000,236127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::readData(char*,qint64)
func (q *QFile) ReadData(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(236000,236128,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::readLineData(char*,qint64)
func (q *QFile) ReadLineData(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(236000,236129,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::readLink()
func (q *QFile) ReadLink() string {
	var __rv string
	q.Drv(236000,236130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::readLink(QString const&)	
func QFileReadLinkWithFilename(fileName string) string {
	var __rv string
	DirectQtDrv(nil,236000,236131,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::readLink(QString const&)
func (q *QFile) ReadLinkWithFilename(fileName string) string {
	var __rv string
	q.Drv(236000,236131,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::remove()
func (q *QFile) Remove() bool {
	var __rv bool
	q.Drv(236000,236132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::remove(QString const&)	
func QFileRemoveWithFilename(fileName string) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236133,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::remove(QString const&)
func (q *QFile) RemoveWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(236000,236133,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::rename(QString const&)
func (q *QFile) Rename(newName string) bool {
	var __rv bool
	q.Drv(236000,236134,unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::rename(QString const&,QString const&)	
func QFileRenameWithOldnameNewname(oldName string,newName string) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236135,unsafe.Pointer(&oldName),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::rename(QString const&,QString const&)
func (q *QFile) RenameWithOldnameNewname(oldName string,newName string) bool {
	var __rv bool
	q.Drv(236000,236135,unsafe.Pointer(&oldName),unsafe.Pointer(&newName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::resize(qint64)
func (q *QFile) Resize(sz int64) bool {
	var __rv bool
	q.Drv(236000,236136,unsafe.Pointer(&sz),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::resize(QString const&,qint64)	
func QFileResizeWithFilenameSz(filename string,sz int64) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236137,unsafe.Pointer(&filename),unsafe.Pointer(&sz),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::resize(QString const&,qint64)
func (q *QFile) ResizeWithFilenameSz(filename string,sz int64) bool {
	var __rv bool
	q.Drv(236000,236137,unsafe.Pointer(&filename),unsafe.Pointer(&sz),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::seek(qint64)
func (q *QFile) Seek(offset int64) bool {
	var __rv bool
	q.Drv(236000,236138,unsafe.Pointer(&offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::setFileName(QString const&)
func (q *QFile) SetFileName(name string)  {
	q.Drv(236000,236139,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFile::setPermissions(QFlags<QFile::Permission>)
func (q *QFile) SetPermissions(permissionSpec QFile_Permission) bool {
	var __rv bool
	q.Drv(236000,236140,unsafe.Pointer(&permissionSpec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::setPermissions(QString const&,QFlags<QFile::Permission>)	
func QFileSetPermissionsWithFilenamePermissionspec(filename string,permissionSpec QFile_Permission) bool {
	var __rv bool
	DirectQtDrv(nil,236000,236141,unsafe.Pointer(&filename),unsafe.Pointer(&permissionSpec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::setPermissions(QString const&,QFlags<QFile::Permission>)
func (q *QFile) SetPermissionsWithFilenamePermissionspec(filename string,permissionSpec QFile_Permission) bool {
	var __rv bool
	q.Drv(236000,236141,unsafe.Pointer(&filename),unsafe.Pointer(&permissionSpec),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::size()
func (q *QFile) Size() int64 {
	var __rv int64
	q.Drv(236000,236142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::symLinkTarget()
func (q *QFile) SymLinkTarget() string {
	var __rv string
	q.Drv(236000,236143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::symLinkTarget(QString const&)	
func QFileSymLinkTargetWithFilename(fileName string) string {
	var __rv string
	DirectQtDrv(nil,236000,236144,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::symLinkTarget(QString const&)
func (q *QFile) SymLinkTargetWithFilename(fileName string) string {
	var __rv string
	q.Drv(236000,236144,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::unmap(unsigned char*)
func (q *QFile) Unmap(address *byte) bool {
	var __rv bool
	q.Drv(236000,236145,unsafe.Pointer(&address),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFile::unsetError()
func (q *QFile) UnsetError()  {
	q.Drv(236000,236146,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFile::writeData(char const*,qint64)
func (q *QFile) WriteData(data *byte,len int64) int64 {
	var __rv int64
	q.Drv(236000,236147,unsafe.Pointer(&data),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
