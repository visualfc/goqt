// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTranslator : QTranslator
type QTranslator struct {
	QObject
}
//QTranslator::QTranslator()	
func NewTranslator() *QTranslator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),384000,384102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTranslator{}
	_p.SetDriver(__rv,384000,false)
	return _p
} 
//QTranslator::QTranslator(QObject*)	
func NewTranslatorWithParent(parent QObjectInterface) *QTranslator {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),384000,384103,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTranslator{}
	_p.SetDriver(__rv,384000,false)
	return _p
} 
//QTranslator::isEmpty()
func (q *QTranslator) IsEmpty() bool {
	var __rv bool
	q.Drv(384000,384104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTranslator::load(QString const&)
func (q *QTranslator) Load(filename string) bool {
	var __rv bool
	q.Drv(384000,384105,unsafe.Pointer(&filename),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTranslator::load(unsigned char const*,int)
func (q *QTranslator) LoadWithDataLen(data *byte,len int) bool {
	var __rv bool
	q.Drv(384000,384106,unsafe.Pointer(&data),unsafe.Pointer(&len),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTranslator::load(QString const&,QString const&,QString const&,QString const&)
func (q *QTranslator) LoadWithFilenameDirectorySearch_delimitersSuffix(filename string,directory string,search_delimiters string,suffix string) bool {
	var __rv bool
	q.Drv(384000,384107,unsafe.Pointer(&filename),unsafe.Pointer(&directory),unsafe.Pointer(&search_delimiters),unsafe.Pointer(&suffix),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTranslator::translate(char const*,char const*,char const*)
func (q *QTranslator) TranslateWithContextSourcetextDisambiguation(context string,sourceText string,disambiguation string) string {
	var __rv string
	q.Drv(384000,384108,unsafe.Pointer(&context),unsafe.Pointer(&sourceText),unsafe.Pointer(&disambiguation),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTranslator::translate(char const*,char const*,char const*,int)
func (q *QTranslator) TranslateWithContextSourcetextDisambiguationN(context string,sourceText string,disambiguation string,n int) string {
	var __rv string
	q.Drv(384000,384109,unsafe.Pointer(&context),unsafe.Pointer(&sourceText),unsafe.Pointer(&disambiguation),unsafe.Pointer(&n),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
