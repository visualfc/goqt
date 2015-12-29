// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextDocument_ResourceType - QTextDocument::ResourceType
type QTextDocument_ResourceType uint32
const (
	QTextDocument_HtmlResource QTextDocument_ResourceType = 1
	QTextDocument_ImageResource QTextDocument_ResourceType = 2
	QTextDocument_StyleSheetResource QTextDocument_ResourceType = 3
	QTextDocument_UserResource QTextDocument_ResourceType = 100
)
//enum QTextDocument_Stacks - QTextDocument::Stacks
type QTextDocument_Stacks uint32
const (
	QTextDocument_UndoStack QTextDocument_Stacks = 0x01
	QTextDocument_RedoStack QTextDocument_Stacks = 0x02
	QTextDocument_UndoAndRedoStacks QTextDocument_Stacks = QTextDocument_UndoStack|QTextDocument_RedoStack
)
//enum QTextDocument_FindFlag - QTextDocument::FindFlag
type QTextDocument_FindFlag uint32
const (
	QTextDocument_FindBackward QTextDocument_FindFlag = 0x00001
	QTextDocument_FindCaseSensitively QTextDocument_FindFlag = 0x00002
	QTextDocument_FindWholeWords QTextDocument_FindFlag = 0x00004
)
//enum QTextDocument_MetaInformation - QTextDocument::MetaInformation
type QTextDocument_MetaInformation uint32
const (
	QTextDocument_DocumentTitle QTextDocument_MetaInformation = 0
	QTextDocument_DocumentUrl QTextDocument_MetaInformation = 1
)
//struct QTextDocument : QTextDocument
type QTextDocument struct {
	QObject
}
func (q *QTextDocument) OnUndoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(372000,372102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnBlockCountChanged(fn func(int)) uintptr {
	var __rv uintptr
	q.Drv(372000,372103,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnModificationChanged(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(372000,372104,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnCursorPositionChanged(fn func(*QTextCursor)) uintptr {
	var __rv uintptr
	q.Drv(372000,372105,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnContentsChange(fn func(int,int,int)) uintptr {
	var __rv uintptr
	q.Drv(372000,372106,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnDocumentLayoutChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(372000,372107,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnUndoCommandAdded(fn func()) uintptr {
	var __rv uintptr
	q.Drv(372000,372108,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnContentsChanged(fn func()) uintptr {
	var __rv uintptr
	q.Drv(372000,372109,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
func (q *QTextDocument) OnRedoAvailable(fn func(bool)) uintptr {
	var __rv uintptr
	q.Drv(372000,372110,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QTextDocument::QTextDocument()	
func NewTextDocument() *QTextDocument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),372000,372111,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocument{}
	_p.SetDriver(__rv,372000,false)
	return _p
} 
//QTextDocument::QTextDocument(QObject*)	
func NewTextDocumentWithParent(parent QObjectInterface) *QTextDocument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),372000,372112,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocument{}
	_p.SetDriver(__rv,372000,false)
	return _p
} 
//QTextDocument::QTextDocument(QString const&,QObject*)	
func NewTextDocumentWithTextParent(text string,parent QObjectInterface) *QTextDocument {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),372000,372113,unsafe.Pointer(&text),Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextDocument{}
	_p.SetDriver(__rv,372000,false)
	return _p
} 
//QTextDocument::addResource(int,QUrl const&,QVariant const&)
func (q *QTextDocument) AddResource(_type int,name *QUrl,resource *QVariant)  {
	q.Drv(372000,372114,unsafe.Pointer(&_type),Native(name),Native(resource),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::adjustSize()
func (q *QTextDocument) AdjustSize()  {
	q.Drv(372000,372115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::allFormats()
func (q *QTextDocument) AllFormats() []*QTextFormat {
	var __rv []*QTextFormat
	q.Drv(372000,372116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::appendUndoItem(QAbstractUndoItem*)
func (q *QTextDocument) AppendUndoItem(value *QAbstractUndoItem)  {
	q.Drv(372000,372117,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::availableRedoSteps()
func (q *QTextDocument) AvailableRedoSteps() int {
	var __rv int
	q.Drv(372000,372118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::availableUndoSteps()
func (q *QTextDocument) AvailableUndoSteps() int {
	var __rv int
	q.Drv(372000,372119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::begin()
func (q *QTextDocument) Begin() *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::blockCount()
func (q *QTextDocument) BlockCount() int {
	var __rv int
	q.Drv(372000,372121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::characterAt(int)
func (q *QTextDocument) CharacterAt(pos int) rune {
	var __rv rune
	q.Drv(372000,372122,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::characterCount()
func (q *QTextDocument) CharacterCount() int {
	var __rv int
	q.Drv(372000,372123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::clear()
func (q *QTextDocument) Clear()  {
	q.Drv(372000,372124,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::clearUndoRedoStacks()
func (q *QTextDocument) ClearUndoRedoStacks()  {
	q.Drv(372000,372125,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::clearUndoRedoStacks(QTextDocument::Stacks)
func (q *QTextDocument) ClearUndoRedoStacksWithHistorytoclear(historyToClear QTextDocument_Stacks)  {
	q.Drv(372000,372126,unsafe.Pointer(&historyToClear),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::clone()
func (q *QTextDocument) Clone() *QTextDocument {
	var __rv uintptr
	q.Drv(372000,372127,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextDocument::clone(QObject*)
func (q *QTextDocument) CloneWithParent(parent QObjectInterface) *QTextDocument {
	var __rv uintptr
	q.Drv(372000,372128,Native(parent),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QTextDocument::createObject(QTextFormat const&)
func (q *QTextDocument) CreateObject(f *QTextFormat) *QTextObject {
	var __rv uintptr
	q.Drv(372000,372129,Native(f),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextObject{}
	_rp.SetDriver(__rv,376000,false)
	return _rp
}	
//QTextDocument::defaultFont()
func (q *QTextDocument) DefaultFont() *QFont {
	var __rv uintptr
	q.Drv(372000,372130,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QTextDocument::defaultStyleSheet()
func (q *QTextDocument) DefaultStyleSheet() string {
	var __rv string
	q.Drv(372000,372131,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::defaultTextOption()
func (q *QTextDocument) DefaultTextOption() *QTextOption {
	var __rv uintptr
	q.Drv(372000,372132,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextOption{}
	_rp.SetDriver(__rv,164000,true)
	return _rp
}	
//QTextDocument::documentLayout()
func (q *QTextDocument) DocumentLayout() *QAbstractTextDocumentLayout {
	var __rv uintptr
	q.Drv(372000,372133,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QAbstractTextDocumentLayout{}
	_rp.SetDriver(__rv,205000,false)
	return _rp
}	
//QTextDocument::documentMargin()
func (q *QTextDocument) DocumentMargin() float64 {
	var __rv float64
	q.Drv(372000,372134,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::drawContents(QPainter*)
func (q *QTextDocument) DrawContents(painter *QPainter)  {
	q.Drv(372000,372135,Native(painter),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::drawContents(QPainter*,QRectF const&)
func (q *QTextDocument) DrawContentsFWithPainterRect(painter *QPainter,rect *QRectF)  {
	q.Drv(372000,372136,Native(painter),Native(rect),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::end()
func (q *QTextDocument) End() *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372137,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::find(QRegExp const&)
func (q *QTextDocument) Find(expr *QRegExp) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372138,Native(expr),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::find(QString const&)
func (q *QTextDocument) FindWithSubstring(subString string) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372139,unsafe.Pointer(&subString),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::find(QRegExp const&,QTextCursor const&,QFlags<QTextDocument::FindFlag>)
func (q *QTextDocument) FindWithExprFromOptions(expr *QRegExp,from *QTextCursor,options QTextDocument_FindFlag) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372140,Native(expr),Native(from),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::find(QRegExp const&,int,QFlags<QTextDocument::FindFlag>)
func (q *QTextDocument) FindWithExprIfromOptions(expr *QRegExp,from int,options QTextDocument_FindFlag) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372141,Native(expr),unsafe.Pointer(&from),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::find(QString const&,QTextCursor const&,QFlags<QTextDocument::FindFlag>)
func (q *QTextDocument) FindWithSubstringFromOptions(subString string,from *QTextCursor,options QTextDocument_FindFlag) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372142,unsafe.Pointer(&subString),Native(from),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::find(QString const&,int,QFlags<QTextDocument::FindFlag>)
func (q *QTextDocument) FindWithSubstringIfromOptions(subString string,from int,options QTextDocument_FindFlag) *QTextCursor {
	var __rv uintptr
	q.Drv(372000,372143,unsafe.Pointer(&subString),unsafe.Pointer(&from),unsafe.Pointer(&options),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCursor{}
	_rp.SetDriver(__rv,145000,true)
	return _rp
}	
//QTextDocument::findBlock(int)
func (q *QTextDocument) FindBlock(pos int) *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372144,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::findBlockByLineNumber(int)
func (q *QTextDocument) FindBlockByLineNumber(blockNumber int) *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372145,unsafe.Pointer(&blockNumber),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::findBlockByNumber(int)
func (q *QTextDocument) FindBlockByNumber(blockNumber int) *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372146,unsafe.Pointer(&blockNumber),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::firstBlock()
func (q *QTextDocument) FirstBlock() *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372147,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::frameAt(int)
func (q *QTextDocument) FrameAt(pos int) *QTextFrame {
	var __rv uintptr
	q.Drv(372000,372148,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextDocument::idealWidth()
func (q *QTextDocument) IdealWidth() float64 {
	var __rv float64
	q.Drv(372000,372149,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::indentWidth()
func (q *QTextDocument) IndentWidth() float64 {
	var __rv float64
	q.Drv(372000,372150,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::isEmpty()
func (q *QTextDocument) IsEmpty() bool {
	var __rv bool
	q.Drv(372000,372151,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::isModified()
func (q *QTextDocument) IsModified() bool {
	var __rv bool
	q.Drv(372000,372152,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::isRedoAvailable()
func (q *QTextDocument) IsRedoAvailable() bool {
	var __rv bool
	q.Drv(372000,372153,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::isUndoAvailable()
func (q *QTextDocument) IsUndoAvailable() bool {
	var __rv bool
	q.Drv(372000,372154,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::isUndoRedoEnabled()
func (q *QTextDocument) IsUndoRedoEnabled() bool {
	var __rv bool
	q.Drv(372000,372155,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::lastBlock()
func (q *QTextDocument) LastBlock() *QTextBlock {
	var __rv uintptr
	q.Drv(372000,372156,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QTextDocument::lineCount()
func (q *QTextDocument) LineCount() int {
	var __rv int
	q.Drv(372000,372157,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::loadResource(int,QUrl const&)
func (q *QTextDocument) LoadResource(_type int,name *QUrl) *QVariant {
	var __rv uintptr
	q.Drv(372000,372158,unsafe.Pointer(&_type),Native(name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextDocument::markContentsDirty(int,int)
func (q *QTextDocument) MarkContentsDirty(from int,length int)  {
	q.Drv(372000,372159,unsafe.Pointer(&from),unsafe.Pointer(&length),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::maximumBlockCount()
func (q *QTextDocument) MaximumBlockCount() int {
	var __rv int
	q.Drv(372000,372160,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::metaInformation(QTextDocument::MetaInformation)
func (q *QTextDocument) MetaInformation(info QTextDocument_MetaInformation) string {
	var __rv string
	q.Drv(372000,372161,unsafe.Pointer(&info),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::object(int)
func (q *QTextDocument) Object(objectIndex int) *QTextObject {
	var __rv uintptr
	q.Drv(372000,372162,unsafe.Pointer(&objectIndex),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextObject{}
	_rp.SetDriver(__rv,376000,false)
	return _rp
}	
//QTextDocument::objectForFormat(QTextFormat const&)
func (q *QTextDocument) ObjectForFormat(value *QTextFormat) *QTextObject {
	var __rv uintptr
	q.Drv(372000,372163,Native(value),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextObject{}
	_rp.SetDriver(__rv,376000,false)
	return _rp
}	
//QTextDocument::pageCount()
func (q *QTextDocument) PageCount() int {
	var __rv int
	q.Drv(372000,372164,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::pageSize()
func (q *QTextDocument) PageSize() *QSizeF {
	var __rv uintptr
	q.Drv(372000,372165,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QTextDocument::print(QPrinter*)
func (q *QTextDocument) Print(printer *QPrinter)  {
	q.Drv(372000,372166,Native(printer),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::redo()
func (q *QTextDocument) Redo()  {
	q.Drv(372000,372167,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::redo(QTextCursor*)
func (q *QTextDocument) RedoWithCursor(cursor *QTextCursor)  {
	q.Drv(372000,372168,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::resource(int,QUrl const&)
func (q *QTextDocument) Resource(_type int,name *QUrl) *QVariant {
	var __rv uintptr
	q.Drv(372000,372169,unsafe.Pointer(&_type),Native(name),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QVariant{}
	_rp.SetDriver(__rv,182000,true)
	return _rp
}	
//QTextDocument::revision()
func (q *QTextDocument) Revision() int {
	var __rv int
	q.Drv(372000,372170,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::rootFrame()
func (q *QTextDocument) RootFrame() *QTextFrame {
	var __rv uintptr
	q.Drv(372000,372171,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextFrame{}
	_rp.SetDriver(__rv,374000,false)
	return _rp
}	
//QTextDocument::setDefaultFont(QFont const&)
func (q *QTextDocument) SetDefaultFont(font *QFont)  {
	q.Drv(372000,372172,Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setDefaultStyleSheet(QString const&)
func (q *QTextDocument) SetDefaultStyleSheet(sheet string)  {
	q.Drv(372000,372173,unsafe.Pointer(&sheet),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setDefaultTextOption(QTextOption const&)
func (q *QTextDocument) SetDefaultTextOption(option *QTextOption)  {
	q.Drv(372000,372174,Native(option),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setDocumentLayout(QAbstractTextDocumentLayout*)
func (q *QTextDocument) SetDocumentLayout(layout *QAbstractTextDocumentLayout)  {
	q.Drv(372000,372175,Native(layout),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setDocumentMargin(double)
func (q *QTextDocument) SetDocumentMargin(margin float64)  {
	q.Drv(372000,372176,unsafe.Pointer(&margin),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setHtml(QString const&)
func (q *QTextDocument) SetHtml(html string)  {
	q.Drv(372000,372177,unsafe.Pointer(&html),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setIndentWidth(double)
func (q *QTextDocument) SetIndentWidth(width float64)  {
	q.Drv(372000,372178,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setMaximumBlockCount(int)
func (q *QTextDocument) SetMaximumBlockCount(maximum int)  {
	q.Drv(372000,372179,unsafe.Pointer(&maximum),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setMetaInformation(QTextDocument::MetaInformation,QString const&)
func (q *QTextDocument) SetMetaInformation(info QTextDocument_MetaInformation,value2 string)  {
	q.Drv(372000,372180,unsafe.Pointer(&info),unsafe.Pointer(&value2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setModified(bool)
func (q *QTextDocument) SetModified(m bool)  {
	q.Drv(372000,372181,unsafe.Pointer(&m),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setPageSize(QSizeF const&)
func (q *QTextDocument) SetPageSize(size *QSizeF)  {
	q.Drv(372000,372182,Native(size),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setPlainText(QString const&)
func (q *QTextDocument) SetPlainText(text string)  {
	q.Drv(372000,372183,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setTextWidth(double)
func (q *QTextDocument) SetTextWidth(width float64)  {
	q.Drv(372000,372184,unsafe.Pointer(&width),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setUndoRedoEnabled(bool)
func (q *QTextDocument) SetUndoRedoEnabled(enable bool)  {
	q.Drv(372000,372185,unsafe.Pointer(&enable),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::setUseDesignMetrics(bool)
func (q *QTextDocument) SetUseDesignMetrics(b bool)  {
	q.Drv(372000,372186,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::size()
func (q *QTextDocument) Size() *QSizeF {
	var __rv uintptr
	q.Drv(372000,372187,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSizeF{}
	_rp.SetDriver(__rv,120000,true)
	return _rp
}	
//QTextDocument::textWidth()
func (q *QTextDocument) TextWidth() float64 {
	var __rv float64
	q.Drv(372000,372188,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::toHtml()
func (q *QTextDocument) ToHtml() string {
	var __rv string
	q.Drv(372000,372189,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::toHtml(QByteArray const&)
func (q *QTextDocument) ToHtmlWithEncoding(encoding []byte) string {
	var __rv string
	q.Drv(372000,372190,unsafe.Pointer(&encoding),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::toPlainText()
func (q *QTextDocument) ToPlainText() string {
	var __rv string
	q.Drv(372000,372191,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextDocument::undo()
func (q *QTextDocument) Undo()  {
	q.Drv(372000,372192,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::undo(QTextCursor*)
func (q *QTextDocument) UndoWithCursor(cursor *QTextCursor)  {
	q.Drv(372000,372193,Native(cursor),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextDocument::useDesignMetrics()
func (q *QTextDocument) UseDesignMetrics() bool {
	var __rv bool
	q.Drv(372000,372194,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
