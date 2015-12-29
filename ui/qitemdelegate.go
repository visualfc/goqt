// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QItemDelegate : QItemDelegate
type QItemDelegate struct {
	QAbstractItemDelegate
}
//QItemDelegate::QItemDelegate()	
func NewItemDelegate() *QItemDelegate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),295000,295102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemDelegate{}
	_p.SetDriver(__rv,295000,false)
	return _p
} 
//QItemDelegate::QItemDelegate(QObject*)	
func NewItemDelegateWithParent(parent QObjectInterface) *QItemDelegate {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),295000,295103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QItemDelegate{}
	_p.SetDriver(__rv,295000,false)
	return _p
} 
//QItemDelegate::eventFilter(QObject*,QEvent*)
func (q *QItemDelegate) EventFilter(object QObjectInterface,event *QEvent) bool {
	var __rv bool
	q.Drv(295000,295104,Native(object),Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemDelegate::hasClipping()
func (q *QItemDelegate) HasClipping() bool {
	var __rv bool
	q.Drv(295000,295105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QItemDelegate::itemEditorFactory()
func (q *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	var __rv uintptr
	q.Drv(295000,295106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QItemEditorFactory{}
	_rp.SetDriver(__rv,61000,true)
	return _rp
}	
//QItemDelegate::selected(QPixmap const&,QPalette const&,bool)
func (q *QItemDelegate) Selected(pixmap *QPixmap,palette *QPalette,enabled bool) *QPixmap {
	var __rv uintptr
	q.Drv(295000,295107,Native(pixmap),Native(palette),unsafe.Pointer(&enabled),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QItemDelegate::setClipping(bool)
func (q *QItemDelegate) SetClipping(clip bool)  {
	q.Drv(295000,295108,unsafe.Pointer(&clip),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemDelegate::setEditorData(QWidget*,QModelIndex const&)
func (q *QItemDelegate) SetEditorData(editor QWidgetInterface,index *QModelIndex)  {
	q.Drv(295000,295109,Native(editor),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemDelegate::setItemEditorFactory(QItemEditorFactory*)
func (q *QItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory)  {
	q.Drv(295000,295110,Native(factory),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemDelegate::setModelData(QWidget*,QAbstractItemModel*,QModelIndex const&)
func (q *QItemDelegate) SetModelData(editor QWidgetInterface,model QAbstractItemModelInterface,index *QModelIndex)  {
	q.Drv(295000,295111,Native(editor),Native(model),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QItemDelegate::textRectangle(QPainter*,QRect const&,QFont const&,QString const&)
func (q *QItemDelegate) TextRectangle(painter *QPainter,rect *QRect,font *QFont,text string) *QRect {
	var __rv uintptr
	q.Drv(295000,295112,Native(painter),Native(rect),Native(font),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
