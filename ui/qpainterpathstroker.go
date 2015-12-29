// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QPainterPathStroker : QPainterPathStroker
type QPainterPathStroker struct {
	BaseDrv
}
//QPainterPathStroker::QPainterPathStroker()	
func NewPainterPathStroker() *QPainterPathStroker {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),90000,90102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPathStroker{}
	_p.SetDriver(__rv,90000,true)
	return _p
} 
//QPainterPathStroker::capStyle()
func (q *QPainterPathStroker) CapStyle() Qt_PenCapStyle {
	var __rv Qt_PenCapStyle
	q.Drv(90000,90103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::createStroke(QPainterPath const&)
func (q *QPainterPathStroker) CreateStroke(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(90000,90104,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPathStroker::curveThreshold()
func (q *QPainterPathStroker) CurveThreshold() float64 {
	var __rv float64
	q.Drv(90000,90105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::dashOffset()
func (q *QPainterPathStroker) DashOffset() float64 {
	var __rv float64
	q.Drv(90000,90106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::dashPattern()
func (q *QPainterPathStroker) DashPattern() []float64 {
	var __rv []float64
	q.Drv(90000,90107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::joinStyle()
func (q *QPainterPathStroker) JoinStyle() Qt_PenJoinStyle {
	var __rv Qt_PenJoinStyle
	q.Drv(90000,90108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::miterLimit()
func (q *QPainterPathStroker) MiterLimit() float64 {
	var __rv float64
	q.Drv(90000,90109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPathStroker::setCapStyle(Qt::PenCapStyle)
func (q *QPainterPathStroker) SetCapStyle(style Qt_PenCapStyle)  {
	q.Drv(90000,90110,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setCurveThreshold(double)
func (q *QPainterPathStroker) SetCurveThreshold(threshold float64)  {
	q.Drv(90000,90111,unsafe.Pointer(&threshold),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setDashOffset(double)
func (q *QPainterPathStroker) SetDashOffset(offset float64)  {
	q.Drv(90000,90112,unsafe.Pointer(&offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setDashPattern(QVector<qreal> const&)
func (q *QPainterPathStroker) SetDashPattern(dashPattern []float64)  {
	q.Drv(90000,90113,unsafe.Pointer(&dashPattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setDashPattern(Qt::PenStyle)
func (q *QPainterPathStroker) SetDashPatternWithPenstyle(value Qt_PenStyle)  {
	q.Drv(90000,90114,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setJoinStyle(Qt::PenJoinStyle)
func (q *QPainterPathStroker) SetJoinStyle(style Qt_PenJoinStyle)  {
	q.Drv(90000,90115,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setMiterLimit(double)
func (q *QPainterPathStroker) SetMiterLimit(length float64)  {
	q.Drv(90000,90116,unsafe.Pointer(&length),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::setWidth(double)
func (q *QPainterPathStroker) SetWidth(width float64)  {
	q.Drv(90000,90117,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPathStroker::width()
func (q *QPainterPathStroker) Width() float64 {
	var __rv float64
	q.Drv(90000,90118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
