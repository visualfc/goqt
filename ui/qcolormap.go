// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QColormap_Mode - QColormap::Mode
type QColormap_Mode uint32
const (
	QColormap_Direct QColormap_Mode = 0
	QColormap_Indexed QColormap_Mode = 1
	QColormap_Gray QColormap_Mode = 2
)
//struct QColormap : QColormap
type QColormap struct {
	BaseDrv
}
//QColormap::QColormap(QColormap const&)	
func NewColormap(colormap *QColormap) *QColormap {
	var __rv uintptr
	err := DirectQtDrv(unsafe.Pointer(&__rv),14000,14102,Native(colormap),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if err != nil || __rv == 0 {
		return nil
	}
	_p := &QColormap{}
	_p.SetDriver(__rv,14000,true)
	return _p
} 
//QColormap::cleanup()	
func QColormapCleanup()  {
	DirectQtDrv(nil,14000,14103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColormap::cleanup()
func (q *QColormap) Cleanup()  {
	q.Drv(14000,14103,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColormap::colorAt(unsigned int)
func (q *QColormap) ColorAt(pixel uint) *QColor {
	var __rv uintptr
	q.Drv(14000,14104,unsafe.Pointer(&pixel),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColor{}
	_rp.SetDriver(__rv,13000,true)
	return _rp
}	
//QColormap::colormap()
func (q *QColormap) Colormap() []*QColor {
	var __rv []*QColor
	q.Drv(14000,14105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColormap::depth()
func (q *QColormap) Depth() int {
	var __rv int
	q.Drv(14000,14106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColormap::initialize()	
func QColormapInitialize()  {
	DirectQtDrv(nil,14000,14107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColormap::initialize()
func (q *QColormap) Initialize()  {
	q.Drv(14000,14107,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
}	
//QColormap::instance()	
func QColormapInstance() *QColormap {
	var __rv uintptr
	DirectQtDrv(nil,14000,14108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColormap{}
	_rp.SetDriver(__rv,14000,true)
	return _rp
}	
//QColormap::instance()
func (q *QColormap) Instance() *QColormap {
	var __rv uintptr
	q.Drv(14000,14108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColormap{}
	_rp.SetDriver(__rv,14000,true)
	return _rp
}	
//QColormap::instance(int)	
func QColormapInstanceWithScreen(screen int) *QColormap {
	var __rv uintptr
	DirectQtDrv(nil,14000,14109,unsafe.Pointer(&screen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColormap{}
	_rp.SetDriver(__rv,14000,true)
	return _rp
}	
//QColormap::instance(int)
func (q *QColormap) InstanceWithScreen(screen int) *QColormap {
	var __rv uintptr
	q.Drv(14000,14109,unsafe.Pointer(&screen),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QColormap{}
	_rp.SetDriver(__rv,14000,true)
	return _rp
}	
//QColormap::mode()
func (q *QColormap) Mode() QColormap_Mode {
	var __rv QColormap_Mode
	q.Drv(14000,14110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColormap::pixel(QColor const&)
func (q *QColormap) Pixel(color *QColor) uint {
	var __rv uint
	q.Drv(14000,14111,Native(color),unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QColormap::size()
func (q *QColormap) Size() int {
	var __rv int
	q.Drv(14000,14112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
