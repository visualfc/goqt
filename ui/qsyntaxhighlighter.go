// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSyntaxHighlighter : QSyntaxHighlighter
type QSyntaxHighlighter struct {
	QObject
}
//QSyntaxHighlighter::currentBlock()
func (q *QSyntaxHighlighter) CurrentBlock() *QTextBlock {
	var __rv uintptr
	q.Drv(360000,360102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlock{}
	_rp.SetDriver(__rv,137000,true)
	return _rp
}	
//QSyntaxHighlighter::currentBlockState()
func (q *QSyntaxHighlighter) CurrentBlockState() int {
	var __rv int
	q.Drv(360000,360103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSyntaxHighlighter::currentBlockUserData()
func (q *QSyntaxHighlighter) CurrentBlockUserData() *QTextBlockUserData {
	var __rv uintptr
	q.Drv(360000,360104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextBlockUserData{}
	_rp.SetDriver(__rv,140000,true)
	return _rp
}	
//QSyntaxHighlighter::document()
func (q *QSyntaxHighlighter) Document() *QTextDocument {
	var __rv uintptr
	q.Drv(360000,360105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextDocument{}
	_rp.SetDriver(__rv,372000,false)
	return _rp
}	
//QSyntaxHighlighter::format(int)
func (q *QSyntaxHighlighter) Format(pos int) *QTextCharFormat {
	var __rv uintptr
	q.Drv(360000,360106,unsafe.Pointer(&pos),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QSyntaxHighlighter::highlightBlock(QString const&)
func (q *QSyntaxHighlighter) HighlightBlock(text string)  {
	q.Drv(360000,360107,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::previousBlockState()
func (q *QSyntaxHighlighter) PreviousBlockState() int {
	var __rv int
	q.Drv(360000,360108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QSyntaxHighlighter::rehighlight()
func (q *QSyntaxHighlighter) Rehighlight()  {
	q.Drv(360000,360109,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::rehighlightBlock(QTextBlock const&)
func (q *QSyntaxHighlighter) RehighlightBlock(block *QTextBlock)  {
	q.Drv(360000,360110,Native(block),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setCurrentBlockState(int)
func (q *QSyntaxHighlighter) SetCurrentBlockState(newState int)  {
	q.Drv(360000,360111,unsafe.Pointer(&newState),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setCurrentBlockUserData(QTextBlockUserData*)
func (q *QSyntaxHighlighter) SetCurrentBlockUserData(data *QTextBlockUserData)  {
	q.Drv(360000,360112,Native(data),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setDocument(QTextDocument*)
func (q *QSyntaxHighlighter) SetDocument(doc *QTextDocument)  {
	q.Drv(360000,360113,Native(doc),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setFormat(int,int,QColor const&)
func (q *QSyntaxHighlighter) SetFormat(start int,count int,color *QColor)  {
	q.Drv(360000,360114,unsafe.Pointer(&start),unsafe.Pointer(&count),Native(color),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setFormat(int,int,QFont const&)
func (q *QSyntaxHighlighter) SetFormatWithStartCountFont(start int,count int,font *QFont)  {
	q.Drv(360000,360115,unsafe.Pointer(&start),unsafe.Pointer(&count),Native(font),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QSyntaxHighlighter::setFormat(int,int,QTextCharFormat const&)
func (q *QSyntaxHighlighter) SetFormatWithStartCountFormat(start int,count int,format *QTextCharFormat)  {
	q.Drv(360000,360116,unsafe.Pointer(&start),unsafe.Pointer(&count),Native(format),nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
