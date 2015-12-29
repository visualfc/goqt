// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QLabel : QLabel
type QLabel struct {
	QFrame
}
func (q *QLabel) OnLinkActivated(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(299000,299102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLabel) OnLinkHovered(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(299000,299103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QLabel::QLabel()	
func NewLabel() *QLabel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),299000,299104,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLabel{}
	_p.SetDriver(__rv,299000,false)
	return _p
} 
//QLabel::QLabel(QWidget*,QFlags<Qt::WindowType>)	
func NewLabelWithParentFlags(parent QWidgetInterface,f Qt_WindowType) *QLabel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),299000,299105,Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLabel{}
	_p.SetDriver(__rv,299000,false)
	return _p
} 
//QLabel::QLabel(QString const&,QWidget*,QFlags<Qt::WindowType>)	
func NewLabelWithTextParentFlags(text string,parent QWidgetInterface,f Qt_WindowType) *QLabel {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),299000,299106,unsafe.Pointer(&text),Native(parent),unsafe.Pointer(&f),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLabel{}
	_p.SetDriver(__rv,299000,false)
	return _p
} 
//QLabel::alignment()
func (q *QLabel) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(299000,299107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::buddy()
func (q *QLabel) Buddy() *QWidget {
	var __rv uintptr
	q.Drv(299000,299108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QWidget{}
	_rp.SetDriver(__rv,395000,false)
	return _rp
}	
//QLabel::changeEvent(QEvent*)
func (q *QLabel) ChangeEvent(value *QEvent)  {
	q.Drv(299000,299109,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::clear()
func (q *QLabel) Clear()  {
	q.Drv(299000,299110,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::contextMenuEvent(QContextMenuEvent*)
func (q *QLabel) ContextMenuEvent(ev *QContextMenuEvent)  {
	q.Drv(299000,299111,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::event(QEvent*)
func (q *QLabel) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(299000,299112,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::focusInEvent(QFocusEvent*)
func (q *QLabel) FocusInEvent(ev *QFocusEvent)  {
	q.Drv(299000,299113,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::focusNextPrevChild(bool)
func (q *QLabel) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(299000,299114,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::focusOutEvent(QFocusEvent*)
func (q *QLabel) FocusOutEvent(ev *QFocusEvent)  {
	q.Drv(299000,299115,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::hasScaledContents()
func (q *QLabel) HasScaledContents() bool {
	var __rv bool
	q.Drv(299000,299116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::hasSelectedText()
func (q *QLabel) HasSelectedText() bool {
	var __rv bool
	q.Drv(299000,299117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::heightForWidth(int)
func (q *QLabel) HeightForWidth(value int) int {
	var __rv int
	q.Drv(299000,299118,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::indent()
func (q *QLabel) Indent() int {
	var __rv int
	q.Drv(299000,299119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::keyPressEvent(QKeyEvent*)
func (q *QLabel) KeyPressEvent(ev *QKeyEvent)  {
	q.Drv(299000,299120,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::margin()
func (q *QLabel) Margin() int {
	var __rv int
	q.Drv(299000,299121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::minimumSizeHint()
func (q *QLabel) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(299000,299122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLabel::mouseMoveEvent(QMouseEvent*)
func (q *QLabel) MouseMoveEvent(ev *QMouseEvent)  {
	q.Drv(299000,299123,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::mousePressEvent(QMouseEvent*)
func (q *QLabel) MousePressEvent(ev *QMouseEvent)  {
	q.Drv(299000,299124,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::mouseReleaseEvent(QMouseEvent*)
func (q *QLabel) MouseReleaseEvent(ev *QMouseEvent)  {
	q.Drv(299000,299125,Native(ev),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::movie()
func (q *QLabel) Movie() *QMovie {
	var __rv uintptr
	q.Drv(299000,299126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMovie{}
	_rp.SetDriver(__rv,313000,false)
	return _rp
}	
//QLabel::openExternalLinks()
func (q *QLabel) OpenExternalLinks() bool {
	var __rv bool
	q.Drv(299000,299127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::paintEvent(QPaintEvent*)
func (q *QLabel) PaintEvent(value *QPaintEvent)  {
	q.Drv(299000,299128,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::picture()
func (q *QLabel) Picture() *QPicture {
	var __rv uintptr
	q.Drv(299000,299129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPicture{}
	_rp.SetDriver(__rv,94000,true)
	return _rp
}	
//QLabel::pixmap()
func (q *QLabel) Pixmap() *QPixmap {
	var __rv uintptr
	q.Drv(299000,299130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPixmap{}
	_rp.SetDriver(__rv,96000,true)
	return _rp
}	
//QLabel::selectedText()
func (q *QLabel) SelectedText() string {
	var __rv string
	q.Drv(299000,299131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::selectionStart()
func (q *QLabel) SelectionStart() int {
	var __rv int
	q.Drv(299000,299132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QLabel) SetAlignment(value Qt_AlignmentFlag)  {
	q.Drv(299000,299133,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setBuddy(QWidget*)
func (q *QLabel) SetBuddy(value QWidgetInterface)  {
	q.Drv(299000,299134,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setIndent(int)
func (q *QLabel) SetIndent(value int)  {
	q.Drv(299000,299135,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setMargin(int)
func (q *QLabel) SetMargin(value int)  {
	q.Drv(299000,299136,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setMovie(QMovie*)
func (q *QLabel) SetMovie(movie *QMovie)  {
	q.Drv(299000,299137,Native(movie),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setNum(double)
func (q *QLabel) SetNum(value float64)  {
	q.Drv(299000,299138,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setNum(int)
func (q *QLabel) SetNumWithInt(value int)  {
	q.Drv(299000,299139,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setOpenExternalLinks(bool)
func (q *QLabel) SetOpenExternalLinks(open bool)  {
	q.Drv(299000,299140,unsafe.Pointer(&open),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setPicture(QPicture const&)
func (q *QLabel) SetPicture(value *QPicture)  {
	q.Drv(299000,299141,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setPixmap(QPixmap const&)
func (q *QLabel) SetPixmap(value *QPixmap)  {
	q.Drv(299000,299142,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setScaledContents(bool)
func (q *QLabel) SetScaledContents(value bool)  {
	q.Drv(299000,299143,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setSelection(int,int)
func (q *QLabel) SetSelection(value2 int,value3 int)  {
	q.Drv(299000,299144,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setText(QString const&)
func (q *QLabel) SetText(value string)  {
	q.Drv(299000,299145,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setTextFormat(Qt::TextFormat)
func (q *QLabel) SetTextFormat(value Qt_TextFormat)  {
	q.Drv(299000,299146,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>)
func (q *QLabel) SetTextInteractionFlags(flags Qt_TextInteractionFlag)  {
	q.Drv(299000,299147,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::setWordWrap(bool)
func (q *QLabel) SetWordWrap(on bool)  {
	q.Drv(299000,299148,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLabel::sizeHint()
func (q *QLabel) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(299000,299149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLabel::text()
func (q *QLabel) Text() string {
	var __rv string
	q.Drv(299000,299150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::textFormat()
func (q *QLabel) TextFormat() Qt_TextFormat {
	var __rv Qt_TextFormat
	q.Drv(299000,299151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::textInteractionFlags()
func (q *QLabel) TextInteractionFlags() Qt_TextInteractionFlag {
	var __rv Qt_TextInteractionFlag
	q.Drv(299000,299152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLabel::wordWrap()
func (q *QLabel) WordWrap() bool {
	var __rv bool
	q.Drv(299000,299153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
