// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextObject : QTextObject
type QTextObject struct {
	QObject
}
//QTextObject::document()
func (q *QTextObject) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(376000,376102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextObject::format()
func (q *QTextObject) Format() *QTextFormat {
	var __rv uintptr
	q.Drv(376000,376103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFormat{}
	_rp.SetDriver(__rv,151000,true)
	return _rp
}	
//QTextObject::formatIndex()
func (q *QTextObject) FormatIndex() int {
	var __rv int
	q.Drv(376000,376104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextObject::objectIndex()
func (q *QTextObject) ObjectIndex() int {
	var __rv int
	q.Drv(376000,376105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextObject::setFormat(QTextFormat const&)
func (q *QTextObject) SetFormat(format *QTextFormat)  {
	q.Drv(376000,376106,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
