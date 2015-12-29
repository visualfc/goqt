// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QLineEdit_EchoMode - QLineEdit::EchoMode
type QLineEdit_EchoMode uint32
const (
	QLineEdit_Normal QLineEdit_EchoMode = 0
	QLineEdit_NoEcho QLineEdit_EchoMode = 1
	QLineEdit_Password QLineEdit_EchoMode = 2
	QLineEdit_PasswordEchoOnEdit QLineEdit_EchoMode = 3
)
//struct QLineEdit : QLineEdit
type QLineEdit struct {
	QWidget
}
func (q *QLineEdit) OnTextEdited(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(302000,302102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLineEdit) OnReturnPressed(fn func()) uintptr {
	var __rv uintptr
	q.Drv(302000,302103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLineEdit) OnCursorPositionChanged(fn func(int,int)) uintptr {
	var __rv uintptr
	q.Drv(302000,302104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLineEdit) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(302000,302105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLineEdit) OnTextChanged(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(302000,302106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QLineEdit) OnEditingFinished(fn func()) uintptr {
	var __rv uintptr
	q.Drv(302000,302107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QLineEdit::QLineEdit()	
func NewLineEdit() *QLineEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),302000,302108,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineEdit{}
	_p.SetDriver(__rv,302000,false)
	return _p
} 
//QLineEdit::QLineEdit(QWidget*)	
func NewLineEditWithParent(parent QWidgetInterface) *QLineEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),302000,302109,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineEdit{}
	_p.SetDriver(__rv,302000,false)
	return _p
} 
//QLineEdit::QLineEdit(QString const&,QWidget*)	
func NewLineEditWithStringParent(value2 string,parent QWidgetInterface) *QLineEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),302000,302110,unsafe.Pointer(&value2),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLineEdit{}
	_p.SetDriver(__rv,302000,false)
	return _p
} 
//QLineEdit::alignment()
func (q *QLineEdit) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(302000,302111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::backspace()
func (q *QLineEdit) Backspace()  {
	q.Drv(302000,302112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::changeEvent(QEvent*)
func (q *QLineEdit) ChangeEvent(value *QEvent)  {
	q.Drv(302000,302113,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::clear()
func (q *QLineEdit) Clear()  {
	q.Drv(302000,302114,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::completer()
func (q *QLineEdit) Completer() *QCompleter {
	var __rv uintptr
	q.Drv(302000,302115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QCompleter{}
	_rp.SetDriver(__rv,220000,false)
	return _rp
}	
//QLineEdit::contextMenuEvent(QContextMenuEvent*)
func (q *QLineEdit) ContextMenuEvent(value *QContextMenuEvent)  {
	q.Drv(302000,302116,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::copy()
func (q *QLineEdit) Copy()  {
	q.Drv(302000,302117,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::createStandardContextMenu()
func (q *QLineEdit) CreateStandardContextMenu() *QMenu {
	var __rv uintptr
	q.Drv(302000,302118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QLineEdit::cursorBackward(bool)
func (q *QLineEdit) CursorBackward(mark bool)  {
	q.Drv(302000,302119,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cursorBackward(bool,int)
func (q *QLineEdit) CursorBackwardWithMarkSteps(mark bool,steps int)  {
	q.Drv(302000,302120,unsafe.Pointer(&mark),unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cursorForward(bool)
func (q *QLineEdit) CursorForward(mark bool)  {
	q.Drv(302000,302121,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cursorForward(bool,int)
func (q *QLineEdit) CursorForwardWithMarkSteps(mark bool,steps int)  {
	q.Drv(302000,302122,unsafe.Pointer(&mark),unsafe.Pointer(&steps),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cursorPosition()
func (q *QLineEdit) CursorPosition() int {
	var __rv int
	q.Drv(302000,302123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::cursorPositionAt(QPoint const&)
func (q *QLineEdit) CursorPositionAt(pos *QPoint) int {
	var __rv int
	q.Drv(302000,302124,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::cursorRect()
func (q *QLineEdit) CursorRect() *QRect {
	var __rv uintptr
	q.Drv(302000,302125,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QLineEdit::cursorWordBackward(bool)
func (q *QLineEdit) CursorWordBackward(mark bool)  {
	q.Drv(302000,302126,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cursorWordForward(bool)
func (q *QLineEdit) CursorWordForward(mark bool)  {
	q.Drv(302000,302127,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::cut()
func (q *QLineEdit) Cut()  {
	q.Drv(302000,302128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::del()
func (q *QLineEdit) Del()  {
	q.Drv(302000,302129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::deselect()
func (q *QLineEdit) Deselect()  {
	q.Drv(302000,302130,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::displayText()
func (q *QLineEdit) DisplayText() string {
	var __rv string
	q.Drv(302000,302131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::dragEnabled()
func (q *QLineEdit) DragEnabled() bool {
	var __rv bool
	q.Drv(302000,302132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::dragEnterEvent(QDragEnterEvent*)
func (q *QLineEdit) DragEnterEvent(value *QDragEnterEvent)  {
	q.Drv(302000,302133,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::dragLeaveEvent(QDragLeaveEvent*)
func (q *QLineEdit) DragLeaveEvent(e *QDragLeaveEvent)  {
	q.Drv(302000,302134,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::dragMoveEvent(QDragMoveEvent*)
func (q *QLineEdit) DragMoveEvent(e *QDragMoveEvent)  {
	q.Drv(302000,302135,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::dropEvent(QDropEvent*)
func (q *QLineEdit) DropEvent(value *QDropEvent)  {
	q.Drv(302000,302136,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::echoMode()
func (q *QLineEdit) EchoMode() QLineEdit_EchoMode {
	var __rv QLineEdit_EchoMode
	q.Drv(302000,302137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::end(bool)
func (q *QLineEdit) End(mark bool)  {
	q.Drv(302000,302138,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::event(QEvent*)
func (q *QLineEdit) Event(value *QEvent) bool {
	var __rv bool
	q.Drv(302000,302139,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::focusInEvent(QFocusEvent*)
func (q *QLineEdit) FocusInEvent(value *QFocusEvent)  {
	q.Drv(302000,302140,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::focusOutEvent(QFocusEvent*)
func (q *QLineEdit) FocusOutEvent(value *QFocusEvent)  {
	q.Drv(302000,302141,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::getTextMargins(int*,int*,int*,int*)
func (q *QLineEdit) GetTextMargins(left *int,top *int,right *int,bottom *int)  {
	q.Drv(302000,302142,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::hasAcceptableInput()
func (q *QLineEdit) HasAcceptableInput() bool {
	var __rv bool
	q.Drv(302000,302143,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::hasFrame()
func (q *QLineEdit) HasFrame() bool {
	var __rv bool
	q.Drv(302000,302144,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::hasSelectedText()
func (q *QLineEdit) HasSelectedText() bool {
	var __rv bool
	q.Drv(302000,302145,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::home(bool)
func (q *QLineEdit) Home(mark bool)  {
	q.Drv(302000,302146,unsafe.Pointer(&mark),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::inputMask()
func (q *QLineEdit) InputMask() string {
	var __rv string
	q.Drv(302000,302147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::inputMethodEvent(QInputMethodEvent*)
func (q *QLineEdit) InputMethodEvent(value *QInputMethodEvent)  {
	q.Drv(302000,302148,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::inputMethodQuery(Qt::InputMethodQuery)
func (q *QLineEdit) InputMethodQuery(value Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(302000,302149,unsafe.Pointer(&value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QLineEdit::insert(QString const&)
func (q *QLineEdit) Insert(value string)  {
	q.Drv(302000,302150,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::isModified()
func (q *QLineEdit) IsModified() bool {
	var __rv bool
	q.Drv(302000,302151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::isReadOnly()
func (q *QLineEdit) IsReadOnly() bool {
	var __rv bool
	q.Drv(302000,302152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::isRedoAvailable()
func (q *QLineEdit) IsRedoAvailable() bool {
	var __rv bool
	q.Drv(302000,302153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::isUndoAvailable()
func (q *QLineEdit) IsUndoAvailable() bool {
	var __rv bool
	q.Drv(302000,302154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::keyPressEvent(QKeyEvent*)
func (q *QLineEdit) KeyPressEvent(value *QKeyEvent)  {
	q.Drv(302000,302155,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::maxLength()
func (q *QLineEdit) MaxLength() int {
	var __rv int
	q.Drv(302000,302156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::minimumSizeHint()
func (q *QLineEdit) MinimumSizeHint() *QSize {
	var __rv uintptr
	q.Drv(302000,302157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLineEdit::mouseDoubleClickEvent(QMouseEvent*)
func (q *QLineEdit) MouseDoubleClickEvent(value *QMouseEvent)  {
	q.Drv(302000,302158,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::mouseMoveEvent(QMouseEvent*)
func (q *QLineEdit) MouseMoveEvent(value *QMouseEvent)  {
	q.Drv(302000,302159,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::mousePressEvent(QMouseEvent*)
func (q *QLineEdit) MousePressEvent(value *QMouseEvent)  {
	q.Drv(302000,302160,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::mouseReleaseEvent(QMouseEvent*)
func (q *QLineEdit) MouseReleaseEvent(value *QMouseEvent)  {
	q.Drv(302000,302161,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::paintEvent(QPaintEvent*)
func (q *QLineEdit) PaintEvent(value *QPaintEvent)  {
	q.Drv(302000,302162,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::paste()
func (q *QLineEdit) Paste()  {
	q.Drv(302000,302163,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::placeholderText()
func (q *QLineEdit) PlaceholderText() string {
	var __rv string
	q.Drv(302000,302164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::redo()
func (q *QLineEdit) Redo()  {
	q.Drv(302000,302165,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::selectAll()
func (q *QLineEdit) SelectAll()  {
	q.Drv(302000,302166,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::selectedText()
func (q *QLineEdit) SelectedText() string {
	var __rv string
	q.Drv(302000,302167,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::selectionStart()
func (q *QLineEdit) SelectionStart() int {
	var __rv int
	q.Drv(302000,302168,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QLineEdit) SetAlignment(flag Qt_AlignmentFlag)  {
	q.Drv(302000,302169,unsafe.Pointer(&flag),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setCompleter(QCompleter*)
func (q *QLineEdit) SetCompleter(completer *QCompleter)  {
	q.Drv(302000,302170,Native(completer),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setCursorPosition(int)
func (q *QLineEdit) SetCursorPosition(value int)  {
	q.Drv(302000,302171,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setDragEnabled(bool)
func (q *QLineEdit) SetDragEnabled(b bool)  {
	q.Drv(302000,302172,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setEchoMode(QLineEdit::EchoMode)
func (q *QLineEdit) SetEchoMode(value QLineEdit_EchoMode)  {
	q.Drv(302000,302173,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setFrame(bool)
func (q *QLineEdit) SetFrame(value bool)  {
	q.Drv(302000,302174,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setInputMask(QString const&)
func (q *QLineEdit) SetInputMask(inputMask string)  {
	q.Drv(302000,302175,unsafe.Pointer(&inputMask),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setMaxLength(int)
func (q *QLineEdit) SetMaxLength(value int)  {
	q.Drv(302000,302176,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setModified(bool)
func (q *QLineEdit) SetModified(value bool)  {
	q.Drv(302000,302177,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setPlaceholderText(QString const&)
func (q *QLineEdit) SetPlaceholderText(value string)  {
	q.Drv(302000,302178,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setReadOnly(bool)
func (q *QLineEdit) SetReadOnly(value bool)  {
	q.Drv(302000,302179,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setSelection(int,int)
func (q *QLineEdit) SetSelection(value2 int,value3 int)  {
	q.Drv(302000,302180,unsafe.Pointer(&value2),unsafe.Pointer(&value3),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setText(QString const&)
func (q *QLineEdit) SetText(value string)  {
	q.Drv(302000,302181,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setTextMargins(QMargins const&)
func (q *QLineEdit) SetTextMargins(margins *QMargins)  {
	q.Drv(302000,302182,Native(margins),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setTextMargins(int,int,int,int)
func (q *QLineEdit) SetTextMarginsWithLeftTopRightBottom(left int,top int,right int,bottom int)  {
	q.Drv(302000,302183,unsafe.Pointer(&left),unsafe.Pointer(&top),unsafe.Pointer(&right),unsafe.Pointer(&bottom),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::setValidator(QValidator const*)
func (q *QLineEdit) SetValidator(value *QValidator)  {
	q.Drv(302000,302184,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::sizeHint()
func (q *QLineEdit) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(302000,302185,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QLineEdit::text()
func (q *QLineEdit) Text() string {
	var __rv string
	q.Drv(302000,302186,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLineEdit::textMargins()
func (q *QLineEdit) TextMargins() *QMargins {
	var __rv uintptr
	q.Drv(302000,302187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMargins{}
	_rp.SetDriver(__rv,73000,true)
	return _rp
}	
//QLineEdit::undo()
func (q *QLineEdit) Undo()  {
	q.Drv(302000,302188,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLineEdit::validator()
func (q *QLineEdit) Validator() *QValidator {
	var __rv uintptr
	q.Drv(302000,302189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QValidator{}
	_rp.SetDriver(__rv,393000,false)
	return _rp
}	
