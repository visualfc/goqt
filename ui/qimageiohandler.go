// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QImageIOHandler_ImageOption - QImageIOHandler::ImageOption
type QImageIOHandler_ImageOption uint32
const (
	QImageIOHandler_Size QImageIOHandler_ImageOption = 0
	QImageIOHandler_ClipRect QImageIOHandler_ImageOption = 1
	QImageIOHandler_Description QImageIOHandler_ImageOption = 2
	QImageIOHandler_ScaledClipRect QImageIOHandler_ImageOption = 3
	QImageIOHandler_ScaledSize QImageIOHandler_ImageOption = 4
	QImageIOHandler_CompressionRatio QImageIOHandler_ImageOption = 5
	QImageIOHandler_Gamma QImageIOHandler_ImageOption = 6
	QImageIOHandler_Quality QImageIOHandler_ImageOption = 7
	QImageIOHandler_Name QImageIOHandler_ImageOption = 8
	QImageIOHandler_SubType QImageIOHandler_ImageOption = 9
	QImageIOHandler_IncrementalReading QImageIOHandler_ImageOption = 10
	QImageIOHandler_Endianness QImageIOHandler_ImageOption = 11
	QImageIOHandler_Animation QImageIOHandler_ImageOption = 12
	QImageIOHandler_BackgroundColor QImageIOHandler_ImageOption = 13
	QImageIOHandler_ImageFormat QImageIOHandler_ImageOption = 14
)
//struct QImageIOHandler : QImageIOHandler
type QImageIOHandler struct {
	BaseDrv
}
//QImageIOHandler::canRead()
func (q *QImageIOHandler) CanRead() bool {
	var __rv bool
	q.Drv(54000,54102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::currentImageNumber()
func (q *QImageIOHandler) CurrentImageNumber() int {
	var __rv int
	q.Drv(54000,54103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::currentImageRect()
func (q *QImageIOHandler) CurrentImageRect() *QRect {
	var __rv uintptr
	q.Drv(54000,54104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QImageIOHandler::device()
func (q *QImageIOHandler) Device() *QIODevice {
	var __rv uintptr
	q.Drv(54000,54105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QImageIOHandler::format()
func (q *QImageIOHandler) Format() []byte {
	var __rv []byte
	q.Drv(54000,54106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::imageCount()
func (q *QImageIOHandler) ImageCount() int {
	var __rv int
	q.Drv(54000,54107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::jumpToImage(int)
func (q *QImageIOHandler) JumpToImage(imageNumber int) bool {
	var __rv bool
	q.Drv(54000,54108,unsafe.Pointer(&imageNumber),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::jumpToNextImage()
func (q *QImageIOHandler) JumpToNextImage() bool {
	var __rv bool
	q.Drv(54000,54109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::loopCount()
func (q *QImageIOHandler) LoopCount() int {
	var __rv int
	q.Drv(54000,54110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::name()
func (q *QImageIOHandler) Name() []byte {
	var __rv []byte
	q.Drv(54000,54111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::nextImageDelay()
func (q *QImageIOHandler) NextImageDelay() int {
	var __rv int
	q.Drv(54000,54112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::option(QImageIOHandler::ImageOption)
func (q *QImageIOHandler) Option(option QImageIOHandler_ImageOption) *QVariant {
	var __rv uintptr
	q.Drv(54000,54113,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QImageIOHandler::read(QImage*)
func (q *QImageIOHandler) Read(image *QImage) bool {
	var __rv bool
	q.Drv(54000,54114,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::setDevice(QIODevice*)
func (q *QImageIOHandler) SetDevice(device QIODeviceInterface)  {
	q.Drv(54000,54115,Native(device),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageIOHandler::setFormat(QByteArray const&)
func (q *QImageIOHandler) SetFormat(format []byte)  {
	q.Drv(54000,54116,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageIOHandler::setOption(QImageIOHandler::ImageOption,QVariant const&)
func (q *QImageIOHandler) SetOption(option QImageIOHandler_ImageOption,value *QVariant)  {
	q.Drv(54000,54117,unsafe.Pointer(&option),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageIOHandler::supportsOption(QImageIOHandler::ImageOption)
func (q *QImageIOHandler) SupportsOption(option QImageIOHandler_ImageOption) bool {
	var __rv bool
	q.Drv(54000,54118,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageIOHandler::write(QImage const&)
func (q *QImageIOHandler) Write(image *QImage) bool {
	var __rv bool
	q.Drv(54000,54119,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
