// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBlockiterator : QTextBlock::iterator
type QTextBlockiterator struct {
	BaseDrv
}
//QTextBlock::iterator::atEnd()
func (q *QTextBlockiterator) AtEnd() bool {
	var __rv bool
	q.Drv(138000,138102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBlock::iterator::fragment()
func (q *QTextBlockiterator) Fragment() *QTextFragment {
	var __rv uintptr
	q.Drv(138000,138103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFragment{}
	_rp.SetDriver(__rv,152000,true)
	return _rp
}	
//QTextBlock::iterator::iterator()	
func NewTextBlockiterator() *QTextBlockiterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),138000,138104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlockiterator{}
	_p.SetDriver(__rv,138000,true)
	return _p
} 
//QTextBlock::iterator::iterator(QTextBlock::iterator const&)	
func NewTextBlockiteratorCopy(o *QTextBlockiterator) *QTextBlockiterator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),138000,138105,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBlockiterator{}
	_p.SetDriver(__rv,138000,true)
	return _p
} 
