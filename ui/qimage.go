// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QImage_InvertMode - QImage::InvertMode
type QImage_InvertMode uint32
const (
	QImage_InvertRgb QImage_InvertMode = 0
	QImage_InvertRgba QImage_InvertMode = 1
)
//enum QImage_Format - QImage::Format
type QImage_Format uint32
const (
	QImage_Format_Invalid QImage_Format = 0
	QImage_Format_Mono QImage_Format = 1
	QImage_Format_MonoLSB QImage_Format = 2
	QImage_Format_Indexed8 QImage_Format = 3
	QImage_Format_RGB32 QImage_Format = 4
	QImage_Format_ARGB32 QImage_Format = 5
	QImage_Format_ARGB32_Premultiplied QImage_Format = 6
	QImage_Format_RGB16 QImage_Format = 7
	QImage_Format_ARGB8565_Premultiplied QImage_Format = 8
	QImage_Format_RGB666 QImage_Format = 9
	QImage_Format_ARGB6666_Premultiplied QImage_Format = 10
	QImage_Format_RGB555 QImage_Format = 11
	QImage_Format_ARGB8555_Premultiplied QImage_Format = 12
	QImage_Format_RGB888 QImage_Format = 13
	QImage_Format_RGB444 QImage_Format = 14
	QImage_Format_ARGB4444_Premultiplied QImage_Format = 15
	QImage_NImageFormats QImage_Format = 16
)
//struct QImage : QImage
type QImage struct {
	QPaintDevice
}
//QImage::QImage()	
func NewImage() *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(QImage const&)	
func NewImageCopy(value *QImage) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(char const*[])	
func NewImageWithXpm(xpm [][]byte) *QImage {
	_xpm:=drvBytesArrayToC(xpm)
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53104,unsafe.Pointer(&_xpm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(QSize const&,QImage::Format)	
func NewImageWithSizeFormat(size *QSize,format QImage_Format) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53105,Native(size),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(QString const&,char const*)	
func NewImageWithFilenameFormat(fileName string,format string) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53106,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(int,int,QImage::Format)	
func NewImageWithWidthHeightFormat(width int,height int,format QImage_Format) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53107,unsafe.Pointer(&width),unsafe.Pointer(&height),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(unsigned char const*,int,int,QImage::Format)	
func NewImageWithDataWidthHeightFormat(data *byte,width int,height int,format QImage_Format) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53108,unsafe.Pointer(&data),unsafe.Pointer(&width),unsafe.Pointer(&height),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::QImage(unsigned char const*,int,int,int,QImage::Format)	
func NewImageWithDataWidthHeightBytesperlineFormat(data *byte,width int,height int,bytesPerLine int,format QImage_Format) *QImage {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),53000,53109,unsafe.Pointer(&data),unsafe.Pointer(&width),unsafe.Pointer(&height),unsafe.Pointer(&bytesPerLine),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QImage{}
	_p.SetDriver(__rv,53000,true)
	return _p
} 
//QImage::allGray()
func (q *QImage) AllGray() bool {
	var __rv bool
	q.Drv(53000,53110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::bitPlaneCount()
func (q *QImage) BitPlaneCount() int {
	var __rv int
	q.Drv(53000,53111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::bits()
func (q *QImage) Bits() *byte {
	var __rv *byte
	q.Drv(53000,53112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::byteCount()
func (q *QImage) ByteCount() int {
	var __rv int
	q.Drv(53000,53113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::bytesPerLine()
func (q *QImage) BytesPerLine() int {
	var __rv int
	q.Drv(53000,53114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::cacheKey()
func (q *QImage) CacheKey() int64 {
	var __rv int64
	q.Drv(53000,53115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::color(int)
func (q *QImage) Color(i int) uint {
	var __rv uint
	q.Drv(53000,53116,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::colorCount()
func (q *QImage) ColorCount() int {
	var __rv int
	q.Drv(53000,53117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::colorTable()
func (q *QImage) ColorTable() []QRgb {
	var __rv []QRgb
	q.Drv(53000,53118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::constBits()
func (q *QImage) ConstBits() *byte {
	var __rv *byte
	q.Drv(53000,53119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::constScanLine(int)
func (q *QImage) ConstScanLine(value int) *byte {
	var __rv *byte
	q.Drv(53000,53120,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::convertToFormat(QImage::Format)
func (q *QImage) ConvertToFormat(f QImage_Format) *QImage {
	var __rv uintptr
	q.Drv(53000,53121,unsafe.Pointer(&f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::convertToFormat(QImage::Format,QFlags<Qt::ImageConversionFlag>)
func (q *QImage) ConvertToFormatWithFormatFlags(f QImage_Format,flags Qt_ImageConversionFlag) *QImage {
	var __rv uintptr
	q.Drv(53000,53122,unsafe.Pointer(&f),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::convertToFormat(QImage::Format,QVector<QRgb> const&,QFlags<Qt::ImageConversionFlag>)
func (q *QImage) ConvertToFormatWithFormatColortableFlags(f QImage_Format,colorTable []QRgb,flags Qt_ImageConversionFlag) *QImage {
	var __rv uintptr
	q.Drv(53000,53123,unsafe.Pointer(&f),unsafe.Pointer(&colorTable),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::copy()
func (q *QImage) Copy() *QImage {
	var __rv uintptr
	q.Drv(53000,53124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::copy(QRect const&)
func (q *QImage) CopyWithRect(rect *QRect) *QImage {
	var __rv uintptr
	q.Drv(53000,53125,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::copy(int,int,int,int)
func (q *QImage) CopyWithXYWidthHeight(x int,y int,w int,h int) *QImage {
	var __rv uintptr
	q.Drv(53000,53126,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::createAlphaMask()
func (q *QImage) CreateAlphaMask() *QImage {
	var __rv uintptr
	q.Drv(53000,53127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::createAlphaMask(QFlags<Qt::ImageConversionFlag>)
func (q *QImage) CreateAlphaMaskWithFlags(flags Qt_ImageConversionFlag) *QImage {
	var __rv uintptr
	q.Drv(53000,53128,unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::createHeuristicMask(bool)
func (q *QImage) CreateHeuristicMask(clipTight bool) *QImage {
	var __rv uintptr
	q.Drv(53000,53129,unsafe.Pointer(&clipTight),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::createMaskFromColor(unsigned int)
func (q *QImage) CreateMaskFromColor(color uint) *QImage {
	var __rv uintptr
	q.Drv(53000,53130,unsafe.Pointer(&color),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::createMaskFromColor(unsigned int,Qt::MaskMode)
func (q *QImage) CreateMaskFromColorWithColorMode(color uint,mode Qt_MaskMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53131,unsafe.Pointer(&color),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::depth()
func (q *QImage) Depth() int {
	var __rv int
	q.Drv(53000,53132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::detach()
func (q *QImage) Detach()  {
	q.Drv(53000,53133,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::devType()
func (q *QImage) DevType() int {
	var __rv int
	q.Drv(53000,53134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::dotsPerMeterX()
func (q *QImage) DotsPerMeterX() int {
	var __rv int
	q.Drv(53000,53135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::dotsPerMeterY()
func (q *QImage) DotsPerMeterY() int {
	var __rv int
	q.Drv(53000,53136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::fill(unsigned int)
func (q *QImage) Fill(pixel uint)  {
	q.Drv(53000,53137,unsafe.Pointer(&pixel),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::format()
func (q *QImage) Format() QImage_Format {
	var __rv QImage_Format
	q.Drv(53000,53138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::fromData(QByteArray const&)	
func QImageFromData(data []byte) *QImage {
	var __rv uintptr
	DirectQtDrv(nil,53000,53139,unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::fromData(QByteArray const&)
func (q *QImage) FromData(data []byte) *QImage {
	var __rv uintptr
	q.Drv(53000,53139,unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::fromData(QByteArray const&,char const*)	
func QImageFromDataWithDataFormat(data []byte,format string) *QImage {
	var __rv uintptr
	DirectQtDrv(nil,53000,53140,unsafe.Pointer(&data),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::fromData(QByteArray const&,char const*)
func (q *QImage) FromDataWithDataFormat(data []byte,format string) *QImage {
	var __rv uintptr
	q.Drv(53000,53140,unsafe.Pointer(&data),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::fromData(unsigned char const*,int,char const*)	
func QImageFromDataWithDataSizeFormat(data *byte,size int,format string) *QImage {
	var __rv uintptr
	DirectQtDrv(nil,53000,53141,unsafe.Pointer(&data),unsafe.Pointer(&size),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::fromData(unsigned char const*,int,char const*)
func (q *QImage) FromDataWithDataSizeFormat(data *byte,size int,format string) *QImage {
	var __rv uintptr
	q.Drv(53000,53141,unsafe.Pointer(&data),unsafe.Pointer(&size),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::hasAlphaChannel()
func (q *QImage) HasAlphaChannel() bool {
	var __rv bool
	q.Drv(53000,53142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::height()
func (q *QImage) Height() int {
	var __rv int
	q.Drv(53000,53143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::invertPixels()
func (q *QImage) InvertPixels()  {
	q.Drv(53000,53144,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::invertPixels(QImage::InvertMode)
func (q *QImage) InvertPixelsWithInvertmode(value QImage_InvertMode)  {
	q.Drv(53000,53145,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::isDetached()
func (q *QImage) IsDetached() bool {
	var __rv bool
	q.Drv(53000,53146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::isGrayscale()
func (q *QImage) IsGrayscale() bool {
	var __rv bool
	q.Drv(53000,53147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::isNull()
func (q *QImage) IsNull() bool {
	var __rv bool
	q.Drv(53000,53148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::load(QString const&)
func (q *QImage) Load(fileName string) bool {
	var __rv bool
	q.Drv(53000,53149,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::load(QIODevice*,char const*)
func (q *QImage) LoadWithDeviceFormat(device QIODeviceInterface,format string) bool {
	var __rv bool
	q.Drv(53000,53150,Native(device),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::load(QString const&,char const*)
func (q *QImage) LoadWithFilenameFormat(fileName string,format string) bool {
	var __rv bool
	q.Drv(53000,53151,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::loadFromData(QByteArray const&)
func (q *QImage) LoadFromData(data []byte) bool {
	var __rv bool
	q.Drv(53000,53152,unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::loadFromData(QByteArray const&,char const*)
func (q *QImage) LoadFromDataWithDataAformat(data []byte,aformat string) bool {
	var __rv bool
	q.Drv(53000,53153,unsafe.Pointer(&data),unsafe.Pointer(&aformat),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::loadFromData(unsigned char const*,int,char const*)
func (q *QImage) LoadFromDataWithBufLenFormat(buf *byte,len int,format string) bool {
	var __rv bool
	q.Drv(53000,53154,unsafe.Pointer(&buf),unsafe.Pointer(&len),unsafe.Pointer(&format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::mirrored()
func (q *QImage) Mirrored() *QImage {
	var __rv uintptr
	q.Drv(53000,53155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::mirrored(bool,bool)
func (q *QImage) MirroredWithHorizontallyVertically(horizontally bool,vertically bool) *QImage {
	var __rv uintptr
	q.Drv(53000,53156,unsafe.Pointer(&horizontally),unsafe.Pointer(&vertically),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::offset()
func (q *QImage) Offset() *QPoint {
	var __rv uintptr
	q.Drv(53000,53157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QImage::paintEngine()
func (q *QImage) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(53000,53158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QImage::pixel(QPoint const&)
func (q *QImage) Pixel(pt *QPoint) uint {
	var __rv uint
	q.Drv(53000,53159,Native(pt),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::pixel(int,int)
func (q *QImage) PixelWithXY(x int,y int) uint {
	var __rv uint
	q.Drv(53000,53160,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::pixelIndex(QPoint const&)
func (q *QImage) PixelIndex(pt *QPoint) int {
	var __rv int
	q.Drv(53000,53161,Native(pt),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::pixelIndex(int,int)
func (q *QImage) PixelIndexWithXY(x int,y int) int {
	var __rv int
	q.Drv(53000,53162,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::rect()
func (q *QImage) Rect() *QRect {
	var __rv uintptr
	q.Drv(53000,53163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QImage::rgbSwapped()
func (q *QImage) RgbSwapped() *QImage {
	var __rv uintptr
	q.Drv(53000,53164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::save(QIODevice*)
func (q *QImage) Save(device QIODeviceInterface) bool {
	var __rv bool
	q.Drv(53000,53165,Native(device),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::save(QString const&)
func (q *QImage) SaveWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(53000,53166,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::save(QIODevice*,char const*,int)
func (q *QImage) SaveWithDeviceFormatQuality(device QIODeviceInterface,format string,quality int) bool {
	var __rv bool
	q.Drv(53000,53167,Native(device),unsafe.Pointer(&format),unsafe.Pointer(&quality),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::save(QString const&,char const*,int)
func (q *QImage) SaveWithFilenameFormatQuality(fileName string,format string,quality int) bool {
	var __rv bool
	q.Drv(53000,53168,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&quality),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::scaled(QSize const&)
func (q *QImage) Scaled(s *QSize) *QImage {
	var __rv uintptr
	q.Drv(53000,53169,Native(s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaled(QSize const&,Qt::AspectRatioMode,Qt::TransformationMode)
func (q *QImage) ScaledWithSizeAspectmodeMode(s *QSize,aspectMode Qt_AspectRatioMode,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53170,Native(s),unsafe.Pointer(&aspectMode),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaled(int,int,Qt::AspectRatioMode,Qt::TransformationMode)
func (q *QImage) ScaledWithWidthHeightAspectmodeMode(w int,h int,aspectMode Qt_AspectRatioMode,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53171,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&aspectMode),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaledToHeight(int)
func (q *QImage) ScaledToHeight(h int) *QImage {
	var __rv uintptr
	q.Drv(53000,53172,unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaledToHeight(int,Qt::TransformationMode)
func (q *QImage) ScaledToHeightWithHeightMode(h int,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53173,unsafe.Pointer(&h),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaledToWidth(int)
func (q *QImage) ScaledToWidth(w int) *QImage {
	var __rv uintptr
	q.Drv(53000,53174,unsafe.Pointer(&w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scaledToWidth(int,Qt::TransformationMode)
func (q *QImage) ScaledToWidthWithWidthMode(w int,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53175,unsafe.Pointer(&w),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::scanLine(int)
func (q *QImage) ScanLine(value int) *byte {
	var __rv *byte
	q.Drv(53000,53176,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::setColor(int,unsigned int)
func (q *QImage) SetColor(i int,c uint)  {
	q.Drv(53000,53177,unsafe.Pointer(&i),unsafe.Pointer(&c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setColorCount(int)
func (q *QImage) SetColorCount(value int)  {
	q.Drv(53000,53178,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setColorTable(QVector<QRgb> const)
func (q *QImage) SetColorTable(colors []QRgb)  {
	q.Drv(53000,53179,unsafe.Pointer(&colors),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setDotsPerMeterX(int)
func (q *QImage) SetDotsPerMeterX(value int)  {
	q.Drv(53000,53180,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setDotsPerMeterY(int)
func (q *QImage) SetDotsPerMeterY(value int)  {
	q.Drv(53000,53181,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setOffset(QPoint const&)
func (q *QImage) SetOffset(value *QPoint)  {
	q.Drv(53000,53182,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setPixel(QPoint const&,unsigned int)
func (q *QImage) SetPixelWithPointIndex_or_rgb(pt *QPoint,index_or_rgb uint)  {
	q.Drv(53000,53183,Native(pt),unsafe.Pointer(&index_or_rgb),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setPixel(int,int,unsigned int)
func (q *QImage) SetPixelWithXYIndex_or_rgb(x int,y int,index_or_rgb uint)  {
	q.Drv(53000,53184,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&index_or_rgb),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::setText(QString const&,QString const&)
func (q *QImage) SetText(key string,value string)  {
	q.Drv(53000,53185,unsafe.Pointer(&key),unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QImage::size()
func (q *QImage) Size() *QSize {
	var __rv uintptr
	q.Drv(53000,53186,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QImage::text()
func (q *QImage) Text() string {
	var __rv string
	q.Drv(53000,53187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::text(QString const&)
func (q *QImage) TextWithKey(key string) string {
	var __rv string
	q.Drv(53000,53188,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::textKeys()
func (q *QImage) TextKeys() []string {
	var __rv []string
	q.Drv(53000,53189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::transformed(QMatrix const&)
func (q *QImage) Transformed(matrix *QMatrix) *QImage {
	var __rv uintptr
	q.Drv(53000,53190,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::transformed(QTransform const&)
func (q *QImage) TransformedWithTransform(matrix *QTransform) *QImage {
	var __rv uintptr
	q.Drv(53000,53191,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::transformed(QMatrix const&,Qt::TransformationMode)
func (q *QImage) TransformedWithMatrixMode(matrix *QMatrix,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53192,Native(matrix),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::transformed(QTransform const&,Qt::TransformationMode)
func (q *QImage) TransformedWithTransformMode(matrix *QTransform,mode Qt_TransformationMode) *QImage {
	var __rv uintptr
	q.Drv(53000,53193,Native(matrix),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QImage::trueMatrix(QMatrix const&,int,int)	
func QImageTrueMatrixWithMatrixWidthHeight(value2 *QMatrix,w int,h int) *QMatrix {
	var __rv uintptr
	DirectQtDrv(nil,53000,53194,Native(value2),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QImage::trueMatrix(QMatrix const&,int,int)
func (q *QImage) TrueMatrixWithMatrixWidthHeight(value2 *QMatrix,w int,h int) *QMatrix {
	var __rv uintptr
	q.Drv(53000,53194,Native(value2),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QImage::trueMatrix(QTransform const&,int,int)	
func QImageTrueMatrixWithTransformWidthHeight(value2 *QTransform,w int,h int) *QTransform {
	var __rv uintptr
	DirectQtDrv(nil,53000,53195,Native(value2),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QImage::trueMatrix(QTransform const&,int,int)
func (q *QImage) TrueMatrixWithTransformWidthHeight(value2 *QTransform,w int,h int) *QTransform {
	var __rv uintptr
	q.Drv(53000,53195,Native(value2),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QImage::valid(QPoint const&)
func (q *QImage) Valid(pt *QPoint) bool {
	var __rv bool
	q.Drv(53000,53196,Native(pt),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::valid(int,int)
func (q *QImage) ValidWithXY(x int,y int) bool {
	var __rv bool
	q.Drv(53000,53197,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QImage::width()
func (q *QImage) Width() int {
	var __rv int
	q.Drv(53000,53198,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
