// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QTextFragment : QTextFragment
type QTextFragment struct {
	BaseDrv
}
//QTextFragment::QTextFragment()	
func NewTextFragment() *QTextFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),152000,152102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFragment{}
	_p.SetDriver(__rv,152000,true)
	return _p
} 
//QTextFragment::QTextFragment(QTextFragment const&)	
func NewTextFragmentCopy(o *QTextFragment) *QTextFragment {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),152000,152103,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextFragment{}
	_p.SetDriver(__rv,152000,true)
	return _p
} 
//QTextFragment::charFormat()
func (q *QTextFragment) CharFormat() *QTextCharFormat {
	var __rv uintptr
	q.Drv(152000,152104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QTextCharFormat{}
	_rp.SetDriver(__rv,142000,true)
	return _rp
}	
//QTextFragment::charFormatIndex()
func (q *QTextFragment) CharFormatIndex() int {
	var __rv int
	q.Drv(152000,152105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFragment::contains(int)
func (q *QTextFragment) Contains(position int) bool {
	var __rv bool
	q.Drv(152000,152106,unsafe.Pointer(&position),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFragment::isValid()
func (q *QTextFragment) IsValid() bool {
	var __rv bool
	q.Drv(152000,152107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFragment::length()
func (q *QTextFragment) Length() int {
	var __rv int
	q.Drv(152000,152108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFragment::position()
func (q *QTextFragment) Position() int {
	var __rv int
	q.Drv(152000,152109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextFragment::text()
func (q *QTextFragment) Text() string {
	var __rv string
	q.Drv(152000,152110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
