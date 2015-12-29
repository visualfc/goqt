// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPixmap_HBitmapFormat - QPixmap::HBitmapFormat
type QPixmap_HBitmapFormat uint32
const (
	QPixmap_NoAlpha QPixmap_HBitmapFormat = 0
	QPixmap_PremultipliedAlpha QPixmap_HBitmapFormat = 1
	QPixmap_Alpha QPixmap_HBitmapFormat = 2
)
//enum QPixmap_Type - QPixmap::Type
type QPixmap_Type uint32
const (
	QPixmap_PixmapType QPixmap_Type = 0
	QPixmap_BitmapType QPixmap_Type = 1
)
//struct QPixmap : QPixmap
type QPixmap struct {
	QPaintDevice
}
//QPixmap::QPixmap()	
func NewPixmap() *QPixmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::QPixmap(QPixmap const&)	
func NewPixmapCopy(value *QPixmap) *QPixmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::QPixmap(QSize const&)	
func NewPixmapWithSize(value *QSize) *QPixmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::QPixmap(char const*[])	
func NewPixmapWithXpm(xpm [][]byte) *QPixmap {
	_xpm:=drvBytesArrayToC(xpm)
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96105,unsafe.Pointer(&_xpm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::QPixmap(int,int)	
func NewPixmapWithWidthHeight(w int,h int) *QPixmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96106,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::QPixmap(QString const&,char const*,QFlags<Qt::ImageConversionFlag>)	
func NewPixmapWithFilenameFormatFlags(fileName string,format string,flags Qt_ImageConversionFlag) *QPixmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),96000,96107,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPixmap{}
	_p.SetDriver(__rv,96000,true)
	return _p
} 
//QPixmap::cacheKey()
func (q *QPixmap) CacheKey() int64 {
	var __rv int64
	q.Drv(96000,96108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::convertFromImage(QImage const&)
func (q *QPixmap) ConvertFromImage(img *QImage) bool {
	var __rv bool
	q.Drv(96000,96109,Native(img),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::convertFromImage(QImage const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) ConvertFromImageWithImgFlags(img *QImage,flags Qt_ImageConversionFlag) bool {
	var __rv bool
	q.Drv(96000,96110,Native(img),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::copy()
func (q *QPixmap) Copy() *QPixmap {
	var __rv uintptr
	q.Drv(96000,96111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::copy(QRect const&)
func (q *QPixmap) CopyWithRect(rect *QRect) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96112,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::copy(int,int,int,int)
func (q *QPixmap) CopyWithXYWidthHeight(x int,y int,width int,height int) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96113,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&width),unsafe.Pointer(&height),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::createHeuristicMask(bool)
func (q *QPixmap) CreateHeuristicMask(clipTight bool) *QBitmap {
	var __rv uintptr
	q.Drv(96000,96114,unsafe.Pointer(&clipTight),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QPixmap::createMaskFromColor(QColor const&)
func (q *QPixmap) CreateMaskFromColor(maskColor *QColor) *QBitmap {
	var __rv uintptr
	q.Drv(96000,96115,Native(maskColor),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QPixmap::createMaskFromColor(QColor const&,Qt::MaskMode)
func (q *QPixmap) CreateMaskFromColorWithColorMode(maskColor *QColor,mode Qt_MaskMode) *QBitmap {
	var __rv uintptr
	q.Drv(96000,96116,Native(maskColor),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QPixmap::defaultDepth()	
func QPixmapDefaultDepth() int {
	var __rv int
	DirectQtDrv(nil,96000,96117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::defaultDepth()
func (q *QPixmap) DefaultDepth() int {
	var __rv int
	q.Drv(96000,96117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::depth()
func (q *QPixmap) Depth() int {
	var __rv int
	q.Drv(96000,96118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::detach()
func (q *QPixmap) Detach()  {
	q.Drv(96000,96119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::devType()
func (q *QPixmap) DevType() int {
	var __rv int
	q.Drv(96000,96120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::fill()
func (q *QPixmap) Fill()  {
	q.Drv(96000,96121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::fill(QColor const&)
func (q *QPixmap) FillWithColor(fillColor *QColor)  {
	q.Drv(96000,96122,Native(fillColor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::fill(QWidget const*,QPoint const&)
func (q *QPixmap) FillWithWidgetOfs(widget QWidgetInterface,ofs *QPoint)  {
	q.Drv(96000,96123,Native(widget),Native(ofs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::fill(QWidget const*,int,int)
func (q *QPixmap) FillWithWidgetXofsYofs(widget QWidgetInterface,xofs int,yofs int)  {
	q.Drv(96000,96124,Native(widget),unsafe.Pointer(&xofs),unsafe.Pointer(&yofs),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::fromImage(QImage const&)	
func QPixmapFromImage(image *QImage) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96125,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImage(QImage const&)
func (q *QPixmap) FromImage(image *QImage) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96125,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImage(QImage const&,QFlags<Qt::ImageConversionFlag>)	
func QPixmapFromImageWithImageFlags(image *QImage,flags Qt_ImageConversionFlag) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96126,Native(image),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImage(QImage const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) FromImageWithImageFlags(image *QImage,flags Qt_ImageConversionFlag) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96126,Native(image),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImageReader(QImageReader*)	
func QPixmapFromImageReader(imageReader *QImageReader) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96127,Native(imageReader),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImageReader(QImageReader*)
func (q *QPixmap) FromImageReader(imageReader *QImageReader) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96127,Native(imageReader),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImageReader(QImageReader*,QFlags<Qt::ImageConversionFlag>)	
func QPixmapFromImageReaderWithImagereaderFlags(imageReader *QImageReader,flags Qt_ImageConversionFlag) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96128,Native(imageReader),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::fromImageReader(QImageReader*,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) FromImageReaderWithImagereaderFlags(imageReader *QImageReader,flags Qt_ImageConversionFlag) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96128,Native(imageReader),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*)	
func QPixmapGrabWidget(widget QWidgetInterface) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96129,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*)
func (q *QPixmap) GrabWidget(widget QWidgetInterface) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96129,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*,QRect const&)	
func QPixmapGrabWidgetWithWidgetRect(widget QWidgetInterface,rect *QRect) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96130,Native(widget),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*,QRect const&)
func (q *QPixmap) GrabWidgetWithWidgetRect(widget QWidgetInterface,rect *QRect) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96130,Native(widget),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*,int,int,int,int)	
func QPixmapGrabWidgetWithWidgetXYWidthHeight(widget QWidgetInterface,x int,y int,w int,h int) *QPixmap {
	var __rv uintptr
	DirectQtDrv(nil,96000,96131,Native(widget),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::grabWidget(QWidget*,int,int,int,int)
func (q *QPixmap) GrabWidgetWithWidgetXYWidthHeight(widget QWidgetInterface,x int,y int,w int,h int) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96131,Native(widget),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::hasAlpha()
func (q *QPixmap) HasAlpha() bool {
	var __rv bool
	q.Drv(96000,96132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::hasAlphaChannel()
func (q *QPixmap) HasAlphaChannel() bool {
	var __rv bool
	q.Drv(96000,96133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::height()
func (q *QPixmap) Height() int {
	var __rv int
	q.Drv(96000,96134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::isDetached()
func (q *QPixmap) IsDetached() bool {
	var __rv bool
	q.Drv(96000,96135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::isNull()
func (q *QPixmap) IsNull() bool {
	var __rv bool
	q.Drv(96000,96136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::isQBitmap()
func (q *QPixmap) IsQBitmap() bool {
	var __rv bool
	q.Drv(96000,96137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::load(QString const&)
func (q *QPixmap) Load(fileName string) bool {
	var __rv bool
	q.Drv(96000,96138,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::load(QString const&,char const*,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) LoadWithFilenameFormatFlags(fileName string,format string,flags Qt_ImageConversionFlag) bool {
	var __rv bool
	q.Drv(96000,96139,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::loadFromData(QByteArray const&)
func (q *QPixmap) LoadFromData(data []byte) bool {
	var __rv bool
	q.Drv(96000,96140,unsafe.Pointer(&data),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::loadFromData(QByteArray const&,char const*,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) LoadFromDataWithDataFormatFlags(data []byte,format string,flags Qt_ImageConversionFlag) bool {
	var __rv bool
	q.Drv(96000,96141,unsafe.Pointer(&data),unsafe.Pointer(&format),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::loadFromData(unsigned char const*,unsigned int,char const*,QFlags<Qt::ImageConversionFlag>)
func (q *QPixmap) LoadFromDataWithBufLenFormatFlags(buf *byte,len uint,format string,flags Qt_ImageConversionFlag) bool {
	var __rv bool
	q.Drv(96000,96142,unsafe.Pointer(&buf),unsafe.Pointer(&len),unsafe.Pointer(&format),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::mask()
func (q *QPixmap) Mask() *QBitmap {
	var __rv uintptr
	q.Drv(96000,96143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QPixmap::paintEngine()
func (q *QPixmap) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(96000,96144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QPixmap::rect()
func (q *QPixmap) Rect() *QRect {
	var __rv uintptr
	q.Drv(96000,96145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPixmap::save(QIODevice*)
func (q *QPixmap) Save(device QIODeviceInterface) bool {
	var __rv bool
	q.Drv(96000,96146,Native(device),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::save(QString const&)
func (q *QPixmap) SaveWithFilename(fileName string) bool {
	var __rv bool
	q.Drv(96000,96147,unsafe.Pointer(&fileName),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::save(QIODevice*,char const*,int)
func (q *QPixmap) SaveWithDeviceFormatQuality(device QIODeviceInterface,format string,quality int) bool {
	var __rv bool
	q.Drv(96000,96148,Native(device),unsafe.Pointer(&format),unsafe.Pointer(&quality),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::save(QString const&,char const*,int)
func (q *QPixmap) SaveWithFilenameFormatQuality(fileName string,format string,quality int) bool {
	var __rv bool
	q.Drv(96000,96149,unsafe.Pointer(&fileName),unsafe.Pointer(&format),unsafe.Pointer(&quality),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPixmap::scaled(QSize const&)
func (q *QPixmap) Scaled(s *QSize) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96150,Native(s),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaled(QSize const&,Qt::AspectRatioMode,Qt::TransformationMode)
func (q *QPixmap) ScaledWithSizeAspectmodeMode(s *QSize,aspectMode Qt_AspectRatioMode,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96151,Native(s),unsafe.Pointer(&aspectMode),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaled(int,int,Qt::AspectRatioMode,Qt::TransformationMode)
func (q *QPixmap) ScaledWithWidthHeightAspectmodeMode(w int,h int,aspectMode Qt_AspectRatioMode,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96152,unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&aspectMode),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaledToHeight(int)
func (q *QPixmap) ScaledToHeight(h int) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96153,unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaledToHeight(int,Qt::TransformationMode)
func (q *QPixmap) ScaledToHeightWithHeightMode(h int,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96154,unsafe.Pointer(&h),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaledToWidth(int)
func (q *QPixmap) ScaledToWidth(w int) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96155,unsafe.Pointer(&w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scaledToWidth(int,Qt::TransformationMode)
func (q *QPixmap) ScaledToWidthWithWidthMode(w int,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96156,unsafe.Pointer(&w),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::scroll(int,int,QRect const&,QRegion*)
func (q *QPixmap) ScrollWithDxDyRectExposed(dx int,dy int,rect *QRect,exposed *QRegion)  {
	q.Drv(96000,96157,unsafe.Pointer(&dx),unsafe.Pointer(&dy),Native(rect),Native(exposed),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::scroll(int,int,int,int,int,int,QRegion*)
func (q *QPixmap) ScrollWithDxDyXYWidthHeightExposed(dx int,dy int,x int,y int,width int,height int,exposed *QRegion)  {
	q.Drv(96000,96158,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&width),unsafe.Pointer(&height),Native(exposed),nil,nil,nil,nil,nil)
}	
//QPixmap::setMask(QBitmap const&)
func (q *QPixmap) SetMask(value *QBitmap)  {
	q.Drv(96000,96159,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPixmap::size()
func (q *QPixmap) Size() *QSize {
	var __rv uintptr
	q.Drv(96000,96160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QPixmap::toImage()
func (q *QPixmap) ToImage() *QImage {
	var __rv uintptr
	q.Drv(96000,96161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QPixmap::transformed(QMatrix const&)
func (q *QPixmap) Transformed(value *QMatrix) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96162,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::transformed(QTransform const&)
func (q *QPixmap) TransformedWithTransform(value *QTransform) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96163,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::transformed(QMatrix const&,Qt::TransformationMode)
func (q *QPixmap) TransformedWithMatrixMode(value2 *QMatrix,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96164,Native(value2),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::transformed(QTransform const&,Qt::TransformationMode)
func (q *QPixmap) TransformedWithTransformMode(value2 *QTransform,mode Qt_TransformationMode) *QPixmap {
	var __rv uintptr
	q.Drv(96000,96165,Native(value2),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QPixmap::trueMatrix(QMatrix const&,int,int)	
func QPixmapTrueMatrixWithMWidthHeight(m *QMatrix,w int,h int) *QMatrix {
	var __rv uintptr
	DirectQtDrv(nil,96000,96166,Native(m),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPixmap::trueMatrix(QMatrix const&,int,int)
func (q *QPixmap) TrueMatrixWithMWidthHeight(m *QMatrix,w int,h int) *QMatrix {
	var __rv uintptr
	q.Drv(96000,96166,Native(m),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPixmap::trueMatrix(QTransform const&,int,int)	
func QPixmapTrueMatrixWithTransformWidthHeight(m *QTransform,w int,h int) *QTransform {
	var __rv uintptr
	DirectQtDrv(nil,96000,96167,Native(m),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QPixmap::trueMatrix(QTransform const&,int,int)
func (q *QPixmap) TrueMatrixWithTransformWidthHeight(m *QTransform,w int,h int) *QTransform {
	var __rv uintptr
	q.Drv(96000,96167,Native(m),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QPixmap::width()
func (q *QPixmap) Width() int {
	var __rv int
	q.Drv(96000,96168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
