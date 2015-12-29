// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QBrush : QBrush
type QBrush struct {
	BaseDrv
}
//QBrush::QBrush()	
func NewBrush() *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QBrush const&)	
func NewBrushCopy(brush *QBrush) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9103,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QGradient const&)	
func NewBrushWithGradient(gradient *QGradient) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9104,Native(gradient),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QImage const&)	
func NewBrushWithImage(image *QImage) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9105,Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QPixmap const&)	
func NewBrushWithPixmap(pixmap *QPixmap) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9106,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(Qt::BrushStyle)	
func NewBrushWithBrushstyle(bs Qt_BrushStyle) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9107,unsafe.Pointer(&bs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QColor const&,QPixmap const&)	
func NewBrushWithColorPixmap(color *QColor,pixmap *QPixmap) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9108,Native(color),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(QColor const&,Qt::BrushStyle)	
func NewBrushWithColorBrushstyle(color *QColor,bs Qt_BrushStyle) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9109,Native(color),unsafe.Pointer(&bs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(Qt::GlobalColor,QPixmap const&)	
func NewBrushWithGlobalcolorPixmap(color Qt_GlobalColor,pixmap *QPixmap) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9110,unsafe.Pointer(&color),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::QBrush(Qt::GlobalColor,Qt::BrushStyle)	
func NewBrushWithGlobalcolorBrushstyle(color Qt_GlobalColor,bs Qt_BrushStyle) *QBrush {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),9000,9111,unsafe.Pointer(&color),unsafe.Pointer(&bs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QBrush{}
	_p.SetDriver(__rv,9000,true)
	return _p
} 
//QBrush::color()
func (q *QBrush) Color() *QColor {
	var __rv uintptr
	q.Drv(9000,9112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QBrush::gradient()
func (q *QBrush) Gradient() *QGradient {
	var __rv uintptr
	q.Drv(9000,9113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGradient{}
	_rp.SetDriver(__rv,46000,true)
	return _rp
}	
//QBrush::isDetached()
func (q *QBrush) IsDetached() bool {
	var __rv bool
	q.Drv(9000,9114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBrush::isOpaque()
func (q *QBrush) IsOpaque() bool {
	var __rv bool
	q.Drv(9000,9115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBrush::matrix()
func (q *QBrush) Matrix() *QMatrix {
	var __rv uintptr
	q.Drv(9000,9116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QBrush::setColor(QColor const&)
func (q *QBrush) SetColor(color *QColor)  {
	q.Drv(9000,9117,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setColor(Qt::GlobalColor)
func (q *QBrush) SetColorWithGlobalcolor(color Qt_GlobalColor)  {
	q.Drv(9000,9118,unsafe.Pointer(&color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setMatrix(QMatrix const&)
func (q *QBrush) SetMatrix(mat *QMatrix)  {
	q.Drv(9000,9119,Native(mat),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setStyle(Qt::BrushStyle)
func (q *QBrush) SetStyle(value Qt_BrushStyle)  {
	q.Drv(9000,9120,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setTexture(QPixmap const&)
func (q *QBrush) SetTexture(pixmap *QPixmap)  {
	q.Drv(9000,9121,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setTextureImage(QImage const&)
func (q *QBrush) SetTextureImage(image *QImage)  {
	q.Drv(9000,9122,Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::setTransform(QTransform const&)
func (q *QBrush) SetTransform(value *QTransform)  {
	q.Drv(9000,9123,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QBrush::style()
func (q *QBrush) Style() Qt_BrushStyle {
	var __rv Qt_BrushStyle
	q.Drv(9000,9124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QBrush::texture()
func (q *QBrush) Texture() *QPixmap {
	var __rv uintptr
	q.Drv(9000,9125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QBrush::textureImage()
func (q *QBrush) TextureImage() *QImage {
	var __rv uintptr
	q.Drv(9000,9126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QImage{}
	_rp.SetDriver(__rv,53000,true)
	return _rp
}	
//QBrush::transform()
func (q *QBrush) Transform() *QTransform {
	var __rv uintptr
	q.Drv(9000,9127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
