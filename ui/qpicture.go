// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPicture : QPicture
type QPicture struct {
	QPaintDevice
}
//QPicture::QPicture()	
func NewPicture() *QPicture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),94000,94102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPicture{}
	_p.SetDriver(__rv,94000,true)
	return _p
} 
//QPicture::QPicture(QPicture const&)	
func NewPictureCopy(value *QPicture) *QPicture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),94000,94103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPicture{}
	_p.SetDriver(__rv,94000,true)
	return _p
} 
//QPicture::QPicture(int)	
func NewPictureWithFormatversion(formatVersion int) *QPicture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),94000,94104,unsafe.Pointer(&formatVersion),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPicture{}
	_p.SetDriver(__rv,94000,true)
	return _p
} 
//QPicture::boundingRect()
func (q *QPicture) BoundingRect() *QRect {
	var __rv uintptr
	q.Drv(94000,94105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPicture::data()
func (q *QPicture) Data() string {
	var __rv string
	q.Drv(94000,94106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::detach()
func (q *QPicture) Detach()  {
	q.Drv(94000,94107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPicture::devType()
func (q *QPicture) DevType() int {
	var __rv int
	q.Drv(94000,94108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::inputFormatList()	
func QPictureInputFormatList() []string {
	var __rv []string
	DirectQtDrv(nil,94000,94109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::inputFormatList()
func (q *QPicture) InputFormatList() []string {
	var __rv []string
	q.Drv(94000,94109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::inputFormats()	
func QPictureInputFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,94000,94110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::inputFormats()
func (q *QPicture) InputFormats() [][]byte {
	var __rv [][]byte
	q.Drv(94000,94110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::isDetached()
func (q *QPicture) IsDetached() bool {
	var __rv bool
	q.Drv(94000,94111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::isNull()
func (q *QPicture) IsNull() bool {
	var __rv bool
	q.Drv(94000,94112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::load(QIODevice*)
func (q *QPicture) Load(dev QIODeviceInterface) bool {
	var __rv bool
	q.Drv(94000,94113,Native(dev),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::load(QString const&)
func (q *QPicture) LoadWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(94000,94114,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::load(QIODevice*,char const*)
func (q *QPicture) LoadWithDevFormat(dev QIODeviceInterface,format string) bool {
	var __rv bool
	q.Drv(94000,94115,Native(dev),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::load(QString const&,char const*)
func (q *QPicture) LoadWithFilenameFormat(fileName string,format string) bool {
	var __rv bool
	q.Drv(94000,94116,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::outputFormatList()	
func QPictureOutputFormatList() []string {
	var __rv []string
	DirectQtDrv(nil,94000,94117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::outputFormatList()
func (q *QPicture) OutputFormatList() []string {
	var __rv []string
	q.Drv(94000,94117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::outputFormats()	
func QPictureOutputFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,94000,94118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::outputFormats()
func (q *QPicture) OutputFormats() [][]byte {
	var __rv [][]byte
	q.Drv(94000,94118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::paintEngine()
func (q *QPicture) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(94000,94119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QPicture::pictureFormat(QString const&)	
func QPicturePictureFormat(fileName string) string {
	var __rv string
	DirectQtDrv(nil,94000,94120,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::pictureFormat(QString const&)
func (q *QPicture) PictureFormat(fileName string) string {
	var __rv string
	q.Drv(94000,94120,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::play(QPainter*)
func (q *QPicture) Play(p *QPainter) bool {
	var __rv bool
	q.Drv(94000,94121,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::save(QIODevice*)
func (q *QPicture) Save(dev QIODeviceInterface) bool {
	var __rv bool
	q.Drv(94000,94122,Native(dev),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::save(QString const&)
func (q *QPicture) SaveWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(94000,94123,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::save(QIODevice*,char const*)
func (q *QPicture) SaveWithDevFormat(dev QIODeviceInterface,format string) bool {
	var __rv bool
	q.Drv(94000,94124,Native(dev),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::save(QString const&,char const*)
func (q *QPicture) SaveWithFilenameFormat(fileName string,format string) bool {
	var __rv bool
	q.Drv(94000,94125,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPicture::setBoundingRect(QRect const&)
func (q *QPicture) SetBoundingRect(r *QRect)  {
	q.Drv(94000,94126,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPicture::setData(char const*,unsigned int)
func (q *QPicture) SetData(data string,size uint)  {
	q.Drv(94000,94127,unsafe.Pointer(&data),unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPicture::size()
func (q *QPicture) Size() uint {
	var __rv uint
	q.Drv(94000,94128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
