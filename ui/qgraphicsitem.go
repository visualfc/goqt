// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QGraphicsItem_CacheMode - QGraphicsItem::CacheMode
type QGraphicsItem_CacheMode uint32
const (
	QGraphicsItem_NoCache QGraphicsItem_CacheMode = 0
	QGraphicsItem_ItemCoordinateCache QGraphicsItem_CacheMode = 1
	QGraphicsItem_DeviceCoordinateCache QGraphicsItem_CacheMode = 2
)
//enum QGraphicsItem_PanelModality - QGraphicsItem::PanelModality
type QGraphicsItem_PanelModality uint32
const (
	QGraphicsItem_NonModal QGraphicsItem_PanelModality = 0
	QGraphicsItem_PanelModal QGraphicsItem_PanelModality = 1
	QGraphicsItem_SceneModal QGraphicsItem_PanelModality = 2
)
//enum QGraphicsItem_enum_1 - QGraphicsItem::enum_1
type QGraphicsItem_enum_1 uint32
const (
	QGraphicsItem_Type QGraphicsItem_enum_1 = 1
	QGraphicsItem_UserType QGraphicsItem_enum_1 = 65536
)
//enum QGraphicsItem_GraphicsItemFlag - QGraphicsItem::GraphicsItemFlag
type QGraphicsItem_GraphicsItemFlag uint32
const (
	QGraphicsItem_ItemIsMovable QGraphicsItem_GraphicsItemFlag = 0x1
	QGraphicsItem_ItemIsSelectable QGraphicsItem_GraphicsItemFlag = 0x2
	QGraphicsItem_ItemIsFocusable QGraphicsItem_GraphicsItemFlag = 0x4
	QGraphicsItem_ItemClipsToShape QGraphicsItem_GraphicsItemFlag = 0x8
	QGraphicsItem_ItemClipsChildrenToShape QGraphicsItem_GraphicsItemFlag = 0x10
	QGraphicsItem_ItemIgnoresTransformations QGraphicsItem_GraphicsItemFlag = 0x20
	QGraphicsItem_ItemIgnoresParentOpacity QGraphicsItem_GraphicsItemFlag = 0x40
	QGraphicsItem_ItemDoesntPropagateOpacityToChildren QGraphicsItem_GraphicsItemFlag = 0x80
	QGraphicsItem_ItemStacksBehindParent QGraphicsItem_GraphicsItemFlag = 0x100
	QGraphicsItem_ItemUsesExtendedStyleOption QGraphicsItem_GraphicsItemFlag = 0x200
	QGraphicsItem_ItemHasNoContents QGraphicsItem_GraphicsItemFlag = 0x400
	QGraphicsItem_ItemSendsGeometryChanges QGraphicsItem_GraphicsItemFlag = 0x800
	QGraphicsItem_ItemAcceptsInputMethod QGraphicsItem_GraphicsItemFlag = 0x1000
	QGraphicsItem_ItemNegativeZStacksBehindParent QGraphicsItem_GraphicsItemFlag = 0x2000
	QGraphicsItem_ItemIsPanel QGraphicsItem_GraphicsItemFlag = 0x4000
	QGraphicsItem_ItemIsFocusScope QGraphicsItem_GraphicsItemFlag = 0x8000
	QGraphicsItem_ItemSendsScenePositionChanges QGraphicsItem_GraphicsItemFlag = 0x10000
	QGraphicsItem_ItemStopsClickFocusPropagation QGraphicsItem_GraphicsItemFlag = 0x20000
	QGraphicsItem_ItemStopsFocusHandling QGraphicsItem_GraphicsItemFlag = 0x40000
)
//enum QGraphicsItem_GraphicsItemChange - QGraphicsItem::GraphicsItemChange
type QGraphicsItem_GraphicsItemChange uint32
const (
	QGraphicsItem_ItemPositionChange QGraphicsItem_GraphicsItemChange = 0
	QGraphicsItem_ItemMatrixChange QGraphicsItem_GraphicsItemChange = 1
	QGraphicsItem_ItemVisibleChange QGraphicsItem_GraphicsItemChange = 2
	QGraphicsItem_ItemEnabledChange QGraphicsItem_GraphicsItemChange = 3
	QGraphicsItem_ItemSelectedChange QGraphicsItem_GraphicsItemChange = 4
	QGraphicsItem_ItemParentChange QGraphicsItem_GraphicsItemChange = 5
	QGraphicsItem_ItemChildAddedChange QGraphicsItem_GraphicsItemChange = 6
	QGraphicsItem_ItemChildRemovedChange QGraphicsItem_GraphicsItemChange = 7
	QGraphicsItem_ItemTransformChange QGraphicsItem_GraphicsItemChange = 8
	QGraphicsItem_ItemPositionHasChanged QGraphicsItem_GraphicsItemChange = 9
	QGraphicsItem_ItemTransformHasChanged QGraphicsItem_GraphicsItemChange = 10
	QGraphicsItem_ItemSceneChange QGraphicsItem_GraphicsItemChange = 11
	QGraphicsItem_ItemVisibleHasChanged QGraphicsItem_GraphicsItemChange = 12
	QGraphicsItem_ItemEnabledHasChanged QGraphicsItem_GraphicsItemChange = 13
	QGraphicsItem_ItemSelectedHasChanged QGraphicsItem_GraphicsItemChange = 14
	QGraphicsItem_ItemParentHasChanged QGraphicsItem_GraphicsItemChange = 15
	QGraphicsItem_ItemSceneHasChanged QGraphicsItem_GraphicsItemChange = 16
	QGraphicsItem_ItemCursorChange QGraphicsItem_GraphicsItemChange = 17
	QGraphicsItem_ItemCursorHasChanged QGraphicsItem_GraphicsItemChange = 18
	QGraphicsItem_ItemToolTipChange QGraphicsItem_GraphicsItemChange = 19
	QGraphicsItem_ItemToolTipHasChanged QGraphicsItem_GraphicsItemChange = 20
	QGraphicsItem_ItemFlagsChange QGraphicsItem_GraphicsItemChange = 21
	QGraphicsItem_ItemFlagsHaveChanged QGraphicsItem_GraphicsItemChange = 22
	QGraphicsItem_ItemZValueChange QGraphicsItem_GraphicsItemChange = 23
	QGraphicsItem_ItemZValueHasChanged QGraphicsItem_GraphicsItemChange = 24
	QGraphicsItem_ItemOpacityChange QGraphicsItem_GraphicsItemChange = 25
	QGraphicsItem_ItemOpacityHasChanged QGraphicsItem_GraphicsItemChange = 26
	QGraphicsItem_ItemScenePositionHasChanged QGraphicsItem_GraphicsItemChange = 27
	QGraphicsItem_ItemRotationChange QGraphicsItem_GraphicsItemChange = 28
	QGraphicsItem_ItemRotationHasChanged QGraphicsItem_GraphicsItemChange = 29
	QGraphicsItem_ItemScaleChange QGraphicsItem_GraphicsItemChange = 30
	QGraphicsItem_ItemScaleHasChanged QGraphicsItem_GraphicsItemChange = 31
	QGraphicsItem_ItemTransformOriginPointChange QGraphicsItem_GraphicsItemChange = 32
	QGraphicsItem_ItemTransformOriginPointHasChanged QGraphicsItem_GraphicsItemChange = 33
)
//enum QGraphicsItem_Extension - QGraphicsItem::Extension
type QGraphicsItem_Extension uint32
const (
	QGraphicsItem_UserExtension QGraphicsItem_Extension = 0x80000000
)
//struct QGraphicsItem : QGraphicsItem
type QGraphicsItem struct {
	BaseDrv
}
//QGraphicsItem::acceptDrops()
func (q *QGraphicsItem) AcceptDrops() bool {
	var __rv bool
	q.Drv(256000,256102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::acceptHoverEvents()
func (q *QGraphicsItem) AcceptHoverEvents() bool {
	var __rv bool
	q.Drv(256000,256103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::acceptTouchEvents()
func (q *QGraphicsItem) AcceptTouchEvents() bool {
	var __rv bool
	q.Drv(256000,256104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::acceptedMouseButtons()
func (q *QGraphicsItem) AcceptedMouseButtons() Qt_MouseButton {
	var __rv Qt_MouseButton
	q.Drv(256000,256105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::advance(int)
func (q *QGraphicsItem) Advance(phase int)  {
	q.Drv(256000,256106,unsafe.Pointer(&phase),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::boundingRect()
func (q *QGraphicsItem) BoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(256000,256107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::boundingRegion(QTransform const&)
func (q *QGraphicsItem) BoundingRegion(itemToDeviceTransform *QTransform) *QRegion {
	var __rv uintptr
	q.Drv(256000,256108,Native(itemToDeviceTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
//QGraphicsItem::boundingRegionGranularity()
func (q *QGraphicsItem) BoundingRegionGranularity() float64 {
	var __rv float64
	q.Drv(256000,256109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::cacheMode()
func (q *QGraphicsItem) CacheMode() QGraphicsItem_CacheMode {
	var __rv QGraphicsItem_CacheMode
	q.Drv(256000,256110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::childItems()
func (q *QGraphicsItem) ChildItems() []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(256000,256111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::childrenBoundingRect()
func (q *QGraphicsItem) ChildrenBoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(256000,256112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::clearFocus()
func (q *QGraphicsItem) ClearFocus()  {
	q.Drv(256000,256113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::clipPath()
func (q *QGraphicsItem) ClipPath() *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::collidesWithItem(QGraphicsItem const*)
func (q *QGraphicsItem) CollidesWithItem(other *QGraphicsItem) bool {
	var __rv bool
	q.Drv(256000,256115,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::collidesWithItem(QGraphicsItem const*,Qt::ItemSelectionMode)
func (q *QGraphicsItem) CollidesWithItemWithOtherMode(other *QGraphicsItem,mode Qt_ItemSelectionMode) bool {
	var __rv bool
	q.Drv(256000,256116,Native(other),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::collidesWithPath(QPainterPath const&)
func (q *QGraphicsItem) CollidesWithPath(path *QPainterPath) bool {
	var __rv bool
	q.Drv(256000,256117,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::collidesWithPath(QPainterPath const&,Qt::ItemSelectionMode)
func (q *QGraphicsItem) CollidesWithPathWithPathMode(path *QPainterPath,mode Qt_ItemSelectionMode) bool {
	var __rv bool
	q.Drv(256000,256118,Native(path),unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::collidingItems()
func (q *QGraphicsItem) CollidingItems() []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(256000,256119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::collidingItems(Qt::ItemSelectionMode)
func (q *QGraphicsItem) CollidingItemsWithMode(mode Qt_ItemSelectionMode) []*QGraphicsItem {
	var __rv []*QGraphicsItem
	q.Drv(256000,256120,unsafe.Pointer(&mode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::commonAncestorItem(QGraphicsItem const*)
func (q *QGraphicsItem) CommonAncestorItem(other *QGraphicsItem) *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256121,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::contains(QPointF const&)
func (q *QGraphicsItem) Contains(point *QPointF) bool {
	var __rv bool
	q.Drv(256000,256122,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::cursor()
func (q *QGraphicsItem) Cursor() *QCursor {
	var __rv uintptr
	q.Drv(256000,256123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCursor{}
	_rp.SetDriver(__rv,18000,true)
	return _rp
}	
//QGraphicsItem::data(int)
func (q *QGraphicsItem) Data(key int) *QVariant {
	var __rv uintptr
	q.Drv(256000,256124,unsafe.Pointer(&key),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QGraphicsItem::deviceTransform(QTransform const&)
func (q *QGraphicsItem) DeviceTransform(viewportTransform *QTransform) *QTransform {
	var __rv uintptr
	q.Drv(256000,256125,Native(viewportTransform),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsItem::effectiveOpacity()
func (q *QGraphicsItem) EffectiveOpacity() float64 {
	var __rv float64
	q.Drv(256000,256126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::ensureVisible()
func (q *QGraphicsItem) EnsureVisible()  {
	q.Drv(256000,256127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::ensureVisible(QRectF const&,int,int)
func (q *QGraphicsItem) EnsureVisibleFWithRectXmarginYmargin(rect *QRectF,xmargin int,ymargin int)  {
	q.Drv(256000,256128,Native(rect),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::ensureVisible(double,double,double,double,int,int)
func (q *QGraphicsItem) EnsureVisibleFWithXYWidthHeightXmarginYmargin(x float64,y float64,w float64,h float64,xmargin int,ymargin int)  {
	q.Drv(256000,256129,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&xmargin),unsafe.Pointer(&ymargin),nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::filtersChildEvents()
func (q *QGraphicsItem) FiltersChildEvents() bool {
	var __rv bool
	q.Drv(256000,256130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::flags()
func (q *QGraphicsItem) Flags() QGraphicsItem_GraphicsItemFlag {
	var __rv QGraphicsItem_GraphicsItemFlag
	q.Drv(256000,256131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::focusItem()
func (q *QGraphicsItem) FocusItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::focusProxy()
func (q *QGraphicsItem) FocusProxy() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::focusScopeItem()
func (q *QGraphicsItem) FocusScopeItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::grabKeyboard()
func (q *QGraphicsItem) GrabKeyboard()  {
	q.Drv(256000,256135,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::grabMouse()
func (q *QGraphicsItem) GrabMouse()  {
	q.Drv(256000,256136,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::graphicsEffect()
func (q *QGraphicsItem) GraphicsEffect() *QGraphicsEffect {
	var __rv uintptr
	q.Drv(256000,256137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsEffect{}
	_rp.SetDriver(__rv,253000,false)
	return _rp
}	
//QGraphicsItem::group()
func (q *QGraphicsItem) Group() *QGraphicsItemGroup {
	var __rv uintptr
	q.Drv(256000,256138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItemGroup{}
	_rp.SetDriver(__rv,258000,true)
	return _rp
}	
//QGraphicsItem::handlesChildEvents()
func (q *QGraphicsItem) HandlesChildEvents() bool {
	var __rv bool
	q.Drv(256000,256139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::hasCursor()
func (q *QGraphicsItem) HasCursor() bool {
	var __rv bool
	q.Drv(256000,256140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::hasFocus()
func (q *QGraphicsItem) HasFocus() bool {
	var __rv bool
	q.Drv(256000,256141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::hide()
func (q *QGraphicsItem) Hide()  {
	q.Drv(256000,256142,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::inputMethodHints()
func (q *QGraphicsItem) InputMethodHints() Qt_InputMethodHint {
	var __rv Qt_InputMethodHint
	q.Drv(256000,256143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::installSceneEventFilter(QGraphicsItem*)
func (q *QGraphicsItem) InstallSceneEventFilter(filterItem *QGraphicsItem)  {
	q.Drv(256000,256144,Native(filterItem),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::isActive()
func (q *QGraphicsItem) IsActive() bool {
	var __rv bool
	q.Drv(256000,256145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isAncestorOf(QGraphicsItem const*)
func (q *QGraphicsItem) IsAncestorOf(child *QGraphicsItem) bool {
	var __rv bool
	q.Drv(256000,256146,Native(child),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isClipped()
func (q *QGraphicsItem) IsClipped() bool {
	var __rv bool
	q.Drv(256000,256147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isEnabled()
func (q *QGraphicsItem) IsEnabled() bool {
	var __rv bool
	q.Drv(256000,256148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isObscured()
func (q *QGraphicsItem) IsObscured() bool {
	var __rv bool
	q.Drv(256000,256149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isObscured(QRectF const&)
func (q *QGraphicsItem) IsObscuredFWithRect(rect *QRectF) bool {
	var __rv bool
	q.Drv(256000,256150,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isObscured(double,double,double,double)
func (q *QGraphicsItem) IsObscuredFWithXYWidthHeight(x float64,y float64,w float64,h float64) bool {
	var __rv bool
	q.Drv(256000,256151,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isObscuredBy(QGraphicsItem const*)
func (q *QGraphicsItem) IsObscuredBy(item *QGraphicsItem) bool {
	var __rv bool
	q.Drv(256000,256152,Native(item),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isPanel()
func (q *QGraphicsItem) IsPanel() bool {
	var __rv bool
	q.Drv(256000,256153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isSelected()
func (q *QGraphicsItem) IsSelected() bool {
	var __rv bool
	q.Drv(256000,256154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isUnderMouse()
func (q *QGraphicsItem) IsUnderMouse() bool {
	var __rv bool
	q.Drv(256000,256155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isVisible()
func (q *QGraphicsItem) IsVisible() bool {
	var __rv bool
	q.Drv(256000,256156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isVisibleTo(QGraphicsItem const*)
func (q *QGraphicsItem) IsVisibleTo(parent *QGraphicsItem) bool {
	var __rv bool
	q.Drv(256000,256157,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isWidget()
func (q *QGraphicsItem) IsWidget() bool {
	var __rv bool
	q.Drv(256000,256158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::isWindow()
func (q *QGraphicsItem) IsWindow() bool {
	var __rv bool
	q.Drv(256000,256159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::itemTransform(QGraphicsItem const*)
func (q *QGraphicsItem) ItemTransform(other *QGraphicsItem) *QTransform {
	var __rv uintptr
	q.Drv(256000,256160,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsItem::itemTransform(QGraphicsItem const*,bool*)
func (q *QGraphicsItem) ItemTransformWithOtherOk(other *QGraphicsItem,ok *bool) *QTransform {
	var __rv uintptr
	q.Drv(256000,256161,Native(other),unsafe.Pointer(&ok),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,QPainterPath const&)
func (q *QGraphicsItem) MapFromItemWithItemPath(item *QGraphicsItem,path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256162,Native(item),Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,QPointF const&)
func (q *QGraphicsItem) MapFromItemFWithItemPoint(item *QGraphicsItem,point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256163,Native(item),Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,QPolygonF const&)
func (q *QGraphicsItem) MapFromItemFWithItemPolygon(item *QGraphicsItem,polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256164,Native(item),Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,QRectF const&)
func (q *QGraphicsItem) MapFromItemFWithItemRect(item *QGraphicsItem,rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256165,Native(item),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,double,double)
func (q *QGraphicsItem) MapFromItemFWithItemXY(item *QGraphicsItem,x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256166,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromItem(QGraphicsItem const*,double,double,double,double)
func (q *QGraphicsItem) MapFromItemFWithItemXYWidthHeight(item *QGraphicsItem,x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256167,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(QPainterPath const&)
func (q *QGraphicsItem) MapFromParent(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256168,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(QPointF const&)
func (q *QGraphicsItem) MapFromParentFWithPoint(point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256169,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(QPolygonF const&)
func (q *QGraphicsItem) MapFromParentFWithPolygon(polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256170,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(QRectF const&)
func (q *QGraphicsItem) MapFromParentFWithRect(rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256171,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(double,double)
func (q *QGraphicsItem) MapFromParentFWithXY(x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256172,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromParent(double,double,double,double)
func (q *QGraphicsItem) MapFromParentFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256173,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(QPainterPath const&)
func (q *QGraphicsItem) MapFromScene(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256174,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(QPointF const&)
func (q *QGraphicsItem) MapFromSceneFWithPoint(point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256175,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(QPolygonF const&)
func (q *QGraphicsItem) MapFromSceneFWithPolygon(polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256176,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(QRectF const&)
func (q *QGraphicsItem) MapFromSceneFWithRect(rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256177,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(double,double)
func (q *QGraphicsItem) MapFromSceneFWithXY(x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256178,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapFromScene(double,double,double,double)
func (q *QGraphicsItem) MapFromSceneFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256179,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromItem(QGraphicsItem const*,QRectF const&)
func (q *QGraphicsItem) MapRectFromItemFWithItemRect(item *QGraphicsItem,rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256180,Native(item),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromItem(QGraphicsItem const*,double,double,double,double)
func (q *QGraphicsItem) MapRectFromItemFWithItemXYWidthHeight(item *QGraphicsItem,x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256181,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromParent(QRectF const&)
func (q *QGraphicsItem) MapRectFromParent(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256182,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromParent(double,double,double,double)
func (q *QGraphicsItem) MapRectFromParentFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256183,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromScene(QRectF const&)
func (q *QGraphicsItem) MapRectFromScene(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256184,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectFromScene(double,double,double,double)
func (q *QGraphicsItem) MapRectFromSceneFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256185,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToItem(QGraphicsItem const*,QRectF const&)
func (q *QGraphicsItem) MapRectToItemFWithItemRect(item *QGraphicsItem,rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256186,Native(item),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToItem(QGraphicsItem const*,double,double,double,double)
func (q *QGraphicsItem) MapRectToItemFWithItemXYWidthHeight(item *QGraphicsItem,x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256187,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToParent(QRectF const&)
func (q *QGraphicsItem) MapRectToParent(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256188,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToParent(double,double,double,double)
func (q *QGraphicsItem) MapRectToParentFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256189,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToScene(QRectF const&)
func (q *QGraphicsItem) MapRectToScene(rect *QRectF) *QRectF {
	var __rv uintptr
	q.Drv(256000,256190,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapRectToScene(double,double,double,double)
func (q *QGraphicsItem) MapRectToSceneFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QRectF {
	var __rv uintptr
	q.Drv(256000,256191,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,QPainterPath const&)
func (q *QGraphicsItem) MapToItemWithItemPath(item *QGraphicsItem,path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256192,Native(item),Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,QPointF const&)
func (q *QGraphicsItem) MapToItemFWithItemPoint(item *QGraphicsItem,point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256193,Native(item),Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,QPolygonF const&)
func (q *QGraphicsItem) MapToItemFWithItemPolygon(item *QGraphicsItem,polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256194,Native(item),Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,QRectF const&)
func (q *QGraphicsItem) MapToItemFWithItemRect(item *QGraphicsItem,rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256195,Native(item),Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,double,double)
func (q *QGraphicsItem) MapToItemFWithItemXY(item *QGraphicsItem,x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256196,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToItem(QGraphicsItem const*,double,double,double,double)
func (q *QGraphicsItem) MapToItemFWithItemXYWidthHeight(item *QGraphicsItem,x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256197,Native(item),unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(QPainterPath const&)
func (q *QGraphicsItem) MapToParent(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256198,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(QPointF const&)
func (q *QGraphicsItem) MapToParentFWithPoint(point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256199,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(QPolygonF const&)
func (q *QGraphicsItem) MapToParentFWithPolygon(polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256200,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(QRectF const&)
func (q *QGraphicsItem) MapToParentFWithRect(rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256201,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(double,double)
func (q *QGraphicsItem) MapToParentFWithXY(x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256202,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToParent(double,double,double,double)
func (q *QGraphicsItem) MapToParentFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256203,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(QPainterPath const&)
func (q *QGraphicsItem) MapToScene(path *QPainterPath) *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256204,Native(path),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(QPointF const&)
func (q *QGraphicsItem) MapToSceneFWithPoint(point *QPointF) *QPointF {
	var __rv uintptr
	q.Drv(256000,256205,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(QPolygonF const&)
func (q *QGraphicsItem) MapToSceneFWithPolygon(polygon *QPolygonF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256206,Native(polygon),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(QRectF const&)
func (q *QGraphicsItem) MapToSceneFWithRect(rect *QRectF) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256207,Native(rect),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(double,double)
func (q *QGraphicsItem) MapToSceneFWithXY(x float64,y float64) *QPointF {
	var __rv uintptr
	q.Drv(256000,256208,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::mapToScene(double,double,double,double)
func (q *QGraphicsItem) MapToSceneFWithXYWidthHeight(x float64,y float64,w float64,h float64) *QPolygonF {
	var __rv uintptr
	q.Drv(256000,256209,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPolygonF{}
	_rp.SetDriver(__rv,102000,true)
	return _rp
}	
//QGraphicsItem::matrix()
func (q *QGraphicsItem) Matrix() *QMatrix {
	var __rv uintptr
	q.Drv(256000,256210,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QGraphicsItem::moveBy(double,double)
func (q *QGraphicsItem) MoveBy(dx float64,dy float64)  {
	q.Drv(256000,256211,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::opacity()
func (q *QGraphicsItem) Opacity() float64 {
	var __rv float64
	q.Drv(256000,256212,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::opaqueArea()
func (q *QGraphicsItem) OpaqueArea() *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::panel()
func (q *QGraphicsItem) Panel() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256214,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::panelModality()
func (q *QGraphicsItem) PanelModality() QGraphicsItem_PanelModality {
	var __rv QGraphicsItem_PanelModality
	q.Drv(256000,256215,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::parentItem()
func (q *QGraphicsItem) ParentItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256216,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::parentObject()
func (q *QGraphicsItem) ParentObject() *QGraphicsObject {
	var __rv uintptr
	q.Drv(256000,256217,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsObject{}
	_rp.SetDriver(__rv,263000,false)
	return _rp
}	
//QGraphicsItem::parentWidget()
func (q *QGraphicsItem) ParentWidget() *QGraphicsWidget {
	var __rv uintptr
	q.Drv(256000,256218,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsWidget{}
	_rp.SetDriver(__rv,286000,false)
	return _rp
}	
//QGraphicsItem::pos()
func (q *QGraphicsItem) Pos() *QPointF {
	var __rv uintptr
	q.Drv(256000,256219,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::removeSceneEventFilter(QGraphicsItem*)
func (q *QGraphicsItem) RemoveSceneEventFilter(filterItem *QGraphicsItem)  {
	q.Drv(256000,256220,Native(filterItem),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::resetMatrix()
func (q *QGraphicsItem) ResetMatrix()  {
	q.Drv(256000,256221,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::resetTransform()
func (q *QGraphicsItem) ResetTransform()  {
	q.Drv(256000,256222,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::rotation()
func (q *QGraphicsItem) Rotation() float64 {
	var __rv float64
	q.Drv(256000,256223,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::scene()
func (q *QGraphicsItem) Scene() *QGraphicsScene {
	var __rv uintptr
	q.Drv(256000,256224,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsScene{}
	_rp.SetDriver(__rv,272000,false)
	return _rp
}	
//QGraphicsItem::sceneBoundingRect()
func (q *QGraphicsItem) SceneBoundingRect() *QRectF {
	var __rv uintptr
	q.Drv(256000,256225,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QGraphicsItem::sceneMatrix()
func (q *QGraphicsItem) SceneMatrix() *QMatrix {
	var __rv uintptr
	q.Drv(256000,256226,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMatrix{}
	_rp.SetDriver(__rv,74000,true)
	return _rp
}	
//QGraphicsItem::scenePos()
func (q *QGraphicsItem) ScenePos() *QPointF {
	var __rv uintptr
	q.Drv(256000,256227,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::sceneTransform()
func (q *QGraphicsItem) SceneTransform() *QTransform {
	var __rv uintptr
	q.Drv(256000,256228,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsItem::scroll(double,double,QRectF const&)
func (q *QGraphicsItem) Scroll(dx float64,dy float64,rect *QRectF)  {
	q.Drv(256000,256229,unsafe.Pointer(&dx),unsafe.Pointer(&dy),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setAcceptDrops(bool)
func (q *QGraphicsItem) SetAcceptDrops(on bool)  {
	q.Drv(256000,256230,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setAcceptHoverEvents(bool)
func (q *QGraphicsItem) SetAcceptHoverEvents(enabled bool)  {
	q.Drv(256000,256231,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setAcceptTouchEvents(bool)
func (q *QGraphicsItem) SetAcceptTouchEvents(enabled bool)  {
	q.Drv(256000,256232,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setAcceptedMouseButtons(QFlags<Qt::MouseButton>)
func (q *QGraphicsItem) SetAcceptedMouseButtons(buttons Qt_MouseButton)  {
	q.Drv(256000,256233,unsafe.Pointer(&buttons),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setActive(bool)
func (q *QGraphicsItem) SetActive(active bool)  {
	q.Drv(256000,256234,unsafe.Pointer(&active),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setBoundingRegionGranularity(double)
func (q *QGraphicsItem) SetBoundingRegionGranularity(granularity float64)  {
	q.Drv(256000,256235,unsafe.Pointer(&granularity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setCacheMode(QGraphicsItem::CacheMode)
func (q *QGraphicsItem) SetCacheMode(mode QGraphicsItem_CacheMode)  {
	q.Drv(256000,256236,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setCacheMode(QGraphicsItem::CacheMode,QSize const&)
func (q *QGraphicsItem) SetCacheModeWithModeCachesize(mode QGraphicsItem_CacheMode,cacheSize *QSize)  {
	q.Drv(256000,256237,unsafe.Pointer(&mode),Native(cacheSize),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setCursor(QCursor const&)
func (q *QGraphicsItem) SetCursor(cursor *QCursor)  {
	q.Drv(256000,256238,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setData(int,QVariant const&)
func (q *QGraphicsItem) SetData(key int,value *QVariant)  {
	q.Drv(256000,256239,unsafe.Pointer(&key),Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setEnabled(bool)
func (q *QGraphicsItem) SetEnabled(enabled bool)  {
	q.Drv(256000,256240,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFiltersChildEvents(bool)
func (q *QGraphicsItem) SetFiltersChildEvents(enabled bool)  {
	q.Drv(256000,256241,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFlag(QGraphicsItem::GraphicsItemFlag)
func (q *QGraphicsItem) SetFlag(flag QGraphicsItem_GraphicsItemFlag)  {
	q.Drv(256000,256242,unsafe.Pointer(&flag),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFlag(QGraphicsItem::GraphicsItemFlag,bool)
func (q *QGraphicsItem) SetFlagWithFlagEnabled(flag QGraphicsItem_GraphicsItemFlag,enabled bool)  {
	q.Drv(256000,256243,unsafe.Pointer(&flag),unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFlags(QFlags<QGraphicsItem::GraphicsItemFlag>)
func (q *QGraphicsItem) SetFlags(flags QGraphicsItem_GraphicsItemFlag)  {
	q.Drv(256000,256244,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFocus()
func (q *QGraphicsItem) SetFocus()  {
	q.Drv(256000,256245,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFocus(Qt::FocusReason)
func (q *QGraphicsItem) SetFocusWithFocusreason(focusReason Qt_FocusReason)  {
	q.Drv(256000,256246,unsafe.Pointer(&focusReason),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setFocusProxy(QGraphicsItem*)
func (q *QGraphicsItem) SetFocusProxy(item *QGraphicsItem)  {
	q.Drv(256000,256247,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setGraphicsEffect(QGraphicsEffect*)
func (q *QGraphicsItem) SetGraphicsEffect(effect *QGraphicsEffect)  {
	q.Drv(256000,256248,Native(effect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setGroup(QGraphicsItemGroup*)
func (q *QGraphicsItem) SetGroup(group *QGraphicsItemGroup)  {
	q.Drv(256000,256249,Native(group),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setHandlesChildEvents(bool)
func (q *QGraphicsItem) SetHandlesChildEvents(enabled bool)  {
	q.Drv(256000,256250,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setInputMethodHints(QFlags<Qt::InputMethodHint>)
func (q *QGraphicsItem) SetInputMethodHints(hints Qt_InputMethodHint)  {
	q.Drv(256000,256251,unsafe.Pointer(&hints),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setMatrix(QMatrix const&)
func (q *QGraphicsItem) SetMatrix(matrix *QMatrix)  {
	q.Drv(256000,256252,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setMatrix(QMatrix const&,bool)
func (q *QGraphicsItem) SetMatrixWithMatrixCombine(matrix *QMatrix,combine bool)  {
	q.Drv(256000,256253,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setOpacity(double)
func (q *QGraphicsItem) SetOpacity(opacity float64)  {
	q.Drv(256000,256254,unsafe.Pointer(&opacity),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setPanelModality(QGraphicsItem::PanelModality)
func (q *QGraphicsItem) SetPanelModality(panelModality QGraphicsItem_PanelModality)  {
	q.Drv(256000,256255,unsafe.Pointer(&panelModality),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setParentItem(QGraphicsItem*)
func (q *QGraphicsItem) SetParentItem(parent *QGraphicsItem)  {
	q.Drv(256000,256256,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setPos(QPointF const&)
func (q *QGraphicsItem) SetPos(pos *QPointF)  {
	q.Drv(256000,256257,Native(pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setPos(double,double)
func (q *QGraphicsItem) SetPosFWithXY(x float64,y float64)  {
	q.Drv(256000,256258,unsafe.Pointer(&x),unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setRotation(double)
func (q *QGraphicsItem) SetRotation(angle float64)  {
	q.Drv(256000,256259,unsafe.Pointer(&angle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setScale(double)
func (q *QGraphicsItem) SetScale(scale float64)  {
	q.Drv(256000,256260,unsafe.Pointer(&scale),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setSelected(bool)
func (q *QGraphicsItem) SetSelected(selected bool)  {
	q.Drv(256000,256261,unsafe.Pointer(&selected),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setToolTip(QString const&)
func (q *QGraphicsItem) SetToolTip(toolTip string)  {
	q.Drv(256000,256262,unsafe.Pointer(&toolTip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setTransform(QTransform const&)
func (q *QGraphicsItem) SetTransform(matrix *QTransform)  {
	q.Drv(256000,256263,Native(matrix),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setTransform(QTransform const&,bool)
func (q *QGraphicsItem) SetTransformWithTransformCombine(matrix *QTransform,combine bool)  {
	q.Drv(256000,256264,Native(matrix),unsafe.Pointer(&combine),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setTransformOriginPoint(QPointF const&)
func (q *QGraphicsItem) SetTransformOriginPoint(origin *QPointF)  {
	q.Drv(256000,256265,Native(origin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setTransformOriginPoint(double,double)
func (q *QGraphicsItem) SetTransformOriginPointFWithAxAy(ax float64,ay float64)  {
	q.Drv(256000,256266,unsafe.Pointer(&ax),unsafe.Pointer(&ay),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setTransformations(QList<QGraphicsTransform*> const&)
func (q *QGraphicsItem) SetTransformations(transformations []*QGraphicsTransform)  {
	q.Drv(256000,256267,unsafe.Pointer(&transformations),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setVisible(bool)
func (q *QGraphicsItem) SetVisible(visible bool)  {
	q.Drv(256000,256268,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setX(double)
func (q *QGraphicsItem) SetX(x float64)  {
	q.Drv(256000,256269,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setY(double)
func (q *QGraphicsItem) SetY(y float64)  {
	q.Drv(256000,256270,unsafe.Pointer(&y),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::setZValue(double)
func (q *QGraphicsItem) SetZValue(z float64)  {
	q.Drv(256000,256271,unsafe.Pointer(&z),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::shape()
func (q *QGraphicsItem) Shape() *QPainterPath {
	var __rv uintptr
	q.Drv(256000,256272,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPainterPath{}
	_rp.SetDriver(__rv,88000,true)
	return _rp
}	
//QGraphicsItem::show()
func (q *QGraphicsItem) Show()  {
	q.Drv(256000,256273,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::stackBefore(QGraphicsItem const*)
func (q *QGraphicsItem) StackBefore(sibling *QGraphicsItem)  {
	q.Drv(256000,256274,Native(sibling),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::toGraphicsObject()
func (q *QGraphicsItem) ToGraphicsObject() *QGraphicsObject {
	var __rv uintptr
	q.Drv(256000,256275,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsObject{}
	_rp.SetDriver(__rv,263000,false)
	return _rp
}	
//QGraphicsItem::toolTip()
func (q *QGraphicsItem) ToolTip() string {
	var __rv string
	q.Drv(256000,256276,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::topLevelItem()
func (q *QGraphicsItem) TopLevelItem() *QGraphicsItem {
	var __rv uintptr
	q.Drv(256000,256277,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsItem{}
	_rp.SetDriver(__rv,256000,true)
	return _rp
}	
//QGraphicsItem::topLevelWidget()
func (q *QGraphicsItem) TopLevelWidget() *QGraphicsWidget {
	var __rv uintptr
	q.Drv(256000,256278,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsWidget{}
	_rp.SetDriver(__rv,286000,false)
	return _rp
}	
//QGraphicsItem::transform()
func (q *QGraphicsItem) Transform() *QTransform {
	var __rv uintptr
	q.Drv(256000,256279,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTransform{}
	_rp.SetDriver(__rv,176000,true)
	return _rp
}	
//QGraphicsItem::transformOriginPoint()
func (q *QGraphicsItem) TransformOriginPoint() *QPointF {
	var __rv uintptr
	q.Drv(256000,256280,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QGraphicsItem::transformations()
func (q *QGraphicsItem) Transformations() []*QGraphicsTransform {
	var __rv []*QGraphicsTransform
	q.Drv(256000,256281,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::type()
func (q *QGraphicsItem) Type() int {
	var __rv int
	q.Drv(256000,256282,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::ungrabKeyboard()
func (q *QGraphicsItem) UngrabKeyboard()  {
	q.Drv(256000,256283,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::ungrabMouse()
func (q *QGraphicsItem) UngrabMouse()  {
	q.Drv(256000,256284,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::unsetCursor()
func (q *QGraphicsItem) UnsetCursor()  {
	q.Drv(256000,256285,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::update()
func (q *QGraphicsItem) Update()  {
	q.Drv(256000,256286,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::update(QRectF const&)
func (q *QGraphicsItem) UpdateFWithRect(rect *QRectF)  {
	q.Drv(256000,256287,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::update(double,double,double,double)
func (q *QGraphicsItem) UpdateFWithXYWidthHeight(x float64,y float64,width float64,height float64)  {
	q.Drv(256000,256288,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&width),unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGraphicsItem::window()
func (q *QGraphicsItem) Window() *QGraphicsWidget {
	var __rv uintptr
	q.Drv(256000,256289,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QGraphicsWidget{}
	_rp.SetDriver(__rv,286000,false)
	return _rp
}	
//QGraphicsItem::x()
func (q *QGraphicsItem) X() float64 {
	var __rv float64
	q.Drv(256000,256290,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::y()
func (q *QGraphicsItem) Y() float64 {
	var __rv float64
	q.Drv(256000,256291,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGraphicsItem::zValue()
func (q *QGraphicsItem) ZValue() float64 {
	var __rv float64
	q.Drv(256000,256292,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
