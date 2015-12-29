// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QAbstractSpinBox_StepEnabledFlag - QAbstractSpinBox::StepEnabledFlag
type QAbstractSpinBox_StepEnabledFlag uint32
const (
	QAbstractSpinBox_StepNone QAbstractSpinBox_StepEnabledFlag = 0x00
	QAbstractSpinBox_StepUpEnabled QAbstractSpinBox_StepEnabledFlag = 0x01
	QAbstractSpinBox_StepDownEnabled QAbstractSpinBox_StepEnabledFlag = 0x02
)
//enum QAbstractSpinBox_CorrectionMode - QAbstractSpinBox::CorrectionMode
type QAbstractSpinBox_CorrectionMode uint32
const (
	QAbstractSpinBox_CorrectToPreviousValue QAbstractSpinBox_CorrectionMode = 0
	QAbstractSpinBox_CorrectToNearestValue QAbstractSpinBox_CorrectionMode = 1
)
//enum QAbstractSpinBox_ButtonSymbols - QAbstractSpinBox::ButtonSymbols
type QAbstractSpinBox_ButtonSymbols uint32
const (
	QAbstractSpinBox_UpDownArrows QAbstractSpinBox_ButtonSymbols = 0
	QAbstractSpinBox_PlusMinus QAbstractSpinBox_ButtonSymbols = 1
	QAbstractSpinBox_NoButtons QAbstractSpinBox_ButtonSymbols = 2
)
//struct QAbstractSpinBox : QAbstractSpinBox
type QAbstractSpinBox struct {
	QWidget
}
func (q *QAbstractSpinBox) OnEditingFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(202000,202102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QAbstractSpinBox::QAbstractSpinBox()	
func NewAbstractSpinBox() *QAbstractSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),202000,202103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractSpinBox{}
	_p.SetDriver(__rv,202000,false)
	return _p
} 
//QAbstractSpinBox::QAbstractSpinBox(QWidget*)	
func NewAbstractSpinBoxWithParent(parent QWidgetInterface) *QAbstractSpinBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),202000,202104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QAbstractSpinBox{}
	_p.SetDriver(__rv,202000,false)
	return _p
} 
//QAbstractSpinBox::alignment()
func (q *QAbstractSpinBox) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(202000,202105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::buttonSymbols()
func (q *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox_ButtonSymbols {
	var __rv QAbstractSpinBox_ButtonSymbols
	q.Drv(202000,202106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::changeEvent(QEvent*)
func (q *QAbstractSpinBox) ChangeEvent(event *QEvent)  {
	q.Drv(202000,202107,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::clear()
func (q *QAbstractSpinBox) Clear()  {
	q.Drv(202000,202108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::closeEvent(QCloseEvent*)
func (q *QAbstractSpinBox) CloseEvent(event *QCloseEvent)  {
	q.Drv(202000,202109,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::contextMenuEvent(QContextMenuEvent*)
func (q *QAbstractSpinBox) ContextMenuEvent(event *QContextMenuEvent)  {
	q.Drv(202000,202110,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::correctionMode()
func (q *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox_CorrectionMode {
	var __rv QAbstractSpinBox_CorrectionMode
	q.Drv(202000,202111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::event(QEvent*)
func (q *QAbstractSpinBox) Event(event *QEvent) bool {
	var __rv bool
	q.Drv(202000,202112,Native(event),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::fixup(QString&)
func (q *QAbstractSpinBox) Fixup(input *string)  {
	q.Drv(202000,202113,unsafe.Pointer(&input),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::focusInEvent(QFocusEvent*)
func (q *QAbstractSpinBox) FocusInEvent(event *QFocusEvent)  {
	q.Drv(202000,202114,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::focusOutEvent(QFocusEvent*)
func (q *QAbstractSpinBox) FocusOutEvent(event *QFocusEvent)  {
	q.Drv(202000,202115,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::hasAcceptableInput()
func (q *QAbstractSpinBox) HasAcceptableInput() bool {
	var __rv bool
	q.Drv(202000,202116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::hasFrame()
func (q *QAbstractSpinBox) HasFrame() bool {
	var __rv bool
	q.Drv(202000,202117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::hideEvent(QHideEvent*)
func (q *QAbstractSpinBox) HideEvent(event *QHideEvent)  {
	q.Drv(202000,202118,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::inputMethodQuery(Qt::InputMethodQuery)
func (q *QAbstractSpinBox) InputMethodQuery(value Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(202000,202119,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QAbstractSpinBox::interpretText()
func (q *QAbstractSpinBox) InterpretText()  {
	q.Drv(202000,202120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::isAccelerated()
func (q *QAbstractSpinBox) IsAccelerated() bool {
	var __rv bool
	q.Drv(202000,202121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::isReadOnly()
func (q *QAbstractSpinBox) IsReadOnly() bool {
	var __rv bool
	q.Drv(202000,202122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::keyPressEvent(QKeyEvent*)
func (q *QAbstractSpinBox) KeyPressEvent(event *QKeyEvent)  {
	q.Drv(202000,202123,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::keyReleaseEvent(QKeyEvent*)
func (q *QAbstractSpinBox) KeyReleaseEvent(event *QKeyEvent)  {
	q.Drv(202000,202124,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::keyboardTracking()
func (q *QAbstractSpinBox) KeyboardTracking() bool {
	var __rv bool
	q.Drv(202000,202125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::lineEdit()
func (q *QAbstractSpinBox) LineEdit() *QLineEdit {
	var __rv uintptr
	q.Drv(202000,202126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLineEdit{}
	_rp.SetDriver(__rv,302000,false)
	return _rp
}	
//QAbstractSpinBox::minimumSizeHint()
func (q *QAbstractSpinBox) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(202000,202127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractSpinBox::mouseMoveEvent(QMouseEvent*)
func (q *QAbstractSpinBox) MouseMoveEvent(event *QMouseEvent)  {
	q.Drv(202000,202128,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::mousePressEvent(QMouseEvent*)
func (q *QAbstractSpinBox) MousePressEvent(event *QMouseEvent)  {
	q.Drv(202000,202129,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::mouseReleaseEvent(QMouseEvent*)
func (q *QAbstractSpinBox) MouseReleaseEvent(event *QMouseEvent)  {
	q.Drv(202000,202130,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::paintEvent(QPaintEvent*)
func (q *QAbstractSpinBox) PaintEvent(event *QPaintEvent)  {
	q.Drv(202000,202131,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::resizeEvent(QResizeEvent*)
func (q *QAbstractSpinBox) ResizeEvent(event *QResizeEvent)  {
	q.Drv(202000,202132,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::selectAll()
func (q *QAbstractSpinBox) SelectAll()  {
	q.Drv(202000,202133,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setAccelerated(bool)
func (q *QAbstractSpinBox) SetAccelerated(on bool)  {
	q.Drv(202000,202134,unsafe.Pointer(&on),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QAbstractSpinBox) SetAlignment(flag Qt_AlignmentFlag)  {
	q.Drv(202000,202135,unsafe.Pointer(&flag),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setButtonSymbols(QAbstractSpinBox::ButtonSymbols)
func (q *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox_ButtonSymbols)  {
	q.Drv(202000,202136,unsafe.Pointer(&bs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setCorrectionMode(QAbstractSpinBox::CorrectionMode)
func (q *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox_CorrectionMode)  {
	q.Drv(202000,202137,unsafe.Pointer(&cm),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setFrame(bool)
func (q *QAbstractSpinBox) SetFrame(value bool)  {
	q.Drv(202000,202138,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setKeyboardTracking(bool)
func (q *QAbstractSpinBox) SetKeyboardTracking(kt bool)  {
	q.Drv(202000,202139,unsafe.Pointer(&kt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setLineEdit(QLineEdit*)
func (q *QAbstractSpinBox) SetLineEdit(edit *QLineEdit)  {
	q.Drv(202000,202140,Native(edit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setReadOnly(bool)
func (q *QAbstractSpinBox) SetReadOnly(r bool)  {
	q.Drv(202000,202141,unsafe.Pointer(&r),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setSpecialValueText(QString const&)
func (q *QAbstractSpinBox) SetSpecialValueText(txt string)  {
	q.Drv(202000,202142,unsafe.Pointer(&txt),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::setWrapping(bool)
func (q *QAbstractSpinBox) SetWrapping(w bool)  {
	q.Drv(202000,202143,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::showEvent(QShowEvent*)
func (q *QAbstractSpinBox) ShowEvent(event *QShowEvent)  {
	q.Drv(202000,202144,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::sizeHint()
func (q *QAbstractSpinBox) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(202000,202145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QAbstractSpinBox::specialValueText()
func (q *QAbstractSpinBox) SpecialValueText() string {
	var __rv string
	q.Drv(202000,202146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::stepBy(int)
func (q *QAbstractSpinBox) StepBy(steps int)  {
	q.Drv(202000,202147,unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::stepDown()
func (q *QAbstractSpinBox) StepDown()  {
	q.Drv(202000,202148,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::stepEnabled()
func (q *QAbstractSpinBox) StepEnabled() QAbstractSpinBox_StepEnabledFlag {
	var __rv QAbstractSpinBox_StepEnabledFlag
	q.Drv(202000,202149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::stepUp()
func (q *QAbstractSpinBox) StepUp()  {
	q.Drv(202000,202150,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::text()
func (q *QAbstractSpinBox) Text() string {
	var __rv string
	q.Drv(202000,202151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::timerEvent(QTimerEvent*)
func (q *QAbstractSpinBox) TimerEvent(event *QTimerEvent)  {
	q.Drv(202000,202152,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::validate(QString&,int&)
func (q *QAbstractSpinBox) Validate(input *string,pos *int) QValidator_State {
	var __rv QValidator_State
	q.Drv(202000,202153,unsafe.Pointer(&input),unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QAbstractSpinBox::wheelEvent(QWheelEvent*)
func (q *QAbstractSpinBox) WheelEvent(event *QWheelEvent)  {
	q.Drv(202000,202154,Native(event),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QAbstractSpinBox::wrapping()
func (q *QAbstractSpinBox) Wrapping() bool {
	var __rv bool
	q.Drv(202000,202155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
