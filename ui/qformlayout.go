// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFormLayout_FieldGrowthPolicy - QFormLayout::FieldGrowthPolicy
type QFormLayout_FieldGrowthPolicy uint32
const (
	QFormLayout_FieldsStayAtSizeHint QFormLayout_FieldGrowthPolicy = 0
	QFormLayout_ExpandingFieldsGrow QFormLayout_FieldGrowthPolicy = 1
	QFormLayout_AllNonFixedFieldsGrow QFormLayout_FieldGrowthPolicy = 2
)
//enum QFormLayout_RowWrapPolicy - QFormLayout::RowWrapPolicy
type QFormLayout_RowWrapPolicy uint32
const (
	QFormLayout_DontWrapRows QFormLayout_RowWrapPolicy = 0
	QFormLayout_WrapLongRows QFormLayout_RowWrapPolicy = 1
	QFormLayout_WrapAllRows QFormLayout_RowWrapPolicy = 2
)
//enum QFormLayout_ItemRole - QFormLayout::ItemRole
type QFormLayout_ItemRole uint32
const (
	QFormLayout_LabelRole QFormLayout_ItemRole = 0
	QFormLayout_FieldRole QFormLayout_ItemRole = 1
	QFormLayout_SpanningRole QFormLayout_ItemRole = 2
)
//struct QFormLayout : QFormLayout
type QFormLayout struct {
	QLayout
}
//QFormLayout::QFormLayout()	
func NewFormLayout() *QFormLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),244000,244102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFormLayout{}
	_p.SetDriver(__rv,244000,false)
	return _p
} 
//QFormLayout::QFormLayout(QWidget*)	
func NewFormLayoutWithParent(parent QWidgetInterface) *QFormLayout {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),244000,244103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFormLayout{}
	_p.SetDriver(__rv,244000,false)
	return _p
} 
//QFormLayout::addItem(QLayoutItem*)
func (q *QFormLayout) AddItem(item *QLayoutItem)  {
	q.Drv(244000,244104,Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QLayout*)
func (q *QFormLayout) AddRow(layout QLayoutInterface)  {
	q.Drv(244000,244105,Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QWidget*)
func (q *QFormLayout) AddRowWithWidget(widget QWidgetInterface)  {
	q.Drv(244000,244106,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QString const&,QLayout*)
func (q *QFormLayout) AddRowWithLabeltextLayout(labelText string,field QLayoutInterface)  {
	q.Drv(244000,244107,unsafe.Pointer(&labelText),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QString const&,QWidget*)
func (q *QFormLayout) AddRowWithLabeltextWidget(labelText string,field QWidgetInterface)  {
	q.Drv(244000,244108,unsafe.Pointer(&labelText),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QWidget*,QLayout*)
func (q *QFormLayout) AddRowWithWidgetLayout(label QWidgetInterface,field QLayoutInterface)  {
	q.Drv(244000,244109,Native(label),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::addRow(QWidget*,QWidget*)
func (q *QFormLayout) AddRowWithWidgetWidget(label QWidgetInterface,field QWidgetInterface)  {
	q.Drv(244000,244110,Native(label),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::count()
func (q *QFormLayout) Count() int {
	var __rv int
	q.Drv(244000,244111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::expandingDirections()
func (q *QFormLayout) ExpandingDirections() Qt_Orientation {
	var __rv Qt_Orientation
	q.Drv(244000,244112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::fieldGrowthPolicy()
func (q *QFormLayout) FieldGrowthPolicy() QFormLayout_FieldGrowthPolicy {
	var __rv QFormLayout_FieldGrowthPolicy
	q.Drv(244000,244113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::formAlignment()
func (q *QFormLayout) FormAlignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(244000,244114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::getItemPosition(int,int*,QFormLayout::ItemRole*)
func (q *QFormLayout) GetItemPosition(index int,rowPtr *int,rolePtr *QFormLayout_ItemRole)  {
	q.Drv(244000,244115,unsafe.Pointer(&index),unsafe.Pointer(&rowPtr),unsafe.Pointer(&rolePtr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::getLayoutPosition(QLayout*,int*,QFormLayout::ItemRole*)
func (q *QFormLayout) GetLayoutPosition(layout QLayoutInterface,rowPtr *int,rolePtr *QFormLayout_ItemRole)  {
	q.Drv(244000,244116,Native(layout),unsafe.Pointer(&rowPtr),unsafe.Pointer(&rolePtr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::getWidgetPosition(QWidget*,int*,QFormLayout::ItemRole*)
func (q *QFormLayout) GetWidgetPosition(widget QWidgetInterface,rowPtr *int,rolePtr *QFormLayout_ItemRole)  {
	q.Drv(244000,244117,Native(widget),unsafe.Pointer(&rowPtr),unsafe.Pointer(&rolePtr),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::hasHeightForWidth()
func (q *QFormLayout) HasHeightForWidth() bool {
	var __rv bool
	q.Drv(244000,244118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::heightForWidth(int)
func (q *QFormLayout) HeightForWidth(width int) int {
	var __rv int
	q.Drv(244000,244119,unsafe.Pointer(&width),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::horizontalSpacing()
func (q *QFormLayout) HorizontalSpacing() int {
	var __rv int
	q.Drv(244000,244120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::insertRow(int,QLayout*)
func (q *QFormLayout) InsertRowWithRowLayout(row int,layout QLayoutInterface)  {
	q.Drv(244000,244121,unsafe.Pointer(&row),Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::insertRow(int,QWidget*)
func (q *QFormLayout) InsertRowWithRowWidget(row int,widget QWidgetInterface)  {
	q.Drv(244000,244122,unsafe.Pointer(&row),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::insertRow(int,QString const&,QLayout*)
func (q *QFormLayout) InsertRowWithRowLabeltextLayout(row int,labelText string,field QLayoutInterface)  {
	q.Drv(244000,244123,unsafe.Pointer(&row),unsafe.Pointer(&labelText),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::insertRow(int,QString const&,QWidget*)
func (q *QFormLayout) InsertRowWithRowLabeltextWidget(row int,labelText string,field QWidgetInterface)  {
	q.Drv(244000,244124,unsafe.Pointer(&row),unsafe.Pointer(&labelText),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::insertRow(int,QWidget*,QLayout*)
func (q *QFormLayout) InsertRowWithRowWidgetLayout(row int,label QWidgetInterface,field QLayoutInterface)  {
	q.Drv(244000,244125,unsafe.Pointer(&row),Native(label),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::insertRow(int,QWidget*,QWidget*)
func (q *QFormLayout) InsertRowWithRowWidgetWidget(row int,label QWidgetInterface,field QWidgetInterface)  {
	q.Drv(244000,244126,unsafe.Pointer(&row),Native(label),Native(field),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::invalidate()
func (q *QFormLayout) Invalidate()  {
	q.Drv(244000,244127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::itemAt(int)
func (q *QFormLayout) ItemAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(244000,244128,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QFormLayout::itemAt(int,QFormLayout::ItemRole)
func (q *QFormLayout) ItemAtWithRowRole(row int,role QFormLayout_ItemRole) *QLayoutItem {
	var __rv uintptr
	q.Drv(244000,244129,unsafe.Pointer(&row),unsafe.Pointer(&role),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QFormLayout::labelAlignment()
func (q *QFormLayout) LabelAlignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(244000,244130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::labelForField(QLayout*)
func (q *QFormLayout) LabelForField(field QLayoutInterface) *QWidget {
	var __rv uintptr
	q.Drv(244000,244131,Native(field),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QFormLayout::labelForField(QWidget*)
func (q *QFormLayout) LabelForFieldWithWidget(field QWidgetInterface) *QWidget {
	var __rv uintptr
	q.Drv(244000,244132,Native(field),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QFormLayout::minimumSize()
func (q *QFormLayout) MinimumSize() *QSize {
	var __rv uintptr
	q.Drv(244000,244133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QFormLayout::rowCount()
func (q *QFormLayout) RowCount() int {
	var __rv int
	q.Drv(244000,244134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::rowWrapPolicy()
func (q *QFormLayout) RowWrapPolicy() QFormLayout_RowWrapPolicy {
	var __rv QFormLayout_RowWrapPolicy
	q.Drv(244000,244135,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::setFieldGrowthPolicy(QFormLayout::FieldGrowthPolicy)
func (q *QFormLayout) SetFieldGrowthPolicy(policy QFormLayout_FieldGrowthPolicy)  {
	q.Drv(244000,244136,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setFormAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QFormLayout) SetFormAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(244000,244137,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setGeometry(QRect const&)
func (q *QFormLayout) SetGeometry(rect *QRect)  {
	q.Drv(244000,244138,Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setHorizontalSpacing(int)
func (q *QFormLayout) SetHorizontalSpacing(spacing int)  {
	q.Drv(244000,244139,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setItem(int,QFormLayout::ItemRole,QLayoutItem*)
func (q *QFormLayout) SetItem(row int,role QFormLayout_ItemRole,item *QLayoutItem)  {
	q.Drv(244000,244140,unsafe.Pointer(&row),unsafe.Pointer(&role),Native(item),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setLabelAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QFormLayout) SetLabelAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(244000,244141,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setLayout(int,QFormLayout::ItemRole,QLayout*)
func (q *QFormLayout) SetLayout(row int,role QFormLayout_ItemRole,layout QLayoutInterface)  {
	q.Drv(244000,244142,unsafe.Pointer(&row),unsafe.Pointer(&role),Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setRowWrapPolicy(QFormLayout::RowWrapPolicy)
func (q *QFormLayout) SetRowWrapPolicy(policy QFormLayout_RowWrapPolicy)  {
	q.Drv(244000,244143,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setSpacing(int)
func (q *QFormLayout) SetSpacing(value int)  {
	q.Drv(244000,244144,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setVerticalSpacing(int)
func (q *QFormLayout) SetVerticalSpacing(spacing int)  {
	q.Drv(244000,244145,unsafe.Pointer(&spacing),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::setWidget(int,QFormLayout::ItemRole,QWidget*)
func (q *QFormLayout) SetWidget(row int,role QFormLayout_ItemRole,widget QWidgetInterface)  {
	q.Drv(244000,244146,unsafe.Pointer(&row),unsafe.Pointer(&role),Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFormLayout::sizeHint()
func (q *QFormLayout) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(244000,244147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QFormLayout::spacing()
func (q *QFormLayout) Spacing() int {
	var __rv int
	q.Drv(244000,244148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFormLayout::takeAt(int)
func (q *QFormLayout) TakeAt(index int) *QLayoutItem {
	var __rv uintptr
	q.Drv(244000,244149,unsafe.Pointer(&index),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLayoutItem{}
	_rp.SetDriver(__rv,66000,true)
	return _rp
}	
//QFormLayout::verticalSpacing()
func (q *QFormLayout) VerticalSpacing() int {
	var __rv int
	q.Drv(244000,244150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
