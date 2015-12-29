// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPlainTextEdit_LineWrapMode - QPlainTextEdit::LineWrapMode
type QPlainTextEdit_LineWrapMode uint32
const (
	QPlainTextEdit_NoWrap QPlainTextEdit_LineWrapMode = 0
	QPlainTextEdit_WidgetWidth QPlainTextEdit_LineWrapMode = 1
)
//struct QPlainTextEdit : QPlainTextEdit
type QPlainTextEdit struct {
	QAbstractScrollArea
}
func (q *QPlainTextEdit) OnCopyAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(321000,321102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnUndoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(321000,321103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnBlockCountChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(321000,321104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnUpdateRequest(fn func(*QRect,int)) uintptr {
	var __rv uintptr
	q.Drv(321000,321105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnModificationChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(321000,321106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnCursorPositionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(321000,321107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnSelectionChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(321000,321108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnTextChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(321000,321109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QPlainTextEdit) OnRedoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(321000,321110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QPlainTextEdit::QPlainTextEdit()	
func NewPlainTextEdit() *QPlainTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),321000,321111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPlainTextEdit{}
	_p.SetDriver(__rv,321000,false)
	return _p
} 
//QPlainTextEdit::QPlainTextEdit(QWidget*)	
func NewPlainTextEditWithParent(parent QWidgetInterface) *QPlainTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),321000,321112,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPlainTextEdit{}
	_p.SetDriver(__rv,321000,false)
	return _p
} 
//QPlainTextEdit::QPlainTextEdit(QString const&,QWidget*)	
func NewPlainTextEditWithTextParent(text string,parent QWidgetInterface) *QPlainTextEdit {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),321000,321113,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QPlainTextEdit{}
	_p.SetDriver(__rv,321000,false)
	return _p
} 
//QPlainTextEdit::anchorAt(QPoint const&)
func (q *QPlainTextEdit) AnchorAt(pos *QPoint) string {
	var __rv string
	q.Drv(321000,321114,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::appendHtml(QString const&)
func (q *QPlainTextEdit) AppendHtml(html string)  {
	q.Drv(321000,321115,unsafe.Pointer(&html),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::appendPlainText(QString const&)
func (q *QPlainTextEdit) AppendPlainText(text string)  {
	q.Drv(321000,321116,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::backgroundVisible()
func (q *QPlainTextEdit) BackgroundVisible() bool {
	var __rv bool
	q.Drv(321000,321117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::blockBoundingGeometry(QTextBlock const&)
func (q *QPlainTextEdit) BlockBoundingGeometry(block *QTextBlock) *QRectF {
	var __rv uintptr
	q.Drv(321000,321118,Native(block),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPlainTextEdit::blockBoundingRect(QTextBlock const&)
func (q *QPlainTextEdit) BlockBoundingRect(block *QTextBlock) *QRectF {
	var __rv uintptr
	q.Drv(321000,321119,Native(block),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRectF{}
	_rp.SetDriver(__rv,111000,true)
	return _rp
}	
//QPlainTextEdit::blockCount()
func (q *QPlainTextEdit) BlockCount() int {
	var __rv int
	q.Drv(321000,321120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::canInsertFromMimeData(QMimeData const*)
func (q *QPlainTextEdit) CanInsertFromMimeData(source *QMimeData) bool {
	var __rv bool
	q.Drv(321000,321121,Native(source),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::canPaste()
func (q *QPlainTextEdit) CanPaste() bool {
	var __rv bool
	q.Drv(321000,321122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::centerCursor()
func (q *QPlainTextEdit) CenterCursor()  {
	q.Drv(321000,321123,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::centerOnScroll()
func (q *QPlainTextEdit) CenterOnScroll() bool {
	var __rv bool
	q.Drv(321000,321124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::changeEvent(QEvent*)
func (q *QPlainTextEdit) ChangeEvent(e *QEvent)  {
	q.Drv(321000,321125,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::clear()
func (q *QPlainTextEdit) Clear()  {
	q.Drv(321000,321126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::contentOffset()
func (q *QPlainTextEdit) ContentOffset() *QPointF {
	var __rv uintptr
	q.Drv(321000,321127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPointF{}
	_rp.SetDriver(__rv,100000,true)
	return _rp
}	
//QPlainTextEdit::contextMenuEvent(QContextMenuEvent*)
func (q *QPlainTextEdit) ContextMenuEvent(e *QContextMenuEvent)  {
	q.Drv(321000,321128,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::copy()
func (q *QPlainTextEdit) Copy()  {
	q.Drv(321000,321129,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::createMimeDataFromSelection()
func (q *QPlainTextEdit) CreateMimeDataFromSelection() *QMimeData {
	var __rv uintptr
	q.Drv(321000,321130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMimeData{}
	_rp.SetDriver(__rv,311000,false)
	return _rp
}	
//QPlainTextEdit::createStandardContextMenu()
func (q *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	var __rv uintptr
	q.Drv(321000,321131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QMenu{}
	_rp.SetDriver(__rv,308000,false)
	return _rp
}	
//QPlainTextEdit::currentCharFormat()
func (q *QPlainTextEdit) CurrentCharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(321000,321132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QPlainTextEdit::cursorForPosition(QPoint const&)
func (q *QPlainTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	var __rv uintptr
	q.Drv(321000,321133,Native(pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QPlainTextEdit::cursorRect()
func (q *QPlainTextEdit) CursorRect() *QRect {
	var __rv uintptr
	q.Drv(321000,321134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPlainTextEdit::cursorRect(QTextCursor const&)
func (q *QPlainTextEdit) CursorRectWithCursor(cursor *QTextCursor) *QRect {
	var __rv uintptr
	q.Drv(321000,321135,Native(cursor),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QRect{}
	_rp.SetDriver(__rv,110000,true)
	return _rp
}	
//QPlainTextEdit::cursorWidth()
func (q *QPlainTextEdit) CursorWidth() int {
	var __rv int
	q.Drv(321000,321136,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::cut()
func (q *QPlainTextEdit) Cut()  {
	q.Drv(321000,321137,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::document()
func (q *QPlainTextEdit) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(321000,321138,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QPlainTextEdit::documentTitle()
func (q *QPlainTextEdit) DocumentTitle() string {
	var __rv string
	q.Drv(321000,321139,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::dragEnterEvent(QDragEnterEvent*)
func (q *QPlainTextEdit) DragEnterEvent(e *QDragEnterEvent)  {
	q.Drv(321000,321140,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::dragLeaveEvent(QDragLeaveEvent*)
func (q *QPlainTextEdit) DragLeaveEvent(e *QDragLeaveEvent)  {
	q.Drv(321000,321141,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::dragMoveEvent(QDragMoveEvent*)
func (q *QPlainTextEdit) DragMoveEvent(e *QDragMoveEvent)  {
	q.Drv(321000,321142,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::dropEvent(QDropEvent*)
func (q *QPlainTextEdit) DropEvent(e *QDropEvent)  {
	q.Drv(321000,321143,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::ensureCursorVisible()
func (q *QPlainTextEdit) EnsureCursorVisible()  {
	q.Drv(321000,321144,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::event(QEvent*)
func (q *QPlainTextEdit) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(321000,321145,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::extraSelections()
func (q *QPlainTextEdit) ExtraSelections() []*QTextEditExtraSelection {
	var __rv []*QTextEditExtraSelection
	q.Drv(321000,321146,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::find(QString const&)
func (q *QPlainTextEdit) Find(exp string) bool {
	var __rv bool
	q.Drv(321000,321147,unsafe.Pointer(&exp),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::find(QString const&,QFlags<QTextDocument::FindFlag>)
func (q *QPlainTextEdit) FindWithExpOptions(exp string,options QTextDocument_FindFlag) bool {
	var __rv bool
	q.Drv(321000,321148,unsafe.Pointer(&exp),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::firstVisibleBlock()
func (q *QPlainTextEdit) FirstVisibleBlock() *QTextBlock {
	var __rv uintptr
	q.Drv(321000,321149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QPlainTextEdit::focusInEvent(QFocusEvent*)
func (q *QPlainTextEdit) FocusInEvent(e *QFocusEvent)  {
	q.Drv(321000,321150,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::focusNextPrevChild(bool)
func (q *QPlainTextEdit) FocusNextPrevChild(next bool) bool {
	var __rv bool
	q.Drv(321000,321151,unsafe.Pointer(&next),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::focusOutEvent(QFocusEvent*)
func (q *QPlainTextEdit) FocusOutEvent(e *QFocusEvent)  {
	q.Drv(321000,321152,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::getPaintContext()
func (q *QPlainTextEdit) GetPaintContext() *QAbstractTextDocumentLayoutPaintContext {
	var __rv uintptr
	q.Drv(321000,321153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractTextDocumentLayoutPaintContext{}
	_rp.SetDriver(__rv,2000,true)
	return _rp
}	
//QPlainTextEdit::inputMethodEvent(QInputMethodEvent*)
func (q *QPlainTextEdit) InputMethodEvent(value *QInputMethodEvent)  {
	q.Drv(321000,321154,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::inputMethodQuery(Qt::InputMethodQuery)
func (q *QPlainTextEdit) InputMethodQuery(property Qt_InputMethodQuery) *QVariant {
	var __rv uintptr
	q.Drv(321000,321155,unsafe.Pointer(&property),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QPlainTextEdit::insertFromMimeData(QMimeData const*)
func (q *QPlainTextEdit) InsertFromMimeData(source *QMimeData)  {
	q.Drv(321000,321156,Native(source),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::insertPlainText(QString const&)
func (q *QPlainTextEdit) InsertPlainText(text string)  {
	q.Drv(321000,321157,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::isReadOnly()
func (q *QPlainTextEdit) IsReadOnly() bool {
	var __rv bool
	q.Drv(321000,321158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::isUndoRedoEnabled()
func (q *QPlainTextEdit) IsUndoRedoEnabled() bool {
	var __rv bool
	q.Drv(321000,321159,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::keyPressEvent(QKeyEvent*)
func (q *QPlainTextEdit) KeyPressEvent(e *QKeyEvent)  {
	q.Drv(321000,321160,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::keyReleaseEvent(QKeyEvent*)
func (q *QPlainTextEdit) KeyReleaseEvent(e *QKeyEvent)  {
	q.Drv(321000,321161,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::lineWrapMode()
func (q *QPlainTextEdit) LineWrapMode() QPlainTextEdit_LineWrapMode {
	var __rv QPlainTextEdit_LineWrapMode
	q.Drv(321000,321162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::loadResource(int,QUrl const&)
func (q *QPlainTextEdit) LoadResource(_type int,name *QUrl) *QVariant {
	var __rv uintptr
	q.Drv(321000,321163,unsafe.Pointer(&_type),Native(name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QPlainTextEdit::maximumBlockCount()
func (q *QPlainTextEdit) MaximumBlockCount() int {
	var __rv int
	q.Drv(321000,321164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::mergeCurrentCharFormat(QTextCharFormat const&)
func (q *QPlainTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat)  {
	q.Drv(321000,321165,Native(modifier),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::mouseDoubleClickEvent(QMouseEvent*)
func (q *QPlainTextEdit) MouseDoubleClickEvent(e *QMouseEvent)  {
	q.Drv(321000,321166,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::mouseMoveEvent(QMouseEvent*)
func (q *QPlainTextEdit) MouseMoveEvent(e *QMouseEvent)  {
	q.Drv(321000,321167,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::mousePressEvent(QMouseEvent*)
func (q *QPlainTextEdit) MousePressEvent(e *QMouseEvent)  {
	q.Drv(321000,321168,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::mouseReleaseEvent(QMouseEvent*)
func (q *QPlainTextEdit) MouseReleaseEvent(e *QMouseEvent)  {
	q.Drv(321000,321169,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::moveCursor(QTextCursor::MoveOperation)
func (q *QPlainTextEdit) MoveCursor(operation QTextCursor_MoveOperation)  {
	q.Drv(321000,321170,unsafe.Pointer(&operation),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::moveCursor(QTextCursor::MoveOperation,QTextCursor::MoveMode)
func (q *QPlainTextEdit) MoveCursorWithOperationMode(operation QTextCursor_MoveOperation,mode QTextCursor_MoveMode)  {
	q.Drv(321000,321171,unsafe.Pointer(&operation),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::overwriteMode()
func (q *QPlainTextEdit) OverwriteMode() bool {
	var __rv bool
	q.Drv(321000,321172,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::paintEvent(QPaintEvent*)
func (q *QPlainTextEdit) PaintEvent(e *QPaintEvent)  {
	q.Drv(321000,321173,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::paste()
func (q *QPlainTextEdit) Paste()  {
	q.Drv(321000,321174,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::print(QPrinter*)
func (q *QPlainTextEdit) Print(printer *QPrinter)  {
	q.Drv(321000,321175,Native(printer),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::redo()
func (q *QPlainTextEdit) Redo()  {
	q.Drv(321000,321176,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::resizeEvent(QResizeEvent*)
func (q *QPlainTextEdit) ResizeEvent(e *QResizeEvent)  {
	q.Drv(321000,321177,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::scrollContentsBy(int,int)
func (q *QPlainTextEdit) ScrollContentsBy(dx int,dy int)  {
	q.Drv(321000,321178,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::selectAll()
func (q *QPlainTextEdit) SelectAll()  {
	q.Drv(321000,321179,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setBackgroundVisible(bool)
func (q *QPlainTextEdit) SetBackgroundVisible(visible bool)  {
	q.Drv(321000,321180,unsafe.Pointer(&visible),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setCenterOnScroll(bool)
func (q *QPlainTextEdit) SetCenterOnScroll(enabled bool)  {
	q.Drv(321000,321181,unsafe.Pointer(&enabled),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setCurrentCharFormat(QTextCharFormat const&)
func (q *QPlainTextEdit) SetCurrentCharFormat(format *QTextCharFormat)  {
	q.Drv(321000,321182,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setCursorWidth(int)
func (q *QPlainTextEdit) SetCursorWidth(width int)  {
	q.Drv(321000,321183,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setDocument(QTextDocument*)
func (q *QPlainTextEdit) SetDocument(document *QTextDocument)  {
	q.Drv(321000,321184,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setDocumentTitle(QString const&)
func (q *QPlainTextEdit) SetDocumentTitle(title string)  {
	q.Drv(321000,321185,unsafe.Pointer(&title),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setExtraSelections(QList<QTextEdit::ExtraSelection> const&)
func (q *QPlainTextEdit) SetExtraSelections(selections []*QTextEditExtraSelection)  {
	q.Drv(321000,321186,unsafe.Pointer(&selections),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setLineWrapMode(QPlainTextEdit::LineWrapMode)
func (q *QPlainTextEdit) SetLineWrapMode(mode QPlainTextEdit_LineWrapMode)  {
	q.Drv(321000,321187,unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setMaximumBlockCount(int)
func (q *QPlainTextEdit) SetMaximumBlockCount(maximum int)  {
	q.Drv(321000,321188,unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setOverwriteMode(bool)
func (q *QPlainTextEdit) SetOverwriteMode(overwrite bool)  {
	q.Drv(321000,321189,unsafe.Pointer(&overwrite),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setPlainText(QString const&)
func (q *QPlainTextEdit) SetPlainText(text string)  {
	q.Drv(321000,321190,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setReadOnly(bool)
func (q *QPlainTextEdit) SetReadOnly(ro bool)  {
	q.Drv(321000,321191,unsafe.Pointer(&ro),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setTabChangesFocus(bool)
func (q *QPlainTextEdit) SetTabChangesFocus(b bool)  {
	q.Drv(321000,321192,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setTabStopWidth(int)
func (q *QPlainTextEdit) SetTabStopWidth(width int)  {
	q.Drv(321000,321193,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setTextCursor(QTextCursor const&)
func (q *QPlainTextEdit) SetTextCursor(cursor *QTextCursor)  {
	q.Drv(321000,321194,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>)
func (q *QPlainTextEdit) SetTextInteractionFlags(flags Qt_TextInteractionFlag)  {
	q.Drv(321000,321195,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setUndoRedoEnabled(bool)
func (q *QPlainTextEdit) SetUndoRedoEnabled(enable bool)  {
	q.Drv(321000,321196,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::setWordWrapMode(QTextOption::WrapMode)
func (q *QPlainTextEdit) SetWordWrapMode(policy QTextOption_WrapMode)  {
	q.Drv(321000,321197,unsafe.Pointer(&policy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::showEvent(QShowEvent*)
func (q *QPlainTextEdit) ShowEvent(value *QShowEvent)  {
	q.Drv(321000,321198,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::tabChangesFocus()
func (q *QPlainTextEdit) TabChangesFocus() bool {
	var __rv bool
	q.Drv(321000,321199,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::tabStopWidth()
func (q *QPlainTextEdit) TabStopWidth() int {
	var __rv int
	q.Drv(321000,321200,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::textCursor()
func (q *QPlainTextEdit) TextCursor() *QTextCursor {
	var __rv uintptr
	q.Drv(321000,321201,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QPlainTextEdit::textInteractionFlags()
func (q *QPlainTextEdit) TextInteractionFlags() Qt_TextInteractionFlag {
	var __rv Qt_TextInteractionFlag
	q.Drv(321000,321202,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::timerEvent(QTimerEvent*)
func (q *QPlainTextEdit) TimerEvent(e *QTimerEvent)  {
	q.Drv(321000,321203,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::toPlainText()
func (q *QPlainTextEdit) ToPlainText() string {
	var __rv string
	q.Drv(321000,321204,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPlainTextEdit::undo()
func (q *QPlainTextEdit) Undo()  {
	q.Drv(321000,321205,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::wheelEvent(QWheelEvent*)
func (q *QPlainTextEdit) WheelEvent(e *QWheelEvent)  {
	q.Drv(321000,321206,Native(e),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QPlainTextEdit::wordWrapMode()
func (q *QPlainTextEdit) WordWrapMode() QTextOption_WrapMode {
	var __rv QTextOption_WrapMode
	q.Drv(321000,321207,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
