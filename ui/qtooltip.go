// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QToolTip : QToolTip
type QToolTip struct {
	BaseDrv
}
//QToolTip::font()	
func QToolTipFont() *QFont {
	var __rv uintptr
	DirectQtDrv(nil,173000,173102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QToolTip::font()
func (q *QToolTip) Font() *QFont {
	var __rv uintptr
	q.Drv(173000,173102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QToolTip::hideText()	
func QToolTipHideText()  {
	DirectQtDrv(nil,173000,173103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::hideText()
func (q *QToolTip) HideText()  {
	q.Drv(173000,173103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::isVisible()	
func QToolTipIsVisible() bool {
	var __rv bool
	DirectQtDrv(nil,173000,173104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolTip::isVisible()
func (q *QToolTip) IsVisible() bool {
	var __rv bool
	q.Drv(173000,173104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolTip::palette()	
func QToolTipPalette() *QPalette {
	var __rv uintptr
	DirectQtDrv(nil,173000,173105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QToolTip::palette()
func (q *QToolTip) Palette() *QPalette {
	var __rv uintptr
	q.Drv(173000,173105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPalette{}
	_rp.SetDriver(__rv,91000,true)
	return _rp
}	
//QToolTip::setFont(QFont const&)	
func QToolTipSetFont(value *QFont)  {
	DirectQtDrv(nil,173000,173106,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::setFont(QFont const&)
func (q *QToolTip) SetFont(value *QFont)  {
	q.Drv(173000,173106,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::setPalette(QPalette const&)	
func QToolTipSetPalette(value *QPalette)  {
	DirectQtDrv(nil,173000,173107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::setPalette(QPalette const&)
func (q *QToolTip) SetPalette(value *QPalette)  {
	q.Drv(173000,173107,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::showText(QPoint const&,QString const&,QWidget*)	
func QToolTipShowTextWithPosTextWidget(pos *QPoint,text string,w QWidgetInterface)  {
	DirectQtDrv(nil,173000,173108,Native(pos),unsafe.Pointer(&text),Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::showText(QPoint const&,QString const&,QWidget*)
func (q *QToolTip) ShowTextWithPosTextWidget(pos *QPoint,text string,w QWidgetInterface)  {
	q.Drv(173000,173108,Native(pos),unsafe.Pointer(&text),Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::showText(QPoint const&,QString const&,QWidget*,QRect const&)	
func QToolTipShowTextWithPosTextWidgetRect(pos *QPoint,text string,w QWidgetInterface,rect *QRect)  {
	DirectQtDrv(nil,173000,173109,Native(pos),unsafe.Pointer(&text),Native(w),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::showText(QPoint const&,QString const&,QWidget*,QRect const&)
func (q *QToolTip) ShowTextWithPosTextWidgetRect(pos *QPoint,text string,w QWidgetInterface,rect *QRect)  {
	q.Drv(173000,173109,Native(pos),unsafe.Pointer(&text),Native(w),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QToolTip::text()	
func QToolTipText() string {
	var __rv string
	DirectQtDrv(nil,173000,173110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QToolTip::text()
func (q *QToolTip) Text() string {
	var __rv string
	q.Drv(173000,173110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
