// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPaintEngine_PolygonDrawMode - QPaintEngine::PolygonDrawMode
type QPaintEngine_PolygonDrawMode uint32
const (
	QPaintEngine_OddEvenMode QPaintEngine_PolygonDrawMode = 0
	QPaintEngine_WindingMode QPaintEngine_PolygonDrawMode = 1
	QPaintEngine_ConvexMode QPaintEngine_PolygonDrawMode = 2
	QPaintEngine_PolylineMode QPaintEngine_PolygonDrawMode = 3
)
//enum QPaintEngine_Type - QPaintEngine::Type
type QPaintEngine_Type uint32
const (
	QPaintEngine_X11 QPaintEngine_Type = 0
	QPaintEngine_Windows QPaintEngine_Type = 1
	QPaintEngine_QuickDraw QPaintEngine_Type = 2
	QPaintEngine_CoreGraphics QPaintEngine_Type = 3
	QPaintEngine_MacPrinter QPaintEngine_Type = 4
	QPaintEngine_QWindowSystem QPaintEngine_Type = 5
	QPaintEngine_PostScript QPaintEngine_Type = 6
	QPaintEngine_OpenGL QPaintEngine_Type = 7
	QPaintEngine_Picture QPaintEngine_Type = 8
	QPaintEngine_SVG QPaintEngine_Type = 9
	QPaintEngine_Raster QPaintEngine_Type = 10
	QPaintEngine_Direct3D QPaintEngine_Type = 11
	QPaintEngine_Pdf QPaintEngine_Type = 12
	QPaintEngine_OpenVG QPaintEngine_Type = 13
	QPaintEngine_OpenGL2 QPaintEngine_Type = 14
	QPaintEngine_PaintBuffer QPaintEngine_Type = 15
	QPaintEngine_User QPaintEngine_Type = 50
	QPaintEngine_MaxUser QPaintEngine_Type = 100
)
//enum QPaintEngine_PaintEngineFeature - QPaintEngine::PaintEngineFeature
type QPaintEngine_PaintEngineFeature uint32
const (
	QPaintEngine_PrimitiveTransform QPaintEngine_PaintEngineFeature = 0x00000001
	QPaintEngine_PatternTransform QPaintEngine_PaintEngineFeature = 0x00000002
	QPaintEngine_PixmapTransform QPaintEngine_PaintEngineFeature = 0x00000004
	QPaintEngine_PatternBrush QPaintEngine_PaintEngineFeature = 0x00000008
	QPaintEngine_LinearGradientFill QPaintEngine_PaintEngineFeature = 0x00000010
	QPaintEngine_RadialGradientFill QPaintEngine_PaintEngineFeature = 0x00000020
	QPaintEngine_ConicalGradientFill QPaintEngine_PaintEngineFeature = 0x00000040
	QPaintEngine_AlphaBlend QPaintEngine_PaintEngineFeature = 0x00000080
	QPaintEngine_PorterDuff QPaintEngine_PaintEngineFeature = 0x00000100
	QPaintEngine_PainterPaths QPaintEngine_PaintEngineFeature = 0x00000200
	QPaintEngine_Antialiasing QPaintEngine_PaintEngineFeature = 0x00000400
	QPaintEngine_BrushStroke QPaintEngine_PaintEngineFeature = 0x00000800
	QPaintEngine_ConstantOpacity QPaintEngine_PaintEngineFeature = 0x00001000
	QPaintEngine_MaskedBrush QPaintEngine_PaintEngineFeature = 0x00002000
	QPaintEngine_PerspectiveTransform QPaintEngine_PaintEngineFeature = 0x00004000
	QPaintEngine_BlendModes QPaintEngine_PaintEngineFeature = 0x00008000
	QPaintEngine_ObjectBoundingModeGradients QPaintEngine_PaintEngineFeature = 0x00010000
	QPaintEngine_RasterOpModes QPaintEngine_PaintEngineFeature = 0x00020000
	QPaintEngine_PaintOutsidePaintEvent QPaintEngine_PaintEngineFeature = 0x20000000
	QPaintEngine_AllFeatures QPaintEngine_PaintEngineFeature = 0xffffffff
)
//enum QPaintEngine_DirtyFlag - QPaintEngine::DirtyFlag
type QPaintEngine_DirtyFlag uint32
const (
	QPaintEngine_DirtyPen QPaintEngine_DirtyFlag = 0x0001
	QPaintEngine_DirtyBrush QPaintEngine_DirtyFlag = 0x0002
	QPaintEngine_DirtyBrushOrigin QPaintEngine_DirtyFlag = 0x0004
	QPaintEngine_DirtyFont QPaintEngine_DirtyFlag = 0x0008
	QPaintEngine_DirtyBackground QPaintEngine_DirtyFlag = 0x0010
	QPaintEngine_DirtyBackgroundMode QPaintEngine_DirtyFlag = 0x0020
	QPaintEngine_DirtyTransform QPaintEngine_DirtyFlag = 0x0040
	QPaintEngine_DirtyClipRegion QPaintEngine_DirtyFlag = 0x0080
	QPaintEngine_DirtyClipPath QPaintEngine_DirtyFlag = 0x0100
	QPaintEngine_DirtyHints QPaintEngine_DirtyFlag = 0x0200
	QPaintEngine_DirtyCompositionMode QPaintEngine_DirtyFlag = 0x0400
	QPaintEngine_DirtyClipEnabled QPaintEngine_DirtyFlag = 0x0800
	QPaintEngine_DirtyOpacity QPaintEngine_DirtyFlag = 0x1000
	QPaintEngine_AllDirty QPaintEngine_DirtyFlag = 0xffff
)
//struct QPaintEngine : QPaintEngine
type QPaintEngine struct {
	BaseDrv
}
//QPaintEngine::begin(QPaintDevice*)
func (q *QPaintEngine) Begin(pdev QPaintDeviceInterface) bool {
	var __rv bool
	q.Drv(83000,83102,unsafe.Pointer(new_pd_head(pdev)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::clearDirty(QFlags<QPaintEngine::DirtyFlag>)
func (q *QPaintEngine) ClearDirty(df QPaintEngine_DirtyFlag)  {
	q.Drv(83000,83103,unsafe.Pointer(&df),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::coordinateOffset()
func (q *QPaintEngine) CoordinateOffset() *QPoint {
	var __rv uintptr
	q.Drv(83000,83104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QPaintEngine::drawEllipse(QRect const&)
func (q *QPaintEngine) DrawEllipse(r *QRect)  {
	q.Drv(83000,83105,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawEllipse(QRectF const&)
func (q *QPaintEngine) DrawEllipseFWithRectf(r *QRectF)  {
	q.Drv(83000,83106,Native(r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawImage(QRectF const&,QImage const&,QRectF const&,QFlags<Qt::ImageConversionFlag>)
func (q *QPaintEngine) DrawImage(r *QRectF,pm *QImage,sr *QRectF,flags Qt_ImageConversionFlag)  {
	q.Drv(83000,83107,Native(r),Native(pm),Native(sr),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawLines(QLine const*,int)
func (q *QPaintEngine) DrawLinesWithLinesLinecount(lines *QLine,lineCount int)  {
	q.Drv(83000,83108,Native(lines),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawLines(QLineF const*,int)
func (q *QPaintEngine) DrawLinesFWithLinesLinecount(lines *QLineF,lineCount int)  {
	q.Drv(83000,83109,Native(lines),unsafe.Pointer(&lineCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPath(QPainterPath const&)
func (q *QPaintEngine) DrawPath(path *QPainterPath)  {
	q.Drv(83000,83110,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPixmap(QRectF const&,QPixmap const&,QRectF const&)
func (q *QPaintEngine) DrawPixmap(r *QRectF,pm *QPixmap,sr *QRectF)  {
	q.Drv(83000,83111,Native(r),Native(pm),Native(sr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPoints(QPoint const*,int)
func (q *QPaintEngine) DrawPointsWithPointsPointcount(points *QPoint,pointCount int)  {
	q.Drv(83000,83112,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPoints(QPointF const*,int)
func (q *QPaintEngine) DrawPointsFWithPointsPointcount(points *QPointF,pointCount int)  {
	q.Drv(83000,83113,Native(points),unsafe.Pointer(&pointCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPolygon(QPoint const*,int,QPaintEngine::PolygonDrawMode)
func (q *QPaintEngine) DrawPolygonWithPointsPointcountMode(points *QPoint,pointCount int,mode QPaintEngine_PolygonDrawMode)  {
	q.Drv(83000,83114,Native(points),unsafe.Pointer(&pointCount),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawPolygon(QPointF const*,int,QPaintEngine::PolygonDrawMode)
func (q *QPaintEngine) DrawPolygonFWithPointsPointcountMode(points *QPointF,pointCount int,mode QPaintEngine_PolygonDrawMode)  {
	q.Drv(83000,83115,Native(points),unsafe.Pointer(&pointCount),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawRects(QRect const*,int)
func (q *QPaintEngine) DrawRectsWithRectsRectcount(rects *QRect,rectCount int)  {
	q.Drv(83000,83116,Native(rects),unsafe.Pointer(&rectCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawRects(QRectF const*,int)
func (q *QPaintEngine) DrawRectsFWithRectsRectcount(rects *QRectF,rectCount int)  {
	q.Drv(83000,83117,Native(rects),unsafe.Pointer(&rectCount),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawTextItem(QPointF const&,QTextItem const&)
func (q *QPaintEngine) DrawTextItem(p *QPointF,textItem *QTextItem)  {
	q.Drv(83000,83118,Native(p),Native(textItem),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::drawTiledPixmap(QRectF const&,QPixmap const&,QPointF const&)
func (q *QPaintEngine) DrawTiledPixmap(r *QRectF,pixmap *QPixmap,s *QPointF)  {
	q.Drv(83000,83119,Native(r),Native(pixmap),Native(s),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::end()
func (q *QPaintEngine) End() bool {
	var __rv bool
	q.Drv(83000,83120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::fix_neg_rect(int*,int*,int*,int*)
func (q *QPaintEngine) Fix_neg_rect(x *int,y *int,w *int,h *int)  {
	q.Drv(83000,83121,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::hasFeature(QFlags<QPaintEngine::PaintEngineFeature>)
func (q *QPaintEngine) HasFeature(feature QPaintEngine_PaintEngineFeature) bool {
	var __rv bool
	q.Drv(83000,83122,unsafe.Pointer(&feature),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::isActive()
func (q *QPaintEngine) IsActive() bool {
	var __rv bool
	q.Drv(83000,83123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::isExtended()
func (q *QPaintEngine) IsExtended() bool {
	var __rv bool
	q.Drv(83000,83124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::paintDevice()
func (q *QPaintEngine) PaintDevice() *QPaintDevice {
	var __rv uintptr
	q.Drv(83000,83125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintDevice{}
	_rp.SetDriver(__rv,82000,true)
	return _rp
}	
//QPaintEngine::painter()
func (q *QPaintEngine) Painter() *QPainter {
	var __rv uintptr
	q.Drv(83000,83126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainter{}
	_rp.SetDriver(__rv,86000,true)
	return _rp
}	
//QPaintEngine::setActive(bool)
func (q *QPaintEngine) SetActive(newState bool)  {
	q.Drv(83000,83127,unsafe.Pointer(&newState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::setDirty(QFlags<QPaintEngine::DirtyFlag>)
func (q *QPaintEngine) SetDirty(df QPaintEngine_DirtyFlag)  {
	q.Drv(83000,83128,unsafe.Pointer(&df),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::setPaintDevice(QPaintDevice*)
func (q *QPaintEngine) SetPaintDevice(device QPaintDeviceInterface)  {
	q.Drv(83000,83129,unsafe.Pointer(new_pd_head(device)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::setSystemClip(QRegion const&)
func (q *QPaintEngine) SetSystemClip(baseClip *QRegion)  {
	q.Drv(83000,83130,Native(baseClip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::setSystemRect(QRect const&)
func (q *QPaintEngine) SetSystemRect(rect *QRect)  {
	q.Drv(83000,83131,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::syncState()
func (q *QPaintEngine) SyncState()  {
	q.Drv(83000,83132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPaintEngine::systemClip()
func (q *QPaintEngine) SystemClip() *QRegion {
	var __rv uintptr
	q.Drv(83000,83133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QPaintEngine::systemRect()
func (q *QPaintEngine) SystemRect() *QRect {
	var __rv uintptr
	q.Drv(83000,83134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPaintEngine::testDirty(QFlags<QPaintEngine::DirtyFlag>)
func (q *QPaintEngine) TestDirty(df QPaintEngine_DirtyFlag) bool {
	var __rv bool
	q.Drv(83000,83135,unsafe.Pointer(&df),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::type()
func (q *QPaintEngine) Type() QPaintEngine_Type {
	var __rv QPaintEngine_Type
	q.Drv(83000,83136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintEngine::updateState(QPaintEngineState const&)
func (q *QPaintEngine) UpdateState(state *QPaintEngineState)  {
	q.Drv(83000,83137,Native(state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
