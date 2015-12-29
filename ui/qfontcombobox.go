// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QFontComboBox_FontFilter - QFontComboBox::FontFilter
type QFontComboBox_FontFilter uint32
const (
	QFontComboBox_AllFonts QFontComboBox_FontFilter = 0
	QFontComboBox_ScalableFonts QFontComboBox_FontFilter = 0x1
	QFontComboBox_NonScalableFonts QFontComboBox_FontFilter = 0x2
	QFontComboBox_MonospacedFonts QFontComboBox_FontFilter = 0x4
	QFontComboBox_ProportionalFonts QFontComboBox_FontFilter = 0x8
)
//struct QFontComboBox : QFontComboBox
type QFontComboBox struct {
	QComboBox
}
func (q *QFontComboBox) OnCurrentFontChanged(fn func(*QFont)) uintptr {
	var __rv uintptr
	q.Drv(242000,242102,unsafe.Pointer(drvNewIfaceRef(fn)),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)	
	signalMap[__rv] = fn
	return __rv
}
//QFontComboBox::QFontComboBox()	
func NewFontComboBox() *QFontComboBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),242000,242103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontComboBox{}
	_p.SetDriver(__rv,242000,false)
	return _p
} 
//QFontComboBox::QFontComboBox(QWidget*)	
func NewFontComboBoxWithParent(parent QWidgetInterface) *QFontComboBox {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),242000,242104,Native(parent),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QFontComboBox{}
	_p.SetDriver(__rv,242000,false)
	return _p
} 
//QFontComboBox::currentFont()
func (q *QFontComboBox) CurrentFont() *QFont {
	var __rv uintptr
	q.Drv(242000,242105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QFont{}
	_rp.SetDriver(__rv,37000,true)
	return _rp
}	
//QFontComboBox::event(QEvent*)
func (q *QFontComboBox) Event(e *QEvent) bool {
	var __rv bool
	q.Drv(242000,242106,Native(e),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontComboBox::fontFilters()
func (q *QFontComboBox) FontFilters() QFontComboBox_FontFilter {
	var __rv QFontComboBox_FontFilter
	q.Drv(242000,242107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QFontComboBox::setCurrentFont(QFont const&)
func (q *QFontComboBox) SetCurrentFont(f *QFont)  {
	q.Drv(242000,242108,Native(f),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontComboBox::setFontFilters(QFlags<QFontComboBox::FontFilter>)
func (q *QFontComboBox) SetFontFilters(filters QFontComboBox_FontFilter)  {
	q.Drv(242000,242109,unsafe.Pointer(&filters),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontComboBox::setWritingSystem(QFontDatabase::WritingSystem)
func (q *QFontComboBox) SetWritingSystem(value QFontDatabase_WritingSystem)  {
	q.Drv(242000,242110,unsafe.Pointer(&value),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QFontComboBox::sizeHint()
func (q *QFontComboBox) SizeHint() *QSize {
	var __rv uintptr
	q.Drv(242000,242111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QSize{}
	_rp.SetDriver(__rv,119000,true)
	return _rp
}	
//QFontComboBox::writingSystem()
func (q *QFontComboBox) WritingSystem() QFontDatabase_WritingSystem {
	var __rv QFontDatabase_WritingSystem
	q.Drv(242000,242112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
