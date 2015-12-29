// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPainter_CompositionMode - QPainter::CompositionMode
type QPainter_CompositionMode uint32
const (
	QPainter_CompositionMode_SourceOver QPainter_CompositionMode = 0
	QPainter_CompositionMode_DestinationOver QPainter_CompositionMode = 1
	QPainter_CompositionMode_Clear QPainter_CompositionMode = 2
	QPainter_CompositionMode_Source QPainter_CompositionMode = 3
	QPainter_CompositionMode_Destination QPainter_CompositionMode = 4
	QPainter_CompositionMode_SourceIn QPainter_CompositionMode = 5
	QPainter_CompositionMode_DestinationIn QPainter_CompositionMode = 6
	QPainter_CompositionMode_SourceOut QPainter_CompositionMode = 7
	QPainter_CompositionMode_DestinationOut QPainter_CompositionMode = 8
	QPainter_CompositionMode_SourceAtop QPainter_CompositionMode = 9
	QPainter_CompositionMode_DestinationAtop QPainter_CompositionMode = 10
	QPainter_CompositionMode_Xor QPainter_CompositionMode = 11
	QPainter_CompositionMode_Plus QPainter_CompositionMode = 12
	QPainter_CompositionMode_Multiply QPainter_CompositionMode = 13
	QPainter_CompositionMode_Screen QPainter_CompositionMode = 14
	QPainter_CompositionMode_Overlay QPainter_CompositionMode = 15
	QPainter_CompositionMode_Darken QPainter_CompositionMode = 16
	QPainter_CompositionMode_Lighten QPainter_CompositionMode = 17
	QPainter_CompositionMode_ColorDodge QPainter_CompositionMode = 18
	QPainter_CompositionMode_ColorBurn QPainter_CompositionMode = 19
	QPainter_CompositionMode_HardLight QPainter_CompositionMode = 20
	QPainter_CompositionMode_SoftLight QPainter_CompositionMode = 21
	QPainter_CompositionMode_Difference QPainter_CompositionMode = 22
	QPainter_CompositionMode_Exclusion QPainter_CompositionMode = 23
	QPainter_RasterOp_SourceOrDestination QPainter_CompositionMode = 24
	QPainter_RasterOp_SourceAndDestination QPainter_CompositionMode = 25
	QPainter_RasterOp_SourceXorDestination QPainter_CompositionMode = 26
	QPainter_RasterOp_NotSourceAndNotDestination QPainter_CompositionMode = 27
	QPainter_RasterOp_NotSourceOrNotDestination QPainter_CompositionMode = 28
	QPainter_RasterOp_NotSourceXorDestination QPainter_CompositionMode = 29
	QPainter_RasterOp_NotSource QPainter_CompositionMode = 30
	QPainter_RasterOp_NotSourceAndDestination QPainter_CompositionMode = 31
	QPainter_RasterOp_SourceAndNotDestination QPainter_CompositionMode = 32
)
//enum QPainter_RenderHint - QPainter::RenderHint
type QPainter_RenderHint uint32
const (
	QPainter_Antialiasing QPainter_RenderHint = 0x01
	QPainter_TextAntialiasing QPainter_RenderHint = 0x02
	QPainter_SmoothPixmapTransform QPainter_RenderHint = 0x04
	QPainter_HighQualityAntialiasing QPainter_RenderHint = 0x08
	QPainter_NonCosmeticDefaultPen QPainter_RenderHint = 0x10
)
//enum QPainter_PixmapFragmentHint - QPainter::PixmapFragmentHint
type QPainter_PixmapFragmentHint uint32
const (
	QPainter_OpaqueHint QPainter_PixmapFragmentHint = 0x01
)
//struct QPainter : QPainter
type QPainter struct {
	BaseDrv
}
//QPainter::QPainter()	
func NewPainter() *QPainter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),86000,86102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainter{}
	_p.SetDriver(__rv,86000,true)
	return _p
} 
//QPainter::QPainter(QPaintDevice*)	
func NewPainterWithPaintDevice(value QPaintDeviceInterface) *QPainter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),86000,86103,unsafe.Pointer(new_pd_head(value)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPainter{}
	_p.SetDriver(__rv,86000,true)
	return _p
} 
//QPainter::background()
func (q *QPainter) Background() *QBrush {
	var __rv uintptr
	q.Drv(86000,86104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPainter::backgroundMode()
func (q *QPainter) BackgroundMode() Qt_BGMode {
	var __rv Qt_BGMode
	q.Drv(86000,86105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::begin(QPaintDevice*)
func (q *QPainter) Begin(value QPaintDeviceInterface) bool {
	var __rv bool
	q.Drv(86000,86106,unsafe.Pointer(new_pd_head(value)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::beginNativePainting()
func (q *QPainter) BeginNativePainting()  {
	q.Drv(86000,86107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::boundingRect(QRect const&,int,QString const&)
func (q *QPainter) BoundingRectWithRectFlagsText(rect *QRect,flags int,text string) *QRect {
	var __rv uintptr
	q.Drv(86000,86108,Native(rect),unsafe.Pointer(&flags),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPainter::boundingRect(QRectF const&,QString const&,QTextOption const&)
func (q *QPainter) BoundingRectFWithRectTextOption(rect *QRectF,text string,o *QTextOption) *QRectF {
	var __rv uintptr
	q.Drv(86000,86109,Native(rect),unsafe.Pointer(&text),Native(o),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPainter::boundingRect(QRectF const&,int,QString const&)
func (q *QPainter) BoundingRectFWithRectFlagsText(rect *QRectF,flags int,text string) *QRectF {
	var __rv uintptr
	q.Drv(86000,86110,Native(rect),unsafe.Pointer(&flags),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPainter::boundingRect(int,int,int,int,int,QString const&)
func (q *QPainter) BoundingRectWithXYWidthHeightFlagsText(x int,y int,w int,h int,flags int,text string) *QRect {
	var __rv uintptr
	q.Drv(86000,86111,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&flags),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPainter::brush()
func (q *QPainter) Brush() *QBrush {
	var __rv uintptr
	q.Drv(86000,86112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QPainter::brushOrigin()
func (q *QPainter) BrushOrigin() *QPoint {
	var __rv uintptr
	q.Drv(86000,86113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QPainter::clipPath()
func (q *QPainter) ClipPath() *QPainterPath {
	var __rv uintptr
	q.Drv(86000,86114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QPainter::clipRegion()
func (q *QPainter) ClipRegion() *QRegion {
	var __rv uintptr
	q.Drv(86000,86115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QPainter::combinedMatrix()
func (q *QPainter) CombinedMatrix() *QMatrix {
	var __rv uintptr
	q.Drv(86000,86116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPainter::combinedTransform()
func (q *QPainter) CombinedTransform() *QTransform {
	var __rv uintptr
	q.Drv(86000,86117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QPainter::compositionMode()
func (q *QPainter) CompositionMode() QPainter_CompositionMode {
	var __rv QPainter_CompositionMode
	q.Drv(86000,86118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::device()
func (q *QPainter) Device() *QPaintDevice {
	var __rv uintptr
	q.Drv(86000,86119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPainter::deviceMatrix()
func (q *QPainter) DeviceMatrix() *QMatrix {
	var __rv uintptr
	q.Drv(86000,86120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPainter::deviceTransform()
func (q *QPainter) DeviceTransform() *QTransform {
	var __rv uintptr
	q.Drv(86000,86121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QPainter::drawArc(QRect const&,int,int)
func (q *QPainter) DrawArcWithRectStartangleSpanangle(value2 *QRect,startAngle int,spanAngle int)  {
	q.Drv(86000,86122,Native(value2),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawArc(QRectF const&,int,int)
func (q *QPainter) DrawArcFWithRectStartangleSpanangle(rect *QRectF,startAngle int,spanAngle int)  {
	q.Drv(86000,86123,Native(rect),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawArc(int,int,int,int,int,int)
func (q *QPainter) DrawArcWithXYWidthHeightStartangleSpanangle(x int,y int,w int,h int,startAngle int,spanAngle int)  {
	q.Drv(86000,86124,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawChord(QRect const&,int,int)
func (q *QPainter) DrawChordWithRectStartangleSpanangle(value2 *QRect,startAngle int,spanAngle int)  {
	q.Drv(86000,86125,Native(value2),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawChord(QRectF const&,int,int)
func (q *QPainter) DrawChordFWithRectStartangleSpanangle(rect *QRectF,startAngle int,spanAngle int)  {
	q.Drv(86000,86126,Native(rect),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawChord(int,int,int,int,int,int)
func (q *QPainter) DrawChordWithXYWidthHeightStartangleSpanangle(x int,y int,w int,h int,startAngle int,spanAngle int)  {
	q.Drv(86000,86127,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawConvexPolygon(QPolygon const&)
func (q *QPainter) DrawConvexPolygon(polygon *QPolygon)  {
	q.Drv(86000,86128,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawConvexPolygon(QPolygonF const&)
func (q *QPainter) DrawConvexPolygonFWithPolygon(polygon *QPolygonF)  {
	q.Drv(86000,86129,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawConvexPolygon(QPoint const*,int)
func (q *QPainter) DrawConvexPolygonWithPointsPointcount(points *QPoint,pointCount int)  {
	q.Drv(86000,86130,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawConvexPolygon(QPointF const*,int)
func (q *QPainter) DrawConvexPolygonFWithPointsPointcount(points *QPointF,pointCount int)  {
	q.Drv(86000,86131,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawEllipse(QRect const&)
func (q *QPainter) DrawEllipse(r *QRect)  {
	q.Drv(86000,86132,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawEllipse(QRectF const&)
func (q *QPainter) DrawEllipseFWithRectf(r *QRectF)  {
	q.Drv(86000,86133,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawEllipse(QPoint const&,int,int)
func (q *QPainter) DrawEllipseWithCenterRxRy(center *QPoint,rx int,ry int)  {
	q.Drv(86000,86134,Native(center),unsafe.Pointer(&rx),unsafe.Pointer(&ry),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawEllipse(QPointF const&,double,double)
func (q *QPainter) DrawEllipseFWithCenterRxRy(center *QPointF,rx float64,ry float64)  {
	q.Drv(86000,86135,Native(center),unsafe.Pointer(&rx),unsafe.Pointer(&ry),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawEllipse(int,int,int,int)
func (q *QPainter) DrawEllipseWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(86000,86136,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QPoint const&,QImage const&)
func (q *QPainter) DrawImageWithPointImage(p *QPoint,image *QImage)  {
	q.Drv(86000,86137,Native(p),Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QPointF const&,QImage const&)
func (q *QPainter) DrawImageFWithPointfImage(p *QPointF,image *QImage)  {
	q.Drv(86000,86138,Native(p),Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QRect const&,QImage const&)
func (q *QPainter) DrawImageWithRectImage(r *QRect,image *QImage)  {
	q.Drv(86000,86139,Native(r),Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QRectF const&,QImage const&)
func (q *QPainter) DrawImageFWithRectfImage(r *QRectF,image *QImage)  {
	q.Drv(86000,86140,Native(r),Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QPoint const&,QImage const&,QRect const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPainter) DrawImageWithPointImageRectFlags(p *QPoint,image *QImage,sr *QRect,flags Qt_ImageConversionFlag)  {
	q.Drv(86000,86141,Native(p),Native(image),Native(sr),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QPointF const&,QImage const&,QRectF const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPainter) DrawImageFWithPointfImageRectfFlags(p *QPointF,image *QImage,sr *QRectF,flags Qt_ImageConversionFlag)  {
	q.Drv(86000,86142,Native(p),Native(image),Native(sr),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QRect const&,QImage const&,QRect const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPainter) DrawImageWithTargetrectImageSourcerectFlags(targetRect *QRect,image *QImage,sourceRect *QRect,flags Qt_ImageConversionFlag)  {
	q.Drv(86000,86143,Native(targetRect),Native(image),Native(sourceRect),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(QRectF const&,QImage const&,QRectF const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPainter) DrawImageFWithTargetrectImageSourcerectFlags(targetRect *QRectF,image *QImage,sourceRect *QRectF,flags Qt_ImageConversionFlag)  {
	q.Drv(86000,86144,Native(targetRect),Native(image),Native(sourceRect),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawImage(int,int,QImage const&,int,int,int,int,QFlags<Qt::ImageConversionFlag>)
func (q *QPainter) DrawImageWithXYImageSxSySwShFlags(x int,y int,image *QImage,sx int,sy int,sw int,sh int,flags Qt_ImageConversionFlag)  {
	q.Drv(86000,86145,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(image),unsafe.Pointer(&sx),unsafe.Pointer(&sy),unsafe.Pointer(&sw),unsafe.Pointer(&sh),unsafe.Pointer(&flags),nil,nil,nil,nil)
}	
//QPainter::drawLine(QLine const&)
func (q *QPainter) DrawLine(line *QLine)  {
	q.Drv(86000,86146,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLine(QLineF const&)
func (q *QPainter) DrawLineFWithLine(line *QLineF)  {
	q.Drv(86000,86147,Native(line),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLine(QPoint const&,QPoint const&)
func (q *QPainter) DrawLineWithPointPoint(p1 *QPoint,p2 *QPoint)  {
	q.Drv(86000,86148,Native(p1),Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLine(QPointF const&,QPointF const&)
func (q *QPainter) DrawLineFWithPointfPointf(p1 *QPointF,p2 *QPointF)  {
	q.Drv(86000,86149,Native(p1),Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLine(int,int,int,int)
func (q *QPainter) DrawLineWithX1Y1X2Y2(x1 int,y1 int,x2 int,y2 int)  {
	q.Drv(86000,86150,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QVector<QLine> const&)
func (q *QPainter) DrawLines(lines []*QLine)  {
	q.Drv(86000,86151,unsafe.Pointer(&lines),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QVector<QLineF> const&)
func (q *QPainter) DrawLinesWithLinefs(lines []*QLineF)  {
	q.Drv(86000,86152,unsafe.Pointer(&lines),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QVector<QPoint> const&)
func (q *QPainter) DrawLinesWithPointpairs(pointPairs []*QPoint)  {
	q.Drv(86000,86153,unsafe.Pointer(&pointPairs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QVector<QPointF> const&)
func (q *QPainter) DrawLinesWithPointfpairs(pointPairs []*QPointF)  {
	q.Drv(86000,86154,unsafe.Pointer(&pointPairs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QLine const*,int)
func (q *QPainter) DrawLinesWithLinesLinecount(lines *QLine,lineCount int)  {
	q.Drv(86000,86155,Native(lines),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QLineF const*,int)
func (q *QPainter) DrawLinesFWithLinesLinecount(lines *QLineF,lineCount int)  {
	q.Drv(86000,86156,Native(lines),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QPoint const*,int)
func (q *QPainter) DrawLinesWithPointpairsLinecount(pointPairs *QPoint,lineCount int)  {
	q.Drv(86000,86157,Native(pointPairs),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawLines(QPointF const*,int)
func (q *QPainter) DrawLinesFWithPointpairsLinecount(pointPairs *QPointF,lineCount int)  {
	q.Drv(86000,86158,Native(pointPairs),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPath(QPainterPath const&)
func (q *QPainter) DrawPath(path *QPainterPath)  {
	q.Drv(86000,86159,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPicture(QPoint const&,QPicture const&)
func (q *QPainter) DrawPictureWithPointPicture(p *QPoint,picture *QPicture)  {
	q.Drv(86000,86160,Native(p),Native(picture),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPicture(QPointF const&,QPicture const&)
func (q *QPainter) DrawPictureFWithPointfPicture(p *QPointF,picture *QPicture)  {
	q.Drv(86000,86161,Native(p),Native(picture),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPicture(int,int,QPicture const&)
func (q *QPainter) DrawPictureWithXYPicture(x int,y int,picture *QPicture)  {
	q.Drv(86000,86162,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(picture),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPie(QRect const&,int,int)
func (q *QPainter) DrawPieWithRectStartangleSpanangle(value2 *QRect,startAngle int,spanAngle int)  {
	q.Drv(86000,86163,Native(value2),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPie(QRectF const&,int,int)
func (q *QPainter) DrawPieFWithRectStartangleSpanangle(rect *QRectF,startAngle int,spanAngle int)  {
	q.Drv(86000,86164,Native(rect),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPie(int,int,int,int,int,int)
func (q *QPainter) DrawPieWithXYWidthHeightStartangleSpanangle(x int,y int,w int,h int,startAngle int,spanAngle int)  {
	q.Drv(86000,86165,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&startAngle),unsafe.Pointer(&spanAngle),nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QPoint const&,QPixmap const&)
func (q *QPainter) DrawPixmapWithPointPixmap(p *QPoint,pm *QPixmap)  {
	q.Drv(86000,86166,Native(p),Native(pm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QPointF const&,QPixmap const&)
func (q *QPainter) DrawPixmapFWithPointfPixmap(p *QPointF,pm *QPixmap)  {
	q.Drv(86000,86167,Native(p),Native(pm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QRect const&,QPixmap const&)
func (q *QPainter) DrawPixmapWithRectPixmap(r *QRect,pm *QPixmap)  {
	q.Drv(86000,86168,Native(r),Native(pm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QPoint const&,QPixmap const&,QRect const&)
func (q *QPainter) DrawPixmapWithPointPixmapRect(p *QPoint,pm *QPixmap,sr *QRect)  {
	q.Drv(86000,86169,Native(p),Native(pm),Native(sr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QPointF const&,QPixmap const&,QRectF const&)
func (q *QPainter) DrawPixmapFWithPointfPixmapRectf(p *QPointF,pm *QPixmap,sr *QRectF)  {
	q.Drv(86000,86170,Native(p),Native(pm),Native(sr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QRect const&,QPixmap const&,QRect const&)
func (q *QPainter) DrawPixmapWithTargetrectPixmapSourcerect(targetRect *QRect,pixmap *QPixmap,sourceRect *QRect)  {
	q.Drv(86000,86171,Native(targetRect),Native(pixmap),Native(sourceRect),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(QRectF const&,QPixmap const&,QRectF const&)
func (q *QPainter) DrawPixmapFWithTargetrectPixmapSourcerect(targetRect *QRectF,pixmap *QPixmap,sourceRect *QRectF)  {
	q.Drv(86000,86172,Native(targetRect),Native(pixmap),Native(sourceRect),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(int,int,QPixmap const&)
func (q *QPainter) DrawPixmapWithXYPixmap(x int,y int,pm *QPixmap)  {
	q.Drv(86000,86173,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(pm),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(int,int,int,int,QPixmap const&)
func (q *QPainter) DrawPixmapWithXYWidthHeightPixmap(x int,y int,w int,h int,pm *QPixmap)  {
	q.Drv(86000,86174,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(pm),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(int,int,QPixmap const&,int,int,int,int)
func (q *QPainter) DrawPixmapWithXYPixmapSxSySwSh(x int,y int,pm *QPixmap,sx int,sy int,sw int,sh int)  {
	q.Drv(86000,86175,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(pm),unsafe.Pointer(&sx),unsafe.Pointer(&sy),unsafe.Pointer(&sw),unsafe.Pointer(&sh),nil,nil,nil,nil,nil)
}	
//QPainter::drawPixmap(int,int,int,int,QPixmap const&,int,int,int,int)
func (q *QPainter) DrawPixmapWithXYWidthHeightPixmapSxSySwSh(x int,y int,w int,h int,pm *QPixmap,sx int,sy int,sw int,sh int)  {
	q.Drv(86000,86176,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(pm),unsafe.Pointer(&sx),unsafe.Pointer(&sy),unsafe.Pointer(&sw),unsafe.Pointer(&sh),nil,nil,nil)
}	
//QPainter::drawPixmapFragments(QPainter::PixmapFragment const*,int,QPixmap const&,QFlags<QPainter::PixmapFragmentHint>)
func (q *QPainter) DrawPixmapFragments(fragments *QPainterPixmapFragment,fragmentCount int,pixmap *QPixmap,hints QPainter_PixmapFragmentHint)  {
	q.Drv(86000,86177,Native(fragments),unsafe.Pointer(&fragmentCount),Native(pixmap),unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoint(QPoint const&)
func (q *QPainter) DrawPoint(p *QPoint)  {
	q.Drv(86000,86178,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoint(QPointF const&)
func (q *QPainter) DrawPointFWithPointf(pt *QPointF)  {
	q.Drv(86000,86179,Native(pt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoint(int,int)
func (q *QPainter) DrawPointWithXY(x int,y int)  {
	q.Drv(86000,86180,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoints(QPolygon const&)
func (q *QPainter) DrawPoints(points *QPolygon)  {
	q.Drv(86000,86181,Native(points),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoints(QPolygonF const&)
func (q *QPainter) DrawPointsFWithPoints(points *QPolygonF)  {
	q.Drv(86000,86182,Native(points),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoints(QPoint const*,int)
func (q *QPainter) DrawPointsWithPointsPointcount(points *QPoint,pointCount int)  {
	q.Drv(86000,86183,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPoints(QPointF const*,int)
func (q *QPainter) DrawPointsFWithPointsPointcount(points *QPointF,pointCount int)  {
	q.Drv(86000,86184,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPolygon const&)
func (q *QPainter) DrawPolygon(polygon *QPolygon)  {
	q.Drv(86000,86185,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPolygonF const&)
func (q *QPainter) DrawPolygonFWithPolygon(polygon *QPolygonF)  {
	q.Drv(86000,86186,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPolygon const&,Qt::FillRule)
func (q *QPainter) DrawPolygonWithPolygonFillrule(polygon *QPolygon,fillRule Qt_FillRule)  {
	q.Drv(86000,86187,Native(polygon),unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPolygonF const&,Qt::FillRule)
func (q *QPainter) DrawPolygonFWithPolygonFillrule(polygon *QPolygonF,fillRule Qt_FillRule)  {
	q.Drv(86000,86188,Native(polygon),unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPoint const*,int,Qt::FillRule)
func (q *QPainter) DrawPolygonWithPointsPointcountFillrule(points *QPoint,pointCount int,fillRule Qt_FillRule)  {
	q.Drv(86000,86189,Native(points),unsafe.Pointer(&pointCount),unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolygon(QPointF const*,int,Qt::FillRule)
func (q *QPainter) DrawPolygonFWithPointsPointcountFillrule(points *QPointF,pointCount int,fillRule Qt_FillRule)  {
	q.Drv(86000,86190,Native(points),unsafe.Pointer(&pointCount),unsafe.Pointer(&fillRule),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolyline(QPolygon const&)
func (q *QPainter) DrawPolyline(polygon *QPolygon)  {
	q.Drv(86000,86191,Native(polygon),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolyline(QPolygonF const&)
func (q *QPainter) DrawPolylineFWithPolyline(polyline *QPolygonF)  {
	q.Drv(86000,86192,Native(polyline),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolyline(QPoint const*,int)
func (q *QPainter) DrawPolylineWithPointsPointcount(points *QPoint,pointCount int)  {
	q.Drv(86000,86193,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawPolyline(QPointF const*,int)
func (q *QPainter) DrawPolylineFWithPointsPointcount(points *QPointF,pointCount int)  {
	q.Drv(86000,86194,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRect(QRect const&)
func (q *QPainter) DrawRect(rect *QRect)  {
	q.Drv(86000,86195,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRect(QRectF const&)
func (q *QPainter) DrawRectFWithRect(rect *QRectF)  {
	q.Drv(86000,86196,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRect(int,int,int,int)
func (q *QPainter) DrawRectWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(86000,86197,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRects(QVector<QRect> const&)
func (q *QPainter) DrawRects(rectangles []*QRect)  {
	q.Drv(86000,86198,unsafe.Pointer(&rectangles),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRects(QVector<QRectF> const&)
func (q *QPainter) DrawRectsWithRectanglefs(rectangles []*QRectF)  {
	q.Drv(86000,86199,unsafe.Pointer(&rectangles),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRects(QRect const*,int)
func (q *QPainter) DrawRectsWithRectsRectcount(rects *QRect,rectCount int)  {
	q.Drv(86000,86200,Native(rects),unsafe.Pointer(&rectCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRects(QRectF const*,int)
func (q *QPainter) DrawRectsFWithRectsRectcount(rects *QRectF,rectCount int)  {
	q.Drv(86000,86201,Native(rects),unsafe.Pointer(&rectCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundRect(QRect const&)
func (q *QPainter) DrawRoundRect(r *QRect)  {
	q.Drv(86000,86202,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundRect(QRectF const&)
func (q *QPainter) DrawRoundRectFWithRectf(r *QRectF)  {
	q.Drv(86000,86203,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundRect(QRect const&,int,int)
func (q *QPainter) DrawRoundRectWithRectXroundYround(r *QRect,xround int,yround int)  {
	q.Drv(86000,86204,Native(r),unsafe.Pointer(&xround),unsafe.Pointer(&yround),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundRect(QRectF const&,int,int)
func (q *QPainter) DrawRoundRectFWithRectfXroundYround(r *QRectF,xround int,yround int)  {
	q.Drv(86000,86205,Native(r),unsafe.Pointer(&xround),unsafe.Pointer(&yround),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundRect(int,int,int,int,int,int)
func (q *QPainter) DrawRoundRectWithXYWidthHeightIntInt(x int,y int,w int,h int,value2 int,value3 int)  {
	q.Drv(86000,86206,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundedRect(QRect const&,double,double,Qt::SizeMode)
func (q *QPainter) DrawRoundedRectWithRectXradiusYradiusMode(rect *QRect,xRadius float64,yRadius float64,mode Qt_SizeMode)  {
	q.Drv(86000,86207,Native(rect),unsafe.Pointer(&xRadius),unsafe.Pointer(&yRadius),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundedRect(QRectF const&,double,double,Qt::SizeMode)
func (q *QPainter) DrawRoundedRectFWithRectXradiusYradiusMode(rect *QRectF,xRadius float64,yRadius float64,mode Qt_SizeMode)  {
	q.Drv(86000,86208,Native(rect),unsafe.Pointer(&xRadius),unsafe.Pointer(&yRadius),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawRoundedRect(int,int,int,int,double,double,Qt::SizeMode)
func (q *QPainter) DrawRoundedRectFWithXYWidthHeightXradiusYradiusMode(x int,y int,w int,h int,xRadius float64,yRadius float64,mode Qt_SizeMode)  {
	q.Drv(86000,86209,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&xRadius),unsafe.Pointer(&yRadius),unsafe.Pointer(&mode),nil,nil,nil,nil,nil)
}	
//QPainter::drawStaticText(QPoint const&,QStaticText const&)
func (q *QPainter) DrawStaticTextWithTopleftpositionStatictext(topLeftPosition *QPoint,staticText *QStaticText)  {
	q.Drv(86000,86210,Native(topLeftPosition),Native(staticText),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawStaticText(QPointF const&,QStaticText const&)
func (q *QPainter) DrawStaticTextFWithTopleftpositionStatictext(topLeftPosition *QPointF,staticText *QStaticText)  {
	q.Drv(86000,86211,Native(topLeftPosition),Native(staticText),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawStaticText(int,int,QStaticText const&)
func (q *QPainter) DrawStaticTextWithLeftTopStatictext(left int,top int,staticText *QStaticText)  {
	q.Drv(86000,86212,unsafe.Pointer(&left),unsafe.Pointer(&top),Native(staticText),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QPoint const&,QString const&)
func (q *QPainter) DrawTextWithPointText(p *QPoint,s string)  {
	q.Drv(86000,86213,Native(p),unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QPointF const&,QString const&)
func (q *QPainter) DrawTextFWithPointfText(p *QPointF,s string)  {
	q.Drv(86000,86214,Native(p),unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QRectF const&,QString const&,QTextOption const&)
func (q *QPainter) DrawTextFWithRectfTextOption(r *QRectF,text string,o *QTextOption)  {
	q.Drv(86000,86215,Native(r),unsafe.Pointer(&text),Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(int,int,QString const&)
func (q *QPainter) DrawTextWithXYText(x int,y int,s string)  {
	q.Drv(86000,86216,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QPointF const&,QString const&,int,int)
func (q *QPainter) DrawTextFWithPointfTextTfJustificationpadding(p *QPointF,str string,tf int,justificationPadding int)  {
	q.Drv(86000,86217,Native(p),unsafe.Pointer(&str),unsafe.Pointer(&tf),unsafe.Pointer(&justificationPadding),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QRect const&,int,QString const&,QRect*)
func (q *QPainter) DrawTextWithRectFlagsTextRect(r *QRect,flags int,text string,br *QRect)  {
	q.Drv(86000,86218,Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&text),Native(br),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(QRectF const&,int,QString const&,QRectF*)
func (q *QPainter) DrawTextFWithRectfFlagsTextRectf(r *QRectF,flags int,text string,br *QRectF)  {
	q.Drv(86000,86219,Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&text),Native(br),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawText(int,int,int,int,int,QString const&,QRect*)
func (q *QPainter) DrawTextWithXYWidthHeightFlagsTextRect(x int,y int,w int,h int,flags int,text string,br *QRect)  {
	q.Drv(86000,86220,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&flags),unsafe.Pointer(&text),Native(br),nil,nil,nil,nil,nil)
}	
//QPainter::drawTextItem(QPoint const&,QTextItem const&)
func (q *QPainter) DrawTextItemWithPointTi(p *QPoint,ti *QTextItem)  {
	q.Drv(86000,86221,Native(p),Native(ti),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawTextItem(QPointF const&,QTextItem const&)
func (q *QPainter) DrawTextItemFWithPointfTi(p *QPointF,ti *QTextItem)  {
	q.Drv(86000,86222,Native(p),Native(ti),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawTextItem(int,int,QTextItem const&)
func (q *QPainter) DrawTextItemWithXYTi(x int,y int,ti *QTextItem)  {
	q.Drv(86000,86223,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(ti),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawTiledPixmap(QRect const&,QPixmap const&,QPoint const&)
func (q *QPainter) DrawTiledPixmapWithRectPixmapPoint(value2 *QRect,value3 *QPixmap,value4 *QPoint)  {
	q.Drv(86000,86224,Native(value2),Native(value3),Native(value4),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawTiledPixmap(QRectF const&,QPixmap const&,QPointF const&)
func (q *QPainter) DrawTiledPixmapFWithRectPixmapOffset(rect *QRectF,pm *QPixmap,offset *QPointF)  {
	q.Drv(86000,86225,Native(rect),Native(pm),Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::drawTiledPixmap(int,int,int,int,QPixmap const&,int,int)
func (q *QPainter) DrawTiledPixmapWithXYWidthHeightPixmapSxSy(x int,y int,w int,h int,value2 *QPixmap,sx int,sy int)  {
	q.Drv(86000,86226,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(value2),unsafe.Pointer(&sx),unsafe.Pointer(&sy),nil,nil,nil,nil,nil)
}	
//QPainter::end()
func (q *QPainter) End() bool {
	var __rv bool
	q.Drv(86000,86227,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::endNativePainting()
func (q *QPainter) EndNativePainting()  {
	q.Drv(86000,86228,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::eraseRect(QRect const&)
func (q *QPainter) EraseRect(value *QRect)  {
	q.Drv(86000,86229,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::eraseRect(QRectF const&)
func (q *QPainter) EraseRectFWithRectf(value *QRectF)  {
	q.Drv(86000,86230,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::eraseRect(int,int,int,int)
func (q *QPainter) EraseRectWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(86000,86231,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillPath(QPainterPath const&,QBrush const&)
func (q *QPainter) FillPath(path *QPainterPath,brush *QBrush)  {
	q.Drv(86000,86232,Native(path),Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRect const&,QBrush const&)
func (q *QPainter) FillRectWithRectBrush(value2 *QRect,value3 *QBrush)  {
	q.Drv(86000,86233,Native(value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRect const&,QColor const&)
func (q *QPainter) FillRectWithRectColor(value2 *QRect,color *QColor)  {
	q.Drv(86000,86234,Native(value2),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRect const&,Qt::BrushStyle)
func (q *QPainter) FillRectWithRectStyle(r *QRect,style Qt_BrushStyle)  {
	q.Drv(86000,86235,Native(r),unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRect const&,Qt::GlobalColor)
func (q *QPainter) FillRectWithRectGlobalcolor(r *QRect,c Qt_GlobalColor)  {
	q.Drv(86000,86236,Native(r),unsafe.Pointer(&c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRectF const&,QBrush const&)
func (q *QPainter) FillRectFWithRectfBrush(value2 *QRectF,value3 *QBrush)  {
	q.Drv(86000,86237,Native(value2),Native(value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRectF const&,QColor const&)
func (q *QPainter) FillRectFWithRectfColor(value2 *QRectF,color *QColor)  {
	q.Drv(86000,86238,Native(value2),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRectF const&,Qt::BrushStyle)
func (q *QPainter) FillRectFWithRectfStyle(r *QRectF,style Qt_BrushStyle)  {
	q.Drv(86000,86239,Native(r),unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(QRectF const&,Qt::GlobalColor)
func (q *QPainter) FillRectFWithRectfGlobalcolor(r *QRectF,c Qt_GlobalColor)  {
	q.Drv(86000,86240,Native(r),unsafe.Pointer(&c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(int,int,int,int,QBrush const&)
func (q *QPainter) FillRectWithXYWidthHeightBrush(x int,y int,w int,h int,value2 *QBrush)  {
	q.Drv(86000,86241,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(value2),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(int,int,int,int,QColor const&)
func (q *QPainter) FillRectWithXYWidthHeightColor(x int,y int,w int,h int,color *QColor)  {
	q.Drv(86000,86242,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(color),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(int,int,int,int,Qt::BrushStyle)
func (q *QPainter) FillRectWithXYWidthHeightStyle(x int,y int,w int,h int,style Qt_BrushStyle)  {
	q.Drv(86000,86243,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::fillRect(int,int,int,int,Qt::GlobalColor)
func (q *QPainter) FillRectWithXYWidthHeightGlobalcolor(x int,y int,w int,h int,c Qt_GlobalColor)  {
	q.Drv(86000,86244,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&c),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::font()
func (q *QPainter) Font() *QFont {
	var __rv uintptr
	q.Drv(86000,86245,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QPainter::fontInfo()
func (q *QPainter) FontInfo() *QFontInfo {
	var __rv uintptr
	q.Drv(86000,86246,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontInfo{}
	_rp.SetDriver(__rv,39000,true)
	return _rp
}	
//QPainter::fontMetrics()
func (q *QPainter) FontMetrics() *QFontMetrics {
	var __rv uintptr
	q.Drv(86000,86247,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFontMetrics{}
	_rp.SetDriver(__rv,40000,true)
	return _rp
}	
//QPainter::hasClipping()
func (q *QPainter) HasClipping() bool {
	var __rv bool
	q.Drv(86000,86248,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::initFrom(QWidget const*)
func (q *QPainter) InitFrom(widget QWidgetInterface)  {
	q.Drv(86000,86249,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::isActive()
func (q *QPainter) IsActive() bool {
	var __rv bool
	q.Drv(86000,86250,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::layoutDirection()
func (q *QPainter) LayoutDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(86000,86251,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::matrix()
func (q *QPainter) Matrix() *QMatrix {
	var __rv uintptr
	q.Drv(86000,86252,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPainter::matrixEnabled()
func (q *QPainter) MatrixEnabled() bool {
	var __rv bool
	q.Drv(86000,86253,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::opacity()
func (q *QPainter) Opacity() float64 {
	var __rv float64
	q.Drv(86000,86254,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::paintEngine()
func (q *QPainter) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(86000,86255,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QPainter::pen()
func (q *QPainter) Pen() *QPen {
	var __rv uintptr
	q.Drv(86000,86256,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPen{}
	_rp.SetDriver(__rv,92000,true)
	return _rp
}	
//QPainter::redirected(QPaintDevice const*)	
func QPainterRedirected(device QPaintDeviceInterface) *QPaintDevice {
	var __rv uintptr
	DirectQtDrv(nil,86000,86257,unsafe.Pointer(new_pd_head(device)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPainter::redirected(QPaintDevice const*)
func (q *QPainter) Redirected(device QPaintDeviceInterface) *QPaintDevice {
	var __rv uintptr
	q.Drv(86000,86257,unsafe.Pointer(new_pd_head(device)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPainter::redirected(QPaintDevice const*,QPoint*)	
func QPainterRedirectedWithPaintDeviceOffset(device QPaintDeviceInterface,offset *QPoint) *QPaintDevice {
	var __rv uintptr
	DirectQtDrv(nil,86000,86258,unsafe.Pointer(new_pd_head(device)),Native(offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPainter::redirected(QPaintDevice const*,QPoint*)
func (q *QPainter) RedirectedWithPaintDeviceOffset(device QPaintDeviceInterface,offset *QPoint) *QPaintDevice {
	var __rv uintptr
	q.Drv(86000,86258,unsafe.Pointer(new_pd_head(device)),Native(offset),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPainter::renderHints()
func (q *QPainter) RenderHints() QPainter_RenderHint {
	var __rv QPainter_RenderHint
	q.Drv(86000,86259,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::resetMatrix()
func (q *QPainter) ResetMatrix()  {
	q.Drv(86000,86260,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::resetTransform()
func (q *QPainter) ResetTransform()  {
	q.Drv(86000,86261,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::restore()
func (q *QPainter) Restore()  {
	q.Drv(86000,86262,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::restoreRedirected(QPaintDevice const*)	
func QPainterRestoreRedirected(device QPaintDeviceInterface)  {
	DirectQtDrv(nil,86000,86263,unsafe.Pointer(new_pd_head(device)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::restoreRedirected(QPaintDevice const*)
func (q *QPainter) RestoreRedirected(device QPaintDeviceInterface)  {
	q.Drv(86000,86263,unsafe.Pointer(new_pd_head(device)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::rotate(double)
func (q *QPainter) Rotate(a float64)  {
	q.Drv(86000,86264,unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::save()
func (q *QPainter) Save()  {
	q.Drv(86000,86265,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::scale(double,double)
func (q *QPainter) Scale(sx float64,sy float64)  {
	q.Drv(86000,86266,unsafe.Pointer(&sx),unsafe.Pointer(&sy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBackground(QBrush const&)
func (q *QPainter) SetBackground(bg *QBrush)  {
	q.Drv(86000,86267,Native(bg),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBackgroundMode(Qt::BGMode)
func (q *QPainter) SetBackgroundMode(mode Qt_BGMode)  {
	q.Drv(86000,86268,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBrush(QBrush const&)
func (q *QPainter) SetBrush(brush *QBrush)  {
	q.Drv(86000,86269,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBrush(Qt::BrushStyle)
func (q *QPainter) SetBrushWithStyle(style Qt_BrushStyle)  {
	q.Drv(86000,86270,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBrushOrigin(QPoint const&)
func (q *QPainter) SetBrushOrigin(value *QPoint)  {
	q.Drv(86000,86271,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBrushOrigin(QPointF const&)
func (q *QPainter) SetBrushOriginFWithPointf(value *QPointF)  {
	q.Drv(86000,86272,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setBrushOrigin(int,int)
func (q *QPainter) SetBrushOriginWithXY(x int,y int)  {
	q.Drv(86000,86273,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipPath(QPainterPath const&)
func (q *QPainter) SetClipPath(path *QPainterPath)  {
	q.Drv(86000,86274,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipPath(QPainterPath const&,Qt::ClipOperation)
func (q *QPainter) SetClipPathWithPathOp(path *QPainterPath,op Qt_ClipOperation)  {
	q.Drv(86000,86275,Native(path),unsafe.Pointer(&op),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRect(QRect const&)
func (q *QPainter) SetClipRect(value *QRect)  {
	q.Drv(86000,86276,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRect(QRectF const&)
func (q *QPainter) SetClipRectFWithRectf(value *QRectF)  {
	q.Drv(86000,86277,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRect(QRect const&,Qt::ClipOperation)
func (q *QPainter) SetClipRectWithRectOp(value2 *QRect,op Qt_ClipOperation)  {
	q.Drv(86000,86278,Native(value2),unsafe.Pointer(&op),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRect(QRectF const&,Qt::ClipOperation)
func (q *QPainter) SetClipRectFWithRectfOp(value2 *QRectF,op Qt_ClipOperation)  {
	q.Drv(86000,86279,Native(value2),unsafe.Pointer(&op),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRect(int,int,int,int,Qt::ClipOperation)
func (q *QPainter) SetClipRectWithXYWidthHeightOp(x int,y int,w int,h int,op Qt_ClipOperation)  {
	q.Drv(86000,86280,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&op),nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRegion(QRegion const&)
func (q *QPainter) SetClipRegion(value *QRegion)  {
	q.Drv(86000,86281,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipRegion(QRegion const&,Qt::ClipOperation)
func (q *QPainter) SetClipRegionWithRegionOp(value2 *QRegion,op Qt_ClipOperation)  {
	q.Drv(86000,86282,Native(value2),unsafe.Pointer(&op),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setClipping(bool)
func (q *QPainter) SetClipping(enable bool)  {
	q.Drv(86000,86283,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setCompositionMode(QPainter::CompositionMode)
func (q *QPainter) SetCompositionMode(mode QPainter_CompositionMode)  {
	q.Drv(86000,86284,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setFont(QFont const&)
func (q *QPainter) SetFont(f *QFont)  {
	q.Drv(86000,86285,Native(f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setLayoutDirection(Qt::LayoutDirection)
func (q *QPainter) SetLayoutDirection(direction Qt_LayoutDirection)  {
	q.Drv(86000,86286,unsafe.Pointer(&direction),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setMatrix(QMatrix const&)
func (q *QPainter) SetMatrix(matrix *QMatrix)  {
	q.Drv(86000,86287,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setMatrix(QMatrix const&,bool)
func (q *QPainter) SetMatrixWithMatrixCombine(matrix *QMatrix,combine bool)  {
	q.Drv(86000,86288,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setMatrixEnabled(bool)
func (q *QPainter) SetMatrixEnabled(enabled bool)  {
	q.Drv(86000,86289,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setOpacity(double)
func (q *QPainter) SetOpacity(opacity float64)  {
	q.Drv(86000,86290,unsafe.Pointer(&opacity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setPen(QColor const&)
func (q *QPainter) SetPen(color *QColor)  {
	q.Drv(86000,86291,Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setPen(QPen const&)
func (q *QPainter) SetPenWithPen(pen *QPen)  {
	q.Drv(86000,86292,Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setPen(Qt::PenStyle)
func (q *QPainter) SetPenWithStyle(style Qt_PenStyle)  {
	q.Drv(86000,86293,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRedirected(QPaintDevice const*,QPaintDevice*,QPoint const&)	
func QPainterSetRedirected(device QPaintDeviceInterface,replacement QPaintDeviceInterface,offset *QPoint)  {
	DirectQtDrv(nil,86000,86294,unsafe.Pointer(new_pd_head(device)),unsafe.Pointer(new_pd_head(replacement)),Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRedirected(QPaintDevice const*,QPaintDevice*,QPoint const&)
func (q *QPainter) SetRedirected(device QPaintDeviceInterface,replacement QPaintDeviceInterface,offset *QPoint)  {
	q.Drv(86000,86294,unsafe.Pointer(new_pd_head(device)),unsafe.Pointer(new_pd_head(replacement)),Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRenderHint(QPainter::RenderHint)
func (q *QPainter) SetRenderHint(hint QPainter_RenderHint)  {
	q.Drv(86000,86295,unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRenderHint(QPainter::RenderHint,bool)
func (q *QPainter) SetRenderHintWithHintOn(hint QPainter_RenderHint,on bool)  {
	q.Drv(86000,86296,unsafe.Pointer(&hint),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRenderHints(QFlags<QPainter::RenderHint>)
func (q *QPainter) SetRenderHints(hints QPainter_RenderHint)  {
	q.Drv(86000,86297,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setRenderHints(QFlags<QPainter::RenderHint>,bool)
func (q *QPainter) SetRenderHintsWithHintsOn(hints QPainter_RenderHint,on bool)  {
	q.Drv(86000,86298,unsafe.Pointer(&hints),unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setTransform(QTransform const&)
func (q *QPainter) SetTransform(transform *QTransform)  {
	q.Drv(86000,86299,Native(transform),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setTransform(QTransform const&,bool)
func (q *QPainter) SetTransformWithTransformCombine(transform *QTransform,combine bool)  {
	q.Drv(86000,86300,Native(transform),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setViewTransformEnabled(bool)
func (q *QPainter) SetViewTransformEnabled(enable bool)  {
	q.Drv(86000,86301,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setViewport(QRect const&)
func (q *QPainter) SetViewport(viewport *QRect)  {
	q.Drv(86000,86302,Native(viewport),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setViewport(int,int,int,int)
func (q *QPainter) SetViewportWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(86000,86303,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWindow(QRect const&)
func (q *QPainter) SetWindow(window *QRect)  {
	q.Drv(86000,86304,Native(window),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWindow(int,int,int,int)
func (q *QPainter) SetWindowWithXYWidthHeight(x int,y int,w int,h int)  {
	q.Drv(86000,86305,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWorldMatrix(QMatrix const&)
func (q *QPainter) SetWorldMatrix(matrix *QMatrix)  {
	q.Drv(86000,86306,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWorldMatrix(QMatrix const&,bool)
func (q *QPainter) SetWorldMatrixWithMatrixCombine(matrix *QMatrix,combine bool)  {
	q.Drv(86000,86307,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWorldMatrixEnabled(bool)
func (q *QPainter) SetWorldMatrixEnabled(enabled bool)  {
	q.Drv(86000,86308,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWorldTransform(QTransform const&)
func (q *QPainter) SetWorldTransform(matrix *QTransform)  {
	q.Drv(86000,86309,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::setWorldTransform(QTransform const&,bool)
func (q *QPainter) SetWorldTransformWithTransformCombine(matrix *QTransform,combine bool)  {
	q.Drv(86000,86310,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::shear(double,double)
func (q *QPainter) Shear(sh float64,sv float64)  {
	q.Drv(86000,86311,unsafe.Pointer(&sh),unsafe.Pointer(&sv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::strokePath(QPainterPath const&,QPen const&)
func (q *QPainter) StrokePath(path *QPainterPath,pen *QPen)  {
	q.Drv(86000,86312,Native(path),Native(pen),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::testRenderHint(QPainter::RenderHint)
func (q *QPainter) TestRenderHint(hint QPainter_RenderHint) bool {
	var __rv bool
	q.Drv(86000,86313,unsafe.Pointer(&hint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::transform()
func (q *QPainter) Transform() *QTransform {
	var __rv uintptr
	q.Drv(86000,86314,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QPainter::translate(QPoint const&)
func (q *QPainter) Translate(offset *QPoint)  {
	q.Drv(86000,86315,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::translate(QPointF const&)
func (q *QPainter) TranslateFWithOffset(offset *QPointF)  {
	q.Drv(86000,86316,Native(offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::translate(double,double)
func (q *QPainter) TranslateFWithDxDy(dx float64,dy float64)  {
	q.Drv(86000,86317,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPainter::viewTransformEnabled()
func (q *QPainter) ViewTransformEnabled() bool {
	var __rv bool
	q.Drv(86000,86318,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::viewport()
func (q *QPainter) Viewport() *QRect {
	var __rv uintptr
	q.Drv(86000,86319,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPainter::window()
func (q *QPainter) Window() *QRect {
	var __rv uintptr
	q.Drv(86000,86320,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPainter::worldMatrix()
func (q *QPainter) WorldMatrix() *QMatrix {
	var __rv uintptr
	q.Drv(86000,86321,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QPainter::worldMatrixEnabled()
func (q *QPainter) WorldMatrixEnabled() bool {
	var __rv bool
	q.Drv(86000,86322,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPainter::worldTransform()
func (q *QPainter) WorldTransform() *QTransform {
	var __rv uintptr
	q.Drv(86000,86323,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
