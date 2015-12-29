// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QByteArrayMatcher : QByteArrayMatcher
type QByteArrayMatcher struct {
	BaseDrv
}
//QByteArrayMatcher::QByteArrayMatcher()	
func NewByteArrayMatcher() *QByteArrayMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),10000,10102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QByteArrayMatcher{}
	_p.SetDriver(__rv,10000,true)
	return _p
} 
//QByteArrayMatcher::QByteArrayMatcher(QByteArray const&)	
func NewByteArrayMatcherWithPattern(pattern []byte) *QByteArrayMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),10000,10103,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QByteArrayMatcher{}
	_p.SetDriver(__rv,10000,true)
	return _p
} 
//QByteArrayMatcher::QByteArrayMatcher(QByteArrayMatcher const&)	
func NewByteArrayMatcherCopy(other *QByteArrayMatcher) *QByteArrayMatcher {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),10000,10104,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QByteArrayMatcher{}
	_p.SetDriver(__rv,10000,true)
	return _p
} 
//QByteArrayMatcher::indexIn(QByteArray const&)
func (q *QByteArrayMatcher) IndexIn(ba []byte) int {
	var __rv int
	q.Drv(10000,10105,unsafe.Pointer(&ba),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QByteArrayMatcher::indexIn(QByteArray const&,int)
func (q *QByteArrayMatcher) IndexInWithBaIfrom(ba []byte,from int) int {
	var __rv int
	q.Drv(10000,10106,unsafe.Pointer(&ba),unsafe.Pointer(&from),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QByteArrayMatcher::pattern()
func (q *QByteArrayMatcher) Pattern() []byte {
	var __rv []byte
	q.Drv(10000,10107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QByteArrayMatcher::setPattern(QByteArray const&)
func (q *QByteArrayMatcher) SetPattern(pattern []byte)  {
	q.Drv(10000,10108,unsafe.Pointer(&pattern),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
