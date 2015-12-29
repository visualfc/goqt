// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QImageReader_ImageReaderError - QImageReader::ImageReaderError
type QImageReader_ImageReaderError uint32
const (
	QImageReader_UnknownError QImageReader_ImageReaderError = 0
	QImageReader_FileNotFoundError QImageReader_ImageReaderError = 1
	QImageReader_DeviceError QImageReader_ImageReaderError = 2
	QImageReader_UnsupportedFormatError QImageReader_ImageReaderError = 3
	QImageReader_InvalidDataError QImageReader_ImageReaderError = 4
)
//struct QImageReader : QImageReader
type QImageReader struct {
	BaseDrv
}
//QImageReader::QImageReader()	
func NewImageReader() *QImageReader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),55000,55102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageReader{}
	_p.SetDriver(__rv,55000,true)
	return _p
} 
//QImageReader::QImageReader(QIODevice*,QByteArray const&)	
func NewImageReaderWithDeviceFormat(device QIODeviceInterface,format []byte) *QImageReader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),55000,55103,Native(device),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageReader{}
	_p.SetDriver(__rv,55000,true)
	return _p
} 
//QImageReader::QImageReader(QString const&,QByteArray const&)	
func NewImageReaderWithFilenameFormat(fileName string,format []byte) *QImageReader {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),55000,55104,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImageReader{}
	_p.SetDriver(__rv,55000,true)
	return _p
} 
//QImageReader::autoDetectImageFormat()
func (q *QImageReader) AutoDetectImageFormat() bool {
	var __rv bool
	q.Drv(55000,55105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::backgroundColor()
func (q *QImageReader) BackgroundColor() *QColor {
	var __rv uintptr
	q.Drv(55000,55106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QImageReader::canRead()
func (q *QImageReader) CanRead() bool {
	var __rv bool
	q.Drv(55000,55107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::clipRect()
func (q *QImageReader) ClipRect() *QRect {
	var __rv uintptr
	q.Drv(55000,55108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QImageReader::currentImageNumber()
func (q *QImageReader) CurrentImageNumber() int {
	var __rv int
	q.Drv(55000,55109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::currentImageRect()
func (q *QImageReader) CurrentImageRect() *QRect {
	var __rv uintptr
	q.Drv(55000,55110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QImageReader::decideFormatFromContent()
func (q *QImageReader) DecideFormatFromContent() bool {
	var __rv bool
	q.Drv(55000,55111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::device()
func (q *QImageReader) Device() *QIODevice {
	var __rv uintptr
	q.Drv(55000,55112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QImageReader::error()
func (q *QImageReader) Error() QImageReader_ImageReaderError {
	var __rv QImageReader_ImageReaderError
	q.Drv(55000,55113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::errorString()
func (q *QImageReader) ErrorString() string {
	var __rv string
	q.Drv(55000,55114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::fileName()
func (q *QImageReader) FileName() string {
	var __rv string
	q.Drv(55000,55115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::format()
func (q *QImageReader) Format() []byte {
	var __rv []byte
	q.Drv(55000,55116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageCount()
func (q *QImageReader) ImageCount() int {
	var __rv int
	q.Drv(55000,55117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageFormat()
func (q *QImageReader) ImageFormat() QImage_Format {
	var __rv QImage_Format
	q.Drv(55000,55118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageFormat(QIODevice*)	
func QImageReaderImageFormatWithDevice(device QIODeviceInterface) []byte {
	var __rv []byte
	DirectQtDrv(nil,55000,55119,Native(device),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageFormat(QIODevice*)
func (q *QImageReader) ImageFormatWithDevice(device QIODeviceInterface) []byte {
	var __rv []byte
	q.Drv(55000,55119,Native(device),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageFormat(QString const&)	
func QImageReaderImageFormatWithFilename(fileName string) []byte {
	var __rv []byte
	DirectQtDrv(nil,55000,55120,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::imageFormat(QString const&)
func (q *QImageReader) ImageFormatWithFilename(fileName string) []byte {
	var __rv []byte
	q.Drv(55000,55120,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::jumpToImage(int)
func (q *QImageReader) JumpToImage(imageNumber int) bool {
	var __rv bool
	q.Drv(55000,55121,unsafe.Pointer(&imageNumber),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::jumpToNextImage()
func (q *QImageReader) JumpToNextImage() bool {
	var __rv bool
	q.Drv(55000,55122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::loopCount()
func (q *QImageReader) LoopCount() int {
	var __rv int
	q.Drv(55000,55123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::nextImageDelay()
func (q *QImageReader) NextImageDelay() int {
	var __rv int
	q.Drv(55000,55124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::quality()
func (q *QImageReader) Quality() int {
	var __rv int
	q.Drv(55000,55125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::read()
func (q *QImageReader) Read() *QImage {
	var __rv uintptr
	q.Drv(55000,55126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImageReader::read(QImage*)
func (q *QImageReader) ReadWithImage(image *QImage) bool {
	var __rv bool
	q.Drv(55000,55127,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::scaledClipRect()
func (q *QImageReader) ScaledClipRect() *QRect {
	var __rv uintptr
	q.Drv(55000,55128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QImageReader::scaledSize()
func (q *QImageReader) ScaledSize() *QSize {
	var __rv uintptr
	q.Drv(55000,55129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QImageReader::setAutoDetectImageFormat(bool)
func (q *QImageReader) SetAutoDetectImageFormat(enabled bool)  {
	q.Drv(55000,55130,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setBackgroundColor(QColor const&)
func (q *QImageReader) SetBackgroundColor(color *QColor)  {
	q.Drv(55000,55131,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setClipRect(QRect const&)
func (q *QImageReader) SetClipRect(rect *QRect)  {
	q.Drv(55000,55132,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setDecideFormatFromContent(bool)
func (q *QImageReader) SetDecideFormatFromContent(ignored bool)  {
	q.Drv(55000,55133,unsafe.Pointer(&ignored),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setDevice(QIODevice*)
func (q *QImageReader) SetDevice(device QIODeviceInterface)  {
	q.Drv(55000,55134,Native(device),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setFileName(QString const&)
func (q *QImageReader) SetFileName(fileName string)  {
	q.Drv(55000,55135,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setFormat(QByteArray const&)
func (q *QImageReader) SetFormat(format []byte)  {
	q.Drv(55000,55136,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setQuality(int)
func (q *QImageReader) SetQuality(quality int)  {
	q.Drv(55000,55137,unsafe.Pointer(&quality),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setScaledClipRect(QRect const&)
func (q *QImageReader) SetScaledClipRect(rect *QRect)  {
	q.Drv(55000,55138,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::setScaledSize(QSize const&)
func (q *QImageReader) SetScaledSize(size *QSize)  {
	q.Drv(55000,55139,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImageReader::size()
func (q *QImageReader) Size() *QSize {
	var __rv uintptr
	q.Drv(55000,55140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QImageReader::supportedImageFormats()	
func QImageReaderSupportedImageFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,55000,55141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::supportedImageFormats()
func (q *QImageReader) SupportedImageFormats() [][]byte {
	var __rv [][]byte
	q.Drv(55000,55141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::supportsAnimation()
func (q *QImageReader) SupportsAnimation() bool {
	var __rv bool
	q.Drv(55000,55142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::supportsOption(QImageIOHandler::ImageOption)
func (q *QImageReader) SupportsOption(option QImageIOHandler_ImageOption) bool {
	var __rv bool
	q.Drv(55000,55143,unsafe.Pointer(&option),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::text(QString const&)
func (q *QImageReader) Text(key string) string {
	var __rv string
	q.Drv(55000,55144,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImageReader::textKeys()
func (q *QImageReader) TextKeys() []string {
	var __rv []string
	q.Drv(55000,55145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
