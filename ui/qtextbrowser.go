// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextBrowser : QTextBrowser
type QTextBrowser struct {
	QTextEdit
}
func (q *QTextBrowser) OnSourceChanged(fn func(*QUrl)) uintptr {
	var __rv uintptr
	q.Drv(371000,371102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnAnchorClicked(fn func(*QUrl)) uintptr {
	var __rv uintptr
	q.Drv(371000,371103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnHighlighted(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(371000,371104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnHighlightedWithUrl(fn func(*QUrl)) uintptr {
	var __rv uintptr
	q.Drv(371000,371105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnHistoryChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(371000,371106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnBackwardAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(371000,371107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextBrowser) OnForwardAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(371000,371108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTextBrowser::QTextBrowser()	
func NewTextBrowser() *QTextBrowser {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),371000,371109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBrowser{}
	_p.SetDriver(__rv,371000,false)
	return _p
} 
//QTextBrowser::QTextBrowser(QWidget*)	
func NewTextBrowserWithParent(parent QWidgetInterface) *QTextBrowser {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),371000,371110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBrowser{}
	_p.SetDriver(__rv,371000,false)
	return _p
} 
//QTextBrowser::backward()
func (q *QTextBrowser) Backward()  {
	q.Drv(371000,371111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::backwardHistoryCount()
func (q *QTextBrowser) BackwardHistoryCount() int {
	var __rv int
	q.Drv(371000,371112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::clearHistory()
func (q *QTextBrowser) ClearHistory()  {
	q.Drv(371000,371113,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::event(QEvent*)
func (q *QTextBrowser) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(371000,371114,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::focusNextPrevChild(bool)
func (q *QTextBrowser) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(371000,371115,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::focusOutEvent(QFocusEvent*)
func (q *QTextBrowser) FocusOutEvent(ev *QFocusEvent)  {
	q.Drv(371000,371116,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::forward()
func (q *QTextBrowser) Forward()  {
	q.Drv(371000,371117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::forwardHistoryCount()
func (q *QTextBrowser) ForwardHistoryCount() int {
	var __rv int
	q.Drv(371000,371118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::historyTitle(int)
func (q *QTextBrowser) HistoryTitle(value int) string {
	var __rv string
	q.Drv(371000,371119,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::historyUrl(int)
func (q *QTextBrowser) HistoryUrl(value int) *QUrl {
	var __rv uintptr
	q.Drv(371000,371120,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
//QTextBrowser::home()
func (q *QTextBrowser) Home()  {
	q.Drv(371000,371121,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::isBackwardAvailable()
func (q *QTextBrowser) IsBackwardAvailable() bool {
	var __rv bool
	q.Drv(371000,371122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::isForwardAvailable()
func (q *QTextBrowser) IsForwardAvailable() bool {
	var __rv bool
	q.Drv(371000,371123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::keyPressEvent(QKeyEvent*)
func (q *QTextBrowser) KeyPressEvent(ev *QKeyEvent)  {
	q.Drv(371000,371124,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::loadResource(int,QUrl const&)
func (q *QTextBrowser) LoadResource(_type int,name *QUrl) *QVariant {
	var __rv uintptr
	q.Drv(371000,371125,unsafe.Pointer(&_type),Native(name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextBrowser::mouseMoveEvent(QMouseEvent*)
func (q *QTextBrowser) MouseMoveEvent(ev *QMouseEvent)  {
	q.Drv(371000,371126,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::mousePressEvent(QMouseEvent*)
func (q *QTextBrowser) MousePressEvent(ev *QMouseEvent)  {
	q.Drv(371000,371127,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::mouseReleaseEvent(QMouseEvent*)
func (q *QTextBrowser) MouseReleaseEvent(ev *QMouseEvent)  {
	q.Drv(371000,371128,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::openExternalLinks()
func (q *QTextBrowser) OpenExternalLinks() bool {
	var __rv bool
	q.Drv(371000,371129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::openLinks()
func (q *QTextBrowser) OpenLinks() bool {
	var __rv bool
	q.Drv(371000,371130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::paintEvent(QPaintEvent*)
func (q *QTextBrowser) PaintEvent(e *QPaintEvent)  {
	q.Drv(371000,371131,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::reload()
func (q *QTextBrowser) Reload()  {
	q.Drv(371000,371132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::searchPaths()
func (q *QTextBrowser) SearchPaths() []string {
	var __rv []string
	q.Drv(371000,371133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBrowser::setOpenExternalLinks(bool)
func (q *QTextBrowser) SetOpenExternalLinks(open bool)  {
	q.Drv(371000,371134,unsafe.Pointer(&open),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::setOpenLinks(bool)
func (q *QTextBrowser) SetOpenLinks(open bool)  {
	q.Drv(371000,371135,unsafe.Pointer(&open),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::setSearchPaths(QStringList const&)
func (q *QTextBrowser) SetSearchPaths(paths []string)  {
	q.Drv(371000,371136,unsafe.Pointer(&paths),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::setSource(QUrl const&)
func (q *QTextBrowser) SetSource(name *QUrl)  {
	q.Drv(371000,371137,Native(name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBrowser::source()
func (q *QTextBrowser) Source() *QUrl {
	var __rv uintptr
	q.Drv(371000,371138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QUrl{}
	_rp.SetDriver(__rv,180000,true)
	return _rp
}	
