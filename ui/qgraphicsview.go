// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsView_OptimizationFlag - QGraphicsView::OptimizationFlag
type QGraphicsView_OptimizationFlag uint32
const (
	QGraphicsView_DontClipPainter QGraphicsView_OptimizationFlag = 0x1
	QGraphicsView_DontSavePainterState QGraphicsView_OptimizationFlag = 0x2
	QGraphicsView_DontAdjustForAntialiasing QGraphicsView_OptimizationFlag = 0x4
	QGraphicsView_IndirectPainting QGraphicsView_OptimizationFlag = 0x8
)
//enum QGraphicsView_ViewportAnchor - QGraphicsView::ViewportAnchor
type QGraphicsView_ViewportAnchor uint32
const (
	QGraphicsView_NoAnchor QGraphicsView_ViewportAnchor = 0
	QGraphicsView_AnchorViewCenter QGraphicsView_ViewportAnchor = 1
	QGraphicsView_AnchorUnderMouse QGraphicsView_ViewportAnchor = 2
)
//enum QGraphicsView_ViewportUpdateMode - QGraphicsView::ViewportUpdateMode
type QGraphicsView_ViewportUpdateMode uint32
const (
	QGraphicsView_FullViewportUpdate QGraphicsView_ViewportUpdateMode = 0
	QGraphicsView_MinimalViewportUpdate QGraphicsView_ViewportUpdateMode = 1
	QGraphicsView_SmartViewportUpdate QGraphicsView_ViewportUpdateMode = 2
	QGraphicsView_NoViewportUpdate QGraphicsView_ViewportUpdateMode = 3
	QGraphicsView_BoundingRectViewportUpdate QGraphicsView_ViewportUpdateMode = 4
)
//enum QGraphicsView_CacheModeFlag - QGraphicsView::CacheModeFlag
type QGraphicsView_CacheModeFlag uint32
const (
	QGraphicsView_CacheNone QGraphicsView_CacheModeFlag = 0x0
	QGraphicsView_CacheBackground QGraphicsView_CacheModeFlag = 0x1
)
//enum QGraphicsView_DragMode - QGraphicsView::DragMode
type QGraphicsView_DragMode uint32
const (
	QGraphicsView_NoDrag QGraphicsView_DragMode = 0
	QGraphicsView_ScrollHandDrag QGraphicsView_DragMode = 1
	QGraphicsView_RubberBandDrag QGraphicsView_DragMode = 2
)
//struct QGraphicsView : QGraphicsView
type QGraphicsView struct {
	QAbstractScrollArea
}
//QGraphicsView::QGraphicsView()	
func NewGraphicsView() *QGraphicsView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),285000,285102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsView{}
	_p.SetDriver(__rv,285000,false)
	return _p
} 
//QGraphicsView::QGraphicsView(QWidget*)	
func NewGraphicsViewWithParent(parent QWidgetInterface) *QGraphicsView {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),285000,285103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsView{}
	_p.SetDriver(__rv,285000,false)
	return _p
} 
//QGraphicsView::alignment()
func (q *QGraphicsView) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(285000,285104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::backgroundBrush()
func (q *QGraphicsView) BackgroundBrush() *QBrush {
	var __rv uintptr
	q.Drv(285000,285105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QGraphicsView::cacheMode()
func (q *QGraphicsView) CacheMode() QGraphicsView_CacheModeFlag {
	var __rv QGraphicsView_CacheModeFlag
	q.Drv(285000,285106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::centerOn(QGraphicsItem const*)
func (q *QGraphicsView) CenterOn(item *QGraphicsItem)  {
	q.Drv(285000,285107,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::centerOn(QPointF const&)
func (q *QGraphicsView) CenterOnFWithPos(pos *QPointF)  {
	q.Drv(285000,285108,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::centerOn(double,double)
func (q *QGraphicsView) CenterOnFWithXY(x float64,y float64)  {
	q.Drv(285000,285109,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::contextMenuEvent(QContextMenuEvent*)
func (q *QGraphicsView) ContextMenuEvent(event *QContextMenuEvent)  {
	q.Drv(285000,285110,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::dragEnterEvent(QDragEnterEvent*)
func (q *QGraphicsView) DragEnterEvent(event *QDragEnterEvent)  {
	q.Drv(285000,285111,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::dragLeaveEvent(QDragLeaveEvent*)
func (q *QGraphicsView) DragLeaveEvent(event *QDragLeaveEvent)  {
	q.Drv(285000,285112,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::dragMode()
func (q *QGraphicsView) DragMode() QGraphicsView_DragMode {
	var __rv QGraphicsView_DragMode
	q.Drv(285000,285113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::dragMoveEvent(QDragMoveEvent*)
func (q *QGraphicsView) DragMoveEvent(event *QDragMoveEvent)  {
	q.Drv(285000,285114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::drawBackground(QPainter*,QRectF const&)
func (q *QGraphicsView) DrawBackground(painter *QPainter,rect *QRectF)  {
	q.Drv(285000,285115,Native(painter),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::drawForeground(QPainter*,QRectF const&)
func (q *QGraphicsView) DrawForeground(painter *QPainter,rect *QRectF)  {
	q.Drv(285000,285116,Native(painter),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::dropEvent(QDropEvent*)
func (q *QGraphicsView) DropEvent(event *QDropEvent)  {
	q.Drv(285000,285117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::ensureVisible(QGraphicsItem const*)
func (q *QGraphicsView) EnsureVisible(item *QGraphicsItem)  {
	q.Drv(285000,285118,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::ensureVisible(QRectF const&)
func (q *QGraphicsView) EnsureVisibleFWithRect(rect *QRectF)  {
	q.Drv(285000,285119,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::ensureVisible(QGraphicsItem const*,int,int)
func (q *QGraphicsView) EnsureVisibleWithItemXmarginYmargin(item *QGraphicsItem,xmargin int,ymargin int)  {
	q.Drv(285000,285120,Native(item),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::ensureVisible(QRectF const&,int,int)
func (q *QGraphicsView) EnsureVisibleFWithRectXmarginYmargin(rect *QRectF,xmargin int,ymargin int)  {
	q.Drv(285000,285121,Native(rect),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::ensureVisible(double,double,double,double,int,int)
func (q *QGraphicsView) EnsureVisibleFWithXYWidthHeightXmarginYmargin(x float64,y float64,w float64,h float64,xmargin int,ymargin int)  {
	q.Drv(285000,285122,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::event(QEvent*)
func (q *QGraphicsView) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(285000,285123,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::fitInView(QGraphicsItem const*)
func (q *QGraphicsView) FitInView(item *QGraphicsItem)  {
	q.Drv(285000,285124,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::fitInView(QRectF const&)
func (q *QGraphicsView) FitInViewFWithRect(rect *QRectF)  {
	q.Drv(285000,285125,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::fitInView(QGraphicsItem const*,Qt::AspectRatioMode)
func (q *QGraphicsView) FitInViewWithItemAspectradiomode(item *QGraphicsItem,aspectRadioMode Qt_AspectRatioMode)  {
	q.Drv(285000,285126,Native(item),unsafe.Pointer(&aspectRadioMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::fitInView(QRectF const&,Qt::AspectRatioMode)
func (q *QGraphicsView) FitInViewFWithRectAspectradiomode(rect *QRectF,aspectRadioMode Qt_AspectRatioMode)  {
	q.Drv(285000,285127,Native(rect),unsafe.Pointer(&aspectRadioMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::fitInView(double,double,double,double,Qt::AspectRatioMode)
func (q *QGraphicsView) FitInViewFWithXYWidthHeightAspectradiomode(x float64,y float64,w float64,h float64,aspectRadioMode Qt_AspectRatioMode)  {
	q.Drv(285000,285128,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&aspectRadioMode),nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::focusInEvent(QFocusEvent*)
func (q *QGraphicsView) FocusInEvent(event *QFocusEvent)  {
	q.Drv(285000,285129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::focusNextPrevChild(bool)
func (q *QGraphicsView) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(285000,285130,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::focusOutEvent(QFocusEvent*)
func (q *QGraphicsView) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(285000,285131,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::foregroundBrush()
func (q *QGraphicsView) ForegroundBrush() *QBrush {
	var __rv uintptr
	q.Drv(285000,285132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QGraphicsView::inputMethodEvent(QInputMethodEvent*)
func (q *QGraphicsView) InputMethodEvent(event *QInputMethodEvent)  {
	q.Drv(285000,285133,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::inputMethodQuery(Qt::InputMethodQuery)
func (q *QGraphicsView) InputMethodQuery(query Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(285000,285134,unsafe.Pointer(&query),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsView::invalidateScene()
func (q *QGraphicsView) InvalidateScene()  {
	q.Drv(285000,285135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::invalidateScene(QRectF const&,QFlags<QGraphicsScene::SceneLayer>)
func (q *QGraphicsView) InvalidateSceneFWithRectLayers(rect *QRectF,layers QGraphicsScene_SceneLayer)  {
	q.Drv(285000,285136,Native(rect),unsafe.Pointer(&layers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::isInteractive()
func (q *QGraphicsView) IsInteractive() bool {
	var __rv bool
	q.Drv(285000,285137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::isTransformed()
func (q *QGraphicsView) IsTransformed() bool {
	var __rv bool
	q.Drv(285000,285138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::itemAt(QPoint const&)
func (q *QGraphicsView) ItemAt(pos *QPoint) *QGraphicsItem {
	var __rv uintptr
	q.Drv(285000,285139,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsView::itemAt(int,int)
func (q *QGraphicsView) ItemAtWithXY(x int,y int) *QGraphicsItem {
	var __rv uintptr
	q.Drv(285000,285140,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsView::items()
func (q *QGraphicsView) Items() []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QPainterPath const&)
func (q *QGraphicsView) ItemsWithPath(path *QPainterPath) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285142,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QPoint const&)
func (q *QGraphicsView) ItemsWithPos(pos *QPoint) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285143,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QPolygon const&)
func (q *QGraphicsView) ItemsWithPolygon(polygon *QPolygon) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285144,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QRect const&)
func (q *QGraphicsView) ItemsWithRect(rect *QRect) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285145,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QPainterPath const&,Qt::ItemSelectionMode)
func (q *QGraphicsView) ItemsWithPathMode(path *QPainterPath,mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285146,Native(path),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QPolygon const&,Qt::ItemSelectionMode)
func (q *QGraphicsView) ItemsWithPolygonMode(polygon *QPolygon,mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285147,Native(polygon),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(QRect const&,Qt::ItemSelectionMode)
func (q *QGraphicsView) ItemsWithRectMode(rect *QRect,mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285148,Native(rect),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(int,int)
func (q *QGraphicsView) ItemsWithXY(x int,y int) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285149,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::items(int,int,int,int,Qt::ItemSelectionMode)
func (q *QGraphicsView) ItemsWithXYWidthHeightMode(x int,y int,w int,h int,mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(285000,285150,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::keyPressEvent(QKeyEvent*)
func (q *QGraphicsView) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(285000,285151,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::keyReleaseEvent(QKeyEvent*)
func (q *QGraphicsView) KeyReleaseEvent(event *QKeyEvent)  {
	q.Drv(285000,285152,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::mapFromScene(QPainterPath const&)
func (q *QGraphicsView) MapFromScene(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(285000,285153,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsView::mapFromScene(QPointF const&)
func (q *QGraphicsView) MapFromSceneFWithPoint(point *QPointF) *QPoint {
	var __rv uintptr
	q.Drv(285000,285154,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsView::mapFromScene(QPolygonF const&)
func (q *QGraphicsView) MapFromSceneFWithPolygon(polygon *QPolygonF) *QPolygon {
	var __rv uintptr
	q.Drv(285000,285155,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QGraphicsView::mapFromScene(QRectF const&)
func (q *QGraphicsView) MapFromSceneFWithRect(rect *QRectF) *QPolygon {
	var __rv uintptr
	q.Drv(285000,285156,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QGraphicsView::mapFromScene(double,double)
func (q *QGraphicsView) MapFromSceneFWithXY(x float64,y float64) *QPoint {
	var __rv uintptr
	q.Drv(285000,285157,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QGraphicsView::mapFromScene(double,double,double,double)
func (q *QGraphicsView) MapFromSceneFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QPolygon {
	var __rv uintptr
	q.Drv(285000,285158,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygon{}
	_rp.SetDriver(__rv,101000,true)
	return _rp
}	
//QGraphicsView::mapToScene(QPainterPath const&)
func (q *QGraphicsView) MapToScene(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(285000,285159,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsView::mapToScene(QPoint const&)
func (q *QGraphicsView) MapToSceneWithPoint(point *QPoint) *QPointF {
	var __rv uintptr
	q.Drv(285000,285160,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsView::mapToScene(QPolygon const&)
func (q *QGraphicsView) MapToSceneWithPolygon(polygon *QPolygon) *QPolygonF {
	var __rv uintptr
	q.Drv(285000,285161,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsView::mapToScene(QRect const&)
func (q *QGraphicsView) MapToSceneWithRect(rect *QRect) *QPolygonF {
	var __rv uintptr
	q.Drv(285000,285162,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsView::mapToScene(int,int)
func (q *QGraphicsView) MapToSceneWithXY(x int,y int) *QPointF {
	var __rv uintptr
	q.Drv(285000,285163,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsView::mapToScene(int,int,int,int)
func (q *QGraphicsView) MapToSceneWithXYWidthHeight(x int,y int,w int,h int) *QPolygonF {
	var __rv uintptr
	q.Drv(285000,285164,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsView::matrix()
func (q *QGraphicsView) Matrix() *QMatrix {
	var __rv uintptr
	q.Drv(285000,285165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QGraphicsView::mouseDoubleClickEvent(QMouseEvent*)
func (q *QGraphicsView) MouseDoubleClickEvent(event *QMouseEvent)  {
	q.Drv(285000,285166,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::mouseMoveEvent(QMouseEvent*)
func (q *QGraphicsView) MouseMoveEvent(event *QMouseEvent)  {
	q.Drv(285000,285167,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::mousePressEvent(QMouseEvent*)
func (q *QGraphicsView) MousePressEvent(event *QMouseEvent)  {
	q.Drv(285000,285168,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::mouseReleaseEvent(QMouseEvent*)
func (q *QGraphicsView) MouseReleaseEvent(event *QMouseEvent)  {
	q.Drv(285000,285169,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::optimizationFlags()
func (q *QGraphicsView) OptimizationFlags() QGraphicsView_OptimizationFlag {
	var __rv QGraphicsView_OptimizationFlag
	q.Drv(285000,285170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::paintEvent(QPaintEvent*)
func (q *QGraphicsView) PaintEvent(event *QPaintEvent)  {
	q.Drv(285000,285171,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::render(QPainter*)
func (q *QGraphicsView) Render(painter *QPainter)  {
	q.Drv(285000,285172,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::render(QPainter*,QRectF const&,QRect const&,Qt::AspectRatioMode)
func (q *QGraphicsView) RenderWithPainterTargetSourceAspectratiomode(painter *QPainter,target *QRectF,source *QRect,aspectRatioMode Qt_AspectRatioMode)  {
	q.Drv(285000,285173,Native(painter),Native(target),Native(source),unsafe.Pointer(&aspectRatioMode),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::renderHints()
func (q *QGraphicsView) RenderHints() QPainter_RenderHint {
	var __rv QPainter_RenderHint
	q.Drv(285000,285174,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::resetCachedContent()
func (q *QGraphicsView) ResetCachedContent()  {
	q.Drv(285000,285175,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::resetMatrix()
func (q *QGraphicsView) ResetMatrix()  {
	q.Drv(285000,285176,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::resetTransform()
func (q *QGraphicsView) ResetTransform()  {
	q.Drv(285000,285177,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::resizeAnchor()
func (q *QGraphicsView) ResizeAnchor() QGraphicsView_ViewportAnchor {
	var __rv QGraphicsView_ViewportAnchor
	q.Drv(285000,285178,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::resizeEvent(QResizeEvent*)
func (q *QGraphicsView) ResizeEvent(event *QResizeEvent)  {
	q.Drv(285000,285179,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::rotate(double)
func (q *QGraphicsView) Rotate(angle float64)  {
	q.Drv(285000,285180,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::rubberBandSelectionMode()
func (q *QGraphicsView) RubberBandSelectionMode() Qt_ItemSelectionMode {
	var __rv Qt_ItemSelectionMode
	q.Drv(285000,285181,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::scale(double,double)
func (q *QGraphicsView) Scale(sx float64,sy float64)  {
	q.Drv(285000,285182,unsafe.Pointer(&sx),unsafe.Pointer(&sy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::scene()
func (q *QGraphicsView) Scene() *QGraphicsScene {
	var __rv uintptr
	q.Drv(285000,285183,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsScene{}
	_rp.SetDriver(__rv,272000,false)
	return _rp
}	
//QGraphicsView::sceneRect()
func (q *QGraphicsView) SceneRect() *QRectF {
	var __rv uintptr
	q.Drv(285000,285184,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsView::scrollContentsBy(int,int)
func (q *QGraphicsView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(285000,285185,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QGraphicsView) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(285000,285186,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setBackgroundBrush(QBrush const&)
func (q *QGraphicsView) SetBackgroundBrush(brush *QBrush)  {
	q.Drv(285000,285187,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setCacheMode(QFlags<QGraphicsView::CacheModeFlag>)
func (q *QGraphicsView) SetCacheMode(mode QGraphicsView_CacheModeFlag)  {
	q.Drv(285000,285188,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setDragMode(QGraphicsView::DragMode)
func (q *QGraphicsView) SetDragMode(mode QGraphicsView_DragMode)  {
	q.Drv(285000,285189,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setForegroundBrush(QBrush const&)
func (q *QGraphicsView) SetForegroundBrush(brush *QBrush)  {
	q.Drv(285000,285190,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setInteractive(bool)
func (q *QGraphicsView) SetInteractive(allowed bool)  {
	q.Drv(285000,285191,unsafe.Pointer(&allowed),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setMatrix(QMatrix const&)
func (q *QGraphicsView) SetMatrix(matrix *QMatrix)  {
	q.Drv(285000,285192,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setMatrix(QMatrix const&,bool)
func (q *QGraphicsView) SetMatrixWithMatrixCombine(matrix *QMatrix,combine bool)  {
	q.Drv(285000,285193,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setOptimizationFlag(QGraphicsView::OptimizationFlag)
func (q *QGraphicsView) SetOptimizationFlag(flag QGraphicsView_OptimizationFlag)  {
	q.Drv(285000,285194,unsafe.Pointer(&flag),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setOptimizationFlag(QGraphicsView::OptimizationFlag,bool)
func (q *QGraphicsView) SetOptimizationFlagWithFlagEnabled(flag QGraphicsView_OptimizationFlag,enabled bool)  {
	q.Drv(285000,285195,unsafe.Pointer(&flag),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setOptimizationFlags(QFlags<QGraphicsView::OptimizationFlag>)
func (q *QGraphicsView) SetOptimizationFlags(flags QGraphicsView_OptimizationFlag)  {
	q.Drv(285000,285196,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setRenderHint(QPainter::RenderHint)
func (q *QGraphicsView) SetRenderHint(hint QPainter_RenderHint)  {
	q.Drv(285000,285197,unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setRenderHint(QPainter::RenderHint,bool)
func (q *QGraphicsView) SetRenderHintWithHintEnabled(hint QPainter_RenderHint,enabled bool)  {
	q.Drv(285000,285198,unsafe.Pointer(&hint),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setRenderHints(QFlags<QPainter::RenderHint>)
func (q *QGraphicsView) SetRenderHints(hints QPainter_RenderHint)  {
	q.Drv(285000,285199,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setResizeAnchor(QGraphicsView::ViewportAnchor)
func (q *QGraphicsView) SetResizeAnchor(anchor QGraphicsView_ViewportAnchor)  {
	q.Drv(285000,285200,unsafe.Pointer(&anchor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setRubberBandSelectionMode(Qt::ItemSelectionMode)
func (q *QGraphicsView) SetRubberBandSelectionMode(mode Qt_ItemSelectionMode)  {
	q.Drv(285000,285201,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setScene(QGraphicsScene*)
func (q *QGraphicsView) SetScene(scene *QGraphicsScene)  {
	q.Drv(285000,285202,Native(scene),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setSceneRect(QRectF const&)
func (q *QGraphicsView) SetSceneRect(rect *QRectF)  {
	q.Drv(285000,285203,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setSceneRect(double,double,double,double)
func (q *QGraphicsView) SetSceneRectFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(285000,285204,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setTransform(QTransform const&)
func (q *QGraphicsView) SetTransform(matrix *QTransform)  {
	q.Drv(285000,285205,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setTransform(QTransform const&,bool)
func (q *QGraphicsView) SetTransformWithTransformCombine(matrix *QTransform,combine bool)  {
	q.Drv(285000,285206,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setTransformationAnchor(QGraphicsView::ViewportAnchor)
func (q *QGraphicsView) SetTransformationAnchor(anchor QGraphicsView_ViewportAnchor)  {
	q.Drv(285000,285207,unsafe.Pointer(&anchor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setViewportUpdateMode(QGraphicsView::ViewportUpdateMode)
func (q *QGraphicsView) SetViewportUpdateMode(mode QGraphicsView_ViewportUpdateMode)  {
	q.Drv(285000,285208,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::setupViewport(QWidget*)
func (q *QGraphicsView) SetupViewport(widget QWidgetInterface)  {
	q.Drv(285000,285209,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::shear(double,double)
func (q *QGraphicsView) Shear(sh float64,sv float64)  {
	q.Drv(285000,285210,unsafe.Pointer(&sh),unsafe.Pointer(&sv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::showEvent(QShowEvent*)
func (q *QGraphicsView) ShowEvent(event *QShowEvent)  {
	q.Drv(285000,285211,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::sizeHint()
func (q *QGraphicsView) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(285000,285212,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QGraphicsView::transform()
func (q *QGraphicsView) Transform() *QTransform {
	var __rv uintptr
	q.Drv(285000,285213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsView::transformationAnchor()
func (q *QGraphicsView) TransformationAnchor() QGraphicsView_ViewportAnchor {
	var __rv QGraphicsView_ViewportAnchor
	q.Drv(285000,285214,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::translate(double,double)
func (q *QGraphicsView) Translate(dx float64,dy float64)  {
	q.Drv(285000,285215,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::updateScene(QList<QRectF> const&)
func (q *QGraphicsView) UpdateScene(rects []*QRectF)  {
	q.Drv(285000,285216,unsafe.Pointer(&rects),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::updateSceneRect(QRectF const&)
func (q *QGraphicsView) UpdateSceneRect(rect *QRectF)  {
	q.Drv(285000,285217,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsView::viewportEvent(QEvent*)
func (q *QGraphicsView) ViewportEvent(event *QEvent) bool {
	var __rv bool
	q.Drv(285000,285218,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::viewportTransform()
func (q *QGraphicsView) ViewportTransform() *QTransform {
	var __rv uintptr
	q.Drv(285000,285219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsView::viewportUpdateMode()
func (q *QGraphicsView) ViewportUpdateMode() QGraphicsView_ViewportUpdateMode {
	var __rv QGraphicsView_ViewportUpdateMode
	q.Drv(285000,285220,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsView::wheelEvent(QWheelEvent*)
func (q *QGraphicsView) WheelEvent(event *QWheelEvent)  {
	q.Drv(285000,285221,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
