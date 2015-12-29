// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QTextBoundaryFinder_BoundaryReason - QTextBoundaryFinder::BoundaryReason
type QTextBoundaryFinder_BoundaryReason uint32
const (
	QTextBoundaryFinder_NotAtBoundary QTextBoundaryFinder_BoundaryReason = 0
	QTextBoundaryFinder_StartWord QTextBoundaryFinder_BoundaryReason = 1
	QTextBoundaryFinder_EndWord QTextBoundaryFinder_BoundaryReason = 2
)
//enum QTextBoundaryFinder_BoundaryType - QTextBoundaryFinder::BoundaryType
type QTextBoundaryFinder_BoundaryType uint32
const (
	QTextBoundaryFinder_Grapheme QTextBoundaryFinder_BoundaryType = 0
	QTextBoundaryFinder_Word QTextBoundaryFinder_BoundaryType = 1
	QTextBoundaryFinder_Line QTextBoundaryFinder_BoundaryType = 2
	QTextBoundaryFinder_Sentence QTextBoundaryFinder_BoundaryType = 3
)
//struct QTextBoundaryFinder : QTextBoundaryFinder
type QTextBoundaryFinder struct {
	BaseDrv
}
//QTextBoundaryFinder::QTextBoundaryFinder()	
func NewTextBoundaryFinder() *QTextBoundaryFinder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),141000,141102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBoundaryFinder{}
	_p.SetDriver(__rv,141000,true)
	return _p
} 
//QTextBoundaryFinder::QTextBoundaryFinder(QTextBoundaryFinder const&)	
func NewTextBoundaryFinderCopy(other *QTextBoundaryFinder) *QTextBoundaryFinder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),141000,141103,Native(other),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBoundaryFinder{}
	_p.SetDriver(__rv,141000,true)
	return _p
} 
//QTextBoundaryFinder::QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType,QString const&)	
func NewTextBoundaryFinderWithTypeString(_type QTextBoundaryFinder_BoundaryType,string string) *QTextBoundaryFinder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),141000,141104,unsafe.Pointer(&_type),unsafe.Pointer(&string),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBoundaryFinder{}
	_p.SetDriver(__rv,141000,true)
	return _p
} 
//QTextBoundaryFinder::QTextBoundaryFinder(QTextBoundaryFinder::BoundaryType,QChar const*,int,unsigned char*,int)	
func NewTextBoundaryFinderWithTypeCharsLengthBufferBuffersize(_type QTextBoundaryFinder_BoundaryType,chars *rune,length int,buffer *byte,bufferSize int) *QTextBoundaryFinder {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),141000,141105,unsafe.Pointer(&_type),unsafe.Pointer(&chars),unsafe.Pointer(&length),unsafe.Pointer(&buffer),unsafe.Pointer(&bufferSize),nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QTextBoundaryFinder{}
	_p.SetDriver(__rv,141000,true)
	return _p
} 
//QTextBoundaryFinder::boundaryReasons()
func (q *QTextBoundaryFinder) BoundaryReasons() QTextBoundaryFinder_BoundaryReason {
	var __rv QTextBoundaryFinder_BoundaryReason
	q.Drv(141000,141106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::isAtBoundary()
func (q *QTextBoundaryFinder) IsAtBoundary() bool {
	var __rv bool
	q.Drv(141000,141107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::isValid()
func (q *QTextBoundaryFinder) IsValid() bool {
	var __rv bool
	q.Drv(141000,141108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::position()
func (q *QTextBoundaryFinder) Position() int {
	var __rv int
	q.Drv(141000,141109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::setPosition(int)
func (q *QTextBoundaryFinder) SetPosition(position int)  {
	q.Drv(141000,141110,unsafe.Pointer(&position),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBoundaryFinder::string()
func (q *QTextBoundaryFinder) String() string {
	var __rv string
	q.Drv(141000,141111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::toEnd()
func (q *QTextBoundaryFinder) ToEnd()  {
	q.Drv(141000,141112,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBoundaryFinder::toNextBoundary()
func (q *QTextBoundaryFinder) ToNextBoundary() int {
	var __rv int
	q.Drv(141000,141113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::toPreviousBoundary()
func (q *QTextBoundaryFinder) ToPreviousBoundary() int {
	var __rv int
	q.Drv(141000,141114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QTextBoundaryFinder::toStart()
func (q *QTextBoundaryFinder) ToStart()  {
	q.Drv(141000,141115,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QTextBoundaryFinder::type()
func (q *QTextBoundaryFinder) Type() QTextBoundaryFinder_BoundaryType {
	var __rv QTextBoundaryFinder_BoundaryType
	q.Drv(141000,141116,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
