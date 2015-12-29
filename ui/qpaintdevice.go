// Copyright 2015-2016 visualfc <visualfc@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"unsafe"
)	

//enum QPaintDevice_PaintDeviceMetric - QPaintDevice::PaintDeviceMetric
type QPaintDevice_PaintDeviceMetric uint32
const (
	QPaintDevice_PdmWidth QPaintDevice_PaintDeviceMetric = 1
	QPaintDevice_PdmHeight QPaintDevice_PaintDeviceMetric = 1+1
	QPaintDevice_PdmWidthMM QPaintDevice_PaintDeviceMetric = 1+1+1
	QPaintDevice_PdmHeightMM QPaintDevice_PaintDeviceMetric = 1+1+1+1
	QPaintDevice_PdmNumColors QPaintDevice_PaintDeviceMetric = 1+1+1+1+1
	QPaintDevice_PdmDepth QPaintDevice_PaintDeviceMetric = 1+1+1+1+1+1
	QPaintDevice_PdmDpiX QPaintDevice_PaintDeviceMetric = 1+1+1+1+1+1+1
	QPaintDevice_PdmDpiY QPaintDevice_PaintDeviceMetric = 1+1+1+1+1+1+1+1
	QPaintDevice_PdmPhysicalDpiX QPaintDevice_PaintDeviceMetric = 1+1+1+1+1+1+1+1+1
	QPaintDevice_PdmPhysicalDpiY QPaintDevice_PaintDeviceMetric = 1+1+1+1+1+1+1+1+1+1
)
//struct QPaintDevice : QPaintDevice
type QPaintDevice struct {
	BaseDrv
}
//QPaintDevice::colorCount()
func (q *QPaintDevice) ColorCount() int {
	var __rv int
	q.Drv(82000,82102,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::depth()
func (q *QPaintDevice) Depth() int {
	var __rv int
	q.Drv(82000,82103,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::devType()
func (q *QPaintDevice) DevType() int {
	var __rv int
	q.Drv(82000,82104,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::height()
func (q *QPaintDevice) Height() int {
	var __rv int
	q.Drv(82000,82105,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::heightMM()
func (q *QPaintDevice) HeightMM() int {
	var __rv int
	q.Drv(82000,82106,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::logicalDpiX()
func (q *QPaintDevice) LogicalDpiX() int {
	var __rv int
	q.Drv(82000,82107,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::logicalDpiY()
func (q *QPaintDevice) LogicalDpiY() int {
	var __rv int
	q.Drv(82000,82108,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::paintEngine()
func (q *QPaintDevice) PaintEngine() *QPaintEngine {
	var __rv uintptr
	q.Drv(82000,82109,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	if __rv == 0 {
		return nil
	}
	_rp := &QPaintEngine{}
	_rp.SetDriver(__rv,83000,true)
	return _rp
}	
//QPaintDevice::paintingActive()
func (q *QPaintDevice) PaintingActive() bool {
	var __rv bool
	q.Drv(82000,82110,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::physicalDpiX()
func (q *QPaintDevice) PhysicalDpiX() int {
	var __rv int
	q.Drv(82000,82111,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::physicalDpiY()
func (q *QPaintDevice) PhysicalDpiY() int {
	var __rv int
	q.Drv(82000,82112,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::width()
func (q *QPaintDevice) Width() int {
	var __rv int
	q.Drv(82000,82113,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
//QPaintDevice::widthMM()
func (q *QPaintDevice) WidthMM() int {
	var __rv int
	q.Drv(82000,82114,unsafe.Pointer(&__rv),nil,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil)
	return __rv
}	
