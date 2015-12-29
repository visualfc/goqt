// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextDocumentWriter : QTextDocumentWriter
type QTextDocumentWriter struct {
	BaseDrv
}
//QTextDocumentWriter::QTextDocumentWriter()	
func NewTextDocumentWriter() *QTextDocumentWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),148000,148102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentWriter{}
	_p.SetDriver(__rv,148000,true)
	return _p
} 
//QTextDocumentWriter::QTextDocumentWriter(QIODevice*,QByteArray const&)	
func NewTextDocumentWriterWithDeviceFormat(device QIODeviceInterface,format []byte) *QTextDocumentWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),148000,148103,Native(device),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentWriter{}
	_p.SetDriver(__rv,148000,true)
	return _p
} 
//QTextDocumentWriter::QTextDocumentWriter(QString const&,QByteArray const&)	
func NewTextDocumentWriterWithFilenameFormat(fileName string,format []byte) *QTextDocumentWriter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),148000,148104,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocumentWriter{}
	_p.SetDriver(__rv,148000,true)
	return _p
} 
//QTextDocumentWriter::codec()
func (q *QTextDocumentWriter) Codec() *QTextCodec {
	var __rv uintptr
	q.Drv(148000,148105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCodec{}
	_rp.SetDriver(__rv,143000,true)
	return _rp
}	
//QTextDocumentWriter::device()
func (q *QTextDocumentWriter) Device() *QIODevice {
	var __rv uintptr
	q.Drv(148000,148106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QTextDocumentWriter::fileName()
func (q *QTextDocumentWriter) FileName() string {
	var __rv string
	q.Drv(148000,148107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentWriter::format()
func (q *QTextDocumentWriter) Format() []byte {
	var __rv []byte
	q.Drv(148000,148108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentWriter::setCodec(QTextCodec*)
func (q *QTextDocumentWriter) SetCodec(codec *QTextCodec)  {
	q.Drv(148000,148109,Native(codec),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocumentWriter::setDevice(QIODevice*)
func (q *QTextDocumentWriter) SetDevice(device QIODeviceInterface)  {
	q.Drv(148000,148110,Native(device),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocumentWriter::setFileName(QString const&)
func (q *QTextDocumentWriter) SetFileName(fileName string)  {
	q.Drv(148000,148111,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocumentWriter::setFormat(QByteArray const&)
func (q *QTextDocumentWriter) SetFormat(format []byte)  {
	q.Drv(148000,148112,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocumentWriter::supportedDocumentFormats()	
func QTextDocumentWriterSupportedDocumentFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,148000,148113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentWriter::supportedDocumentFormats()
func (q *QTextDocumentWriter) SupportedDocumentFormats() [][]byte {
	var __rv [][]byte
	q.Drv(148000,148113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentWriter::write(QTextDocument const*)
func (q *QTextDocumentWriter) Write(document *QTextDocument) bool {
	var __rv bool
	q.Drv(148000,148114,Native(document),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocumentWriter::write(QTextDocumentFragment const&)
func (q *QTextDocumentWriter) WriteWithFragment(fragment *QTextDocumentFragment) bool {
	var __rv bool
	q.Drv(148000,148115,Native(fragment),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
