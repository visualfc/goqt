// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QMovie_CacheMode - QMovie::CacheMode
type QMovie_CacheMode uint32
const (
	QMovie_CacheNone QMovie_CacheMode = 0
	QMovie_CacheAll QMovie_CacheMode = 1
)
//enum QMovie_MovieState - QMovie::MovieState
type QMovie_MovieState uint32
const (
	QMovie_NotRunning QMovie_MovieState = 0
	QMovie_Paused QMovie_MovieState = 1
	QMovie_Running QMovie_MovieState = 2
)
//struct QMovie : QMovie
type QMovie struct {
	QObject
}
func (q *QMovie) OnFrameChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(313000,313102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnError(fn func(QImageReader_ImageReaderError)) uintptr {
	var __rv uintptr
	q.Drv(313000,313103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnResized(fn func(*QSize)) uintptr {
	var __rv uintptr
	q.Drv(313000,313104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnStateChanged(fn func(QMovie_MovieState)) uintptr {
	var __rv uintptr
	q.Drv(313000,313105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnStarted(fn func()) uintptr {
	var __rv uintptr
	q.Drv(313000,313106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnUpdated(fn func(*QRect)) uintptr {
	var __rv uintptr
	q.Drv(313000,313107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QMovie) OnFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(313000,313108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QMovie::QMovie()	
func NewMovie() *QMovie {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),313000,313109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMovie{}
	_p.SetDriver(__rv,313000,false)
	return _p
} 
//QMovie::QMovie(QObject*)	
func NewMovieWithParent(parent QObjectInterface) *QMovie {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),313000,313110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMovie{}
	_p.SetDriver(__rv,313000,false)
	return _p
} 
//QMovie::QMovie(QIODevice*,QByteArray const&,QObject*)	
func NewMovieWithDeviceFormatParent(device QIODeviceInterface,format []byte,parent QObjectInterface) *QMovie {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),313000,313111,Native(device),unsafe.Pointer(&format),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMovie{}
	_p.SetDriver(__rv,313000,false)
	return _p
} 
//QMovie::QMovie(QString const&,QByteArray const&,QObject*)	
func NewMovieWithFilenameFormatParent(fileName string,format []byte,parent QObjectInterface) *QMovie {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),313000,313112,unsafe.Pointer(&fileName),unsafe.Pointer(&format),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QMovie{}
	_p.SetDriver(__rv,313000,false)
	return _p
} 
//QMovie::backgroundColor()
func (q *QMovie) BackgroundColor() *QColor {
	var __rv uintptr
	q.Drv(313000,313113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QMovie::cacheMode()
func (q *QMovie) CacheMode() QMovie_CacheMode {
	var __rv QMovie_CacheMode
	q.Drv(313000,313114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::currentFrameNumber()
func (q *QMovie) CurrentFrameNumber() int {
	var __rv int
	q.Drv(313000,313115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::currentImage()
func (q *QMovie) CurrentImage() *QImage {
	var __rv uintptr
	q.Drv(313000,313116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QMovie::currentPixmap()
func (q *QMovie) CurrentPixmap() *QPixmap {
	var __rv uintptr
	q.Drv(313000,313117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QMovie::device()
func (q *QMovie) Device() *QIODevice {
	var __rv uintptr
	q.Drv(313000,313118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QIODevice{}
	_rp.SetDriver(__rv,292000,false)
	return _rp
}	
//QMovie::fileName()
func (q *QMovie) FileName() string {
	var __rv string
	q.Drv(313000,313119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::format()
func (q *QMovie) Format() []byte {
	var __rv []byte
	q.Drv(313000,313120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::frameCount()
func (q *QMovie) FrameCount() int {
	var __rv int
	q.Drv(313000,313121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::frameRect()
func (q *QMovie) FrameRect() *QRect {
	var __rv uintptr
	q.Drv(313000,313122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QMovie::isValid()
func (q *QMovie) IsValid() bool {
	var __rv bool
	q.Drv(313000,313123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::jumpToFrame(int)
func (q *QMovie) JumpToFrame(frameNumber int) bool {
	var __rv bool
	q.Drv(313000,313124,unsafe.Pointer(&frameNumber),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::jumpToNextFrame()
func (q *QMovie) JumpToNextFrame() bool {
	var __rv bool
	q.Drv(313000,313125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::loopCount()
func (q *QMovie) LoopCount() int {
	var __rv int
	q.Drv(313000,313126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::nextFrameDelay()
func (q *QMovie) NextFrameDelay() int {
	var __rv int
	q.Drv(313000,313127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::scaledSize()
func (q *QMovie) ScaledSize() *QSize {
	var __rv uintptr
	q.Drv(313000,313128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QMovie::setBackgroundColor(QColor const&)
func (q *QMovie) SetBackgroundColor(color *QColor)  {
	q.Drv(313000,313129,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setCacheMode(QMovie::CacheMode)
func (q *QMovie) SetCacheMode(mode QMovie_CacheMode)  {
	q.Drv(313000,313130,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setDevice(QIODevice*)
func (q *QMovie) SetDevice(device QIODeviceInterface)  {
	q.Drv(313000,313131,Native(device),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setFileName(QString const&)
func (q *QMovie) SetFileName(fileName string)  {
	q.Drv(313000,313132,unsafe.Pointer(&fileName),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setFormat(QByteArray const&)
func (q *QMovie) SetFormat(format []byte)  {
	q.Drv(313000,313133,unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setPaused(bool)
func (q *QMovie) SetPaused(paused bool)  {
	q.Drv(313000,313134,unsafe.Pointer(&paused),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setScaledSize(QSize const&)
func (q *QMovie) SetScaledSize(size *QSize)  {
	q.Drv(313000,313135,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::setSpeed(int)
func (q *QMovie) SetSpeed(percentSpeed int)  {
	q.Drv(313000,313136,unsafe.Pointer(&percentSpeed),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::speed()
func (q *QMovie) Speed() int {
	var __rv int
	q.Drv(313000,313137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::start()
func (q *QMovie) Start()  {
	q.Drv(313000,313138,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::state()
func (q *QMovie) State() QMovie_MovieState {
	var __rv QMovie_MovieState
	q.Drv(313000,313139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::stop()
func (q *QMovie) Stop()  {
	q.Drv(313000,313140,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QMovie::supportedFormats()	
func QMovieSupportedFormats() [][]byte {
	var __rv [][]byte
	DirectQtDrv(nil,313000,313141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QMovie::supportedFormats()
func (q *QMovie) SupportedFormats() [][]byte {
	var __rv [][]byte
	q.Drv(313000,313141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
