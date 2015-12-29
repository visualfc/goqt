// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPaintEvent : QPaintEvent
type QPaintEvent struct {
	QEvent
}
//QPaintEvent::QPaintEvent(QRect const&)	
func NewPaintEvent(paintRect *QRect) *QPaintEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),85000,85102,Native(paintRect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPaintEvent{}
	_p.SetDriver(__rv,85000,true)
	return _p
} 
//QPaintEvent::QPaintEvent(QRegion const&)	
func NewPaintEventWithPaintregion(paintRegion *QRegion) *QPaintEvent {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),85000,85103,Native(paintRegion),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPaintEvent{}
	_p.SetDriver(__rv,85000,true)
	return _p
} 
//QPaintEvent::rect()
func (q *QPaintEvent) Rect() *QRect {
	var __rv uintptr
	q.Drv(85000,85104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPaintEvent::region()
func (q *QPaintEvent) Region() *QRegion {
	var __rv uintptr
	q.Drv(85000,85105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
