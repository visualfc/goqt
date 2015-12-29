// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBlock : QTextBlock
type QTextBlock struct {
	BaseDrv
}
//QTextBlock::QTextBlock()	
func NewTextBlock() *QTextBlock {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),137000,137102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlock{}
	_p.SetDriver(__rv,137000,true)
	return _p
} 
//QTextBlock::QTextBlock(QTextBlock const&)	
func NewTextBlockCopy(o *QTextBlock) *QTextBlock {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),137000,137103,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlock{}
	_p.SetDriver(__rv,137000,true)
	return _p
} 
//QTextBlock::begin()
func (q *QTextBlock) Begin() *QTextBlockiterator {
	var __rv uintptr
	q.Drv(137000,137104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockiterator{}
	_rp.SetDriver(__rv,138000,true)
	return _rp
}	
//QTextBlock::blockFormat()
func (q *QTextBlock) BlockFormat() *QTextBlockFormat {
	var __rv uintptr
	q.Drv(137000,137105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockFormat{}
	_rp.SetDriver(__rv,139000,true)
	return _rp
}	
//QTextBlock::blockFormatIndex()
func (q *QTextBlock) BlockFormatIndex() int {
	var __rv int
	q.Drv(137000,137106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::blockNumber()
func (q *QTextBlock) BlockNumber() int {
	var __rv int
	q.Drv(137000,137107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::charFormat()
func (q *QTextBlock) CharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(137000,137108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextBlock::charFormatIndex()
func (q *QTextBlock) CharFormatIndex() int {
	var __rv int
	q.Drv(137000,137109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::clearLayout()
func (q *QTextBlock) ClearLayout()  {
	q.Drv(137000,137110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::contains(int)
func (q *QTextBlock) Contains(position int) bool {
	var __rv bool
	q.Drv(137000,137111,unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::document()
func (q *QTextBlock) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(137000,137112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextBlock::end()
func (q *QTextBlock) End() *QTextBlockiterator {
	var __rv uintptr
	q.Drv(137000,137113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockiterator{}
	_rp.SetDriver(__rv,138000,true)
	return _rp
}	
//QTextBlock::firstLineNumber()
func (q *QTextBlock) FirstLineNumber() int {
	var __rv int
	q.Drv(137000,137114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::fragmentIndex()
func (q *QTextBlock) FragmentIndex() int {
	var __rv int
	q.Drv(137000,137115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::isValid()
func (q *QTextBlock) IsValid() bool {
	var __rv bool
	q.Drv(137000,137116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::isVisible()
func (q *QTextBlock) IsVisible() bool {
	var __rv bool
	q.Drv(137000,137117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::layout()
func (q *QTextBlock) Layout() *QTextLayout {
	var __rv uintptr
	q.Drv(137000,137118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextLayout{}
	_rp.SetDriver(__rv,158000,true)
	return _rp
}	
//QTextBlock::length()
func (q *QTextBlock) Length() int {
	var __rv int
	q.Drv(137000,137119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::lineCount()
func (q *QTextBlock) LineCount() int {
	var __rv int
	q.Drv(137000,137120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::next()
func (q *QTextBlock) Next() *QTextBlock {
	var __rv uintptr
	q.Drv(137000,137121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextBlock::position()
func (q *QTextBlock) Position() int {
	var __rv int
	q.Drv(137000,137122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::previous()
func (q *QTextBlock) Previous() *QTextBlock {
	var __rv uintptr
	q.Drv(137000,137123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextBlock::revision()
func (q *QTextBlock) Revision() int {
	var __rv int
	q.Drv(137000,137124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::setLineCount(int)
func (q *QTextBlock) SetLineCount(count int)  {
	q.Drv(137000,137125,unsafe.Pointer(&count),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::setRevision(int)
func (q *QTextBlock) SetRevision(rev int)  {
	q.Drv(137000,137126,unsafe.Pointer(&rev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::setUserData(QTextBlockUserData*)
func (q *QTextBlock) SetUserData(data *QTextBlockUserData)  {
	q.Drv(137000,137127,Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::setUserState(int)
func (q *QTextBlock) SetUserState(state int)  {
	q.Drv(137000,137128,unsafe.Pointer(&state),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::setVisible(bool)
func (q *QTextBlock) SetVisible(visible bool)  {
	q.Drv(137000,137129,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBlock::text()
func (q *QTextBlock) Text() string {
	var __rv string
	q.Drv(137000,137130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::textDirection()
func (q *QTextBlock) TextDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(137000,137131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::textList()
func (q *QTextBlock) TextList() *QTextList {
	var __rv uintptr
	q.Drv(137000,137132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextBlock::userData()
func (q *QTextBlock) UserData() *QTextBlockUserData {
	var __rv uintptr
	q.Drv(137000,137133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockUserData{}
	_rp.SetDriver(__rv,140000,true)
	return _rp
}	
//QTextBlock::userState()
func (q *QTextBlock) UserState() int {
	var __rv int
	q.Drv(137000,137134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
