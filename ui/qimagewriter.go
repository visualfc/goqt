// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QImageWriter_ImageWriterError - QImageWriter::ImageWriterError
type QImageWriter_ImageWriterError uint32
const (
	QImageWriter_UnknownError QImageWriter_ImageWriterError = 0
	QImageWriter_DeviceError QImageWriter_ImageWriterError = 1
	QImageWriter_UnsupportedFormatError QImageWriter_ImageWriterError = 2
)
//struct QImageWriter : QImageWriter
type QImageWriter struct {
	BaseDrv
}
//QImageWriter::QImageWriter()	
func NewImageWriter() *QImageWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),56000,56102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageWriter{}
	_p.SetDriver(__rv,56000,true)
	return _p
} 
//QImageWriter::QImageWriter(QIODevice*,QByteArray const&)	
func NewImageWriterWithDeviceFormat(device QIODeviceInterface,format []byte) *QImageWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),56000,56103,Native(device),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageWriter{}
	_p.SetDriver(__rv,56000,true)
	return _p
} 
//QImageWriter::QImageWriter(QString const&,QByteArray const&)	
func NewImageWriterWithFilenameFormat(fileName string,format []byte) *QImageWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),56000,56104,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageWriter{}
	_p.SetDriver(__rv,56000,true)
	return _p
} 
//QImageWriter::canWrite()
func (q *QImageWriter) CanWrite() bool {
	var __rv bool
	q.Drv(56000,56105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::compression()
func (q *QImageWriter) Compression() int {
	var __rv int
	q.Drv(56000,56106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::description()
func (q *QImageWriter) Description() string {
	var __rv string
	q.Drv(56000,56107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::device()
func (q *QImageWriter) Device() *QIODevice {
	var __rv uintptr
	q.Drv(56000,56108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QImageWriter::error()
func (q *QImageWriter) Error() QImageWriter_ImageWriterError {
	var __rv QImageWriter_ImageWriterError
	q.Drv(56000,56109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::errorString()
func (q *QImageWriter) ErrorString() string {
	var __rv string
	q.Drv(56000,56110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::fileName()
func (q *QImageWriter) FileName() string {
	var __rv string
	q.Drv(56000,56111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::format()
func (q *QImageWriter) Format() []byte {
	var __rv []byte
	q.Drv(56000,56112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::gamma()
func (q *QImageWriter) Gamma() float32 {
	var __rv float32
	q.Drv(56000,56113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::quality()
func (q *QImageWriter) Quality() int {
	var __rv int
	q.Drv(56000,56114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::setCompression(int)
func (q *QImageWriter) SetCompression(compression int)  {
	q.Drv(56000,56115,unsafe.Pointer(&compression),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setDescription(QString const&)
func (q *QImageWriter) SetDescription(description string)  {
	q.Drv(56000,56116,unsafe.Pointer(&description),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setDevice(QIODevice*)
func (q *QImageWriter) SetDevice(device QIODeviceInterface)  {
	q.Drv(56000,56117,Native(device),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setFileName(QString const&)
func (q *QImageWriter) SetFileName(fileName string)  {
	q.Drv(56000,56118,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setFormat(QByteArray const&)
func (q *QImageWriter) SetFormat(format []byte)  {
	q.Drv(56000,56119,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setGamma(float)
func (q *QImageWriter) SetGamma(gamma float32)  {
	q.Drv(56000,56120,unsafe.Pointer(&gamma),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setQuality(int)
func (q *QImageWriter) SetQuality(quality int)  {
	q.Drv(56000,56121,unsafe.Pointer(&quality),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::setText(QString const&,QString const&)
func (q *QImageWriter) SetText(key string,text string)  {
	q.Drv(56000,56122,unsafe.Pointer(&key),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageWriter::supportedImageFormats()	
func QImageWriterSupportedImageFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,56000,56123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::supportedImageFormats()
func (q *QImageWriter) SupportedImageFormats() [][]byte {
	var __rv [][]byte
	q.Drv(56000,56123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::supportsOption(QImageIOHandler::ImageOption)
func (q *QImageWriter) SupportsOption(option QImageIOHandler_ImageOption) bool {
	var __rv bool
	q.Drv(56000,56124,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageWriter::write(QImage const&)
func (q *QImageWriter) Write(image *QImage) bool {
	var __rv bool
	q.Drv(56000,56125,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
