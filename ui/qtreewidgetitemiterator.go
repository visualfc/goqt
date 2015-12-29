// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTreeWidgetItemIterator_IteratorFlag - QTreeWidgetItemIterator::IteratorFlag
type QTreeWidgetItemIterator_IteratorFlag uint32
const (
	QTreeWidgetItemIterator_All QTreeWidgetItemIterator_IteratorFlag = 0x00000000
	QTreeWidgetItemIterator_Hidden QTreeWidgetItemIterator_IteratorFlag = 0x00000001
	QTreeWidgetItemIterator_NotHidden QTreeWidgetItemIterator_IteratorFlag = 0x00000002
	QTreeWidgetItemIterator_Selected QTreeWidgetItemIterator_IteratorFlag = 0x00000004
	QTreeWidgetItemIterator_Unselected QTreeWidgetItemIterator_IteratorFlag = 0x00000008
	QTreeWidgetItemIterator_Selectable QTreeWidgetItemIterator_IteratorFlag = 0x00000010
	QTreeWidgetItemIterator_NotSelectable QTreeWidgetItemIterator_IteratorFlag = 0x00000020
	QTreeWidgetItemIterator_DragEnabled QTreeWidgetItemIterator_IteratorFlag = 0x00000040
	QTreeWidgetItemIterator_DragDisabled QTreeWidgetItemIterator_IteratorFlag = 0x00000080
	QTreeWidgetItemIterator_DropEnabled QTreeWidgetItemIterator_IteratorFlag = 0x00000100
	QTreeWidgetItemIterator_DropDisabled QTreeWidgetItemIterator_IteratorFlag = 0x00000200
	QTreeWidgetItemIterator_HasChildren QTreeWidgetItemIterator_IteratorFlag = 0x00000400
	QTreeWidgetItemIterator_NoChildren QTreeWidgetItemIterator_IteratorFlag = 0x00000800
	QTreeWidgetItemIterator_Checked QTreeWidgetItemIterator_IteratorFlag = 0x00001000
	QTreeWidgetItemIterator_NotChecked QTreeWidgetItemIterator_IteratorFlag = 0x00002000
	QTreeWidgetItemIterator_Enabled QTreeWidgetItemIterator_IteratorFlag = 0x00004000
	QTreeWidgetItemIterator_Disabled QTreeWidgetItemIterator_IteratorFlag = 0x00008000
	QTreeWidgetItemIterator_Editable QTreeWidgetItemIterator_IteratorFlag = 0x00010000
	QTreeWidgetItemIterator_NotEditable QTreeWidgetItemIterator_IteratorFlag = 0x00020000
	QTreeWidgetItemIterator_UserFlag QTreeWidgetItemIterator_IteratorFlag = 0x01000000
)
//struct QTreeWidgetItemIterator : QTreeWidgetItemIterator
type QTreeWidgetItemIterator struct {
	BaseDrv
}
//QTreeWidgetItemIterator::QTreeWidgetItemIterator(QTreeWidgetItemIterator const&)	
func NewTreeWidgetItemIterator(it *QTreeWidgetItemIterator) *QTreeWidgetItemIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),178000,178102,Native(it),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItemIterator{}
	_p.SetDriver(__rv,178000,true)
	return _p
} 
//QTreeWidgetItemIterator::QTreeWidgetItemIterator(QTreeWidget*,QFlags<QTreeWidgetItemIterator::IteratorFlag>)	
func NewTreeWidgetItemIteratorWithWidgetFlags(widget *QTreeWidget,flags QTreeWidgetItemIterator_IteratorFlag) *QTreeWidgetItemIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),178000,178103,Native(widget),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItemIterator{}
	_p.SetDriver(__rv,178000,true)
	return _p
} 
//QTreeWidgetItemIterator::QTreeWidgetItemIterator(QTreeWidgetItem*,QFlags<QTreeWidgetItemIterator::IteratorFlag>)	
func NewTreeWidgetItemIteratorWithItemFlags(item *QTreeWidgetItem,flags QTreeWidgetItemIterator_IteratorFlag) *QTreeWidgetItemIterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),178000,178104,Native(item),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTreeWidgetItemIterator{}
	_p.SetDriver(__rv,178000,true)
	return _p
} 
