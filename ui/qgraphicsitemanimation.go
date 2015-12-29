// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGraphicsItemAnimation : QGraphicsItemAnimation
type QGraphicsItemAnimation struct {
	QObject
}
//QGraphicsItemAnimation::QGraphicsItemAnimation()	
func NewGraphicsItemAnimation() *QGraphicsItemAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),257000,257102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsItemAnimation{}
	_p.SetDriver(__rv,257000,false)
	return _p
} 
//QGraphicsItemAnimation::QGraphicsItemAnimation(QObject*)	
func NewGraphicsItemAnimationWithParent(parent QObjectInterface) *QGraphicsItemAnimation {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),257000,257103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsItemAnimation{}
	_p.SetDriver(__rv,257000,false)
	return _p
} 
//QGraphicsItemAnimation::afterAnimationStep(double)
func (q *QGraphicsItemAnimation) AfterAnimationStep(step float64)  {
	q.Drv(257000,257104,unsafe.Pointer(&step),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::beforeAnimationStep(double)
func (q *QGraphicsItemAnimation) BeforeAnimationStep(step float64)  {
	q.Drv(257000,257105,unsafe.Pointer(&step),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::clear()
func (q *QGraphicsItemAnimation) Clear()  {
	q.Drv(257000,257106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::horizontalScaleAt(double)
func (q *QGraphicsItemAnimation) HorizontalScaleAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257107,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::horizontalShearAt(double)
func (q *QGraphicsItemAnimation) HorizontalShearAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257108,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::item()
func (q *QGraphicsItemAnimation) Item() *QGraphicsItem {
	var __rv uintptr
	q.Drv(257000,257109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItemAnimation::matrixAt(double)
func (q *QGraphicsItemAnimation) MatrixAt(step float64) *QMatrix {
	var __rv uintptr
	q.Drv(257000,257110,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QGraphicsItemAnimation::posAt(double)
func (q *QGraphicsItemAnimation) PosAt(step float64) *QPointF {
	var __rv uintptr
	q.Drv(257000,257111,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItemAnimation::reset()
func (q *QGraphicsItemAnimation) Reset()  {
	q.Drv(257000,257112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::rotationAt(double)
func (q *QGraphicsItemAnimation) RotationAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257113,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::setItem(QGraphicsItem*)
func (q *QGraphicsItemAnimation) SetItem(item *QGraphicsItem)  {
	q.Drv(257000,257114,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setPosAt(double,QPointF const&)
func (q *QGraphicsItemAnimation) SetPosAt(step float64,pos *QPointF)  {
	q.Drv(257000,257115,unsafe.Pointer(&step),Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setRotationAt(double,double)
func (q *QGraphicsItemAnimation) SetRotationAt(step float64,angle float64)  {
	q.Drv(257000,257116,unsafe.Pointer(&step),unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setScaleAt(double,double,double)
func (q *QGraphicsItemAnimation) SetScaleAt(step float64,sx float64,sy float64)  {
	q.Drv(257000,257117,unsafe.Pointer(&step),unsafe.Pointer(&sx),unsafe.Pointer(&sy),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setShearAt(double,double,double)
func (q *QGraphicsItemAnimation) SetShearAt(step float64,sh float64,sv float64)  {
	q.Drv(257000,257118,unsafe.Pointer(&step),unsafe.Pointer(&sh),unsafe.Pointer(&sv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setStep(double)
func (q *QGraphicsItemAnimation) SetStep(x float64)  {
	q.Drv(257000,257119,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setTimeLine(QTimeLine*)
func (q *QGraphicsItemAnimation) SetTimeLine(timeLine *QTimeLine)  {
	q.Drv(257000,257120,Native(timeLine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::setTranslationAt(double,double,double)
func (q *QGraphicsItemAnimation) SetTranslationAt(step float64,dx float64,dy float64)  {
	q.Drv(257000,257121,unsafe.Pointer(&step),unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItemAnimation::timeLine()
func (q *QGraphicsItemAnimation) TimeLine() *QTimeLine {
	var __rv uintptr
	q.Drv(257000,257122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTimeLine{}
	_rp.SetDriver(__rv,379000,false)
	return _rp
}	
//QGraphicsItemAnimation::verticalScaleAt(double)
func (q *QGraphicsItemAnimation) VerticalScaleAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257123,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::verticalShearAt(double)
func (q *QGraphicsItemAnimation) VerticalShearAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257124,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::xTranslationAt(double)
func (q *QGraphicsItemAnimation) XTranslationAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257125,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItemAnimation::yTranslationAt(double)
func (q *QGraphicsItemAnimation) YTranslationAt(step float64) float64 {
	var __rv float64
	q.Drv(257000,257126,unsafe.Pointer(&step),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
