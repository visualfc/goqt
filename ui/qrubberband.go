// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QRubberBand_Shape - QRubberBand::Shape
type QRubberBand_Shape uint32
const (
	QRubberBand_Line QRubberBand_Shape = 0
	QRubberBand_Rectangle QRubberBand_Shape = 1
)
//struct QRubberBand : QRubberBand
type QRubberBand struct {
	QWidget
}
//QRubberBand::QRubberBand(QRubberBand::Shape,QWidget*)	
func NewRubberBand(value2 QRubberBand_Shape,value3 QWidgetInterface) *QRubberBand {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),334000,334102,unsafe.Pointer(&value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRubberBand{}
	_p.SetDriver(__rv,334000,false)
	return _p
} 
//QRubberBand::changeEvent(QEvent*)
func (q *QRubberBand) ChangeEvent(value *QEvent)  {
	q.Drv(334000,334103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::event(QEvent*)
func (q *QRubberBand) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(334000,334104,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRubberBand::move(QPoint const&)
func (q *QRubberBand) Move(p *QPoint)  {
	q.Drv(334000,334105,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::move(int,int)
func (q *QRubberBand) MoveWithXY(x int,y int)  {
	q.Drv(334000,334106,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::moveEvent(QMoveEvent*)
func (q *QRubberBand) MoveEvent(value *QMoveEvent)  {
	q.Drv(334000,334107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::paintEvent(QPaintEvent*)
func (q *QRubberBand) PaintEvent(value *QPaintEvent)  {
	q.Drv(334000,334108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::resize(QSize const&)
func (q *QRubberBand) Resize(s *QSize)  {
	q.Drv(334000,334109,Native(s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::resize(int,int)
func (q *QRubberBand) ResizeWithWidthHeight(w int,h int)  {
	q.Drv(334000,334110,unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::resizeEvent(QResizeEvent*)
func (q *QRubberBand) ResizeEvent(value *QResizeEvent)  {
	q.Drv(334000,334111,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::setGeometry(QRect const&)
func (q *QRubberBand) SetGeometry(r *QRect)  {
	q.Drv(334000,334112,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::setGeometry(int,int,int,int)
func (q *QRubberBand) SetGeometryWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(334000,334113,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRubberBand::shape()
func (q *QRubberBand) Shape() QRubberBand_Shape {
	var __rv QRubberBand_Shape
	q.Drv(334000,334114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRubberBand::showEvent(QShowEvent*)
func (q *QRubberBand) ShowEvent(value *QShowEvent)  {
	q.Drv(334000,334115,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
