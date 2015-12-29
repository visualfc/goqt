// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPrintPreviewWidget_ZoomMode - QPrintPreviewWidget::ZoomMode
type QPrintPreviewWidget_ZoomMode uint32
const (
	QPrintPreviewWidget_CustomZoom QPrintPreviewWidget_ZoomMode = 0
	QPrintPreviewWidget_FitToWidth QPrintPreviewWidget_ZoomMode = 1
	QPrintPreviewWidget_FitInView QPrintPreviewWidget_ZoomMode = 2
)
//enum QPrintPreviewWidget_ViewMode - QPrintPreviewWidget::ViewMode
type QPrintPreviewWidget_ViewMode uint32
const (
	QPrintPreviewWidget_SinglePageView QPrintPreviewWidget_ViewMode = 0
	QPrintPreviewWidget_FacingPagesView QPrintPreviewWidget_ViewMode = 1
	QPrintPreviewWidget_AllPagesView QPrintPreviewWidget_ViewMode = 2
)
//struct QPrintPreviewWidget : QPrintPreviewWidget
type QPrintPreviewWidget struct {
	QWidget
}
func (q *QPrintPreviewWidget) OnPreviewChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(325000,325102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPrintPreviewWidget) OnPaintRequested(fn func(*QPrinter)) uintptr {
	var __rv uintptr
	q.Drv(325000,325103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QPrintPreviewWidget::QPrintPreviewWidget()	
func NewPrintPreviewWidget() *QPrintPreviewWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),325000,325104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewWidget{}
	_p.SetDriver(__rv,325000,false)
	return _p
} 
//QPrintPreviewWidget::QPrintPreviewWidget(QWidget*,QFlags<Qt::WindowType>)	
func NewPrintPreviewWidgetWithParentFlags(parent QWidgetInterface,flags Qt_WindowType) *QPrintPreviewWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),325000,325105,Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewWidget{}
	_p.SetDriver(__rv,325000,false)
	return _p
} 
//QPrintPreviewWidget::QPrintPreviewWidget(QPrinter*,QWidget*,QFlags<Qt::WindowType>)	
func NewPrintPreviewWidgetWithPrinterParentFlags(printer *QPrinter,parent QWidgetInterface,flags Qt_WindowType) *QPrintPreviewWidget {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),325000,325106,Native(printer),Native(parent),unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPrintPreviewWidget{}
	_p.SetDriver(__rv,325000,false)
	return _p
} 
//QPrintPreviewWidget::currentPage()
func (q *QPrintPreviewWidget) CurrentPage() int {
	var __rv int
	q.Drv(325000,325107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::fitInView()
func (q *QPrintPreviewWidget) FitInView()  {
	q.Drv(325000,325108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::fitToWidth()
func (q *QPrintPreviewWidget) FitToWidth()  {
	q.Drv(325000,325109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::orientation()
func (q *QPrintPreviewWidget) Orientation() QPrinter_Orientation {
	var __rv QPrinter_Orientation
	q.Drv(325000,325110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::pageCount()
func (q *QPrintPreviewWidget) PageCount() int {
	var __rv int
	q.Drv(325000,325111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::print()
func (q *QPrintPreviewWidget) Print()  {
	q.Drv(325000,325112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setAllPagesViewMode()
func (q *QPrintPreviewWidget) SetAllPagesViewMode()  {
	q.Drv(325000,325113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setCurrentPage(int)
func (q *QPrintPreviewWidget) SetCurrentPage(pageNumber int)  {
	q.Drv(325000,325114,unsafe.Pointer(&pageNumber),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setFacingPagesViewMode()
func (q *QPrintPreviewWidget) SetFacingPagesViewMode()  {
	q.Drv(325000,325115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setLandscapeOrientation()
func (q *QPrintPreviewWidget) SetLandscapeOrientation()  {
	q.Drv(325000,325116,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setOrientation(QPrinter::Orientation)
func (q *QPrintPreviewWidget) SetOrientation(orientation QPrinter_Orientation)  {
	q.Drv(325000,325117,unsafe.Pointer(&orientation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setPortraitOrientation()
func (q *QPrintPreviewWidget) SetPortraitOrientation()  {
	q.Drv(325000,325118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setSinglePageViewMode()
func (q *QPrintPreviewWidget) SetSinglePageViewMode()  {
	q.Drv(325000,325119,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setViewMode(QPrintPreviewWidget::ViewMode)
func (q *QPrintPreviewWidget) SetViewMode(viewMode QPrintPreviewWidget_ViewMode)  {
	q.Drv(325000,325120,unsafe.Pointer(&viewMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setVisible(bool)
func (q *QPrintPreviewWidget) SetVisible(visible bool)  {
	q.Drv(325000,325121,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setZoomFactor(double)
func (q *QPrintPreviewWidget) SetZoomFactor(zoomFactor float64)  {
	q.Drv(325000,325122,unsafe.Pointer(&zoomFactor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::setZoomMode(QPrintPreviewWidget::ZoomMode)
func (q *QPrintPreviewWidget) SetZoomMode(zoomMode QPrintPreviewWidget_ZoomMode)  {
	q.Drv(325000,325123,unsafe.Pointer(&zoomMode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::updatePreview()
func (q *QPrintPreviewWidget) UpdatePreview()  {
	q.Drv(325000,325124,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::viewMode()
func (q *QPrintPreviewWidget) ViewMode() QPrintPreviewWidget_ViewMode {
	var __rv QPrintPreviewWidget_ViewMode
	q.Drv(325000,325125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::zoomFactor()
func (q *QPrintPreviewWidget) ZoomFactor() float64 {
	var __rv float64
	q.Drv(325000,325126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::zoomIn()
func (q *QPrintPreviewWidget) ZoomIn()  {
	q.Drv(325000,325127,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::zoomIn(double)
func (q *QPrintPreviewWidget) ZoomInFWithZoom(zoom float64)  {
	q.Drv(325000,325128,unsafe.Pointer(&zoom),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::zoomMode()
func (q *QPrintPreviewWidget) ZoomMode() QPrintPreviewWidget_ZoomMode {
	var __rv QPrintPreviewWidget_ZoomMode
	q.Drv(325000,325129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPrintPreviewWidget::zoomOut()
func (q *QPrintPreviewWidget) ZoomOut()  {
	q.Drv(325000,325130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPrintPreviewWidget::zoomOut(double)
func (q *QPrintPreviewWidget) ZoomOutFWithZoom(zoom float64)  {
	q.Drv(325000,325131,unsafe.Pointer(&zoom),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
