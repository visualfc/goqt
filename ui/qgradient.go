// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGradient_CoordinateMode - QGradient::CoordinateMode
type QGradient_CoordinateMode uint32
const (
	QGradient_LogicalMode QGradient_CoordinateMode = 0
	QGradient_StretchToDeviceMode QGradient_CoordinateMode = 1
	QGradient_ObjectBoundingMode QGradient_CoordinateMode = 2
)
//enum QGradient_Spread - QGradient::Spread
type QGradient_Spread uint32
const (
	QGradient_PadSpread QGradient_Spread = 0
	QGradient_ReflectSpread QGradient_Spread = 1
	QGradient_RepeatSpread QGradient_Spread = 2
)
//enum QGradient_Type - QGradient::Type
type QGradient_Type uint32
const (
	QGradient_LinearGradient QGradient_Type = 0
	QGradient_RadialGradient QGradient_Type = 1
	QGradient_ConicalGradient QGradient_Type = 2
	QGradient_NoGradient QGradient_Type = 3
)
//enum QGradient_InterpolationMode - QGradient::InterpolationMode
type QGradient_InterpolationMode uint32
const (
	QGradient_ColorInterpolation QGradient_InterpolationMode = 0
	QGradient_ComponentInterpolation QGradient_InterpolationMode = 1
)
//struct QGradient : QGradient
type QGradient struct {
	BaseDrv
}
//QGradient::QGradient()	
func NewGradient() *QGradient {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),46000,46102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGradient{}
	_p.SetDriver(__rv,46000,true)
	return _p
} 
//QGradient::coordinateMode()
func (q *QGradient) CoordinateMode() QGradient_CoordinateMode {
	var __rv QGradient_CoordinateMode
	q.Drv(46000,46103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGradient::interpolationMode()
func (q *QGradient) InterpolationMode() QGradient_InterpolationMode {
	var __rv QGradient_InterpolationMode
	q.Drv(46000,46104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGradient::setColorAt(double,QColor const&)
func (q *QGradient) SetColorAt(pos float64,color *QColor)  {
	q.Drv(46000,46105,unsafe.Pointer(&pos),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGradient::setCoordinateMode(QGradient::CoordinateMode)
func (q *QGradient) SetCoordinateMode(mode QGradient_CoordinateMode)  {
	q.Drv(46000,46106,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGradient::setInterpolationMode(QGradient::InterpolationMode)
func (q *QGradient) SetInterpolationMode(mode QGradient_InterpolationMode)  {
	q.Drv(46000,46107,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGradient::setSpread(QGradient::Spread)
func (q *QGradient) SetSpread(spread QGradient_Spread)  {
	q.Drv(46000,46108,unsafe.Pointer(&spread),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGradient::spread()
func (q *QGradient) Spread() QGradient_Spread {
	var __rv QGradient_Spread
	q.Drv(46000,46109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGradient::type()
func (q *QGradient) Type() QGradient_Type {
	var __rv QGradient_Type
	q.Drv(46000,46110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
