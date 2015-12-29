// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QIODevice_OpenModeFlag - QIODevice::OpenModeFlag
type QIODevice_OpenModeFlag uint32
const (
	QIODevice_NotOpen QIODevice_OpenModeFlag = 0x0000
	QIODevice_ReadOnly QIODevice_OpenModeFlag = 0x0001
	QIODevice_WriteOnly QIODevice_OpenModeFlag = 0x0002
	QIODevice_ReadWrite QIODevice_OpenModeFlag = QIODevice_ReadOnly|QIODevice_WriteOnly
	QIODevice_Append QIODevice_OpenModeFlag = 0x0004
	QIODevice_Truncate QIODevice_OpenModeFlag = 0x0008
	QIODevice_Text QIODevice_OpenModeFlag = 0x0010
	QIODevice_Unbuffered QIODevice_OpenModeFlag = 0x0020
)
//struct QIODevice : QIODevice
type QIODevice struct {
	QObject
}
func (q *QIODevice) OnBytesWritten(fn func(int64)) uintptr {
	var __rv uintptr
	q.Drv(292000,292102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QIODevice) OnReadyRead(fn func()) uintptr {
	var __rv uintptr
	q.Drv(292000,292103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QIODevice) OnReadChannelFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(292000,292104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QIODevice) OnAboutToClose(fn func()) uintptr {
	var __rv uintptr
	q.Drv(292000,292105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QIODevice::atEnd()
func (q *QIODevice) AtEnd() bool {
	var __rv bool
	q.Drv(292000,292106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::bytesAvailable()
func (q *QIODevice) BytesAvailable() int64 {
	var __rv int64
	q.Drv(292000,292107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::bytesToWrite()
func (q *QIODevice) BytesToWrite() int64 {
	var __rv int64
	q.Drv(292000,292108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::canReadLine()
func (q *QIODevice) CanReadLine() bool {
	var __rv bool
	q.Drv(292000,292109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::close()
func (q *QIODevice) Close()  {
	q.Drv(292000,292110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIODevice::errorString()
func (q *QIODevice) ErrorString() string {
	var __rv string
	q.Drv(292000,292111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::getChar(char*)
func (q *QIODevice) GetChar(c *byte) bool {
	var __rv bool
	q.Drv(292000,292112,unsafe.Pointer(&c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::isOpen()
func (q *QIODevice) IsOpen() bool {
	var __rv bool
	q.Drv(292000,292113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::isReadable()
func (q *QIODevice) IsReadable() bool {
	var __rv bool
	q.Drv(292000,292114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::isSequential()
func (q *QIODevice) IsSequential() bool {
	var __rv bool
	q.Drv(292000,292115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::isTextModeEnabled()
func (q *QIODevice) IsTextModeEnabled() bool {
	var __rv bool
	q.Drv(292000,292116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::isWritable()
func (q *QIODevice) IsWritable() bool {
	var __rv bool
	q.Drv(292000,292117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::open(QFlags<QIODevice::OpenModeFlag>)
func (q *QIODevice) Open(mode QIODevice_OpenModeFlag) bool {
	var __rv bool
	q.Drv(292000,292118,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::openMode()
func (q *QIODevice) OpenMode() QIODevice_OpenModeFlag {
	var __rv QIODevice_OpenModeFlag
	q.Drv(292000,292119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::peek(qint64)
func (q *QIODevice) Peek(maxlen int64) []byte {
	var __rv []byte
	q.Drv(292000,292120,unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::peek(char*,qint64)
func (q *QIODevice) PeekWithDataMaxlen(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(292000,292121,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::pos()
func (q *QIODevice) Pos() int64 {
	var __rv int64
	q.Drv(292000,292122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::putChar(char)
func (q *QIODevice) PutChar(c byte) bool {
	var __rv bool
	q.Drv(292000,292123,unsafe.Pointer(&c),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::read(qint64)
func (q *QIODevice) Read(maxlen int64) []byte {
	var __rv []byte
	q.Drv(292000,292124,unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::read(char*,qint64)
func (q *QIODevice) ReadWithDataMaxlen(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(292000,292125,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readAll()
func (q *QIODevice) ReadAll() []byte {
	var __rv []byte
	q.Drv(292000,292126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readData(char*,qint64)
func (q *QIODevice) ReadData(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(292000,292127,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readLine()
func (q *QIODevice) ReadLine() []byte {
	var __rv []byte
	q.Drv(292000,292128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readLine(qint64)
func (q *QIODevice) ReadLineWithMaxlen(maxlen int64) []byte {
	var __rv []byte
	q.Drv(292000,292129,unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readLine(char*,qint64)
func (q *QIODevice) ReadLineWithDataMaxlen(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(292000,292130,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::readLineData(char*,qint64)
func (q *QIODevice) ReadLineData(data *byte,maxlen int64) int64 {
	var __rv int64
	q.Drv(292000,292131,unsafe.Pointer(&data),unsafe.Pointer(&maxlen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::reset()
func (q *QIODevice) Reset() bool {
	var __rv bool
	q.Drv(292000,292132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::seek(qint64)
func (q *QIODevice) Seek(pos int64) bool {
	var __rv bool
	q.Drv(292000,292133,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::setErrorString(QString const&)
func (q *QIODevice) SetErrorString(errorString string)  {
	q.Drv(292000,292134,unsafe.Pointer(&errorString),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIODevice::setOpenMode(QFlags<QIODevice::OpenModeFlag>)
func (q *QIODevice) SetOpenMode(openMode QIODevice_OpenModeFlag)  {
	q.Drv(292000,292135,unsafe.Pointer(&openMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIODevice::setTextModeEnabled(bool)
func (q *QIODevice) SetTextModeEnabled(enabled bool)  {
	q.Drv(292000,292136,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIODevice::size()
func (q *QIODevice) Size() int64 {
	var __rv int64
	q.Drv(292000,292137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::ungetChar(char)
func (q *QIODevice) UngetChar(c byte)  {
	q.Drv(292000,292138,unsafe.Pointer(&c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QIODevice::waitForBytesWritten(int)
func (q *QIODevice) WaitForBytesWritten(msecs int) bool {
	var __rv bool
	q.Drv(292000,292139,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::waitForReadyRead(int)
func (q *QIODevice) WaitForReadyRead(msecs int) bool {
	var __rv bool
	q.Drv(292000,292140,unsafe.Pointer(&msecs),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::write(QByteArray const&)
func (q *QIODevice) Write(data []byte) int64 {
	var __rv int64
	q.Drv(292000,292141,unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::write(char const*,qint64)
func (q *QIODevice) WriteWithDataLen(data *byte,len int64) int64 {
	var __rv int64
	q.Drv(292000,292142,unsafe.Pointer(&data),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QIODevice::writeData(char const*,qint64)
func (q *QIODevice) WriteData(data *byte,len int64) int64 {
	var __rv int64
	q.Drv(292000,292143,unsafe.Pointer(&data),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
