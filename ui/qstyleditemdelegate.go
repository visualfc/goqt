// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStyledItemDelegate : QStyledItemDelegate
type QStyledItemDelegate struct {
	QAbstractItemDelegate
}
//QStyledItemDelegate::QStyledItemDelegate()	
func NewStyledItemDelegate() *QStyledItemDelegate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),358000,358102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyledItemDelegate{}
	_p.SetDriver(__rv,358000,false)
	return _p
} 
//QStyledItemDelegate::QStyledItemDelegate(QObject*)	
func NewStyledItemDelegateWithParent(parent QObjectInterface) *QStyledItemDelegate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),358000,358103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStyledItemDelegate{}
	_p.SetDriver(__rv,358000,false)
	return _p
} 
//QStyledItemDelegate::displayText(QVariant const&,QLocale const&)
func (q *QStyledItemDelegate) DisplayText(value *QVariant,locale *QLocale) string {
	var __rv string
	q.Drv(358000,358104,Native(value),Native(locale),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyledItemDelegate::eventFilter(QObject*,QEvent*)
func (q *QStyledItemDelegate) EventFilter(object QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(358000,358105,Native(object),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStyledItemDelegate::itemEditorFactory()
func (q *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	var __rv uintptr
	q.Drv(358000,358106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemEditorFactory{}
	_rp.SetDriver(__rv,61000,true)
	return _rp
}	
//QStyledItemDelegate::setEditorData(QWidget*,QModelIndex const&)
func (q *QStyledItemDelegate) SetEditorData(editor QWidgetInterface,index *QModelIndex)  {
	q.Drv(358000,358107,Native(editor),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyledItemDelegate::setItemEditorFactory(QItemEditorFactory*)
func (q *QStyledItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory)  {
	q.Drv(358000,358108,Native(factory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStyledItemDelegate::setModelData(QWidget*,QAbstractItemModel*,QModelIndex const&)
func (q *QStyledItemDelegate) SetModelData(editor QWidgetInterface,model QAbstractItemModelInterface,index *QModelIndex)  {
	q.Drv(358000,358109,Native(editor),Native(model),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
