// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSplashScreen : QSplashScreen
type QSplashScreen struct {
	QWidget
}
func (q *QSplashScreen) OnMessageChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(347000,347102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSplashScreen::QSplashScreen()	
func NewSplashScreen() *QSplashScreen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),347000,347103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplashScreen{}
	_p.SetDriver(__rv,347000,false)
	return _p
} 
//QSplashScreen::QSplashScreen(QPixmap const&,QFlags<Qt::WindowType>)	
func NewSplashScreenWithPixmapFlags(pixmap *QPixmap,f Qt_WindowType) *QSplashScreen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),347000,347104,Native(pixmap),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplashScreen{}
	_p.SetDriver(__rv,347000,false)
	return _p
} 
//QSplashScreen::QSplashScreen(QWidget*,QPixmap const&,QFlags<Qt::WindowType>)	
func NewSplashScreenWithParentPixmapFlags(parent QWidgetInterface,pixmap *QPixmap,f Qt_WindowType) *QSplashScreen {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),347000,347105,Native(parent),Native(pixmap),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSplashScreen{}
	_p.SetDriver(__rv,347000,false)
	return _p
} 
//QSplashScreen::clearMessage()
func (q *QSplashScreen) ClearMessage()  {
	q.Drv(347000,347106,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::drawContents(QPainter*)
func (q *QSplashScreen) DrawContents(painter *QPainter)  {
	q.Drv(347000,347107,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::event(QEvent*)
func (q *QSplashScreen) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(347000,347108,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSplashScreen::finish(QWidget*)
func (q *QSplashScreen) Finish(w QWidgetInterface)  {
	q.Drv(347000,347109,Native(w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::mousePressEvent(QMouseEvent*)
func (q *QSplashScreen) MousePressEvent(value *QMouseEvent)  {
	q.Drv(347000,347110,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::pixmap()
func (q *QSplashScreen) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(347000,347111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QSplashScreen::repaint()
func (q *QSplashScreen) Repaint()  {
	q.Drv(347000,347112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::setPixmap(QPixmap const&)
func (q *QSplashScreen) SetPixmap(pixmap *QPixmap)  {
	q.Drv(347000,347113,Native(pixmap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::showMessage(QString const&)
func (q *QSplashScreen) ShowMessage(message string)  {
	q.Drv(347000,347114,unsafe.Pointer(&message),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSplashScreen::showMessage(QString const&,int,QColor const&)
func (q *QSplashScreen) ShowMessageWithMessageAlignmentColor(message string,alignment int,color *QColor)  {
	q.Drv(347000,347115,unsafe.Pointer(&message),unsafe.Pointer(&alignment),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
