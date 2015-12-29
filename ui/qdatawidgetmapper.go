// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDataWidgetMapper_SubmitPolicy - QDataWidgetMapper::SubmitPolicy
type QDataWidgetMapper_SubmitPolicy uint32
const (
	QDataWidgetMapper_AutoSubmit QDataWidgetMapper_SubmitPolicy = 0
	QDataWidgetMapper_ManualSubmit QDataWidgetMapper_SubmitPolicy = 1
)
//struct QDataWidgetMapper : QDataWidgetMapper
type QDataWidgetMapper struct {
	QObject
}
func (q *QDataWidgetMapper) OnCurrentIndexChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(221000,221102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDataWidgetMapper::QDataWidgetMapper()	
func NewDataWidgetMapper() *QDataWidgetMapper {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),221000,221103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDataWidgetMapper{}
	_p.SetDriver(__rv,221000,false)
	return _p
} 
//QDataWidgetMapper::QDataWidgetMapper(QObject*)	
func NewDataWidgetMapperWithParent(parent QObjectInterface) *QDataWidgetMapper {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),221000,221104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDataWidgetMapper{}
	_p.SetDriver(__rv,221000,false)
	return _p
} 
//QDataWidgetMapper::addMapping(QWidget*,int)
func (q *QDataWidgetMapper) AddMappingWithWidgetSection(widget QWidgetInterface,section int)  {
	q.Drv(221000,221105,Native(widget),unsafe.Pointer(&section),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::addMapping(QWidget*,int,QByteArray const&)
func (q *QDataWidgetMapper) AddMappingWithWidgetSectionPropertyname(widget QWidgetInterface,section int,propertyName []byte)  {
	q.Drv(221000,221106,Native(widget),unsafe.Pointer(&section),unsafe.Pointer(&propertyName),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::clearMapping()
func (q *QDataWidgetMapper) ClearMapping()  {
	q.Drv(221000,221107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::currentIndex()
func (q *QDataWidgetMapper) CurrentIndex() int {
	var __rv int
	q.Drv(221000,221108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::itemDelegate()
func (q *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	var __rv uintptr
	q.Drv(221000,221109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemDelegate{}
	_rp.SetDriver(__rv,194000,false)
	return _rp
}	
//QDataWidgetMapper::mappedPropertyName(QWidget*)
func (q *QDataWidgetMapper) MappedPropertyName(widget QWidgetInterface) []byte {
	var __rv []byte
	q.Drv(221000,221110,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::mappedSection(QWidget*)
func (q *QDataWidgetMapper) MappedSection(widget QWidgetInterface) int {
	var __rv int
	q.Drv(221000,221111,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::mappedWidgetAt(int)
func (q *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	var __rv uintptr
	q.Drv(221000,221112,unsafe.Pointer(&section),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDataWidgetMapper::model()
func (q *QDataWidgetMapper) Model() *QAbstractItemModel {
	var __rv uintptr
	q.Drv(221000,221113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractItemModel{}
	_rp.SetDriver(__rv,195000,false)
	return _rp
}	
//QDataWidgetMapper::orientation()
func (q *QDataWidgetMapper) Orientation() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(221000,221114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::removeMapping(QWidget*)
func (q *QDataWidgetMapper) RemoveMapping(widget QWidgetInterface)  {
	q.Drv(221000,221115,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::revert()
func (q *QDataWidgetMapper) Revert()  {
	q.Drv(221000,221116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::rootIndex()
func (q *QDataWidgetMapper) RootIndex() *QModelIndex {
	var __rv uintptr
	q.Drv(221000,221117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QModelIndex{}
	_rp.SetDriver(__rv,79000,true)
	return _rp
}	
//QDataWidgetMapper::setCurrentIndex(int)
func (q *QDataWidgetMapper) SetCurrentIndex(index int)  {
	q.Drv(221000,221118,unsafe.Pointer(&index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setCurrentModelIndex(QModelIndex const&)
func (q *QDataWidgetMapper) SetCurrentModelIndex(index *QModelIndex)  {
	q.Drv(221000,221119,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setItemDelegate(QAbstractItemDelegate*)
func (q *QDataWidgetMapper) SetItemDelegate(delegate *QAbstractItemDelegate)  {
	q.Drv(221000,221120,Native(delegate),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setModel(QAbstractItemModel*)
func (q *QDataWidgetMapper) SetModel(model QAbstractItemModelInterface)  {
	q.Drv(221000,221121,Native(model),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setOrientation(Qt::Orientation)
func (q *QDataWidgetMapper) SetOrientation(aOrientation Qt_Orientation)  {
	q.Drv(221000,221122,unsafe.Pointer(&aOrientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setRootIndex(QModelIndex const&)
func (q *QDataWidgetMapper) SetRootIndex(index *QModelIndex)  {
	q.Drv(221000,221123,Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::setSubmitPolicy(QDataWidgetMapper::SubmitPolicy)
func (q *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper_SubmitPolicy)  {
	q.Drv(221000,221124,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::submit()
func (q *QDataWidgetMapper) Submit() bool {
	var __rv bool
	q.Drv(221000,221125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::submitPolicy()
func (q *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper_SubmitPolicy {
	var __rv QDataWidgetMapper_SubmitPolicy
	q.Drv(221000,221126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDataWidgetMapper::toFirst()
func (q *QDataWidgetMapper) ToFirst()  {
	q.Drv(221000,221127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::toLast()
func (q *QDataWidgetMapper) ToLast()  {
	q.Drv(221000,221128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::toNext()
func (q *QDataWidgetMapper) ToNext()  {
	q.Drv(221000,221129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDataWidgetMapper::toPrevious()
func (q *QDataWidgetMapper) ToPrevious()  {
	q.Drv(221000,221130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
