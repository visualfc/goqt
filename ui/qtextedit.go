// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextEdit_AutoFormattingFlag - QTextEdit::AutoFormattingFlag
type QTextEdit_AutoFormattingFlag uint32
const (
	QTextEdit_AutoNone QTextEdit_AutoFormattingFlag = 0
	QTextEdit_AutoBulletList QTextEdit_AutoFormattingFlag = 0x00000001
	QTextEdit_AutoAll QTextEdit_AutoFormattingFlag = 0xffffffff
)
//enum QTextEdit_LineWrapMode - QTextEdit::LineWrapMode
type QTextEdit_LineWrapMode uint32
const (
	QTextEdit_NoWrap QTextEdit_LineWrapMode = 0
	QTextEdit_WidgetWidth QTextEdit_LineWrapMode = 1
	QTextEdit_FixedPixelWidth QTextEdit_LineWrapMode = 2
	QTextEdit_FixedColumnWidth QTextEdit_LineWrapMode = 3
)
//struct QTextEdit : QTextEdit
type QTextEdit struct {
	QAbstractScrollArea
}
func (q *QTextEdit) OnCopyAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(373000,373102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnUndoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(373000,373103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnCursorPositionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(373000,373104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(373000,373105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnTextChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(373000,373106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnRedoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(373000,373107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextEdit) OnCurrentCharFormatChanged(fn func(*QTextCharFormat)) uintptr {
	var __rv uintptr
	q.Drv(373000,373108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTextEdit::QTextEdit()	
func NewTextEdit() *QTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),373000,373109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextEdit{}
	_p.SetDriver(__rv,373000,false)
	return _p
} 
//QTextEdit::QTextEdit(QWidget*)	
func NewTextEditWithParent(parent QWidgetInterface) *QTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),373000,373110,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextEdit{}
	_p.SetDriver(__rv,373000,false)
	return _p
} 
//QTextEdit::QTextEdit(QString const&,QWidget*)	
func NewTextEditWithTextParent(text string,parent QWidgetInterface) *QTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),373000,373111,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextEdit{}
	_p.SetDriver(__rv,373000,false)
	return _p
} 
//QTextEdit::acceptRichText()
func (q *QTextEdit) AcceptRichText() bool {
	var __rv bool
	q.Drv(373000,373112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::alignment()
func (q *QTextEdit) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(373000,373113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::anchorAt(QPoint const&)
func (q *QTextEdit) AnchorAt(pos *QPoint) string {
	var __rv string
	q.Drv(373000,373114,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::append(QString const&)
func (q *QTextEdit) Append(text string)  {
	q.Drv(373000,373115,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::autoFormatting()
func (q *QTextEdit) AutoFormatting() QTextEdit_AutoFormattingFlag {
	var __rv QTextEdit_AutoFormattingFlag
	q.Drv(373000,373116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::canInsertFromMimeData(QMimeData const*)
func (q *QTextEdit) CanInsertFromMimeData(source *QMimeData) bool {
	var __rv bool
	q.Drv(373000,373117,Native(source),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::canPaste()
func (q *QTextEdit) CanPaste() bool {
	var __rv bool
	q.Drv(373000,373118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::changeEvent(QEvent*)
func (q *QTextEdit) ChangeEvent(e *QEvent)  {
	q.Drv(373000,373119,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::clear()
func (q *QTextEdit) Clear()  {
	q.Drv(373000,373120,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::contextMenuEvent(QContextMenuEvent*)
func (q *QTextEdit) ContextMenuEvent(e *QContextMenuEvent)  {
	q.Drv(373000,373121,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::copy()
func (q *QTextEdit) Copy()  {
	q.Drv(373000,373122,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::createMimeDataFromSelection()
func (q *QTextEdit) CreateMimeDataFromSelection() *QMimeData {
	var __rv uintptr
	q.Drv(373000,373123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QTextEdit::createStandardContextMenu()
func (q *QTextEdit) CreateStandardContextMenu() *QMenu {
	var __rv uintptr
	q.Drv(373000,373124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QTextEdit::createStandardContextMenu(QPoint const&)
func (q *QTextEdit) CreateStandardContextMenuWithPosition(position *QPoint) *QMenu {
	var __rv uintptr
	q.Drv(373000,373125,Native(position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QTextEdit::currentCharFormat()
func (q *QTextEdit) CurrentCharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(373000,373126,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextEdit::currentFont()
func (q *QTextEdit) CurrentFont() *QFont {
	var __rv uintptr
	q.Drv(373000,373127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTextEdit::cursorForPosition(QPoint const&)
func (q *QTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	var __rv uintptr
	q.Drv(373000,373128,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextEdit::cursorRect()
func (q *QTextEdit) CursorRect() *QRect {
	var __rv uintptr
	q.Drv(373000,373129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTextEdit::cursorRect(QTextCursor const&)
func (q *QTextEdit) CursorRectWithCursor(cursor *QTextCursor) *QRect {
	var __rv uintptr
	q.Drv(373000,373130,Native(cursor),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QTextEdit::cursorWidth()
func (q *QTextEdit) CursorWidth() int {
	var __rv int
	q.Drv(373000,373131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::cut()
func (q *QTextEdit) Cut()  {
	q.Drv(373000,373132,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::document()
func (q *QTextEdit) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(373000,373133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextEdit::documentTitle()
func (q *QTextEdit) DocumentTitle() string {
	var __rv string
	q.Drv(373000,373134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::dragEnterEvent(QDragEnterEvent*)
func (q *QTextEdit) DragEnterEvent(e *QDragEnterEvent)  {
	q.Drv(373000,373135,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::dragLeaveEvent(QDragLeaveEvent*)
func (q *QTextEdit) DragLeaveEvent(e *QDragLeaveEvent)  {
	q.Drv(373000,373136,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::dragMoveEvent(QDragMoveEvent*)
func (q *QTextEdit) DragMoveEvent(e *QDragMoveEvent)  {
	q.Drv(373000,373137,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::dropEvent(QDropEvent*)
func (q *QTextEdit) DropEvent(e *QDropEvent)  {
	q.Drv(373000,373138,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::ensureCursorVisible()
func (q *QTextEdit) EnsureCursorVisible()  {
	q.Drv(373000,373139,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::event(QEvent*)
func (q *QTextEdit) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(373000,373140,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::extraSelections()
func (q *QTextEdit) ExtraSelections() []*QTextEditExtraSelection {
	var __rv []*QTextEditExtraSelection
	q.Drv(373000,373141,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::find(QString const&)
func (q *QTextEdit) Find(exp string) bool {
	var __rv bool
	q.Drv(373000,373142,unsafe.Pointer(&exp),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::find(QString const&,QFlags<QTextDocument::FindFlag>)
func (q *QTextEdit) FindWithExpOptions(exp string,options QTextDocument_FindFlag) bool {
	var __rv bool
	q.Drv(373000,373143,unsafe.Pointer(&exp),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::focusInEvent(QFocusEvent*)
func (q *QTextEdit) FocusInEvent(e *QFocusEvent)  {
	q.Drv(373000,373144,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::focusNextPrevChild(bool)
func (q *QTextEdit) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(373000,373145,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::focusOutEvent(QFocusEvent*)
func (q *QTextEdit) FocusOutEvent(e *QFocusEvent)  {
	q.Drv(373000,373146,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::fontFamily()
func (q *QTextEdit) FontFamily() string {
	var __rv string
	q.Drv(373000,373147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::fontItalic()
func (q *QTextEdit) FontItalic() bool {
	var __rv bool
	q.Drv(373000,373148,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::fontPointSize()
func (q *QTextEdit) FontPointSize() float64 {
	var __rv float64
	q.Drv(373000,373149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::fontUnderline()
func (q *QTextEdit) FontUnderline() bool {
	var __rv bool
	q.Drv(373000,373150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::fontWeight()
func (q *QTextEdit) FontWeight() int {
	var __rv int
	q.Drv(373000,373151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::inputMethodEvent(QInputMethodEvent*)
func (q *QTextEdit) InputMethodEvent(value *QInputMethodEvent)  {
	q.Drv(373000,373152,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::inputMethodQuery(Qt::InputMethodQuery)
func (q *QTextEdit) InputMethodQuery(property Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(373000,373153,unsafe.Pointer(&property),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextEdit::insertFromMimeData(QMimeData const*)
func (q *QTextEdit) InsertFromMimeData(source *QMimeData)  {
	q.Drv(373000,373154,Native(source),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::insertHtml(QString const&)
func (q *QTextEdit) InsertHtml(text string)  {
	q.Drv(373000,373155,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::insertPlainText(QString const&)
func (q *QTextEdit) InsertPlainText(text string)  {
	q.Drv(373000,373156,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::isReadOnly()
func (q *QTextEdit) IsReadOnly() bool {
	var __rv bool
	q.Drv(373000,373157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::isUndoRedoEnabled()
func (q *QTextEdit) IsUndoRedoEnabled() bool {
	var __rv bool
	q.Drv(373000,373158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::keyPressEvent(QKeyEvent*)
func (q *QTextEdit) KeyPressEvent(e *QKeyEvent)  {
	q.Drv(373000,373159,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::keyReleaseEvent(QKeyEvent*)
func (q *QTextEdit) KeyReleaseEvent(e *QKeyEvent)  {
	q.Drv(373000,373160,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::lineWrapColumnOrWidth()
func (q *QTextEdit) LineWrapColumnOrWidth() int {
	var __rv int
	q.Drv(373000,373161,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::lineWrapMode()
func (q *QTextEdit) LineWrapMode() QTextEdit_LineWrapMode {
	var __rv QTextEdit_LineWrapMode
	q.Drv(373000,373162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::loadResource(int,QUrl const&)
func (q *QTextEdit) LoadResource(_type int,name *QUrl) *QVariant {
	var __rv uintptr
	q.Drv(373000,373163,unsafe.Pointer(&_type),Native(name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextEdit::mergeCurrentCharFormat(QTextCharFormat const&)
func (q *QTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat)  {
	q.Drv(373000,373164,Native(modifier),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::mouseDoubleClickEvent(QMouseEvent*)
func (q *QTextEdit) MouseDoubleClickEvent(e *QMouseEvent)  {
	q.Drv(373000,373165,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::mouseMoveEvent(QMouseEvent*)
func (q *QTextEdit) MouseMoveEvent(e *QMouseEvent)  {
	q.Drv(373000,373166,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::mousePressEvent(QMouseEvent*)
func (q *QTextEdit) MousePressEvent(e *QMouseEvent)  {
	q.Drv(373000,373167,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::mouseReleaseEvent(QMouseEvent*)
func (q *QTextEdit) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(373000,373168,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::moveCursor(QTextCursor::MoveOperation)
func (q *QTextEdit) MoveCursor(operation QTextCursor_MoveOperation)  {
	q.Drv(373000,373169,unsafe.Pointer(&operation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::moveCursor(QTextCursor::MoveOperation,QTextCursor::MoveMode)
func (q *QTextEdit) MoveCursorWithOperationMode(operation QTextCursor_MoveOperation,mode QTextCursor_MoveMode)  {
	q.Drv(373000,373170,unsafe.Pointer(&operation),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::overwriteMode()
func (q *QTextEdit) OverwriteMode() bool {
	var __rv bool
	q.Drv(373000,373171,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::paintEvent(QPaintEvent*)
func (q *QTextEdit) PaintEvent(e *QPaintEvent)  {
	q.Drv(373000,373172,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::paste()
func (q *QTextEdit) Paste()  {
	q.Drv(373000,373173,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::print(QPrinter*)
func (q *QTextEdit) Print(printer *QPrinter)  {
	q.Drv(373000,373174,Native(printer),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::redo()
func (q *QTextEdit) Redo()  {
	q.Drv(373000,373175,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::resizeEvent(QResizeEvent*)
func (q *QTextEdit) ResizeEvent(e *QResizeEvent)  {
	q.Drv(373000,373176,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::scrollContentsBy(int,int)
func (q *QTextEdit) ScrollContentsBy(dx int,dy int)  {
	q.Drv(373000,373177,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::scrollToAnchor(QString const&)
func (q *QTextEdit) ScrollToAnchor(name string)  {
	q.Drv(373000,373178,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::selectAll()
func (q *QTextEdit) SelectAll()  {
	q.Drv(373000,373179,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setAcceptRichText(bool)
func (q *QTextEdit) SetAcceptRichText(accept bool)  {
	q.Drv(373000,373180,unsafe.Pointer(&accept),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QTextEdit) SetAlignment(a Qt_AlignmentFlag)  {
	q.Drv(373000,373181,unsafe.Pointer(&a),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setAutoFormatting(QFlags<QTextEdit::AutoFormattingFlag>)
func (q *QTextEdit) SetAutoFormatting(features QTextEdit_AutoFormattingFlag)  {
	q.Drv(373000,373182,unsafe.Pointer(&features),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setCurrentCharFormat(QTextCharFormat const&)
func (q *QTextEdit) SetCurrentCharFormat(format *QTextCharFormat)  {
	q.Drv(373000,373183,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setCurrentFont(QFont const&)
func (q *QTextEdit) SetCurrentFont(f *QFont)  {
	q.Drv(373000,373184,Native(f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setCursorWidth(int)
func (q *QTextEdit) SetCursorWidth(width int)  {
	q.Drv(373000,373185,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setDocument(QTextDocument*)
func (q *QTextEdit) SetDocument(document *QTextDocument)  {
	q.Drv(373000,373186,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setDocumentTitle(QString const&)
func (q *QTextEdit) SetDocumentTitle(title string)  {
	q.Drv(373000,373187,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setExtraSelections(QList<QTextEdit::ExtraSelection> const&)
func (q *QTextEdit) SetExtraSelections(selections []*QTextEditExtraSelection)  {
	q.Drv(373000,373188,unsafe.Pointer(&selections),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setFontFamily(QString const&)
func (q *QTextEdit) SetFontFamily(fontFamily string)  {
	q.Drv(373000,373189,unsafe.Pointer(&fontFamily),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setFontItalic(bool)
func (q *QTextEdit) SetFontItalic(b bool)  {
	q.Drv(373000,373190,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setFontPointSize(double)
func (q *QTextEdit) SetFontPointSize(s float64)  {
	q.Drv(373000,373191,unsafe.Pointer(&s),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setFontUnderline(bool)
func (q *QTextEdit) SetFontUnderline(b bool)  {
	q.Drv(373000,373192,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setFontWeight(int)
func (q *QTextEdit) SetFontWeight(w int)  {
	q.Drv(373000,373193,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setHtml(QString const&)
func (q *QTextEdit) SetHtml(text string)  {
	q.Drv(373000,373194,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setLineWrapColumnOrWidth(int)
func (q *QTextEdit) SetLineWrapColumnOrWidth(w int)  {
	q.Drv(373000,373195,unsafe.Pointer(&w),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setLineWrapMode(QTextEdit::LineWrapMode)
func (q *QTextEdit) SetLineWrapMode(mode QTextEdit_LineWrapMode)  {
	q.Drv(373000,373196,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setOverwriteMode(bool)
func (q *QTextEdit) SetOverwriteMode(overwrite bool)  {
	q.Drv(373000,373197,unsafe.Pointer(&overwrite),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setPlainText(QString const&)
func (q *QTextEdit) SetPlainText(text string)  {
	q.Drv(373000,373198,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setReadOnly(bool)
func (q *QTextEdit) SetReadOnly(ro bool)  {
	q.Drv(373000,373199,unsafe.Pointer(&ro),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTabChangesFocus(bool)
func (q *QTextEdit) SetTabChangesFocus(b bool)  {
	q.Drv(373000,373200,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTabStopWidth(int)
func (q *QTextEdit) SetTabStopWidth(width int)  {
	q.Drv(373000,373201,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setText(QString const&)
func (q *QTextEdit) SetText(text string)  {
	q.Drv(373000,373202,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTextBackgroundColor(QColor const&)
func (q *QTextEdit) SetTextBackgroundColor(c *QColor)  {
	q.Drv(373000,373203,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTextColor(QColor const&)
func (q *QTextEdit) SetTextColor(c *QColor)  {
	q.Drv(373000,373204,Native(c),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTextCursor(QTextCursor const&)
func (q *QTextEdit) SetTextCursor(cursor *QTextCursor)  {
	q.Drv(373000,373205,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>)
func (q *QTextEdit) SetTextInteractionFlags(flags Qt_TextInteractionFlag)  {
	q.Drv(373000,373206,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setUndoRedoEnabled(bool)
func (q *QTextEdit) SetUndoRedoEnabled(enable bool)  {
	q.Drv(373000,373207,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::setWordWrapMode(QTextOption::WrapMode)
func (q *QTextEdit) SetWordWrapMode(policy QTextOption_WrapMode)  {
	q.Drv(373000,373208,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::showEvent(QShowEvent*)
func (q *QTextEdit) ShowEvent(value *QShowEvent)  {
	q.Drv(373000,373209,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::tabChangesFocus()
func (q *QTextEdit) TabChangesFocus() bool {
	var __rv bool
	q.Drv(373000,373210,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::tabStopWidth()
func (q *QTextEdit) TabStopWidth() int {
	var __rv int
	q.Drv(373000,373211,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::textBackgroundColor()
func (q *QTextEdit) TextBackgroundColor() *QColor {
	var __rv uintptr
	q.Drv(373000,373212,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTextEdit::textColor()
func (q *QTextEdit) TextColor() *QColor {
	var __rv uintptr
	q.Drv(373000,373213,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QTextEdit::textCursor()
func (q *QTextEdit) TextCursor() *QTextCursor {
	var __rv uintptr
	q.Drv(373000,373214,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextEdit::textInteractionFlags()
func (q *QTextEdit) TextInteractionFlags() Qt_TextInteractionFlag {
	var __rv Qt_TextInteractionFlag
	q.Drv(373000,373215,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::timerEvent(QTimerEvent*)
func (q *QTextEdit) TimerEvent(e *QTimerEvent)  {
	q.Drv(373000,373216,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::toHtml()
func (q *QTextEdit) ToHtml() string {
	var __rv string
	q.Drv(373000,373217,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::toPlainText()
func (q *QTextEdit) ToPlainText() string {
	var __rv string
	q.Drv(373000,373218,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::undo()
func (q *QTextEdit) Undo()  {
	q.Drv(373000,373219,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::wheelEvent(QWheelEvent*)
func (q *QTextEdit) WheelEvent(e *QWheelEvent)  {
	q.Drv(373000,373220,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::wordWrapMode()
func (q *QTextEdit) WordWrapMode() QTextOption_WrapMode {
	var __rv QTextOption_WrapMode
	q.Drv(373000,373221,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextEdit::zoomIn()
func (q *QTextEdit) ZoomIn()  {
	q.Drv(373000,373222,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::zoomIn(int)
func (q *QTextEdit) ZoomInWithRange(_range int)  {
	q.Drv(373000,373223,unsafe.Pointer(&_range),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::zoomOut()
func (q *QTextEdit) ZoomOut()  {
	q.Drv(373000,373224,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextEdit::zoomOut(int)
func (q *QTextEdit) ZoomOutWithRange(_range int)  {
	q.Drv(373000,373225,unsafe.Pointer(&_range),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
