// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QAbstractTextDocumentLayoutPaintContext : QAbstractTextDocumentLayout::PaintContext
type QAbstractTextDocumentLayoutPaintContext struct {
	BaseDrv
}
//QAbstractTextDocumentLayout::PaintContext::PaintContext()	
func NewAbstractTextDocumentLayoutPaintContext() *QAbstractTextDocumentLayoutPaintContext {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),2000,2102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractTextDocumentLayoutPaintContext{}
	_p.SetDriver(__rv,2000,true)
	return _p
} 
