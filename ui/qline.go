// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//struct QLine : QLine
type QLine struct {
	BaseDrv
}
//QLine::QLine()	
func NewLine() *QLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),68000,68102,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLine{}
	_p.SetDriver(__rv,68000,true)
	return _p
} 
//QLine::QLine(QPoint const&,QPoint const&)	
func NewLineWithPt1Pt2(pt1 *QPoint,pt2 *QPoint) *QLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),68000,68103,Native(pt1),Native(pt2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLine{}
	_p.SetDriver(__rv,68000,true)
	return _p
} 
//QLine::QLine(int,int,int,int)	
func NewLineWithX1Y1X2Y2(x1 int,y1 int,x2 int,y2 int) *QLine {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),68000,68104,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QLine{}
	_p.SetDriver(__rv,68000,true)
	return _p
} 
//QLine::dx()
func (q *QLine) Dx() int {
	var __rv int
	q.Drv(68000,68105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::dy()
func (q *QLine) Dy() int {
	var __rv int
	q.Drv(68000,68106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::isNull()
func (q *QLine) IsNull() bool {
	var __rv bool
	q.Drv(68000,68107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::p1()
func (q *QLine) P1() *QPoint {
	var __rv uintptr
	q.Drv(68000,68108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QLine::p2()
func (q *QLine) P2() *QPoint {
	var __rv uintptr
	q.Drv(68000,68109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPoint{}
	_rp.SetDriver(__rv,99000,true)
	return _rp
}	
//QLine::setLine(int,int,int,int)
func (q *QLine) SetLine(x1 int,y1 int,x2 int,y2 int)  {
	q.Drv(68000,68110,unsafe.Pointer(&x1),unsafe.Pointer(&y1),unsafe.Pointer(&x2),unsafe.Pointer(&y2),nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::setP1(QPoint const&)
func (q *QLine) SetP1(p1 *QPoint)  {
	q.Drv(68000,68111,Native(p1),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::setP2(QPoint const&)
func (q *QLine) SetP2(p2 *QPoint)  {
	q.Drv(68000,68112,Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::setPoints(QPoint const&,QPoint const&)
func (q *QLine) SetPoints(p1 *QPoint,p2 *QPoint)  {
	q.Drv(68000,68113,Native(p1),Native(p2),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::translate(QPoint const&)
func (q *QLine) Translate(p *QPoint)  {
	q.Drv(68000,68114,Native(p),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::translate(int,int)
func (q *QLine) TranslateWithDxDy(dx int,dy int)  {
	q.Drv(68000,68115,unsafe.Pointer(&dx),unsafe.Pointer(&dy),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QLine::translated(QPoint const&)
func (q *QLine) Translated(p *QPoint) *QLine {
	var __rv uintptr
	q.Drv(68000,68116,Native(p),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLine{}
	_rp.SetDriver(__rv,68000,true)
	return _rp
}	
//QLine::translated(int,int)
func (q *QLine) TranslatedWithDxDy(dx int,dy int) *QLine {
	var __rv uintptr
	q.Drv(68000,68117,unsafe.Pointer(&dx),unsafe.Pointer(&dy),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QLine{}
	_rp.SetDriver(__rv,68000,true)
	return _rp
}	
//QLine::x1()
func (q *QLine) X1() int {
	var __rv int
	q.Drv(68000,68118,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::x2()
func (q *QLine) X2() int {
	var __rv int
	q.Drv(68000,68119,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::y1()
func (q *QLine) Y1() int {
	var __rv int
	q.Drv(68000,68120,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QLine::y2()
func (q *QLine) Y2() int {
	var __rv int
	q.Drv(68000,68121,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
