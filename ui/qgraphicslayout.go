// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsLayout : QGraphicsLayout
type QGraphicsLayout struct {
	QGraphicsLayoutItem
}
//QGraphicsLayout::activate()
func (q *QGraphicsLayout) Activate()  {
	q.Drv(259000,259102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::count()
func (q *QGraphicsLayout) Count() int {
	var __rv int
	q.Drv(259000,259103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayout::getContentsMargins(double*,double*,double*,double*)
func (q *QGraphicsLayout) GetContentsMargins(left *float64,top *float64,right *float64,bottom *float64)  {
	q.Drv(259000,259104,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::instantInvalidatePropagation()	
func QGraphicsLayoutInstantInvalidatePropagation() bool {
	var __rv bool
	DirectQtDrv(nil,259000,259105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayout::instantInvalidatePropagation()
func (q *QGraphicsLayout) InstantInvalidatePropagation() bool {
	var __rv bool
	q.Drv(259000,259105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayout::invalidate()
func (q *QGraphicsLayout) Invalidate()  {
	q.Drv(259000,259106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::isActivated()
func (q *QGraphicsLayout) IsActivated() bool {
	var __rv bool
	q.Drv(259000,259107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsLayout::itemAt(int)
func (q *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem {
	var __rv uintptr
	q.Drv(259000,259108,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLayoutItem{}
	_rp.SetDriver(__rv,260000,true)
	return _rp
}	
//QGraphicsLayout::removeAt(int)
func (q *QGraphicsLayout) RemoveAt(index int)  {
	q.Drv(259000,259109,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::setContentsMargins(double,double,double,double)
func (q *QGraphicsLayout) SetContentsMargins(left float64,top float64,right float64,bottom float64)  {
	q.Drv(259000,259110,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::setInstantInvalidatePropagation(bool)	
func QGraphicsLayoutSetInstantInvalidatePropagation(enable bool)  {
	DirectQtDrv(nil,259000,259111,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::setInstantInvalidatePropagation(bool)
func (q *QGraphicsLayout) SetInstantInvalidatePropagation(enable bool)  {
	q.Drv(259000,259111,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::updateGeometry()
func (q *QGraphicsLayout) UpdateGeometry()  {
	q.Drv(259000,259112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsLayout::widgetEvent(QEvent*)
func (q *QGraphicsLayout) WidgetEvent(e *QEvent)  {
	q.Drv(259000,259113,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
