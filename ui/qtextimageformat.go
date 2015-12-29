// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextImageFormat : QTextImageFormat
type QTextImageFormat struct {
	QTextCharFormat
}
//QTextImageFormat::QTextImageFormat()	
func NewTextImageFormat() *QTextImageFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),155000,155102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextImageFormat{}
	_p.SetDriver(__rv,155000,true)
	return _p
} 
//QTextImageFormat::height()
func (q *QTextImageFormat) Height() float64 {
	var __rv float64
	q.Drv(155000,155103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextImageFormat::isValid()
func (q *QTextImageFormat) IsValid() bool {
	var __rv bool
	q.Drv(155000,155104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextImageFormat::name()
func (q *QTextImageFormat) Name() string {
	var __rv string
	q.Drv(155000,155105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextImageFormat::setHeight(double)
func (q *QTextImageFormat) SetHeight(height float64)  {
	q.Drv(155000,155106,unsafe.Pointer(&height),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextImageFormat::setName(QString const&)
func (q *QTextImageFormat) SetName(name string)  {
	q.Drv(155000,155107,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextImageFormat::setWidth(double)
func (q *QTextImageFormat) SetWidth(width float64)  {
	q.Drv(155000,155108,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextImageFormat::width()
func (q *QTextImageFormat) Width() float64 {
	var __rv float64
	q.Drv(155000,155109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
