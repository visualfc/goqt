// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QHeaderView_ResizeMode - QHeaderView::ResizeMode
type QHeaderView_ResizeMode uint32
const (
	QHeaderView_Interactive QHeaderView_ResizeMode = 0
	QHeaderView_Stretch QHeaderView_ResizeMode = 1
	QHeaderView_Fixed QHeaderView_ResizeMode = 2
	QHeaderView_ResizeToContents QHeaderView_ResizeMode = 3
	QHeaderView_Custom QHeaderView_ResizeMode = QHeaderView_Fixed
)
//struct QHeaderView : QHeaderView
type QHeaderView struct {
	QAbstractItemView
}
func (q *QHeaderView) OnSectionEntered(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionPressed(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionHandleDoubleClicked(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnGeometriesChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(290000,290105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionResized(fn func(int,int,int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSortIndicatorChanged(fn func(int,Qt_SortOrder)) uintptr {
	var __rv uintptr
	q.Drv(290000,290107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionMoved(fn func(int,int,int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionAutoResize(fn func(int,QHeaderView_ResizeMode)) uintptr {
	var __rv uintptr
	q.Drv(290000,290109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionDoubleClicked(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionClicked(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290111,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QHeaderView) OnSectionCountChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(290000,290112,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QHeaderView::cascadingSectionResizes()
func (q *QHeaderView) CascadingSectionResizes() bool {
	var __rv bool
	q.Drv(290000,290113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::count()
func (q *QHeaderView) Count() int {
	var __rv int
	q.Drv(290000,290114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::currentChanged(QModelIndex const&,QModelIndex const&)
func (q *QHeaderView) CurrentChanged(current *QModelIndex,old *QModelIndex)  {
	q.Drv(290000,290115,Native(current),Native(old),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::dataChanged(QModelIndex const&,QModelIndex const&)
func (q *QHeaderView) DataChanged(topLeft *QModelIndex,bottomRight *QModelIndex)  {
	q.Drv(290000,290116,Native(topLeft),Native(bottomRight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::defaultAlignment()
func (q *QHeaderView) DefaultAlignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(290000,290117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::defaultSectionSize()
func (q *QHeaderView) DefaultSectionSize() int {
	var __rv int
	q.Drv(290000,290118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::doItemsLayout()
func (q *QHeaderView) DoItemsLayout()  {
	q.Drv(290000,290119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::event(QEvent*)
func (q *QHeaderView) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(290000,290120,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::headerDataChanged(Qt::Orientation,int,int)
func (q *QHeaderView) HeaderDataChanged(orientation Qt_Orientation,logicalFirst int,logicalLast int)  {
	q.Drv(290000,290121,unsafe.Pointer(&orientation),unsafe.Pointer(&logicalFirst),unsafe.Pointer(&logicalLast),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::hiddenSectionCount()
func (q *QHeaderView) HiddenSectionCount() int {
	var __rv int
	q.Drv(290000,290122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::hideSection(int)
func (q *QHeaderView) HideSection(logicalIndex int)  {
	q.Drv(290000,290123,unsafe.Pointer(&logicalIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::highlightSections()
func (q *QHeaderView) HighlightSections() bool {
	var __rv bool
	q.Drv(290000,290124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::horizontalOffset()
func (q *QHeaderView) HorizontalOffset() int {
	var __rv int
	q.Drv(290000,290125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::indexAt(QPoint const&)
func (q *QHeaderView) IndexAt(p *QPoint) *QModelIndex {
	var __rv uintptr
	q.Drv(290000,290126,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QHeaderView::initialize()
func (q *QHeaderView) Initialize()  {
	q.Drv(290000,290127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::initializeSections()
func (q *QHeaderView) InitializeSections()  {
	q.Drv(290000,290128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::initializeSections(int,int)
func (q *QHeaderView) InitializeSectionsWithStartEnd(start int,end int)  {
	q.Drv(290000,290129,unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::isClickable()
func (q *QHeaderView) SectionsClickable() bool {
	var __rv bool
	q.Drv(290000,290130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::isIndexHidden(QModelIndex const&)
func (q *QHeaderView) IsIndexHidden(index *QModelIndex) bool {
	var __rv bool
	q.Drv(290000,290131,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::isMovable()
func (q *QHeaderView) SectionsMovable() bool {
	var __rv bool
	q.Drv(290000,290132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::isSectionHidden(int)
func (q *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	var __rv bool
	q.Drv(290000,290133,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::isSortIndicatorShown()
func (q *QHeaderView) IsSortIndicatorShown() bool {
	var __rv bool
	q.Drv(290000,290134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::length()
func (q *QHeaderView) Length() int {
	var __rv int
	q.Drv(290000,290135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::logicalIndex(int)
func (q *QHeaderView) LogicalIndex(visualIndex int) int {
	var __rv int
	q.Drv(290000,290136,unsafe.Pointer(&visualIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::logicalIndexAt(QPoint const&)
func (q *QHeaderView) LogicalIndexAt(pos *QPoint) int {
	var __rv int
	q.Drv(290000,290137,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::logicalIndexAt(int)
func (q *QHeaderView) LogicalIndexAtWithPosition(position int) int {
	var __rv int
	q.Drv(290000,290138,unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::logicalIndexAt(int,int)
func (q *QHeaderView) LogicalIndexAtWithXY(x int,y int) int {
	var __rv int
	q.Drv(290000,290139,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::minimumSectionSize()
func (q *QHeaderView) MinimumSectionSize() int {
	var __rv int
	q.Drv(290000,290140,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::mouseDoubleClickEvent(QMouseEvent*)
func (q *QHeaderView) MouseDoubleClickEvent(e *QMouseEvent)  {
	q.Drv(290000,290141,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::mouseMoveEvent(QMouseEvent*)
func (q *QHeaderView) MouseMoveEvent(e *QMouseEvent)  {
	q.Drv(290000,290142,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::mousePressEvent(QMouseEvent*)
func (q *QHeaderView) MousePressEvent(e *QMouseEvent)  {
	q.Drv(290000,290143,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::mouseReleaseEvent(QMouseEvent*)
func (q *QHeaderView) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(290000,290144,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::moveCursor(QAbstractItemView::CursorAction,QFlags<Qt::KeyboardModifier>)
func (q *QHeaderView) MoveCursor(value2 QAbstractItemView_CursorAction,value3 Qt_KeyboardModifier) *QModelIndex {
	var __rv uintptr
	q.Drv(290000,290145,unsafe.Pointer(&value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QHeaderView::moveSection(int,int)
func (q *QHeaderView) MoveSection(from int,to int)  {
	q.Drv(290000,290146,unsafe.Pointer(&from),unsafe.Pointer(&to),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::offset()
func (q *QHeaderView) Offset() int {
	var __rv int
	q.Drv(290000,290147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::orientation()
func (q *QHeaderView) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(290000,290148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::paintEvent(QPaintEvent*)
func (q *QHeaderView) PaintEvent(e *QPaintEvent)  {
	q.Drv(290000,290149,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::paintSection(QPainter*,QRect const&,int)
func (q *QHeaderView) PaintSection(painter *QPainter,rect *QRect,logicalIndex int)  {
	q.Drv(290000,290150,Native(painter),Native(rect),unsafe.Pointer(&logicalIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::reset()
func (q *QHeaderView) Reset()  {
	q.Drv(290000,290151,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::resizeMode(int)
func (q *QHeaderView) SectionResizeMode(logicalIndex int) QHeaderView_ResizeMode {
	var __rv QHeaderView_ResizeMode
	q.Drv(290000,290152,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::resizeSection(int,int)
func (q *QHeaderView) ResizeSection(logicalIndex int,size int)  {
	q.Drv(290000,290153,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::resizeSections()
func (q *QHeaderView) ResizeSections()  {
	q.Drv(290000,290154,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::resizeSections(QHeaderView::ResizeMode)
func (q *QHeaderView) ResizeSectionsWithMode(mode QHeaderView_ResizeMode)  {
	q.Drv(290000,290155,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::restoreState(QByteArray const&)
func (q *QHeaderView) RestoreState(state []byte) bool {
	var __rv bool
	q.Drv(290000,290156,unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::rowsInserted(QModelIndex const&,int,int)
func (q *QHeaderView) RowsInserted(parent *QModelIndex,start int,end int)  {
	q.Drv(290000,290157,Native(parent),unsafe.Pointer(&start),unsafe.Pointer(&end),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::saveState()
func (q *QHeaderView) SaveState() []byte {
	var __rv []byte
	q.Drv(290000,290158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::scrollContentsBy(int,int)
func (q *QHeaderView) ScrollContentsBy(dx int,dy int)  {
	q.Drv(290000,290159,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::scrollTo(QModelIndex const&,QAbstractItemView::ScrollHint)
func (q *QHeaderView) ScrollTo(index *QModelIndex,hint QAbstractItemView_ScrollHint)  {
	q.Drv(290000,290160,Native(index),unsafe.Pointer(&hint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::sectionPosition(int)
func (q *QHeaderView) SectionPosition(logicalIndex int) int {
	var __rv int
	q.Drv(290000,290161,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sectionSize(int)
func (q *QHeaderView) SectionSize(logicalIndex int) int {
	var __rv int
	q.Drv(290000,290162,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sectionSizeFromContents(int)
func (q *QHeaderView) SectionSizeFromContents(logicalIndex int) *QSize {
	var __rv uintptr
	q.Drv(290000,290163,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QHeaderView::sectionSizeHint(int)
func (q *QHeaderView) SectionSizeHint(logicalIndex int) int {
	var __rv int
	q.Drv(290000,290164,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sectionViewportPosition(int)
func (q *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	var __rv int
	q.Drv(290000,290165,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sectionsAboutToBeRemoved(QModelIndex const&,int,int)
func (q *QHeaderView) SectionsAboutToBeRemoved(parent *QModelIndex,logicalFirst int,logicalLast int)  {
	q.Drv(290000,290166,Native(parent),unsafe.Pointer(&logicalFirst),unsafe.Pointer(&logicalLast),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::sectionsHidden()
func (q *QHeaderView) SectionsHidden() bool {
	var __rv bool
	q.Drv(290000,290167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sectionsInserted(QModelIndex const&,int,int)
func (q *QHeaderView) SectionsInserted(parent *QModelIndex,logicalFirst int,logicalLast int)  {
	q.Drv(290000,290168,Native(parent),unsafe.Pointer(&logicalFirst),unsafe.Pointer(&logicalLast),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::sectionsMoved()
func (q *QHeaderView) SectionsMoved() bool {
	var __rv bool
	q.Drv(290000,290169,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::setCascadingSectionResizes(bool)
func (q *QHeaderView) SetCascadingSectionResizes(enable bool)  {
	q.Drv(290000,290170,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setClickable(bool)
func (q *QHeaderView) SetSectionsClickable(clickable bool)  {
	q.Drv(290000,290171,unsafe.Pointer(&clickable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setDefaultAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QHeaderView) SetDefaultAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(290000,290172,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setDefaultSectionSize(int)
func (q *QHeaderView) SetDefaultSectionSize(size int)  {
	q.Drv(290000,290173,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setHighlightSections(bool)
func (q *QHeaderView) SetHighlightSections(highlight bool)  {
	q.Drv(290000,290174,unsafe.Pointer(&highlight),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setMinimumSectionSize(int)
func (q *QHeaderView) SetMinimumSectionSize(size int)  {
	q.Drv(290000,290175,unsafe.Pointer(&size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setModel(QAbstractItemModel*)
func (q *QHeaderView) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(290000,290176,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setMovable(bool)
func (q *QHeaderView) SetSectionsMovable(movable bool)  {
	q.Drv(290000,290177,unsafe.Pointer(&movable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setOffset(int)
func (q *QHeaderView) SetOffset(offset int)  {
	q.Drv(290000,290178,unsafe.Pointer(&offset),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setOffsetToLastSection()
func (q *QHeaderView) SetOffsetToLastSection()  {
	q.Drv(290000,290179,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setOffsetToSectionPosition(int)
func (q *QHeaderView) SetOffsetToSectionPosition(visualIndex int)  {
	q.Drv(290000,290180,unsafe.Pointer(&visualIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setResizeMode(QHeaderView::ResizeMode)
func (q *QHeaderView) SetSectionResizeMode(mode QHeaderView_ResizeMode)  {
	q.Drv(290000,290181,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setResizeMode(int,QHeaderView::ResizeMode)
func (q *QHeaderView) SetSectionResizeModeWithLogicalindexMode(logicalIndex int,mode QHeaderView_ResizeMode)  {
	q.Drv(290000,290182,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setSectionHidden(int,bool)
func (q *QHeaderView) SetSectionHidden(logicalIndex int,hide bool)  {
	q.Drv(290000,290183,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&hide),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setSelection(QRect const&,QFlags<QItemSelectionModel::SelectionFlag>)
func (q *QHeaderView) SetSelection(rect *QRect,flags QItemSelectionModel_SelectionFlag)  {
	q.Drv(290000,290184,Native(rect),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setSortIndicator(int,Qt::SortOrder)
func (q *QHeaderView) SetSortIndicator(logicalIndex int,order Qt_SortOrder)  {
	q.Drv(290000,290185,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&order),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setSortIndicatorShown(bool)
func (q *QHeaderView) SetSortIndicatorShown(show bool)  {
	q.Drv(290000,290186,unsafe.Pointer(&show),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::setStretchLastSection(bool)
func (q *QHeaderView) SetStretchLastSection(stretch bool)  {
	q.Drv(290000,290187,unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::showSection(int)
func (q *QHeaderView) ShowSection(logicalIndex int)  {
	q.Drv(290000,290188,unsafe.Pointer(&logicalIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::sizeHint()
func (q *QHeaderView) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(290000,290189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QHeaderView::sortIndicatorOrder()
func (q *QHeaderView) SortIndicatorOrder() Qt_SortOrder {
	var __rv Qt_SortOrder
	q.Drv(290000,290190,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::sortIndicatorSection()
func (q *QHeaderView) SortIndicatorSection() int {
	var __rv int
	q.Drv(290000,290191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::stretchLastSection()
func (q *QHeaderView) StretchLastSection() bool {
	var __rv bool
	q.Drv(290000,290192,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::stretchSectionCount()
func (q *QHeaderView) StretchSectionCount() int {
	var __rv int
	q.Drv(290000,290193,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::swapSections(int,int)
func (q *QHeaderView) SwapSections(first int,second int)  {
	q.Drv(290000,290194,unsafe.Pointer(&first),unsafe.Pointer(&second),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::updateGeometries()
func (q *QHeaderView) UpdateGeometries()  {
	q.Drv(290000,290195,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::updateSection(int)
func (q *QHeaderView) UpdateSection(logicalIndex int)  {
	q.Drv(290000,290196,unsafe.Pointer(&logicalIndex),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QHeaderView::verticalOffset()
func (q *QHeaderView) VerticalOffset() int {
	var __rv int
	q.Drv(290000,290197,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::viewportEvent(QEvent*)
func (q *QHeaderView) ViewportEvent(e *QEvent) bool {
	var __rv bool
	q.Drv(290000,290198,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::visualIndex(int)
func (q *QHeaderView) VisualIndex(logicalIndex int) int {
	var __rv int
	q.Drv(290000,290199,unsafe.Pointer(&logicalIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::visualIndexAt(int)
func (q *QHeaderView) VisualIndexAt(position int) int {
	var __rv int
	q.Drv(290000,290200,unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QHeaderView::visualRect(QModelIndex const&)
func (q *QHeaderView) VisualRect(index *QModelIndex) *QRect {
	var __rv uintptr
	q.Drv(290000,290201,Native(index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QHeaderView::visualRegionForSelection(QItemSelection const&)
func (q *QHeaderView) VisualRegionForSelection(selection *QItemSelection) *QRegion {
	var __rv uintptr
	q.Drv(290000,290202,Native(selection),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRegion{}
	_rp.SetDriver(__rv,113000,true)
	return _rp
}	
