// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextListFormat_Style - QTextListFormat::Style
type QTextListFormat_Style int32
const (
	QTextListFormat_ListDisc QTextListFormat_Style = -1
	QTextListFormat_ListCircle QTextListFormat_Style = -2
	QTextListFormat_ListSquare QTextListFormat_Style = -3
	QTextListFormat_ListDecimal QTextListFormat_Style = -4
	QTextListFormat_ListLowerAlpha QTextListFormat_Style = -5
	QTextListFormat_ListUpperAlpha QTextListFormat_Style = -6
	QTextListFormat_ListLowerRoman QTextListFormat_Style = -7
	QTextListFormat_ListUpperRoman QTextListFormat_Style = -8
	QTextListFormat_ListStyleUndefined QTextListFormat_Style = 0
)
//struct QTextListFormat : QTextListFormat
type QTextListFormat struct {
	QTextFormat
}
//QTextListFormat::QTextListFormat()	
func NewTextListFormat() *QTextListFormat {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),162000,162102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextListFormat{}
	_p.SetDriver(__rv,162000,true)
	return _p
} 
//QTextListFormat::indent()
func (q *QTextListFormat) Indent() int {
	var __rv int
	q.Drv(162000,162103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextListFormat::isValid()
func (q *QTextListFormat) IsValid() bool {
	var __rv bool
	q.Drv(162000,162104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextListFormat::setIndent(int)
func (q *QTextListFormat) SetIndent(indent int)  {
	q.Drv(162000,162105,unsafe.Pointer(&indent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextListFormat::setStyle(QTextListFormat::Style)
func (q *QTextListFormat) SetStyle(style QTextListFormat_Style)  {
	q.Drv(162000,162106,unsafe.Pointer(&style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextListFormat::style()
func (q *QTextListFormat) Style() QTextListFormat_Style {
	var __rv QTextListFormat_Style
	q.Drv(162000,162107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
