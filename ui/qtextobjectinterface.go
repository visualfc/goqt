// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextObjectInterface : QTextObjectInterface
type QTextObjectInterface struct {
	BaseDrv
}
//QTextObjectInterface::drawObject(QPainter*,QRectF const&,QTextDocument*,int,QTextFormat const&)
func (q *QTextObjectInterface) DrawObject(painter *QPainter,rect *QRectF,doc *QTextDocument,posInDocument int,format *QTextFormat)  {
	q.Drv(163000,163102,Native(painter),Native(rect),Native(doc),unsafe.Pointer(&posInDocument),Native(format),nil,nil,nil,nil,nil,nil,nil)
}	
//QTextObjectInterface::intrinsicSize(QTextDocument*,int,QTextFormat const&)
func (q *QTextObjectInterface) IntrinsicSize(doc *QTextDocument,posInDocument int,format *QTextFormat) *QSizeF {
	var __rv uintptr
	q.Drv(163000,163103,Native(doc),unsafe.Pointer(&posInDocument),Native(format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
