// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPainterPath_ElementType - QPainterPath::ElementType
type QPainterPath_ElementType uint32
const (
	QPainterPath_MoveToElement QPainterPath_ElementType = 0
	QPainterPath_LineToElement QPainterPath_ElementType = 1
	QPainterPath_CurveToElement QPainterPath_ElementType = 2
	QPainterPath_CurveToDataElement QPainterPath_ElementType = 3
)
//struct QPainterPath : QPainterPath
type QPainterPath struct {
	BaseDrv
}
//QPainterPath::QPainterPath()	
func NewPainterPath() *QPainterPath {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),88000,88102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPath{}
	_p.SetDriver(__rv,88000,true)
	return _p
} 
//QPainterPath::QPainterPath(QPainterPath const&)	
func NewPainterPathCopy(other *QPainterPath) *QPainterPath {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),88000,88103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPath{}
	_p.SetDriver(__rv,88000,true)
	return _p
} 
//QPainterPath::QPainterPath(QPointF const&)	
func NewPainterPathWithStartpoint(startPoint *QPointF) *QPainterPath {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),88000,88104,Native(startPoint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainterPath{}
	_p.SetDriver(__rv,88000,true)
	return _p
} 
//QPainterPath::addEllipse(QRectF const&)
func (q *QPainterPath) AddEllipse(rect *QRectF)  {
	q.Drv(88000,88105,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addEllipse(QPointF const&,double,double)
func (q *QPainterPath) AddEllipseFWithCenterRxRy(center *QPointF,rx float64,ry float64)  {
	q.Drv(88000,88106,Native(center),unsafe.Pointer(&rx),unsafe.Pointer(&ry),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addEllipse(double,double,double,double)
func (q *QPainterPath) AddEllipseFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(88000,88107,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addPath(QPainterPath const&)
func (q *QPainterPath) AddPath(path *QPainterPath)  {
	q.Drv(88000,88108,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addPolygon(QPolygonF const&)
func (q *QPainterPath) AddPolygon(polygon *QPolygonF)  {
	q.Drv(88000,88109,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRect(QRectF const&)
func (q *QPainterPath) AddRect(rect *QRectF)  {
	q.Drv(88000,88110,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRect(double,double,double,double)
func (q *QPainterPath) AddRectFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(88000,88111,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRegion(QRegion const&)
func (q *QPainterPath) AddRegion(region *QRegion)  {
	q.Drv(88000,88112,Native(region),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundRect(QRectF const&,int)
func (q *QPainterPath) AddRoundRectFWithRectRoundness(rect *QRectF,roundness int)  {
	q.Drv(88000,88113,Native(rect),unsafe.Pointer(&roundness),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundRect(QRectF const&,int,int)
func (q *QPainterPath) AddRoundRectFWithRectXrndYrnd(rect *QRectF,xRnd int,yRnd int)  {
	q.Drv(88000,88114,Native(rect),unsafe.Pointer(&xRnd),unsafe.Pointer(&yRnd),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundRect(double,double,double,double,int)
func (q *QPainterPath) AddRoundRectFWithXYWidthHeightRoundness(x float64,y float64,w float64,h float64,roundness int)  {
	q.Drv(88000,88115,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&roundness),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundRect(double,double,double,double,int,int)
func (q *QPainterPath) AddRoundRectFWithXYWidthHeightXrndYrnd(x float64,y float64,w float64,h float64,xRnd int,yRnd int)  {
	q.Drv(88000,88116,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&xRnd),unsafe.Pointer(&yRnd),nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundedRect(QRectF const&,double,double,Qt::SizeMode)
func (q *QPainterPath) AddRoundedRectFWithRectXradiusYradiusMode(rect *QRectF,xRadius float64,yRadius float64,mode Qt_SizeMode)  {
	q.Drv(88000,88117,Native(rect),unsafe.Pointer(&xRadius),unsafe.Pointer(&yRadius),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addRoundedRect(double,double,double,double,double,double,Qt::SizeMode)
func (q *QPainterPath) AddRoundedRectFWithXYWidthHeightXradiusYradiusMode(x float64,y float64,w float64,h float64,xRadius float64,yRadius float64,mode Qt_SizeMode)  {
	q.Drv(88000,88118,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&xRadius),unsafe.Pointer(&yRadius),unsafe.Pointer(&mode),nil,nil,nil,nil,nil)
}	
//QPainterPath::addText(QPointF const&,QFont const&,QString const&)
func (q *QPainterPath) AddTextFWithPointFText(point *QPointF,f *QFont,text string)  {
	q.Drv(88000,88119,Native(point),Native(f),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::addText(double,double,QFont const&,QString const&)
func (q *QPainterPath) AddTextFWithXYFText(x float64,y float64,f *QFont,text string)  {
	q.Drv(88000,88120,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(f),unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::angleAtPercent(double)
func (q *QPainterPath) AngleAtPercent(t float64) float64 {
	var __rv float64
	q.Drv(88000,88121,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::arcMoveTo(QRectF const&,double)
func (q *QPainterPath) ArcMoveToFWithRectAngle(rect *QRectF,angle float64)  {
	q.Drv(88000,88122,Native(rect),unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::arcMoveTo(double,double,double,double,double)
func (q *QPainterPath) ArcMoveToFWithXYWidthHeightAngle(x float64,y float64,w float64,h float64,angle float64)  {
	q.Drv(88000,88123,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::arcTo(QRectF const&,double,double)
func (q *QPainterPath) ArcToFWithRectStartangleArclength(rect *QRectF,startAngle float64,arcLength float64)  {
	q.Drv(88000,88124,Native(rect),unsafe.Pointer(&startAngle),unsafe.Pointer(&arcLength),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::arcTo(double,double,double,double,double,double)
func (q *QPainterPath) ArcToFWithXYWidthHeightStartangleArclength(x float64,y float64,w float64,h float64,startAngle float64,arcLength float64)  {
	q.Drv(88000,88125,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&startAngle),unsafe.Pointer(&arcLength),nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::boundingRect()
func (q *QPainterPath) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(88000,88126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPainterPath::closeSubpath()
func (q *QPainterPath) CloseSubpath()  {
	q.Drv(88000,88127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::connectPath(QPainterPath const&)
func (q *QPainterPath) ConnectPath(path *QPainterPath)  {
	q.Drv(88000,88128,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::contains(QPainterPath const&)
func (q *QPainterPath) Contains(p *QPainterPath) bool {
	var __rv bool
	q.Drv(88000,88129,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::contains(QPointF const&)
func (q *QPainterPath) ContainsFWithPointf(pt *QPointF) bool {
	var __rv bool
	q.Drv(88000,88130,Native(pt),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::contains(QRectF const&)
func (q *QPainterPath) ContainsFWithRect(rect *QRectF) bool {
	var __rv bool
	q.Drv(88000,88131,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::controlPointRect()
func (q *QPainterPath) ControlPointRect() *QRectF {
	var __rv uintptr
	q.Drv(88000,88132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPainterPath::cubicTo(QPointF const&,QPointF const&,QPointF const&)
func (q *QPainterPath) CubicToFWithCtrlpt1Ctrlpt2Endpt(ctrlPt1 *QPointF,ctrlPt2 *QPointF,endPt *QPointF)  {
	q.Drv(88000,88133,Native(ctrlPt1),Native(ctrlPt2),Native(endPt),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::cubicTo(double,double,double,double,double,double)
func (q *QPainterPath) CubicToFWithCtrlpt1xCtrlpt1yCtrlpt2xCtrlpt2yEndptxEndpty(ctrlPt1x float64,ctrlPt1y float64,ctrlPt2x float64,ctrlPt2y float64,endPtx float64,endPty float64)  {
	q.Drv(88000,88134,unsafe.Pointer(&ctrlPt1x),unsafe.Pointer(&ctrlPt1y),unsafe.Pointer(&ctrlPt2x),unsafe.Pointer(&ctrlPt2y),unsafe.Pointer(&endPtx),unsafe.Pointer(&endPty),nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::currentPosition()
func (q *QPainterPath) CurrentPosition() *QPointF {
	var __rv uintptr
	q.Drv(88000,88135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPainterPath::elementAt(int)
func (q *QPainterPath) ElementAt(i int) *QPainterPathElement {
	var __rv uintptr
	q.Drv(88000,88136,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPathElement{}
	_rp.SetDriver(__rv,89000,true)
	return _rp
}	
//QPainterPath::elementCount()
func (q *QPainterPath) ElementCount() int {
	var __rv int
	q.Drv(88000,88137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::fillRule()
func (q *QPainterPath) FillRule() Qt_FillRule {
	var __rv Qt_FillRule
	q.Drv(88000,88138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::intersected(QPainterPath const&)
func (q *QPainterPath) Intersected(r *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88139,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::intersects(QPainterPath const&)
func (q *QPainterPath) Intersects(p *QPainterPath) bool {
	var __rv bool
	q.Drv(88000,88140,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::intersects(QRectF const&)
func (q *QPainterPath) IntersectsFWithRect(rect *QRectF) bool {
	var __rv bool
	q.Drv(88000,88141,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::isEmpty()
func (q *QPainterPath) IsEmpty() bool {
	var __rv bool
	q.Drv(88000,88142,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::length()
func (q *QPainterPath) Length() float64 {
	var __rv float64
	q.Drv(88000,88143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::lineTo(QPointF const&)
func (q *QPainterPath) LineTo(p *QPointF)  {
	q.Drv(88000,88144,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::lineTo(double,double)
func (q *QPainterPath) LineToFWithXY(x float64,y float64)  {
	q.Drv(88000,88145,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::moveTo(QPointF const&)
func (q *QPainterPath) MoveTo(p *QPointF)  {
	q.Drv(88000,88146,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::moveTo(double,double)
func (q *QPainterPath) MoveToFWithXY(x float64,y float64)  {
	q.Drv(88000,88147,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::percentAtLength(double)
func (q *QPainterPath) PercentAtLength(t float64) float64 {
	var __rv float64
	q.Drv(88000,88148,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::pointAtPercent(double)
func (q *QPainterPath) PointAtPercent(t float64) *QPointF {
	var __rv uintptr
	q.Drv(88000,88149,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPainterPath::quadTo(QPointF const&,QPointF const&)
func (q *QPainterPath) QuadToFWithCtrlptEndpt(ctrlPt *QPointF,endPt *QPointF)  {
	q.Drv(88000,88150,Native(ctrlPt),Native(endPt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::quadTo(double,double,double,double)
func (q *QPainterPath) QuadToFWithCtrlptxCtrlptyEndptxEndpty(ctrlPtx float64,ctrlPty float64,endPtx float64,endPty float64)  {
	q.Drv(88000,88151,unsafe.Pointer(&ctrlPtx),unsafe.Pointer(&ctrlPty),unsafe.Pointer(&endPtx),unsafe.Pointer(&endPty),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::setElementPositionAt(int,double,double)
func (q *QPainterPath) SetElementPositionAt(i int,x float64,y float64)  {
	q.Drv(88000,88152,unsafe.Pointer(&i),unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::setFillRule(Qt::FillRule)
func (q *QPainterPath) SetFillRule(fillRule Qt_FillRule)  {
	q.Drv(88000,88153,unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::simplified()
func (q *QPainterPath) Simplified() *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::slopeAtPercent(double)
func (q *QPainterPath) SlopeAtPercent(t float64) float64 {
	var __rv float64
	q.Drv(88000,88155,unsafe.Pointer(&t),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::subtracted(QPainterPath const&)
func (q *QPainterPath) Subtracted(r *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88156,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::subtractedInverted(QPainterPath const&)
func (q *QPainterPath) SubtractedInverted(r *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88157,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::toFillPolygon()
func (q *QPainterPath) ToFillPolygon() *QPolygonF {
	var __rv uintptr
	q.Drv(88000,88158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPainterPath::toFillPolygon(QMatrix const&)
func (q *QPainterPath) ToFillPolygonWithMatrix(matrix *QMatrix) *QPolygonF {
	var __rv uintptr
	q.Drv(88000,88159,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPainterPath::toFillPolygon(QTransform const&)
func (q *QPainterPath) ToFillPolygonWithTransform(matrix *QTransform) *QPolygonF {
	var __rv uintptr
	q.Drv(88000,88160,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QPainterPath::toFillPolygons()
func (q *QPainterPath) ToFillPolygons() []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::toFillPolygons(QMatrix const&)
func (q *QPainterPath) ToFillPolygonsWithMatrix(matrix *QMatrix) []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88162,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::toFillPolygons(QTransform const&)
func (q *QPainterPath) ToFillPolygonsWithTransform(matrix *QTransform) []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88163,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::toReversed()
func (q *QPainterPath) ToReversed() *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::toSubpathPolygons()
func (q *QPainterPath) ToSubpathPolygons() []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::toSubpathPolygons(QMatrix const&)
func (q *QPainterPath) ToSubpathPolygonsWithMatrix(matrix *QMatrix) []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88166,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::toSubpathPolygons(QTransform const&)
func (q *QPainterPath) ToSubpathPolygonsWithTransform(matrix *QTransform) []*QPolygonF {
	var __rv []*QPolygonF
	q.Drv(88000,88167,Native(matrix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainterPath::translate(QPointF const&)
func (q *QPainterPath) Translate(offset *QPointF)  {
	q.Drv(88000,88168,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::translate(double,double)
func (q *QPainterPath) TranslateFWithDxDy(dx float64,dy float64)  {
	q.Drv(88000,88169,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainterPath::translated(QPointF const&)
func (q *QPainterPath) Translated(offset *QPointF) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88170,Native(offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::translated(double,double)
func (q *QPainterPath) TranslatedFWithDxDy(dx float64,dy float64) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88171,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainterPath::united(QPainterPath const&)
func (q *QPainterPath) United(r *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(88000,88172,Native(r),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
