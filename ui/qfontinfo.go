// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QFontInfo : QFontInfo
type QFontInfo struct {
	BaseDrv
}
//QFontInfo::QFontInfo(QFont const&)	
func NewFontInfo(value *QFont) *QFontInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),39000,39102,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontInfo{}
	_p.SetDriver(__rv,39000,true)
	return _p
} 
//QFontInfo::QFontInfo(QFontInfo const&)	
func NewFontInfoCopy(value *QFontInfo) *QFontInfo {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),39000,39103,Native(value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontInfo{}
	_p.SetDriver(__rv,39000,true)
	return _p
} 
//QFontInfo::bold()
func (q *QFontInfo) Bold() bool {
	var __rv bool
	q.Drv(39000,39104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::exactMatch()
func (q *QFontInfo) ExactMatch() bool {
	var __rv bool
	q.Drv(39000,39105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::family()
func (q *QFontInfo) Family() string {
	var __rv string
	q.Drv(39000,39106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::fixedPitch()
func (q *QFontInfo) FixedPitch() bool {
	var __rv bool
	q.Drv(39000,39107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::italic()
func (q *QFontInfo) Italic() bool {
	var __rv bool
	q.Drv(39000,39108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::overline()
func (q *QFontInfo) Overline() bool {
	var __rv bool
	q.Drv(39000,39109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::pixelSize()
func (q *QFontInfo) PixelSize() int {
	var __rv int
	q.Drv(39000,39110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::pointSize()
func (q *QFontInfo) PointSize() int {
	var __rv int
	q.Drv(39000,39111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::pointSizeF()
func (q *QFontInfo) PointSizeF() float64 {
	var __rv float64
	q.Drv(39000,39112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::rawMode()
func (q *QFontInfo) RawMode() bool {
	var __rv bool
	q.Drv(39000,39113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::strikeOut()
func (q *QFontInfo) StrikeOut() bool {
	var __rv bool
	q.Drv(39000,39114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::style()
func (q *QFontInfo) Style() QFont_Style {
	var __rv QFont_Style
	q.Drv(39000,39115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::styleHint()
func (q *QFontInfo) StyleHint() QFont_StyleHint {
	var __rv QFont_StyleHint
	q.Drv(39000,39116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::underline()
func (q *QFontInfo) Underline() bool {
	var __rv bool
	q.Drv(39000,39117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontInfo::weight()
func (q *QFontInfo) Weight() int {
	var __rv int
	q.Drv(39000,39118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
