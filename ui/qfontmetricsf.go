// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFontMetricsF : QFontMetricsF
type QFontMetricsF struct {
	BaseDrv
}
//QFontMetricsF::QFontMetricsF(QFont const&)	
func NewFontMetricsF(value *QFont) *QFontMetricsF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),41000,41102,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetricsF{}
	_p.SetDriver(__rv,41000,true)
	return _p
} 
//QFontMetricsF::QFontMetricsF(QFontMetrics const&)	
func NewFontMetricsFWithFontmetrics(value *QFontMetrics) *QFontMetricsF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),41000,41103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetricsF{}
	_p.SetDriver(__rv,41000,true)
	return _p
} 
//QFontMetricsF::QFontMetricsF(QFontMetricsF const&)	
func NewFontMetricsFCopy(value *QFontMetricsF) *QFontMetricsF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),41000,41104,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetricsF{}
	_p.SetDriver(__rv,41000,true)
	return _p
} 
//QFontMetricsF::QFontMetricsF(QFont const&,QPaintDevice*)	
func NewFontMetricsFWithFontPaintDevice(value2 *QFont,pd QPaintDeviceInterface) *QFontMetricsF {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),41000,41105,Native(value2),unsafe.Pointer(new_pd_head(pd)),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontMetricsF{}
	_p.SetDriver(__rv,41000,true)
	return _p
} 
//QFontMetricsF::ascent()
func (q *QFontMetricsF) Ascent() float64 {
	var __rv float64
	q.Drv(41000,41106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::averageCharWidth()
func (q *QFontMetricsF) AverageCharWidth() float64 {
	var __rv float64
	q.Drv(41000,41107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::boundingRect(QChar)
func (q *QFontMetricsF) BoundingRect(value rune) *QRectF {
	var __rv uintptr
	q.Drv(41000,41108,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QFontMetricsF::boundingRect(QString const&)
func (q *QFontMetricsF) BoundingRectWithString(string string) *QRectF {
	var __rv uintptr
	q.Drv(41000,41109,unsafe.Pointer(&string),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QFontMetricsF::boundingRect(QRectF const&,int,QString const&,int,int*)
func (q *QFontMetricsF) BoundingRectFWithRectfFlagsStringTabstopsTabarray(r *QRectF,flags int,string string,tabstops int,tabarray *int) *QRectF {
	var __rv uintptr
	q.Drv(41000,41110,Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&string),unsafe.Pointer(&tabstops),unsafe.Pointer(&tabarray),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QFontMetricsF::descent()
func (q *QFontMetricsF) Descent() float64 {
	var __rv float64
	q.Drv(41000,41111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::elidedText(QString const&,Qt::TextElideMode,double,int)
func (q *QFontMetricsF) ElidedText(text string,mode Qt_TextElideMode,width float64,flags int) string {
	var __rv string
	q.Drv(41000,41112,unsafe.Pointer(&text),unsafe.Pointer(&mode),unsafe.Pointer(&width),unsafe.Pointer(&flags),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::height()
func (q *QFontMetricsF) Height() float64 {
	var __rv float64
	q.Drv(41000,41113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::inFont(QChar)
func (q *QFontMetricsF) InFont(value rune) bool {
	var __rv bool
	q.Drv(41000,41114,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::leading()
func (q *QFontMetricsF) Leading() float64 {
	var __rv float64
	q.Drv(41000,41115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::leftBearing(QChar)
func (q *QFontMetricsF) LeftBearing(value rune) float64 {
	var __rv float64
	q.Drv(41000,41116,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::lineSpacing()
func (q *QFontMetricsF) LineSpacing() float64 {
	var __rv float64
	q.Drv(41000,41117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::lineWidth()
func (q *QFontMetricsF) LineWidth() float64 {
	var __rv float64
	q.Drv(41000,41118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::maxWidth()
func (q *QFontMetricsF) MaxWidth() float64 {
	var __rv float64
	q.Drv(41000,41119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::minLeftBearing()
func (q *QFontMetricsF) MinLeftBearing() float64 {
	var __rv float64
	q.Drv(41000,41120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::minRightBearing()
func (q *QFontMetricsF) MinRightBearing() float64 {
	var __rv float64
	q.Drv(41000,41121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::overlinePos()
func (q *QFontMetricsF) OverlinePos() float64 {
	var __rv float64
	q.Drv(41000,41122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::rightBearing(QChar)
func (q *QFontMetricsF) RightBearing(value rune) float64 {
	var __rv float64
	q.Drv(41000,41123,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::size(int,QString const&,int,int*)
func (q *QFontMetricsF) Size(flags int,str string,tabstops int,tabarray *int) *QSizeF {
	var __rv uintptr
	q.Drv(41000,41124,unsafe.Pointer(&flags),unsafe.Pointer(&str),unsafe.Pointer(&tabstops),unsafe.Pointer(&tabarray),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QFontMetricsF::strikeOutPos()
func (q *QFontMetricsF) StrikeOutPos() float64 {
	var __rv float64
	q.Drv(41000,41125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::tightBoundingRect(QString const&)
func (q *QFontMetricsF) TightBoundingRect(text string) *QRectF {
	var __rv uintptr
	q.Drv(41000,41126,unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QFontMetricsF::underlinePos()
func (q *QFontMetricsF) UnderlinePos() float64 {
	var __rv float64
	q.Drv(41000,41127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::width(QChar)
func (q *QFontMetricsF) Width(value rune) float64 {
	var __rv float64
	q.Drv(41000,41128,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::width(QString const&)
func (q *QFontMetricsF) WidthWithString(string string) float64 {
	var __rv float64
	q.Drv(41000,41129,unsafe.Pointer(&string),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontMetricsF::xHeight()
func (q *QFontMetricsF) XHeight() float64 {
	var __rv float64
	q.Drv(41000,41130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
