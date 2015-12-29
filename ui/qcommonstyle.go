// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QCommonStyle : QCommonStyle
type QCommonStyle struct {
	QStyle
}
//QCommonStyle::QCommonStyle()	
func NewCommonStyle() *QCommonStyle {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),219000,219102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QCommonStyle{}
	_p.SetDriver(__rv,219000,false)
	return _p
} 
//QCommonStyle::pixelMetric(QStyle::PixelMetric)
func (q *QCommonStyle) PixelMetric(m QStyle_PixelMetric) int {
	var __rv int
	q.Drv(219000,219103,unsafe.Pointer(&m),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCommonStyle::polish(QApplication*)
func (q *QCommonStyle) Polish(app *QApplication)  {
	q.Drv(219000,219104,Native(app),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommonStyle::polish(QPalette&)
func (q *QCommonStyle) PolishWithPalette(value *QPalette)  {
	q.Drv(219000,219105,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommonStyle::polish(QWidget*)
func (q *QCommonStyle) PolishWithWidget(widget QWidgetInterface)  {
	q.Drv(219000,219106,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommonStyle::standardPixmap(QStyle::StandardPixmap)
func (q *QCommonStyle) StandardPixmap(sp QStyle_StandardPixmap) *QPixmap {
	var __rv uintptr
	q.Drv(219000,219107,unsafe.Pointer(&sp),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QCommonStyle::styleHint(QStyle::StyleHint)
func (q *QCommonStyle) StyleHint(sh QStyle_StyleHint) int {
	var __rv int
	q.Drv(219000,219108,unsafe.Pointer(&sh),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QCommonStyle::unpolish(QApplication*)
func (q *QCommonStyle) Unpolish(application *QApplication)  {
	q.Drv(219000,219109,Native(application),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QCommonStyle::unpolish(QWidget*)
func (q *QCommonStyle) UnpolishWithWidget(widget QWidgetInterface)  {
	q.Drv(219000,219110,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
