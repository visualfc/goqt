// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QStaticText_PerformanceHint - QStaticText::PerformanceHint
type QStaticText_PerformanceHint uint32
const (
	QStaticText_ModerateCaching QStaticText_PerformanceHint = 0
	QStaticText_AggressiveCaching QStaticText_PerformanceHint = 1
)
//struct QStaticText : QStaticText
type QStaticText struct {
	BaseDrv
}
//QStaticText::QStaticText()	
func NewStaticText() *QStaticText {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),126000,126102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStaticText{}
	_p.SetDriver(__rv,126000,true)
	return _p
} 
//QStaticText::QStaticText(QStaticText const&)	
func NewStaticTextCopy(other *QStaticText) *QStaticText {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),126000,126103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStaticText{}
	_p.SetDriver(__rv,126000,true)
	return _p
} 
//QStaticText::QStaticText(QString const&)	
func NewStaticTextWithText(text string) *QStaticText {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),126000,126104,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStaticText{}
	_p.SetDriver(__rv,126000,true)
	return _p
} 
//QStaticText::performanceHint()
func (q *QStaticText) PerformanceHint() QStaticText_PerformanceHint {
	var __rv QStaticText_PerformanceHint
	q.Drv(126000,126105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStaticText::prepare()
func (q *QStaticText) Prepare()  {
	q.Drv(126000,126106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::prepare(QTransform const&,QFont const&)
func (q *QStaticText) PrepareWithTransformFont(matrix *QTransform,font *QFont)  {
	q.Drv(126000,126107,Native(matrix),Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::setPerformanceHint(QStaticText::PerformanceHint)
func (q *QStaticText) SetPerformanceHint(performanceHint QStaticText_PerformanceHint)  {
	q.Drv(126000,126108,unsafe.Pointer(&performanceHint),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::setText(QString const&)
func (q *QStaticText) SetText(text string)  {
	q.Drv(126000,126109,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::setTextFormat(Qt::TextFormat)
func (q *QStaticText) SetTextFormat(textFormat Qt_TextFormat)  {
	q.Drv(126000,126110,unsafe.Pointer(&textFormat),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::setTextOption(QTextOption const&)
func (q *QStaticText) SetTextOption(textOption *QTextOption)  {
	q.Drv(126000,126111,Native(textOption),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::setTextWidth(double)
func (q *QStaticText) SetTextWidth(textWidth float64)  {
	q.Drv(126000,126112,unsafe.Pointer(&textWidth),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStaticText::size()
func (q *QStaticText) Size() *QSizeF {
	var __rv uintptr
	q.Drv(126000,126113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QStaticText::text()
func (q *QStaticText) Text() string {
	var __rv string
	q.Drv(126000,126114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStaticText::textFormat()
func (q *QStaticText) TextFormat() Qt_TextFormat {
	var __rv Qt_TextFormat
	q.Drv(126000,126115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStaticText::textOption()
func (q *QStaticText) TextOption() *QTextOption {
	var __rv uintptr
	q.Drv(126000,126116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextOption{}
	_rp.SetDriver(__rv,164000,true)
	return _rp
}	
//QStaticText::textWidth()
func (q *QStaticText) TextWidth() float64 {
	var __rv float64
	q.Drv(126000,126117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
