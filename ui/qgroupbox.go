// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QGroupBox : QGroupBox
type QGroupBox struct {
	QWidget
}
func (q *QGroupBox) OnClicked(fn func()) uintptr {
	var __rv uintptr
	q.Drv(288000,288102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGroupBox) OnClickedEx(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(288000,288103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QGroupBox) OnToggled(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(288000,288104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QGroupBox::QGroupBox()	
func NewGroupBox() *QGroupBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),288000,288105,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGroupBox{}
	_p.SetDriver(__rv,288000,false)
	return _p
} 
//QGroupBox::QGroupBox(QWidget*)	
func NewGroupBoxWithParent(parent QWidgetInterface) *QGroupBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),288000,288106,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGroupBox{}
	_p.SetDriver(__rv,288000,false)
	return _p
} 
//QGroupBox::QGroupBox(QString const&,QWidget*)	
func NewGroupBoxWithTitleParent(title string,parent QWidgetInterface) *QGroupBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),288000,288107,unsafe.Pointer(&title),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QGroupBox{}
	_p.SetDriver(__rv,288000,false)
	return _p
} 
//QGroupBox::alignment()
func (q *QGroupBox) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(288000,288108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGroupBox::changeEvent(QEvent*)
func (q *QGroupBox) ChangeEvent(event *QEvent)  {
	q.Drv(288000,288109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::childEvent(QChildEvent*)
func (q *QGroupBox) ChildEvent(event *QChildEvent)  {
	q.Drv(288000,288110,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::event(QEvent*)
func (q *QGroupBox) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(288000,288111,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGroupBox::focusInEvent(QFocusEvent*)
func (q *QGroupBox) FocusInEvent(event *QFocusEvent)  {
	q.Drv(288000,288112,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::isCheckable()
func (q *QGroupBox) IsCheckable() bool {
	var __rv bool
	q.Drv(288000,288113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGroupBox::isChecked()
func (q *QGroupBox) IsChecked() bool {
	var __rv bool
	q.Drv(288000,288114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGroupBox::isFlat()
func (q *QGroupBox) IsFlat() bool {
	var __rv bool
	q.Drv(288000,288115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QGroupBox::minimumSizeHint()
func (q *QGroupBox) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(288000,288116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QGroupBox::mouseMoveEvent(QMouseEvent*)
func (q *QGroupBox) MouseMoveEvent(event *QMouseEvent)  {
	q.Drv(288000,288117,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::mousePressEvent(QMouseEvent*)
func (q *QGroupBox) MousePressEvent(event *QMouseEvent)  {
	q.Drv(288000,288118,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::mouseReleaseEvent(QMouseEvent*)
func (q *QGroupBox) MouseReleaseEvent(event *QMouseEvent)  {
	q.Drv(288000,288119,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::paintEvent(QPaintEvent*)
func (q *QGroupBox) PaintEvent(event *QPaintEvent)  {
	q.Drv(288000,288120,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::resizeEvent(QResizeEvent*)
func (q *QGroupBox) ResizeEvent(event *QResizeEvent)  {
	q.Drv(288000,288121,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::setAlignment(int)
func (q *QGroupBox) SetAlignment(alignment int)  {
	q.Drv(288000,288122,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::setCheckable(bool)
func (q *QGroupBox) SetCheckable(checkable bool)  {
	q.Drv(288000,288123,unsafe.Pointer(&checkable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::setChecked(bool)
func (q *QGroupBox) SetChecked(checked bool)  {
	q.Drv(288000,288124,unsafe.Pointer(&checked),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::setFlat(bool)
func (q *QGroupBox) SetFlat(flat bool)  {
	q.Drv(288000,288125,unsafe.Pointer(&flat),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::setTitle(QString const&)
func (q *QGroupBox) SetTitle(title string)  {
	q.Drv(288000,288126,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QGroupBox::title()
func (q *QGroupBox) Title() string {
	var __rv string
	q.Drv(288000,288127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
