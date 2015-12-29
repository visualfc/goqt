// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPainterPixmapFragment : QPainter::PixmapFragment
type QPainterPixmapFragment struct {
	BaseDrv
}
//QPainter::PixmapFragment::PixmapFragment()	
func NewPainterPixmapFragment() *QPainterPixmapFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),87000,87102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPixmapFragment{}
	_p.SetDriver(__rv,87000,true)
	return _p
} 
//QPainter::PixmapFragment::create(QPointF const&,QRectF const&,double,double,double,double)	
func QPainterPixmapFragmentCreate(pos *QPointF,sourceRect *QRectF,scaleX float64,scaleY float64,rotation float64,opacity float64) *QPainterPixmapFragment {
	var __rv uintptr
	DirectQtDrv(nil,87000,87103,Native(pos),Native(sourceRect),unsafe.Pointer(&scaleX),unsafe.Pointer(&scaleY),unsafe.Pointer(&rotation),unsafe.Pointer(&opacity),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPixmapFragment{}
	_rp.SetDriver(__rv,87000,true)
	return _rp
}	
//QPainter::PixmapFragment::create(QPointF const&,QRectF const&,double,double,double,double)
func (q *QPainterPixmapFragment) Create(pos *QPointF,sourceRect *QRectF,scaleX float64,scaleY float64,rotation float64,opacity float64) *QPainterPixmapFragment {
	var __rv uintptr
	q.Drv(87000,87103,Native(pos),Native(sourceRect),unsafe.Pointer(&scaleX),unsafe.Pointer(&scaleY),unsafe.Pointer(&rotation),unsafe.Pointer(&opacity),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPixmapFragment{}
	_rp.SetDriver(__rv,87000,true)
	return _rp
}	
