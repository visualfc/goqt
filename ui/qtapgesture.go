// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTapGesture : QTapGesture
type QTapGesture struct {
	QGesture
}
//QTapGesture::QTapGesture()	
func NewTapGesture() *QTapGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),368000,368102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTapGesture{}
	_p.SetDriver(__rv,368000,false)
	return _p
} 
//QTapGesture::QTapGesture(QObject*)	
func NewTapGestureWithParent(parent QObjectInterface) *QTapGesture {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),368000,368103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTapGesture{}
	_p.SetDriver(__rv,368000,false)
	return _p
} 
//QTapGesture::position()
func (q *QTapGesture) Position() *QPointF {
	var __rv uintptr
	q.Drv(368000,368104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QTapGesture::setPosition(QPointF const&)
func (q *QTapGesture) SetPosition(pos *QPointF)  {
	q.Drv(368000,368105,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
