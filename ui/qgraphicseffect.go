// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsEffect_PixmapPadMode - QGraphicsEffect::PixmapPadMode
type QGraphicsEffect_PixmapPadMode uint32
const (
	QGraphicsEffect_NoPad QGraphicsEffect_PixmapPadMode = 0
	QGraphicsEffect_PadToTransparentBorder QGraphicsEffect_PixmapPadMode = 1
	QGraphicsEffect_PadToEffectiveBoundingRect QGraphicsEffect_PixmapPadMode = 2
)
//enum QGraphicsEffect_ChangeFlag - QGraphicsEffect::ChangeFlag
type QGraphicsEffect_ChangeFlag uint32
const (
	QGraphicsEffect_SourceAttached QGraphicsEffect_ChangeFlag = 0x1
	QGraphicsEffect_SourceDetached QGraphicsEffect_ChangeFlag = 0x2
	QGraphicsEffect_SourceBoundingRectChanged QGraphicsEffect_ChangeFlag = 0x4
	QGraphicsEffect_SourceInvalidated QGraphicsEffect_ChangeFlag = 0x8
)
//struct QGraphicsEffect : QGraphicsEffect
type QGraphicsEffect struct {
	QObject
}
func (q *QGraphicsEffect) OnEnabledChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(253000,253102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsEffect::boundingRect()
func (q *QGraphicsEffect) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(253000,253103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsEffect::boundingRectFor(QRectF const&)
func (q *QGraphicsEffect) BoundingRectFor(sourceRect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(253000,253104,Native(sourceRect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsEffect::draw(QPainter*)
func (q *QGraphicsEffect) Draw(painter *QPainter)  {
	q.Drv(253000,253105,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEffect::drawSource(QPainter*)
func (q *QGraphicsEffect) DrawSource(painter *QPainter)  {
	q.Drv(253000,253106,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEffect::isEnabled()
func (q *QGraphicsEffect) IsEnabled() bool {
	var __rv bool
	q.Drv(253000,253107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEffect::setEnabled(bool)
func (q *QGraphicsEffect) SetEnabled(enable bool)  {
	q.Drv(253000,253108,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEffect::sourceBoundingRect(Qt::CoordinateSystem)
func (q *QGraphicsEffect) SourceBoundingRect(system Qt_CoordinateSystem) *QRectF {
	var __rv uintptr
	q.Drv(253000,253109,unsafe.Pointer(&system),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsEffect::sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>)
func (q *QGraphicsEffect) SourceChanged(flags QGraphicsEffect_ChangeFlag)  {
	q.Drv(253000,253110,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEffect::sourceIsPixmap()
func (q *QGraphicsEffect) SourceIsPixmap() bool {
	var __rv bool
	q.Drv(253000,253111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsEffect::sourcePixmap(Qt::CoordinateSystem,QPoint*,QGraphicsEffect::PixmapPadMode)
func (q *QGraphicsEffect) SourcePixmap(system Qt_CoordinateSystem,offset *QPoint,mode QGraphicsEffect_PixmapPadMode) *QPixmap {
	var __rv uintptr
	q.Drv(253000,253112,unsafe.Pointer(&system),Native(offset),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QGraphicsEffect::update()
func (q *QGraphicsEffect) Update()  {
	q.Drv(253000,253113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsEffect::updateBoundingRect()
func (q *QGraphicsEffect) UpdateBoundingRect()  {
	q.Drv(253000,253114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
