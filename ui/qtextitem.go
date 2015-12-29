// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextItem_RenderFlag - QTextItem::RenderFlag
type QTextItem_RenderFlag uint32
const (
	QTextItem_RightToLeft QTextItem_RenderFlag = 0x1
	QTextItem_Overline QTextItem_RenderFlag = 0x10
	QTextItem_Underline QTextItem_RenderFlag = 0x20
	QTextItem_StrikeOut QTextItem_RenderFlag = 0x40
	QTextItem_Dummy QTextItem_RenderFlag = 0xffffffff
)
//struct QTextItem : QTextItem
type QTextItem struct {
	BaseDrv
}
//QTextItem::QTextItem()	
func NewTextItem() *QTextItem {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),157000,157102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextItem{}
	_p.SetDriver(__rv,157000,true)
	return _p
} 
//QTextItem::ascent()
func (q *QTextItem) Ascent() float64 {
	var __rv float64
	q.Drv(157000,157103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextItem::descent()
func (q *QTextItem) Descent() float64 {
	var __rv float64
	q.Drv(157000,157104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextItem::font()
func (q *QTextItem) Font() *QFont {
	var __rv uintptr
	q.Drv(157000,157105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTextItem::renderFlags()
func (q *QTextItem) RenderFlags() QTextItem_RenderFlag {
	var __rv QTextItem_RenderFlag
	q.Drv(157000,157106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextItem::text()
func (q *QTextItem) Text() string {
	var __rv string
	q.Drv(157000,157107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextItem::width()
func (q *QTextItem) Width() float64 {
	var __rv float64
	q.Drv(157000,157108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
