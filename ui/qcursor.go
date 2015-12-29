// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QCursor : QCursor
type QCursor struct {
	BaseDrv
}
//QCursor::QCursor()	
func NewCursor() *QCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),18000,18102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCursor{}
	_p.SetDriver(__rv,18000,true)
	return _p
} 
//QCursor::QCursor(QCursor const&)	
func NewCursorCopy(cursor *QCursor) *QCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),18000,18103,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCursor{}
	_p.SetDriver(__rv,18000,true)
	return _p
} 
//QCursor::QCursor(Qt::CursorShape)	
func NewCursorWithShape(shape Qt_CursorShape) *QCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),18000,18104,unsafe.Pointer(&shape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCursor{}
	_p.SetDriver(__rv,18000,true)
	return _p
} 
//QCursor::QCursor(QPixmap const&,int,int)	
func NewCursorWithPixmapHotxHoty(pixmap *QPixmap,hotX int,hotY int) *QCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),18000,18105,Native(pixmap),unsafe.Pointer(&hotX),unsafe.Pointer(&hotY),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCursor{}
	_p.SetDriver(__rv,18000,true)
	return _p
} 
//QCursor::QCursor(QBitmap const&,QBitmap const&,int,int)	
func NewCursorWithBitmapMaskHotxHoty(bitmap *QBitmap,mask *QBitmap,hotX int,hotY int) *QCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),18000,18106,Native(bitmap),Native(mask),unsafe.Pointer(&hotX),unsafe.Pointer(&hotY),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCursor{}
	_p.SetDriver(__rv,18000,true)
	return _p
} 
//QCursor::bitmap()
func (q *QCursor) Bitmap() *QBitmap {
	var __rv uintptr
	q.Drv(18000,18107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QCursor::hotSpot()
func (q *QCursor) HotSpot() *QPoint {
	var __rv uintptr
	q.Drv(18000,18108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QCursor::mask()
func (q *QCursor) Mask() *QBitmap {
	var __rv uintptr
	q.Drv(18000,18109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBitmap{}
	_rp.SetDriver(__rv,8000,true)
	return _rp
}	
//QCursor::pixmap()
func (q *QCursor) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(18000,18110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QCursor::pos()	
func QCursorPos() *QPoint {
	var __rv uintptr
	DirectQtDrv(nil,18000,18111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QCursor::pos()
func (q *QCursor) Pos() *QPoint {
	var __rv uintptr
	q.Drv(18000,18111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QCursor::setPos(QPoint const&)	
func QCursorSetPos(p *QPoint)  {
	DirectQtDrv(nil,18000,18112,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCursor::setPos(QPoint const&)
func (q *QCursor) SetPos(p *QPoint)  {
	q.Drv(18000,18112,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCursor::setPos(int,int)	
func QCursorSetPosWithXY(x int,y int)  {
	DirectQtDrv(nil,18000,18113,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCursor::setPos(int,int)
func (q *QCursor) SetPosWithXY(x int,y int)  {
	q.Drv(18000,18113,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCursor::setShape(Qt::CursorShape)
func (q *QCursor) SetShape(newShape Qt_CursorShape)  {
	q.Drv(18000,18114,unsafe.Pointer(&newShape),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCursor::shape()
func (q *QCursor) Shape() Qt_CursorShape {
	var __rv Qt_CursorShape
	q.Drv(18000,18115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
