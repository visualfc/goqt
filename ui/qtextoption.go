// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextOption_TabType - QTextOption::TabType
type QTextOption_TabType uint32
const (
	QTextOption_LeftTab QTextOption_TabType = 0
	QTextOption_RightTab QTextOption_TabType = 1
	QTextOption_CenterTab QTextOption_TabType = 2
	QTextOption_DelimiterTab QTextOption_TabType = 3
)
//enum QTextOption_Flag - QTextOption::Flag
type QTextOption_Flag uint32
const (
	QTextOption_ShowTabsAndSpaces QTextOption_Flag = 0x1
	QTextOption_ShowLineAndParagraphSeparators QTextOption_Flag = 0x2
	QTextOption_AddSpaceForLineAndParagraphSeparators QTextOption_Flag = 0x4
	QTextOption_SuppressColors QTextOption_Flag = 0x8
	QTextOption_IncludeTrailingSpaces QTextOption_Flag = 0x80000000
)
//enum QTextOption_WrapMode - QTextOption::WrapMode
type QTextOption_WrapMode uint32
const (
	QTextOption_NoWrap QTextOption_WrapMode = 0
	QTextOption_WordWrap QTextOption_WrapMode = 1
	QTextOption_ManualWrap QTextOption_WrapMode = 2
	QTextOption_WrapAnywhere QTextOption_WrapMode = 3
	QTextOption_WrapAtWordBoundaryOrAnywhere QTextOption_WrapMode = 4
)
//struct QTextOption : QTextOption
type QTextOption struct {
	BaseDrv
}
//QTextOption::QTextOption()	
func NewTextOption() *QTextOption {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),164000,164102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextOption{}
	_p.SetDriver(__rv,164000,true)
	return _p
} 
//QTextOption::QTextOption(QFlags<Qt::AlignmentFlag>)	
func NewTextOptionWithAlignment(alignment Qt_AlignmentFlag) *QTextOption {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),164000,164103,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextOption{}
	_p.SetDriver(__rv,164000,true)
	return _p
} 
//QTextOption::QTextOption(QTextOption const&)	
func NewTextOptionCopy(o *QTextOption) *QTextOption {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),164000,164104,Native(o),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextOption{}
	_p.SetDriver(__rv,164000,true)
	return _p
} 
//QTextOption::alignment()
func (q *QTextOption) Alignment() Qt_AlignmentFlag {
	var __rv Qt_AlignmentFlag
	q.Drv(164000,164105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::flags()
func (q *QTextOption) Flags() QTextOption_Flag {
	var __rv QTextOption_Flag
	q.Drv(164000,164106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::setAlignment(QFlags<Qt::AlignmentFlag>)
func (q *QTextOption) SetAlignment(alignment Qt_AlignmentFlag)  {
	q.Drv(164000,164107,unsafe.Pointer(&alignment),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setFlags(QFlags<QTextOption::Flag>)
func (q *QTextOption) SetFlags(flags QTextOption_Flag)  {
	q.Drv(164000,164108,unsafe.Pointer(&flags),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setTabArray(QList<qreal>)
func (q *QTextOption) SetTabArray(tabStops []float64)  {
	q.Drv(164000,164109,unsafe.Pointer(&tabStops),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setTabStop(double)
func (q *QTextOption) SetTabStop(tabStop float64)  {
	q.Drv(164000,164110,unsafe.Pointer(&tabStop),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setTabs(QList<QTextOption::Tab>)
func (q *QTextOption) SetTabs(tabStops []*QTextOptionTab)  {
	q.Drv(164000,164111,unsafe.Pointer(&tabStops),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setTextDirection(Qt::LayoutDirection)
func (q *QTextOption) SetTextDirection(aDirection Qt_LayoutDirection)  {
	q.Drv(164000,164112,unsafe.Pointer(&aDirection),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setUseDesignMetrics(bool)
func (q *QTextOption) SetUseDesignMetrics(b bool)  {
	q.Drv(164000,164113,unsafe.Pointer(&b),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::setWrapMode(QTextOption::WrapMode)
func (q *QTextOption) SetWrapMode(wrap QTextOption_WrapMode)  {
	q.Drv(164000,164114,unsafe.Pointer(&wrap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextOption::tabArray()
func (q *QTextOption) TabArray() []float64 {
	var __rv []float64
	q.Drv(164000,164115,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::tabStop()
func (q *QTextOption) TabStop() float64 {
	var __rv float64
	q.Drv(164000,164116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::tabs()
func (q *QTextOption) Tabs() []*QTextOptionTab {
	var __rv []*QTextOptionTab
	q.Drv(164000,164117,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::textDirection()
func (q *QTextOption) TextDirection() Qt_LayoutDirection {
	var __rv Qt_LayoutDirection
	q.Drv(164000,164118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::useDesignMetrics()
func (q *QTextOption) UseDesignMetrics() bool {
	var __rv bool
	q.Drv(164000,164119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextOption::wrapMode()
func (q *QTextOption) WrapMode() QTextOption_WrapMode {
	var __rv QTextOption_WrapMode
	q.Drv(164000,164120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
