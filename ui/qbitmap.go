// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QBitmap : QBitmap
type QBitmap struct {
	QPixmap
}
//QBitmap::QBitmap()	
func NewBitmap() *QBitmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),8000,8102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBitmap{}
	_p.SetDriver(__rv,8000,true)
	return _p
} 
//QBitmap::QBitmap(QPixmap const&)	
func NewBitmapWithPixmap(value *QPixmap) *QBitmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),8000,8103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBitmap{}
	_p.SetDriver(__rv,8000,true)
	return _p
} 
//QBitmap::QBitmap(QSize const&)	
func NewBitmapWithSize(value *QSize) *QBitmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),8000,8104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBitmap{}
	_p.SetDriver(__rv,8000,true)
	return _p
} 
//QBitmap::QBitmap(QString const&,char const*)	
func NewBitmapWithFilenameFormat(fileName string,format string) *QBitmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),8000,8105,unsafe.Pointer(&fileName),unsafe.Pointer(&format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBitmap{}
	_p.SetDriver(__rv,8000,true)
	return _p
} 
//QBitmap::QBitmap(int,int)	
func NewBitmapWithWidthHeight(w int,h int) *QBitmap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),8000,8106,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBitmap{}
	_p.SetDriver(__rv,8000,true)
	return _p
} 
//QBitmap::clear()
func (q *QBitmap) Clear()  {
	q.Drv(8000,8107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBitmap::fromData(QSize const&,unsigned char const*,QImage::Format)	
func QBitmapFromData(size *QSize,bits *byte,monoFormat QImage_Format) *QBitmap {
	var __rv uintptr
	DirectQtDrv(nil,8000,8108,Native(size),unsafe.Pointer(&bits),unsafe.Pointer(&monoFormat),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::fromData(QSize const&,unsigned char const*,QImage::Format)
func (q *QBitmap) FromData(size *QSize,bits *byte,monoFormat QImage_Format) *QBitmap {
	var __rv uintptr
	q.Drv(8000,8108,Native(size),unsafe.Pointer(&bits),unsafe.Pointer(&monoFormat),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::fromImage(QImage const&)	
func QBitmapFromImage(image *QImage) *QBitmap {
	var __rv uintptr
	DirectQtDrv(nil,8000,8109,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::fromImage(QImage const&)
func (q *QBitmap) FromImage(image *QImage) *QBitmap {
	var __rv uintptr
	q.Drv(8000,8109,Native(image),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::fromImage(QImage const&,QFlags<Qt::ImageConversionFlag>)	
func QBitmapFromImageWithImageFlags(image *QImage,flags Qt_ImageConversionFlag) *QBitmap {
	var __rv uintptr
	DirectQtDrv(nil,8000,8110,Native(image),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::fromImage(QImage const&,QFlags<Qt::ImageConversionFlag>)
func (q *QBitmap) FromImageWithImageFlags(image *QImage,flags Qt_ImageConversionFlag) *QBitmap {
	var __rv uintptr
	q.Drv(8000,8110,Native(image),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::transformed(QMatrix const&)
func (q *QBitmap) Transformed(value *QMatrix) *QBitmap {
	var __rv uintptr
	q.Drv(8000,8111,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QBitmap::transformed(QTransform const&)
func (q *QBitmap) TransformedWithTransform(matrix *QTransform) *QBitmap {
	var __rv uintptr
	q.Drv(8000,8112,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
