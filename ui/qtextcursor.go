// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextCursor_MoveMode - QTextCursor::MoveMode
type QTextCursor_MoveMode uint32
const (
	QTextCursor_MoveAnchor QTextCursor_MoveMode = 0
	QTextCursor_KeepAnchor QTextCursor_MoveMode = 1
)
//enum QTextCursor_MoveOperation - QTextCursor::MoveOperation
type QTextCursor_MoveOperation uint32
const (
	QTextCursor_NoMove QTextCursor_MoveOperation = 0
	QTextCursor_Start QTextCursor_MoveOperation = 1
	QTextCursor_Up QTextCursor_MoveOperation = 2
	QTextCursor_StartOfLine QTextCursor_MoveOperation = 3
	QTextCursor_StartOfBlock QTextCursor_MoveOperation = 4
	QTextCursor_StartOfWord QTextCursor_MoveOperation = 5
	QTextCursor_PreviousBlock QTextCursor_MoveOperation = 6
	QTextCursor_PreviousCharacter QTextCursor_MoveOperation = 7
	QTextCursor_PreviousWord QTextCursor_MoveOperation = 8
	QTextCursor_Left QTextCursor_MoveOperation = 9
	QTextCursor_WordLeft QTextCursor_MoveOperation = 10
	QTextCursor_End QTextCursor_MoveOperation = 11
	QTextCursor_Down QTextCursor_MoveOperation = 12
	QTextCursor_EndOfLine QTextCursor_MoveOperation = 13
	QTextCursor_EndOfWord QTextCursor_MoveOperation = 14
	QTextCursor_EndOfBlock QTextCursor_MoveOperation = 15
	QTextCursor_NextBlock QTextCursor_MoveOperation = 16
	QTextCursor_NextCharacter QTextCursor_MoveOperation = 17
	QTextCursor_NextWord QTextCursor_MoveOperation = 18
	QTextCursor_Right QTextCursor_MoveOperation = 19
	QTextCursor_WordRight QTextCursor_MoveOperation = 20
	QTextCursor_NextCell QTextCursor_MoveOperation = 21
	QTextCursor_PreviousCell QTextCursor_MoveOperation = 22
	QTextCursor_NextRow QTextCursor_MoveOperation = 23
	QTextCursor_PreviousRow QTextCursor_MoveOperation = 24
)
//enum QTextCursor_SelectionType - QTextCursor::SelectionType
type QTextCursor_SelectionType uint32
const (
	QTextCursor_WordUnderCursor QTextCursor_SelectionType = 0
	QTextCursor_LineUnderCursor QTextCursor_SelectionType = 1
	QTextCursor_BlockUnderCursor QTextCursor_SelectionType = 2
	QTextCursor_Document QTextCursor_SelectionType = 3
)
//struct QTextCursor : QTextCursor
type QTextCursor struct {
	BaseDrv
}
//QTextCursor::QTextCursor()	
func NewTextCursor() *QTextCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),145000,145102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCursor{}
	_p.SetDriver(__rv,145000,true)
	return _p
} 
//QTextCursor::QTextCursor(QTextBlock const&)	
func NewTextCursorWithBlock(block *QTextBlock) *QTextCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),145000,145103,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCursor{}
	_p.SetDriver(__rv,145000,true)
	return _p
} 
//QTextCursor::QTextCursor(QTextCursor const&)	
func NewTextCursorCopy(cursor *QTextCursor) *QTextCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),145000,145104,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCursor{}
	_p.SetDriver(__rv,145000,true)
	return _p
} 
//QTextCursor::QTextCursor(QTextDocument*)	
func NewTextCursorWithDocument(document *QTextDocument) *QTextCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),145000,145105,Native(document),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCursor{}
	_p.SetDriver(__rv,145000,true)
	return _p
} 
//QTextCursor::QTextCursor(QTextFrame*)	
func NewTextCursorWithFrame(frame *QTextFrame) *QTextCursor {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),145000,145106,Native(frame),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextCursor{}
	_p.SetDriver(__rv,145000,true)
	return _p
} 
//QTextCursor::anchor()
func (q *QTextCursor) Anchor() int {
	var __rv int
	q.Drv(145000,145107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::atBlockEnd()
func (q *QTextCursor) AtBlockEnd() bool {
	var __rv bool
	q.Drv(145000,145108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::atBlockStart()
func (q *QTextCursor) AtBlockStart() bool {
	var __rv bool
	q.Drv(145000,145109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::atEnd()
func (q *QTextCursor) AtEnd() bool {
	var __rv bool
	q.Drv(145000,145110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::atStart()
func (q *QTextCursor) AtStart() bool {
	var __rv bool
	q.Drv(145000,145111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::beginEditBlock()
func (q *QTextCursor) BeginEditBlock()  {
	q.Drv(145000,145112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::block()
func (q *QTextCursor) Block() *QTextBlock {
	var __rv uintptr
	q.Drv(145000,145113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextCursor::blockCharFormat()
func (q *QTextCursor) BlockCharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(145000,145114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextCursor::blockFormat()
func (q *QTextCursor) BlockFormat() *QTextBlockFormat {
	var __rv uintptr
	q.Drv(145000,145115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockFormat{}
	_rp.SetDriver(__rv,139000,true)
	return _rp
}	
//QTextCursor::blockNumber()
func (q *QTextCursor) BlockNumber() int {
	var __rv int
	q.Drv(145000,145116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::charFormat()
func (q *QTextCursor) CharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(145000,145117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextCursor::clearSelection()
func (q *QTextCursor) ClearSelection()  {
	q.Drv(145000,145118,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::columnNumber()
func (q *QTextCursor) ColumnNumber() int {
	var __rv int
	q.Drv(145000,145119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::createList(QTextListFormat const&)
func (q *QTextCursor) CreateList(format *QTextListFormat) *QTextList {
	var __rv uintptr
	q.Drv(145000,145120,Native(format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextCursor::createList(QTextListFormat::Style)
func (q *QTextCursor) CreateListWithStyle(style QTextListFormat_Style) *QTextList {
	var __rv uintptr
	q.Drv(145000,145121,unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextCursor::currentFrame()
func (q *QTextCursor) CurrentFrame() *QTextFrame {
	var __rv uintptr
	q.Drv(145000,145122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextCursor::currentList()
func (q *QTextCursor) CurrentList() *QTextList {
	var __rv uintptr
	q.Drv(145000,145123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextCursor::currentTable()
func (q *QTextCursor) CurrentTable() *QTextTable {
	var __rv uintptr
	q.Drv(145000,145124,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTable{}
	_rp.SetDriver(__rv,377000,false)
	return _rp
}	
//QTextCursor::deleteChar()
func (q *QTextCursor) DeleteChar()  {
	q.Drv(145000,145125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::deletePreviousChar()
func (q *QTextCursor) DeletePreviousChar()  {
	q.Drv(145000,145126,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::document()
func (q *QTextCursor) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(145000,145127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextCursor::endEditBlock()
func (q *QTextCursor) EndEditBlock()  {
	q.Drv(145000,145128,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::hasComplexSelection()
func (q *QTextCursor) HasComplexSelection() bool {
	var __rv bool
	q.Drv(145000,145129,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::hasSelection()
func (q *QTextCursor) HasSelection() bool {
	var __rv bool
	q.Drv(145000,145130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::insertBlock()
func (q *QTextCursor) InsertBlock()  {
	q.Drv(145000,145131,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertBlock(QTextBlockFormat const&)
func (q *QTextCursor) InsertBlockWithFormat(format *QTextBlockFormat)  {
	q.Drv(145000,145132,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertBlock(QTextBlockFormat const&,QTextCharFormat const&)
func (q *QTextCursor) InsertBlockWithFormatCharformat(format *QTextBlockFormat,charFormat *QTextCharFormat)  {
	q.Drv(145000,145133,Native(format),Native(charFormat),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertFragment(QTextDocumentFragment const&)
func (q *QTextCursor) InsertFragment(fragment *QTextDocumentFragment)  {
	q.Drv(145000,145134,Native(fragment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertFrame(QTextFrameFormat const&)
func (q *QTextCursor) InsertFrame(format *QTextFrameFormat) *QTextFrame {
	var __rv uintptr
	q.Drv(145000,145135,Native(format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextCursor::insertHtml(QString const&)
func (q *QTextCursor) InsertHtml(html string)  {
	q.Drv(145000,145136,unsafe.Pointer(&html),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertImage(QImage const&)
func (q *QTextCursor) InsertImage(image *QImage)  {
	q.Drv(145000,145137,Native(image),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertImage(QString const&)
func (q *QTextCursor) InsertImageWithName(name string)  {
	q.Drv(145000,145138,unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertImage(QTextImageFormat const&)
func (q *QTextCursor) InsertImageWithFormat(format *QTextImageFormat)  {
	q.Drv(145000,145139,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertImage(QImage const&,QString const&)
func (q *QTextCursor) InsertImageWithImageName(image *QImage,name string)  {
	q.Drv(145000,145140,Native(image),unsafe.Pointer(&name),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertImage(QTextImageFormat const&,QTextFrameFormat::Position)
func (q *QTextCursor) InsertImageWithFormatAlignment(format *QTextImageFormat,alignment QTextFrameFormat_Position)  {
	q.Drv(145000,145141,Native(format),unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertList(QTextListFormat const&)
func (q *QTextCursor) InsertList(format *QTextListFormat) *QTextList {
	var __rv uintptr
	q.Drv(145000,145142,Native(format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextCursor::insertList(QTextListFormat::Style)
func (q *QTextCursor) InsertListWithStyle(style QTextListFormat_Style) *QTextList {
	var __rv uintptr
	q.Drv(145000,145143,unsafe.Pointer(&style),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextList{}
	_rp.SetDriver(__rv,375000,false)
	return _rp
}	
//QTextCursor::insertTable(int,int)
func (q *QTextCursor) InsertTableWithRowsCols(rows int,cols int) *QTextTable {
	var __rv uintptr
	q.Drv(145000,145144,unsafe.Pointer(&rows),unsafe.Pointer(&cols),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTable{}
	_rp.SetDriver(__rv,377000,false)
	return _rp
}	
//QTextCursor::insertTable(int,int,QTextTableFormat const&)
func (q *QTextCursor) InsertTableWithRowsColsFormat(rows int,cols int,format *QTextTableFormat) *QTextTable {
	var __rv uintptr
	q.Drv(145000,145145,unsafe.Pointer(&rows),unsafe.Pointer(&cols),Native(format),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextTable{}
	_rp.SetDriver(__rv,377000,false)
	return _rp
}	
//QTextCursor::insertText(QString const&)
func (q *QTextCursor) InsertText(text string)  {
	q.Drv(145000,145146,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::insertText(QString const&,QTextCharFormat const&)
func (q *QTextCursor) InsertTextWithTextFormat(text string,format *QTextCharFormat)  {
	q.Drv(145000,145147,unsafe.Pointer(&text),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::isCopyOf(QTextCursor const&)
func (q *QTextCursor) IsCopyOf(other *QTextCursor) bool {
	var __rv bool
	q.Drv(145000,145148,Native(other),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::isNull()
func (q *QTextCursor) IsNull() bool {
	var __rv bool
	q.Drv(145000,145149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::joinPreviousEditBlock()
func (q *QTextCursor) JoinPreviousEditBlock()  {
	q.Drv(145000,145150,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::keepPositionOnInsert()
func (q *QTextCursor) KeepPositionOnInsert() bool {
	var __rv bool
	q.Drv(145000,145151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::mergeBlockCharFormat(QTextCharFormat const&)
func (q *QTextCursor) MergeBlockCharFormat(modifier *QTextCharFormat)  {
	q.Drv(145000,145152,Native(modifier),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::mergeBlockFormat(QTextBlockFormat const&)
func (q *QTextCursor) MergeBlockFormat(modifier *QTextBlockFormat)  {
	q.Drv(145000,145153,Native(modifier),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::mergeCharFormat(QTextCharFormat const&)
func (q *QTextCursor) MergeCharFormat(modifier *QTextCharFormat)  {
	q.Drv(145000,145154,Native(modifier),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::movePosition(QTextCursor::MoveOperation)
func (q *QTextCursor) MovePosition(op QTextCursor_MoveOperation) bool {
	var __rv bool
	q.Drv(145000,145155,unsafe.Pointer(&op),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::movePosition(QTextCursor::MoveOperation,QTextCursor::MoveMode,int)
func (q *QTextCursor) MovePositionWithOpMovemodeN(op QTextCursor_MoveOperation,value2 QTextCursor_MoveMode,n int) bool {
	var __rv bool
	q.Drv(145000,145156,unsafe.Pointer(&op),unsafe.Pointer(&value2),unsafe.Pointer(&n),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::position()
func (q *QTextCursor) Position() int {
	var __rv int
	q.Drv(145000,145157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::positionInBlock()
func (q *QTextCursor) PositionInBlock() int {
	var __rv int
	q.Drv(145000,145158,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::removeSelectedText()
func (q *QTextCursor) RemoveSelectedText()  {
	q.Drv(145000,145159,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::select(QTextCursor::SelectionType)
func (q *QTextCursor) Select(selection QTextCursor_SelectionType)  {
	q.Drv(145000,145160,unsafe.Pointer(&selection),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::selectedTableCells(int*,int*,int*,int*)
func (q *QTextCursor) SelectedTableCells(firstRow *int,numRows *int,firstColumn *int,numColumns *int)  {
	q.Drv(145000,145161,unsafe.Pointer(&firstRow),unsafe.Pointer(&numRows),unsafe.Pointer(&firstColumn),unsafe.Pointer(&numColumns),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::selectedText()
func (q *QTextCursor) SelectedText() string {
	var __rv string
	q.Drv(145000,145162,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::selection()
func (q *QTextCursor) Selection() *QTextDocumentFragment {
	var __rv uintptr
	q.Drv(145000,145163,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocumentFragment{}
	_rp.SetDriver(__rv,147000,true)
	return _rp
}	
//QTextCursor::selectionEnd()
func (q *QTextCursor) SelectionEnd() int {
	var __rv int
	q.Drv(145000,145164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::selectionStart()
func (q *QTextCursor) SelectionStart() int {
	var __rv int
	q.Drv(145000,145165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::setBlockCharFormat(QTextCharFormat const&)
func (q *QTextCursor) SetBlockCharFormat(format *QTextCharFormat)  {
	q.Drv(145000,145166,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setBlockFormat(QTextBlockFormat const&)
func (q *QTextCursor) SetBlockFormat(format *QTextBlockFormat)  {
	q.Drv(145000,145167,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setCharFormat(QTextCharFormat const&)
func (q *QTextCursor) SetCharFormat(format *QTextCharFormat)  {
	q.Drv(145000,145168,Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setKeepPositionOnInsert(bool)
func (q *QTextCursor) SetKeepPositionOnInsert(b bool)  {
	q.Drv(145000,145169,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setPosition(int)
func (q *QTextCursor) SetPosition(pos int)  {
	q.Drv(145000,145170,unsafe.Pointer(&pos),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setPosition(int,QTextCursor::MoveMode)
func (q *QTextCursor) SetPositionWithPosMode(pos int,mode QTextCursor_MoveMode)  {
	q.Drv(145000,145171,unsafe.Pointer(&pos),unsafe.Pointer(&mode),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setVerticalMovementX(int)
func (q *QTextCursor) SetVerticalMovementX(x int)  {
	q.Drv(145000,145172,unsafe.Pointer(&x),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::setVisualNavigation(bool)
func (q *QTextCursor) SetVisualNavigation(b bool)  {
	q.Drv(145000,145173,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextCursor::verticalMovementX()
func (q *QTextCursor) VerticalMovementX() int {
	var __rv int
	q.Drv(145000,145174,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextCursor::visualNavigation()
func (q *QTextCursor) VisualNavigation() bool {
	var __rv bool
	q.Drv(145000,145175,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
