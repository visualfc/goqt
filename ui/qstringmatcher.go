// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QStringMatcher : QStringMatcher
type QStringMatcher struct {
	BaseDrv
}
//QStringMatcher::QStringMatcher()	
func NewStringMatcher() *QStringMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),128000,128102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStringMatcher{}
	_p.SetDriver(__rv,128000,true)
	return _p
} 
//QStringMatcher::QStringMatcher(QStringMatcher const&)	
func NewStringMatcherCopy(other *QStringMatcher) *QStringMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),128000,128103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStringMatcher{}
	_p.SetDriver(__rv,128000,true)
	return _p
} 
//QStringMatcher::QStringMatcher(QString const&,Qt::CaseSensitivity)	
func NewStringMatcherWithPatternCs(pattern string,cs Qt_CaseSensitivity) *QStringMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),128000,128104,unsafe.Pointer(&pattern),unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStringMatcher{}
	_p.SetDriver(__rv,128000,true)
	return _p
} 
//QStringMatcher::QStringMatcher(QChar const*,int,Qt::CaseSensitivity)	
func NewStringMatcherWithUcLenCs(uc *rune,len int,cs Qt_CaseSensitivity) *QStringMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),128000,128105,unsafe.Pointer(&uc),unsafe.Pointer(&len),unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QStringMatcher{}
	_p.SetDriver(__rv,128000,true)
	return _p
} 
//QStringMatcher::caseSensitivity()
func (q *QStringMatcher) CaseSensitivity() Qt_CaseSensitivity {
	var __rv Qt_CaseSensitivity
	q.Drv(128000,128106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringMatcher::indexIn(QString const&)
func (q *QStringMatcher) IndexIn(str string) int {
	var __rv int
	q.Drv(128000,128107,unsafe.Pointer(&str),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringMatcher::indexIn(QString const&,int)
func (q *QStringMatcher) IndexInWithTextIfrom(str string,from int) int {
	var __rv int
	q.Drv(128000,128108,unsafe.Pointer(&str),unsafe.Pointer(&from),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringMatcher::indexIn(QChar const*,int,int)
func (q *QStringMatcher) IndexInWithStrLengthIfrom(str *rune,length int,from int) int {
	var __rv int
	q.Drv(128000,128109,unsafe.Pointer(&str),unsafe.Pointer(&length),unsafe.Pointer(&from),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringMatcher::pattern()
func (q *QStringMatcher) Pattern() string {
	var __rv string
	q.Drv(128000,128110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QStringMatcher::setCaseSensitivity(Qt::CaseSensitivity)
func (q *QStringMatcher) SetCaseSensitivity(cs Qt_CaseSensitivity)  {
	q.Drv(128000,128111,unsafe.Pointer(&cs),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QStringMatcher::setPattern(QString const&)
func (q *QStringMatcher) SetPattern(pattern string)  {
	q.Drv(128000,128112,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
