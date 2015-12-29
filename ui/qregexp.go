// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QRegExp_PatternSyntax - QRegExp::PatternSyntax
type QRegExp_PatternSyntax uint32
const (
	QRegExp_RegExp QRegExp_PatternSyntax = 0
	QRegExp_Wildcard QRegExp_PatternSyntax = 1
	QRegExp_FixedString QRegExp_PatternSyntax = 2
	QRegExp_RegExp2 QRegExp_PatternSyntax = 3
	QRegExp_WildcardUnix QRegExp_PatternSyntax = 4
	QRegExp_W3CXmlSchema11 QRegExp_PatternSyntax = 5
)
//enum QRegExp_CaretMode - QRegExp::CaretMode
type QRegExp_CaretMode uint32
const (
	QRegExp_CaretAtZero QRegExp_CaretMode = 0
	QRegExp_CaretAtOffset QRegExp_CaretMode = 1
	QRegExp_CaretWontMatch QRegExp_CaretMode = 2
)
//struct QRegExp : QRegExp
type QRegExp struct {
	BaseDrv
}
//QRegExp::QRegExp()	
func NewRegExp() *QRegExp {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),112000,112102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExp{}
	_p.SetDriver(__rv,112000,true)
	return _p
} 
//QRegExp::QRegExp(QRegExp const&)	
func NewRegExpCopy(rx *QRegExp) *QRegExp {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),112000,112103,Native(rx),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExp{}
	_p.SetDriver(__rv,112000,true)
	return _p
} 
//QRegExp::QRegExp(QString const&,Qt::CaseSensitivity,QRegExp::PatternSyntax)	
func NewRegExpWithPatternCsSyntax(pattern string,cs Qt_CaseSensitivity,syntax QRegExp_PatternSyntax) *QRegExp {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),112000,112104,unsafe.Pointer(&pattern),unsafe.Pointer(&cs),unsafe.Pointer(&syntax),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QRegExp{}
	_p.SetDriver(__rv,112000,true)
	return _p
} 
//QRegExp::cap()
func (q *QRegExp) Cap() string {
	var __rv string
	q.Drv(112000,112105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::cap(int)
func (q *QRegExp) CapWithNth(nth int) string {
	var __rv string
	q.Drv(112000,112106,unsafe.Pointer(&nth),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::captureCount()
func (q *QRegExp) CaptureCount() int {
	var __rv int
	q.Drv(112000,112107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::capturedTexts()
func (q *QRegExp) CapturedTexts() []string {
	var __rv []string
	q.Drv(112000,112108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::caseSensitivity()
func (q *QRegExp) CaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(112000,112109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::errorString()
func (q *QRegExp) ErrorString() string {
	var __rv string
	q.Drv(112000,112110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::escape(QString const&)	
func QRegExpEscape(str string) string {
	var __rv string
	DirectQtDrv(nil,112000,112111,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::escape(QString const&)
func (q *QRegExp) Escape(str string) string {
	var __rv string
	q.Drv(112000,112111,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::exactMatch(QString const&)
func (q *QRegExp) ExactMatch(str string) bool {
	var __rv bool
	q.Drv(112000,112112,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::indexIn(QString const&)
func (q *QRegExp) IndexIn(str string) int {
	var __rv int
	q.Drv(112000,112113,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::indexIn(QString const&,int,QRegExp::CaretMode)
func (q *QRegExp) IndexInWithTextOffsetCaretmode(str string,offset int,caretMode QRegExp_CaretMode) int {
	var __rv int
	q.Drv(112000,112114,unsafe.Pointer(&str),unsafe.Pointer(&offset),unsafe.Pointer(&caretMode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::isEmpty()
func (q *QRegExp) IsEmpty() bool {
	var __rv bool
	q.Drv(112000,112115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::isMinimal()
func (q *QRegExp) IsMinimal() bool {
	var __rv bool
	q.Drv(112000,112116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::isValid()
func (q *QRegExp) IsValid() bool {
	var __rv bool
	q.Drv(112000,112117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::lastIndexIn(QString const&)
func (q *QRegExp) LastIndexIn(str string) int {
	var __rv int
	q.Drv(112000,112118,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::lastIndexIn(QString const&,int,QRegExp::CaretMode)
func (q *QRegExp) LastIndexInWithTextOffsetCaretmode(str string,offset int,caretMode QRegExp_CaretMode) int {
	var __rv int
	q.Drv(112000,112119,unsafe.Pointer(&str),unsafe.Pointer(&offset),unsafe.Pointer(&caretMode),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::matchedLength()
func (q *QRegExp) MatchedLength() int {
	var __rv int
	q.Drv(112000,112120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::pattern()
func (q *QRegExp) Pattern() string {
	var __rv string
	q.Drv(112000,112121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::patternSyntax()
func (q *QRegExp) PatternSyntax() QRegExp_PatternSyntax {
	var __rv QRegExp_PatternSyntax
	q.Drv(112000,112122,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::pos()
func (q *QRegExp) Pos() int {
	var __rv int
	q.Drv(112000,112123,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::pos(int)
func (q *QRegExp) PosWithNth(nth int) int {
	var __rv int
	q.Drv(112000,112124,unsafe.Pointer(&nth),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QRegExp::setCaseSensitivity(Qt::CaseSensitivity)
func (q *QRegExp) SetCaseSensitivity(cs Qt_CaseSensitivity)  {
	q.Drv(112000,112125,unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegExp::setMinimal(bool)
func (q *QRegExp) SetMinimal(minimal bool)  {
	q.Drv(112000,112126,unsafe.Pointer(&minimal),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegExp::setPattern(QString const&)
func (q *QRegExp) SetPattern(pattern string)  {
	q.Drv(112000,112127,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QRegExp::setPatternSyntax(QRegExp::PatternSyntax)
func (q *QRegExp) SetPatternSyntax(syntax QRegExp_PatternSyntax)  {
	q.Drv(112000,112128,unsafe.Pointer(&syntax),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
