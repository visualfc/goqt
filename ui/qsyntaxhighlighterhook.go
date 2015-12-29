// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QSyntaxHighlighterHook : QSyntaxHighlighterHook
type QSyntaxHighlighterHook struct {
	QSyntaxHighlighter
}
func (q *QSyntaxHighlighterHook) OnHighlightBlock(fn func(string)) uintptr {
	var __rv uintptr
	q.Drv(361000,361102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QSyntaxHighlighterHook::QSyntaxHighlighterHook(QObject*)	
func NewSyntaxHighlighterHook(parent QObjectInterface) *QSyntaxHighlighterHook {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),361000,361103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSyntaxHighlighterHook{}
	_p.SetDriver(__rv,361000,false)
	return _p
} 
//QSyntaxHighlighterHook::QSyntaxHighlighterHook(QTextDocument*)	
func NewSyntaxHighlighterHookWithDoc(doc *QTextDocument) *QSyntaxHighlighterHook {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),361000,361104,Native(doc),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSyntaxHighlighterHook{}
	_p.SetDriver(__rv,361000,false)
	return _p
} 
//QSyntaxHighlighterHook::QSyntaxHighlighterHook(QTextEdit*)	
func NewSyntaxHighlighterHookWithEdit(edit *QTextEdit) *QSyntaxHighlighterHook {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),361000,361105,Native(edit),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QSyntaxHighlighterHook{}
	_p.SetDriver(__rv,361000,false)
	return _p
} 
//QSyntaxHighlighterHook::highlightBlock(QString const&)
func (q *QSyntaxHighlighterHook) HighlightBlock(text string)  {
	q.Drv(361000,361106,unsafe.Pointer(&text),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
