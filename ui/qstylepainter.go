// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStylePainter : QStylePainter
type QStylePainter struct {
	QPainter
}
//QStylePainter::QStylePainter()	
func NewStylePainter() *QStylePainter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),133000,133102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStylePainter{}
	_p.SetDriver(__rv,133000,true)
	return _p
} 
//QStylePainter::QStylePainter(QWidget*)	
func NewStylePainterWithWidget(w QWidgetInterface) *QStylePainter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),133000,133103,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStylePainter{}
	_p.SetDriver(__rv,133000,true)
	return _p
} 
//QStylePainter::QStylePainter(QPaintDevice*,QWidget*)	
func NewStylePainterWithPaintDeviceWidget(pd QPaintDeviceInterface,w QWidgetInterface) *QStylePainter {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),133000,133104,unsafe.Pointer(new_pd_head(pd)),Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStylePainter{}
	_p.SetDriver(__rv,133000,true)
	return _p
} 
//QStylePainter::begin(QWidget*)
func (q *QStylePainter) Begin(w QWidgetInterface) bool {
	var __rv bool
	q.Drv(133000,133105,Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStylePainter::begin(QPaintDevice*,QWidget*)
func (q *QStylePainter) BeginWithPaintDeviceWidget(pd QPaintDeviceInterface,w QWidgetInterface) bool {
	var __rv bool
	q.Drv(133000,133106,unsafe.Pointer(new_pd_head(pd)),Native(w),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStylePainter::drawItemPixmap(QRect const&,int,QPixmap const&)
func (q *QStylePainter) DrawItemPixmap(r *QRect,flags int,pixmap *QPixmap)  {
	q.Drv(133000,133107,Native(r),unsafe.Pointer(&flags),Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStylePainter::drawItemText(QRect const&,int,QPalette const&,bool,QString const&,QPalette::ColorRole)
func (q *QStylePainter) DrawItemText(r *QRect,flags int,pal *QPalette,enabled bool,text string,textRole QPalette_ColorRole)  {
	q.Drv(133000,133108,Native(r),unsafe.Pointer(&flags),Native(pal),unsafe.Pointer(&enabled),unsafe.Pointer(&text),unsafe.Pointer(&textRole),nil,nil,nil,nil,nil,nil)
}	
//QStylePainter::style()
func (q *QStylePainter) Style() *QStyle {
	var __rv uintptr
	q.Drv(133000,133109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QStyle{}
	_rp.SetDriver(__rv,357000,false)
	return _rp
}	
