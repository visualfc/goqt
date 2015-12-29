// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFontMetrics : QFontMetrics
type QFontMetrics struct {
	BaseDrv
}
//QFontMetrics::QFontMetrics(QFont const&)	
func NewFontMetrics(value *QFont) *QFontMetrics {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),40000,40102,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetrics{}
	_p.SetDriver(__rv,40000,true)
	return _p
} 
//QFontMetrics::QFontMetrics(QFontMetrics const&)	
func NewFontMetricsCopy(value *QFontMetrics) *QFontMetrics {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),40000,40103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetrics{}
	_p.SetDriver(__rv,40000,true)
	return _p
} 
//QFontMetrics::QFontMetrics(QFont const&,QPaintDevice*)	
func NewFontMetricsWithFontPaintDevice(value2 *QFont,pd QPaintDeviceInterface) *QFontMetrics {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),40000,40104,Native(value2),unsafe.Pointer(new_pd_head(pd)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetrics{}
	_p.SetDriver(__rv,40000,true)
	return _p
} 
//QFontMetrics::ascent()
func (q *QFontMetrics) Ascent() int {
	var __rv int
	q.Drv(40000,40105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::averageCharWidth()
func (q *QFontMetrics) AverageCharWidth() int {
	var __rv int
	q.Drv(40000,40106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::boundingRect(QChar)
func (q *QFontMetrics) BoundingRect(value rune) *QRect {
	var __rv uintptr
	q.Drv(40000,40107,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFontMetrics::boundingRect(QString const&)
func (q *QFontMetrics) BoundingRectWithText(text string) *QRect {
	var __rv uintptr
	q.Drv(40000,40108,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFontMetrics::boundingRect(QRect const&,int,QString const&,int,int*)
func (q *QFontMetrics) BoundingRectWithRectFlagsTextTabstopsTabarray(r *QRect,flags int,text string,tabstops int,tabarray *int) *QRect {
	var __rv uintptr
	q.Drv(40000,40109,Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&text),unsafe.Pointer(&tabstops),unsafe.Pointer(&tabarray),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFontMetrics::boundingRect(int,int,int,int,int,QString const&,int,int*)
func (q *QFontMetrics) BoundingRectWithXYWidthHeightFlagsTextTabstopsTabarray(x int,y int,w int,h int,flags int,text string,tabstops int,tabarray *int) *QRect {
	var __rv uintptr
	q.Drv(40000,40110,unsafe.Pointer(&x),unsafe.Pointer(&y),unsafe.Pointer(&w),unsafe.Pointer(&h),unsafe.Pointer(&flags),unsafe.Pointer(&text),unsafe.Pointer(&tabstops),unsafe.Pointer(&tabarray),unsafe.Pointer(&__rv),nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFontMetrics::charWidth(QString const&,int)
func (q *QFontMetrics) CharWidth(str string,pos int) int {
	var __rv int
	q.Drv(40000,40111,unsafe.Pointer(&str),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::descent()
func (q *QFontMetrics) Descent() int {
	var __rv int
	q.Drv(40000,40112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::elidedText(QString const&,Qt::TextElideMode,int,int)
func (q *QFontMetrics) ElidedText(text string,mode Qt_TextElideMode,width int,flags int) string {
	var __rv string
	q.Drv(40000,40113,unsafe.Pointer(&text),unsafe.Pointer(&mode),unsafe.Pointer(&width),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::height()
func (q *QFontMetrics) Height() int {
	var __rv int
	q.Drv(40000,40114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::inFont(QChar)
func (q *QFontMetrics) InFont(value rune) bool {
	var __rv bool
	q.Drv(40000,40115,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::leading()
func (q *QFontMetrics) Leading() int {
	var __rv int
	q.Drv(40000,40116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::leftBearing(QChar)
func (q *QFontMetrics) LeftBearing(value rune) int {
	var __rv int
	q.Drv(40000,40117,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::lineSpacing()
func (q *QFontMetrics) LineSpacing() int {
	var __rv int
	q.Drv(40000,40118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::lineWidth()
func (q *QFontMetrics) LineWidth() int {
	var __rv int
	q.Drv(40000,40119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::maxWidth()
func (q *QFontMetrics) MaxWidth() int {
	var __rv int
	q.Drv(40000,40120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::minLeftBearing()
func (q *QFontMetrics) MinLeftBearing() int {
	var __rv int
	q.Drv(40000,40121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::minRightBearing()
func (q *QFontMetrics) MinRightBearing() int {
	var __rv int
	q.Drv(40000,40122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::overlinePos()
func (q *QFontMetrics) OverlinePos() int {
	var __rv int
	q.Drv(40000,40123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::rightBearing(QChar)
func (q *QFontMetrics) RightBearing(value rune) int {
	var __rv int
	q.Drv(40000,40124,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::size(int,QString const&,int,int*)
func (q *QFontMetrics) Size(flags int,str string,tabstops int,tabarray *int) *QSize {
	var __rv uintptr
	q.Drv(40000,40125,unsafe.Pointer(&flags),unsafe.Pointer(&str),unsafe.Pointer(&tabstops),unsafe.Pointer(&tabarray),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QFontMetrics::strikeOutPos()
func (q *QFontMetrics) StrikeOutPos() int {
	var __rv int
	q.Drv(40000,40126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::tightBoundingRect(QString const&)
func (q *QFontMetrics) TightBoundingRect(text string) *QRect {
	var __rv uintptr
	q.Drv(40000,40127,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QFontMetrics::underlinePos()
func (q *QFontMetrics) UnderlinePos() int {
	var __rv int
	q.Drv(40000,40128,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::width(QChar)
func (q *QFontMetrics) Width(value rune) int {
	var __rv int
	q.Drv(40000,40129,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::width(QString const&)
func (q *QFontMetrics) WidthWithString(value string) int {
	var __rv int
	q.Drv(40000,40130,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::width(QString const&,int)
func (q *QFontMetrics) WidthWithStringLen(value2 string,len int) int {
	var __rv int
	q.Drv(40000,40131,unsafe.Pointer(&value2),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::width(QString const&,int,int)
func (q *QFontMetrics) WidthWithStringLenFlags(value2 string,len int,flags int) int {
	var __rv int
	q.Drv(40000,40132,unsafe.Pointer(&value2),unsafe.Pointer(&len),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetrics::xHeight()
func (q *QFontMetrics) XHeight() int {
	var __rv int
	q.Drv(40000,40133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
