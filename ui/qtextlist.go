// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextList : QTextList
type QTextList struct {
	QTextBlockGroup
}
//QTextList::QTextList(QTextDocument*)	
func NewTextList(doc *QTextDocument) *QTextList {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),375000,375102,Native(doc),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextList{}
	_p.SetDriver(__rv,375000,false)
	return _p
} 
//QTextList::add(QTextBlock const&)
func (q *QTextList) Add(block *QTextBlock)  {
	q.Drv(375000,375103,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextList::count()
func (q *QTextList) Count() int {
	var __rv int
	q.Drv(375000,375104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextList::format()
func (q *QTextList) Format() *QTextListFormat {
	var __rv uintptr
	q.Drv(375000,375105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextListFormat{}
	_rp.SetDriver(__rv,162000,true)
	return _rp
}	
//QTextList::isEmpty()
func (q *QTextList) IsEmpty() bool {
	var __rv bool
	q.Drv(375000,375106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextList::item(int)
func (q *QTextList) Item(i int) *QTextBlock {
	var __rv uintptr
	q.Drv(375000,375107,unsafe.Pointer(&i),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextList::itemNumber(QTextBlock const&)
func (q *QTextList) ItemNumber(value *QTextBlock) int {
	var __rv int
	q.Drv(375000,375108,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextList::itemText(QTextBlock const&)
func (q *QTextList) ItemText(value *QTextBlock) string {
	var __rv string
	q.Drv(375000,375109,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextList::remove(QTextBlock const&)
func (q *QTextList) Remove(value *QTextBlock)  {
	q.Drv(375000,375110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextList::removeItem(int)
func (q *QTextList) RemoveItem(i int)  {
	q.Drv(375000,375111,unsafe.Pointer(&i),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextList::setFormat(QTextListFormat const&)
func (q *QTextList) SetFormat(format *QTextListFormat)  {
	q.Drv(375000,375112,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
