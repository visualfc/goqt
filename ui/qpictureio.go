// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPictureIO : QPictureIO
type QPictureIO struct {
	BaseDrv
}
//QPictureIO::QPictureIO()	
func NewPictureIO() *QPictureIO {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),95000,95102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPictureIO{}
	_p.SetDriver(__rv,95000,true)
	return _p
} 
//QPictureIO::QPictureIO(QIODevice*,char const*)	
func NewPictureIOWithIodeviceFormat(ioDevice QIODeviceInterface,format string) *QPictureIO {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),95000,95103,Native(ioDevice),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPictureIO{}
	_p.SetDriver(__rv,95000,true)
	return _p
} 
//QPictureIO::QPictureIO(QString const&,char const*)	
func NewPictureIOWithFilenameFormat(fileName string,format string) *QPictureIO {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),95000,95104,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPictureIO{}
	_p.SetDriver(__rv,95000,true)
	return _p
} 
//QPictureIO::description()
func (q *QPictureIO) Description() string {
	var __rv string
	q.Drv(95000,95105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::fileName()
func (q *QPictureIO) FileName() string {
	var __rv string
	q.Drv(95000,95106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::format()
func (q *QPictureIO) Format() string {
	var __rv string
	q.Drv(95000,95107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::gamma()
func (q *QPictureIO) Gamma() float32 {
	var __rv float32
	q.Drv(95000,95108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::inputFormats()	
func QPictureIOInputFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,95000,95109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::inputFormats()
func (q *QPictureIO) InputFormats() [][]byte {
	var __rv [][]byte
	q.Drv(95000,95109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::ioDevice()
func (q *QPictureIO) IoDevice() *QIODevice {
	var __rv uintptr
	q.Drv(95000,95110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QPictureIO::outputFormats()	
func QPictureIOOutputFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,95000,95111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::outputFormats()
func (q *QPictureIO) OutputFormats() [][]byte {
	var __rv [][]byte
	q.Drv(95000,95111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::parameters()
func (q *QPictureIO) Parameters() string {
	var __rv string
	q.Drv(95000,95112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::picture()
func (q *QPictureIO) Picture() *QPicture {
	var __rv uintptr
	q.Drv(95000,95113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPicture{}
	_rp.SetDriver(__rv,94000,true)
	return _rp
}	
//QPictureIO::pictureFormat(QIODevice*)	
func QPictureIOPictureFormat(value QIODeviceInterface) []byte {
	var __rv []byte
	DirectQtDrv(nil,95000,95114,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::pictureFormat(QIODevice*)
func (q *QPictureIO) PictureFormat(value QIODeviceInterface) []byte {
	var __rv []byte
	q.Drv(95000,95114,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::pictureFormat(QString const&)	
func QPictureIOPictureFormatWithFilename(fileName string) []byte {
	var __rv []byte
	DirectQtDrv(nil,95000,95115,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::pictureFormat(QString const&)
func (q *QPictureIO) PictureFormatWithFilename(fileName string) []byte {
	var __rv []byte
	q.Drv(95000,95115,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::quality()
func (q *QPictureIO) Quality() int {
	var __rv int
	q.Drv(95000,95116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::read()
func (q *QPictureIO) Read() bool {
	var __rv bool
	q.Drv(95000,95117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::setDescription(QString const&)
func (q *QPictureIO) SetDescription(value string)  {
	q.Drv(95000,95118,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setFileName(QString const&)
func (q *QPictureIO) SetFileName(value string)  {
	q.Drv(95000,95119,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setFormat(char const*)
func (q *QPictureIO) SetFormat(value string)  {
	q.Drv(95000,95120,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setGamma(float)
func (q *QPictureIO) SetGamma(value float32)  {
	q.Drv(95000,95121,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setIODevice(QIODevice*)
func (q *QPictureIO) SetIODevice(value QIODeviceInterface)  {
	q.Drv(95000,95122,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setParameters(char const*)
func (q *QPictureIO) SetParameters(value string)  {
	q.Drv(95000,95123,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setPicture(QPicture const&)
func (q *QPictureIO) SetPicture(value *QPicture)  {
	q.Drv(95000,95124,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setQuality(int)
func (q *QPictureIO) SetQuality(value int)  {
	q.Drv(95000,95125,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::setStatus(int)
func (q *QPictureIO) SetStatus(value int)  {
	q.Drv(95000,95126,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPictureIO::status()
func (q *QPictureIO) Status() int {
	var __rv int
	q.Drv(95000,95127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPictureIO::write()
func (q *QPictureIO) Write() bool {
	var __rv bool
	q.Drv(95000,95128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
