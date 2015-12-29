// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSplitterHandle : QSplitterHandle
type QSplitterHandle struct {
	QWidget
}
//QSplitterHandle::QSplitterHandle(Qt::Orientation,QSplitter*)	
func NewSplitterHandle(o Qt_Orientation,parent *QSplitter) *QSplitterHandle {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),349000,349102,unsafe.Pointer(&o),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplitterHandle{}
	_p.SetDriver(__rv,349000,false)
	return _p
} 
//QSplitterHandle::closestLegalPosition(int)
func (q *QSplitterHandle) ClosestLegalPosition(p int) int {
	var __rv int
	q.Drv(349000,349103,unsafe.Pointer(&p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitterHandle::event(QEvent*)
func (q *QSplitterHandle) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(349000,349104,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitterHandle::mouseMoveEvent(QMouseEvent*)
func (q *QSplitterHandle) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(349000,349105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::mousePressEvent(QMouseEvent*)
func (q *QSplitterHandle) MousePressEvent(value *QMouseEvent)  {
	q.Drv(349000,349106,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::mouseReleaseEvent(QMouseEvent*)
func (q *QSplitterHandle) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(349000,349107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::moveSplitter(int)
func (q *QSplitterHandle) MoveSplitter(p int)  {
	q.Drv(349000,349108,unsafe.Pointer(&p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::opaqueResize()
func (q *QSplitterHandle) OpaqueResize() bool {
	var __rv bool
	q.Drv(349000,349109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitterHandle::orientation()
func (q *QSplitterHandle) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(349000,349110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitterHandle::paintEvent(QPaintEvent*)
func (q *QSplitterHandle) PaintEvent(value *QPaintEvent)  {
	q.Drv(349000,349111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::resizeEvent(QResizeEvent*)
func (q *QSplitterHandle) ResizeEvent(value *QResizeEvent)  {
	q.Drv(349000,349112,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::setOrientation(Qt::Orientation)
func (q *QSplitterHandle) SetOrientation(o Qt_Orientation)  {
	q.Drv(349000,349113,unsafe.Pointer(&o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitterHandle::sizeHint()
func (q *QSplitterHandle) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(349000,349114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSplitterHandle::splitter()
func (q *QSplitterHandle) Splitter() *QSplitter {
	var __rv uintptr
	q.Drv(349000,349115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSplitter{}
	_rp.SetDriver(__rv,348000,false)
	return _rp
}	
