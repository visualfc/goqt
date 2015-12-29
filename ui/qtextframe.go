// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextFrame : QTextFrame
type QTextFrame struct {
	QTextObject
}
//QTextFrame::QTextFrame(QTextDocument*)	
func NewTextFrame(doc *QTextDocument) *QTextFrame {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),374000,374102,Native(doc),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFrame{}
	_p.SetDriver(__rv,374000,false)
	return _p
} 
//QTextFrame::begin()
func (q *QTextFrame) Begin() *QTextFrameiterator {
	var __rv uintptr
	q.Drv(374000,374103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameiterator{}
	_rp.SetDriver(__rv,153000,true)
	return _rp
}	
//QTextFrame::childFrames()
func (q *QTextFrame) ChildFrames() []*QTextFrame {
	var __rv []*QTextFrame
	q.Drv(374000,374104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrame::end()
func (q *QTextFrame) End() *QTextFrameiterator {
	var __rv uintptr
	q.Drv(374000,374105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameiterator{}
	_rp.SetDriver(__rv,153000,true)
	return _rp
}	
//QTextFrame::firstCursorPosition()
func (q *QTextFrame) FirstCursorPosition() *QTextCursor {
	var __rv uintptr
	q.Drv(374000,374106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextFrame::firstPosition()
func (q *QTextFrame) FirstPosition() int {
	var __rv int
	q.Drv(374000,374107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrame::frameFormat()
func (q *QTextFrame) FrameFormat() *QTextFrameFormat {
	var __rv uintptr
	q.Drv(374000,374108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrameFormat{}
	_rp.SetDriver(__rv,154000,true)
	return _rp
}	
//QTextFrame::lastCursorPosition()
func (q *QTextFrame) LastCursorPosition() *QTextCursor {
	var __rv uintptr
	q.Drv(374000,374109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextFrame::lastPosition()
func (q *QTextFrame) LastPosition() int {
	var __rv int
	q.Drv(374000,374110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrame::parentFrame()
func (q *QTextFrame) ParentFrame() *QTextFrame {
	var __rv uintptr
	q.Drv(374000,374111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextFrame::setFrameFormat(QTextFrameFormat const&)
func (q *QTextFrame) SetFrameFormat(format *QTextFrameFormat)  {
	q.Drv(374000,374112,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
