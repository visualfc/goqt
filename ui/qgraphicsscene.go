// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsScene_SceneLayer - QGraphicsScene::SceneLayer
type QGraphicsScene_SceneLayer uint32
const (
	QGraphicsScene_ItemLayer QGraphicsScene_SceneLayer = 0x1
	QGraphicsScene_BackgroundLayer QGraphicsScene_SceneLayer = 0x2
	QGraphicsScene_ForegroundLayer QGraphicsScene_SceneLayer = 0x4
	QGraphicsScene_AllLayers QGraphicsScene_SceneLayer = 0xffff
)
//enum QGraphicsScene_ItemIndexMethod - QGraphicsScene::ItemIndexMethod
type QGraphicsScene_ItemIndexMethod int32
const (
	QGraphicsScene_BspTreeIndex QGraphicsScene_ItemIndexMethod = 0
	QGraphicsScene_NoIndex QGraphicsScene_ItemIndexMethod = -1
)
//struct QGraphicsScene : QGraphicsScene
type QGraphicsScene struct {
	QObject
}
func (q *QGraphicsScene) OnSceneRectChanged(fn func(*QRectF)) uintptr {
	var __rv uintptr
	q.Drv(272000,272102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScene) OnChanged(fn func([]*QRectF)) uintptr {
	var __rv uintptr
	q.Drv(272000,272103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGraphicsScene) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(272000,272104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGraphicsScene::QGraphicsScene()	
func NewGraphicsScene() *QGraphicsScene {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),272000,272105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScene{}
	_p.SetDriver(__rv,272000,false)
	return _p
} 
//QGraphicsScene::QGraphicsScene(QObject*)	
func NewGraphicsSceneWithParent(parent QObjectInterface) *QGraphicsScene {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),272000,272106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScene{}
	_p.SetDriver(__rv,272000,false)
	return _p
} 
//QGraphicsScene::QGraphicsScene(QRectF const&,QObject*)	
func NewGraphicsSceneWithScenerectParent(sceneRect *QRectF,parent QObjectInterface) *QGraphicsScene {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),272000,272107,Native(sceneRect),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScene{}
	_p.SetDriver(__rv,272000,false)
	return _p
} 
//QGraphicsScene::QGraphicsScene(double,double,double,double,QObject*)	
func NewGraphicsSceneWithXYWidthHeightParent(x float64,y float64,width float64,height float64,parent QObjectInterface) *QGraphicsScene {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),272000,272108,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&width),unsafe.Pointer(&height),Native(parent),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGraphicsScene{}
	_p.SetDriver(__rv,272000,false)
	return _p
} 
//QGraphicsScene::activePanel()
func (q *QGraphicsScene) ActivePanel() *QGraphicsItem {
	var __rv uintptr
	q.Drv(272000,272109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsScene::activeWindow()
func (q *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	var __rv uintptr
	q.Drv(272000,272110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsWidget{}
	_rp.SetDriver(__rv,286000,false)
	return _rp
}	
//QGraphicsScene::addEllipse(QRectF const&)
func (q *QGraphicsScene) AddEllipse(rect *QRectF) *QGraphicsEllipseItem {
	var __rv uintptr
	q.Drv(272000,272111,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsEllipseItem{}
	_rp.SetDriver(__rv,254000,true)
	return _rp
}	
//QGraphicsScene::addEllipse(QRectF const&,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddEllipseFWithRectPenBrush(rect *QRectF,pen *QPen,brush *QBrush) *QGraphicsEllipseItem {
	var __rv uintptr
	q.Drv(272000,272112,Native(rect),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsEllipseItem{}
	_rp.SetDriver(__rv,254000,true)
	return _rp
}	
//QGraphicsScene::addEllipse(double,double,double,double,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddEllipseFWithXYWidthHeightPenBrush(x float64,y float64,w float64,h float64,pen *QPen,brush *QBrush) *QGraphicsEllipseItem {
	var __rv uintptr
	q.Drv(272000,272113,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsEllipseItem{}
	_rp.SetDriver(__rv,254000,true)
	return _rp
}	
//QGraphicsScene::addItem(QGraphicsItem*)
func (q *QGraphicsScene) AddItem(item *QGraphicsItem)  {
	q.Drv(272000,272114,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::addLine(QLineF const&)
func (q *QGraphicsScene) AddLine(line *QLineF) *QGraphicsLineItem {
	var __rv uintptr
	q.Drv(272000,272115,Native(line),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLineItem{}
	_rp.SetDriver(__rv,261000,true)
	return _rp
}	
//QGraphicsScene::addLine(QLineF const&,QPen const&)
func (q *QGraphicsScene) AddLineFWithLinePen(line *QLineF,pen *QPen) *QGraphicsLineItem {
	var __rv uintptr
	q.Drv(272000,272116,Native(line),Native(pen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLineItem{}
	_rp.SetDriver(__rv,261000,true)
	return _rp
}	
//QGraphicsScene::addLine(double,double,double,double,QPen const&)
func (q *QGraphicsScene) AddLineFWithX1Y1X2Y2Pen(x1 float64,y1 float64,x2 float64,y2 float64,pen *QPen) *QGraphicsLineItem {
	var __rv uintptr
	q.Drv(272000,272117,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),Native(pen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsLineItem{}
	_rp.SetDriver(__rv,261000,true)
	return _rp
}	
//QGraphicsScene::addPath(QPainterPath const&)
func (q *QGraphicsScene) AddPath(path *QPainterPath) *QGraphicsPathItem {
	var __rv uintptr
	q.Drv(272000,272118,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsPathItem{}
	_rp.SetDriver(__rv,265000,true)
	return _rp
}	
//QGraphicsScene::addPath(QPainterPath const&,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddPathWithPathPenBrush(path *QPainterPath,pen *QPen,brush *QBrush) *QGraphicsPathItem {
	var __rv uintptr
	q.Drv(272000,272119,Native(path),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsPathItem{}
	_rp.SetDriver(__rv,265000,true)
	return _rp
}	
//QGraphicsScene::addPixmap(QPixmap const&)
func (q *QGraphicsScene) AddPixmap(pixmap *QPixmap) *QGraphicsPixmapItem {
	var __rv uintptr
	q.Drv(272000,272120,Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsPixmapItem{}
	_rp.SetDriver(__rv,266000,true)
	return _rp
}	
//QGraphicsScene::addPolygon(QPolygonF const&)
func (q *QGraphicsScene) AddPolygon(polygon *QPolygonF) *QGraphicsPolygonItem {
	var __rv uintptr
	q.Drv(272000,272121,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsPolygonItem{}
	_rp.SetDriver(__rv,267000,true)
	return _rp
}	
//QGraphicsScene::addPolygon(QPolygonF const&,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddPolygonFWithPolygonPenBrush(polygon *QPolygonF,pen *QPen,brush *QBrush) *QGraphicsPolygonItem {
	var __rv uintptr
	q.Drv(272000,272122,Native(polygon),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsPolygonItem{}
	_rp.SetDriver(__rv,267000,true)
	return _rp
}	
//QGraphicsScene::addRect(QRectF const&)
func (q *QGraphicsScene) AddRect(rect *QRectF) *QGraphicsRectItem {
	var __rv uintptr
	q.Drv(272000,272123,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsRectItem{}
	_rp.SetDriver(__rv,269000,true)
	return _rp
}	
//QGraphicsScene::addRect(QRectF const&,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddRectFWithRectPenBrush(rect *QRectF,pen *QPen,brush *QBrush) *QGraphicsRectItem {
	var __rv uintptr
	q.Drv(272000,272124,Native(rect),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsRectItem{}
	_rp.SetDriver(__rv,269000,true)
	return _rp
}	
//QGraphicsScene::addRect(double,double,double,double,QPen const&,QBrush const&)
func (q *QGraphicsScene) AddRectFWithXYWidthHeightPenBrush(x float64,y float64,w float64,h float64,pen *QPen,brush *QBrush) *QGraphicsRectItem {
	var __rv uintptr
	q.Drv(272000,272125,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),Native(pen),Native(brush),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsRectItem{}
	_rp.SetDriver(__rv,269000,true)
	return _rp
}	
//QGraphicsScene::addSimpleText(QString const&)
func (q *QGraphicsScene) AddSimpleText(text string) *QGraphicsSimpleTextItem {
	var __rv uintptr
	q.Drv(272000,272126,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsSimpleTextItem{}
	_rp.SetDriver(__rv,282000,true)
	return _rp
}	
//QGraphicsScene::addSimpleText(QString const&,QFont const&)
func (q *QGraphicsScene) AddSimpleTextWithTextFont(text string,font *QFont) *QGraphicsSimpleTextItem {
	var __rv uintptr
	q.Drv(272000,272127,unsafe.Pointer(&text),Native(font),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsSimpleTextItem{}
	_rp.SetDriver(__rv,282000,true)
	return _rp
}	
//QGraphicsScene::addText(QString const&)
func (q *QGraphicsScene) AddText(text string) *QGraphicsTextItem {
	var __rv uintptr
	q.Drv(272000,272128,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsTextItem{}
	_rp.SetDriver(__rv,283000,false)
	return _rp
}	
//QGraphicsScene::addText(QString const&,QFont const&)
func (q *QGraphicsScene) AddTextWithTextFont(text string,font *QFont) *QGraphicsTextItem {
	var __rv uintptr
	q.Drv(272000,272129,unsafe.Pointer(&text),Native(font),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsTextItem{}
	_rp.SetDriver(__rv,283000,false)
	return _rp
}	
//QGraphicsScene::addWidget(QWidget*)
func (q *QGraphicsScene) AddWidget(widget QWidgetInterface) *QGraphicsProxyWidget {
	var __rv uintptr
	q.Drv(272000,272130,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsProxyWidget{}
	_rp.SetDriver(__rv,268000,false)
	return _rp
}	
//QGraphicsScene::addWidget(QWidget*,QFlags<Qt::WindowType>)
func (q *QGraphicsScene) AddWidgetWithWidgetFlags(widget QWidgetInterface,wFlags Qt_WindowType) *QGraphicsProxyWidget {
	var __rv uintptr
	q.Drv(272000,272131,Native(widget),unsafe.Pointer(&wFlags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsProxyWidget{}
	_rp.SetDriver(__rv,268000,false)
	return _rp
}	
//QGraphicsScene::advance()
func (q *QGraphicsScene) Advance()  {
	q.Drv(272000,272132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::backgroundBrush()
func (q *QGraphicsScene) BackgroundBrush() *QBrush {
	var __rv uintptr
	q.Drv(272000,272133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QGraphicsScene::bspTreeDepth()
func (q *QGraphicsScene) BspTreeDepth() int {
	var __rv int
	q.Drv(272000,272134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::clear()
func (q *QGraphicsScene) Clear()  {
	q.Drv(272000,272135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::clearFocus()
func (q *QGraphicsScene) ClearFocus()  {
	q.Drv(272000,272136,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::clearSelection()
func (q *QGraphicsScene) ClearSelection()  {
	q.Drv(272000,272137,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::collidingItems(QGraphicsItem const*)
func (q *QGraphicsScene) CollidingItems(item *QGraphicsItem) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272138,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::collidingItems(QGraphicsItem const*,Qt::ItemSelectionMode)
func (q *QGraphicsScene) CollidingItemsWithItemMode(item *QGraphicsItem,mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272139,Native(item),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::contextMenuEvent(QGraphicsSceneContextMenuEvent*)
func (q *QGraphicsScene) ContextMenuEvent(event *QGraphicsSceneContextMenuEvent)  {
	q.Drv(272000,272140,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::createItemGroup(QList<QGraphicsItem*> const&)
func (q *QGraphicsScene) CreateItemGroup(items []*QGraphicsItem) *QGraphicsItemGroup {
	var __rv uintptr
	q.Drv(272000,272141,unsafe.Pointer(&items),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItemGroup{}
	_rp.SetDriver(__rv,258000,true)
	return _rp
}	
//QGraphicsScene::destroyItemGroup(QGraphicsItemGroup*)
func (q *QGraphicsScene) DestroyItemGroup(group *QGraphicsItemGroup)  {
	q.Drv(272000,272142,Native(group),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::dragEnterEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsScene) DragEnterEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(272000,272143,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::dragLeaveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsScene) DragLeaveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(272000,272144,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::dragMoveEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsScene) DragMoveEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(272000,272145,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::drawBackground(QPainter*,QRectF const&)
func (q *QGraphicsScene) DrawBackground(painter *QPainter,rect *QRectF)  {
	q.Drv(272000,272146,Native(painter),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::drawForeground(QPainter*,QRectF const&)
func (q *QGraphicsScene) DrawForeground(painter *QPainter,rect *QRectF)  {
	q.Drv(272000,272147,Native(painter),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::dropEvent(QGraphicsSceneDragDropEvent*)
func (q *QGraphicsScene) DropEvent(event *QGraphicsSceneDragDropEvent)  {
	q.Drv(272000,272148,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::event(QEvent*)
func (q *QGraphicsScene) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(272000,272149,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::eventFilter(QObject*,QEvent*)
func (q *QGraphicsScene) EventFilter(watched QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(272000,272150,Native(watched),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::focusInEvent(QFocusEvent*)
func (q *QGraphicsScene) FocusInEvent(event *QFocusEvent)  {
	q.Drv(272000,272151,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::focusItem()
func (q *QGraphicsScene) FocusItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(272000,272152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsScene::focusNextPrevChild(bool)
func (q *QGraphicsScene) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(272000,272153,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::focusOutEvent(QFocusEvent*)
func (q *QGraphicsScene) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(272000,272154,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::font()
func (q *QGraphicsScene) Font() *QFont {
	var __rv uintptr
	q.Drv(272000,272155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QGraphicsScene::foregroundBrush()
func (q *QGraphicsScene) ForegroundBrush() *QBrush {
	var __rv uintptr
	q.Drv(272000,272156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QBrush{}
	_rp.SetDriver(__rv,9000,true)
	return _rp
}	
//QGraphicsScene::hasFocus()
func (q *QGraphicsScene) HasFocus() bool {
	var __rv bool
	q.Drv(272000,272157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::height()
func (q *QGraphicsScene) Height() float64 {
	var __rv float64
	q.Drv(272000,272158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::helpEvent(QGraphicsSceneHelpEvent*)
func (q *QGraphicsScene) HelpEvent(event *QGraphicsSceneHelpEvent)  {
	q.Drv(272000,272159,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::inputMethodEvent(QInputMethodEvent*)
func (q *QGraphicsScene) InputMethodEvent(event *QInputMethodEvent)  {
	q.Drv(272000,272160,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::inputMethodQuery(Qt::InputMethodQuery)
func (q *QGraphicsScene) InputMethodQuery(query Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(272000,272161,unsafe.Pointer(&query),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsScene::invalidate()
func (q *QGraphicsScene) Invalidate()  {
	q.Drv(272000,272162,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::invalidate(QRectF const&,QFlags<QGraphicsScene::SceneLayer>)
func (q *QGraphicsScene) InvalidateFWithRectLayers(rect *QRectF,layers QGraphicsScene_SceneLayer)  {
	q.Drv(272000,272163,Native(rect),unsafe.Pointer(&layers),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::invalidate(double,double,double,double,QFlags<QGraphicsScene::SceneLayer>)
func (q *QGraphicsScene) InvalidateFWithXYWidthHeightLayers(x float64,y float64,w float64,h float64,layers QGraphicsScene_SceneLayer)  {
	q.Drv(272000,272164,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&layers),nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::isActive()
func (q *QGraphicsScene) IsActive() bool {
	var __rv bool
	q.Drv(272000,272165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::isSortCacheEnabled()
func (q *QGraphicsScene) IsSortCacheEnabled() bool {
	var __rv bool
	q.Drv(272000,272166,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::itemAt(QPointF const&,QTransform const&)
func (q *QGraphicsScene) ItemAtFWithPosTransform(pos *QPointF,deviceTransform *QTransform) *QGraphicsItem {
	var __rv uintptr
	q.Drv(272000,272167,Native(pos),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsScene::itemAt(double,double,QTransform const&)
func (q *QGraphicsScene) ItemAtFWithXYTransform(x float64,y float64,deviceTransform *QTransform) *QGraphicsItem {
	var __rv uintptr
	q.Drv(272000,272168,unsafe.Pointer(&x),unsafe.Pointer(&y),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsScene::itemIndexMethod()
func (q *QGraphicsScene) ItemIndexMethod() QGraphicsScene_ItemIndexMethod {
	var __rv QGraphicsScene_ItemIndexMethod
	q.Drv(272000,272169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items()
func (q *QGraphicsScene) Items() []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(Qt::SortOrder)
func (q *QGraphicsScene) ItemsWithOrder(order Qt_SortOrder) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272171,unsafe.Pointer(&order),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(QPainterPath const&,Qt::ItemSelectionMode,Qt::SortOrder,QTransform const&)
func (q *QGraphicsScene) ItemsWithPathModeOrderTransform(path *QPainterPath,mode Qt_ItemSelectionMode,order Qt_SortOrder,deviceTransform *QTransform) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272172,Native(path),unsafe.Pointer(&mode),unsafe.Pointer(&order),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(QPointF const&,Qt::ItemSelectionMode,Qt::SortOrder,QTransform const&)
func (q *QGraphicsScene) ItemsFWithPosModeOrderTransform(pos *QPointF,mode Qt_ItemSelectionMode,order Qt_SortOrder,deviceTransform *QTransform) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272173,Native(pos),unsafe.Pointer(&mode),unsafe.Pointer(&order),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(QPolygonF const&,Qt::ItemSelectionMode,Qt::SortOrder,QTransform const&)
func (q *QGraphicsScene) ItemsFWithPolygonModeOrderTransform(polygon *QPolygonF,mode Qt_ItemSelectionMode,order Qt_SortOrder,deviceTransform *QTransform) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272174,Native(polygon),unsafe.Pointer(&mode),unsafe.Pointer(&order),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(QRectF const&,Qt::ItemSelectionMode,Qt::SortOrder,QTransform const&)
func (q *QGraphicsScene) ItemsFWithRectModeOrderTransform(rect *QRectF,mode Qt_ItemSelectionMode,order Qt_SortOrder,deviceTransform *QTransform) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272175,Native(rect),unsafe.Pointer(&mode),unsafe.Pointer(&order),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::items(double,double,double,double,Qt::ItemSelectionMode,Qt::SortOrder,QTransform const&)
func (q *QGraphicsScene) ItemsFWithXYWidthHeightModeOrderTransform(x float64,y float64,w float64,h float64,mode Qt_ItemSelectionMode,order Qt_SortOrder,deviceTransform *QTransform) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272176,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&mode),unsafe.Pointer(&order),Native(deviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::itemsBoundingRect()
func (q *QGraphicsScene) ItemsBoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(272000,272177,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsScene::keyPressEvent(QKeyEvent*)
func (q *QGraphicsScene) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(272000,272178,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::keyReleaseEvent(QKeyEvent*)
func (q *QGraphicsScene) KeyReleaseEvent(event *QKeyEvent)  {
	q.Drv(272000,272179,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::mouseDoubleClickEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsScene) MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(272000,272180,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::mouseGrabberItem()
func (q *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(272000,272181,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsScene::mouseMoveEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsScene) MouseMoveEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(272000,272182,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::mousePressEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsScene) MousePressEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(272000,272183,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::mouseReleaseEvent(QGraphicsSceneMouseEvent*)
func (q *QGraphicsScene) MouseReleaseEvent(event *QGraphicsSceneMouseEvent)  {
	q.Drv(272000,272184,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::palette()
func (q *QGraphicsScene) Palette() *QPalette {
	var __rv uintptr
	q.Drv(272000,272185,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QGraphicsScene::removeItem(QGraphicsItem*)
func (q *QGraphicsScene) RemoveItem(item *QGraphicsItem)  {
	q.Drv(272000,272186,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::render(QPainter*)
func (q *QGraphicsScene) Render(painter *QPainter)  {
	q.Drv(272000,272187,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::render(QPainter*,QRectF const&,QRectF const&,Qt::AspectRatioMode)
func (q *QGraphicsScene) RenderFWithPainterTargetSourceAspectratiomode(painter *QPainter,target *QRectF,source *QRectF,aspectRatioMode Qt_AspectRatioMode)  {
	q.Drv(272000,272188,Native(painter),Native(target),Native(source),unsafe.Pointer(&aspectRatioMode),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::sceneRect()
func (q *QGraphicsScene) SceneRect() *QRectF {
	var __rv uintptr
	q.Drv(272000,272189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsScene::selectedItems()
func (q *QGraphicsScene) SelectedItems() []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(272000,272190,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::selectionArea()
func (q *QGraphicsScene) SelectionArea() *QPainterPath {
	var __rv uintptr
	q.Drv(272000,272191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsScene::sendEvent(QGraphicsItem*,QEvent*)
func (q *QGraphicsScene) SendEvent(item *QGraphicsItem,event *QEvent) bool {
	var __rv bool
	q.Drv(272000,272192,Native(item),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::setActivePanel(QGraphicsItem*)
func (q *QGraphicsScene) SetActivePanel(item *QGraphicsItem)  {
	q.Drv(272000,272193,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setActiveWindow(QGraphicsWidget*)
func (q *QGraphicsScene) SetActiveWindow(widget *QGraphicsWidget)  {
	q.Drv(272000,272194,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setBackgroundBrush(QBrush const&)
func (q *QGraphicsScene) SetBackgroundBrush(brush *QBrush)  {
	q.Drv(272000,272195,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setBspTreeDepth(int)
func (q *QGraphicsScene) SetBspTreeDepth(depth int)  {
	q.Drv(272000,272196,unsafe.Pointer(&depth),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setFocus()
func (q *QGraphicsScene) SetFocus()  {
	q.Drv(272000,272197,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setFocus(Qt::FocusReason)
func (q *QGraphicsScene) SetFocusWithFocusreason(focusReason Qt_FocusReason)  {
	q.Drv(272000,272198,unsafe.Pointer(&focusReason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setFocusItem(QGraphicsItem*)
func (q *QGraphicsScene) SetFocusItem(item *QGraphicsItem)  {
	q.Drv(272000,272199,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setFocusItem(QGraphicsItem*,Qt::FocusReason)
func (q *QGraphicsScene) SetFocusItemWithItemFocusreason(item *QGraphicsItem,focusReason Qt_FocusReason)  {
	q.Drv(272000,272200,Native(item),unsafe.Pointer(&focusReason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setFont(QFont const&)
func (q *QGraphicsScene) SetFont(font *QFont)  {
	q.Drv(272000,272201,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setForegroundBrush(QBrush const&)
func (q *QGraphicsScene) SetForegroundBrush(brush *QBrush)  {
	q.Drv(272000,272202,Native(brush),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setItemIndexMethod(QGraphicsScene::ItemIndexMethod)
func (q *QGraphicsScene) SetItemIndexMethod(method QGraphicsScene_ItemIndexMethod)  {
	q.Drv(272000,272203,unsafe.Pointer(&method),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setPalette(QPalette const&)
func (q *QGraphicsScene) SetPalette(palette *QPalette)  {
	q.Drv(272000,272204,Native(palette),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSceneRect(QRectF const&)
func (q *QGraphicsScene) SetSceneRect(rect *QRectF)  {
	q.Drv(272000,272205,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSceneRect(double,double,double,double)
func (q *QGraphicsScene) SetSceneRectFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(272000,272206,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSelectionArea(QPainterPath const&)
func (q *QGraphicsScene) SetSelectionArea(path *QPainterPath)  {
	q.Drv(272000,272207,Native(path),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSelectionArea(QPainterPath const&,QTransform const&)
func (q *QGraphicsScene) SetSelectionAreaWithPathTransform(path *QPainterPath,deviceTransform *QTransform)  {
	q.Drv(272000,272208,Native(path),Native(deviceTransform),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSelectionArea(QPainterPath const&,Qt::ItemSelectionMode)
func (q *QGraphicsScene) SetSelectionAreaWithPathMode(path *QPainterPath,mode Qt_ItemSelectionMode)  {
	q.Drv(272000,272209,Native(path),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSelectionArea(QPainterPath const&,Qt::ItemSelectionMode,QTransform const&)
func (q *QGraphicsScene) SetSelectionAreaWithPathModeTransform(path *QPainterPath,mode Qt_ItemSelectionMode,deviceTransform *QTransform)  {
	q.Drv(272000,272210,Native(path),unsafe.Pointer(&mode),Native(deviceTransform),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setSortCacheEnabled(bool)
func (q *QGraphicsScene) SetSortCacheEnabled(enabled bool)  {
	q.Drv(272000,272211,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setStickyFocus(bool)
func (q *QGraphicsScene) SetStickyFocus(enabled bool)  {
	q.Drv(272000,272212,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::setStyle(QStyle*)
func (q *QGraphicsScene) SetStyle(style *QStyle)  {
	q.Drv(272000,272213,Native(style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::stickyFocus()
func (q *QGraphicsScene) StickyFocus() bool {
	var __rv bool
	q.Drv(272000,272214,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::style()
func (q *QGraphicsScene) Style() *QStyle {
	var __rv uintptr
	q.Drv(272000,272215,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QGraphicsScene::update()
func (q *QGraphicsScene) Update()  {
	q.Drv(272000,272216,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::update(QRectF const&)
func (q *QGraphicsScene) UpdateFWithRect(rect *QRectF)  {
	q.Drv(272000,272217,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::update(double,double,double,double)
func (q *QGraphicsScene) UpdateFWithXYWidthHeight(x float64,y float64,w float64,h float64)  {
	q.Drv(272000,272218,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::views()
func (q *QGraphicsScene) Views() []*QGraphicsView {
	var __rv []*QGraphicsView
	q.Drv(272000,272219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsScene::wheelEvent(QGraphicsSceneWheelEvent*)
func (q *QGraphicsScene) WheelEvent(event *QGraphicsSceneWheelEvent)  {
	q.Drv(272000,272220,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsScene::width()
func (q *QGraphicsScene) Width() float64 {
	var __rv float64
	q.Drv(272000,272221,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
