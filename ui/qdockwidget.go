// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QDockWidget_DockWidgetFeature - QDockWidget::DockWidgetFeature
type QDockWidget_DockWidgetFeature uint32
const (
	QDockWidget_DockWidgetClosable QDockWidget_DockWidgetFeature = 0x01
	QDockWidget_DockWidgetMovable QDockWidget_DockWidgetFeature = 0x02
	QDockWidget_DockWidgetFloatable QDockWidget_DockWidgetFeature = 0x04
	QDockWidget_DockWidgetVerticalTitleBar QDockWidget_DockWidgetFeature = 0x08
	QDockWidget_DockWidgetFeatureMask QDockWidget_DockWidgetFeature = 0x0f
	QDockWidget_AllDockWidgetFeatures QDockWidget_DockWidgetFeature = QDockWidget_DockWidgetClosable|QDockWidget_DockWidgetMovable|QDockWidget_DockWidgetFloatable
	QDockWidget_NoDockWidgetFeatures QDockWidget_DockWidgetFeature = 0x00
	QDockWidget_Reserved QDockWidget_DockWidgetFeature = 0xff
)
//struct QDockWidget : QDockWidget
type QDockWidget struct {
	QWidget
}
func (q *QDockWidget) OnTopLevelChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(229000,229102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDockWidget) OnFeaturesChanged(fn func(QDockWidget_DockWidgetFeature)) uintptr {
	var __rv uintptr
	q.Drv(229000,229103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDockWidget) OnAllowedAreasChanged(fn func(Qt_DockWidgetArea)) uintptr {
	var __rv uintptr
	q.Drv(229000,229104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDockWidget) OnDockLocationChanged(fn func(Qt_DockWidgetArea)) uintptr {
	var __rv uintptr
	q.Drv(229000,229105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QDockWidget) OnVisibilityChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(229000,229106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QDockWidget::QDockWidget()	
func NewDockWidget() *QDockWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),229000,229107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDockWidget{}
	_p.SetDriver(__rv,229000,false)
	return _p
} 
//QDockWidget::QDockWidget(QWidget*,QFlags<Qt::WindowType>)	
func NewDockWidgetWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QDockWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),229000,229108,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDockWidget{}
	_p.SetDriver(__rv,229000,false)
	return _p
} 
//QDockWidget::QDockWidget(QString const&,QWidget*,QFlags<Qt::WindowType>)	
func NewDockWidgetWithTitleParentFlags(title string,parent QWidgetInterface,flags Qt_WindowType) *QDockWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),229000,229109,unsafe.Pointer(&title),Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QDockWidget{}
	_p.SetDriver(__rv,229000,false)
	return _p
} 
//QDockWidget::allowedAreas()
func (q *QDockWidget) AllowedAreas() Qt_DockWidgetArea {
	var __rv Qt_DockWidgetArea
	q.Drv(229000,229110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDockWidget::changeEvent(QEvent*)
func (q *QDockWidget) ChangeEvent(event *QEvent)  {
	q.Drv(229000,229111,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::closeEvent(QCloseEvent*)
func (q *QDockWidget) CloseEvent(event *QCloseEvent)  {
	q.Drv(229000,229112,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::event(QEvent*)
func (q *QDockWidget) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(229000,229113,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDockWidget::features()
func (q *QDockWidget) Features() QDockWidget_DockWidgetFeature {
	var __rv QDockWidget_DockWidgetFeature
	q.Drv(229000,229114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDockWidget::isAreaAllowed(Qt::DockWidgetArea)
func (q *QDockWidget) IsAreaAllowed(area Qt_DockWidgetArea) bool {
	var __rv bool
	q.Drv(229000,229115,unsafe.Pointer(&area),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDockWidget::isFloating()
func (q *QDockWidget) IsFloating() bool {
	var __rv bool
	q.Drv(229000,229116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QDockWidget::paintEvent(QPaintEvent*)
func (q *QDockWidget) PaintEvent(event *QPaintEvent)  {
	q.Drv(229000,229117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::setAllowedAreas(QFlags<Qt::DockWidgetArea>)
func (q *QDockWidget) SetAllowedAreas(areas Qt_DockWidgetArea)  {
	q.Drv(229000,229118,unsafe.Pointer(&areas),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::setFeatures(QFlags<QDockWidget::DockWidgetFeature>)
func (q *QDockWidget) SetFeatures(features QDockWidget_DockWidgetFeature)  {
	q.Drv(229000,229119,unsafe.Pointer(&features),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::setFloating(bool)
func (q *QDockWidget) SetFloating(floating bool)  {
	q.Drv(229000,229120,unsafe.Pointer(&floating),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::setTitleBarWidget(QWidget*)
func (q *QDockWidget) SetTitleBarWidget(widget QWidgetInterface)  {
	q.Drv(229000,229121,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::setWidget(QWidget*)
func (q *QDockWidget) SetWidget(widget QWidgetInterface)  {
	q.Drv(229000,229122,Native(widget),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QDockWidget::titleBarWidget()
func (q *QDockWidget) TitleBarWidget() *QWidget {
	var __rv uintptr
	q.Drv(229000,229123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QDockWidget::toggleViewAction()
func (q *QDockWidget) ToggleViewAction() *QAction {
	var __rv uintptr
	q.Drv(229000,229124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAction{}
	_rp.SetDriver(__rv,207000,false)
	return _rp
}	
//QDockWidget::widget()
func (q *QDockWidget) Widget() *QWidget {
	var __rv uintptr
	q.Drv(229000,229125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
