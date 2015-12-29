// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractItemDelegate_EndEditHint - QAbstractItemDelegate::EndEditHint
type QAbstractItemDelegate_EndEditHint uint32
const (
	QAbstractItemDelegate_NoHint QAbstractItemDelegate_EndEditHint = 0
	QAbstractItemDelegate_EditNextItem QAbstractItemDelegate_EndEditHint = 1
	QAbstractItemDelegate_EditPreviousItem QAbstractItemDelegate_EndEditHint = 2
	QAbstractItemDelegate_SubmitModelCache QAbstractItemDelegate_EndEditHint = 3
	QAbstractItemDelegate_RevertModelCache QAbstractItemDelegate_EndEditHint = 4
)
//struct QAbstractItemDelegate : QAbstractItemDelegate
type QAbstractItemDelegate struct {
	QObject
}
func (q *QAbstractItemDelegate) OnSizeHintChanged(fn func(*QModelIndex)) uintptr {
	var __rv uintptr
	q.Drv(194000,194102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemDelegate) OnCommitData(fn func(*QWidget)) uintptr {
	var __rv uintptr
	q.Drv(194000,194103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QAbstractItemDelegate) OnCloseEditor(fn func(*QWidget,QAbstractItemDelegate_EndEditHint)) uintptr {
	var __rv uintptr
	q.Drv(194000,194104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractItemDelegate::elidedText(QFontMetrics const&,int,Qt::TextElideMode,QString const&)	
func QAbstractItemDelegateElidedText(fontMetrics *QFontMetrics,width int,mode Qt_TextElideMode,text string) string {
	var __rv string
	DirectQtDrv(nil,194000,194105,Native(fontMetrics),unsafe.Pointer(&width),unsafe.Pointer(&mode),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemDelegate::elidedText(QFontMetrics const&,int,Qt::TextElideMode,QString const&)
func (q *QAbstractItemDelegate) ElidedText(fontMetrics *QFontMetrics,width int,mode Qt_TextElideMode,text string) string {
	var __rv string
	q.Drv(194000,194105,Native(fontMetrics),unsafe.Pointer(&width),unsafe.Pointer(&mode),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractItemDelegate::setEditorData(QWidget*,QModelIndex const&)
func (q *QAbstractItemDelegate) SetEditorData(editor QWidgetInterface,index *QModelIndex)  {
	q.Drv(194000,194106,Native(editor),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractItemDelegate::setModelData(QWidget*,QAbstractItemModel*,QModelIndex const&)
func (q *QAbstractItemDelegate) SetModelData(editor QWidgetInterface,model QAbstractItemModelInterface,index *QModelIndex)  {
	q.Drv(194000,194107,Native(editor),Native(model),Native(index),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
