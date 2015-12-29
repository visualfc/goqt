// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextFrameiterator : QTextFrame::iterator
type QTextFrameiterator struct {
	BaseDrv
}
//QTextFrame::iterator::atEnd()
func (q *QTextFrameiterator) AtEnd() bool {
	var __rv bool
	q.Drv(153000,153102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFrame::iterator::currentBlock()
func (q *QTextFrameiterator) CurrentBlock() *QTextBlock {
	var __rv uintptr
	q.Drv(153000,153103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextFrame::iterator::currentFrame()
func (q *QTextFrameiterator) CurrentFrame() *QTextFrame {
	var __rv uintptr
	q.Drv(153000,153104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextFrame::iterator::iterator()	
func NewTextFrameiterator() *QTextFrameiterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),153000,153105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFrameiterator{}
	_p.SetDriver(__rv,153000,true)
	return _p
} 
//QTextFrame::iterator::iterator(QTextFrame::iterator const&)	
func NewTextFrameiteratorCopy(o *QTextFrameiterator) *QTextFrameiterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),153000,153106,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFrameiterator{}
	_p.SetDriver(__rv,153000,true)
	return _p
} 
//QTextFrame::iterator::parentFrame()
func (q *QTextFrameiterator) ParentFrame() *QTextFrame {
	var __rv uintptr
	q.Drv(153000,153107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
