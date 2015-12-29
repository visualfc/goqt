// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QProxyStyle : QProxyStyle
type QProxyStyle struct {
	QCommonStyle
}
//QProxyStyle::QProxyStyle()	
func NewProxyStyle() *QProxyStyle {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),330000,330102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProxyStyle{}
	_p.SetDriver(__rv,330000,false)
	return _p
} 
//QProxyStyle::QProxyStyle(QStyle*)	
func NewProxyStyleWithBasestyle(baseStyle *QStyle) *QProxyStyle {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),330000,330103,Native(baseStyle),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QProxyStyle{}
	_p.SetDriver(__rv,330000,false)
	return _p
} 
//QProxyStyle::baseStyle()
func (q *QProxyStyle) BaseStyle() *QStyle {
	var __rv uintptr
	q.Drv(330000,330104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
//QProxyStyle::drawItemPixmap(QPainter*,QRect const&,int,QPixmap const&)
func (q *QProxyStyle) DrawItemPixmap(painter *QPainter,rect *QRect,alignment int,pixmap *QPixmap)  {
	q.Drv(330000,330105,Native(painter),Native(rect),unsafe.Pointer(&alignment),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::drawItemText(QPainter*,QRect const&,int,QPalette const&,bool,QString const&,QPalette::ColorRole)
func (q *QProxyStyle) DrawItemText(painter *QPainter,rect *QRect,flags int,pal *QPalette,enabled bool,text string,textRole QPalette_ColorRole)  {
	q.Drv(330000,330106,Native(painter),Native(rect),unsafe.Pointer(&flags),Native(pal),unsafe.Pointer(&enabled),unsafe.Pointer(&text),unsafe.Pointer(&textRole),nil,nil,nil,nil,nil)
}	
//QProxyStyle::event(QEvent*)
func (q *QProxyStyle) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(330000,330107,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProxyStyle::itemPixmapRect(QRect const&,int,QPixmap const&)
func (q *QProxyStyle) ItemPixmapRect(r *QRect,flags int,pixmap *QPixmap) *QRect {
	var __rv uintptr
	q.Drv(330000,330108,Native(r),unsafe.Pointer(&flags),Native(pixmap),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QProxyStyle::itemTextRect(QFontMetrics const&,QRect const&,int,bool,QString const&)
func (q *QProxyStyle) ItemTextRect(fm *QFontMetrics,r *QRect,flags int,enabled bool,text string) *QRect {
	var __rv uintptr
	q.Drv(330000,330109,Native(fm),Native(r),unsafe.Pointer(&flags),unsafe.Pointer(&enabled),unsafe.Pointer(&text),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QProxyStyle::pixelMetric(QStyle::PixelMetric)
func (q *QProxyStyle) PixelMetric(metric QStyle_PixelMetric) int {
	var __rv int
	q.Drv(330000,330110,unsafe.Pointer(&metric),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProxyStyle::polish(QApplication*)
func (q *QProxyStyle) Polish(app *QApplication)  {
	q.Drv(330000,330111,Native(app),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::polish(QPalette&)
func (q *QProxyStyle) PolishWithPal(pal *QPalette)  {
	q.Drv(330000,330112,Native(pal),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::polish(QWidget*)
func (q *QProxyStyle) PolishWithWidget(widget QWidgetInterface)  {
	q.Drv(330000,330113,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::setBaseStyle(QStyle*)
func (q *QProxyStyle) SetBaseStyle(style *QStyle)  {
	q.Drv(330000,330114,Native(style),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::standardPalette()
func (q *QProxyStyle) StandardPalette() *QPalette {
	var __rv uintptr
	q.Drv(330000,330115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QProxyStyle::styleHint(QStyle::StyleHint)
func (q *QProxyStyle) StyleHint(hint QStyle_StyleHint) int {
	var __rv int
	q.Drv(330000,330116,unsafe.Pointer(&hint),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QProxyStyle::unpolish(QApplication*)
func (q *QProxyStyle) Unpolish(app *QApplication)  {
	q.Drv(330000,330117,Native(app),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QProxyStyle::unpolish(QWidget*)
func (q *QProxyStyle) UnpolishWithWidget(widget QWidgetInterface)  {
	q.Drv(330000,330118,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
