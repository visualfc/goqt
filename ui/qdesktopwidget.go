// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QDesktopWidget : QDesktopWidget
type QDesktopWidget struct {
	QWidget
}
func (q *QDesktopWidget) OnResized(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(224000,224102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDesktopWidget) OnWorkAreaResized(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(224000,224103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDesktopWidget) OnScreenCountChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(224000,224104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDesktopWidget::QDesktopWidget()	
func NewDesktopWidget() *QDesktopWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),224000,224105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDesktopWidget{}
	_p.SetDriver(__rv,224000,false)
	return _p
} 
//QDesktopWidget::availableGeometry()
func (q *QDesktopWidget) AvailableGeometry() *QRect {
	var __rv uintptr
	q.Drv(224000,224106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::availableGeometry(QPoint const&)
func (q *QDesktopWidget) AvailableGeometryWithPoint(point *QPoint) *QRect {
	var __rv uintptr
	q.Drv(224000,224107,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::availableGeometry(QWidget const*)
func (q *QDesktopWidget) AvailableGeometryWithWidget(widget QWidgetInterface) *QRect {
	var __rv uintptr
	q.Drv(224000,224108,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::availableGeometry(int)
func (q *QDesktopWidget) AvailableGeometryWithScreen(screen int) *QRect {
	var __rv uintptr
	q.Drv(224000,224109,unsafe.Pointer(&screen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::isVirtualDesktop()
func (q *QDesktopWidget) IsVirtualDesktop() bool {
	var __rv bool
	q.Drv(224000,224110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::numScreens()
func (q *QDesktopWidget) NumScreens() int {
	var __rv int
	q.Drv(224000,224111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::primaryScreen()
func (q *QDesktopWidget) PrimaryScreen() int {
	var __rv int
	q.Drv(224000,224112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::resizeEvent(QResizeEvent*)
func (q *QDesktopWidget) ResizeEvent(e *QResizeEvent)  {
	q.Drv(224000,224113,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDesktopWidget::screen()
func (q *QDesktopWidget) Screen() *QWidget {
	var __rv uintptr
	q.Drv(224000,224114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDesktopWidget::screen(int)
func (q *QDesktopWidget) ScreenWithScreen(screen int) *QWidget {
	var __rv uintptr
	q.Drv(224000,224115,unsafe.Pointer(&screen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDesktopWidget::screenCount()
func (q *QDesktopWidget) ScreenCount() int {
	var __rv int
	q.Drv(224000,224116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::screenGeometry()
func (q *QDesktopWidget) ScreenGeometry() *QRect {
	var __rv uintptr
	q.Drv(224000,224117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::screenGeometry(QPoint const&)
func (q *QDesktopWidget) ScreenGeometryWithPoint(point *QPoint) *QRect {
	var __rv uintptr
	q.Drv(224000,224118,Native(point),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::screenGeometry(QWidget const*)
func (q *QDesktopWidget) ScreenGeometryWithWidget(widget QWidgetInterface) *QRect {
	var __rv uintptr
	q.Drv(224000,224119,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::screenGeometry(int)
func (q *QDesktopWidget) ScreenGeometryWithScreen(screen int) *QRect {
	var __rv uintptr
	q.Drv(224000,224120,unsafe.Pointer(&screen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QDesktopWidget::screenNumber()
func (q *QDesktopWidget) ScreenNumber() int {
	var __rv int
	q.Drv(224000,224121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::screenNumber(QPoint const&)
func (q *QDesktopWidget) ScreenNumberWithPoint(value *QPoint) int {
	var __rv int
	q.Drv(224000,224122,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDesktopWidget::screenNumber(QWidget const*)
func (q *QDesktopWidget) ScreenNumberWithWidget(widget QWidgetInterface) int {
	var __rv int
	q.Drv(224000,224123,Native(widget),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
