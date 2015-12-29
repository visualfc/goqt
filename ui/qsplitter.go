// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSplitter : QSplitter
type QSplitter struct {
	QFrame
}
func (q *QSplitter) OnSplitterMoved(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(348000,348102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSplitter::QSplitter()	
func NewSplitter() *QSplitter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),348000,348103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplitter{}
	_p.SetDriver(__rv,348000,false)
	return _p
} 
//QSplitter::QSplitter(QWidget*)	
func NewSplitterWithParent(parent QWidgetInterface) *QSplitter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),348000,348104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplitter{}
	_p.SetDriver(__rv,348000,false)
	return _p
} 
//QSplitter::QSplitter(Qt::Orientation,QWidget*)	
func NewSplitterWithOrientationParent(value2 Qt_Orientation,parent QWidgetInterface) *QSplitter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),348000,348105,unsafe.Pointer(&value2),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplitter{}
	_p.SetDriver(__rv,348000,false)
	return _p
} 
//QSplitter::addWidget(QWidget*)
func (q *QSplitter) AddWidget(widget QWidgetInterface)  {
	q.Drv(348000,348106,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::changeEvent(QEvent*)
func (q *QSplitter) ChangeEvent(value *QEvent)  {
	q.Drv(348000,348107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::childEvent(QChildEvent*)
func (q *QSplitter) ChildEvent(value *QChildEvent)  {
	q.Drv(348000,348108,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::childrenCollapsible()
func (q *QSplitter) ChildrenCollapsible() bool {
	var __rv bool
	q.Drv(348000,348109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::closestLegalPosition(int,int)
func (q *QSplitter) ClosestLegalPosition(value2 int,value3 int) int {
	var __rv int
	q.Drv(348000,348110,unsafe.Pointer(&value2),unsafe.Pointer(&value3),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::count()
func (q *QSplitter) Count() int {
	var __rv int
	q.Drv(348000,348111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::createHandle()
func (q *QSplitter) CreateHandle() *QSplitterHandle {
	var __rv uintptr
	q.Drv(348000,348112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSplitterHandle{}
	_rp.SetDriver(__rv,349000,false)
	return _rp
}	
//QSplitter::event(QEvent*)
func (q *QSplitter) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(348000,348113,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::getRange(int,int*,int*)
func (q *QSplitter) GetRange(index int,value2 *int,value3 *int)  {
	q.Drv(348000,348114,unsafe.Pointer(&index),unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::handle(int)
func (q *QSplitter) Handle(index int) *QSplitterHandle {
	var __rv uintptr
	q.Drv(348000,348115,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSplitterHandle{}
	_rp.SetDriver(__rv,349000,false)
	return _rp
}	
//QSplitter::handleWidth()
func (q *QSplitter) HandleWidth() int {
	var __rv int
	q.Drv(348000,348116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::indexOf(QWidget*)
func (q *QSplitter) IndexOf(w QWidgetInterface) int {
	var __rv int
	q.Drv(348000,348117,Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::insertWidget(int,QWidget*)
func (q *QSplitter) InsertWidget(index int,widget QWidgetInterface)  {
	q.Drv(348000,348118,unsafe.Pointer(&index),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::isCollapsible(int)
func (q *QSplitter) IsCollapsible(index int) bool {
	var __rv bool
	q.Drv(348000,348119,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::minimumSizeHint()
func (q *QSplitter) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(348000,348120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSplitter::moveSplitter(int,int)
func (q *QSplitter) MoveSplitter(pos int,index int)  {
	q.Drv(348000,348121,unsafe.Pointer(&pos),unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::opaqueResize()
func (q *QSplitter) OpaqueResize() bool {
	var __rv bool
	q.Drv(348000,348122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::orientation()
func (q *QSplitter) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(348000,348123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::refresh()
func (q *QSplitter) Refresh()  {
	q.Drv(348000,348124,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::resizeEvent(QResizeEvent*)
func (q *QSplitter) ResizeEvent(value *QResizeEvent)  {
	q.Drv(348000,348125,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::restoreState(QByteArray const&)
func (q *QSplitter) RestoreState(state []byte) bool {
	var __rv bool
	q.Drv(348000,348126,unsafe.Pointer(&state),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::saveState()
func (q *QSplitter) SaveState() []byte {
	var __rv []byte
	q.Drv(348000,348127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::setChildrenCollapsible(bool)
func (q *QSplitter) SetChildrenCollapsible(value bool)  {
	q.Drv(348000,348128,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setCollapsible(int,bool)
func (q *QSplitter) SetCollapsible(index int,value2 bool)  {
	q.Drv(348000,348129,unsafe.Pointer(&index),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setHandleWidth(int)
func (q *QSplitter) SetHandleWidth(value int)  {
	q.Drv(348000,348130,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setOpaqueResize(bool)
func (q *QSplitter) SetOpaqueResize(opaque bool)  {
	q.Drv(348000,348131,unsafe.Pointer(&opaque),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setOrientation(Qt::Orientation)
func (q *QSplitter) SetOrientation(value Qt_Orientation)  {
	q.Drv(348000,348132,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setRubberBand(int)
func (q *QSplitter) SetRubberBand(position int)  {
	q.Drv(348000,348133,unsafe.Pointer(&position),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setSizes(QList<int> const&)
func (q *QSplitter) SetSizes(list []int)  {
	q.Drv(348000,348134,unsafe.Pointer(&list),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::setStretchFactor(int,int)
func (q *QSplitter) SetStretchFactor(index int,stretch int)  {
	q.Drv(348000,348135,unsafe.Pointer(&index),unsafe.Pointer(&stretch),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplitter::sizeHint()
func (q *QSplitter) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(348000,348136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QSplitter::sizes()
func (q *QSplitter) Sizes() []int {
	var __rv []int
	q.Drv(348000,348137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplitter::widget(int)
func (q *QSplitter) Widget(index int) *QWidget {
	var __rv uintptr
	q.Drv(348000,348138,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
